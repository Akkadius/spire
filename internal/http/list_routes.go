package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sort"
)

// list internal routes
func listRoutes(c echo.Context) error {
	type Route struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}

	routes := make([]Route, 0)
	for _, r := range c.Echo().Routes() {
		routes = append(routes, Route{
			Method: r.Method,
			Path:   r.Path,
		})
	}

	sort.Slice(routes, func(i, j int) bool {
		if routes[i].Path != routes[j].Path {
			return routes[i].Path < routes[j].Path
		}

		return routes[i].Method < routes[j].Method
	})

	return c.JSON(http.StatusOK, routes)
}
