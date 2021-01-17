package console

import (
	"fmt"
	"github.com/muesli/termenv"
	"strings"
)

// internal line command
func line(msg string, a ...interface{}) string {
	if len(a) > 0 {
		return fmt.Sprintf(msg+"\n", a...)
	}
	return fmt.Sprintf(msg + "\n")
}

// internal message no newline
func message(msg string, a ...interface{}) string {
	if len(a) > 0 {
		return fmt.Sprintf(msg, a...)
	}
	return fmt.Sprintf(msg)
}

// console colors
const (
	colorWhite  = "#FFFFFF"
	colorYellow = "#00ff00"
	colorOrange = "#FFA500"
	colorRed    = "#FF0000"
)

// console info
func Info(msg string, a ...interface{}) {
	fmt.Print(termenv.String(line(msg, a...)).Foreground(termenv.ColorProfile().Color(colorYellow)).String())
}

// console line
func Line(msg string, a ...interface{}) {
	fmt.Print(termenv.String(line(msg, a...)).Foreground(termenv.ColorProfile().Color(colorWhite)).String())
}

// console warn
func Warn(msg string, a ...interface{}) {
	fmt.Print(termenv.String(line(msg, a...)).Foreground(termenv.ColorProfile().Color(colorOrange)).String())
}

// console error
func Error(msg string, a ...interface{}) {
	p := termenv.ColorProfile()
	fmt.Print(termenv.String(message("\n\n  [ERROR] "+msg+"\n", a...)).
		Background(p.Color(colorRed)).
		Foreground(p.Color(colorWhite)).
		Bold().
		String())
	termenv.Reset()
	fmt.Print("\n\n")
}

func PrintBanner(msg string, bannerLength int) {
	fmt.Println(strings.Repeat("-", bannerLength))
	fmt.Println(fmt.Sprintf("| %v", msg))
	fmt.Println(strings.Repeat("-", bannerLength))
}
