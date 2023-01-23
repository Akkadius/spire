package eqemuchangelog

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
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
		routes.RegisterRoute(http.MethodGet, "changelog/:days", a.getChangelog, nil),
	}
}

func (a *EqemuChangelogController) getChangelog(c echo.Context) error {
	daysParam := c.Param("days")
	days, err := strconv.Atoi(daysParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": a.changelog.BuildChangelog(
				a.changelog.getCommitsDaysBack(time.Duration(days)),
			),
		},
	)
}
