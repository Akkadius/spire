package controllers

import (
	"github.com/Akkadius/spire/http/routes"
	"fmt"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HelloWorldController struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewHelloWorldController(db *gorm.DB, logger *logrus.Logger) *HelloWorldController {
	return &HelloWorldController{db: db, logger: logger}
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
