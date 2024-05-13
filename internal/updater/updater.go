package updater

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/google/go-github/v41/github"
	"github.com/mattn/go-isatty"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

// Updater is a service that checks for updates to the app
type Updater struct {
	// this is the package.json embedded in the binary which contains the app version
	packageJson []byte
	logger      *logger.AppLogger
}

// NewUpdater creates a new updater service
func NewUpdater(packageJson []byte) *Updater {
	return &Updater{
		packageJson: packageJson,
		logger:      logger.ProvideAppLogger(),
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

	// get executable name and path
	executableName := filepath.Base(os.Args[0])
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	executablePath := filepath.Dir(ex)

	// check if a .old version exists, delete it if does
	oldExecutable := fmt.Sprintf("%s\\%s.old", executablePath, executableName)
	if runtime.GOOS == "linux" {
		oldExecutable = fmt.Sprintf("%s/%s.old", executablePath, executableName)
	}
	if _, err := os.Stat(oldExecutable); err == nil {
		e := os.Remove(oldExecutable)
		if e != nil {
			log.Fatal(e)
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
	s.logger.Debug().Msgf("Running as executablePath [%v]", executablePath)

	// get releases
	client := github.NewClient(&http.Client{Timeout: 5 * time.Second})
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "Akkadius", "spire")
	if err != nil {
		log.Println(err)
		return false
	}

	// get app version
	err, e := s.getAppVersion()
	if err != nil {
		log.Println(err)
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
			log.Fatal(e)
		}
	}

	fmt.Printf("Local version [%s] latest [%v]", localVersion, releaseVersion)

	for _, asset := range release.Assets {
		assetName := *asset.Name
		downloadUrl := *asset.BrowserDownloadURL
		targetFileNameZipped := fmt.Sprintf("spire-%s-%s.zip", runtime.GOOS, runtime.GOARCH)
		if runtime.GOOS == "windows" {
			targetFileNameZipped = fmt.Sprintf("spire-%s-%s.exe.zip", runtime.GOOS, runtime.GOARCH)
		}
		targetFileName := fmt.Sprintf("spire-%s-%s", runtime.GOOS, runtime.GOARCH)

		s.logger.Debug().Msgf("Looping assets assetName [%v] targetFileNameZipped [%v]", assetName, targetFileNameZipped)

		if assetName == targetFileNameZipped {
			fmt.Printf("Found matching release [%s]", assetName)

			// download
			file := path.Base(downloadUrl)
			downloadPath := filepath.Join(os.TempDir(), file)
			err := download.WithProgress(downloadPath, downloadUrl)
			if err != nil {
				log.Println(err)
			}

			// linux
			// todo: Move these checks to use platform agnostic filepath.Join calls
			// these will be risky to refactor so will need testing
			if runtime.GOOS == "linux" {

				// unzip
				tempFileZipped := fmt.Sprintf("%s/%s", os.TempDir(), targetFileNameZipped)
				uz := unzip.New(tempFileZipped, os.TempDir())
				err = uz.Extract()
				if err != nil {
					log.Println(err)
				}

				// rename running process to .old
				err := os.Rename(
					fmt.Sprintf("%s/%s", executablePath, executableName),
					fmt.Sprintf("%s/%s.old", executablePath, executableName),
				)
				if err != nil {
					log.Fatal(err)
				}

				// relink
				tempFile := fmt.Sprintf("%s/%s", os.TempDir(), targetFileName)
				newExecutable := fmt.Sprintf("%s/%s", executablePath, executableName)
				err = moveFile(tempFile, newExecutable)
				if err != nil {
					log.Println(err)
				}

				err = os.Chmod(newExecutable, 0755)
				if err != nil {
					log.Println(err)
				}
			}

			// windows
			// todo: Move these checks to use platform agnostic filepath.Join calls
			// these will be risky to refactor so will need testing
			if runtime.GOOS == "windows" {
				// unzip
				tempFileZipped := fmt.Sprintf("%s\\%s", os.TempDir(), targetFileNameZipped)
				uz := unzip.New(tempFileZipped, os.TempDir())
				err = uz.Extract()
				if err != nil {
					log.Println(err)
				}

				// rename running process to .old
				err := os.Rename(
					fmt.Sprintf("%s\\%s", executablePath, executableName),
					fmt.Sprintf("%s\\%s.old", executablePath, executableName),
				)
				if err != nil {
					log.Fatal(err)
				}

				// relink
				tempFile := fmt.Sprintf("%s\\%s.exe", os.TempDir(), targetFileName)
				newExecutable := fmt.Sprintf("%s\\%s", executablePath, executableName)
				err = moveFile(tempFile, newExecutable)
				if err != nil {
					log.Println(err)
				}

				err = os.Chmod(newExecutable, 0755)
				if err != nil {
					log.Println(err)
				}
			}

			// if terminal, wait for user input
			if isatty.IsTerminal(os.Stdout.Fd()) && interactive {
				fmt.Println("")
				s.logger.Info().Msgf("Spire updated to version [%s] you must relaunch Spire manually", releaseVersion)
				fmt.Println("")
				fmt.Print("Press [Enter] to exit spire...")
				fmt.Println("")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				return true
			} else {
				s.logger.Info().Msgf("Spire updated to version [%s] you must relaunch Spire manually", releaseVersion)
				return true
			}
		}
	}
	return false
}
