package console

import "github.com/spf13/pflag"

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
