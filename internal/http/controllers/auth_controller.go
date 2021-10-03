package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/danilopolani/gocialite"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type AuthController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
	gocial *gocialite.Dispatcher
}

func NewAuthController(db *database.DatabaseResolver, logger *logrus.Logger) *AuthController {
	return &AuthController{
		db:     db,
		logger: logger,
		gocial: gocialite.NewDispatcher(),
	}
}

func (a *AuthController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "github", a.githubRedirectHandler, nil),
		routes.RegisterRoute(http.MethodGet, "github/callback", a.githubCallbackHandler, nil),
	}
}

func createJwtToken(userId string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = userId

	atClaims["exp"] = time.Now().Add(time.Hour * 87600).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthController) githubRedirectHandler(c echo.Context) error {
	authURL, err := a.gocial.New().
		Driver("github").
		Scopes([]string{""}).
		Redirect(
			os.Getenv("GITHUB_CLIENT_ID"),
			os.Getenv("GITHUB_CLIENT_SECRET"),
			os.Getenv("GITHUB_REDIRECT_CALLBACK"),
		)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(authURL)

	return c.Redirect(http.StatusFound, authURL) // Redirect with 302 HTTP code
}

type JwtTokenResponse struct {
	Id          int    `json:"id"`
	UserName    string `json:"name"`
	GithubToken string `json:"github_token"`
	Jwt         string `json:"jwt"`
}

func (a *AuthController) githubCallbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	user, token, err := a.gocial.Handle(state, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	//spew.Dump(user)
	spew.Dump(token)

	var newUser models.User
	a.db.GetSpireDb().FirstOrCreate(
		&newUser, models.User{
			UserName:  user.Username,
			FullName:  user.FullName,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Provider:  "github",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	)

	newToken, _ := createJwtToken(fmt.Sprintf("%v", newUser.ID))
	callbackUrl := fmt.Sprintf("%s/fe-auth-callback?jwt=%s", os.Getenv("VUE_APP_FRONTEND_BASE_URL"), newToken)

	fmt.Println(callbackUrl)

	return c.Redirect(http.StatusMovedPermanently, callbackUrl)
}
