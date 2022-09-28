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

type AccountFlagController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAccountFlagController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AccountFlagController {
	return &AccountFlagController{
		db:	 db,
		logger: logger,
	}
}

func (e *AccountFlagController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "account_flag/:pAccid", e.getAccountFlag, nil),
		routes.RegisterRoute(http.MethodGet, "account_flags", e.listAccountFlags, nil),
		routes.RegisterRoute(http.MethodPut, "account_flag", e.createAccountFlag, nil),
		routes.RegisterRoute(http.MethodDelete, "account_flag/:pAccid", e.deleteAccountFlag, nil),
		routes.RegisterRoute(http.MethodPatch, "account_flag/:pAccid", e.updateAccountFlag, nil),
		routes.RegisterRoute(http.MethodPost, "account_flags/bulk", e.getAccountFlagsBulk, nil),
	}
}

// listAccountFlags godoc
// @Id listAccountFlags
// @Summary Lists AccountFlags
// @Accept json
// @Produce json
// @Tags AccountFlag
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AccountFlag
// @Failure 500 {string} string "Bad query request"
// @Router /account_flags [get]
func (e *AccountFlagController) listAccountFlags(c echo.Context) error {
	var results []models.AccountFlag
	err := e.db.QueryContext(models.AccountFlag{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAccountFlag godoc
// @Id getAccountFlag
// @Summary Gets AccountFlag
// @Accept json
// @Produce json
// @Tags AccountFlag
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AccountFlag
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /account_flag/{id} [get]
func (e *AccountFlagController) getAccountFlag(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	pAccid, err := strconv.Atoi(c.Param("pAccid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PAccid]"})
	}
	params = append(params, pAccid)
	keys = append(keys, "p_accid = ?")

	// key param [p_flag] position [2] type [varchar]
	if len(c.QueryParam("p_flag")) > 0 {
		pFlagParam, err := strconv.Atoi(c.QueryParam("p_flag"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [p_flag] err [%s]", err.Error())})
		}

		params = append(params, pFlagParam)
		keys = append(keys, "p_flag = ?")
	}

	// query builder
	var result models.AccountFlag
	query := e.db.QueryContext(models.AccountFlag{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.PAccid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAccountFlag godoc
// @Id updateAccountFlag
// @Summary Updates AccountFlag
// @Accept json
// @Produce json
// @Tags AccountFlag
// @Param id path int true "Id"
// @Param account_flag body models.AccountFlag true "AccountFlag"
// @Success 200 {array} models.AccountFlag
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /account_flag/{id} [patch]
func (e *AccountFlagController) updateAccountFlag(c echo.Context) error {
	request := new(models.AccountFlag)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	pAccid, err := strconv.Atoi(c.Param("pAccid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PAccid]"})
	}
	params = append(params, pAccid)
	keys = append(keys, "p_accid = ?")

	// key param [p_flag] position [2] type [varchar]
	if len(c.QueryParam("p_flag")) > 0 {
		pFlagParam, err := strconv.Atoi(c.QueryParam("p_flag"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [p_flag] err [%s]", err.Error())})
		}

		params = append(params, pFlagParam)
		keys = append(keys, "p_flag = ?")
	}

	// query builder
	var result models.AccountFlag
	query := e.db.QueryContext(models.AccountFlag{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AccountFlag{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAccountFlag godoc
// @Id createAccountFlag
// @Summary Creates AccountFlag
// @Accept json
// @Produce json
// @Param account_flag body models.AccountFlag true "AccountFlag"
// @Tags AccountFlag
// @Success 200 {array} models.AccountFlag
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /account_flag [put]
func (e *AccountFlagController) createAccountFlag(c echo.Context) error {
	accountFlag := new(models.AccountFlag)
	if err := c.Bind(accountFlag); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AccountFlag{}, c).Model(&models.AccountFlag{}).Create(&accountFlag).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, accountFlag)
}

// deleteAccountFlag godoc
// @Id deleteAccountFlag
// @Summary Deletes AccountFlag
// @Accept json
// @Produce json
// @Tags AccountFlag
// @Param id path int true "pAccid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /account_flag/{id} [delete]
func (e *AccountFlagController) deleteAccountFlag(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	pAccid, err := strconv.Atoi(c.Param("pAccid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, pAccid)
	keys = append(keys, "p_accid = ?")

	// key param [p_flag] position [2] type [varchar]
	if len(c.QueryParam("p_flag")) > 0 {
		pFlagParam, err := strconv.Atoi(c.QueryParam("p_flag"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [p_flag] err [%s]", err.Error())})
		}

		params = append(params, pFlagParam)
		keys = append(keys, "p_flag = ?")
	}

	// query builder
	var result models.AccountFlag
	query := e.db.QueryContext(models.AccountFlag{}, c)
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

// getAccountFlagsBulk godoc
// @Id getAccountFlagsBulk
// @Summary Gets AccountFlags in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AccountFlag
// @Success 200 {array} models.AccountFlag
// @Failure 500 {string} string "Bad query request"
// @Router /account_flags/bulk [post]
func (e *AccountFlagController) getAccountFlagsBulk(c echo.Context) error {
	var results []models.AccountFlag

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

	err := e.db.QueryContext(models.AccountFlag{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
