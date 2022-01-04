package updater

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/google/go-github/v41/github"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type UpdaterService struct {
	// this is the package.json embedded in the binary which contains the app version
	packageJson []byte
}

func NewUpdaterService(packageJson []byte) *UpdaterService {
	return &UpdaterService{
		packageJson: packageJson,
	}
}

type EnvResponse struct {
	Env     string `json:"env"`
	Version string `json:"version"`
}

type PackageJson struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
}

func (s UpdaterService) getAppVersion() (error, EnvResponse) {
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

func (s UpdaterService) CheckForUpdates() {

	// get executable name and path
	executableName := filepath.Base(os.Args[0])
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	executablePath := filepath.Dir(ex)

	// check if a .old version exists, delete it if does
	oldExecutable :=  fmt.Sprintf("%s\\%s.old", executablePath, executableName)
	if _, err := os.Stat(oldExecutable); err == nil {
		e := os.Remove(oldExecutable)
		if e != nil {
			log.Fatal(e)
		}
	}

	// if being ran from go run main.go
	if executableName == "main.exe" {
		fmt.Println("[Update] Running as go run main.go, ignoring...")
		return
	}

	// internet connection check
	if !isconnected() {
		fmt.Printf("[Update] Not connected to the internet\n")
		return
	}

	fmt.Printf("[Update] Checking for updates...\n")
	fmt.Printf("[Update] Running as binary [%v]\n", executableName)
	debug(fmt.Sprintf("[Update] Running as executablePath [%v]\n", executablePath))

	// get releases
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "Akkadius", "spire")
	if err != nil {
		//log.Println(err)
		return
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
		fmt.Printf("[Update] Spire is already up to date @ [%v]\n", localVersion)
		return
	}

	fmt.Printf("Local version [%s] latest [%v]\n", localVersion, releaseVersion)

	for _, asset := range release.Assets {
		assetName := *asset.Name
		downloadUrl := *asset.BrowserDownloadURL
		targetFileNameZipped := fmt.Sprintf("spire-%s-%s.zip", runtime.GOOS, runtime.GOARCH)
		if runtime.GOOS == "windows" {
			targetFileNameZipped = fmt.Sprintf("spire-%s-%s.exe.zip", runtime.GOOS, runtime.GOARCH)
		}
		targetFileName := fmt.Sprintf("spire-%s-%s", runtime.GOOS, runtime.GOARCH)

		debug(fmt.Sprintf("[Update] Looping assets assetName [%v] targetFileNameZipped [%v]\n", assetName, targetFileNameZipped))

		if assetName == targetFileNameZipped {
			fmt.Printf("Found matching release [%s]\n", assetName)

			// download
			downloadFile(downloadUrl, os.TempDir())

			// linux
			if runtime.GOOS == "linux" {
				// unzip
				tempFileZipped := fmt.Sprintf("%s/%s", os.TempDir(), targetFileNameZipped)
				uz := unzip.New(tempFileZipped, os.TempDir())
				err = uz.Extract()
				if err != nil {
					log.Println(err)
				}

				// relink
				tempFile := fmt.Sprintf("%s/%s", os.TempDir(), targetFileName)
				err := moveFile(tempFile, "spire")
				if err != nil {
					log.Println(err)
				}

				err = os.Chmod("spire", 0755)
				if err != nil {
					log.Println(err)
				}
			}

			// windows
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

			fmt.Printf("[update] Spire updated to version [%s] restart Spire\n", releaseVersion)
			fmt.Print("Press 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}

func debug(msg string) {
	if len(os.Getenv("DEBUG")) > 0 {
		fmt.Printf("[Debug] " + msg)
	}
}
