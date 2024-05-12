package questapi

import (
	"fmt"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	parser  *ParseService
	sourcer *ExamplesGithubSourcer
}

func NewController(
	parser *ParseService,
	sourcer *ExamplesGithubSourcer,
) *Controller {
	return &Controller{parser: parser, sourcer: sourcer}
}

func (d *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "quest-api/definitions", d.getQuestDefinitions, nil),
		routes.RegisterRoute(http.MethodGet, "quest-api/vscode-snippets", d.getSnippets, nil),
		routes.RegisterRoute(http.MethodPost, "quest-api/webhook-update-vscode-snippets", d.webhookVscodeSnippetsUpdate, nil),
		routes.RegisterRoute(http.MethodPost, "quest-api/refresh-definitions", d.webhookSourceDefinitionsUpdateApi, nil),
		routes.RegisterRoute(http.MethodPost, "quest-api/webhook-update-api", d.webhookSourceDefinitionsUpdateApi, nil),
		routes.RegisterRoute(http.MethodPost, "quest-api/webhook-update-source-examples/org/:org/repo/:repo/branch/:branch", d.webhookSourceExamplesUpdateApi, nil),
		routes.RegisterRoute(
			http.MethodPost,
			"quest-api/source-examples/org/:org/repo/:repo/branch/:branch",
			d.searchGithubExamples,
			nil,
		),
	}
}

func (d *Controller) getQuestDefinitions(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"data": d.parser.Parse(false)})
}

func (d *Controller) getSnippets(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"data": d.parser.GetSnippets()})
}

type SearchTermRequest struct {
	SearchTerms []string `json:"search_terms"`
	Language    string   `json:"language"`
}

// searches quest examples
func (d *Controller) searchGithubExamples(c echo.Context) error {
	// body - bind
	p := new(SearchTermRequest)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// params
	org := c.Param("org")
	repo := c.Param("repo")
	branch := c.Param("branch")

	// result
	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": d.sourcer.Search(org, repo, branch, p.SearchTerms, p.Language, false),
		},
	)
}

// ingests a webhook from Github and updates the repo data locally
func (d *Controller) webhookSourceDefinitionsUpdateApi(c echo.Context) error {
	// todo: verify signature later

	fmt.Println("Received definitions update request...")

	isGithubRequest := c.Request().Header.Get("X-Github-Event") != "" &&
		c.Request().Header.Get("X-Github-Delivery") != ""

	if isGithubRequest && env.IsAppEnvProduction() {
		d.parser.Parse(true)
	}

	if !isGithubRequest && env.IsAppEnvLocal() {
		d.parser.Parse(true)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": "Ok"})
}

// ingests a webhook from Github and updates the repo data locally
func (d *Controller) webhookVscodeSnippetsUpdate(c echo.Context) error {
	// todo: verify signature later

	fmt.Println("Received vscode quest snippets update request...")

	isGithubRequest := c.Request().Header.Get("X-Github-Event") != "" &&
		c.Request().Header.Get("X-Github-Delivery") != ""

	if isGithubRequest && env.IsAppEnvProduction() {
		d.parser.SourceSnippets("EQEmu", "spire-quest-snippets", "main", true)
	}

	if !isGithubRequest && env.IsAppEnvLocal() {
		d.parser.SourceSnippets("EQEmu", "spire-quest-snippets", "main", true)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": "Ok"})
}

// ingests a webhook from Github and updates the repo data locally
func (d *Controller) webhookSourceExamplesUpdateApi(c echo.Context) error {
	// todo: verify signature later
	if c.Request().Header.Get("X-Github-Event") != "" &&
		c.Request().Header.Get("X-Github-Delivery") != "" {
		// params
		org := c.Param("org")
		repo := c.Param("repo")
		branch := c.Param("branch")

		d.sourcer.Source(org, repo, branch, true)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": "Ok"})
}
