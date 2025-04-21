package model

import (
	"database/sql"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/volatiletech/null/v8"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type DbSchemaLookup struct {
	db     *sql.DB
	logger *logger.AppLogger
}

func NewDbSchemaLookup(db *sql.DB, logger *logger.AppLogger) *DbSchemaLookup {
	return &DbSchemaLookup{db: db, logger: logger}
}

// DbSchemaRowResult represents a row in the database schema
type DbSchemaRowResult struct {
	Table           string      `yaml:",omitempty"`
	Column          string      `yaml:"column"`
	DataType        string      `yaml:"data_type"`
	ColumnKey       null.String `yaml:"column_key"`
	OrdinalPosition string      `yaml:"ordinal_position"`
	IsNullable      string      `yaml:"is_nullable"`
	ColumnType      string      `yaml:"column_type"`
	ColumnDefault   null.String `yaml:"column_default"`
}

// schemas is a map that holds the database schema for each table (cached)
var schemas map[string][]DbSchemaRowResult

// GetSchemas retrieves the database schema for a given database name
func (c *DbSchemaLookup) GetSchemas() (map[string][]DbSchemaRowResult, error) {
	if schemas != nil {
		return schemas, nil
	}

	var dbName string
	err := c.db.QueryRow("SELECT database()").Scan(&dbName)
	if err != nil {
		log.Fatal("failed to get database name:", err)
	}

	rows, err := c.db.Query(
		fmt.Sprintf(
			`
		SELECT
		  TABLE_NAME,
		  COLUMN_NAME,
		  COLUMN_TYPE,
		  COLUMN_KEY,
		  ORDINAL_POSITION,
		  IS_NULLABLE,
		  COLUMN_TYPE,
		  COLUMN_DEFAULT
		FROM
		  INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_SCHEMA = '%s'
	  `,
			dbName,
		),
	)
	if err != nil {
		return nil, err
	}

	tableColumnResponses := make(map[string][]DbSchemaRowResult)

	defer rows.Close()
	for rows.Next() {
		var row DbSchemaRowResult
		err = rows.Scan(
			&row.Table,
			&row.Column,
			&row.DataType,
			&row.ColumnKey,
			&row.OrdinalPosition,
			&row.IsNullable,
			&row.ColumnType,
			&row.ColumnDefault,
		)
		if err != nil {
			return nil, err
		}

		table := row.Table
		row.Table = ""

		isTableIgnored := false
		for _, ignoreTable := range GetGenerateModelConfig().Database.IgnoreTables {
			if strings.Contains(table, ignoreTable) {
				isTableIgnored = true
			}
		}

		if isTableIgnored {
			continue
		}

		tableColumnResponses[table] = append(tableColumnResponses[table], row)
	}

	schemas = tableColumnResponses

	return tableColumnResponses, err
}

// GetTableNames retrieves the names of all tables in the database
func (c *DbSchemaLookup) GetTableNames() ([]string, error) {
	schemas, err := c.GetSchemas()
	if err != nil {
		return nil, err
	}

	var tableNames []string
	for tableName := range schemas {
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}

// GetSchemasByTableName retrieves the database schema for a given table name
func (c *DbSchemaLookup) GetSchemasByTableName(tableName string) ([]DbSchemaRowResult, error) {
	schemas, err := c.GetSchemas()
	if err != nil {
		return nil, err
	}

	if _, ok := schemas[tableName]; !ok {
		return nil, fmt.Errorf("table %s not found", tableName)
	}

	return schemas[tableName], nil
}

type GenerateModelConfig struct {
	Database struct {
		IgnoreTables     []string            `yaml:"ignore_tables"`
		TableConnections map[string][]string `yaml:"table_connections"`
	} `yaml:"database"`
}

const generateConfig = "./.generate-model-config.yml"

// GetGenerateModelConfig loads the generate config from yaml file
func GetGenerateModelConfig() GenerateModelConfig {
	m := GenerateModelConfig{}

	config, err := os.ReadFile(generateConfig)
	if err != nil {
		log.Fatal(err)
	}

	// load yaml
	err = yaml.Unmarshal(config, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

// GetConnectionByTableName gets connection from config by table name
func GetConnectionByTableName(tableName string) string {
	m := GetGenerateModelConfig()

	for connection := range m.Database.TableConnections {
		for _, table := range m.Database.TableConnections[connection] {
			if table == tableName {
				return connection
			}
		}
	}

	return ""
}
