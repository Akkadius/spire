package models

import (
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Controller is the controller for the app
type Controller struct {
}

// NewController returns a new app controller
func NewController() *Controller {
	return &Controller{}
}

// Routes returns the routes for the app controller
func (d *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "models", d.getModels, nil),
	}
}

func (d *Controller) getModels(c echo.Context) error {
	models := GetModels()
	modelNames := GetModelNames()

	type Model struct {
		Table         string   `json:"table"`
		Name          string   `json:"model_name"`
		Relationships []string `json:"relationships"`
	}

	var modelList []Model
	for i, model := range models {
		modelList = append(modelList, Model{
			Table:         model.TableName(),
			Name:          modelNames[i],
			Relationships: model.Relationships(),
		})
	}

	return c.JSON(http.StatusOK, modelList)
}
