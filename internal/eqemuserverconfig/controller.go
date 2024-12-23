package eqemuserverconfig

import (
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	serverconfig *Config
}

func NewController(
	serverconfig *Config,
) *Controller {
	return &Controller{
		serverconfig: serverconfig,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "admin/serverconfig", a.get, nil),
		routes.RegisterRoute(http.MethodPost, "admin/serverconfig", a.save, nil),
		routes.RegisterRoute(http.MethodGet, "admin/launcherconfig", a.getLauncherConfig, nil),
		routes.RegisterRoute(http.MethodPost, "admin/launcherconfig", a.saveLauncherConfig, nil),
	}
}

func (a *Controller) get(c echo.Context) error {
	cfg, _ := a.serverconfig.Get()
	return c.JSON(http.StatusOK, cfg)
}

func (a *Controller) save(c echo.Context) error {
	var config EQEmuConfigJson
	err := c.Bind(&config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to bind config")
	}

	if len(config.Server.World.Longname) == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Configuration long name is empty!"})
	}

	if len(config.Server.World.Shortname) == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Configuration short name is empty!"})
	}

	err = a.serverconfig.Save(config)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, config)
}

type LauncherConfig struct {
	RunSharedMemory             bool   `json:"runSharedMemory"`
	RunLoginserver              bool   `json:"runLoginserver"`
	RunQueryServ                bool   `json:"runQueryServ"`
	RunUcs                      bool   `json:"runUcs"`
	IsRunning                   bool   `json:"isRunning"`
	SpireLauncherStart          bool   `json:"spireLauncherStart"`
	MinZoneProcesses            int    `json:"minZoneProcesses"`
	StaticZones                 string `json:"staticZones"`
	UpdateOpcodesOnStart        bool   `json:"updateOpcodesOnStart"`
	DeleteLogFilesOlderThanDays int    `json:"deleteLogFilesOlderThanDays"`
}

func (a *Controller) getLauncherConfig(c echo.Context) error {
	cfg, _ := a.serverconfig.Get()
	l := LauncherConfig{
		RunSharedMemory:             cfg.WebAdmin.Launcher.RunSharedMemory,
		RunLoginserver:              cfg.WebAdmin.Launcher.RunLoginserver,
		RunQueryServ:                cfg.WebAdmin.Launcher.RunQueryServ,
		RunUcs:                      cfg.WebAdmin.Launcher.RunUcs,
		IsRunning:                   cfg.WebAdmin.Launcher.IsRunning,
		SpireLauncherStart:          cfg.Spire.LauncherStart,
		MinZoneProcesses:            cfg.WebAdmin.Launcher.MinZoneProcesses,
		StaticZones:                 cfg.WebAdmin.Launcher.StaticZones,
		UpdateOpcodesOnStart:        cfg.WebAdmin.Launcher.UpdateOpcodesOnStart,
		DeleteLogFilesOlderThanDays: cfg.WebAdmin.Launcher.DeleteLogFilesOlderThanDays,
	}

	return c.JSON(http.StatusOK, l)
}

func (a *Controller) saveLauncherConfig(c echo.Context) error {
	var config LauncherConfig
	err := c.Bind(&config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to bind config")
	}

	cfg, _ := a.serverconfig.Get()
	cfg.WebAdmin.Launcher.RunSharedMemory = config.RunSharedMemory
	cfg.WebAdmin.Launcher.RunLoginserver = config.RunLoginserver
	cfg.WebAdmin.Launcher.RunQueryServ = config.RunQueryServ
	cfg.WebAdmin.Launcher.RunUcs = config.RunUcs
	cfg.WebAdmin.Launcher.IsRunning = config.IsRunning
	cfg.WebAdmin.Launcher.MinZoneProcesses = config.MinZoneProcesses
	cfg.WebAdmin.Launcher.StaticZones = config.StaticZones
	cfg.WebAdmin.Launcher.UpdateOpcodesOnStart = config.UpdateOpcodesOnStart
	cfg.WebAdmin.Launcher.DeleteLogFilesOlderThanDays = config.DeleteLogFilesOlderThanDays
	cfg.Spire.LauncherStart = config.SpireLauncherStart

	// both launchers shouldn't be running at the same time
	if cfg.Spire.LauncherStart {
		cfg.WebAdmin.Launcher.IsRunning = false
	}

	err = a.serverconfig.Save(cfg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, config)
}
