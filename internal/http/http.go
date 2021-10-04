package http

import (
	"errors"
	"fmt"
	"github.com/Akkadius/spire/docs"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/http/spa"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"unicode"

	_ "github.com/Akkadius/spire/docs"
)

// @title Spire
// @version 3.0
// @description Spire API documentation

// @contact.name Akkadius
// @contact.url TODO
// @contact.email akkadius1@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @comment-host localhost:3000 // this is commented to exclude host for now
// @BasePath /api/v1

// Serve runs a http server
func Serve(port uint, logger *logrus.Logger, router *routes.Router) error {
	e := echo.New()

	BootstrapMiddleware(e, router)
	if err := BootstrapControllers(e, router.ControllerGroups()...); err != nil {
		logger.Fatal(err)
	}

	e.GET("/swagger/*", docs.WrapHandler)
	e.GET("/api/v1/routes", listRoutes)

	// serve spa as embedded static assets
	s := spa.NewSpirePackagedSpaService(logger)
	e.GET("/*", s.Spa().Handler())
	e.Use(s.Spa().MiddlewareHandler())

	e.HTTPErrorHandler = errorHandler

	return e.Start(fmt.Sprintf(":%v", port))
}

func first(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func errorHandler(err error, c echo.Context) {
	_ = c.JSON(http.StatusUnprocessableEntity, echo.Map{"error": first(err.Error())})
}

func BootstrapMiddleware(e *echo.Echo, router *routes.Router) {
	for _, m := range router.GlobalPreMiddlewares() {
		e.Pre(m)
	}
	for _, m := range router.GlobalMiddlewares() {
		e.Use(m)
	}
}

func BootstrapControllers(e *echo.Echo, controllerGroups ...*routes.ControllerGroup) error {
	for _, cg := range controllerGroups {
		g := e.Group(cg.RoutePrefix(), cg.Middlewares()...)
		for _, c := range cg.Controllers() {
			if err := registerRoutes(c, g); err != nil {
				return err
			}
		}
	}

	return nil
}

func registerRoutes(controller routes.Controller, g *echo.Group) error {
	for _, r := range controller.Routes() {
		switch r.Method() {
		case http.MethodGet:
			g.GET(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodPost:
			g.POST(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodPut:
			g.PUT(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodPatch:
			g.PATCH(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodDelete:
			g.DELETE(r.Route(), r.Handler(), r.Middlewares()...)
		default:
			return errors.New("invalid r method specified")
		}
	}

	return nil
}
