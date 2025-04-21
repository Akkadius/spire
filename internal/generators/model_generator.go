package generators

import (
	"bytes"
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"os"
	"sort"
	"strings"
	"text/template"
)

type ModelGenerator struct {
	logger       *logger.AppLogger
	gorm         *gorm.DB
	pluralize    *pluralize.Client
	debugEnabled bool

	// attributes
	models        []string
	relationships []ForeignKeyMappings
}

func NewGenerateModels(logger *logger.AppLogger, gorm *gorm.DB) *ModelGenerator {
	return &ModelGenerator{
		logger:    logger,
		gorm:      gorm,
		pluralize: pluralize.NewClient(),
		models:    make([]string, 0),
	}
}

// templates
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

// ModelRelationships is a struct that represents a model and its relationships
type ModelRelationships struct {
	Table         string   `json:"table"`
	ModelName     string   `json:"model_name"`
	Relationships []string `json:"relationships"`
}

// Generate generates the models for the given tables
// if no tables are provided, it will generate models for all tables
// it will also generate a models.go file in the models directory
// the models.go file will contain a list of models that implement the Modelable interface
// and a list of model names
func (g *ModelGenerator) Generate(tables []string) {
	g.relationships = g.loadRelationships()
	tableNames := g.getTableNames()
	g.models = make([]string, 0)

	tablesToGenerate := make([]string, 0)

	// if no argument pull from relationships
	if len(tables) == 0 {
		tablesToGenerate = GetDatabaseTables()
	}

	// TablesToGenerate is just a list of tables table1,table2
	for _, genModel := range tablesToGenerate {
		for _, table := range tableNames {
			if genModel != "all" && table != genModel {
				continue
			}

			err := g.generateModelFileForTable(table)
			if err != nil {
				g.logger.Error().Err(err).Msgf("error generating model for table %s", table)
				continue
			}
		}
	}

	err := g.generateModelsFile()
	if err != nil {
		g.logger.Error().Err(err).Msg("error generating models file")
	}
}

// return relationship type model prefix
func (g *ModelGenerator) getRelationshipTypeModelAttributePrefix(r ForeignKeyMappings) string {
	if r.RelationType == RelationshipType1to1 {
		return "*"
	}
	if r.RelationType == RelationshipType1toMany {
		return "[]"
	}

	return ""
}

func exists(a []string, element string) bool {
	for _, e := range a {
		if e == element {
			return true
		}
	}

	return false
}

