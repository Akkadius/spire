package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type LoginAccountController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoginAccountController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoginAccountController {
	return &LoginAccountController{
		db:	 db,
		logger: logger,
	}
}

func (e *LoginAccountController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "login_account/:id", e.getLoginAccount, nil),
		routes.RegisterRoute(http.MethodGet, "login_accounts", e.listLoginAccounts, nil),
		routes.RegisterRoute(http.MethodPut, "login_account", e.createLoginAccount, nil),
		routes.RegisterRoute(http.MethodDelete, "login_account/:id", e.deleteLoginAccount, nil),
		routes.RegisterRoute(http.MethodPatch, "login_account/:id", e.updateLoginAccount, nil),
		routes.RegisterRoute(http.MethodPost, "login_accounts/bulk", e.getLoginAccountsBulk, nil),
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
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
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
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.LoginAccount
	query := e.db.QueryContext(models.LoginAccount{}, c)
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
	request := new(models.LoginAccount)
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.LoginAccount
	query := e.db.QueryContext(models.LoginAccount{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.LoginAccount{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LoginAccount{}, c).Model(&models.LoginAccount{}).Create(&loginAccount).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
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
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /login_account/{id} [delete]
func (e *LoginAccountController) deleteLoginAccount(c echo.Context) error {
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
	var result models.LoginAccount
	query := e.db.QueryContext(models.LoginAccount{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = query.Limit(10000).Delete(&result).Error
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
