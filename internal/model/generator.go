package model

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type Generator struct {
	// deps
	logger       *logger.AppLogger
	gorm         *gorm.DB
	pluralize    *pluralize.Client
	schemaLookup *DbLookup // schema lookup

	// options
	withControllers bool // generate withControllers

	// member variable
	models        []string
	relationships []ForeignKeyMappings
}

// NewGenerator creates a new generator
func NewGenerator(logger *logger.AppLogger, gorm *gorm.DB) *Generator {
	db, _ := gorm.DB()
	return &Generator{
		logger:       logger,
		gorm:         gorm,
		schemaLookup: NewDbLookup(db, logger),
		pluralize:    pluralize.NewClient(),
		models:       make([]string, 0),
	}
}

const (
	RelationshipType1to1    = "1-1" // RelationshipType1to1 1 to 1
	RelationshipType1toMany = "1-*" // RelationshipType1toMany 1 to many
	RelationshipTypeManyTo1 = "*-1" // RelationshipTypeManyTo1 many to 1
)

// ForeignKeyMappings is a struct that represents a foreign key mapping
type ForeignKeyMappings struct {
	Table        string `json:"table"`         // local Table
	Key          string `json:"key"`           // local Key
	RemoteTable  string `json:"remote_table"`  // remote Table
	RemoteKey    string `json:"remote_key"`    // remote Key
	RelationType string `json:"relation_type"` // relationship type
}

// Relationships is a struct that represents a model and its relationships
type Relationships struct {
	Table         string   `json:"table"`
	ModelName     string   `json:"model_name"`
	Relationships []string `json:"relationships"`
}

// Generate generates the models for the given tables
// if no tables are provided, it will generate models for all tables
// it will also generate a models.go file in the models directory
// the models.go file will contain a list of models that implement the Modelable interface
// and a list of model names
func (g *Generator) Generate(tables []string) {
	g.relationships = g.loadRelationships()
	tableNames := g.getTableNames()

	allTables := len(tables) == 0

	tables, err := g.resolveTablesToGenerate(tables)
	if err != nil {
		return // or handle accordingly
	}

	for _, genModel := range tables {
		for _, table := range tableNames {
			if genModel != "all" && table != genModel {
				continue
			}

			err := g.generateModel(table)
			if err != nil {
				g.logger.Error().Err(err).Msgf("error generating model for table %s", table)
				continue
			}

			if g.withControllers {
				g.MakeController(table)
			}
		}
	}

	if allTables {
		err = g.generateModelsFile()
		if err != nil {
			g.logger.Error().Err(err).Msg("error generating models file")
		}
	}
}

// return relationship type model prefix
func (g *Generator) getRelationshipTypeModelAttributePrefix(r ForeignKeyMappings) string {
	if r.RelationType == RelationshipType1to1 {
		return "*"
	}
	if r.RelationType == RelationshipType1toMany {
		return "[]"
	}

	return ""
}

// check if element exists in slice
func exists(a []string, element string) bool {
	for _, e := range a {
		if e == element {
			return true
		}
	}

	return false
}

// getNestedRelations returns a list of relationships for a table
// it will recursively check for relationships in the table
// it will skip relationships that are already in the parent tables
func (g *Generator) getNestedRelations(table string, prefix string, parentTables []string, level int) []string {
	relationshipNames := make([]string, 0)

	if prefix != "" {
		prefix = fmt.Sprintf("%v.", prefix)
	}

	g.logger.Debug().Msgf("-- [%v] table [%v]", level, table)

	if exists(parentTables, table) && level > 0 {
		return relationshipNames
	}

	parentTables = append(parentTables, table)
	currentLevel := level + 1
	for _, relation := range g.relationships {
		if table != relation.Table {
			continue
		}

		g.logger.Debug().Msgf("-- [%v] table [%v] relation [%v] remote [%v]", level, table, relation, relation.RemoteTable)

		if len(parentTables) > 0 {
			for _, parentTable := range parentTables {
				if relation.RemoteTable == parentTable {
					g.logger.Debug().Msgf(
						"---- [%v] remote table [%v] is a parent table [%v] skipping",
						level,
						relation.RemoteTable,
						parentTable,
					)
					continue
				}
			}
		}

		relationshipAttributeName := ""

		switch relation.RelationType {
		case RelationshipType1to1:
			relationshipAttributeName = strcase.ToCamel(g.pluralize.Singular(relation.RemoteTable))
		case RelationshipType1toMany:
			relationshipAttributeName = strcase.ToCamel(g.pluralize.Plural(relation.RemoteTable))
		case RelationshipTypeManyTo1:
			// todo: inverse
		}

		relationshipNames = append(relationshipNames, fmt.Sprintf("%v%v", prefix, relationshipAttributeName))
		passedDownPrefix := fmt.Sprintf("%v%v", prefix, relationshipAttributeName)
		relationshipNames = append(
			relationshipNames,
			g.getNestedRelations(relation.RemoteTable, passedDownPrefix, parentTables, currentLevel)...,
		)
	}

	sort.Strings(relationshipNames)

	return relationshipNames
}

