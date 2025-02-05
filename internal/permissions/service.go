package permissions

import (
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/k0kubun/pp/v3"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Service struct {
	logger    *logger.AppLogger
	db        *database.Resolver
	cache     *gocache.Cache
	pluralize *pluralize.Client
	debug     int
}

func NewService(
	db *database.Resolver,
	cache *gocache.Cache,
	logger *logger.AppLogger,
	pluralize *pluralize.Client,
) *Service {
	return &Service{
		db:        db,
		cache:     cache,
		logger:    logger,
		pluralize: pluralize,
		debug:     env.GetInt("PERMISSIONS_DEBUG", "0"),
	}
}

// RegisterManualResources
// these are routes that manipulate the database that do not fall under CRUD routes
func (s *Service) RegisterManualResources() map[string][]string {
	return map[string][]string{
		"Client Files": {"client-file"},

		// admin server
		"Server Configuration":    {"admin/serverconfig"},
		"Server System Resources": {"admin/system", "eqemuserver/system-all", "eqemuserver/get-websocket-auth"},
		"Server Update / Build": {
			"eqemuserver/build",
			"eqemuserver/install-release",
			"eqemuserver/get-build-info",
			"eqemuserver/update-type",
			"eqemuserver/version",
		},
		"Server Players Online":    {"eqemuserver/client-list"},
		"Server Dashboard Stats":   {"eqemuserver/dashboard-stats"},
		"Server File Logs":         {"eqemuserver/log"},
		"Server Manual Backup":     {"eqemuserver/manual-backup"},
		"Server Pre-Flight Checks": {"eqemuserver/pre-flight"},
		"Server Reload API":        {"eqemuserver/reload"},
		"Server Lock":              {"eqemuserver/get-lock-status", "eqemuserver/toggle-server-lock"},
		"Server Process Management": {
			"eqemuserver/server/start",
			"eqemuserver/server/stop",
			"eqemuserver/server/restart",
			"eqemuserver/server/stop-cancel",
		},
		"Server Process Stats": {"eqemuserver/server-stats"},
		"Server Zone Servers": {
			"eqemuserver/zone-list",
			"eqemuserver/zoneserver-list",
			"eqemuserver/server/process-kill",
		},
		"Spire Settings": {"spire/setting"},
	}
}

// Resource is a representation of a permission resource
// contains an identifier with route match prefixes
// @example
//
//	 permissions.Resource{
//	  Name:               "Npc Type",
//	  Identifier:         "NPC_TYPE",
//	  RouteMatchPrefixes: []string{
//	    "npc_type",
//	    "npc_types",
//	  },
//	},
type Resource struct {
	Name               string   `json:"name"`       // used for visual display in the UI
	Identifier         string   `json:"identifier"` // used for backend storage identification
	RouteMatchPrefixes []string `json:"-"`          // used for matching routes in middleware
	CanRead            bool     `json:"-"`          // internal: Used in middleware
	CanWrite           bool     `json:"-"`          // internal: Used in middleware
}

// GetResources returns a complete list of resources
// it automatically builds resources from echo routes
// also applies manual routes using RegisterManualResources()
func (s *Service) GetResources(routes []*echo.Route) []Resource {

	// below logic build routes from echo
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
	if s.debug >= 3 {
		console.PrintBanner("Routes", bannerLength)
	}

	resources := map[string][]string{}
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
		if s.debug >= 3 {
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
				resource = s.pluralize.Singular(params[3])
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

				addToMap := true

				// don't add auto resources for things that get
				// manually registered
				for _, manualResources := range s.RegisterManualResources() {
					for _, mr := range manualResources {
						// if the resource route path contains a manually registered route
						// lets exclude it from being added to the auto build map
						if strings.Contains(route.Path, mr) {
							addToMap = false
							continue
						}
					}
				}

				if addToMap {
					// if resource not added to map
					if _, ok := resources[title]; !ok {
						resources[title] = []string{
							resource,
							s.pluralize.Plural(resource),
						}
					}
				}
			}
		}

	}

	// delete non-crud resources
	// crud routes have at least 4 routes per resource
	for resource, _ := range resources {
		//pp.Println(resource, resourceRouteCount[resource])

		// contains([]string{"Quest-Api", "Changelog", "Analytics", "App", "Query"}, resource)

		if resourceRouteCount[resource] != 5 {
			if s.debug >= 3 {
				pp.Printf("Deleting resources [%v]\n", resource)
			}
			delete(resources, resource)
		}
	}

	// apply manually registered resources
	// ex "Client Files": {"client-file"},
	for resource, val := range s.RegisterManualResources() {
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

	if s.debug >= 3 {
		console.PrintBanner(fmt.Sprintf("Total Routes [%v] Unique Routes [%v]", len(routes), len(r)), bannerLength)

		//pp.Println("# Resources")
		//pp.Println(res)
	}

	return res
}

// root permissions object for users when running through access control list logic
// gets cached on first access
type userPermissions struct {
	isConnectionOwner bool
	canReadAll        bool
	canWriteAll       bool
	permissions       []Resource
}

