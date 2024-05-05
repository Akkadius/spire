package logger

import (
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// DebugLogger represents a debug logger.
type DebugLogger struct {
	logger     *zerolog.Logger // The logger.
	debugLevel int             // The debug level. 1,2,3
}

// NewDebugLogger returns a new DebugLogger.
func NewDebugLogger() *DebugLogger {
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
	log.Logger = logger

	debugLogger := DebugLogger{
		logger:     &logger,
		debugLevel: 0,
	}

	return &debugLogger
}

// Debug is -v level logging
func (l *DebugLogger) Debug() *zerolog.Event {
	if l.debugLevel >= 1 {
		return l.logger.Debug()
	}
	return nil
}

// DebugVv is -vv level logging
func (l *DebugLogger) DebugVv() *zerolog.Event {
	if l.debugLevel >= 2 {
		return l.logger.Debug()
	}
	return nil
}

// DebugVvv is -vvv level logging
func (l *DebugLogger) DebugVvv() *zerolog.Event {
	if l.debugLevel >= 3 {
		return l.logger.Debug()
	}
	return nil
}

// SetDebugLevel sets the debug level (passed in from -v flags)
func (l *DebugLogger) SetDebugLevel(level int) {
	l.debugLevel = level
}
