package cmd

import (
	"fmt"
	"github.com/Akkadius/spire/console"
	"github.com/Akkadius/spire/internal/http"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

type RoutesListCommand struct {
	router  *routes.Router
	command *cobra.Command
	logger  *logrus.Logger
}

func (c *RoutesListCommand) Command() *cobra.Command {
	return c.command
}

func NewRoutesListCommand(router *routes.Router, logger *logrus.Logger) *RoutesListCommand {
	i := &RoutesListCommand{
		logger: logger,
		command: &cobra.Command{
			Use:   "routes:list",
			Short: "Lists application routes",
		},
		router: router,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *RoutesListCommand) Handle(_ *cobra.Command, _ []string) {

	// bring up echo instance
	e := echo.New()
	http.BootstrapMiddleware(e, c.router)
	if err := http.BootstrapControllers(e, c.router.ControllerGroups()...); err != nil {
		c.logger.Fatal(err)
	}

	type Route struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	}

	routeDefined := make(map[string]bool)
	routeMethods := make(map[string][]string)
	maxRouteCharacterLength := 0
	maxRouteMethodsLength := 0

	// Loop through echo r
	r := make([]Route, 0)
	for _, route := range e.Routes() {

		// Only unique append r once
		if !routeDefined[route.Path] {
			r = append(r, Route{
				Path: route.Path,
			})

			routeDefined[route.Path] = true
		}

		// calculate max character length for terminal output
		if len(route.Path) > maxRouteCharacterLength {
			maxRouteCharacterLength = len(route.Path)
		}

		// Append multiple methods to the same path
		routeMethods[route.Path] = append(routeMethods[route.Path], route.Method)

		// calculate max character length for terminal output
		methods := strings.Join(routeMethods[route.Path], "|")
		if len(methods) > maxRouteMethodsLength {
			maxRouteMethodsLength = len(methods)
		}
	}

	// sort r by path name
	sort.Slice(r, func(i, j int) bool {
		if r[i].Path != r[j].Path {
			return r[i].Path < r[j].Path
		}

		return r[i].Method < r[j].Method
	})

	// terminal output
	bannerLength := maxRouteMethodsLength + maxRouteCharacterLength + 7
	console.PrintBanner("Routes", bannerLength)

	for _, route := range r {
		fmt.Printf(
			"| %-*v | %-*v | \n",
			maxRouteCharacterLength,
			route.Path,
			maxRouteMethodsLength,
			strings.Join(routeMethods[route.Path], "|"),
		)
	}

	console.PrintBanner(fmt.Sprintf("Total Routes [%v] Unique Routes [%v]", len(e.Routes()), len(r)), bannerLength)
}

// Validate implementation of the Command interface
func (c *RoutesListCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
