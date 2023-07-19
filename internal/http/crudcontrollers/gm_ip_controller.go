package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type GmIpController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewGmIpController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *GmIpController {
	return &GmIpController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *GmIpController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "gm_ip/:accountId", e.getGmIp, nil),
		routes.RegisterRoute(http.MethodGet, "gm_ips", e.listGmIps, nil),
		routes.RegisterRoute(http.MethodGet, "gm_ips/count", e.getGmIpsCount, nil),
		routes.RegisterRoute(http.MethodPut, "gm_ip", e.createGmIp, nil),
		routes.RegisterRoute(http.MethodDelete, "gm_ip/:accountId", e.deleteGmIp, nil),
		routes.RegisterRoute(http.MethodPatch, "gm_ip/:accountId", e.updateGmIp, nil),
		routes.RegisterRoute(http.MethodPost, "gm_ips/bulk", e.getGmIpsBulk, nil),
	}
}

// listGmIps godoc
// @Id listGmIps
// @Summary Lists GmIps
// @Accept json
// @Produce json
// @Tags GmIp
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GmIp
// @Failure 500 {string} string "Bad query request"
// @Router /gm_ips [get]
func (e *GmIpController) listGmIps(c echo.Context) error {
	var results []models.GmIp
	err := e.db.QueryContext(models.GmIp{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGmIp godoc
// @Id getGmIp
// @Summary Gets GmIp
// @Accept json
// @Produce json
// @Tags GmIp
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GmIp
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /gm_ip/{id} [get]
func (e *GmIpController) getGmIp(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [AccountId]"})
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [ip_address] position [3] type [varchar]
	if len(c.QueryParam("ip_address")) > 0 {
		ipAddressParam, err := strconv.Atoi(c.QueryParam("ip_address"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ip_address] err [%s]", err.Error())})
		}

		params = append(params, ipAddressParam)
		keys = append(keys, "ip_address = ?")
	}

	// query builder
	var result models.GmIp
	query := e.db.QueryContext(models.GmIp{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.AccountId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGmIp godoc
// @Id updateGmIp
// @Summary Updates GmIp
// @Accept json
// @Produce json
// @Tags GmIp
// @Param id path int true "Id"
// @Param gm_ip body models.GmIp true "GmIp"
// @Success 200 {array} models.GmIp
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /gm_ip/{id} [patch]
func (e *GmIpController) updateGmIp(c echo.Context) error {
	request := new(models.GmIp)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [AccountId]"})
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [ip_address] position [3] type [varchar]
	if len(c.QueryParam("ip_address")) > 0 {
		ipAddressParam, err := strconv.Atoi(c.QueryParam("ip_address"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ip_address] err [%s]", err.Error())})
		}

		params = append(params, ipAddressParam)
		keys = append(keys, "ip_address = ?")
	}

	// query builder
	var result models.GmIp
	query := e.db.QueryContext(models.GmIp{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [GmIp] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createGmIp godoc
// @Id createGmIp
// @Summary Creates GmIp
// @Accept json
// @Produce json
// @Param gm_ip body models.GmIp true "GmIp"
// @Tags GmIp
// @Success 200 {array} models.GmIp
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /gm_ip [put]
func (e *GmIpController) createGmIp(c echo.Context) error {
	gmIp := new(models.GmIp)
	if err := c.Bind(gmIp); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.GmIp{}, c).Model(&models.GmIp{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
        db = db.Omit(clause.Associations)
    }

	err := db.Create(&gmIp).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GmIp{}, gmIp)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GmIp] [%v] data [%v]", gmIp.AccountId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, gmIp)
}

// deleteGmIp godoc
// @Id deleteGmIp
// @Summary Deletes GmIp
// @Accept json
// @Produce json
// @Tags GmIp
// @Param id path int true "accountId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /gm_ip/{id} [delete]
func (e *GmIpController) deleteGmIp(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [ip_address] position [3] type [varchar]
	if len(c.QueryParam("ip_address")) > 0 {
		ipAddressParam, err := strconv.Atoi(c.QueryParam("ip_address"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ip_address] err [%s]", err.Error())})
		}

		params = append(params, ipAddressParam)
		keys = append(keys, "ip_address = ?")
	}

	// query builder
	var result models.GmIp
	query := e.db.QueryContext(models.GmIp{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [GmIp] [%v] keys [%v]", result.AccountId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGmIpsBulk godoc
// @Id getGmIpsBulk
// @Summary Gets GmIps in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GmIp
// @Success 200 {array} models.GmIp
// @Failure 500 {string} string "Bad query request"
// @Router /gm_ips/bulk [post]
func (e *GmIpController) getGmIpsBulk(c echo.Context) error {
	var results []models.GmIp

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

	err := e.db.QueryContext(models.GmIp{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGmIpsCount godoc
// @Id getGmIpsCount
// @Summary Counts GmIps
// @Accept json
// @Produce json
// @Tags GmIp
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GmIp
// @Failure 500 {string} string "Bad query request"
// @Router /gm_ips/count [get]
func (e *GmIpController) getGmIpsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.GmIp{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}