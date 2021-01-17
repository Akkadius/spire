package routes

import (
	"github.com/labstack/echo/v4"
)

func NewHttpRouter(
	globalPreMiddlewares []echo.MiddlewareFunc,
	globalMiddlewares []echo.MiddlewareFunc,
	controllerGroups []*ControllerGroup,
) *Router {
	return &Router{
		globalPreMiddlewares: globalPreMiddlewares,
		globalMiddlewares:    globalMiddlewares,
		controllerGroups:     controllerGroups,
	}
}

type Router struct {
	globalPreMiddlewares []echo.MiddlewareFunc
	globalMiddlewares    []echo.MiddlewareFunc
	controllerGroups     []*ControllerGroup
}

func (r *Router) GlobalPreMiddlewares() []echo.MiddlewareFunc {
	return r.globalPreMiddlewares
}

func (r *Router) GlobalMiddlewares() []echo.MiddlewareFunc {
	return r.globalMiddlewares
}

func (r *Router) ControllerGroups() []*ControllerGroup {
	return r.controllerGroups
}
