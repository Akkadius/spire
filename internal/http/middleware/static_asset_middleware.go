package middleware

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	// StaticConfig defines the config for Static middleware.
	StaticConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Root directory from where the static content is served.
		// Required.
		Root string `yaml:"root"`

		// Root directory from where the static content is served.
		// Required.
		StripPrefix string `yaml:"strip_prefix"`

		// Index file for serving a directory.
		// Optional. Default value "index.html".
		Index string `yaml:"index"`

		// Enable directory browsing.
		// Optional. Default value false.
		Browse bool `yaml:"browse"`

		// Enable ignoring of the base of the URL path.
		// Example: when assigning a static middleware to a non root path group,
		// the filesystem path is not doubled
		// Optional. Default value false.
		IgnoreBase bool `yaml:"ignoreBase"`

		// Filesystem provides access to the static content.
		// Optional. Defaults to http.Dir(config.Root)
		Filesystem http.FileSystem `yaml:"-"`
	}
)

var (
	// DefaultStaticConfig is the default Static middleware config.
	DefaultStaticConfig = StaticConfig{
		Skipper: middleware.DefaultSkipper,
		Index:   "index.html",
	}
)

// StaticAsset returns a Static middleware with config.
// See `Static()`.
func StaticAsset(config StaticConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Root == "" {
		config.Root = "." // For security we want to restrict to CWD.
	}
	if config.Skipper == nil {
		config.Skipper = DefaultStaticConfig.Skipper
	}
	if config.Index == "" {
		config.Index = DefaultStaticConfig.Index
	}
	if config.Filesystem == nil {
		config.Filesystem = http.Dir(config.Root)
		config.Root = "."
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			p := c.Request().URL.Path

			if strings.Contains(p, "/swagger/") {
				return next(c)
			}

			if strings.HasSuffix(c.Path(), "*") { // When serving from a group, e.g. `/static*`.
				p = c.Param("*")
			}
			p, err = url.PathUnescape(p)
			if err != nil {
				return next(c)
			}
			name := filepath.Join(config.Root, filepath.Clean("/"+p)) // "/"+ for security

			if len(config.StripPrefix) > 0 {
				name = strings.ReplaceAll(name, config.StripPrefix, "")
			}

			file, err := openFile(config.Filesystem, name)
			if err != nil {
				return next(c)
			}

			defer file.Close()

			info, err := file.Stat()
			if err != nil {
				return next(c)
			}

			if info.IsDir() {
				return next(c)
			}

			return serveFile(c, file, info)
		}
	}
}

func openFile(fs http.FileSystem, name string) (http.File, error) {
	pathWithSlashes := filepath.ToSlash(name)
	return fs.Open(pathWithSlashes)
}

func serveFile(c echo.Context, file http.File, info os.FileInfo) error {
	c.Response().Header().Set("Vary", "Accept-Encoding")
	c.Response().Header().Set("Cache-Control", "public, max-age=7776000")

	http.ServeContent(c.Response(), c.Request(), info.Name(), info.ModTime(), file)
	return nil
}
