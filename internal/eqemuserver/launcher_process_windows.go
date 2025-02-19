//go:build windows
// +build windows

package eqemuserver

import (
	"github.com/Akkadius/spire/internal/filepathcheck"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

const (
	// Process creation flags.
	CREATE_BREAKAWAY_FROM_JOB        = 0x01000000
	CREATE_DEFAULT_ERROR_MODE        = 0x04000000
	CREATE_NEW_CONSOLE               = 0x00000010
	CREATE_NEW_PROCESS_GROUP         = 0x00000200
	CREATE_NO_WINDOW                 = 0x08000000
	CREATE_PROTECTED_PROCESS         = 0x00040000
	CREATE_PRESERVE_CODE_AUTHZ_LEVEL = 0x02000000
	CREATE_SEPARATE_WOW_VDM          = 0x00000800
	CREATE_SHARED_WOW_VDM            = 0x00001000
	CREATE_SUSPENDED                 = 0x00000004
	CREATE_UNICODE_ENVIRONMENT       = 0x00000400
	DEBUG_ONLY_THIS_PROCESS          = 0x00000002
	DEBUG_PROCESS                    = 0x00000001
	DETACHED_PROCESS                 = 0x00000008
	EXTENDED_STARTUPINFO_PRESENT     = 0x00080000
	INHERIT_PARENT_AFFINITY          = 0x00010000
)

func (l *Launcher) startServerProcess(name string, args ...string) error {
	bin, err := exec.LookPath(filepath.Join(l.pathmgmt.GetEQEmuServerPath(), "bin", name))
	if err != nil {
		return err
	}

	err = filepathcheck.IsValid(strings.Join(args, " "))
	if err != nil {
		return err
	}

	cmd := exec.Command(bin, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

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
	tmpBin := filepath.Join(os.TempDir(), filepath.Base(bin))
	if err := copyFile(bin, tmpBin); err != nil {
		return err
	}

	// chmod
	if err := os.Chmod(tmpBin, 0755); err != nil {
		return err
	}

	cmd := exec.Command(tmpBin, "eqemu-server:launcher", "start")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: DETACHED_PROCESS | CREATE_NEW_PROCESS_GROUP,
		HideWindow:    true,
	}
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
