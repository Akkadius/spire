package controllers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PermissionsController struct {
	logger      *logrus.Logger
	db          *database.DatabaseResolver
	permissions *permissions.Service
}

func NewPermissionsController(
	logger *logrus.Logger,
	db *database.DatabaseResolver,
	permissions *permissions.Service,
) *PermissionsController {
	return &PermissionsController{
		logger:      logger,
		db:          db,
		permissions: permissions,
	}
}

func (p *PermissionsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "permissions/resources", p.getPermissionResources, nil),
	}
}

func (p *PermissionsController) getPermissionResources(c echo.Context) error {
	return c.JSON(http.StatusOK, p.permissions.GetResources(c.Echo().Routes()))
}
