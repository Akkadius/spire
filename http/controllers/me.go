package controllers

import (
	"eoc/http/routes"
	"eoc/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MeController struct{}

func NewMeController() *MeController {
	return &MeController{}
}

func (a *MeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "me", a.me, nil),
	}
}

// me godoc
// @Id me
// @Summary Returns current logged in user information
// @Accept json
// @Produce json
// @Tags Me
// @Success 200 {array} models.User
// @Failure 500 {string} string "Bad query request"
// @Router /me [get]
func (a *MeController) me(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Get("user").(models.User))
}
