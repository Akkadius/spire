package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type AccountController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAccountController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AccountController {
	return &AccountController {
		db:     db,
		logger: logger,
	}
}

func (e *AccountController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "account/:account", e.deleteAccount, nil),
		routes.RegisterRoute(http.MethodGet, "account/:account", e.getAccount, nil),
		routes.RegisterRoute(http.MethodGet, "accounts", e.listAccounts, nil),
		routes.RegisterRoute(http.MethodPatch, "account/:account", e.updateAccount, nil),
		routes.RegisterRoute(http.MethodPut, "account", e.createAccount, nil),
	}
}

// listAccounts godoc
// @Id listAccounts
// @Summary Lists Accounts
// @Accept json
// @Produce json
// @Tags Account
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AccountFlags<br>AccountIps<br>AccountRewards<br>BugReports<br>Sharedbanks"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Account
// @Failure 500 {string} string "Bad query request"
// @Router /accounts [get]
func (e *AccountController) listAccounts(c echo.Context) error {
	var results []models.Account
	err := e.db.QueryContext(models.Account{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAccount godoc
// @Id getAccount
// @Summary Gets Account
// @Accept json
// @Produce json
// @Tags Account
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AccountFlags<br>AccountIps<br>AccountRewards<br>BugReports<br>Sharedbanks"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Account
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /account/{id} [get]
func (e *AccountController) getAccount(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("account"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Account
	err = e.db.QueryContext(models.Account{}, c).First(&result, accountId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAccount godoc
// @Id updateAccount
// @Summary Updates Account
// @Accept json
// @Produce json
// @Tags Account
// @Param id path int true "Id"
// @Param account body models.Account true "Account"
// @Success 200 {array} models.Account
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /account/{id} [patch]
func (e *AccountController) updateAccount(c echo.Context) error {
	account := new(models.Account)
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Account{}, c).Model(&models.Account{}).First(&models.Account{}, account.ID).Error
	if err != nil || account.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Account{}, c).Model(&models.Account{}).Updates(&account).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, account)
}

// createAccount godoc
// @Id createAccount
// @Summary Creates Account
// @Accept json
// @Produce json
// @Param account body models.Account true "Account"
// @Tags Account
// @Success 200 {array} models.Account
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /account [put]
func (e *AccountController) createAccount(c echo.Context) error {
	account := new(models.Account)
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Account{}, c).Model(&models.Account{}).Create(&account).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, account)
}

// deleteAccount godoc
// @Id deleteAccount
// @Summary Deletes Account
// @Accept json
// @Produce json
// @Tags Account
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /account/{id} [delete]
func (e *AccountController) deleteAccount(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("account"))
	if err != nil {
		e.logger.Error(err)
	}

	account := new(models.Account)
	err = e.db.Get(models.Account{}, c).Model(&models.Account{}).First(&account, accountId).Error
	if err != nil || account.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Account{}, c).Model(&models.Account{}).Delete(&account).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
