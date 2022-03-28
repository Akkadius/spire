package permissions

import (
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	"github.com/Akkadius/spire/internal/http"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/gertd/go-pluralize"
	"github.com/k0kubun/pp/v3"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"sort"
	"strings"
)

type ResourceList struct {
	logger *logrus.Logger
	router *routes.Router
	debug  bool
}

func NewResourceList(logger *logrus.Logger, router *routes.Router) *ResourceList {
	return &ResourceList{
		logger: logger,
		router: router,
		debug:  len(os.Getenv("DEBUG")) > 0,
	}
}

// RegisterManualResources
// these are routes that manipulate the database that do not fall under CRUD routes
func (c *ResourceList) RegisterManualResources() map[string][]string {
	return map[string][]string{
		"Client Files": {"client-file"},
	}
}

type Resource struct {
	name    string
	matches []string
}

// Get returns a complete list of resources
func (c *ResourceList) Get() []Resource {

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
	if c.debug {
		console.PrintBanner("Routes", bannerLength)
	}

	resources := map[string][]string{}
	pluralizeClient := pluralize.NewClient()
	resourceRouteCount := map[string]int{}

	// loop through routes
	for _, route := range r {
		if c.debug {
			fmt.Printf(
				"| %-*v | %-*v | \n",
				maxRouteCharacterLength,
				route.Path,
				maxRouteMethodsLength,
				strings.Join(routeMethods[route.Path], "|"),
			)
		}

		// match on v1 routes
		if strings.Contains(route.Path, "/api/v1/") {
			params := strings.Split(route.Path, "/")
			resource := ""
			if len(params) > 0 {
				resource = pluralizeClient.Singular(params[3])
			}

			// if resource not empty
			if resource != "" {

				// title case
				title := resource
				title = strings.ReplaceAll(title, "_", " ")
				title = strings.Title(title)

				// count routes per resource
				resourceRouteCount[title]++

				// if resource not added to map
				if _, ok := resources[title]; !ok {
					resources[title] = []string{
						resource,
						pluralizeClient.Plural(resource),
					}
				}
			}
		}

	}

	// delete non-crud resources
	// crud routes have at least 4 routes per resource
	for resource, _ := range resources {
		if resourceRouteCount[resource] < 4 {
			if c.debug {
				pp.Printf("Deleting resources [%v]\n", resource)
			}
			delete(resources, resource)
		}
	}

	// apply manually registered resources
	for resource, val := range c.RegisterManualResources() {
		resources[resource] = val
	}

	// list of resources
	res := []Resource{}
	for resource, val := range resources {
		res = append(res, Resource{
			name:    resource,
			matches: val,
		})
	}

	// sort resources
	sort.Slice(res, func(i, j int) bool {
		return res[i].name < res[j].name
	})

	if c.debug {
		console.PrintBanner(fmt.Sprintf("Total Routes [%v] Unique Routes [%v]", len(e.Routes()), len(r)), bannerLength)

		pp.Println("# Resources")
		pp.Println(res)
	}

	return res
}
