package controllers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/danilopolani/gocialite"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type AuthController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
	gocial *gocialite.Dispatcher
	cache  *gocache.Cache
	user   *spire.UserService
}

func NewAuthController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	user *spire.UserService,
	cache *gocache.Cache,
) *AuthController {
	return &AuthController{
		db:     db,
		logger: logger,
		cache:  cache,
		gocial: gocialite.NewDispatcher(),
		user:   user,
	}
}

func (a *AuthController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "login", a.login, nil),
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

	github, _, err := a.gocial.Handle(state, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	//spew.Dump(github)
	//spew.Dump(token)

	var user models.User
	a.db.GetSpireDb().Where("user_name = ? and provider = ?", github.Username, spire.LoginProviderGithub).First(&user)

	if user.ID == 0 {
		u := models.User{
			UserName:  github.Username,
			FullName:  github.FullName,
			FirstName: github.FirstName,
			LastName:  github.LastName,
			Email:     github.Email,
			Avatar:    github.Avatar,
			Provider:  spire.LoginProviderGithub,
		}

		// new github
		newUser, err := a.user.CreateUser(u)
		if err != nil {
			a.logger.Error(err)
		}

		if newUser.ID > 0 {
			user.ID = newUser.ID
		}
	}

	newToken, _ := createJwtToken(fmt.Sprintf("%v", user.ID))
	callbackUrl := fmt.Sprintf("%s/fe-auth-callback?jwt=%s", os.Getenv("VUE_APP_FRONTEND_BASE_URL"), newToken)

	// clear users cache on logging in if exists
	userKey := fmt.Sprintf("user-%v", user.ID)
	a.cache.Delete(userKey)

	fmt.Println(callbackUrl)

	return c.Redirect(http.StatusMovedPermanently, callbackUrl)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthController) login(c echo.Context) error {
	u := new(LoginRequest)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	passwordValid, err, user := a.user.CheckUserLogin(u.Username, u.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if !passwordValid {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Password invalid"})
	}

	newToken, _ := createJwtToken(fmt.Sprintf("%v", user.ID))

	// clear users cache on logging in if exists
	userKey := fmt.Sprintf("user-%v", user.ID)
	a.cache.Delete(userKey)

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": echo.Map{
				"message": "Login success",
				"token":   newToken,
			},
		},
	)
}
