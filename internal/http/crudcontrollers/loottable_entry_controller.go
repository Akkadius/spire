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

type LoottableEntryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewLoottableEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *LoottableEntryController {
	return &LoottableEntryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *LoottableEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "loottable_entry/:loottableId", e.getLoottableEntry, nil),
		routes.RegisterRoute(http.MethodGet, "loottable_entries", e.listLoottableEntries, nil),
		routes.RegisterRoute(http.MethodGet, "loottable_entries/count", e.getLoottableEntriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "loottable_entry", e.createLoottableEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "loottable_entry/:loottableId", e.deleteLoottableEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "loottable_entry/:loottableId", e.updateLoottableEntry, nil),
		routes.RegisterRoute(http.MethodPost, "loottable_entries/bulk", e.getLoottableEntriesBulk, nil),
	}
}

// listLoottableEntries godoc
// @Id listLoottableEntries
// @Summary Lists LoottableEntries
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entries [get]
func (e *LoottableEntryController) listLoottableEntries(c echo.Context) error {
	var results []models.LoottableEntry
	err := e.db.QueryContext(models.LoottableEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoottableEntry godoc
// @Id getLoottableEntry
// @Summary Gets LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoottableEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entry/{id} [get]
func (e *LoottableEntryController) getLoottableEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LoottableId]"})
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.LoottableId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLoottableEntry godoc
// @Id updateLoottableEntry
// @Summary Updates LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "Id"
// @Param loottable_entry body models.LoottableEntry true "LoottableEntry"
// @Success 200 {array} models.LoottableEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /loottable_entry/{id} [patch]
func (e *LoottableEntryController) updateLoottableEntry(c echo.Context) error {
	request := new(models.LoottableEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LoottableId]"})
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
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
		event := fmt.Sprintf("Updated [LoottableEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createLoottableEntry godoc
// @Id createLoottableEntry
// @Summary Creates LoottableEntry
// @Accept json
// @Produce json
// @Param loottable_entry body models.LoottableEntry true "LoottableEntry"
// @Tags LoottableEntry
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /loottable_entry [put]
func (e *LoottableEntryController) createLoottableEntry(c echo.Context) error {
	loottableEntry := new(models.LoottableEntry)
	if err := c.Bind(loottableEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.LoottableEntry{}, c).Model(&models.LoottableEntry{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&loottableEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.LoottableEntry{}, loottableEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [LoottableEntry] [%v] data [%v]", loottableEntry.LoottableId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, loottableEntry)
}

// deleteLoottableEntry godoc
// @Id deleteLoottableEntry
// @Summary Deletes LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "loottableId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /loottable_entry/{id} [delete]
func (e *LoottableEntryController) deleteLoottableEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
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
		event := fmt.Sprintf("Deleted [LoottableEntry] [%v] keys [%v]", result.LoottableId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLoottableEntriesBulk godoc
// @Id getLoottableEntriesBulk
// @Summary Gets LoottableEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoottableEntry
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entries/bulk [post]
func (e *LoottableEntryController) getLoottableEntriesBulk(c echo.Context) error {
	var results []models.LoottableEntry

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

	err := e.db.QueryContext(models.LoottableEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoottableEntriesCount godoc
// @Id getLoottableEntriesCount
// @Summary Counts LoottableEntries
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entries/count [get]
func (e *LoottableEntryController) getLoottableEntriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.LoottableEntry{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}