package spire

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UsersController struct {
	db        *database.DatabaseResolver
	logger    *logrus.Logger
	spireuser *UserService
	crypt     *encryption.Encrypter
}

func NewUsersController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	spireuser *UserService,
	crypt *encryption.Encrypter,
) *UsersController {
	return &UsersController{
		db:        db,
		logger:    logger,
		spireuser: spireuser,
		crypt:     crypt,
	}
}

func (a *UsersController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "users", a.list, nil),
		routes.RegisterRoute(http.MethodPut, "user", a.create, nil),
		routes.RegisterRoute(http.MethodDelete, "user/:id", a.delete, nil),
		routes.RegisterRoute(http.MethodPost, "user/:id/password-reset", a.passwordReset, nil),
	}
}

func (a *UsersController) list(c echo.Context) error {
	var results []models.User
	if a.db.GetSpireDb() != nil {
		query := a.db.GetSpireDb().Model(&models.User{})

		// param
		search := c.QueryParam("q")
		if len(search) > 0 {
			query.Or("user_name LIKE ?", fmt.Sprintf("%%%v%%", search))
		}

		query.Limit(100)

		err := query.Find(&results).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
		}

		return c.JSON(http.StatusOK, results)
	}

	return c.JSON(http.StatusOK, echo.Map{"error": "No users found"})
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// create endpoint is currently primarily in use for local spire users
// assumptions in logic are based around its use
func (a *UsersController) create(c echo.Context) error {

	// validation: not admin
	u := request.GetUser(c)
	if !u.IsAdmin {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "You are not an administrator"})
	}

	r := new(UserCreateRequest)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// new user
	user := models.User{
		UserName: r.Username,
		FullName: r.Username,
		Password: r.Password,
		Provider: LoginProviderLocal,
		IsAdmin:  false,
	}

	newUser, err := a.spireuser.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// associate user with connection
	var uc models.UserServerDatabaseConnection
	uc.UserId = newUser.ID
	uc.Active = 1
	uc.ServerDatabaseConnectionId = 1
	uc.CreatedBy = u.ID
	err = a.db.GetSpireDb().Create(&uc).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, "User created successfully!")
}

// delete endpoint is currently primarily in use for local spire users
// assumptions in logic are based around its use
func (a *UsersController) delete(c echo.Context) error {

	// validation: not admin
	u := request.GetUser(c)
	if !u.IsAdmin {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "You are not an administrator"})
	}

	userId := c.Param("id")

	// validation: can't delete admin
	var checkUser models.User
	err := a.db.GetSpireDb().Where("id = ?", userId).First(&checkUser).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if checkUser.IsAdmin {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Cannot delete an admin!"})
	}

	// validation: can't delete self
	if userId == fmt.Sprintf("%v", u.ID) {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Cannot delete self!"})
	}

	// delete user
	err = a.db.GetSpireDb().Where("id = ?", userId).Delete(&models.User{}).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// hard-coded local database injected connection
	connectionId := 1

	// Delete
	err = a.db.GetSpireDb().
		Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).
		Delete(&models.UserServerDatabaseConnection{}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	// Delete entries first
	err = a.db.GetSpireDb().
		Where("user_id = ? and server_database_connection_id = ?", userId, connectionId).
		Delete(&models.UserServerResourcePermission{}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, "User deleted successfully!")
}

type PasswordResetRequest struct {
	Password string `json:"password"`
}

func (a *UsersController) passwordReset(c echo.Context) error {
	// validation: not admin
	u := request.GetUser(c)
	if !u.IsAdmin {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "You are not an administrator"})
	}

	r := new(PasswordResetRequest)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userId := c.Param("id")

	var user models.User
	err := a.db.GetSpireDb().Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	hash, err := a.crypt.GeneratePassword(r.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user.Password = hash

	a.db.GetSpireDb().Save(&user)

	return c.JSON(http.StatusOK, "User password reset successfully!")
}
