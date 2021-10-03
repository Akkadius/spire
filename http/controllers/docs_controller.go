package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DocsController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewDocsController(db *database.DatabaseResolver, logger *logrus.Logger) *DocsController {
	return &DocsController{
		db:     db,
		logger: logger,
	}
}

func (d *DocsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "doc/:page", d.page, nil),
	}
}

func (d *DocsController) page(c echo.Context) error {
	request, err := url.QueryUnescape(c.Param("page"))
	if err != nil {
		d.logger.Error(err)
	}

	filePath := fmt.Sprintf("%v/%v.md", "docs-quest-api", request)

	page, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{"data": string(page)})
}
