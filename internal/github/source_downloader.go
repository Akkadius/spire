package github

import (
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/unzip"
	gocache "github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SourceDownloader struct {
	cache          *gocache.Cache
	sourceUserDir  bool // if set, sources files to user directory
	sourcedDirPath string
	readFiles      bool // if set, will open files and return contents
	unzipper       *unzip.Unzipper
}

func (g *SourceDownloader) SetReadFiles(readFiles bool) {
	g.readFiles = readFiles
}

func (g *SourceDownloader) SourceToUserCacheDir(sourceUserDir bool) {
	g.sourceUserDir = sourceUserDir
}

func NewGithubSourceDownloader(cache *gocache.Cache, unzipper *unzip.Unzipper) *SourceDownloader {
	return &SourceDownloader{
		cache:     cache,
		unzipper:  unzipper,
		readFiles: true,
	}
}

type SourceResult struct {
	Files      map[string]string
	ZippedPath string
}

// Source downloads files and returns filename:contents
func (g *SourceDownloader) Source(org string, repo string, branch string, forceRefresh bool) SourceResult {
	lockKey := fmt.Sprintf("%v-%v-%v", org, repo, branch)
	// if lock set, return
	_, found := g.cache.Get(lockKey)
	if found {
		return SourceResult{}
	}

	g.cache.Set(lockKey, 1, time.Minute*10)

	// repo params
	repoDir := g.GetSourcedDirPath(repo, branch)
	repoZipUrl := fmt.Sprintf("https://github.com/%v/%v/archive/%v.zip", org, repo, branch)
	zipFileLocalLoc := filepath.Join(g.GetSourceRoot(), fmt.Sprintf("%v.zip", repo))

	if forceRefresh {
		err := os.RemoveAll(repoDir)
		if err != nil {
			fmt.Println(err)
		}
	}

	// if not exist, extract
	if _, err := os.Stat(repoDir); os.IsNotExist(err) || forceRefresh {
		err := g.downloadFile(zipFileLocalLoc, repoZipUrl)
		if err != nil {
			fmt.Println(err)
		}

		err = g.unzipper.Extract(zipFileLocalLoc, repoDir)
		if err != nil {
			fmt.Println(err)
		}
	}

	zippedPath := ""
	files, _ := os.ReadDir(repoDir)
	if len(files) > 0 {
		zippedPath = filepath.Join(repoDir, files[0].Name())
	}

	var unzippedFiles = map[string]string{}

	// walk files
	if g.readFiles {
		err := filepath.Walk(
			repoDir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// stat file
				fi, err := os.Stat(path)
				if err != nil {
					fmt.Println(err)
				}

				// if regular file - not dir
				if fi.Mode().IsRegular() {
					data, err := ioutil.ReadFile(path)
					if err != nil {
						fmt.Println(err)
					}

					// construct relative path
					// remove base file path
					fileName := strings.ReplaceAll(path, repoDir, "")
					// strip top two directory levels
					dirs := strings.Split(fileName, string(filepath.Separator))
					dirs = dirs[2:]

					fileName = strings.Join(dirs, string(filepath.Separator))

					unzippedFiles[fileName] = string(data)
				}

				return nil
			},
		)
		if err != nil {
			fmt.Println(err)
		}
	}

	// file exists
	if _, err := os.Stat(zipFileLocalLoc); err == nil {
		err = os.Remove(zipFileLocalLoc)
		if err != nil {
			log.Fatal(err)
		}
	}

	// delete lock
	g.cache.Delete(lockKey)

	return SourceResult{
		Files:      unzippedFiles,
		ZippedPath: zippedPath + "/",
	}
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func (g *SourceDownloader) downloadFile(filepath string, url string) error {
	err := download.WithProgress(filepath, url)
	if err != nil {
		return err
	}

	return err
}

// GetSourceRoot returns user cache dir or temp dir
// eg /home/user/.cache/ or /tmp/
func (g *SourceDownloader) GetSourceRoot() string {
	userDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println(err)
	}
	if len(userDir) > 0 {
		return userDir
	}

	return os.TempDir()
}

// GetSourcedDirPath returns the sourced download path
// eg /home/user/.cache/<repo>-<branch> or /tmp/<repo>-<branch>
func (g *SourceDownloader) GetSourcedDirPath(repo string, branch string) string {
	return filepath.Join(g.GetSourceRoot(), fmt.Sprintf("%v-%v", repo, branch))
}
