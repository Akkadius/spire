package database

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/volatiletech/null/v8"
	gormMysql "gorm.io/driver/mysql"
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

func GetTableColumnsOrdered(db *gorm.DB, tableName string) []string {
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

func GetTableSchema(db *gorm.DB, tableName string) ([]DbSchemaRowResult, error) {
	databaseName := ""
	f, ok := db.Dialector.(*gormMysql.Dialector)
	if ok {
		p, err := mysql.ParseDSN(f.DSN)
		if err != nil {
			return nil, err
		}

		databaseName = p.DBName
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	rows, err := sqlDb.Query(
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
		  AND TABLE_SCHEMA = '%v'
		ORDER BY ORDINAL_POSITION
	  `,
			tableName,
			databaseName,
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
