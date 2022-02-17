package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type AccountController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAccountController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AccountController {
	return &AccountController{
		db:	 db,
		logger: logger,
	}
}

func (e *AccountController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "account/:id", e.getAccount, nil),
		routes.RegisterRoute(http.MethodGet, "accounts", e.listAccounts, nil),
		routes.RegisterRoute(http.MethodPut, "account", e.createAccount, nil),
		routes.RegisterRoute(http.MethodDelete, "account/:id", e.deleteAccount, nil),
		routes.RegisterRoute(http.MethodPatch, "account/:id", e.updateAccount, nil),
		routes.RegisterRoute(http.MethodPost, "accounts/bulk", e.getAccountsBulk, nil),
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
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
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
// @Param id path int true "id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AccountFlags<br>AccountIps<br>AccountRewards<br>BugReports<br>Sharedbanks"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Account
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /account/{id} [get]
func (e *AccountController) getAccount(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ID]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.Account
	query := e.db.QueryContext(models.Account{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
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
// @Param ID path int true "ID"
// @Param account body models.Account true "Account"
// @Success 200 {array} models.Account
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /account/{id} [patch]
func (e *AccountController) updateAccount(c echo.Context) error {
	request := new(models.Account)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ID]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.Account
	query := e.db.QueryContext(models.Account{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
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
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Account{}, c).Model(&models.Account{}).Create(&account).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, account)
}

// deleteAccount godoc
// @Id deleteAccount
// @Summary Deletes Account
// @Accept json
// @Produce json
// @Tags Account
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /account/{id} [delete]
func (e *AccountController) deleteAccount(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.Account
	query := e.db.QueryContext(models.Account{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.Account{}, c).Model(&models.Account{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getAccountsBulk godoc
// @Id getAccountsBulk
// @Summary Gets Accounts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Account
// @Success 200 {array} models.Account
// @Failure 500 {string} string "Bad query request"
// @Router /accounts/bulk [post]
func (e *AccountController) getAccountsBulk(c echo.Context) error {
	var results []models.Account

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.Account{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