func (s *Service) IsWriteRequest(c echo.Context) bool {
	return contains(
		[]string{
			http.MethodPatch,
			http.MethodPost,
			http.MethodPut,
		},
		c.Request().Method,
	) && !strings.Contains(c.Request().URL.Path, "/bulk")
}

func (s *Service) CanAccessResource(c echo.Context, user models.User, connectionId uint) bool {
	p := s.getUserPermissions(c, user, connectionId)

	if s.debug >= 3 {
		s.logger.Debug().
			Any("permissions", p).
			Any("user", user).
			Msg("Checking permissions")
	}

	// if connection owner, they can access everything
	if p.isConnectionOwner {
		return true
	}

	isWriteRequest := s.IsWriteRequest(c)

	// global level rules
	if p.canReadAll && !isWriteRequest {
		return true
	}
	if p.canWriteAll && isWriteRequest {
		return true
	}

	// match on v1 routes
	if strings.Contains(c.Request().URL.Path, "/api/v1/") {
		// the first parameter after `v1` is what we're describing as the "resource"
		params := strings.Split(c.Request().URL.Path, "/")
		resource := ""
		if len(params) > 0 {
			resource = strings.TrimSpace(params[3])
		}

		for _, up := range p.permissions {
			for _, prefix := range up.RouteMatchPrefixes {
				// first check looks for CRUD generated resources eg task/tasks
				// second check looks for things that would be in manually registered prefixes

				isManualRouteMatch := strings.Contains(c.Request().URL.Path, prefix) && strings.Contains(prefix, "/")

				if prefix == resource || isManualRouteMatch {
					if s.debug >= 3 {
						s.logger.Debug().
							Any("permissions", p).
							Any("user", user).
							Str("prefix", prefix).
							Str("route", c.Request().URL.Path).
							Msg("Found match for prefix")
					}

					if isWriteRequest && up.CanWrite {
						if s.debug >= 3 {
							s.logger.Debug().
								Any("permissions", p).
								Any("user", user).
								Str("prefix", prefix).
								Str("route", c.Request().URL.Path).
								Msg("Can write, passing rules")
						}

						return true
					}

					if !isWriteRequest && up.CanRead {
						if s.debug >= 3 {
							s.logger.Debug().
								Any("permissions", p).
								Any("user", user).
								Str("prefix", prefix).
								Str("route", c.Request().URL.Path).
								Msg("Can read, passing rules")
						}

						return true
					}
				}
			}
		}
	}

	return false
}

func (s *Service) ClearUserPermissionsCache(userId uint64) {
	if s.debug >= 3 {
		s.logger.Debug().Any("userId", userId).Msg("Clearing cache permissions for user id")
	}

	cacheKey := fmt.Sprintf("user-permissions-%v", userId)
	s.cache.Delete(cacheKey)
}

func (s *Service) getUserPermissions(c echo.Context, user models.User, connectionId uint) userPermissions {

	// return cached if exist
	cacheKey := fmt.Sprintf("user-permissions-%v", user.ID)
	cached, found := s.cache.Get(cacheKey)
	if found {
		return cached.(userPermissions)
	}

	var p userPermissions

	p.canReadAll = false
	p.canWriteAll = false

	// fetch permissions
	var userServerResourcePermissions []models.UserServerResourcePermission
	s.db.GetSpireDb().
		Model(&models.UserServerResourcePermission{}).
		Where("user_id = ? and server_database_connection_id = ?", user.ID, connectionId).
		Find(&userServerResourcePermissions)

	// validate ownership
	var serverDatabaseConn models.ServerDatabaseConnection
	_ = s.db.GetSpireDb().Where("created_by = ? and id = ?", user.ID, connectionId).First(&serverDatabaseConn).Error
	if serverDatabaseConn.ID > 0 {
		p.isConnectionOwner = true
	}

	for _, up := range userServerResourcePermissions {
		if up.ResourceName == "ALL" {
			if up.CanWrite == 1 {
				p.canWriteAll = true
			}
			if up.CanRead == 1 {
				p.canReadAll = true
			}
		}
	}

	// if connection owner, we don't need to check permissions
	if !p.isConnectionOwner {
		var userResourcePermissions []Resource
		resources := s.GetResources(c.Echo().Routes())
		for _, up := range userServerResourcePermissions {

			// detailed logging
			if s.debug >= 3 {
				s.logger.Debug().Msgf(
					"Server ID [%v] Resource [%v] Read [%v] Write [%v]\n",
					up.ServerDatabaseConnectionId,
					up.ResourceName,
					up.CanRead,
					up.CanWrite,
				)
			}

			// loop through resources and hydrate for user
			for _, r := range resources {
				if r.Identifier == up.ResourceName {
					r.CanRead = up.CanRead == 1
					r.CanWrite = up.CanWrite == 1

					userResourcePermissions = append(userResourcePermissions, r)
				}
			}
		}

		p.permissions = userResourcePermissions
	}

	s.cache.Set(cacheKey, p, 10*time.Minute)

	return p
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
