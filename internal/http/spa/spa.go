package spa

import (
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/gobuffalo/packr"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Packaged SPA is a single page application that is compiled into the resulting
// go application binary and ultimately served using echo in our monorepo framework
//
// Packaged bundler (https://github.com/gobuffalo/packr)
type PackedSpaService struct {
	fs      http.Handler
	handler echo.HandlerFunc
	box     packr.Box
	logger *logrus.Logger
	config PackedSpaServeConfig
}

func (s PackedSpaService) Handler() echo.HandlerFunc {
	return s.handler
}

// The configuration required to serve the packaged SPA app
type PackedSpaServeConfig struct {
	BasePath      string   // Root path from where the SPA is served
	LocalBasePath string   // The path that the SPA is located from the relative context of this provider
	SpaIndex      string   // SPA index - where requests will fallback to
	SkipPaths     []string // Route prefixes to skip entirely
}

// New up an instance of packaged spa service
func NewPackedSpaService(logger *logrus.Logger, config PackedSpaServeConfig) *PackedSpaService {
	s := &PackedSpaService{}
	s.config = config
	s.logger = logger
	s.box = packr.NewBox(s.config.LocalBasePath)
	s.fs = http.FileServer(s.box)
	//s.handler = echo.WrapHandler(http.StripPrefix(s.config.BasePath, s.fs))
	s.handler = WrapCachedHandler(http.StripPrefix(s.config.BasePath, s.fs))

	return s
}

func WrapCachedHandler(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		if contains([]string{".js", ".css", ".png", ".woff", ".ttf", ".jpg", ".gif", ".svg", ".ico"}, c.Request().RequestURI) {
			c.Response().Header().Set("Vary", "Accept-Encoding")
			c.Response().Header().Set("Cache-Control", "public, max-age=7776000")
		}
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func contains(slice []string, val string) (bool) {
	for _, item := range slice {
		if strings.Contains(val, item) {
			return true
		}
	}
	return false
}

// This middleware handler is for the most part a static middleware handler for handling static assets
// The main difference between this middleware and a generic static middleware is that it provides assets that live
// within a static asset compiler that gets bundled within the compiled resulting binary /packed-spa/index.html
func (s PackedSpaService) MiddlewareHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// if we explicitly ignore paths, process routing as normal
			for _, skipPath := range s.config.SkipPaths {
				if strings.Contains(c.Request().URL.Path, skipPath) {
					return next(c)
				}
			}

			p := c.Request().URL.Path
			if strings.HasSuffix(c.Path(), "*") { // When serving from a group, e.g. `/static*`.
				p = c.Param("*")
			}
			p, err = url.PathUnescape(p)
			if err != nil {
				return
			}

			name := filepath.Join(s.config.BasePath, path.Clean("/"+p)) // "/"+ for security

			if name == "/" || name == "\\" {
				index, err := s.box.Find(s.config.SpaIndex)
				if err != nil {
					s.logger.Error(err)
				}
				return c.HTML(http.StatusOK, string(index))
			}

			_, err = os.Stat(name)
			if err != nil {
				if os.IsNotExist(err) {

					// If we find a valid non-index file in the box, continue the request as normal
					// and let the static asset handler pick up the request later
					fileRequest := strings.Replace(c.Request().RequestURI, s.config.BasePath, "", -1)
					_, err := s.box.Find(fileRequest)
					if err == nil {
						return next(c)
					}

					// If we didn't find a non-index asset at this point, we need to return the
					// spa index when nested SPA route requests are made
					// eg: /spa/nested/route
					index, err := s.box.Find(s.config.SpaIndex)
					if err != nil {
						s.logger.Error(err)
					}

					return c.HTML(http.StatusOK, string(index))
				}
				return
			}

			return c.File(name)
		}
	}
}

// Packaged Spa Controller is a provider to easily handle registration of the necessary route handlers
type PackagedSpaController struct {
	service *PackedSpaService
}

func NewSpaController(s *PackedSpaService) *PackagedSpaController {
	return &PackagedSpaController{
		service: s,
	}
}

// Routes registers controller specific routes
func (c *PackagedSpaController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "/*", c.service.Handler(), nil),
	}
}

func (s PackedSpaService) Controller() *PackagedSpaController {
	return NewSpaController(&s)
}
