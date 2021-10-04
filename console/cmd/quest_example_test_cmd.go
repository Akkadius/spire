package cmd

import (
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/questapi"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

type QuestExampleTestCommand struct {
	logger  *logrus.Logger
	command *cobra.Command
	sourcer *questapi.QuestExamplesGithubSourcer
}

func (c *QuestExampleTestCommand) Command() *cobra.Command {
	return c.command
}

func NewQuestExampleTestCommand(
	logger *logrus.Logger,
	sourcer *questapi.QuestExamplesGithubSourcer,
) *QuestExampleTestCommand {
	i := &QuestExampleTestCommand{
		logger:  logger,
		sourcer: sourcer,
		command: &cobra.Command{
			Use:   "quest:example-test",
			Short: "Parses ProjectEQ repo for examples",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *QuestExampleTestCommand) Handle(_ *cobra.Command, args []string) {
	searchTerms := []string{}
	if len(args) > 0 {
		searchTerms = strings.Split(args[0], ",")
	}

	language := args[1]

	// inputs
	org := "ProjectEQ"
	repo := "projecteqquests"
	branch := "master"

	result := c.sourcer.Search(org, repo, branch, searchTerms, language, true)

	fmt.Println(result)
}

func (c *QuestExampleTestCommand) Validate(_ *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("args required [methods] [language]")
	}

	return nil
}
