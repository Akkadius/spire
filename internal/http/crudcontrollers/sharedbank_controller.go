package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type SharedbankController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewSharedbankController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *SharedbankController {
	return &SharedbankController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *SharedbankController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "sharedbank/:accountId", e.getSharedbank, nil),
		routes.RegisterRoute(http.MethodGet, "sharedbanks", e.listSharedbanks, nil),
		routes.RegisterRoute(http.MethodGet, "sharedbanks/count", e.getSharedbanksCount, nil),
		routes.RegisterRoute(http.MethodPut, "sharedbank", e.createSharedbank, nil),
		routes.RegisterRoute(http.MethodDelete, "sharedbank/:accountId", e.deleteSharedbank, nil),
		routes.RegisterRoute(http.MethodPatch, "sharedbank/:accountId", e.updateSharedbank, nil),
		routes.RegisterRoute(http.MethodPost, "sharedbanks/bulk", e.getSharedbanksBulk, nil),
	}
}

// listSharedbanks godoc
// @Id listSharedbanks
// @Summary Lists Sharedbanks
// @Accept json
// @Produce json
// @Tags Sharedbank
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Sharedbank
// @Failure 500 {string} string "Bad query request"
// @Router /sharedbanks [get]
func (e *SharedbankController) listSharedbanks(c echo.Context) error {
	var results []models.Sharedbank
	err := e.db.QueryContext(models.Sharedbank{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getSharedbank godoc
// @Id getSharedbank
// @Summary Gets Sharedbank
// @Accept json
// @Produce json
// @Tags Sharedbank
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Sharedbank
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /sharedbank/{id} [get]
func (e *SharedbankController) getSharedbank(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [AccountId]"})
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [slot_id] position [2] type [mediumint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.Sharedbank
	query := e.db.QueryContext(models.Sharedbank{}, c)
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

// updateSharedbank godoc
// @Id updateSharedbank
// @Summary Updates Sharedbank
// @Accept json
// @Produce json
// @Tags Sharedbank
// @Param id path int true "Id"
// @Param sharedbank body models.Sharedbank true "Sharedbank"
// @Success 200 {array} models.Sharedbank
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /sharedbank/{id} [patch]
func (e *SharedbankController) updateSharedbank(c echo.Context) error {
	request := new(models.Sharedbank)
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

	// key param [slot_id] position [2] type [mediumint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.Sharedbank
	query := e.db.QueryContext(models.Sharedbank{}, c)
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
		event := fmt.Sprintf("Updated [Sharedbank] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createSharedbank godoc
// @Id createSharedbank
// @Summary Creates Sharedbank
// @Accept json
// @Produce json
// @Param sharedbank body models.Sharedbank true "Sharedbank"
// @Tags Sharedbank
// @Success 200 {array} models.Sharedbank
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /sharedbank [put]
func (e *SharedbankController) createSharedbank(c echo.Context) error {
	sharedbank := new(models.Sharedbank)
	if err := c.Bind(sharedbank); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.Sharedbank{}, c).Model(&models.Sharedbank{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&sharedbank).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.Sharedbank{}, sharedbank)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [Sharedbank] [%v] data [%v]", sharedbank.AccountId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, sharedbank)
}

// deleteSharedbank godoc
// @Id deleteSharedbank
// @Summary Deletes Sharedbank
// @Accept json
// @Produce json
// @Tags Sharedbank
// @Param id path int true "accountId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /sharedbank/{id} [delete]
func (e *SharedbankController) deleteSharedbank(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [slot_id] position [2] type [mediumint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.Sharedbank
	query := e.db.QueryContext(models.Sharedbank{}, c)
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
		event := fmt.Sprintf("Deleted [Sharedbank] [%v] keys [%v]", result.AccountId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSharedbanksBulk godoc
// @Id getSharedbanksBulk
// @Summary Gets Sharedbanks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Sharedbank
// @Success 200 {array} models.Sharedbank
// @Failure 500 {string} string "Bad query request"
// @Router /sharedbanks/bulk [post]
func (e *SharedbankController) getSharedbanksBulk(c echo.Context) error {
	var results []models.Sharedbank

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

	err := e.db.QueryContext(models.Sharedbank{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getSharedbanksCount godoc
// @Id getSharedbanksCount
// @Summary Counts Sharedbanks
// @Accept json
// @Produce json
// @Tags Sharedbank
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Sharedbank
// @Failure 500 {string} string "Bad query request"
// @Router /sharedbanks/count [get]
func (e *SharedbankController) getSharedbanksCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.Sharedbank{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}