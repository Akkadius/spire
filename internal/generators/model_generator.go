package generators

import (
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
)

type GenerateModelContext struct {
	TablesToGenerate []string
	Relationships    []ForeignKeyMappings
}

type GenerateModel struct {
	options      GenerateModelContext
	logger       *logger.AppLogger
	gorm         *gorm.DB
	pluralize    *pluralize.Client
	debugEnabled bool

	models []string
}

func NewGenerateModel(options GenerateModelContext, logger *logger.AppLogger, gorm *gorm.DB) *GenerateModel {
	return &GenerateModel{
		options:   options,
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

type ForeignKeyMappings struct {
	Table        string `json:"table"`         // local Table
	Key          string `json:"key"`           // local Key
	RemoteTable  string `json:"remote_table"`  // remote Table
	RemoteKey    string `json:"remote_key"`    // remote Key
	RelationType string `json:"relation_type"` // relationship type
}

type ModelRelationships struct {
	Table         string   `json:"table"`
	ModelName     string   `json:"model_name"`
	Relationships []string `json:"relationships"`
}

func (g *GenerateModel) Generate() {
	g.options.Relationships = g.loadRelationships()
	relationships := g.options.Relationships
	tableNames := g.getTableNames()
	g.models = make([]string, 0)

	// if no argument pull from relationships
	if len(g.options.TablesToGenerate) == 0 {
		g.options.TablesToGenerate = GetDatabaseTables()
	}

	var modelRelationships []ModelRelationships

	// TablesToGenerate is just a list of tables table1,table2
	for _, genModel := range g.options.TablesToGenerate {
		g.debug(fmt.Sprintf("Generating model for table [%v]", genModel))

		for _, table := range tableNames {
			if genModel != "all" && table != genModel {
				continue
			}

			g.debug(fmt.Sprintf("-- Checking [%v]", table))

			// base gorm template file that we will use to generate
			t := BaseGormModelTemplate
			// separate template for handling and injecting imports for models
			importTemplate := BaseDependencyImportTemplate
			t = strings.ReplaceAll(t, "{{model_name}}", g.pluralize.Singular(strcase.ToCamel(table)))

			g.models = append(g.models, g.pluralize.Singular(strcase.ToCamel(table)))

			// calculate field and type length for formatting
			maxColumnLengthInTable := 0
			maxDataTypeLengthInTable := 0
			for _, def := range g.getColumnDefinitions(table) {
				if len(def.Field) >= maxColumnLengthInTable {
					maxColumnLengthInTable = len(def.Field)
				}

				translatedType := g.translateDataType(def)
				if len(translatedType) >= maxDataTypeLengthInTable {
					maxDataTypeLengthInTable = len(translatedType)
				}
			}

			// calculate relationship field and type lengths
			for _, relation := range relationships {
				if table != relation.Table {
					continue
				}

				col := strcase.ToCamel(g.pluralize.Plural(relation.RemoteTable))
				colType := g.getRelationshipTypeModelAttributePrefix(relation) + strcase.ToCamel(g.pluralize.Singular(relation.RemoteTable))

				if len(col) >= maxColumnLengthInTable {
					maxColumnLengthInTable = len(col)
				}

				if len(colType) >= maxDataTypeLengthInTable {
					maxDataTypeLengthInTable = len(colType)
				}

				g.debug(fmt.Sprintf("-- col [%v] colType [%v]", col, colType))
			}

			g.debug(fmt.Sprintf("-- maxDataTypeLengthInTable [%v]", maxDataTypeLengthInTable))
			g.debug(fmt.Sprintf("-- maxColumnLengthInTable [%v]", maxColumnLengthInTable))

			// gen model fields
			modelFields := ""
			structFieldNames := map[string]bool{}
			for _, def := range g.getColumnDefinitions(table) {

				g.debug(fmt.Sprintf("-- def [%v] table [%v]", def, table))

				structFieldName := ""
				jsonFieldName := strcase.ToSnake(def.Field)
				if def.Field == "id" {
					structFieldName = "ID"
				} else {
					structFieldName = strcase.ToCamel(def.Field)
				}

				lastCharacterColumn := def.Field[len(def.Field)-1:]
				if lastCharacterColumn == "_" {
					structFieldName += "2"
				}

				modelField := fmt.Sprintf(
					"\t%-*s%-*s `json:\"%v\" gorm:\"Column:%v\"`\n",
					maxColumnLengthInTable+1,
					structFieldName,
					maxDataTypeLengthInTable,
					g.translateDataType(def),
					jsonFieldName,
					def.Field,
				)

				structFieldNames[structFieldName] = true

				g.debug(fmt.Sprintf("-- modelField [%v]", strings.TrimSpace(modelField)))

				modelFields += modelField
			}

			// write relationships to model attributes
			for _, relation := range relationships {
				if table != relation.Table {
					continue
				}

				g.debug(fmt.Sprintf("-- relationships [%v]", relation))

				maxColInTable := maxColumnLengthInTable + 1
				relationshipAttributeTypeName := fmt.Sprintf(
					"%v%v",
					g.getRelationshipTypeModelAttributePrefix(relation),
					strcase.ToCamel(g.pluralize.Singular(relation.RemoteTable)),
				)
				relationshipFieldNameSnakeCase := strcase.ToSnake(g.pluralize.Singular(relation.RemoteTable))

				switch relation.RelationType {
				case RelationshipType1to1:
					relationshipAttributeName := strcase.ToCamel(g.pluralize.Singular(relation.RemoteTable))
					// TODO: Revisit later
					//if _, ok := structFieldNames[relationshipAttributeName]; ok {
					//	relationshipAttributeName = relationshipAttributeName + "Relation"
					//	relationshipFieldNameSnakeCase = relationshipFieldNameSnakeCase + "_relation"
					//}

					modelFields += fmt.Sprintf(
						"\t%-*s%-*s `json:\"%v,omitempty\" gorm:\"foreignKey:%v;references:%v\"`\n",
						maxColInTable,
						relationshipAttributeName,
						maxDataTypeLengthInTable,
						relationshipAttributeTypeName,
						relationshipFieldNameSnakeCase,
						relation.Key,
						relation.RemoteKey,
					)
				case RelationshipType1toMany:
					relationshipAttributeName := strcase.ToCamel(g.pluralize.Plural(relation.RemoteTable))
					// TODO: Revisit later
					//if _, ok := structFieldNames[relationshipAttributeName]; ok {
					//	relationshipAttributeName = relationshipAttributeName + "Relation"
					//	relationshipFieldNameSnakeCase = relationshipFieldNameSnakeCase + "_relation"
					//}

					modelFields += fmt.Sprintf(
						"\t%-*s%-*s `json:\"%v,omitempty\" gorm:\"foreignKey:%v;references:%v\"`\n",
						maxColInTable,
						relationshipAttributeName,
						maxDataTypeLengthInTable,
						relationshipAttributeTypeName,
						g.pluralize.Plural(relationshipFieldNameSnakeCase),
						relation.RemoteKey,
						relation.Key,
					)
				case RelationshipTypeManyTo1:
					// todo: inverse
				}
			}

			g.debug(fmt.Sprintf("-- writing model fields"))

			// write model fields
			t = strings.ReplaceAll(t, "{{model_fields}}", modelFields)

			// nested relationships
			rt := BaseGormModelRelationshipTemplate
			relationshipEntries := ""
			nestedRelationships := g.getNestedRelationshipsFromTable(table, "", []string{table}, 0)
			if len(nestedRelationships) > 0 {
				relationshipEntries = "\n"
				for _, nested := range nestedRelationships {
					relationshipEntries = fmt.Sprintf("%v\t\t\"%v\",\n", relationshipEntries, nested)
				}
				relationshipEntries += "\t"
			}
			rt = strings.ReplaceAll(rt, "{{model_name}}", g.pluralize.Singular(strcase.ToCamel(table)))
			rt = strings.ReplaceAll(rt, "{{relationships}}", relationshipEntries)

			// handle imports
			imports := ""
			if strings.Contains(t, "null") {
				imports += fmt.Sprintf("\n\t%v", "\"github.com/volatiletech/null/v8\"")
			}
			if strings.Contains(t, "time.Time") {
				imports += fmt.Sprintf("\n\t%v", "\"time\"")
			}
			if strings.Contains(t, "types.") {
				imports += fmt.Sprintf("\n\t%v", "\"github.com/volatiletech/sqlboiler/v4/types\"")
			}

			if len(imports) > 0 {
				imports += "\n"
				importTemplate = strings.ReplaceAll(importTemplate, "{{imports}}", imports)
				t = strings.ReplaceAll(t, "{{imports}}", "\n"+importTemplate+"\n")
			} else {
				t = strings.ReplaceAll(t, "{{imports}}", "")
			}

			// connection
			ct := BaseGormModelConnectionTemplate
			ct = strings.ReplaceAll(ct, "{{model_name}}", g.pluralize.Singular(strcase.ToCamel(table)))
			ct = strings.ReplaceAll(ct, "{{connection_name}}", GetConnectionByTableName(table))

			// final template
			t = strings.ReplaceAll(t, "{{table_name}}", table)
			t = strings.ReplaceAll(t, "{{relationships}}", rt)
			t = strings.ReplaceAll(t, "{{connection}}", ct)

			// write file
			fileName := "./internal/models/" + strcase.ToSnake(table) + ".go"

			// Create new cmd
			file, err := os.Create(fileName)
			if err != nil {
				g.logger.Fatal().Err(err).Msg("error creating file")
			}

			defer file.Close()

			_, err = file.WriteString(t)
			if err != nil {
				g.logger.Fatal().Err(err).Msg("error writing to file")
			}

			fmt.Println(fmt.Sprintf("Generated [%v]", fileName))

			modelRelationships = append(modelRelationships, ModelRelationships{
				Table:         table,
				ModelName:     g.pluralize.Singular(strcase.ToCamel(table)),
				Relationships: nestedRelationships,
			})
		}
	}

	g.generateModelsFile()

	sort.Slice(modelRelationships, func(i, j int) bool {
		return modelRelationships[i].Table < modelRelationships[j].Table
	})
}

// return relationship type model prefix
func (g *GenerateModel) getRelationshipTypeModelAttributePrefix(r ForeignKeyMappings) string {
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

func (g *GenerateModel) getNestedRelationshipsFromTable(table string, prefix string, parentTables []string, level int) []string {
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
	for _, relation := range g.options.Relationships {
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
func (g *GenerateModel) getTableNames() []string {
	tableNames := make([]string, 0)
	g.gorm.Raw("SHOW TABLES").Scan(&tableNames)

	return tableNames
}

type ShowColumns struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func (g *GenerateModel) getColumnDefinitions(tableName string) []ShowColumns {
	columnDefs := make([]ShowColumns, 0)
	g.gorm.Raw(fmt.Sprintf("SHOW COLUMNS FROM %v", tableName)).Scan(&columnDefs)

	return columnDefs
}

func (g *GenerateModel) translateDataType(column ShowColumns) string {
	unsigned := strings.Contains(column.Type, "unsigned")
	nullable := strings.Contains(column.Null, "YES")
	columnType := strings.Split(column.Type, "(")[0]

	if nullable {
		switch columnType {
		case "tinyint":
			if columnType == "tinyint(1)" {
				return "null.Bool"
			} else if unsigned {
				return "null.Uint8"
			} else {
				return "null.Int8"
			}
		case "smallint":
			if unsigned {
				return "null.Uint16"
			} else {
				return "null.Int16"
			}
		case "mediumint":
			if unsigned {
				return "null.Uint32"
			} else {
				return "null.Int32"
			}
		case "int", "integer":
			if unsigned {
				return "null.Uint"
			} else {
				return "null.Int"
			}
		case "bigint":
			if unsigned {
				return "null.Uint64"
			} else {
				return "null.Int64"
			}
		case "float":
			return "null.Float32"
		case "double", "double precision", "real":
			return "null.Float64"
		case "boolean", "bool":
			return "null.Bool"
		case "date", "datetime", "timestamp":
			return "null.Time"
		case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
			return "null.Bytes"
		case "numeric", "decimal", "dec", "fixed":
			return "types.NullDecimal"
		case "json":
			return "null.JSON"
		default:
			return "null.String"
		}
	} else {
		switch columnType {
		case "tinyint":
			if columnType == "tinyint(1)" {
				return "bool"
			} else if unsigned {
				return "uint8"
			} else {
				return "int8"
			}
		case "smallint":
			if unsigned {
				return "uint16"
			} else {
				return "int16"
			}
		case "mediumint":
			if unsigned {
				return "uint32"
			} else {
				return "int32"
			}
		case "int", "integer":
			if unsigned {
				return "uint"
			} else {
				return "int"
			}
		case "bigint":
			if unsigned {
				return "uint64"
			} else {
				return "int64"
			}
		case "float":
			return "float32"
		case "double", "double precision", "real":
			return "float64"
		case "boolean", "bool":
			return "bool"
		case "date", "datetime", "timestamp":
			return "time.Time"
		case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
			return "[]byte"
		case "numeric", "decimal", "dec", "fixed":
			return "float32"
		case "json":
			return "types.JSON"
		default:
			return "string"
		}
	}
}

const dbRelationshipConfig = "./internal/generators/config/db-relationships.yml"

func (g *GenerateModel) loadRelationships() []ForeignKeyMappings {
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

func (g *GenerateModel) isValidRelationshipType(relationshipType string) bool {
	switch relationshipType {
	case RelationshipType1to1:
		return true
	case RelationshipType1toMany:
		return true
	}

	return false
}

func (g *GenerateModel) debug(msg string) {
	if g.debugEnabled || env.GetInt("DEBUG", "0") > 0 {
		g.logger.Debug().Msg(msg)
	}
}

func (g *GenerateModel) generateModelsFile() {
	// sort models
	sort.Slice(g.models, func(i, j int) bool {
		return g.models[i] < g.models[j]
	})

	modelData := `package models

func GetModels() []Modelable {
	return []Modelable{
		{{models}},
	}
}

func GetModelNames() []string {
	return []string{
		{{modelNames}},
	}
}
`

	// write models in a func GetModels() .go file in ./models/models.go
	f, err := os.Create("internal/models/models.go")
	if err != nil {
		g.logger.Fatal().Err(err).Msg("error creating file")
	}

	tempModels := make([]string, len(g.models))
	for i, m := range g.models {
		tempModels[i] = "\"" + m + "\""
	}

	modelData = strings.ReplaceAll(modelData, "{{modelNames}}", strings.Join(tempModels, ",\n\t\t"))

	for i, m := range g.models {
		g.models[i] = "&" + m + "{}"
	}

	modelData = strings.ReplaceAll(modelData, "{{models}}", strings.Join(g.models, ",\n\t\t"))

	// write
	_, err = f.Write([]byte(modelData))
	if err != nil {
		g.logger.Fatal().Err(err).Msg("error writing to file")
	}

	defer f.Close()

}
