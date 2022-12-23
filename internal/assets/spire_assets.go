package assets

import (
	"fmt"
	"github.com/Akkadius/spire/internal/github"
	appmiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

type SpireAssets struct {
	logger     *logrus.Logger
	cache      *gocache.Cache
	downloader *github.GithubSourceDownloader
}

const (
	organization = "Akkadius"
	repository   = "eq-asset-preview"
	version      = 1
)

func NewSpireAssets(
	logger *logrus.Logger,
	cache *gocache.Cache,
	downloader *github.GithubSourceDownloader,
) *SpireAssets {
	return &SpireAssets{
		logger:     logger,
		cache:      cache,
		downloader: downloader,
	}
}

func (a SpireAssets) ServeStatic() echo.MiddlewareFunc {
	// cleanup old versions
	for i := 1; i < version; i++ {
		oldPath := filepath.Join(a.downloader.GetSourceRoot(), fmt.Sprintf("%v-v%v", repository, i))
		fmt.Printf("Deleting old assets [%v]\n", oldPath)
		err := os.RemoveAll(oldPath)
		if err != nil {
			a.logger.Fatal(err)
		}
	}

	// source
	branch := fmt.Sprintf("v%v", version)
	a.downloader.SourceToUserCacheDir(true)
	r := a.downloader.Source(organization, repository, branch, false)

	// serve
	return appmiddleware.StaticAsset(appmiddleware.StaticConfig{
		Root:        "/",
		StripPrefix: string(filepath.Separator) + "eq-asset-preview-master",
		Filesystem:  http.Dir(r.ZippedPath),
	})
}
