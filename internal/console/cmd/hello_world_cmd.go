package cmd

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"strings"
)

type HelloWorldCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(db *gorm.DB, logger *logrus.Logger) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "hello:hello-world",
			Short: "Says hello world",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, args []string) {
	resp, err := soup.Get("https://everquest.allakhazam.com/db/zone.html?ztype=0")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(resp)

	zone := ""
	if len(args) > 0 {
		zone = args[0]
	}
	images := []string{}
	doc := soup.HTMLParse(resp)
	for _, link := range doc.FindAll("a") {
		if strings.Contains(link.Attrs()["href"], "zstrat") && strings.Contains(link.Text(), zone) {
			fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])

			// grab zone level
			r, err := soup.Get("https://everquest.allakhazam.com" + link.Attrs()["href"])
			if err != nil {
				fmt.Println(err)
			}

			d := soup.HTMLParse(r)
			for _, l := range d.FindAll("a") {
				if strings.Contains(l.Attrs()["href"], "scenery") {
					images = append(images, l.Attrs()["href"])
				}
			}
		}
	}

	fmt.Println(images)
}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
