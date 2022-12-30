package controllers

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/spireuser"
	"github.com/k0kubun/pp/v3"
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
	spireuser  *spireuser.UserService
	settings   *spire.Settings
}

func NewAppController(
	cache *gocache.Cache,
	logger *logrus.Logger,
	onboarding *spire.SpireInit,
	spireuser *spireuser.UserService,
	settings *spire.Settings,
) *AppController {
	return &AppController{
		cache:      cache,
		logger:     logger,
		onboarding: onboarding,
		spireuser:  spireuser,
		settings:   settings,
	}
}

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

type OnboardInitializeAppRequestStruct struct {
	AuthEnabled int    `json:"auth_enabled"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

func (d *AppController) initializeApp(c echo.Context) error {
	// body - bind
	r := new(OnboardInitializeAppRequestStruct)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// new user
	user := models.User{
		UserName: r.Username,
		FullName: r.Username,
		Password: r.Password,
		Provider: spireuser.LOGIN_PROVIDER_LOCAL,
		IsAdmin:  true,
	}

	_, err := d.spireuser.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// auth
	if r.AuthEnabled == 1 {
		d.settings.EnableSetting(spire.SETTING_AUTH_ENABLED)
	}

	// re-initialize again as if we just started up the app
	d.onboarding.Init()

	pp.Println(r)

	return c.JSON(http.StatusOK, echo.Map{"data": "Ok"})
}