// return Table names from database
func (g *Generator) getTableNames() []string {
	tableNames := make([]string, 0)
	g.gorm.Raw("SHOW TABLES").Scan(&tableNames)

	return tableNames
}

const dbRelationshipConfig = ".generate-db-relationships.yml"

// loadRelationships loads the relationships from the yaml file
// and returns a list of ForeignKeyMappings
// the yaml file should be in the format:
// table_name:
//   - relation_type local_key->remote_table:remote_key
func (g *Generator) loadRelationships() []ForeignKeyMappings {
	// load yaml
	databaseSchemaYaml, err := os.ReadFile(dbRelationshipConfig)
	if err != nil {
		g.logger.Fatal().Err(err).Msg("error creating file")
	}

	// unmarshal yaml
	dbRelationships := make(map[string][]string, 0)
	err = yaml.Unmarshal(databaseSchemaYaml, &dbRelationships)
	if err != nil {
		g.logger.Fatal().Err(err).Msg("error unmarshalling yaml")
	}

	relationships := []ForeignKeyMappings{}
	for table, relations := range dbRelationships {
		for _, relationData := range relations {
			// unpack values from config row
			// relation_type local_key->remote_table:remote_key

			// split: relation_type local_key->remote_table:remote_key
			split := strings.Split(relationData, " ")
			relationType := strings.TrimSpace(split[0])
			relationSignature := strings.TrimSpace(split[1])

			if !g.isValidRelationshipType(relationType) {
				g.logger.Info().
					Any("relationType", relationType).
					Any("relationSignature", relationSignature).
					Any("dbRelationshipConfig", dbRelationshipConfig).
					Msg("Invalid relationship type")
			}

			// split: local_key->remote_table:remote_key
			keySplit := strings.Split(relationSignature, "->")
			localKey := strings.TrimSpace(keySplit[0])

			// this will need to handle multiple keys at some point
			remoteSplit := strings.Split(keySplit[1], ":")
			remoteTable := strings.TrimSpace(remoteSplit[0])
			remoteKey := strings.TrimSpace(remoteSplit[1])

			relationships = append(
				relationships, ForeignKeyMappings{
					Table:        table,
					Key:          localKey,
					RemoteTable:  remoteTable,
					RemoteKey:    remoteKey,
					RelationType: relationType,
				},
			)
		}
	}

	return relationships
}

// isValidRelationshipType checks if the relationship type is valid
func (g *Generator) isValidRelationshipType(relationshipType string) bool {
	switch relationshipType {
	case RelationshipType1to1:
		return true
	case RelationshipType1toMany:
		return true
	}

	return false
}

// generateModelsFile generates the models file
// contains a list of models that implement the Modelable interface
// and a list of model names
// outputs a file in ./internal/models/models.go
func (g *Generator) generateModelsFile() error {
	sort.Strings(g.models)

	const modelsFileTemplate = `package models

func GetModels() []Modelable {
	return []Modelable{
		{{- range .Models }}
		&{{ . }}{},
		{{- end }}
	}
}

func GetModelNames() []string {
	return []string{
		{{- range .ModelNames }}
		"{{ . }}",
		{{- end }}
	}
}
`

	tmpl, err := template.New("modelsFile").Parse(modelsFileTemplate)
	if err != nil {
		return err
	}

	data := struct {
		Models     []string
		ModelNames []string
	}{
		Models:     g.models,
		ModelNames: g.models,
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, data); err != nil {
		return err
	}

	err = g.writeToFile("internal/models/models.go", out.String())
	if err != nil {
		return err
	}

	g.logger.Info().Msg("Generated models file")
	return nil
}

