package eqemuchangelog

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type EqemuChangelogController struct {
	logger    *logrus.Logger
	db        *database.DatabaseResolver
	changelog *Changelog
}

func NewEqemuChangelogController(
	logger *logrus.Logger,
	db *database.DatabaseResolver,
	changelog *Changelog,
) *EqemuChangelogController {
	return &EqemuChangelogController{
		logger:    logger,
		db:        db,
		changelog: changelog,
	}
}

func (a *EqemuChangelogController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "changelog", a.getChangelog, nil),
	}
}

func (a *EqemuChangelogController) getChangelog(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": a.changelog.BuildChangelog(
				a.changelog.getCommitsDaysBack(),
			),
		},
	)
}
