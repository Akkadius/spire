package boot

import (
	"github.com/Akkadius/spire/http/controllers"
	appmiddleware "github.com/Akkadius/spire/http/middleware"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

var httpSet = wire.NewSet(
	appmiddleware.NewUserContextMiddleware,
	appmiddleware.NewRequestLogMiddleware,
	controllers.NewAnalyticsController,
	controllers.NewHelloWorldController,
	controllers.NewConnectionsController,
	controllers.NewMeController,
	controllers.NewAuthController,
	controllers.NewDocsController,
	controllers.NewQuestApiController,
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
	logMiddleware *appmiddleware.RequestLogMiddleware,
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
			middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
			}),
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
			routes.NewControllerGroup("/auth/", cg.authControllers, []echo.MiddlewareFunc{}...),
			routes.NewControllerGroup("/api/v1/", cg.v1controllers, userContextMiddleware.Handle(), v1RateLimit()),
			routes.NewControllerGroup(
				"/api/v1/",
				cg.v1controllersNoAuth,
				middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}),
				v1RateLimit(),
			),
			routes.NewControllerGroup("/api/v1/", crudc.routes, userContextMiddleware.Handle(), v1RateLimit()),
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
		},
		v1controllersNoAuth: []routes.Controller{
			quest,
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
				Max:      401,
				Duration: time.Second * 1,
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
