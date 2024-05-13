package generators

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/volatiletech/null/v8"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type DbSchemaConfig struct {
	db     *sql.DB
	logger *logger.AppLogger
}

func NewDbSchemaConfig(db *sql.DB, logger *logger.AppLogger) *DbSchemaConfig {
	return &DbSchemaConfig{db: db, logger: logger}
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
func (c *DbSchemaConfig) Generate(dbName string) error {
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
		c.logger.Fatal().Err(err).Msg("Failed to create file")
	}

	defer file.Close()

	// write config
	_, err = file.Write(b.Bytes())
	if err != nil {
		c.logger.Fatal().Err(err).Msg("Failed to create file")
	}

	c.logger.Info().Any("dbSchemaConfig", dbSchemaConfig).Msg("Wrote configuration")

	return nil
}

func (c *DbSchemaConfig) GenerateKeys() error {
	keys, err := getDbTableKeys()
	if err != nil {
		c.logger.Fatal().Err(err).Msg("Failed to get db table keys")
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
		c.logger.Fatal().Err(err).Msg("Failed to write file")
	}

	c.logger.Info().Any("dbSchemaKeysConfig", dbSchemaKeysConfig).Msg("Wrote configuration")

	return nil
}

var dbSchemaConfigCached = map[string][]DbSchemaRowResult{}

// get db schema config
func GetDbSchemaConfig() map[string][]DbSchemaRowResult {
	if len(dbSchemaConfigCached) != 0 {
		return dbSchemaConfigCached
	}

	m := map[string][]DbSchemaRowResult{}

	config, err := os.ReadFile(dbSchemaConfig)
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

	config, err := os.ReadFile(dbSchemaKeysConfig)
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
	var tables []string
	for table := range GetDbSchemaConfig() {
		// skip spire tables
		if strings.Contains("spire_", table) {
			continue
		}

		tables = append(tables, table)
	}
	return tables
}
