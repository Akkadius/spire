package controllers

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"runtime"
)

// AppController is the controller for the app
type AppController struct {
	cache     *gocache.Cache
	logger    *logrus.Logger
	spireinit *spire.Init
	spireuser *spire.UserService
	settings  *spire.Settings
}

// NewAppController returns a new app controller
func NewAppController(
	cache *gocache.Cache,
	logger *logrus.Logger,
	spireinit *spire.Init,
	spireuser *spire.UserService,
	settings *spire.Settings,
) *AppController {
	return &AppController{
		cache:     cache,
		logger:    logger,
		spireinit: spireinit,
		spireuser: spireuser,
		settings:  settings,
	}
}

// Routes returns the routes for the app controller
func (d *AppController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "app/onboarding-info", d.getOnboardingInfo, nil),
		routes.RegisterRoute(http.MethodPost, "app/onboard-initialize", d.initializeApp, nil),
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

// EnvResponse is a struct to hold the response for the env endpoint
type EnvResponse struct {
	Env              string           `json:"env"`
	Version          string           `json:"version"`
	OS               string           `json:"os"`
	Features         Features         `json:"features"`
	Settings         []models.Setting `json:"settings"`
	SpireInitialized bool             `json:"is_spire_initialized"`
}

// PackageJson is a struct to hold the package.json file
type PackageJson struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
}

// env returns the environment variables for the app
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
			OS:      runtime.GOOS,
			Version: pkg.Version,
			Features: Features{
				GithubAuthEnabled: len(os.Getenv("GITHUB_CLIENT_ID")) > 0,
			},
			Settings:         d.settings.GetSettings(),
			SpireInitialized: d.spireinit.IsInitialized(),
		}

		return c.JSON(http.StatusOK, echo.Map{"data": response})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unknown error"})
}

// getOnboardingInfo is used to get the spireinit info
func (d *AppController) getOnboardingInfo(c echo.Context) error {
	return c.JSON(http.StatusOK,
		echo.Map{
			"data": echo.Map{
				"connection_info": d.spireinit.GetConnectionInfo(),
				"tables":          d.spireinit.GetInstallationTables(),
			},
		},
	)
}

// initializeApp is used to initialize the app
func (d *AppController) initializeApp(c echo.Context) error {
	r := new(spire.InitAppRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Initialize the app
	err := d.spireinit.InitApp(r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": "Ok"})
}
