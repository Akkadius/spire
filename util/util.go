package util

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/go-errors/errors"
	"github.com/spf13/pflag"
)

func RecoveryToError(r interface{}) error {
	var err error
	switch value := r.(type) {
	case string:
		err = errors.New(value)
	case error:
		err = value
	default:
		err = errors.New("unknown panic")
	}

	return err
}

func ParseJSON(v interface{}) (*gabs.Container, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return gabs.ParseJSON(b)
}

func StringToJSON(v string) (*gabs.Container, error) {
	return gabs.ParseJSON([]byte(v))
}

func ErrorMatch(err error, v string) bool {
	return strings.Contains(err.Error(), v)
}

func StringFromFlags(flags *pflag.FlagSet, key string) string {
	value, _ := flags.GetString(key)

	return value
}

func StringSliceFromFlags(flags *pflag.FlagSet, key string) []string {
	values, _ := flags.GetStringSlice(key)

	return values
}

func UintFromFlags(flags *pflag.FlagSet, key string) uint {
	value, _ := flags.GetUint(key)

	return value
}

func Uint64FromFlags(flags *pflag.FlagSet, key string) uint64 {
	value, _ := flags.GetUint64(key)

	return value
}

func Int64FromFlags(flags *pflag.FlagSet, key string) int64 {
	value, _ := flags.GetInt64(key)

	return value
}

func InStringSlice(needle string, haystack []string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}

	return false
}

// StringToInt converts a string to an integer
// defaulting to 0 if the string cannot
// be converted to an int
func StringToInt(value string) int {
	v, _ := strconv.Atoi(value)

	return v
}
