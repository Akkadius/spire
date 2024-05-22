package assets

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/env"
	appmiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SpireAssets struct {
	pathmanager *pathmgmt.PathManagement
}

func NewSpireAssets(
	pathmanager *pathmgmt.PathManagement,
) *SpireAssets {
	return &SpireAssets{
		pathmanager: pathmanager,
	}
}

const (
	assetRepo = "Akkadius/eq-asset-preview"
)

func (a SpireAssets) ServeStatic() echo.MiddlewareFunc {
	// get cachedir
	cachedir := filepath.Join(a.getCacheDir(), "spire", "assets")

	// check for assets
	if len(os.Getenv("SKIP_ASSET_CHECK")) == 0 {
		a.CheckForAssets()
	}

	// in development, perform a symlink between the downloaded assets and the frontend public directory
	// the reason for this is that in development we run the Vue development web server
	// and assets resolve relative to the webserver running assets there - so they need
	// to be available via the FE web dev server
	// this is a convenience
	if len(cachedir) > 0 {
		if env.IsAppEnvDev() {
			symlinkTarget := filepath.Join("./frontend/public/eq-asset-preview-master")

			// always remove the link if it is there
			_ = os.Remove(symlinkTarget)

			// relink
			err := os.Symlink(cachedir, symlinkTarget)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// serve
	return appmiddleware.StaticAsset(appmiddleware.StaticConfig{
		Root:        "/",
		StripPrefix: string(filepath.Separator) + "eq-asset-preview-master",
		Filesystem:  http.Dir(cachedir),
	})
}

func (a SpireAssets) downloadAssets(cachedir string) {
	fmt.Printf("Downloading [eq-asset-preview] latest release\n")

	// zip file path
	dumpZip := filepath.Join(os.TempDir(), "/build.zip")

	// download the zip file
	err := download.WithProgress(
		dumpZip,
		fmt.Sprintf("https://github.com/%v/releases/latest/download/build.zip", assetRepo),
	)
	if err != nil {
		log.Fatal(err)
	}

	// unzip the file
	err = unzip.New(dumpZip, cachedir).Extract()
	if err != nil {
		log.Fatal(err)
	}

	// remove the zip file
	_ = os.Remove(dumpZip)
}

func (a SpireAssets) getCacheDir() string {
	userDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Printf("could not get user cache dir: %v\n", err)
	}
	if len(userDir) > 0 {
		return userDir
	}

	return os.TempDir()
}

func (a SpireAssets) CheckForAssets() {
	cachedir := filepath.Join(a.getCacheDir(), "spire", "assets")

	// check if cachedir exists
	if _, err := os.Stat(cachedir); os.IsNotExist(err) {
		a.downloadAssets(cachedir)
	}

	// check if we're running a command
	// github rate limits us if we check too often
	// 60 requests per hour
	// write a cache file to check last time we checked
	// if we checked within the last hour, don't check again
	// if we checked more than an hour ago, check again
	tmpFile := filepath.Join(os.TempDir(), "spire_asset_last_check")
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		// file doesn't exist, create it
		_, err := os.Create(tmpFile)
		if err != nil {
			log.Fatal(err)
		}

		_ = os.Chtimes(tmpFile, time.Now().Add(-time.Hour*1), time.Now().Add(-time.Hour*1))
	}

	// check if file modified over an hour ago
	fileInfo, err := os.Stat(tmpFile)
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now()

	assetCheckPassedTime := currentTime.Sub(fileInfo.ModTime()).Hours() < 1
	if assetCheckPassedTime {
		return
	}

	// GitHub release struct
	type GitHubRelease struct {
		TagName string `json:"tag_name"`
	}

	// get latest release version
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%v/releases/latest", assetRepo))
	if err != nil {
		fmt.Printf("could not get latest release version for [%v] %v", assetRepo, err)
	}

	if resp == nil {
		fmt.Printf("could not get latest release version for [%v] nil response", assetRepo)
		return
	}

	defer resp.Body.Close()

	// check if response is 200
	if err == nil {
		// read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("could not read response body: %v", err)
		}

		// bind body to struct
		var release GitHubRelease
		err = json.Unmarshal(body, &release)
		if err != nil {
			log.Fatalf("could not unmarshal response body: %v", err)
		}

		type PackageJson struct {
			Version string `json:"version"`
		}

		// get current version from package.json
		file := filepath.Join(cachedir, "package.json")

		// read file package.json contents into PackageJson struct
		packageJson, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("could not read PackageJson file: %v", err)
		}

		// bind package.json to struct
		var packageJsonStruct PackageJson
		err = json.Unmarshal(packageJson, &packageJsonStruct)
		if err != nil {
			log.Fatalf("could not unmarshal package.json: %v", err)
		}

		// check if current version is the same as the latest release version
		remoteRelease := strings.ReplaceAll(release.TagName, "v", "")
		if len(remoteRelease) > 0 && packageJsonStruct.Version != remoteRelease {
			fmt.Printf(
				"New version available, downloading [eq-asset-preview] release [%v]\n",
				release.TagName,
			)
			a.downloadAssets(cachedir)
		}

		// update the file modified time
		_ = os.Chtimes(tmpFile, time.Now(), time.Now())
	}
}
