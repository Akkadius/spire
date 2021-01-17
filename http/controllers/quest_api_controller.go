package controllers

import (
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/questapi"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type QuestApiController struct {
	logger *logrus.Logger
	parser *questapi.ParseService
}

func NewQuestApiController(logger *logrus.Logger, parser *questapi.ParseService) *QuestApiController {
	return &QuestApiController{logger: logger, parser: parser}
}

func (d *QuestApiController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "quest-api/methods", d.methods, nil),
	}
}

func (d *QuestApiController) methods(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"data": d.parser.Parse(false)})
}
