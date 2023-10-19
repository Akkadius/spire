package generators

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type GenerateConfig struct {
	Database struct {
		IgnoreTables     []string            `yaml:"ignore_tables"`
		TableConnections map[string][]string `yaml:"table_connections"`
	} `yaml:"database"`
}

const generateConfig = "./internal/generators/config/generate-config.yml"

func GetGenerateConfig() GenerateConfig {
	m := GenerateConfig{}

	config, err := ioutil.ReadFile(generateConfig)
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

// gets connection from config by table name
func GetConnectionByTableName(tableName string) string {
	m := GetGenerateConfig()

	for connection := range m.Database.TableConnections {
		for _, table := range m.Database.TableConnections[connection] {
			if table == tableName {
				return connection
			}
		}
	}

	return ""
}
