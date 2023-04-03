package env

import (
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"runtime"
	"strings"
)

// EnvMaxDirectorySeekLevels is the number of directory
// levels a .env file needs to be searched in
const EnvMaxDirectorySeekLevels int = 10

var envLoaded = false

// LoadEnvFileIfExists loads environment file .env locally
// loads .env.testing if invoked from the context of a test file
// loads .env.debug.host if invoked from the context of MacOS which references variables to communicate back to the docker network
func LoadEnvFileIfExists() error {
	if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
		_ = os.Setenv("APP_ENV", "local")
	}

	if !envLoaded {
		// use dev or testing envs depending on the environment
		envLoadMsg := ""
		envFile := ".env"
		if IsAppEnvTesting() {
			envFile = ".env.testing"
		}

		// load top-level .env
		if loadEnvFile(envFile) {
			envLoadMsg = fmt.Sprintf("[LoadEnv] APP_ENV [%v] ENV_FILE [%v]", os.Getenv("APP_ENV"), envFile)
		}

		// display env: [LoadEnv] APP_ENV [local] ENV_FILE [.env]
		if GetInt("DEBUG", "0") >= 3 {
			fmt.Println(envLoadMsg)
		}

		envLoaded = true

		// check if running in docker
		isInDocker := true
		if _, err := os.Stat("/.dockerenv"); errors.Is(err, os.ErrNotExist) {
			isInDocker = false
		}

		// search for global .env.debug.host
		// we're likely talking from host -> container network
		// used from IDEs
		if !isInDocker {
			env := ".env.debug.host"
			if loadEnvFile(env) {
				if len(os.Getenv("DEBUG")) > 0 {
					fmt.Println(fmt.Sprintf("[LoadEnv] (Host) APP_ENV [%v] ENV_FILE [%v]", os.Getenv("APP_ENV"), env))
				}
			}
		}
	}

	return nil
}

func IsEnvLoaded() bool {
	return envLoaded
}

// searches for .env file through directory traversal
// loads .env file if found
func loadEnvFile(envFile string) bool {
	var path string
	found := false
	for i := 0; i < EnvMaxDirectorySeekLevels; i++ {
		if _, err := os.Stat(path + envFile); err == nil {
			path += envFile
			found = true
			break
		}
		path += "../"
	}

	if found {
		if err := godotenv.Overload(path); err != nil {
			panic(err)
		}
	}

	return found
}

// environment helpers
const (
	AppEnvTesting    = "testing"
	AppEnvLocal      = "local"
	AppEnvDev        = "dev" // effectively the same as (local)
	AppEnvDesktop    = "desktop"
	AppEnvStaging    = "staging"
	AppEnvProduction = "production"
)

func IsAppEnvDev() bool {
	return os.Getenv("APP_ENV") == AppEnvDev
}

func IsAppEnvLocal() bool {
	return os.Getenv("APP_ENV") == AppEnvLocal ||
		os.Getenv("APP_ENV") == AppEnvDev ||
		os.Getenv("APP_ENV") == AppEnvDesktop
}

func IsAppEnvTesting() bool {
	return os.Getenv("APP_ENV") == AppEnvTesting ||
		flag.Lookup("test.v") != nil ||
		isTestSuffixFromArguments()
}

func IsAppEnvLocalOrTesting() bool {
	return IsAppEnvLocal() || IsAppEnvTesting()
}

func IsAppEnvStaging() bool {
	return os.Getenv("APP_ENV") == AppEnvStaging
}

func IsAppEnvProduction() bool {
	return os.Getenv("APP_ENV") == AppEnvProduction
}

// IsHostedReadOnlyModeEnabled is when the primary database connection becomes read only
// whereas under desktop or local modes the primary connection would always be write mode
// hosted spire has PEQ as a "view" of data and should not be able to be written to
func IsHostedReadOnlyModeEnabled() bool {
	return len(os.Getenv("IS_HOSTED_READ_ONLY_MODE")) > 0 &&
		os.Getenv("IS_HOSTED_READ_ONLY_MODE") != "0"
}

func IsAppEnvStagingOrProduction() bool {
	return IsAppEnvStaging() || IsAppEnvProduction()
}

func isTestSuffixFromArguments() bool {
	anyArgumentContainsTestSuffix := false

	for _, arg := range os.Args {
		if strings.HasSuffix(arg, ".test") || strings.HasSuffix(arg, "-test.run") {
			anyArgumentContainsTestSuffix = true
		}
	}

	return anyArgumentContainsTestSuffix
}
