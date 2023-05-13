package occulus

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/google/go-github/v41/github"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type ProcessManagement struct {
	port     int
	logger   *logrus.Logger
	pathmgmt *pathmgmt.PathManagement
}

func (m *ProcessManagement) Port() int {
	return m.port
}

func (m *ProcessManagement) SetPort(port int) {
	m.port = port
}

func NewProcessManagement(pathmgmt *pathmgmt.PathManagement, logger *logrus.Logger) *ProcessManagement {
	i := &ProcessManagement{
		pathmgmt: pathmgmt,
		logger:   logger,
	}

	return i
}

func checkIfPortAvailable(port int) (status bool, err error) {
	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false, err
	}

	// close the server
	server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return true, nil

}

func (m *ProcessManagement) Run() error {
	// kill existing
	err := m.KillExistingRunningProcesses()
	if err != nil {
		return err
	}

	// check if we have a binary
	// if not, download it
	downloadPath, err := m.FetchOcculusAndGetBinaryPath()
	if err != nil {
		return err
	}

	// run binary
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)

			if len(downloadPath) == 0 {
				m.logger.Info("[Occulus.ProcessManagement] No binary found")
				continue
			}

			// free port
			port := m.FindFreePort()
			if port > 0 {
				m.logger.Infof("[Occulus.ProcessManagement] Found free port @ [%v]\n", port)
				m.SetPort(port)
			}

			if port == 0 {
				m.logger.Fatalf("[Occulus.ProcessManagement] Failed to find free port\n")
			}

			runPath := downloadPath
			developmentRunCmd := env.Get("OCCULUS_DEV_CMD", "")
			if len(developmentRunCmd) > 0 {
				runPath = developmentRunCmd
			}

			m.logger.Infof("[Occulus.ProcessManagement] Running binary/command [%v] via port [%v]\n", runPath, port)
			cmd := exec.Command(runPath, "web", fmt.Sprintf("%v", port))

			stdout, err := cmd.StdoutPipe()
			if err != nil {
				m.logger.Fatalf("[Occulus.ProcessManagement] could not get stdout pipe: %v", err)
			}

			cmd.Stderr = cmd.Stdout

			err = cmd.Start()
			if err != nil {
				m.logger.Error(err)
			}

			merged := io.MultiReader(stdout)
			scanner := bufio.NewScanner(merged)
			for scanner.Scan() {
				m.logger.Printf("[Occulus] %v\n", scanner.Text())
			}

			err = cmd.Wait()
			if err != nil {
				m.logger.Error(err)
			}
		}
	}()

	return nil
}

func (m *ProcessManagement) FindFreePort() int {
	// if we want to override the port over the statically defined port
	envPort := env.GetInt("OCCULUS_PORT", "0")
	if envPort > 0 {
		return envPort
	}

	// pull port dynamically from range
	port := 0
	for i := 49152; i <= 65535; i++ {
		found, err := checkIfPortAvailable(i)
		if found && err == nil {
			port = i
			break
		}
	}

	return port
}

// GetCurrentOcculusBinaryPath will return the path to the current occulus binary
func (m *ProcessManagement) GetCurrentOcculusBinaryPath() (string, error) {
	serverBinDirectory := filepath.Join(m.pathmgmt.GetEQEmuServerPath(), "bin")
	files, err := os.ReadDir(serverBinDirectory)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "occulus") {
			m.logger.Infof("[Occulus.ProcessManagement] Found binary [%v]\n", file.Name())
			binPath := filepath.Join(serverBinDirectory, file.Name())
			return binPath, nil
		}
	}

	return "", nil
}

// CleanupOldVersions will remove any old versions of occulus
func (m *ProcessManagement) CleanupOldVersions(version string) error {
	versionTag := strings.ReplaceAll(version, ".", "-")
	currentBinaryName := fmt.Sprintf("%v-%v", "occulus", versionTag)
	currentBinaryName = strings.ReplaceAll(currentBinaryName, ".", "-")
	if runtime.GOOS == "windows" {
		currentBinaryName += ".exe"
	}
	serverBinDirectory := filepath.Join(m.pathmgmt.GetEQEmuServerPath(), "bin")

	files, err := os.ReadDir(serverBinDirectory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "occulus") && file.Name() != currentBinaryName {
			m.logger.Infof("[Occulus.ProcessManagement] Removing old binary [%v] current [%v]\n", file.Name(), currentBinaryName)
			binPath := filepath.Join(serverBinDirectory, file.Name())
			if err := os.Remove(binPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// KillExistingRunningProcesses kills any existing running processes
func (m *ProcessManagement) KillExistingRunningProcesses() error {
	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			return err
		}
		if strings.Contains(cmdline, "occulus") && strings.Contains(cmdline, "web") {
			processName, _ := p.Name()
			m.logger.Infof("[Occulus.ProcessManagement] Found existing running process, killing [%v] PID (%v)\n", processName, p.Pid)

			err := p.Kill()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// FetchOcculusAndGetBinaryPath fetches the latest occulus binary and returns the path
func (m *ProcessManagement) FetchOcculusAndGetBinaryPath() (string, error) {
	client := github.NewClient(nil)
	if len(os.Getenv("GITHUB_TOKEN")) > 0 {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		tc := &http.Client{
			Timeout: 5 * time.Second,
			Transport: &oauth2.Transport{
				Source: ts,
			},
		}
		client = github.NewClient(tc)
	}

	release, _, err := client.Repositories.GetLatestRelease(
		context.Background(),
		"Akkadius",
		"Occulus",
	)

	var downloadPath string

	if err != nil {
		m.logger.Info(err)
		downloadPath, err = m.GetCurrentOcculusBinaryPath()
		m.logger.Infof("[Occulus.ProcessManagement] Using existing download path @ [%v]\n", downloadPath)
		if err != nil {
			return "", err
		}
	}

	isWindows := runtime.GOOS == "windows"
	if err == nil && len(downloadPath) == 0 {
		// build binary target name from asset name
		// eg. occulus-v2-1-0
		tagName := *release.TagName
		binaryName := fmt.Sprintf("%v-%v", "occulus", tagName)
		binaryName = strings.ReplaceAll(binaryName, ".", "-")

		downloadPath = filepath.Join(m.pathmgmt.GetEQEmuServerPath(), "bin", binaryName)

		// cleanup
		err = m.CleanupOldVersions(tagName)
		if err != nil {
			return "", err
		}

		if isWindows {
			downloadPath += ".exe"
		}

		// check if binary exists before we try to download it
		if _, err := os.Stat(downloadPath); errors.Is(err, os.ErrNotExist) {
			// loop through latest release assets
			for _, asset := range release.Assets {
				releaseAssetName := *asset.Name
				releaseDownloadUrl := *asset.BrowserDownloadURL

				// Occulus assets use `-win` suffix
				assetMatch := runtime.GOOS
				if isWindows {
					assetMatch = "win"
				}

				// find asset / release matching the operating system
				if strings.Contains(releaseAssetName, assetMatch) {
					m.logger.Infof("[Occulus.ProcessManagement] Downloading new binary @ [%v]\n", downloadPath)
					err := download.WithProgress(downloadPath, releaseDownloadUrl)
					if err != nil {
						return "", err
					}
				}
			}
		}
	}

	// windows is a strange beast
	if isWindows {
		downloadPath = strings.ReplaceAll(downloadPath, ".exe", "")
	}

	return downloadPath, nil
}
