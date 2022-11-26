package generators

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type GenerateDbSchemaConfig struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewGenerateDbSchemaConfig(db *sql.DB, logger *logrus.Logger) *GenerateDbSchemaConfig {
	return &GenerateDbSchemaConfig{db: db, logger: logger}
}

const (
	dbSchemaConfig     = "./internal/generators/config/db-schema.yml"
	dbSchemaKeysConfig = "./internal/generators/config/db-schema-keys.yml"
)

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

// Generate db schema config
func (c *GenerateDbSchemaConfig) Generate(dbName string) error {
	rows, err := c.db.Query(
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
		  TABLE_SCHEMA = '%v';
	  `, dbName,
		),
	)
	if err != nil {
		return err
	}

	tableColumnResponses := make(map[string][]DbSchemaRowResult, 0)

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
			return err
		}

		table := row.Table
		row.Table = ""

		isTableIgnored := false
		for _, ignoreTable := range GetGenerateConfig().Database.IgnoreTables {
			if strings.Contains(table, ignoreTable) {
				isTableIgnored = true
			}
		}

		if isTableIgnored {
			continue
		}

		tableColumnResponses[table] = append(tableColumnResponses[table], row)
	}

	var b bytes.Buffer
	config := yaml.NewEncoder(&b)
	config.SetIndent(2)

	if err := config.Encode(&tableColumnResponses); err != nil {
		return err
	}

	// create config
	file, err := os.Create(dbSchemaConfig)
	if err != nil {
		c.logger.Fatal(err)
	}

	defer file.Close()

	// write config
	_, err = file.Write(b.Bytes())
	if err != nil {
		c.logger.Fatal(err)
	}

	c.logger.Infof("Wrote configuration [%v]", dbSchemaConfig)

	return nil
}

func (c *GenerateDbSchemaConfig) GenerateKeys() error {
	keys, err := getDbTableKeys()
	if err != nil {
		c.logger.Fatal(err)
	}

	// yaml
	var b bytes.Buffer
	config := yaml.NewEncoder(&b)
	config.SetIndent(2)
	if err := config.Encode(&keys); err != nil {
		return err
	}

	// create config
	file, err := os.Create(dbSchemaKeysConfig)
	if err != nil {
		return err
	}

	defer file.Close()

	// write config
	_, err = file.Write(b.Bytes())
	if err != nil {
		c.logger.Fatal(err)
	}

	c.logger.Infof("Wrote configuration [%v]", dbSchemaKeysConfig)

	return nil
}

var dbSchemaConfigCached = map[string][]DbSchemaRowResult{}

// get db schema config
func GetDbSchemaConfig() map[string][]DbSchemaRowResult {
	if len(dbSchemaConfigCached) != 0 {
		return dbSchemaConfigCached
	}

	m := map[string][]DbSchemaRowResult{}

	config, err := ioutil.ReadFile(dbSchemaConfig)
	if err != nil {
		log.Fatal(err)
	}

	// load yaml
	err = yaml.Unmarshal(config, &m)
	if err != nil {
		log.Fatal(err)
	}

	dbSchemaConfigCached = m

	return m
}

// get db schema config for table
func GetDbSchemaConfigTable(table string) []DbSchemaRowResult {
	return GetDbSchemaConfig()[table]
}

// get db schema config for table
func getDbTableKeys() (map[string][]DbSchemaRowResult, error) {
	r := map[string][]DbSchemaRowResult{}
	for table, entries := range GetDbSchemaConfig() {
		for _, row := range entries {
			if row.ColumnKey.String == "PRI" {
				r[table] = append(r[table], row)
			}
		}
	}

	return r, nil
}

var dbSchemaConfigKeysCached = map[string][]DbSchemaRowResult{}

// get db schema keys config
func GetDbSchemaKeysConfig() map[string][]DbSchemaRowResult {
	if len(dbSchemaConfigKeysCached) != 0 {
		return dbSchemaConfigKeysCached
	}

	m := map[string][]DbSchemaRowResult{}

	config, err := ioutil.ReadFile(dbSchemaKeysConfig)
	if err != nil {
		log.Fatal(err)
	}

	// load yaml
	err = yaml.Unmarshal(config, &m)
	if err != nil {
		log.Fatal(err)
	}

	dbSchemaConfigKeysCached = m

	return m
}

// get db schema keys config for table
func GetDbSchemaKeysConfigTable(table string) []DbSchemaRowResult {
	return GetDbSchemaKeysConfig()[table]
}

// get database tables from config
func GetDatabaseTables() []string {
	tables := []string{}
	for table := range GetDbSchemaConfig() {
		// skip spire tables
		if strings.Contains("spire_", table) {
			continue
		}
		
		tables = append(tables, table)
	}
	return tables
}
