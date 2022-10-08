package permissions

import (
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/k0kubun/pp/v3"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"sort"
	"strings"
)

type ResourceList struct {
	logger *logrus.Logger
	debug  bool
}

func NewResourceList(logger *logrus.Logger) *ResourceList {
	return &ResourceList{
		logger: logger,
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
	Name               string   `json:"name"`       // used for visual display in the UI
	Identifier         string   `json:"identifier"` // used for backend storage identification
	RouteMatchPrefixes []string `json:"-"`          // used for matching routes in middleware
}

// GetResources returns a complete list of resources
//
// @example
//   permissions.Resource{
//    Name:               "Npc Type",
//    Identifier:         "NPC_TYPE",
//    RouteMatchPrefixes: []string{
//      "npc_type",
//      "npc_types",
//    },
//  },
func (c *ResourceList) GetResources(routes []*echo.Route) []Resource {
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
	for _, route := range routes {

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

	// sort r by path Name
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

	// manual string transformations for resource names
	t := make(map[string]string)
	t["Ip"] = "IP"
	t["Aa"] = "AA"
	t["Gm"] = "GM"
	t["Id"] = "ID"

	// loop through routes
	for _, route := range r {

		// Print routes preview
		// @example
		// | /api/v1/zones      | GET  |
		// | /api/v1/zones/bulk | POST |
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
			// the first parameter after `v1` is what we're describing as the "resource"
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
				title = cases.Title(language.English).String(title)

				// manual transforms
				for o, n := range t {
					title = strings.ReplaceAll(title, o, n)
				}

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
	var res []Resource
	for resource, val := range resources {
		res = append(res, Resource{
			Name:               resource,
			RouteMatchPrefixes: val,
			Identifier:         strcase.ToScreamingSnake(resource),
		})
	}

	// sort resources by name
	sort.Slice(res, func(i, j int) bool {
		return res[i].Name < res[j].Name
	})

	if c.debug {
		console.PrintBanner(fmt.Sprintf("Total Routes [%v] Unique Routes [%v]", len(routes), len(r)), bannerLength)

		pp.Println("# Resources")
		pp.Println(res)
	}

	return res
}
