package routes

import "github.com/labstack/echo/v4"

// RegisterRoute registers a new route
func RegisterRoute(method string, route string, handler echo.HandlerFunc, middlewares []echo.MiddlewareFunc) *Route {
	return &Route{
		method:      method,
		route:       route,
		handler:     handler,
		middlewares: middlewares,
	}
}

// Route represents a http route with
// its associated request handler
type Route struct {
	method      string
	route       string
	handler     echo.HandlerFunc
	middlewares []echo.MiddlewareFunc
}

// Method returns the HTTP method
func (r *Route) Method() string {
	return r.method
}

// Route returns the path of the route
func (r *Route) Route() string {
	return r.route
}

// Handler returns the handler method
func (r *Route) Handler() echo.HandlerFunc {
	return r.handler
}

// Middlewares returns the middlewares
// to be applied to the route
func (r *Route) Middlewares() []echo.MiddlewareFunc {
	if len(r.middlewares) > 0 {
		return r.middlewares
	}

	return []echo.MiddlewareFunc{}
}
