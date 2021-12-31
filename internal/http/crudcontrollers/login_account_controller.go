package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type LoginAccountController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoginAccountController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoginAccountController {
	return &LoginAccountController{
		db:     db,
		logger: logger,
	}
}

func (e *LoginAccountController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "login_account/:login_account", e.deleteLoginAccount, nil),
		routes.RegisterRoute(http.MethodGet, "login_account/:login_account", e.getLoginAccount, nil),
		routes.RegisterRoute(http.MethodGet, "login_accounts", e.listLoginAccounts, nil),
		routes.RegisterRoute(http.MethodPost, "login_accounts/bulk", e.getLoginAccountsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "login_account/:login_account", e.updateLoginAccount, nil),
		routes.RegisterRoute(http.MethodPut, "login_account", e.createLoginAccount, nil),
	}
}

// listLoginAccounts godoc
// @Id listLoginAccounts
// @Summary Lists LoginAccounts
// @Accept json
// @Produce json
// @Tags LoginAccount
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginAccount
// @Failure 500 {string} string "Bad query request"
// @Router /login_accounts [get]
func (e *LoginAccountController) listLoginAccounts(c echo.Context) error {
	var results []models.LoginAccount
	err := e.db.QueryContext(models.LoginAccount{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoginAccount godoc
// @Id getLoginAccount
// @Summary Gets LoginAccount
// @Accept json
// @Produce json
// @Tags LoginAccount
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginAccount
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /login_account/{id} [get]
func (e *LoginAccountController) getLoginAccount(c echo.Context) error {
	loginAccountId, err := strconv.Atoi(c.Param("login_account"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.LoginAccount
	err = e.db.QueryContext(models.LoginAccount{}, c).First(&result, loginAccountId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLoginAccount godoc
// @Id updateLoginAccount
// @Summary Updates LoginAccount
// @Accept json
// @Produce json
// @Tags LoginAccount
// @Param id path int true "Id"
// @Param login_account body models.LoginAccount true "LoginAccount"
// @Success 200 {array} models.LoginAccount
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /login_account/{id} [patch]
func (e *LoginAccountController) updateLoginAccount(c echo.Context) error {
	loginAccount := new(models.LoginAccount)
	if err := c.Bind(loginAccount); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.LoginAccount{}
	err := e.db.Get(models.LoginAccount{}, c).Model(&models.LoginAccount{}).First(&entity, loginAccount.ID).Error
	if err != nil || loginAccount.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LoginAccount{}, c).Model(&entity).Select("*").Updates(&loginAccount).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, loginAccount)
}

// createLoginAccount godoc
// @Id createLoginAccount
// @Summary Creates LoginAccount
// @Accept json
// @Produce json
// @Param login_account body models.LoginAccount true "LoginAccount"
// @Tags LoginAccount
// @Success 200 {array} models.LoginAccount
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /login_account [put]
func (e *LoginAccountController) createLoginAccount(c echo.Context) error {
	loginAccount := new(models.LoginAccount)
	if err := c.Bind(loginAccount); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.LoginAccount{}, c).Model(&models.LoginAccount{}).Create(&loginAccount).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, loginAccount)
}

// deleteLoginAccount godoc
// @Id deleteLoginAccount
// @Summary Deletes LoginAccount
// @Accept json
// @Produce json
// @Tags LoginAccount
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /login_account/{id} [delete]
func (e *LoginAccountController) deleteLoginAccount(c echo.Context) error {
	loginAccountId, err := strconv.Atoi(c.Param("login_account"))
	if err != nil {
		e.logger.Error(err)
	}

	loginAccount := new(models.LoginAccount)
	err = e.db.Get(models.LoginAccount{}, c).Model(&models.LoginAccount{}).First(&loginAccount, loginAccountId).Error
	if err != nil || loginAccount.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LoginAccount{}, c).Model(&models.LoginAccount{}).Delete(&loginAccount).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLoginAccountsBulk godoc
// @Id getLoginAccountsBulk
// @Summary Gets LoginAccounts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoginAccount
// @Success 200 {array} models.LoginAccount
// @Failure 500 {string} string "Bad query request"
// @Router /login_accounts/bulk [post]
func (e *LoginAccountController) getLoginAccountsBulk(c echo.Context) error {
	var results []models.LoginAccount

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.LoginAccount{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
