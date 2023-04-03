package cmd

import (
	"fmt"
	"github.com/Akkadius/spire/internal/backup"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

type HelloWorldCommand struct {
	db       *gorm.DB
	logger   *logrus.Logger
	command  *cobra.Command
	backup   *backup.Mysql
	pathmgmt *pathmgmt.PathManagement
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(
	db *gorm.DB,
	logger *logrus.Logger,
	backup *backup.Mysql,
	pathmgmt *pathmgmt.PathManagement,
) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db:     db,
		logger: logger,
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
	pp.Println(c.pathmgmt.GetLogsDirPath())

	var logFile string
	err := filepath.Walk(c.pathmgmt.GetLogsDirPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, ".log") {
			logFile = path
			return nil
		}

		//fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	pp.Println(logFile)

	f, err := os.Open(logFile)
	if err != nil {
		pp.Println(err)
	}

	b1 := make([]byte, 100*1024*1024)
	n1, err := f.Read(b1)
	if err != nil {
		pp.Println(err)
	}

	pp.Println(n1)

}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
