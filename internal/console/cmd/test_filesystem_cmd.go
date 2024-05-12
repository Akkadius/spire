package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

type TestFilesystemCommand struct {
	command  *cobra.Command
	pathmgmt *pathmgmt.PathManagement
}

func (c *TestFilesystemCommand) Command() *cobra.Command {
	return c.command
}

func NewTestFilesystemCommand(
	pathmgmt *pathmgmt.PathManagement,
) *TestFilesystemCommand {
	i := &TestFilesystemCommand{
		pathmgmt: pathmgmt,
		command: &cobra.Command{
			Use:   "test:filesystem",
			Short: "Development test command, tests filesystem",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *TestFilesystemCommand) Handle(_ *cobra.Command, args []string) {

	type ServerVersionInfo struct {
		BotsDatabaseVersion int    `json:"bots_database_version"`
		CompileDate         string `json:"compile_date"`
		CompileTime         string `json:"compile_time"`
		DatabaseVersion     int    `json:"database_version"`
		ServerVersion       string `json:"server_version"`
	}

	worldBin := c.pathmgmt.GetWorldBinPath()
	pp.Printf("World bin path is [%v]\n", worldBin)
	if _, err := os.Stat(worldBin); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Failed to find World binary to fetch version")
	}
	pp.Printf("World bin path exists [%v]\n", worldBin)

	pp.Printf("Running command [%v %v]\n", worldBin, "world:version")
	cmd := exec.Command(worldBin, "world:version")
	cmd.Dir = c.pathmgmt.GetEQEmuServerPath()
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// not all binaries simply output json alone
	// there was an output bug
	o := string(output)
	var n string
	startWatch := false
	for _, s := range strings.Split(o, "\n") {
		if strings.Contains(s, "{") {
			startWatch = true
		}
		if startWatch {
			n += s
		}
	}

	var v ServerVersionInfo
	err = json.Unmarshal([]byte(n), &v)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Server version")
	pp.Println(v)

}

func (c *TestFilesystemCommand) Validate(_ *cobra.Command, args []string) error {
	return nil
}
