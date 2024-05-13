package spire

import (
	"errors"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SettingsController struct {
	db       *database.Resolver
	settings *Settings
	crypt    *encryption.Encrypter
}

func NewSettingController(
	db *database.Resolver,
	crypt *encryption.Encrypter,
	settings *Settings,
) *SettingsController {
	return &SettingsController{
		db:       db,
		crypt:    crypt,
		settings: settings,
	}
}

func (s *SettingsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spire/settings", s.getSettings, nil),
		routes.RegisterRoute(http.MethodGet, "spire/setting/:setting", s.getSetting, nil),
		//routes.RegisterRoute(http.MethodDelete, "spire/setting/:setting", s.deleteSetting, nil),
		routes.RegisterRoute(http.MethodPost, "spire/setting/:setting/value/:value", s.create, nil),
		routes.RegisterRoute(http.MethodPatch, "spire/setting/:setting/value/:value", s.update, nil),
	}
}

func (s *SettingsController) create(c echo.Context) error {
	s.settings.SetSetting(c.Param("setting"), c.Param("value"))

	return nil
}

func (s *SettingsController) update(c echo.Context) error {
	if s.settings.GetSetting(c.Param("setting")).ID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": errors.New("setting does not exist").Error()})
	}

	s.settings.SetSetting(c.Param("setting"), c.Param("value"))

	return nil
}

func (s *SettingsController) getSetting(c echo.Context) error {
	if s.settings.GetSetting(c.Param("setting")).ID == 0 {
		return c.JSON(
			http.StatusBadRequest,
			echo.Map{"error": errors.New("setting does not exist").Error()},
		)
	}

	return c.JSON(
		http.StatusOK,
		s.settings.GetSetting(c.Param("setting")),
	)
}

func (s *SettingsController) getSettings(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		s.settings.GetSettings(),
	)
}
