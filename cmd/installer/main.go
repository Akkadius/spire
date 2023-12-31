package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"log"
	"os"
	"time"
)

func main() {
	// default
	_ = os.Setenv("APP_ENV", "local")

	// load env
	if err := env.LoadEnvFileIfExists(); err != nil {
		Fatal(err)
	}

	// installer logic
	// check if eqemu_config.json exists in current directory
	// if it doesn't exist, run the eqemu server installer
	err := eqemuserver.NewInstaller().Install()
	if err != nil {
		fmt.Printf("\nInstallation error: %s\n\n", err)
		fmt.Print("Press [Enter] to continue...")
		_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}
}

func Fatal(err error) {
	log.Println(err)

	// only hang if we're not running a command
	if len(os.Args) == 1 {
		fmt.Print("Automatically shutting down in 10 seconds...")
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}
}
