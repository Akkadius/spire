package controllers

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type AppController struct {
	cache      *gocache.Cache
	logger     *logrus.Logger
	onboarding *spire.SpireInit
}

func NewAppController(cache *gocache.Cache, logger *logrus.Logger, onboarding *spire.SpireInit) *AppController {
	return &AppController{
		cache:      cache,
		logger:     logger,
		onboarding: onboarding,
	}
}

func (d *AppController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "app/onboarding-info", d.getOnboardingInfo, nil),
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
	Env              string   `json:"env"`
	Version          string   `json:"version"`
	Features         Features `json:"features"`
	SpireInitialized bool     `json:"is_spire_initialized"`
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
			SpireInitialized: d.onboarding.IsInitialized(),
		}

		return c.JSON(http.StatusOK, echo.Map{"data": response})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unknown error"})
}

func (d *AppController) getOnboardingInfo(c echo.Context) error {
	return c.JSON(http.StatusOK,
		echo.Map{
			"data": echo.Map{
				"connection_info": d.onboarding.GetConnectionInfo(),
				"tables":          d.onboarding.GetInstallationTables(),
			},
		},
	)
}
