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

type InventorySnapshotController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewInventorySnapshotController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *InventorySnapshotController {
	return &InventorySnapshotController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *InventorySnapshotController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "inventory_snapshot/:timeIndex", e.getInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodGet, "inventory_snapshots", e.listInventorySnapshots, nil),
		routes.RegisterRoute(http.MethodGet, "inventory_snapshots/count", e.getInventorySnapshotsCount, nil),
		routes.RegisterRoute(http.MethodPut, "inventory_snapshot", e.createInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodDelete, "inventory_snapshot/:timeIndex", e.deleteInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodPatch, "inventory_snapshot/:timeIndex", e.updateInventorySnapshot, nil),
		routes.RegisterRoute(http.MethodPost, "inventory_snapshots/bulk", e.getInventorySnapshotsBulk, nil),
	}
}

// listInventorySnapshots godoc
// @Id listInventorySnapshots
// @Summary Lists InventorySnapshots
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InventorySnapshot
// @Failure 500 {string} string "Bad query request"
// @Router /inventory_snapshots [get]
func (e *InventorySnapshotController) listInventorySnapshots(c echo.Context) error {
	var results []models.InventorySnapshot
	err := e.db.QueryContext(models.InventorySnapshot{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getInventorySnapshot godoc
// @Id getInventorySnapshot
// @Summary Gets InventorySnapshot
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InventorySnapshot
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /inventory_snapshot/{id} [get]
func (e *InventorySnapshotController) getInventorySnapshot(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	timeIndex, err := strconv.Atoi(c.Param("timeIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TimeIndex]"})
	}
	params = append(params, timeIndex)
	keys = append(keys, "time_index = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [slotid] position [3] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.InventorySnapshot
	query := e.db.QueryContext(models.InventorySnapshot{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.TimeIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateInventorySnapshot godoc
// @Id updateInventorySnapshot
// @Summary Updates InventorySnapshot
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param id path int true "Id"
// @Param inventory_snapshot body models.InventorySnapshot true "InventorySnapshot"
// @Success 200 {array} models.InventorySnapshot
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /inventory_snapshot/{id} [patch]
func (e *InventorySnapshotController) updateInventorySnapshot(c echo.Context) error {
	request := new(models.InventorySnapshot)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	timeIndex, err := strconv.Atoi(c.Param("timeIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TimeIndex]"})
	}
	params = append(params, timeIndex)
	keys = append(keys, "time_index = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [slotid] position [3] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.InventorySnapshot
	query := e.db.QueryContext(models.InventorySnapshot{}, c)
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
		event := fmt.Sprintf("Updated [InventorySnapshot] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createInventorySnapshot godoc
// @Id createInventorySnapshot
// @Summary Creates InventorySnapshot
// @Accept json
// @Produce json
// @Param inventory_snapshot body models.InventorySnapshot true "InventorySnapshot"
// @Tags InventorySnapshot
// @Success 200 {array} models.InventorySnapshot
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /inventory_snapshot [put]
func (e *InventorySnapshotController) createInventorySnapshot(c echo.Context) error {
	inventorySnapshot := new(models.InventorySnapshot)
	if err := c.Bind(inventorySnapshot); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.InventorySnapshot{}, c).Model(&models.InventorySnapshot{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&inventorySnapshot).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.InventorySnapshot{}, inventorySnapshot)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [InventorySnapshot] [%v] data [%v]", inventorySnapshot.TimeIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, inventorySnapshot)
}

// deleteInventorySnapshot godoc
// @Id deleteInventorySnapshot
// @Summary Deletes InventorySnapshot
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param id path int true "timeIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /inventory_snapshot/{id} [delete]
func (e *InventorySnapshotController) deleteInventorySnapshot(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	timeIndex, err := strconv.Atoi(c.Param("timeIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, timeIndex)
	keys = append(keys, "time_index = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [slotid] position [3] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.InventorySnapshot
	query := e.db.QueryContext(models.InventorySnapshot{}, c)
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
		event := fmt.Sprintf("Deleted [InventorySnapshot] [%v] keys [%v]", result.TimeIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getInventorySnapshotsBulk godoc
// @Id getInventorySnapshotsBulk
// @Summary Gets InventorySnapshots in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags InventorySnapshot
// @Success 200 {array} models.InventorySnapshot
// @Failure 500 {string} string "Bad query request"
// @Router /inventory_snapshots/bulk [post]
func (e *InventorySnapshotController) getInventorySnapshotsBulk(c echo.Context) error {
	var results []models.InventorySnapshot

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

	err := e.db.QueryContext(models.InventorySnapshot{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getInventorySnapshotsCount godoc
// @Id getInventorySnapshotsCount
// @Summary Counts InventorySnapshots
// @Accept json
// @Produce json
// @Tags InventorySnapshot
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InventorySnapshot
// @Failure 500 {string} string "Bad query request"
// @Router /inventory_snapshots/count [get]
func (e *InventorySnapshotController) getInventorySnapshotsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.InventorySnapshot{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}