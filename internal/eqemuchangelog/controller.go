package eqemuchangelog

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	db        *database.Resolver
	changelog *Changelog
}

func NewController(
	db *database.Resolver,
	changelog *Changelog,
) *Controller {
	return &Controller{
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
