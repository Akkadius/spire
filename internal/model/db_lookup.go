package model

import (
	"database/sql"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/volatiletech/null/v8"
)

type DbLookup struct {
	db     *sql.DB
	logger *logger.AppLogger
}

func NewDbLookup(db *sql.DB, logger *logger.AppLogger) *DbLookup {
	return &DbLookup{db: db, logger: logger}
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
func (c *DbLookup) GetSchemas() (map[string][]DbSchemaRowResult, error) {
	if schemas != nil {
		return schemas, nil
	}

	var dbName string
	err := c.db.QueryRow("SELECT database()").Scan(&dbName)
	if err != nil {
		c.logger.Fatal().Err(err).Msg("failed to get database name")
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

		tableColumnResponses[table] = append(tableColumnResponses[table], row)
	}

	schemas = tableColumnResponses

	return tableColumnResponses, err
}

// GetTableNames retrieves the names of all tables in the database
func (c *DbLookup) GetTableNames() ([]string, error) {
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
func (c *DbLookup) GetSchemasByTableName(tableName string) ([]DbSchemaRowResult, error) {
	schemas, err := c.GetSchemas()
	if err != nil {
		return nil, err
	}

	if _, ok := schemas[tableName]; !ok {
		return nil, fmt.Errorf("table %s not found", tableName)
	}

	return schemas[tableName], nil
}

// GetTableKeys retrieves the keys for a given table name
func (c *DbLookup) GetTableKeys(tableName string) ([]DbSchemaRowResult, error) {
	schemas, err := c.GetSchemas()
	if err != nil {
		return nil, err
	}

	if _, ok := schemas[tableName]; !ok {
		return nil, fmt.Errorf("table %s not found", tableName)
	}

	var keys []DbSchemaRowResult
	for _, schema := range schemas[tableName] {
		if schema.ColumnKey.String == "PRI" {
			keys = append(keys, schema)
		}
	}

	return keys, nil
}
