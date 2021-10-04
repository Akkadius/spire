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

func (m UserContextMiddleware) Handle() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SuccessHandler: func(c echo.Context) {

			// signature validation
			m.logger.Debugln("JWT validation success")
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*jwtCustomClaims)
			userId, err := strconv.Atoi(claims.UserId)
			if err != nil {
				panic("Failed to convert JWT UserID")
			}

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
		},
		Skipper: func(c echo.Context) bool {

			// If we send an authorization header then we will go through our JWT
			// signature validation logic
			if c.Request().Header.Get("Authorization") != "" {
				return false
			}

			// When we don't send an authorization header, we assume that the request
			// is not an authenticated user so we will simply use the local default
			// PEQ database instance

			return true
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		Claims:     &jwtCustomClaims{},
	})
}
