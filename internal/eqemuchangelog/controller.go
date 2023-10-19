package eqemuchangelog

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	logger    *logrus.Logger
	db        *database.Resolver
	changelog *Changelog
}

func NewController(
	logger *logrus.Logger,
	db *database.Resolver,
	changelog *Changelog,
) *Controller {
	return &Controller{
		logger:    logger,
		db:        db,
		changelog: changelog,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "changelog", a.getChangelog, nil),
	}
}

func (a *Controller) getChangelog(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": a.changelog.BuildChangelog(
				a.changelog.getCommitsDaysBack(),
			),
		},
	)
}
