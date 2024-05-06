//go:build !windows
// +build !windows

package eqemuserver

import (
	"os/exec"
	"path/filepath"
	"syscall"
)

func (l *Launcher) startServerProcess(name string, args ...string) error {
	bin, err := exec.LookPath(filepath.Join(l.pathmgmt.GetEQEmuServerPath(), "bin", name))
	if err != nil {
		return err
	}
	cmd := exec.Command(bin, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	// windows HideWindow: true

	cmd.Dir = l.pathmgmt.GetEQEmuServerPath()
	if err := cmd.Start(); err != nil {
		return err
	}

	err = cmd.Process.Release()
	if err != nil {
		return err
	}

	return nil
}
