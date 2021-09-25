package cmd

import (
	"fmt"
	"github.com/Akkadius/spire/generators"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strings"
)

type GenerateControllersCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (g *GenerateControllersCommand) Command() *cobra.Command {
	return g.command
}

func NewGenerateControllersCommand(db *gorm.DB, logger *logrus.Logger) *GenerateControllersCommand {
	i := &GenerateControllersCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "generate:controllers [all|table_name]",
			Short: "Generates application controller from template",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (g *GenerateControllersCommand) Handle(_ *cobra.Command, args []string) {
	//plural := pluralize.NewClient()

	g.writeTypesFile()

	for table, keys := range generators.GetDbSchemaKeysConfig() {
		if len(args) > 0 && table != args[0] {
			continue
		}

		for _, key := range keys {

			// todo: this needs to change later
			//isKey := strings.Contains(key.ColumnType, "int") && !strings.Contains(key.ColumnType, "unsigned")
			isKey := key.ColumnKey.String == "PRI" &&
				key.Column == "id" && strings.Contains(key.ColumnType, "int")
			if isKey {
				fmt.Printf("Generating models and controllers for [%v]\n", table)

				modelMeta := generators.NewGenerateModel(
					generators.GenerateModelContext{
						TablesToGenerate: []string{table},
					},
					g.logger,
					g.db,
				).Generate()

				relationships := ""
				for _, meta := range modelMeta {
					if meta.Table == table {
						relationships = strings.Join(meta.GormRelationships, "<br>")
					}
				}

				generators.NewGenerateController(
					generators.GenerateControllerContext{
						EntityName:           table,
						RelationshipsComment: relationships,
					},
					g.logger,
				).Generate()
			}
		}
	}

	generators.NewSyncControllersToInjector(g.logger).Sync()
}

func (g *GenerateControllersCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}

func (g *GenerateControllersCommand) writeTypesFile() {
	b, err := ioutil.ReadFile("./generators/templates/crud_controller_types.go")
	if err != nil {
		g.logger.Fatal(err)
	}

	// create file
	file, err := os.Create("./http/crudcontrollers/a_types.go")
	if err != nil {
		g.logger.Fatal(err)
	}

	defer file.Close()

	// write contents
	_, err = file.WriteString(string(b))
	if err != nil {
		g.logger.Fatal(err)
	}
}
