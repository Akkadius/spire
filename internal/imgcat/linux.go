//go:build linux
// +build linux

package imgcat

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func disableEcho() *unix.Termios {
	termios, err := unix.IoctlGetTermios(int(os.Stdout.Fd()), unix.TCGETS)
	if err != nil {
		log.Fatalf("failed to get the termios: %v", err)
	}

	newState := *termios
	newState.Lflag &^= unix.ECHO
	newState.Lflag |= unix.ICANON | unix.ISIG
	newState.Iflag |= unix.ICRNL
	if err := unix.IoctlSetTermios(int(os.Stdout.Fd()), unix.TCSETS, &newState); err != nil {
		log.Fatalf("failed to set the termios: %v", err)
	}

	return termios
}

func enableEcho(termios *unix.Termios) {
	if err := unix.IoctlSetTermios(int(os.Stdout.Fd()), unix.TCSETS, termios); err != nil {
		log.Fatalf("failed to set the termios: %v", err)
	}
}
