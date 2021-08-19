package generators

import (
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type GenerateModelContext struct {
	TablesToGenerate []string
	Relationships    []ForeignKeyMappings
}

type GenerateModel struct {
	options   GenerateModelContext
	logger    *logrus.Logger
	gorm      *gorm.DB
	pluralize *pluralize.Client
	debug     bool
}

func NewGenerateModel(options GenerateModelContext, logger *logrus.Logger, gorm *gorm.DB) *GenerateModel {
	return &GenerateModel{
		options:   options,
		logger:    logger,
		gorm:      gorm,
		pluralize: pluralize.NewClient(),
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

type ModelGenerateResponse struct {
	Table             string
	GormRelationships []string
}

func (gm *GenerateModel) Generate() []ModelGenerateResponse {
	gm.options.Relationships = gm.loadRelationships()
	relationships := gm.options.Relationships
	tableNames := gm.getTableNames()

	// if no argument pull from relationships
	if len(gm.options.TablesToGenerate) == 0 {
		gm.options.TablesToGenerate = GetDatabaseTables()
	}

	// metadata to respond with
	modelGenerateResponse := make([]ModelGenerateResponse, len(gm.options.TablesToGenerate))

	for _, genModel := range gm.options.TablesToGenerate {
		for _, table := range tableNames {
			if genModel != "all" && table != genModel {
				continue
			}

			t := BaseGormModelTemplate
			importTemplate := BaseDependencyImportTemplate
			t = strings.ReplaceAll(t, "{{model_name}}", gm.pluralize.Singular(strcase.ToCamel(table)))

			// calculate field and type length for formatting
			maxColumnLengthInTable := 0
			maxDataTypeLengthInTable := 0
			for _, def := range gm.getColumnDefinitions(table) {
				if len(def.Field) >= maxColumnLengthInTable {
					maxColumnLengthInTable = len(def.Field)
				}

				translatedType := gm.translateDataType(def)
				if len(translatedType) >= maxDataTypeLengthInTable {
					maxDataTypeLengthInTable = len(translatedType)
				}
			}

			// calculate relationship field and type lengths
			for _, relation := range relationships {
				if table != relation.Table {
					continue
				}

				col := strcase.ToCamel(gm.pluralize.Plural(relation.RemoteTable))
				colType := gm.getRelationshipTypeModelAttributePrefix(relation) + strcase.ToCamel(gm.pluralize.Singular(relation.RemoteTable))

				if len(col) >= maxColumnLengthInTable {
					maxColumnLengthInTable = len(col)
				}

				if len(colType) >= maxDataTypeLengthInTable {
					maxDataTypeLengthInTable = len(colType)
				}
			}

			// gen model fields
			modelFields := ""
			for _, def := range gm.getColumnDefinitions(table) {

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

				modelFields += fmt.Sprintf(
					"\t%-*s%-*s `json:\"%v\" gorm:\"Column:%v\"`\n",
					maxColumnLengthInTable+1,
					structFieldName,
					maxDataTypeLengthInTable,
					gm.translateDataType(def),
					jsonFieldName,
					def.Field,
				)
			}

			// write relationships to model attributes
			for _, relation := range relationships {
				if table != relation.Table {
					continue
				}

				maxColInTable := maxColumnLengthInTable + 1
				relationshipAttributeTypeName := fmt.Sprintf(
					"%v%v",
					gm.getRelationshipTypeModelAttributePrefix(relation),
					strcase.ToCamel(gm.pluralize.Singular(relation.RemoteTable)),
				)
				relationshipFieldNameSnakeCase := strcase.ToSnake(gm.pluralize.Singular(relation.RemoteTable))

				switch relation.RelationType {
				case RelationshipType1to1:
					relationshipAttributeName := strcase.ToCamel(gm.pluralize.Singular(relation.RemoteTable))
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
					relationshipAttributeName := strcase.ToCamel(gm.pluralize.Plural(relation.RemoteTable))
					modelFields += fmt.Sprintf(
						"\t%-*s%-*s `json:\"%v,omitempty\" gorm:\"foreignKey:%v;association_foreignkey:%v\"`\n",
						maxColInTable,
						relationshipAttributeName,
						maxDataTypeLengthInTable,
						relationshipAttributeTypeName,
						gm.pluralize.Plural(relationshipFieldNameSnakeCase),
						relation.RemoteKey,
						relation.Key,
					)
				case RelationshipTypeManyTo1:
					// todo: inverse
				}
			}

			// write model fields
			t = strings.ReplaceAll(t, "{{model_fields}}", modelFields)

			// nested relationships
			rt := BaseGormModelRelationshipTemplate
			relationshipEntries := ""
			nestedRelationships := gm.getNestedRelationshipsFromTable(table, "")
			if len(nestedRelationships) > 0 {
				relationshipEntries = "\n"
				for _, nested := range nestedRelationships {
					relationshipEntries = fmt.Sprintf("%v\t\t\"%v\",\n", relationshipEntries, nested)
				}
				relationshipEntries += "\t"
			}
			rt = strings.ReplaceAll(rt, "{{model_name}}", gm.pluralize.Singular(strcase.ToCamel(table)))
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
			ct = strings.ReplaceAll(ct, "{{model_name}}", gm.pluralize.Singular(strcase.ToCamel(table)))
			ct = strings.ReplaceAll(ct, "{{connection_name}}", GetConnectionByTableName(table))

			// final template
			t = strings.ReplaceAll(t, "{{table_name}}", table)
			t = strings.ReplaceAll(t, "{{relationships}}", rt)
			t = strings.ReplaceAll(t, "{{connection}}", ct)

			// write file
			fileName := "./models/" + strcase.ToSnake(table) + ".go"

			// Create new cmd
			file, err := os.Create(fileName)
			if err != nil {
				gm.logger.Fatal(err)
			}

			defer file.Close()

			_, err = file.WriteString(t)
			if err != nil {
				gm.logger.Fatal(err)
			}

			fmt.Println(fmt.Sprintf("Generated [%v]", fileName))

			modelGenerateResponse = append(
				modelGenerateResponse, ModelGenerateResponse{
					Table:             table,
					GormRelationships: nestedRelationships,
				},
			)

		}
	}

	return modelGenerateResponse
}

// return relationship type model prefix
func (gm *GenerateModel) getRelationshipTypeModelAttributePrefix(r ForeignKeyMappings) string {
	if r.RelationType == RelationshipType1to1 {
		return "*"
	}
	if r.RelationType == RelationshipType1toMany {
		return "[]"
	}

	return ""
}

func (gm *GenerateModel) getNestedRelationshipsFromTable(table string, prefix string) []string {
	relationshipNames := make([]string, 0)

	if prefix != "" {
		prefix = fmt.Sprintf("%v.", prefix)
	}

	for _, relation := range gm.options.Relationships {
		if table != relation.Table {
			continue
		}

		relationshipAttributeName := ""

		switch relation.RelationType {
		case RelationshipType1to1:
			relationshipAttributeName = strcase.ToCamel(gm.pluralize.Singular(relation.RemoteTable))
		case RelationshipType1toMany:
			relationshipAttributeName = strcase.ToCamel(gm.pluralize.Plural(relation.RemoteTable))
		case RelationshipTypeManyTo1:
			// todo: inverse
		}

		relationshipNames = append(relationshipNames, fmt.Sprintf("%v%v", prefix, relationshipAttributeName))
		passedDownPrefix := fmt.Sprintf("%v%v", prefix, relationshipAttributeName)
		relationshipNames = append(
			relationshipNames,
			gm.getNestedRelationshipsFromTable(relation.RemoteTable, passedDownPrefix)...,
		)
	}

	sort.Strings(relationshipNames)

	return relationshipNames
}

// return Table names from database
func (gm *GenerateModel) getTableNames() []string {
	rows, err := gm.gorm.DB().Query("SHOW TABLES")
	if err != nil {
		gm.logger.Warn(err)
	}

	tableNames := make([]string, 0)

	defer rows.Close()
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			gm.logger.Warn(err)
		}

		tableNames = append(tableNames, column)
	}

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

func (gm *GenerateModel) getColumnDefinitions(tableName string) []ShowColumns {
	rows, err := gm.gorm.DB().Query(fmt.Sprintf("SHOW COLUMNS FROM %v", tableName))
	if err != nil {
		gm.logger.Warn(err)
	}

	columnDefs := make([]ShowColumns, 0)

	defer rows.Close()
	for rows.Next() {
		var showColumns ShowColumns
		_ = rows.Scan(
			&showColumns.Field,
			&showColumns.Type,
			&showColumns.Null,
			&showColumns.Key,
			&showColumns.Default,
			&showColumns.Extra,
		)

		columnDefs = append(columnDefs, showColumns)
	}

	return columnDefs
}

func (gm *GenerateModel) translateDataType(column ShowColumns) string {
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
			return "types.Decimal"
		case "json":
			return "types.JSON"
		default:
			return "string"
		}
	}
}

const dbRelationshipConfig = "./generators/config/db-relationships.yml"

func (gm *GenerateModel) loadRelationships() []ForeignKeyMappings {
	// load yaml
	databaseSchemaYaml, err := ioutil.ReadFile(dbRelationshipConfig)
	if err != nil {
		gm.logger.Fatal(err)
	}

	// unmarshal yaml
	dbRelationships := make(map[string][]string, 0)
	err = yaml.Unmarshal(databaseSchemaYaml, &dbRelationships)
	if err != nil {
		gm.logger.Fatalf("error: %v", err)
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

			if !gm.isValidRelationshipType(relationType) {
				gm.logger.Fatalf(
					"Invalid relationship type [%v] [%v] in [%v]!\n",
					relationType,
					relationSignature,
					dbRelationshipConfig,
				)
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

func (gm *GenerateModel) isValidRelationshipType(relationshipType string) bool {
	switch relationshipType {
	case RelationshipType1to1:
		return true
	case RelationshipType1toMany:
		return true
	}

	return false
}
