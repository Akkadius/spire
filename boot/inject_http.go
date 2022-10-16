package boot

import (
	"github.com/Akkadius/spire/internal/http/controllers"
	appmiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/http/staticmaps"
	"github.com/Akkadius/spire/internal/models"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gocache "github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var httpSet = wire.NewSet(
	appmiddleware.NewUserContextMiddleware,
	appmiddleware.NewRequestLogMiddleware,
	appmiddleware.NewReadOnlyMiddleware,
	appmiddleware.NewPermissionsMiddleware,
	controllers.NewAnalyticsController,
	controllers.NewHelloWorldController,
	controllers.NewConnectionsController,
	controllers.NewMeController,
	controllers.NewAuthController,
	controllers.NewDocsController,
	controllers.NewQuestApiController,
	controllers.NewAppController,
	controllers.NewQueryController,
	controllers.NewQuestFileApiController,
	controllers.NewClientFilesController,
	controllers.NewAssetsController,
	controllers.NewPermissionsController,
	controllers.NewUsersController,
	staticmaps.NewStaticMapController,
	provideControllers,
	NewRouter,
)

type appControllerGroups struct {
	authControllers       []routes.Controller
	helloWorldControllers []routes.Controller
	v1controllers         []routes.Controller
	v1controllersNoAuth   []routes.Controller
}

func NewRouter(
	cg *appControllerGroups,
	crudc *crudControllers,
	userContextMiddleware *appmiddleware.UserContextMiddleware,
	readOnlyModeMiddleware *appmiddleware.ReadOnlyMiddleware,
	permissionsMiddleware *appmiddleware.PermissionsMiddleware,
	logMiddleware *appmiddleware.RequestLogMiddleware,
	cache *gocache.Cache,
) *routes.Router {
	return routes.NewHttpRouter(

		// pre middleware
		[]echo.MiddlewareFunc{
			middleware.RemoveTrailingSlash(),
		},

		// post middleware
		[]echo.MiddlewareFunc{
			logMiddleware.Handle(),
			//middleware.Logger(), // json logger
			middleware.Recover(),
			middleware.CORSWithConfig(
				middleware.CORSConfig{
					AllowOrigins: []string{"*"},
					AllowMethods: []string{
						http.MethodGet,
						http.MethodHead,
						http.MethodPut,
						http.MethodPatch,
						http.MethodPost,
						http.MethodDelete,
						http.MethodOptions,
					},
				},
			),
		},

		// controller groups
		[]*routes.ControllerGroup{
			routes.NewControllerGroup(
				"/auth/",
				cg.authControllers,
				[]echo.MiddlewareFunc{}...,
			),
			routes.NewControllerGroup(
				"/api/v1/",
				cg.v1controllers,
				userContextMiddleware.HandleHeader(),
				userContextMiddleware.HandleQuerystring(),
				readOnlyModeMiddleware.Handle(),
				v1RateLimit(),
			),
			routes.NewControllerGroup(
				"/api/v1/",
				cg.v1controllersNoAuth,
				v1RateLimit(),
				middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}),
			),
			routes.NewControllerGroup(
				"/api/v1/",
				crudc.routes,
				userContextMiddleware.HandleHeader(),
				readOnlyModeMiddleware.Handle(),
				permissionsMiddleware.Handle(),
				v1RateLimit(),
				middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}),
			),
		},
	)
}

// controllers provider
func provideControllers(
	hello *controllers.HelloWorldController,
	auth *controllers.AuthController,
	me *controllers.MeController,
	analytics *controllers.AnalyticsController,
	connections *controllers.ConnectionsController,
	docs *controllers.DocsController,
	quest *controllers.QuestApiController,
	app *controllers.AppController,
	query *controllers.QueryController,
	questFileApi *controllers.QuestFileApiController,
	clientFilesController *controllers.ClientFilesController,
	staticMaps *staticmaps.StaticMapController,
	assetsController *controllers.AssetsController,
	permissionsController *controllers.PermissionsController,
	usersController *controllers.UsersController,
) *appControllerGroups {
	return &appControllerGroups{
		authControllers: []routes.Controller{
			auth,
		},
		v1controllers: []routes.Controller{
			me,
			analytics,
			connections,
			hello,
			docs,
			query,
			clientFilesController,
			permissionsController,
			usersController,
		},
		v1controllersNoAuth: []routes.Controller{
			quest,
			app,
			questFileApi,
			staticMaps,
			assetsController,
		},
	}
}

func v1RateLimit() echo.MiddlewareFunc {
	return appmiddleware.RateLimiterWithConfig(
		appmiddleware.RateLimiterConfig{
			Skipper: func(c echo.Context) bool {

				// if there is a validate user - skip the middleware
				user, ok := c.Get("user").(models.User)
				if ok {
					if user.ID > 0 {
						return true
					}
				}

				return false
			},
			LimitConfig: appmiddleware.LimiterConfig{
				Max:      5000,
				Duration: time.Minute * 1,
				Strategy: "ip",
				Key:      "",
			},
			Prefix:                       "LIMIT",
			Client:                       nil,
			SkipRateLimiterInternalError: false,
			OnRateLimit:                  nil,
		},
	)
}
