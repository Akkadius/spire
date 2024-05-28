package http

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/banner"
	"github.com/Akkadius/spire/internal/console"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"github.com/Akkadius/spire/internal/imgcat"
	"github.com/Akkadius/spire/internal/logger"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"unicode"

	"github.com/Akkadius/spire/docs"
	_ "github.com/Akkadius/spire/docs"
	spiremiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/http/spa"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	logger   *logger.AppLogger
	router   *routes.Router
	watcher  *eqemuserver.QuestHotReloadWatcher
	launcher *eqemuserver.Launcher
}

func NewServer(
	logger *logger.AppLogger,
	router *routes.Router,
	watcher *eqemuserver.QuestHotReloadWatcher,
	launcher *eqemuserver.Launcher,
) *Server {
	return &Server{
		logger:   logger,
		router:   router,
		watcher:  watcher,
		launcher: launcher,
	}
}

// @title Spire
// @version 3.0
// @description Spire API documentation

// @contact.name Akkadius
// @contact.url TODO
// @contact.email akkadius1@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @comment-host localhost:3000 // this is commented to exclude host for now
// @BasePath /api/v1

// Serve runs a http server
func (c *Server) Serve(port uint) error {
	e := echo.New()

	env.SetAppModeWebserver()

	BootstrapMiddleware(e, c.router)
	if err := BootstrapControllers(e, c.router.ControllerGroups()...); err != nil {
		c.logger.Fatal().Err(err).Msg("Failed to bootstrap controllers")
	}

	// basic auth if env passed
	if len(os.Getenv("BASIC_AUTH_USER")) > 0 && len(os.Getenv("BASIC_AUTH_PASSWORD")) > 0 {
		e.Use(spiremiddleware.BasicAuthWithConfig(spiremiddleware.BasicAuthConfig{
			Validator: func(username string, password string, c echo.Context) (bool, error) {
				// Be careful to use constant time comparison to prevent timing attacks
				if subtle.ConstantTimeCompare([]byte(username), []byte(os.Getenv("BASIC_AUTH_USER"))) == 1 &&
					subtle.ConstantTimeCompare([]byte(password), []byte(os.Getenv("BASIC_AUTH_PASSWORD"))) == 1 {
					return true, nil
				}
				return false, nil
			},
			Realm: "basic",
		}))
	}

	e.Use(spiremiddleware.LoggerWithConfig(spiremiddleware.LoggerConfig{
		//Format: "[${time_rfc3339}] [${status}] [${method}] [${uri}] [${latency_human}]\n",
		Format: fmt.Sprintf(
			"%sSpire › API ›%s [${status}] [${method}] [${uri}] [${latency_human}]\n",
			console.BoldWhite,
			console.Reset,
		),
		Output: e.Logger.Output(),
	}))

	e.GET("/swagger/*", docs.WrapHandler)
	e.GET("/api/v1/routes", routes.List)

	// Proxy requests to eqsage live site

	// Setup proxy for /eqsage with path rewrite
	eqsageTarget, _ := url.Parse("https://eqsage.vercel.app")
	e.Group("/eqsage").Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			proxy := httputil.NewSingleHostReverseProxy(eqsageTarget)
			proxy.Director = func(req *http.Request) {
				originalPath := req.URL.Path
				targetPath := regexp.MustCompile("^/eqsage").ReplaceAllString(originalPath, "")
				req.URL.Scheme = eqsageTarget.Scheme
				req.URL.Host = eqsageTarget.Host
				req.URL.Path = targetPath
				if targetPath == "" {
					req.URL.Path = "/"
				}
				// Set the Host header to the target host
				req.Host = eqsageTarget.Host
			}
			proxy.ServeHTTP(c.Response(), c.Request())
			return nil
		}
	})

	// Proxy middleware for /static without path rewrite but with Host header adjustment
	e.Group("/static").Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			proxy := httputil.NewSingleHostReverseProxy(eqsageTarget)
			proxy.Director = func(req *http.Request) {
				req.URL.Scheme = eqsageTarget.Scheme
				req.URL.Host = eqsageTarget.Host
				// No path rewrite needed for /static, but we ensure the Host header is set
				req.Host = eqsageTarget.Host
			}
			proxy.ServeHTTP(c.Response(), c.Request())
			return nil
		}
	})

	c.watcher.Run()

	// serve spa as embedded static assets
	s := spa.NewSpa(c.logger)
	e.GET("/*", s.Spa().Handler(), middleware.GzipWithConfig(middleware.GzipConfig{Level: 1}))
	e.Use(s.Spa().MiddlewareHandler())

	e.HTTPErrorHandler = errorHandler
	e.HideBanner = true
	e.HidePort = true

	imgcat.Render(banner.GetLogo())

	go func() {
		c.launcher.ServerProcessLauncherWatchdog()
	}()

	c.logger.Info().Any("port", port).Msgf("Starting Spire HTTP server")

	return e.Start(fmt.Sprintf(":%v", port))
}

func first(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func errorHandler(err error, c echo.Context) {
	_ = c.JSON(http.StatusUnprocessableEntity, echo.Map{"error": first(err.Error())})
}

func BootstrapMiddleware(e *echo.Echo, router *routes.Router) {
	for _, m := range router.GlobalPreMiddlewares() {
		e.Pre(m)
	}
	for _, m := range router.GlobalMiddlewares() {
		e.Use(m)
	}
}

func BootstrapControllers(e *echo.Echo, controllerGroups ...*routes.ControllerGroup) error {
	for _, cg := range controllerGroups {
		g := e.Group(cg.RoutePrefix(), cg.Middlewares()...)
		for _, c := range cg.Controllers() {
			if err := registerRoutes(c, g); err != nil {
				return err
			}
		}
	}

	return nil
}

func registerRoutes(controller routes.Controller, g *echo.Group) error {
	for _, r := range controller.Routes() {
		switch r.Method() {
		case "ANY":
			g.Any(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodGet:
			g.GET(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodPost:
			g.POST(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodPut:
			g.PUT(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodPatch:
			g.PATCH(r.Route(), r.Handler(), r.Middlewares()...)
		case http.MethodDelete:
			g.DELETE(r.Route(), r.Handler(), r.Middlewares()...)
		default:
			return errors.New("invalid r method specified")
		}
	}

	return nil
}
