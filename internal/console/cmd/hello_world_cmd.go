package cmd

import (
	"github.com/Akkadius/spire/internal/backup"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
	"strings"
)

type HelloWorldCommand struct {
	db       *gorm.DB
	command  *cobra.Command
	backup   *backup.Mysql
	pathmgmt *pathmgmt.PathManagement
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(
	db *gorm.DB,
	backup *backup.Mysql,
	pathmgmt *pathmgmt.PathManagement,
) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db: db,
		command: &cobra.Command{
			Use:   "hello:hello-world",
			Short: "Says hello world",
		},
		backup:   backup,
		pathmgmt: pathmgmt,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, args []string) {
	// open races.txt
	// read the file
	file := "races.txt"
	contents, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	startListeningMaleHeight := false
	startListeningFemaleHeight := false
	startListeningRaceSizes := false
	maleHeightEntries := make([]string, 0)
	femaleHeightEntries := make([]string, 0)
	raceSizesIndex := 0

	for _, s := range strings.Split(string(contents), "\n") {
		if startListeningMaleHeight {
			if strings.Contains(s, ",") {
				entries := strings.Split(s, ",")
				for _, entry := range entries {
					e := strings.TrimSpace(entry)
					if e == "" {
						continue
					}
					maleHeightEntries = append(maleHeightEntries, e)
				}
			}
			if strings.Contains(s, "}") {
				startListeningMaleHeight = false
			}
		}
		if startListeningFemaleHeight {
			if strings.Contains(s, ",") {
				entries := strings.Split(s, ",")
				for _, entry := range entries {
					e := strings.TrimSpace(entry)
					if e == "" {
						continue
					}
					femaleHeightEntries = append(femaleHeightEntries, e)
				}
			}
			if strings.Contains(s, "}") {
				startListeningFemaleHeight = false
			}
		}
		if startListeningRaceSizes {
			// 	{ Race::Human,                 { 7.0f,   7.0f }},
			// pull both numeric values
			if strings.Contains(s, "Race::") {
				s = strings.ReplaceAll(s, "{", "")
				s = strings.ReplaceAll(s, "}", "")
				splitValues := strings.Split(s, ",")

				race := strings.TrimSpace(splitValues[0])
				male := strings.TrimSpace(splitValues[1])
				female := strings.TrimSpace(splitValues[2])

				origMale := maleHeightEntries[raceSizesIndex]
				origFemale := femaleHeightEntries[raceSizesIndex]

				if origMale != male {
					pp.Printf(
						"race [%v] (%v) origMale does not match male [%v] -> [%v]\n",
						race,
						raceSizesIndex+1,
						origMale,
						male,
					)
				}

				if origFemale != female {
					pp.Printf(
						"race [%v] (%v) origFemale does not match female [%v] -> [%v]\n",
						race,
						raceSizesIndex+1,
						origFemale,
						female,
					)
				}

				raceSizesIndex++
			}

			if strings.Contains(s, "}") {
				startListeningRaceSizes = false
			}
		}

		if strings.Contains(s, "male_height") {
			startListeningMaleHeight = true
		}

		if strings.Contains(s, "female_height") {
			startListeningFemaleHeight = true
		}
		if strings.Contains(s, "race_sizes") {
			startListeningRaceSizes = true
		}
	}

	//pp.Println(maleHeightEntries)
	//pp.Println(femaleHeightEntries)

}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
