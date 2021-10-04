package generators

import (
	"bytes"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"os"
	"text/template"
)

type GenerateControllerContext struct {
	EntityName           string
	RelationshipsComment string
}

type GenerateController struct {
	options GenerateControllerContext
	logger  *logrus.Logger
	pluralize *pluralize.Client
}

func NewGenerateController(options GenerateControllerContext, logger *logrus.Logger) *GenerateController {
	return &GenerateController{
		options:   options,
		logger:    logger,
		pluralize: pluralize.NewClient(),
	}
}

type templateData struct {
	RelationshipsComment  string
	KeyName               string
	EntityName            string
	EntityNamePlural      string
	EntityNameSnake       string
	EntityNameSnakePlural string
	EntityNameCamel       string
	EntityNameCamelPlural string
}

const (
	crudControllerPath = "./internal/http/crudcontrollers/"
)

func (gc *GenerateController) Generate() {
	keyName := ""
	keys := GetDbSchemaKeysConfigTable(gc.options.EntityName)

	// first pass grab id if it exists
	for _, key := range keys {
		if key.Column == "id" {
			keyName = strcase.ToCamel(key.Column)
			break
		}
	}

	// second pass if not found
	if len(keyName) == 0 {
		for _, key := range keys {
			fmt.Println(key)
			if key.ColumnKey.String == "PRI" {
				keyName = strcase.ToCamel(key.Column)
				break
			}

			if key.OrdinalPosition == "1" {
				keyName = strcase.ToCamel(key.Column)
				break
			}
		}
	}

	if keyName == "Id" {
		keyName = "ID"
	}

	entityName := gc.pluralize.Singular(gc.options.EntityName)

	tpl, err := template.ParseFiles("./internal/generators/templates/crud_controller.tmpl")
	templateData := templateData{
		RelationshipsComment:  gc.options.RelationshipsComment,
		EntityName:            strcase.ToCamel(entityName),
		KeyName:               keyName,
		EntityNamePlural:      gc.pluralize.Plural(strcase.ToCamel(entityName)),
		EntityNameSnake:       strcase.ToSnake(entityName),
		EntityNameSnakePlural: gc.pluralize.Plural(strcase.ToSnake(entityName)),
		EntityNameCamel:       strcase.ToLowerCamel(entityName),
		EntityNameCamelPlural: gc.pluralize.Plural(strcase.ToLowerCamel(entityName)),
	}

	var out bytes.Buffer
	err = tpl.ExecuteTemplate(&out, "crud_controller.tmpl", templateData)
	if err != nil {
		fmt.Println(err)
	}

	// write file
	fileName := fmt.Sprintf("%v%v_controller.go", crudControllerPath, strcase.ToSnake(entityName))

	// create file
	file, err := os.Create(fileName)
	if err != nil {
		gc.logger.Fatal(err)
	}

	defer file.Close()

	// write contents
	_, err = file.WriteString(out.String())
	if err != nil {
		gc.logger.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Generated [%v] from Entity [%v]\n", fileName, templateData.EntityName))
}
