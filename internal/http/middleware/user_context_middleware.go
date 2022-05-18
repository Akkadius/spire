package middleware

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserContextMiddleware struct {
	db     *database.DatabaseResolver
	cache  *gocache.Cache
	logger *logrus.Logger
}

func NewUserContextMiddleware(
	db *database.DatabaseResolver,
	cache *gocache.Cache,
	logger *logrus.Logger) *UserContextMiddleware {
	return &UserContextMiddleware{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

type jwtCustomClaims struct {
	Authorized bool   `json:"authorized"`
	UserId     string `json:"userId"`
	jwt.StandardClaims
}

func (m UserContextMiddleware) HandleHeader() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: m.successHandler,
		Skipper:        m.skipperHeader,
		SigningKey:     []byte(os.Getenv("JWT_SECRET_KEY")),
		Claims:         &jwtCustomClaims{},
	})
}

func (m UserContextMiddleware) HandleQuerystring() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: m.successHandler,
		TokenLookup:    "query:jwt",
		Skipper:        m.skipperQuery,
		SigningKey:     []byte(os.Getenv("JWT_SECRET_KEY")),
		Claims:         &jwtCustomClaims{},
	})
}

func (m UserContextMiddleware) skipperHeader(c echo.Context) bool {
	// Skip JWT signature validation if authenticated with basic auth
	if strings.Contains(c.Request().Header.Get("Authorization"), "Basic") {
		return true
	}

	// If we send an authorization header then we will go through our JWT signature validation logic
	if c.Request().Header.Get("Authorization") != "" {
		return false
	}

	// When we don't send an authorization header, we assume that the request
	// is not an authenticated user we will simply use the local default
	// PEQ database instance

	return true
}

func (m UserContextMiddleware) skipperQuery(c echo.Context) bool {
	// If we send a query param header then we will go through our JWT signature validation logic
	if c.QueryParam("jwt") != "" {
		return false
	}

	// When we don't send an authorization header, we assume that the request
	// is not an authenticated user we will simply use the local default
	// PEQ database instance

	return true
}

func (m UserContextMiddleware) successHandler(c echo.Context) {
	// signature validation
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	userId, err := strconv.Atoi(claims.UserId)
	if err != nil {
		panic("Failed to convert JWT UserID")
	}

	m.logger.Debugln(fmt.Sprintf("JWT validation success as user [%v]", userId))

	// cache user context for 10 minutes to refrain from repeated DB hits
	userKey := fmt.Sprintf("user-%v", userId)
	cachedUser, found := m.cache.Get(userKey)
	if found {
		c.Set("user", cachedUser)
		return
	}

	// fetch user if not in cache and set in request context
	var dbUser models.User
	m.db.GetSpireDb().Find(&dbUser, userId)
	if dbUser.ID > 0 {
		c.Set("user", dbUser)
		m.cache.Set(userKey, dbUser, 10*time.Minute)
	}
}
