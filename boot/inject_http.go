package boot

import (
	"github.com/Akkadius/spire/internal/analytics"
	"github.com/Akkadius/spire/internal/app"
	"github.com/Akkadius/spire/internal/assets"
	"github.com/Akkadius/spire/internal/auth"
	"github.com/Akkadius/spire/internal/backup"
	"github.com/Akkadius/spire/internal/clientfiles"
	"github.com/Akkadius/spire/internal/eqemuanalytics"
	"github.com/Akkadius/spire/internal/eqemuchangelog"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	apphttp "github.com/Akkadius/spire/internal/http"
	"github.com/Akkadius/spire/internal/http/controllers"
	appmiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/http/staticmaps"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/Akkadius/spire/internal/query"
	"github.com/Akkadius/spire/internal/questapi"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/system"
	"github.com/Akkadius/spire/internal/user"
	"github.com/Akkadius/spire/internal/websocket"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

var httpSet = wire.NewSet(
	apphttp.NewServer,
	user.NewContextMiddleware,
	appmiddleware.NewRequestLogMiddleware,
	appmiddleware.NewReadOnlyMiddleware,
	appmiddleware.NewPermissionsMiddleware,
	appmiddleware.NewLocalUserAuthMiddleware,
	analytics.NewController,
	controllers.NewHelloWorldController,
	controllers.NewConnectionsController,
	user.NewMeController,
	auth.NewController,
	questapi.NewController,
	app.NewController,
	query.NewController,
	eqemuanalytics.NewController,
	eqemuanalytics.NewAuthedController,
	eqemuchangelog.NewController,
	clientfiles.NewController,
	assets.NewController,
	permissions.NewController,
	user.NewController,
	spire.NewSettingController,
	staticmaps.NewStaticMapController,
	eqemuserver.NewController,
	eqemuserver.NewPublicController,
	eqemuserverconfig.NewController,
	websocket.NewController,
	backup.NewController,
	system.NewController,
	models.NewController,
	provideControllers,
	NewRouter,
)

type appControllerGroups struct {
	authControllers            []routes.Controller
	helloWorldControllers      []routes.Controller
	v1controllersNoPermissions []routes.Controller
	v1controllers              []routes.Controller
	v1controllersNoAuth        []routes.Controller
	v1Analytics                []routes.Controller
}

func NewRouter(
	cg *appControllerGroups,
	crudc *crudControllers,
	userContextMiddleware *user.ContextMiddleware,
	readOnlyModeMiddleware *appmiddleware.ReadOnlyMiddleware,
	permissionsMiddleware *appmiddleware.PermissionsMiddleware,
	logMiddleware *appmiddleware.RequestLogMiddleware,
	localUserAuthMiddleware *appmiddleware.LocalUserAuthMiddleware,
	assets *assets.SpireAssets,
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
			assets.ServeStatic(),
		},

		// controller groups
		[]*routes.ControllerGroup{
			routes.NewControllerGroup(
				"/auth/",
				cg.authControllers,
				[]echo.MiddlewareFunc{}...,
			),
			// non-crud v1 routes that don't need
			// to go through permissions middleware
			routes.NewControllerGroup(
				"/api/v1/",
				cg.v1controllersNoPermissions,
				userContextMiddleware.HandleHeader(),
				userContextMiddleware.HandleQuerystring(),
				localUserAuthMiddleware.Handle(),
				readOnlyModeMiddleware.Handle(),
				v1RateLimit(),
			),
			// non-crud v1 routes subject to permissions
			routes.NewControllerGroup(
				"/api/v1/",
				cg.v1controllers,
				userContextMiddleware.HandleHeader(),
				userContextMiddleware.HandleQuerystring(),
				localUserAuthMiddleware.Handle(),
				readOnlyModeMiddleware.Handle(),
				permissionsMiddleware.Handle(),
				v1RateLimit(),
			),
			// v1 controllers that require no auth
			routes.NewControllerGroup(
				"/api/v1/",
				cg.v1controllersNoAuth,
				v1RateLimit(),
				middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}),
			),
			// v1 analytics
			routes.NewControllerGroup(
				"/api/v1/",
				cg.v1Analytics,
				middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}),
			),
			// v1 crud code generated routes
			routes.NewControllerGroup(
				"/api/v1/",
				crudc.routes,
				userContextMiddleware.HandleHeader(),
				readOnlyModeMiddleware.Handle(),
				localUserAuthMiddleware.Handle(),
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
	auth *auth.Controller,
	me *user.MeController,
	analytics *analytics.Controller,
	connections *controllers.ConnectionsController,
	quest *questapi.Controller,
	app *app.Controller,
	query *query.Controller,
	clientFilesController *clientfiles.Controller,
	staticMaps *staticmaps.StaticMapController,
	analyticsController *eqemuanalytics.Controller,
	authedAnalyticsController *eqemuanalytics.AuthedController,
	changelogController *eqemuchangelog.Controller,
	assetsController *assets.Controller,
	permissionsController *permissions.Controller,
	usersController *user.Controller,
	settingsController *spire.SettingsController,
	eqemuserverController *eqemuserver.Controller,
	eqemuserverPublicController *eqemuserver.PublicController,
	serverconfigController *eqemuserverconfig.Controller,
	backupController *backup.Controller,
	websocketController *websocket.Controller,
	systemController *system.Controller,
	modelController *models.Controller,
) *appControllerGroups {
	return &appControllerGroups{
		authControllers: []routes.Controller{
			auth,
		},
		v1controllersNoPermissions: []routes.Controller{
			me,
			analytics,
			connections,
			hello,
			query,
		},
		v1controllers: []routes.Controller{
			settingsController,
			clientFilesController,
			permissionsController,
			usersController,
			eqemuserverController,
			serverconfigController,
			backupController,
			websocketController,
			systemController,
			authedAnalyticsController,
		},
		v1controllersNoAuth: []routes.Controller{
			quest,
			app,
			staticMaps,
			assetsController,
			changelogController,
			eqemuserverPublicController,
			modelController,
		},
		v1Analytics: []routes.Controller{
			analyticsController,
		},
	}
}

func v1RateLimit() echo.MiddlewareFunc {
	return appmiddleware.RateLimiterWithConfig(
		appmiddleware.RateLimiterConfig{
			Skipper: func(c echo.Context) bool {

				// if there is a valid user - skip the middleware
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
