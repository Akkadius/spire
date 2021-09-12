package github

import (
	"fmt"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type GithubSourceDownloader struct {
	logger *logrus.Logger
}

func NewGithubSourceDownloader(logger *logrus.Logger) *GithubSourceDownloader {
	return &GithubSourceDownloader{logger: logger}
}

// sources quest files and returns filename:contents
func (g *GithubSourceDownloader) Source(org string, repo string, branch string, forceRefresh bool) map[string]string {

	// repo params
	repoDir := fmt.Sprintf("%v/%v-%v/", os.TempDir(), repo, branch)
	if runtime.GOOS == "windows" {
		repoDir = fmt.Sprintf("%v\\%v-%v\\", os.TempDir(), repo, branch)
	}

	repoZipUrl := fmt.Sprintf("https://github.com/%v/%v/archive/%v.zip", org, repo, branch)

	//unzipLoc := fmt.Sprintf("%v/quests/", os.TempDir())
	zipFileLocalLoc := fmt.Sprintf("%v/%v.zip", os.TempDir(), repo)

	if forceRefresh {
		err := os.RemoveAll(repoDir)
		if err != nil {
			g.logger.Error(err)
		}
	}

	// if not exist, extract
	if _, err := os.Stat(repoDir); os.IsNotExist(err) || forceRefresh {
		err := g.downloadFile(zipFileLocalLoc, repoZipUrl)
		if err != nil {
			g.logger.Error(err)
		}

		uz := unzip.New(zipFileLocalLoc, repoDir)
		err = uz.Extract()
		if err != nil {
			g.logger.Error(err)
		}
	}

	var unzippedFiles = map[string]string{}

	// walk files
	err := filepath.Walk(
		repoDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// stat file
			fi, err := os.Stat(path)
			if err != nil {
				g.logger.Fatal(err)
			}

			// if regular file - not dir
			if fi.Mode().IsRegular() {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					g.logger.Fatal(err)
				}

				fileName := strings.ReplaceAll(path, fmt.Sprintf("%v%v-%v/", repoDir, repo, branch), "")
				if runtime.GOOS == "windows" {
					fileName = strings.ReplaceAll(path, fmt.Sprintf("%v%v-%v\\", repoDir, repo, branch), "")
					fileName = strings.ReplaceAll(fileName, "\\", "/")
				}

				unzippedFiles[fileName] = string(data)
			}

			return nil
		},
	)
	if err != nil {
		g.logger.Error(err)
	}

	return unzippedFiles
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (g *GithubSourceDownloader) downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
