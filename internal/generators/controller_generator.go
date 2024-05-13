package generators

import (
	"bytes"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/k0kubun/pp/v3"
	"os"
	"strings"
	"text/template"
)

type GenerateControllerContext struct {
	EntityName           string
	RelationshipsComment string
}

type ControllerGenerator struct {
	options   GenerateControllerContext
	logger    *logger.AppLogger
	pluralize *pluralize.Client
}

func NewControllerGenerator(options GenerateControllerContext, logger *logger.AppLogger) *ControllerGenerator {
	return &ControllerGenerator{
		options:   options,
		logger:    logger,
		pluralize: pluralize.NewClient(),
	}
}

type templateData struct {
	RelationshipsComment  string
	KeyNameModelField     string
	KeyName               string
	KeyColumn             string
	KeyNameLowerCamel     string
	EntityName            string
	EntityNamePlural      string
	EntityNameSnake       string
	EntityNameSnakePlural string
	EntityNameCamel       string
	EntityNameCamelPlural string
	ParamLine             string
	QueryParams           string
}

const (
	crudControllerPath = "./internal/http/crudcontrollers/"
)

func (gc *ControllerGenerator) Generate() {
	keyName := ""
	keys := GetDbSchemaKeysConfigTable(gc.options.EntityName)

	if len(os.Getenv("DEBUG")) > 0 {
		pp.Println("# Keys")
		pp.Println(keys)
	}

	priKey := DbSchemaRowResult{}

	// first pass grab id if it exists
	for _, key := range keys {
		if key.Column == "id" {
			keyName = strcase.ToCamel(key.Column)
			priKey = key
			break
		}
	}

	// second pass if not found
	if len(keyName) == 0 {
		for _, key := range keys {
			if key.ColumnKey.String == "PRI" {
				keyName = strcase.ToCamel(key.Column)
				priKey = key
				break
			}

			if key.OrdinalPosition == "1" {
				keyName = strcase.ToCamel(key.Column)
				priKey = key
				break
			}
		}
	}

	newKeyName := keyName
	// gorm uses capital "ID"
	if keyName == "Id" {
		keyName = "ID"
	}

	// type is a reserved word
	if keyName == "Type" {
		newKeyName = "TypeId"
	}

	if len(os.Getenv("DEBUG")) > 0 {
		pp.Println("keyName")
		pp.Println(keyName)
		pp.Println(priKey)
	}

	// build primary key param line
	paramLine := ""
	if strings.Contains(priKey.DataType, "int") {
		paramLine = fmt.Sprintf(
			"%s, err := strconv.Atoi(c.Param(\"%s\"))",
			strcase.ToLowerCamel(newKeyName),
			strcase.ToLowerCamel(newKeyName),
		)
	}

	queryParams := ""
	// loop through secondary keys (skip first)
	for i, key := range keys {
		if i != 0 {
			if key.Column == "type" {
				key.Column = "typeId"
			}

			if len(os.Getenv("DEBUG")) > 0 {
				pp.Println("key.Column")
				pp.Println(key.Column)
			}

			// add type lines (uint / int etc.)
			param := fmt.Sprintf(`
	// key param [%s] position [%s] type [%s]
	if len(c.QueryParam("%s")) > 0 {
		%sParam, err := strconv.Atoi(c.QueryParam("%s"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [%s] err [%%s]", err.Error())})
		}

		params = append(params, %sParam)
		keys = append(keys, "%s = ?")
	}
`,
				key.Column,
				key.OrdinalPosition,
				key.DataType,
				key.Column,
				strcase.ToLowerCamel(key.Column),
				key.Column,
				key.Column,
				strcase.ToLowerCamel(key.Column),
				key.Column,
			)

			queryParams += param
			if len(os.Getenv("DEBUG")) > 0 {
				fmt.Println(param)
			}
		}
	}

	entityName := gc.pluralize.Singular(gc.options.EntityName)
	tpl, err := template.ParseFiles("./internal/generators/templates/crud_controller.tmpl")
	templateData := templateData{
		RelationshipsComment:  gc.options.RelationshipsComment,
		EntityName:            strcase.ToCamel(entityName),
		KeyNameModelField:     keyName,
		KeyName:               newKeyName,
		KeyNameLowerCamel:     strcase.ToLowerCamel(newKeyName),
		KeyColumn:             priKey.Column,
		EntityNamePlural:      gc.pluralize.Plural(strcase.ToCamel(entityName)),
		EntityNameSnake:       strcase.ToSnake(entityName),
		EntityNameSnakePlural: gc.pluralize.Plural(strcase.ToSnake(entityName)),
		EntityNameCamel:       strcase.ToLowerCamel(entityName),
		EntityNameCamelPlural: gc.pluralize.Plural(strcase.ToLowerCamel(entityName)),
		ParamLine:             paramLine,
		QueryParams:           queryParams,
	}

	if len(os.Getenv("DEBUG")) > 0 {
		pp.Println("# Template Data")
		pp.Println(templateData)
	}

	var out bytes.Buffer
	err = tpl.ExecuteTemplate(&out, "crud_controller.tmpl", templateData)
	if err != nil {
		gc.logger.Fatal().Err(err).Msg("Failed to execute template")
	}

	// write file
	fileName := fmt.Sprintf("%v%v_controller.go", crudControllerPath, strcase.ToSnake(entityName))

	// create file
	file, err := os.Create(fileName)
	if err != nil {
		gc.logger.Fatal().Err(err).Msg("Failed to create file")
	}

	defer file.Close()

	// write contents
	_, err = file.WriteString(out.String())
	if err != nil {
		gc.logger.Fatal().Err(err).Msg("Failed to write file")
	}

	fmt.Println(fmt.Sprintf("Generated [%v] from Entity [%v]\n", fileName, templateData.EntityName))
}
