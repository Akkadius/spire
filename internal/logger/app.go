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
		filename := filepath.Base(os.Args[0])
		return fmt.Sprintf(
			"   %s%s%s %s(%s)%s",
			console.HighIntensityBlack,
			i,
			console.Reset,
			console.FadedGray,
			filename,
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

func getPackage(levelsUp int) string {
	// Retrieve the caller's function name
	pc, _, _, _ := runtime.Caller(levelsUp)
	fn := runtime.FuncForPC(pc)

	// Get the package name from the function name
	if fn != nil {
		fullFuncName := fn.Name()
		lastDot := strings.LastIndex(fullFuncName, ".")
		if lastDot > 0 {
			lastDot = strings.LastIndex(fullFuncName[:lastDot], ".")
			if lastDot > 0 {
				return fullFuncName[:lastDot]
			}
		}
	}

	return ""
}

func newInfoLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stderr}
	output.FormatLevel = func(i interface{}) string {
		return ""
	}
	output.FormatMessage = func(i interface{}) string {
		callerType := ""
		callerPackage := ""
		//now := time.Now()
		for i := 0; i < 20; i++ {
			pkg := getPackage(i)
			//pp.Println(getPackage(i))
			//fmt.Println("took ", time.Since(now).String())
			if strings.Contains(pkg, "(*") {
				callerType = pkg

				// extract type from github.com/Akkadius/spire/internal/eqemuserver.(*QuestHotReloadWatcher)
				// to QuestHotReloadWatcher
				split := strings.Split(pkg, "(*")
				if len(split) > 1 {
					callerType = split[1]
					callerType = strings.TrimSuffix(callerType, ")")
					callerType = strings.TrimSpace(callerType)

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

		return fmt.Sprintf(
			"%sSpire › %s%s%s%s%s",
			console.BoldWhite,
			callerMeta,
			console.Reset,
			console.White,
			i,
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

func (l *AppLogger) Info() *zerolog.Event {
	return l.infoLogger.Info()
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