// getNestedRelationshipsFromTable returns a list of relationships for a table
// it will recursively check for relationships in the table
// it will skip relationships that are already in the parent tables
func (g *ModelGenerator) getNestedRelationshipsFromTable(table string, prefix string, parentTables []string, level int) []string {
	relationshipNames := make([]string, 0)

	if prefix != "" {
		prefix = fmt.Sprintf("%v.", prefix)
	}

	g.debug(fmt.Sprintf("-- [getNestedRelationshipsFromTable] [%v] table [%v]", level, table))

	if exists(parentTables, table) && level > 0 {
		return relationshipNames
	}

	parentTables = append(parentTables, table)
	currentLevel := level + 1
	for _, relation := range g.relationships {
		if table != relation.Table {
			continue
		}

		g.debug(fmt.Sprintf("-- [getNestedRelationshipsFromTable] [%v] table [%v] relation [%v] remote [%v]", level, table, relation, relation.RemoteTable))

		if len(parentTables) > 0 {
			for _, parentTable := range parentTables {
				if relation.RemoteTable == parentTable {
					g.debug(
						fmt.Sprintf(
							"---- [getNestedRelationshipsFromTable] [%v] remote table [%v] is a parent table [%v] skipping",
							level,
							relation.RemoteTable,
							parentTable,
						),
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
			g.getNestedRelationshipsFromTable(relation.RemoteTable, passedDownPrefix, parentTables, currentLevel)...,
		)
	}

	sort.Strings(relationshipNames)

	//g.debug(fmt.Sprintf("-- [getNestedRelationshipsFromTable] relationshipNames [%v]", relationshipNames))

	return relationshipNames
}

// return Table names from database
func (g *ModelGenerator) getTableNames() []string {
	tableNames := make([]string, 0)
	g.gorm.Raw("SHOW TABLES").Scan(&tableNames)

	return tableNames
}

// ShowColumns is a struct that represents a row in the SHOW COLUMNS
type ShowColumns struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

// getColumnDefinitions returns the column definitions for a table
func (g *ModelGenerator) getColumnDefinitions(tableName string) []ShowColumns {
	columnDefs := make([]ShowColumns, 0)
	g.gorm.Raw(fmt.Sprintf("SHOW COLUMNS FROM %v", tableName)).Scan(&columnDefs)

	return columnDefs
}

const dbRelationshipConfig = "./internal/generators/config/db-relationships.yml"

// loadRelationships loads the relationships from the yaml file
// and returns a list of ForeignKeyMappings
// the yaml file should be in the format:
// table_name:
//   - relation_type local_key->remote_table:remote_key
func (g *ModelGenerator) loadRelationships() []ForeignKeyMappings {
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
func (g *ModelGenerator) isValidRelationshipType(relationshipType string) bool {
	switch relationshipType {
	case RelationshipType1to1:
		return true
	case RelationshipType1toMany:
		return true
	}

	return false
}

// debug is a helper function to log debug messages
// it checks if debug is enabled in the config or if the DEBUG environment variable is set to 1
// and logs the message if so
func (g *ModelGenerator) debug(msg string) {
	if g.debugEnabled || env.GetInt("DEBUG", "0") > 0 {
		g.logger.Debug().Msg(msg)
	}
}

// generateModelsFile generates the models file
// contains a list of models that implement the Modelable interface
// and a list of model names
// outputs a file in ./internal/models/models.go
func (g *ModelGenerator) generateModelsFile() error {
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
func (g *ModelGenerator) writeToFile(path, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

// getMaxFieldAndTypeLengths returns the max field and type lengths for a table
func (g *ModelGenerator) getMaxFieldAndTypeLengths(table string, defs []ShowColumns) (int, int) {
	maxFieldLen := 0
	maxTypeLen := 0

	for _, col := range defs {
		if len(col.Field) > maxFieldLen {
			maxFieldLen = len(col.Field)
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

		g.debug(fmt.Sprintf("-- col [%v] colType [%v]", col, colType))
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
func (g *ModelGenerator) buildModelFields(columns []ShowColumns, maxColLen, maxTypeLen int) string {
	var b strings.Builder

	for _, def := range columns {
		g.debug(fmt.Sprintf("-- def [%v] table [%v]", def, def.Field))

		structFieldName := g.getStructFieldName(def.Field)
		jsonFieldName := strcase.ToSnake(def.Field)
		goType := translateDataType(def)

		b.WriteString(fmt.Sprintf(
			"\t%-*s%-*s `json:\"%v\" gorm:\"Column:%v\"`\n",
			maxColLen+1,
			structFieldName,
			maxTypeLen,
			goType,
			jsonFieldName,
			def.Field,
		))

	}

	return b.String()
}

// translateDataType translates the data type from the database to Go
func (g *ModelGenerator) getStructFieldName(field string) string {
	if field == "id" {
		return "ID"
	}

	name := strcase.ToCamel(field)
	if strings.HasSuffix(field, "_") {
		name += "2"
	}
	return name
}

// buildRelationshipFields builds the relationship fields for a table
func (g *ModelGenerator) buildRelationshipFields(table string, maxColLen int, maxTypeLen int) string {
	var b strings.Builder

	for _, relation := range g.relationships {
		if relation.Table != table {
			continue
		}

		g.debug(fmt.Sprintf("-- relationships [%v]", relation))

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

// generateModelFileForTable generates the model file for a table
func (g *ModelGenerator) generateModelFileForTable(table string) error {
	modelName := g.pluralize.Singular(strcase.ToCamel(table))
	fileName := "./internal/models/" + strcase.ToSnake(table) + ".go"

	g.debug(fmt.Sprintf("-- Generating model for [%s] → [%s]", table, fileName))

	// Template skeleton
	t := BaseGormModelTemplate
	importTemplate := BaseDependencyImportTemplate

	// Get column data
	defs := g.getColumnDefinitions(table)
	maxColLen, maxTypeLen := g.getMaxFieldAndTypeLengths(table, defs)

	// Build model content
	modelFields := g.buildModelFields(defs, maxColLen, maxTypeLen)
	modelFields += g.buildRelationshipFields(table, maxColLen, maxTypeLen)

	// Replace model fields
	t = strings.ReplaceAll(t, "{{model_fields}}", modelFields)
	t = strings.ReplaceAll(t, "{{model_name}}", modelName)
	t = strings.ReplaceAll(t, "{{table_name}}", table)

	// Generate nested relationships
	nestedRelationships := g.getNestedRelationshipsFromTable(table, "", []string{table}, 0)

	rt := BaseGormModelRelationshipTemplate
	relationshipEntries := ""
	if len(nestedRelationships) > 0 {
		relationshipEntries = "\n"
		for _, nested := range nestedRelationships {
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

	fmt.Printf("Generated [%s]\n", fileName)
	g.models = append(g.models, modelName)

	return nil
}
