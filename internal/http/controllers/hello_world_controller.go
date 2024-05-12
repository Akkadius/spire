package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type HelloWorldController struct {
	db *gorm.DB
}

func NewHelloWorldController(db *gorm.DB) *HelloWorldController {
	return &HelloWorldController{db: db}
}

func (h *HelloWorldController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "hello-world", h.HelloWorld, nil),
	}
}

func (h *HelloWorldController) HelloWorld(c echo.Context) error {
	fmt.Println("HELLO")

	return c.JSON(http.StatusOK, echo.Map{"message": "hello!"})
}
