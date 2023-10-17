package assets

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/env"
	appmiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SpireAssets struct {
	logger      *logrus.Logger
	pathmanager *pathmgmt.PathManagement
}

func NewSpireAssets(
	logger *logrus.Logger,
	pathmanager *pathmgmt.PathManagement,
) *SpireAssets {
	return &SpireAssets{
		logger:      logger,
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
		a.checkForAssets()
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
				a.logger.Fatal(err)
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
	a.logger.Infof("Downloading [eq-asset-preview] latest release\n")

	attempt := 1
	for {
		err := a.doDownloadAssets(cachedir)
		if err != nil {
			if attempt < 3 {
				attempt++
				continue
			}
			a.logger.Fatal(err)
		}
		break
	}

	dumpZip := filepath.Join(os.TempDir(), "/build.zip")

	// remove the zip file
	_ = os.Remove(dumpZip)
}

func (a SpireAssets) doDownloadAssets(cachedir string) error {
	// zip file path
	dumpZip := filepath.Join(os.TempDir(), "/build.zip")

	// download the zip file
	err := download.WithProgress(
		dumpZip,
		fmt.Sprintf("https://github.com/%v/releases/latest/download/build.zip", assetRepo),
	)
	if err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	// unzip the file
	a.logger.Infof("Downloaded zip to [%v]\n", dumpZip)
	err = unzip.New(dumpZip, cachedir).Extract()
	if err != nil {
		return errors.New(fmt.Sprintf("could not extract zip: %v", err))
	}

	return nil
}

func (a SpireAssets) getCacheDir() string {
	userDir, err := os.UserCacheDir()
	if err != nil {
		a.logger.Error(err)
	}
	if len(userDir) > 0 {
		return userDir
	}

	return os.TempDir()
}

func (a SpireAssets) checkForAssets() {
	cachedir := filepath.Join(a.getCacheDir(), "spire", "assets")

	// check if cachedir exists
	if _, err := os.Stat(cachedir); os.IsNotExist(err) {
		a.downloadAssets(cachedir)
	}

	// GitHub release struct
	type GitHubRelease struct {
		TagName string `json:"tag_name"`
	}

	// get latest release version
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%v/releases/latest", assetRepo))
	if err != nil {
		a.logger.Info("could not get latest release version for [%v] %v", assetRepo, err)
	}

	defer resp.Body.Close()

	// check if response is 200
	if err == nil {
		// read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			a.logger.Fatalf("could not read response body: %v", err)
		}

		// bind body to struct
		var release GitHubRelease
		err = json.Unmarshal(body, &release)
		if err != nil {
			a.logger.Fatalf("could not unmarshal response body: %v", err)
		}

		type PackageJson struct {
			Version string `json:"version"`
		}

		// get current version from package.json
		file := filepath.Join(cachedir, "package.json")

		// read file package.json contents into PackageJson struct
		packageJson, err := os.ReadFile(file)
		if err != nil {
			a.logger.Fatalf("could not read PackageJson file: %v", err)
		}

		// bind package.json to struct
		var packageJsonStruct PackageJson
		err = json.Unmarshal(packageJson, &packageJsonStruct)
		if err != nil {
			a.logger.Fatalf("could not unmarshal package.json: %v", err)
		}

		// check if current version is the same as the latest release version
		remoteRelease := strings.ReplaceAll(release.TagName, "v", "")
		if len(remoteRelease) > 0 && packageJsonStruct.Version != remoteRelease {
			a.logger.Infof(
				"New version available, downloading [eq-asset-preview] release [%v]\n",
				release.TagName,
			)
			a.downloadAssets(cachedir)
		}
	}
}