// writeToFile is a helper function to write content to a file
func (g *Generator) writeToFile(path, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

// getMaxFieldAndTypeLengths returns the max field and type lengths for a table
func (g *Generator) getMaxFieldAndTypeLengths(table string, defs []DbSchemaRowResult) (int, int) {
	maxFieldLen := 0
	maxTypeLen := 0

	for _, col := range defs {
		if len(col.Column) > maxFieldLen {
			maxFieldLen = len(col.Column)
		}
		translated := translateDataType(col)
		if len(translated) > maxTypeLen {
			maxTypeLen = len(translated)
		}
	}

	for _, relation := range g.relationships {
		if relation.Table != table {
			continue
		}

		col := strcase.ToCamel(g.pluralize.Plural(relation.RemoteTable))
		colType := g.getRelationshipTypeModelAttributePrefix(relation) +
			strcase.ToCamel(g.pluralize.Singular(relation.RemoteTable))

		if len(col) > maxFieldLen {
			maxFieldLen = len(col)
		}
		if len(colType) > maxTypeLen {
			maxTypeLen = len(colType)
		}

		g.logger.Debug().Msgf("-- col [%v] colType [%v]", col, colType)
	}

	return maxFieldLen, maxTypeLen
}

// buildModelFields builds the model fields for a table
// it takes the column definitions and the max field and type lengths
// and returns a string with the model fields and a map of struct field names
// example output:
//
//	ID   int    `json:"id" gorm:"Column:id"`
//	Name string `json:"name" gorm:"Column:name"`
func (g *Generator) buildModelFields(columns []DbSchemaRowResult, maxColLen, maxTypeLen int) string {
	var b strings.Builder

	for _, def := range columns {
		g.logger.Debug().Msgf("-- def [%v] table [%v]", def, def.Column)

		structFieldName := g.getStructFieldName(def.Column)
		jsonFieldName := strcase.ToSnake(def.Column)
		goType := translateDataType(def)

		b.WriteString(fmt.Sprintf(
			"\t%-*s%-*s `json:\"%v\" gorm:\"Column:%v\"`\n",
			maxColLen+1,
			structFieldName,
			maxTypeLen,
			goType,
			jsonFieldName,
			def.Column,
		))

	}

	return b.String()
}

// translateDataType translates the data type from the database to Go
func (g *Generator) getStructFieldName(field string) string {
	if field == "id" {
		return "ID"
	}

	name := strcase.ToCamel(field)
	if strings.HasSuffix(field, "_") {
		name += "2"
	}
	return name
}

// buildRelationModelFields builds the relationship fields for a table
// it takes the table name and the max field and type lengths
// and returns a string with the model fields
// example output:
//
//	*User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
func (g *Generator) buildRelationModelFields(table string, maxColLen int, maxTypeLen int) string {
	var b strings.Builder

	for _, relation := range g.relationships {
		if relation.Table != table {
			continue
		}

		g.logger.Debug().Msgf("-- relationships [%v]", relation)

		prefix := g.getRelationshipTypeModelAttributePrefix(relation)
		remoteSingular := g.pluralize.Singular(relation.RemoteTable)
		remotePlural := g.pluralize.Plural(relation.RemoteTable)

		// Field name and type
		var attributeName, jsonName, foreignKey, referenceKey string

		switch relation.RelationType {
		case RelationshipType1to1:
			attributeName = strcase.ToCamel(remoteSingular)
			jsonName = strcase.ToSnake(remoteSingular)
			foreignKey = relation.Key
			referenceKey = relation.RemoteKey
		case RelationshipType1toMany:
			attributeName = strcase.ToCamel(remotePlural)
			jsonName = strcase.ToSnake(remotePlural)
			foreignKey = relation.RemoteKey
			referenceKey = relation.Key
		default:
			continue
		}

		// Suffix logic if name conflict (inactive but ready)
		// if _, exists := structFieldNames[attributeName]; exists {
		//     attributeName += "Relation"
		//     jsonName += "_relation"
		// }

		fieldType := prefix + strcase.ToCamel(remoteSingular)

		b.WriteString(fmt.Sprintf(
			"\t%-*s%-*s `json:\"%s,omitempty\" gorm:\"foreignKey:%s;references:%s\"`\n",
			maxColLen+1,
			attributeName,
			maxTypeLen,
			fieldType,
			jsonName,
			foreignKey,
			referenceKey,
		))
	}

	return b.String()
}

const (
	BaseGormModelTemplate = `package models
{{imports}}
type {{model_name}} struct {
{{model_fields}}}

func ({{model_name}}) TableName() string {
    return "{{table_name}}"
}
{{relationships}}
{{connection}}
`

	BaseGormModelRelationshipTemplate = `
func ({{model_name}}) Relationships() []string {
    return []string{{{relationships}}}
}`
	BaseGormModelConnectionTemplate = `
func ({{model_name}}) Connection() string {
    return "{{connection_name}}"
}`

	BaseDependencyImportTemplate = `import ({{imports}})`
)

// generateModel generates the model file for a table
func (g *Generator) generateModel(table string) error {
	modelName := g.pluralize.Singular(strcase.ToCamel(table))
	fileName := filepath.Join("./internal/models/", strcase.ToSnake(table)+".go")

	g.logger.Debug().Msgf("-- Generating model for [%s] → [%s]", table, fileName)

	// Template skeleton
	t := BaseGormModelTemplate
	importTemplate := BaseDependencyImportTemplate

	// Get column data
	defs, err := g.schemaLookup.GetSchemasByTableName(table)
	if err != nil {
		g.logger.Error().Err(err).Msgf("error getting schema for table %s", table)
	}
	maxColLen, maxTypeLen := g.getMaxFieldAndTypeLengths(table, defs)

	// Build model content
	modelFields := g.buildModelFields(defs, maxColLen, maxTypeLen)
	modelFields += g.buildRelationModelFields(table, maxColLen, maxTypeLen)

	// Replace model fields
	t = strings.ReplaceAll(t, "{{model_fields}}", modelFields)
	t = strings.ReplaceAll(t, "{{model_name}}", modelName)
	t = strings.ReplaceAll(t, "{{table_name}}", table)

	// Get nested relationships
	r := g.getNestedRelations(table, "", []string{table}, 0)

	rt := BaseGormModelRelationshipTemplate
	relationshipEntries := ""
	if len(r) > 0 {
		relationshipEntries = "\n"
		for _, nested := range r {
			relationshipEntries += fmt.Sprintf("\t\t\"%v\",\n", nested)
		}
		relationshipEntries += "\t"
	}
	rt = strings.ReplaceAll(rt, "{{model_name}}", modelName)
	rt = strings.ReplaceAll(rt, "{{relationships}}", relationshipEntries)
	t = strings.ReplaceAll(t, "{{relationships}}", rt)

	// Connection method
	ct := BaseGormModelConnectionTemplate
	ct = strings.ReplaceAll(ct, "{{model_name}}", modelName)
	ct = strings.ReplaceAll(ct, "{{connection_name}}", GetConnectionByTableName(table))
	t = strings.ReplaceAll(t, "{{connection}}", ct)

	// Auto-import detection
	imports := ""
	if strings.Contains(t, "null.") {
		imports += "\n\t\"github.com/volatiletech/null/v8\""
	}
	if strings.Contains(t, "time.Time") {
		imports += "\n\t\"time\""
	}
	if strings.Contains(t, "types.") {
		imports += "\n\t\"github.com/volatiletech/sqlboiler/v4/types\""
	}
	if imports != "" {
		imports += "\n"
		importTemplate = strings.ReplaceAll(importTemplate, "{{imports}}", imports)
		t = strings.ReplaceAll(t, "{{imports}}", "\n"+importTemplate+"\n")
	} else {
		t = strings.ReplaceAll(t, "{{imports}}", "")
	}

	// Write to file
	if err := g.writeToFile(fileName, t); err != nil {
		return fmt.Errorf("writing model file for %s failed: %w", table, err)
	}

	g.logger.Info().Any("fileName", fileName).Msg("Generated model")

	g.models = append(g.models, modelName)

	return nil
}

// resolveTablesToGenerate resolves the tables to generate
func (g *Generator) resolveTablesToGenerate(tables []string) ([]string, error) {
	if len(tables) > 0 {
		return tables, nil
	}

	schemaTables, err := g.schemaLookup.GetTableNames()
	if err != nil {
		g.logger.Error().Err(err).Msg("error getting table names")
		return nil, err
	}

	cfg := GetGenerateModelConfig()

	finalTables := make([]string, 0, len(schemaTables))
	for _, table := range schemaTables {
		ignore := false
		for _, iTable := range cfg.Database.IgnoreTables {
			if strings.Contains(table, iTable) {
				g.logger.Debug().Msgf("Ignoring table [%s] because it is in the ignore list", table)
				ignore = true
				break
			}
		}
		if !ignore {
			finalTables = append(finalTables, table)
		}
	}

	// sort the tables
	sort.Strings(finalTables)

	return finalTables, nil
}

// SetWithControllers sets the withControllers option
func (g *Generator) SetWithControllers(b bool) {
	g.logger.Info().Msg("Setting withControllers option")
	g.withControllers = b
}

const (
	crudControllerPath = "./internal/http/crudcontrollers/"
)

// resolvePrimaryKey resolves the primary key for a table
func (g *Generator) resolvePrimaryKey(keys []DbSchemaRowResult) DbSchemaRowResult {
	for _, key := range keys {
		if key.Column == "id" {
			return key
		}
	}
	for _, key := range keys {
		if key.ColumnKey.String == "PRI" || key.OrdinalPosition == "1" {
			return key
		}
	}
	return keys[0]
}

// adjustKeyName adjusts the key name for a table
func (g *Generator) adjustKeyName(col string) (string, string) {
	keyName := strcase.ToCamel(col)
	newKeyName := keyName
	if keyName == "Id" {
		keyName = "ID"
	}
	if keyName == "Type" {
		newKeyName = "TypeId"
	}
	return keyName, newKeyName
}

// buildParamLine builds the param line for a table
func (g *Generator) buildParamLine(name, dataType string) string {
	if strings.Contains(dataType, "int") {
		return fmt.Sprintf(
			"%s, err := strconv.Atoi(c.Param(\"%s\"))",
			strcase.ToLowerCamel(name),
			strcase.ToLowerCamel(name),
		)
	}
	return ""
}

// buildQueryParams builds the query params for a table
func (g *Generator) buildQueryParams(keys []DbSchemaRowResult) string {
	var b strings.Builder
	for i, key := range keys {
		if i == 0 {
			continue
		}
		col := key.Column
		if col == "type" {
			col = "typeId"
		}

		dataType := key.DataType
		if strings.Contains(dataType, "(") {
			dataType = strings.Split(dataType, "(")[0]
		}

		b.WriteString(fmt.Sprintf(`
	// key param [%s] position [%s] type [%s]
	if len(c.QueryParam("%s")) > 0 {
		%sParam, err := strconv.Atoi(c.QueryParam("%s"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [%s] err [%%s]", err.Error())})
		}

		params = append(params, %sParam)
		keys = append(keys, "%s = ?")
	}
`, key.Column, key.OrdinalPosition, dataType,
			key.Column,
			strcase.ToLowerCamel(col),
			key.Column,
			key.Column,
			strcase.ToLowerCamel(col),
			key.Column))
	}
	return b.String()
}

// MakeController generates a controller for a table
func (g *Generator) MakeController(table string) {
	keys, err := g.schemaLookup.GetTableKeys(table)
	if err != nil {
		g.logger.Error().Err(err).Msgf("Error getting keys for table %s", table)
		return
	}
	if len(keys) == 0 {
		g.logger.Info().Msgf("No keys found for table %s", table)
		return
	}

	primaryKey := g.resolvePrimaryKey(keys)
	keyName, newKeyName := g.adjustKeyName(primaryKey.Column)

	paramLine := g.buildParamLine(newKeyName, primaryKey.DataType)
	queryParams := g.buildQueryParams(keys)

	entity := g.pluralize.Singular(table)
	entityCamel := strcase.ToCamel(entity)
	entitySnake := strcase.ToSnake(entity)

	data := struct {
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
	}{
		KeyNameModelField:     keyName,
		KeyName:               newKeyName,
		KeyColumn:             primaryKey.Column,
		KeyNameLowerCamel:     strcase.ToLowerCamel(newKeyName),
		EntityName:            entityCamel,
		EntityNamePlural:      g.pluralize.Plural(entityCamel),
		EntityNameSnake:       entitySnake,
		EntityNameSnakePlural: g.pluralize.Plural(entitySnake),
		EntityNameCamel:       strcase.ToLowerCamel(entity),
		EntityNameCamelPlural: g.pluralize.Plural(strcase.ToLowerCamel(entity)),
		ParamLine:             paramLine,
		QueryParams:           queryParams,
	}

	tpl, err := template.ParseFiles("./internal/model/templates/crud_controller.tmpl")
	if err != nil {
		g.logger.Fatal().Err(err).Msg("Failed to parse controller template")
	}

	var out bytes.Buffer
	if err := tpl.ExecuteTemplate(&out, "crud_controller.tmpl", data); err != nil {
		g.logger.Fatal().Err(err).Msg("Failed to execute template")
	}

	fileName := fmt.Sprintf("%v%v_controller.go", crudControllerPath, entitySnake)
	if err := os.WriteFile(fileName, out.Bytes(), 0644); err != nil {
		g.logger.Fatal().Err(err).Msg("Failed to write controller file")
	}

	g.logger.Info().
		Str("fileName", fileName).
		Str("entityName", data.EntityName).
		Msg("Generated controller")
}

// getCrudControllerNames returns the names of the crud controllers
func (g *Generator) getCrudControllerNames() ([]string, error) {
	files, err := os.ReadDir(crudControllerPath)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	var names []string
	for _, f := range files {
		path := crudControllerPath + f.Name()
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open file %s: %w", path, err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "type") && strings.Contains(line, "struct") && strings.Contains(line, "Controller") {
				line = strings.ReplaceAll(line, "type ", "")
				line = strings.ReplaceAll(line, "struct {", "")
				names = append(names, strings.TrimSpace(line))
			}
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("scan file %s: %w", path, err)
		}
	}
	return names, nil
}

