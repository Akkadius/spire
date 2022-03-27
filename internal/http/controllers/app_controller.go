package controllers

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type AppController struct {
	cache  *gocache.Cache
	logger *logrus.Logger
}

func NewAppController(cache *gocache.Cache, logger *logrus.Logger) *AppController {
	return &AppController{
		cache:  cache,
		logger: logger,
	}
}

func (d *AppController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "app/changelog", d.changelog, nil),
		routes.RegisterRoute(http.MethodGet, "app/env", d.env, nil),
	}
}

func (d *AppController) changelog(c echo.Context) error {
	changelog, _ := d.cache.Get("changelog")
	return c.JSON(200, echo.Map{"data": changelog})
}

type Features struct {
	GithubAuthEnabled bool `json:"github_auth_enabled"`
}

type EnvResponse struct {
	Env      string   `json:"env"`
	Version  string   `json:"version"`
	Features Features `json:"features"`
}

type PackageJson struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
}

func (d *AppController) env(c echo.Context) error {
	data, _ := d.cache.Get("packageJson")
	pJson, ok := data.([]byte)
	if ok {
		var pkg PackageJson
		err := json.Unmarshal(pJson, &pkg)
		if err != nil {
			return err
		}

		response := EnvResponse{
			Env:     env.Get("APP_ENV", "local"),
			Version: pkg.Version,
			Features: Features{
				GithubAuthEnabled: len(os.Getenv("GITHUB_CLIENT_ID")) > 0,
			},
		}

		return c.JSON(200, echo.Map{"data": response})
	}

	return c.JSON(500, echo.Map{"error": "Unknown error"})
}
