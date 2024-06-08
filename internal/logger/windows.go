//go:build windows
// +build windows

package logger

import (
	"golang.org/x/sys/windows"
	"os"
)

func init() {
	// Get the standard output handle
	stdout := windows.Handle(os.Stdout.Fd())

	// Retrieve current console mode
	var mode uint32
	windows.GetConsoleMode(stdout, &mode)

	// Enable virtual terminal processing
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	windows.SetConsoleMode(stdout, mode)
}
