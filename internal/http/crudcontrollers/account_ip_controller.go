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

type AccountIpController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAccountIpController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AccountIpController {
	return &AccountIpController{
		db:	 db,
		logger: logger,
	}
}

func (e *AccountIpController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "account_ip/:accid", e.getAccountIp, nil),
		routes.RegisterRoute(http.MethodGet, "account_ips", e.listAccountIps, nil),
		routes.RegisterRoute(http.MethodPut, "account_ip", e.createAccountIp, nil),
		routes.RegisterRoute(http.MethodDelete, "account_ip/:accid", e.deleteAccountIp, nil),
		routes.RegisterRoute(http.MethodPatch, "account_ip/:accid", e.updateAccountIp, nil),
		routes.RegisterRoute(http.MethodPost, "account_ips/bulk", e.getAccountIpsBulk, nil),
	}
}

// listAccountIps godoc
// @Id listAccountIps
// @Summary Lists AccountIps
// @Accept json
// @Produce json
// @Tags AccountIp
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AccountIp
// @Failure 500 {string} string "Bad query request"
// @Router /account_ips [get]
func (e *AccountIpController) listAccountIps(c echo.Context) error {
	var results []models.AccountIp
	err := e.db.QueryContext(models.AccountIp{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAccountIp godoc
// @Id getAccountIp
// @Summary Gets AccountIp
// @Accept json
// @Produce json
// @Tags AccountIp
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AccountIp
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /account_ip/{id} [get]
func (e *AccountIpController) getAccountIp(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accid, err := strconv.Atoi(c.Param("accid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Accid]"})
	}
	params = append(params, accid)
	keys = append(keys, "accid = ?")

	// key param [ip] position [2] type [varchar]
	if len(c.QueryParam("ip")) > 0 {
		ipParam, err := strconv.Atoi(c.QueryParam("ip"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ip] err [%s]", err.Error())})
		}

		params = append(params, ipParam)
		keys = append(keys, "ip = ?")
	}

	// query builder
	var result models.AccountIp
	query := e.db.QueryContext(models.AccountIp{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Accid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAccountIp godoc
// @Id updateAccountIp
// @Summary Updates AccountIp
// @Accept json
// @Produce json
// @Tags AccountIp
// @Param id path int true "Id"
// @Param account_ip body models.AccountIp true "AccountIp"
// @Success 200 {array} models.AccountIp
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /account_ip/{id} [patch]
func (e *AccountIpController) updateAccountIp(c echo.Context) error {
	request := new(models.AccountIp)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	accid, err := strconv.Atoi(c.Param("accid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Accid]"})
	}
	params = append(params, accid)
	keys = append(keys, "accid = ?")

	// key param [ip] position [2] type [varchar]
	if len(c.QueryParam("ip")) > 0 {
		ipParam, err := strconv.Atoi(c.QueryParam("ip"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ip] err [%s]", err.Error())})
		}

		params = append(params, ipParam)
		keys = append(keys, "ip = ?")
	}

	// query builder
	var result models.AccountIp
	query := e.db.QueryContext(models.AccountIp{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AccountIp{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAccountIp godoc
// @Id createAccountIp
// @Summary Creates AccountIp
// @Accept json
// @Produce json
// @Param account_ip body models.AccountIp true "AccountIp"
// @Tags AccountIp
// @Success 200 {array} models.AccountIp
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /account_ip [put]
func (e *AccountIpController) createAccountIp(c echo.Context) error {
	accountIp := new(models.AccountIp)
	if err := c.Bind(accountIp); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AccountIp{}, c).Model(&models.AccountIp{}).Create(&accountIp).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, accountIp)
}

// deleteAccountIp godoc
// @Id deleteAccountIp
// @Summary Deletes AccountIp
// @Accept json
// @Produce json
// @Tags AccountIp
// @Param id path int true "accid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /account_ip/{id} [delete]
func (e *AccountIpController) deleteAccountIp(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accid, err := strconv.Atoi(c.Param("accid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, accid)
	keys = append(keys, "accid = ?")

	// key param [ip] position [2] type [varchar]
	if len(c.QueryParam("ip")) > 0 {
		ipParam, err := strconv.Atoi(c.QueryParam("ip"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ip] err [%s]", err.Error())})
		}

		params = append(params, ipParam)
		keys = append(keys, "ip = ?")
	}

	// query builder
	var result models.AccountIp
	query := e.db.QueryContext(models.AccountIp{}, c)
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

// getAccountIpsBulk godoc
// @Id getAccountIpsBulk
// @Summary Gets AccountIps in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AccountIp
// @Success 200 {array} models.AccountIp
// @Failure 500 {string} string "Bad query request"
// @Router /account_ips/bulk [post]
func (e *AccountIpController) getAccountIpsBulk(c echo.Context) error {
	var results []models.AccountIp

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

	err := e.db.QueryContext(models.AccountIp{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
