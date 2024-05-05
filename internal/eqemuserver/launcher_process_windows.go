//go:build windows
// +build windows

package eqemuserver

import (
	"os/exec"
	"path/filepath"
	"syscall"
)

func (l *Launcher) startServerProcess(name string, args ...string) {
	bin := filepath.Join(l.pathmgmt.GetEQEmuServerPath(), "bin", name)
	cmd := exec.Command(bin, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true, HideWindow: true}
	// windows HideWindow: true

	cmd.Dir = l.pathmgmt.GetEQEmuServerPath()
	if err := cmd.Start(); err != nil {
		return // handle error
	}

	cmd.Process.Release()
}
