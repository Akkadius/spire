package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UsersController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewUsersController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *UsersController {
	return &UsersController{
		db:     db,
		logger: logger,
	}
}

func (a *UsersController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "users", a.list, nil),
	}
}

func (a *UsersController) list(c echo.Context) error {
	var results []models.User
	if a.db.GetSpireDb() != nil {
		query := a.db.GetSpireDb().Model(&models.User{})

		// param
		search := c.QueryParam("q")
		if len(search) > 0 {
			query.Or("user_name LIKE ?", fmt.Sprintf("%%%v%%", search))
		}

		query.Limit(100)

		err := query.Find(&results).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		}

		return c.JSON(http.StatusOK, results)
	}

	return c.JSON(http.StatusOK, echo.Map{"error": "No users found"})
}
