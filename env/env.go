package env

import (
	"os"
	"strconv"
)

// Get environment variable with a fallback (string)
// Example: Get("LOGGING_FORMAT", "text")
func Get(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}

// Get environment variable with a fallback (int)
// Example: GetInt("MYSQL_MAX_OPEN_CONNECTIONS", "150")
func GetInt(key, fallback string) int {
	val := Get(key, fallback)
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return ret
}

// Get environment variable with a fallback (bool)
// Example: GetBool("MYSQL_QUERY_LOGGING", "false")
func GetBool(key, fallback string) bool {
	val := Get(key, fallback)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		panic(err)
	}
	return ret
}
