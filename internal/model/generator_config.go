package model

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// GenerateModelConfig represents the configuration for generating models
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
