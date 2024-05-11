package logger

import (
	"github.com/Akkadius/spire/internal/env"
	"os"
	"strings"
)

func ProvideAppLogger() *AppLogger {
	l := NewAppLogger()

	// If the user supplies -v, -vv, -vvv, etc, set the logger verbosity.
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-v") {
			l.SetDebugLevel(strings.Count(arg, "v"))
		}
	}

	// If the user supplies DEBUG=1-3 to set the logger verbosity.
	if len(os.Getenv("DEBUG")) > 0 {
		l.SetDebugLevel(env.GetInt("DEBUG", "0"))
	}

	return l
}
