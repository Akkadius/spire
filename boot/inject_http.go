package boot

import (
	controllers2 "github.com/Akkadius/spire/internal/http/controllers"
	middleware2 "github.com/Akkadius/spire/internal/http/middleware"
	routes2 "github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/models"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gocache "github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

var httpSet = wire.NewSet(
	middleware2.NewUserContextMiddleware,
	middleware2.NewRequestLogMiddleware,
	controllers2.NewAnalyticsController,
	controllers2.NewHelloWorldController,
	controllers2.NewConnectionsController,
	controllers2.NewMeController,
	controllers2.NewAuthController,
	controllers2.NewDocsController,
	controllers2.NewQuestApiController,
	provideControllers,
	NewRouter,
)

type appControllerGroups struct {
	authControllers       []routes2.Controller
	helloWorldControllers []routes2.Controller
	v1controllers         []routes2.Controller
	v1controllersNoAuth   []routes2.Controller
}

func NewRouter(
	cg *appControllerGroups,
	crudc *crudControllers,
	userContextMiddleware *middleware2.UserContextMiddleware,
	logMiddleware *middleware2.RequestLogMiddleware,
	cache *gocache.Cache,
) *routes2.Router {
	return routes2.NewHttpRouter(

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
		[]*routes2.ControllerGroup{
			routes2.NewControllerGroup("/auth/", cg.authControllers, []echo.MiddlewareFunc{}...),
			routes2.NewControllerGroup("/api/v1/", cg.v1controllers, userContextMiddleware.Handle(), v1RateLimit()),
			routes2.NewControllerGroup(
				"/api/v1/",
				cg.v1controllersNoAuth,
				v1RateLimit(),
				middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}),
			),
			routes2.NewControllerGroup(
				"/api/v1/",
				crudc.routes,
				userContextMiddleware.Handle(),
				v1RateLimit(),
				middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}),
			),
		},
	)
}

// controllers provider
func provideControllers(
	hello *controllers2.HelloWorldController,
	auth *controllers2.AuthController,
	me *controllers2.MeController,
	analytics *controllers2.AnalyticsController,
	connections *controllers2.ConnectionsController,
	docs *controllers2.DocsController,
	quest *controllers2.QuestApiController,
) *appControllerGroups {
	return &appControllerGroups{
		authControllers: []routes2.Controller{
			auth,
		},
		v1controllers: []routes2.Controller{
			me,
			analytics,
			connections,
			hello,
			docs,
		},
		v1controllersNoAuth: []routes2.Controller{
			quest,
		},
	}
}

func v1RateLimit() echo.MiddlewareFunc {
	return middleware2.RateLimiterWithConfig(
		middleware2.RateLimiterConfig{
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
			LimitConfig: middleware2.LimiterConfig{
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
