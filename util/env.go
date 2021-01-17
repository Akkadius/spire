package util

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Get environment variable with a fallback (string)
// Example: GetEnv("LOGGING_FORMAT", "text")
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

// Get environment variable with a fallback (int)
// Example: GetIntEnv("MYSQL_MAX_OPEN_CONNECTIONS", "150")
func GetIntEnv(key, fallback string) int {
	val := GetEnv(key, fallback)
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return ret
}

// Get environment variable with a fallback (bool)
// Example: GetBoolEnv("MYSQL_QUERY_LOGGING", "false")
func GetBoolEnv(key, fallback string) bool {
	val := GetEnv(key, fallback)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		panic(err)
	}
	return ret
}

// EnvMaxDirectorySeekLevels is the number of directory
// levels a .env file needs to be searched in
const EnvMaxDirectorySeekLevels int = 10

var envLoaded = false

// loads environment file .env locally
// loads .env.testing if invoked from the context of a test file
// loads .env.debug.host if invoked from the context of MacOS which references variables to communicate back to the docker network
func LoadEnvFileIfExists() error {
	if runtime.GOOS == "darwin" {
		_ = os.Setenv("APP_ENV", "local")
	}

	if IsAppEnvLocalOrTesting() && !envLoaded {
		// use dev or testing envs depending on the environment
		envLoadMsg := ""
		envFile := ""
		if IsAppEnvLocal() {
			envFile = ".env"
		}
		if IsAppEnvTesting() {
			envFile = ".env.testing"
		}

		// load top-level .env
		if loadEnvFile(envFile) {
			envLoadMsg = fmt.Sprintf("[LoadEnv] APP_ENV [%v] ENV_FILE [%v]", os.Getenv("APP_ENV"), envFile)
		}

		// display env: [LoadEnv] APP_ENV [local] ENV_FILE [.env]
		fmt.Println(envLoadMsg)
		envLoaded = true

		// search for global .env.debug.host
		// if OS is darwin; we're likely talking from host -> container network
		// used from IDEs
		if runtime.GOOS == "darwin" {
			env := ".env.debug.host"
			if loadEnvFile(env) {
				fmt.Println(fmt.Sprintf("[LoadEnv] (Host) APP_ENV [%v] ENV_FILE [%v]", os.Getenv("APP_ENV"), env))
			}
		}
	}

	return nil
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
	AppEnvStaging    = "staging"
	AppEnvProduction = "production"
)

func IsAppEnvLocal() bool {
	return os.Getenv("APP_ENV") == AppEnvLocal
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