// SyncControllersToInjector syncs the controllers to the dependency injector
func (g *Generator) SyncControllersToInjector() {
	const (
		crudInjectFile   = "./boot/inject_http_crud_controllers.go"
		crudTemplatePath = "./internal/model/templates/inject_http_crud_controller.tmpl"
		templateName     = "inject_http_crud_controller.tmpl"
	)

	controllerNames, err := g.getCrudControllerNames()
	if err != nil {
		g.logger.Fatal().Err(err).Msg("failed to collect controller names")
	}

	var newControllers, controllerParams, controllerRegisters strings.Builder
	for _, name := range controllerNames {
		newControllers.WriteString(fmt.Sprintf("\tcrudcontrollers.New%v,\n", name))
		controllerParams.WriteString(fmt.Sprintf("\t%v *crudcontrollers.%v,\n", strcase.ToLowerCamel(name), name))
		controllerRegisters.WriteString(fmt.Sprintf("\t\t\t%v,\n", strcase.ToLowerCamel(name)))
	}

	tpl, err := template.ParseFiles(crudTemplatePath)
	if err != nil {
		g.logger.Fatal().Err(err).Msg("failed to parse template")
	}

	var out bytes.Buffer
	err = tpl.ExecuteTemplate(&out, templateName, struct {
		NewControllers      string
		ControllersParam    string
		ControllersRegister string
	}{
		NewControllers:      strings.TrimSuffix(newControllers.String(), "\n"),
		ControllersParam:    strings.TrimSuffix(controllerParams.String(), "\n"),
		ControllersRegister: strings.TrimSuffix(controllerRegisters.String(), "\n"),
	})
	if err != nil {
		g.logger.Fatal().Err(err).Msg("failed to execute template")
	}

	if err := os.WriteFile(crudInjectFile, out.Bytes(), 0644); err != nil {
		g.logger.Fatal().Err(err).Msg("failed to write injector file")
	}

	g.logger.Info().
		Str("path", crudInjectFile).
		Int("count", len(controllerNames)).
		Msg("Synced controllers to injector")
}
