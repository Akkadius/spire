package controllers

import (
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/questapi"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type QuestApiController struct {
	logger     *logrus.Logger
	parser     *questapi.ParseService
	peqSourcer *questapi.QuestExamplesProjectEqSourcer
}

func NewQuestApiController(
	logger *logrus.Logger,
	parser *questapi.ParseService,
	peqSourcer *questapi.QuestExamplesProjectEqSourcer,
) *QuestApiController {
	return &QuestApiController{logger: logger, parser: parser, peqSourcer: peqSourcer}
}

func (d *QuestApiController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "quest-api/methods", d.methods, nil),
		routes.RegisterRoute(http.MethodPost, "quest-api/examples/search-projecteq", d.searchProjectEqExamples, nil),
	}
}

func (d *QuestApiController) methods(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"data": d.parser.Parse(false)})
}

type ProjectEqSearchTermRequest struct {
	SearchTerms []string `json:"search_terms"`
	Language    string   `json:"language"`
}

// searches projecteq quest examples
func (d *QuestApiController) searchProjectEqExamples(c echo.Context) error {
	p := new(ProjectEqSearchTermRequest)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": d.peqSourcer.Search(p.SearchTerms, p.Language),
		},
	)
}
