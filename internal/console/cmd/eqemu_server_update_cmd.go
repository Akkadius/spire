package cmd

import (
	"fmt"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/hashicorp/go-version"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type EQEmuServerUpdateCommand struct {
	logger         *logrus.Logger
	command        *cobra.Command
	serverconfig   *eqemuserverconfig.Config
	settings       *spire.Settings
	pathmgmt       *pathmgmt.PathManagement
	processmanager *eqemuserver.ProcessManager
	updater        *eqemuserver.Updater
}

func (c *EQEmuServerUpdateCommand) Command() *cobra.Command {
	return c.command
}

var useReleaseBinaries bool

// NewEQEmuServerUpdateCommand creates a new spire:init command
func NewEQEmuServerUpdateCommand(
	logger *logrus.Logger,
	serverconfig *eqemuserverconfig.Config,
	settings *spire.Settings,
	pathmgmt *pathmgmt.PathManagement,
	processmanager *eqemuserver.ProcessManager,
	updater *eqemuserver.Updater,

) *EQEmuServerUpdateCommand {
	i := &EQEmuServerUpdateCommand{
		logger:         logger,
		serverconfig:   serverconfig,
		settings:       settings,
		pathmgmt:       pathmgmt,
		processmanager: processmanager,
		updater:        updater,
		command: &cobra.Command{
			Use:   "eqemu-server:update",
			Short: "Extends Spire Admin server update functionality",
		},
	}

	// add flag --compile-server
	// add flag --compile-build-location
	// add flag --use-release-binaries
	i.command.PersistentFlags().BoolVarP(&useReleaseBinaries, "release-binaries", "r", false, "Whether or not to use release binaries")
	i.command.PersistentFlags().BoolVarP(&compileServer, "compile-server", "c", false, "Whether or not the server should be compiled (default: false - uses releases)")
	i.command.PersistentFlags().StringVarP(&compileLocation, "compile-build-location", "l", "~/code/build", "Determines where the binaries should be built")

	i.command.Run = i.Handle

	return i
}

func (c *EQEmuServerUpdateCommand) Handle(_ *cobra.Command, args []string) {
	if useReleaseBinaries {
		fmt.Println("Setting to use release binaries")
		c.settings.SetSetting(spire.SettingUpdateType, spire.UpdateTypeRelease)
	}
	if compileServer {
		fmt.Println("Setting to compile server")
		c.settings.SetSetting(spire.SettingUpdateType, spire.UpdateTypeSelfCompiled)

		info, err := c.updater.GetBuildInfo()
		if err != nil {
			c.logger.Fatal(err)
		}

		// set build cores if not set
		if info.BuildCores == "" {
			// default to 1 core
			cores := 1

			// get system memory available
			memory, err := mem.VirtualMemory()
			if err != nil {
				c.logger.Fatal(err)
			}

			// get system memory available in GB
			memoryAvailableGb := memory.Available / 1024 / 1024 / 1024
			if memoryAvailableGb >= 10 {
				cores = runtime.NumCPU() - 4
				if cores < 1 {
					cores = 1
				}
			} else if memoryAvailableGb >= 6 {
				cores = 4
			}

			c.settings.SetSetting(spire.SettingBuildCores, strconv.Itoa(cores))
		}

		if compileLocation != "" {
			c.logger.Info("Setting compile location: " + compileLocation)

			c.settings.SetSetting(spire.SettingBuildLocation, compileLocation)
		}
	}

	// check if server is online
	online := false
	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, _ := p.Cmdline()
		if strings.Contains(cmdline, "world") {
			online = true
		}
	}

	padSize := 15

	// print current settings
	// string repeat -
	c.LineBreak()
	fmt.Println(" Current settings")
	c.LineBreak()
	updateType := c.updater.GetUpdateType()
	fmt.Printf(" %*s | %v\n", padSize, "Update Type", updateType)

	if updateType == spire.UpdateTypeSelfCompiled {
		fmt.Printf(" %*s | %v\n", padSize, "Build Location", c.updater.GetBuildLocation())
	}
	info, _ := c.updater.GetVersionInfo()

	fmt.Printf(
		" %*s | %v\n",
		padSize,
		"Current",
		info.ServerVersion+" ("+info.CompileDate+" "+info.CompileTime+")",
	)

	c.LineBreak()

	// handle release updates
	if updateType == spire.UpdateTypeRelease {
		latestVer, err := c.updater.GetLatestReleaseVersion()
		if err != nil {
			c.logger.Fatal(err)
		}

		fmt.Printf(" %*s | %v\n", padSize, "Latest Release", latestVer)

		c.LineBreak()

		localVersion, err := version.NewVersion(info.ServerVersion)
		if err != nil {
			c.logger.Fatal(err)
		}

		latestVersion, err := version.NewVersion(latestVer)
		if err != nil {
			c.logger.Fatal(err)
		}

		if localVersion.GreaterThanOrEqual(latestVersion) {
			fmt.Println("Already up to date")
			os.Exit(0)
		} else if online {
			fmt.Println("Server is online, in order to update you need to stop the server first")
			// prompt user to stop server y/n
			// if yes, stop server
			var answer string
			fmt.Printf("Stop server? [y/n] ")
			fmt.Scanln(&answer)
			fmt.Printf("\n")

			if answer == "y" {
				fmt.Println("Stopping server")
				c.processmanager.Stop()
			}
		}

		// update server
		if localVersion.LessThan(latestVersion) {
			fmt.Println("Updating server")
			release := fmt.Sprintf(
				"https://github.com/EQEmu/Server/releases/download/v%s/eqemu-server-linux-x64.zip",
				latestVersion.String(),
			)
			err := c.updater.InstallRelease(release)
			if err != nil {
				c.logger.Fatal(err)
			}

			var answer string
			fmt.Printf("Start server? [y/n] ")
			fmt.Scanln(&answer)
			fmt.Printf("\n")

			if answer == "y" {
				fmt.Println("Starting server")
				c.processmanager.Start()
			}
		} else {
			fmt.Println("Server is already up to date")
		}
	}

	// handle self compiled
	if updateType == spire.UpdateTypeSelfCompiled {
		if runtime.GOOS != "linux" {
			c.logger.Fatal("Self compiled updates are only supported on linux")
		}

		fmt.Println("Starting")
		c.LineBreak()

		info, _ := c.updater.GetBuildInfo()

		fmt.Println("git pull")
		c.LineBreak()

		cmd := exec.Command("git", "pull")
		cmd.Stdout = os.Stdout
		cmd.Dir = info.SourceDirectory
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}

		c.LineBreak()
		fmt.Println(info.BuildTool, "-j"+info.BuildCores)
		c.LineBreak()

		cmd = exec.Command(info.BuildTool, "-j"+info.BuildCores)
		cmd.Dir = info.SourceDirectory
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			log.Println(err)
		}

		c.LineBreak()
	}
}

func (c *EQEmuServerUpdateCommand) LineBreak() {
	fmt.Println(strings.Repeat("-", 40))
}
