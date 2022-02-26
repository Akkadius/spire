package database

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
	"log"
)

type ShowColumns struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func GetColumnDefinitions(db *gorm.DB, tableName string) []ShowColumns {
	columnDefs := make([]ShowColumns, 0)
	db.Raw(fmt.Sprintf("SHOW COLUMNS FROM %v", tableName)).Scan(&columnDefs)

	return columnDefs
}

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

func GetTableColumnsOrdered(db *sql.DB, tableName string) []string {
	columns, err := GetTableSchema(db, tableName)
	if err != nil {
		log.Fatal(err)
	}

	names := []string{}
	for _, column := range columns {
		names = append(names, column.Column)
	}

	return names
}

func GetTableSchema(db *sql.DB, tableName string) ([]DbSchemaRowResult, error) {
	rows, err := db.Query(
		fmt.Sprintf(
			`
		SELECT
		  TABLE_NAME,
		  COLUMN_NAME,
		  DATA_TYPE,
		  COLUMN_KEY,
		  ORDINAL_POSITION,
		  IS_NULLABLE,
		  COLUMN_TYPE,
		  COLUMN_DEFAULT
		FROM
		  INFORMATION_SCHEMA.COLUMNS
		WHERE
		  TABLE_NAME = '%v'
		ORDER BY ORDINAL_POSITION
	  `, tableName,
		),
	)
	if err != nil {
		return nil, err
	}

	tableColumnResponses := make([]DbSchemaRowResult, 0)

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

		tableColumnResponses = append(tableColumnResponses, row)
	}

	return tableColumnResponses, nil
}
