package updater

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/google/go-github/v41/github"
	"github.com/mattn/go-isatty"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

// Updater is a service that checks for updates to the app
type Updater struct {
	// this is the package.json embedded in the binary which contains the app version
	packageJson  []byte
	logger       *logger.AppLogger
	serverconfig *eqemuserverconfig.Config
	unzipper     *unzip.Unzipper
}

// NewUpdater creates a new updater service
func NewUpdater(packageJson []byte) *Updater {
	appLogger := logger.ProvideAppLogger()
	pathmgr := pathmgmt.NewPathManagement(appLogger)
	return &Updater{
		packageJson: packageJson,
		logger:      appLogger,
		serverconfig: eqemuserverconfig.NewConfig(
			appLogger,
			pathmgr,
		),
		unzipper: unzip.NewUnzipper(appLogger),
	}
}

// EnvResponse is the response from the env endpoint
type EnvResponse struct {
	Env     string `json:"env"`
	Version string `json:"version"`
}

// PackageJson is the package.json file
type PackageJson struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
}

// getAppVersion gets the app version from the package.json embedded in the binary
func (s *Updater) getAppVersion() (error, EnvResponse) {
	var pkg PackageJson
	err := json.Unmarshal(s.packageJson, &pkg)
	if err != nil {
		return err, EnvResponse{}
	}

	return nil, EnvResponse{
		Env:     env.Get("APP_ENV", "local"),
		Version: pkg.Version,
	}
}

// CheckForUpdates checks for updates to the app
func (s *Updater) CheckForUpdates(interactive bool) bool {
	config, err := s.serverconfig.Get()
	if err == nil && config.Spire.DisableAutoUpdates && interactive {
		s.logger.Info().
			Any("spire.disable_auto_updates", config.Spire.DisableAutoUpdates).
			Msg("Auto updates are disabled via config")
		return false
	}

	// get executable name and path
	executableName := filepath.Base(os.Args[0])
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	executablePath := filepath.Dir(ex)

	s.logger.Debug().
		Any("executableName", executableName).
		Any("executablePath", executablePath).
		Msg("Checking for updates")

	// check if a .old version exists, delete it if does
	oldExecutable := filepath.Join(executablePath, fmt.Sprintf("%s.old", executableName))
	if _, err := os.Stat(oldExecutable); err == nil {
		s.logger.Info().Any("oldExecutable", oldExecutable).Msg("Removing old executable")
		e := os.Remove(oldExecutable)
		if e != nil {
			s.logger.Fatal().Err(e).Msg("Failed to remove old executable")
		}
	}

	// if being ran from go run main.go
	if executableName == "main.exe" || executableName == "main" {
		s.logger.Info().Msg("Running as go run main.go, ignoring...")
		return false
	}

	// internet connection check
	if !isconnected() {
		s.logger.Info().Msgf("Not connected to the internet")
		return false
	}

	s.logger.Info().Msg("Checking for updates")
	s.logger.Info().Any("executableName", executableName).Msg("Running as binary")
	s.logger.Debug().Any("executableName", executableName).Msg("Checking for updates")

	// get releases
	client := github.NewClient(&http.Client{Timeout: 5 * time.Second})
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "Akkadius", "spire")
	if err != nil {
		s.logger.Info().Err(err).Msg("Failed to get latest release")
		return false
	}

	// get app version
	err, e := s.getAppVersion()
	if err != nil {
		s.logger.Info().Err(err).Msg("Failed to get app version")
	}

	localVersion := fmt.Sprintf("v%v", e.Version)
	releaseVersion := *release.TagName

	// already up to date
	if releaseVersion == localVersion {
		s.logger.Info().Any("version", localVersion).Msgf("Spire is already up to date")
		return false
	}

	// remove asset check file if we have an update
	tmpFile := filepath.Join(os.TempDir(), "spire_asset_last_check")
	if _, err := os.Stat(tmpFile); err == nil {
		e := os.Remove(tmpFile)
		if e != nil {
			s.logger.Fatal().Err(e).Msg("Failed to remove asset check file")
		}
	}

	s.logger.Info().
		Any("local", localVersion).
		Any("latest", releaseVersion).
		Msg("Comparing local version to latest version")

	for _, asset := range release.Assets {
		assetName := *asset.Name
		downloadUrl := *asset.BrowserDownloadURL
		targetFileNameZipped := fmt.Sprintf("spire-%s-%s.zip", runtime.GOOS, runtime.GOARCH)
		if runtime.GOOS == "windows" {
			targetFileNameZipped = fmt.Sprintf("spire-%s-%s.exe.zip", runtime.GOOS, runtime.GOARCH)
		}
		targetFileName := fmt.Sprintf("spire-%s-%s", runtime.GOOS, runtime.GOARCH)

		s.logger.Debug().
			Any("assetName", assetName).
			Any("targetFileNameZipped", targetFileNameZipped).
			Msg("Checking asset")

		if assetName == targetFileNameZipped {
			s.logger.Info().Any("assetName", assetName).Msg("Found matching release")

			// download
			file := path.Base(downloadUrl)
			downloadPath := filepath.Join(os.TempDir(), file)
			err := download.WithProgress(downloadPath, downloadUrl)
			if err != nil {
				s.logger.Fatal().Err(err).Msg("Failed to download asset")
			}

			// unzip
			tempFileZipped := filepath.Join(os.TempDir(), targetFileNameZipped)
			err = s.unzipper.Extract(tempFileZipped, os.TempDir())
			if err != nil {
				s.logger.Fatal().Err(err).Msg("Failed to extract zip")
			}

			// rename running process to .old
			err = os.Rename(
				filepath.Join(executablePath, executableName),
				filepath.Join(executablePath, fmt.Sprintf("%s.old", executableName)),
			)
			if err != nil {
				s.logger.Fatal().Err(err).Msg("Failed to rename executable")
			}

			// relink
			lookTempFile := filepath.Join(os.TempDir(), targetFileName)
			tempFile, err := exec.LookPath(lookTempFile)
			if err != nil {
				s.logger.Fatal().Err(err).Msg("Failed to find executable")
			}

			newExecutable := fmt.Sprintf("%s%s%s", executablePath, string(filepath.Separator), executableName)
			err = moveFile(tempFile, newExecutable)
			if err != nil {
				s.logger.Fatal().Err(err).Msg("Failed to move executable")
			}

			err = os.Chmod(newExecutable, 0755)
			if err != nil {
				s.logger.Fatal().Err(err).Msg("Failed to chmod executable")
			}

			// if terminal, wait for user input
			if isatty.IsTerminal(os.Stdout.Fd()) && interactive {
				s.logger.Info().
					Any("version", releaseVersion).
					Msgf("Spire successfully updated, you must relaunch Spire manually")
				s.logger.Info().Msg("Press any key to continue...")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				return true
			} else {
				s.logger.Info().
					Any("version", releaseVersion).
					Msgf("Spire successfully updated, you must relaunch Spire manually")
				return true
			}
		}
	}
	return false
}
