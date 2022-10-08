package routes

import (
	"github.com/labstack/echo/v4"
)

// Controller is the contract to be implemented
// by http request handler structs
type Controller interface {
	// Routes registers controller specific routes
	Routes() []*Route
}

// NewControllerGroup wraps controllers and their accompanied middleware
func NewControllerGroup(
	routePrefix string,
	controllers []Controller,
	middlewares ...echo.MiddlewareFunc,
) *ControllerGroup {
	return &ControllerGroup{
		routePrefix: routePrefix,
		controllers: controllers,
		middlewares: middlewares,
	}
}

type ControllerGroup struct {
	routePrefix string
	controllers []Controller
	middlewares []echo.MiddlewareFunc
}

func (c *ControllerGroup) RoutePrefix() string {
	return c.routePrefix
}

func (c ControllerGroup) Controllers() []Controller {
	return c.controllers
}

func (c *ControllerGroup) Middlewares() []echo.MiddlewareFunc {
	return c.middlewares
}
