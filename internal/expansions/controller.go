package expansions

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	db  *database.Resolver
	api *eqemuserver.Client
}

func NewController(
	db *database.Resolver,
	api *eqemuserver.Client,
) *Controller {
	return &Controller{
		db:  db,
		api: api,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "expansions", a.get, nil),
		routes.RegisterRoute(http.MethodPost, "expansion", a.post, nil),
	}
}

func (a *Controller) get(c echo.Context) error {
	return c.JSON(http.StatusOK, expansions)
}

type SetExpansionRequest struct {
	Expansion int `json:"expansion"`
	Rules     []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"rules"`
}

func (a *Controller) post(c echo.Context) error {
	r := new(SetExpansionRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var results []models.RuleValue
	err := a.db.QueryContext(models.RuleValue{}, c).Find(&results).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var newRules []models.RuleValue
	for _, rule := range r.Rules {
		for _, result := range results {
			if result.RuleName == rule.Name {
				newRules = append(newRules, models.RuleValue{
					RulesetId: result.RulesetId,
					RuleName:  rule.Name,
					RuleValue: rule.Value,
					Notes:     result.Notes,
				})
			}
		}
	}

	// update changed rules in bulk
	err = a.db.QueryContext(models.RuleValue{}, c).Select("*").Save(&newRules).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	reloads := []string{"content_flags"}

	for _, reload := range reloads {
		_, _ = a.api.Reload(reload)
	}

	return c.JSON(http.StatusOK, "Expansion set successfully")
}
