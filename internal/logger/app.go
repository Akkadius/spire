package logger

import (
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// AppLogger represents a debug logger.
type AppLogger struct {
	debugLogger *zerolog.Logger // The debugLogger.
	infoLogger  *zerolog.Logger // The infoLogger.
	debugLevel  int             // The debugLogger level. 1,2,3
}

// NewAppLogger returns a new AppLogger.
func NewAppLogger() *AppLogger {
	appLogger := AppLogger{
		debugLogger: newDebugLogger(),
		infoLogger:  newInfoLogger(),
		debugLevel:  0,
	}

	return &appLogger
}

func newDebugLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	output.FormatLevel = func(i interface{}) string {
		return ""
	}
	output.FormatMessage = func(i interface{}) string {
		callerMeta := getCallerMeta()
		filename := filepath.Base(os.Args[0])
		appMode := os.Getenv("APP_MODE")
		return fmt.Sprintf(
			"   %s%s%s%s %s(%s) (%s)%s",
			console.HighIntensityBlack,
			callerMeta,
			i,
			console.Reset,
			console.FadedGray,
			filename,
			appMode,
			console.Reset,
		)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("\n      %s> %s%s ", console.HighIntensityBlack, i, console.Reset)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s%s%s ", console.HighIntensityGreen, i, console.Reset)
	}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}

	logger := zerolog.New(output)
	return &logger
}

func newInfoLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	output.FormatLevel = func(i interface{}) string {
		return ""
	}
	output.FormatMessage = func(i interface{}) string {
		callerMeta := getCallerMeta()
		appMode := os.Getenv("APP_MODE")

		return fmt.Sprintf(
			"%sSpire › %s%s%s%s (%s)%s",
			console.BoldWhite,
			callerMeta,
			console.Reset,
			console.White,
			i,
			appMode,
			console.Reset,
		)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s> %s%s ", console.HighIntensityBlack, i, console.Reset)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s%s%s", console.HighIntensityGreen, i, console.Reset)
	}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}

	logger := zerolog.New(output)
	return &logger
}

// getCallerMeta returns the caller type and package
// Example: QuestHotReloadWatcher (eqemuserver) ›
func getCallerMeta() string {
	pc := make([]uintptr, 20) // adjust the number of frames to retrieve
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])

	callerType := ""
	callerPackage := ""
	for {
		frame, more := frames.Next()
		if strings.Contains(frame.Function, "log") {
			continue
		}

		//fmt.Printf("- %s\n", frame.Function)
		if !more {
			break
		}

		pkg := frame.Function
		if strings.Contains(pkg, "(*") {
			callerType = pkg

			// extract type from github.com/Akkadius/spire/internal/eqemuserver.(*QuestHotReloadWatcher)
			// to QuestHotReloadWatcher
			split := strings.Split(pkg, "(*")
			if len(split) > 1 {
				callerType = split[1]
				callerType = strings.TrimSuffix(callerType, ")")
				callerType = strings.TrimSpace(callerType)
				callerType = strings.ReplaceAll(callerType, ")", "")

				// get package
				callerSplit := strings.Split(split[0], "/")
				if len(callerSplit) > 0 {
					callerPackage = callerSplit[len(callerSplit)-1]
					callerPackage = strings.ReplaceAll(callerPackage, ".", "")
				}
			}

			break
		}
	}

	var callerMeta string
	if callerType != "" {
		callerMeta = fmt.Sprintf("%s (%s) › ", callerType, callerPackage)
	}

	return callerMeta
}

//func (l *AppLogger) Logger() *zerolog.Logger {
//	return l.infoLogger
//}

func (l *AppLogger) GetWriter() zerolog.Logger {
	return l.infoLogger.With().Caller().Logger()
}

func (l *AppLogger) Info() *zerolog.Event {
	return l.infoLogger.Info()
}

func (l *AppLogger) Error() *zerolog.Event {
	return l.infoLogger.Error()
}

func (l *AppLogger) Fatal() *zerolog.Event {
	return l.infoLogger.Fatal()
}

func (l *AppLogger) Warn() *zerolog.Event {
	return l.infoLogger.Warn()
}

// Debug is -v level logging
func (l *AppLogger) Debug() *zerolog.Event {
	if l.debugLevel >= 1 {
		return l.debugLogger.Debug()
	}
	return nil
}

// DebugVv is -vv level logging
func (l *AppLogger) DebugVv() *zerolog.Event {
	if l.debugLevel >= 2 {
		return l.debugLogger.Debug()
	}
	return nil
}

// DebugVvv is -vvv level logging
func (l *AppLogger) DebugVvv() *zerolog.Event {
	if l.debugLevel >= 3 {
		return l.debugLogger.Debug()
	}
	return nil
}

// SetDebugLevel sets the debug level (passed in from -v flags)
func (l *AppLogger) SetDebugLevel(level int) {
	l.debugLevel = level
}
