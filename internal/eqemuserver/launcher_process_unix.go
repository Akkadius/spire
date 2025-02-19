//go:build !windows
// +build !windows

package eqemuserver

import (
	"os"
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

	cmd.Dir = l.pathmgmt.GetEQEmuServerPath()
	if err := cmd.Start(); err != nil {
		return err
	}

	l.pollProcessMutex.Lock()
	l.currentProcessCounts[name]++
	l.pollProcessMutex.Unlock()

	// we do this otherwise we get a zombie process
	go func() {
		_ = cmd.Wait()
	}()

	return nil
}

func (l *Launcher) startLauncherProcess() error {
	// Get the path to the currently running executable
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	// Resolve any symbolic links to get the actual path
	exPath, err := filepath.EvalSymlinks(ex)
	if err != nil {
		return err
	}

	bin, err := exec.LookPath(exPath)
	if err != nil {
		return err
	}

	// copy the binary to os tmp dir
	tmpBin := filepath.Join(os.TempDir(), filepath.Base(bin)+"-launcher")
	if err := copyFile(bin, tmpBin); err != nil {
		return err
	}

	// chmod
	if err := os.Chmod(tmpBin, 0755); err != nil {
		return err
	}

	cmd := exec.Command(tmpBin, "eqemu-server:launcher", "start")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	cmd.Dir = l.pathmgmt.GetEQEmuServerPath()
	if err := cmd.Start(); err != nil {
		return err
	}

	// we do this otherwise we get a zombie process
	go func() {
		_ = cmd.Wait()
	}()

	return nil
}
