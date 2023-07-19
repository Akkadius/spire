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

type LootdropEntryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewLootdropEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *LootdropEntryController {
	return &LootdropEntryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *LootdropEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "lootdrop_entry/:lootdropId", e.getLootdropEntry, nil),
		routes.RegisterRoute(http.MethodGet, "lootdrop_entries", e.listLootdropEntries, nil),
		routes.RegisterRoute(http.MethodGet, "lootdrop_entries/count", e.getLootdropEntriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "lootdrop_entry", e.createLootdropEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "lootdrop_entry/:lootdropId", e.deleteLootdropEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "lootdrop_entry/:lootdropId", e.updateLootdropEntry, nil),
		routes.RegisterRoute(http.MethodPost, "lootdrop_entries/bulk", e.getLootdropEntriesBulk, nil),
	}
}

// listLootdropEntries godoc
// @Id listLootdropEntries
// @Summary Lists LootdropEntries
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entries [get]
func (e *LootdropEntryController) listLootdropEntries(c echo.Context) error {
	var results []models.LootdropEntry
	err := e.db.QueryContext(models.LootdropEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLootdropEntry godoc
// @Id getLootdropEntry
// @Summary Gets LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LootdropEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entry/{id} [get]
func (e *LootdropEntryController) getLootdropEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LootdropId]"})
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.LootdropId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLootdropEntry godoc
// @Id updateLootdropEntry
// @Summary Updates LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "Id"
// @Param lootdrop_entry body models.LootdropEntry true "LootdropEntry"
// @Success 200 {array} models.LootdropEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /lootdrop_entry/{id} [patch]
func (e *LootdropEntryController) updateLootdropEntry(c echo.Context) error {
	request := new(models.LootdropEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LootdropId]"})
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
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
		event := fmt.Sprintf("Updated [LootdropEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createLootdropEntry godoc
// @Id createLootdropEntry
// @Summary Creates LootdropEntry
// @Accept json
// @Produce json
// @Param lootdrop_entry body models.LootdropEntry true "LootdropEntry"
// @Tags LootdropEntry
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /lootdrop_entry [put]
func (e *LootdropEntryController) createLootdropEntry(c echo.Context) error {
	lootdropEntry := new(models.LootdropEntry)
	if err := c.Bind(lootdropEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.LootdropEntry{}, c).Model(&models.LootdropEntry{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&lootdropEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.LootdropEntry{}, lootdropEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [LootdropEntry] [%v] data [%v]", lootdropEntry.LootdropId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, lootdropEntry)
}

// deleteLootdropEntry godoc
// @Id deleteLootdropEntry
// @Summary Deletes LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "lootdropId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /lootdrop_entry/{id} [delete]
func (e *LootdropEntryController) deleteLootdropEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
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
		event := fmt.Sprintf("Deleted [LootdropEntry] [%v] keys [%v]", result.LootdropId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLootdropEntriesBulk godoc
// @Id getLootdropEntriesBulk
// @Summary Gets LootdropEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LootdropEntry
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entries/bulk [post]
func (e *LootdropEntryController) getLootdropEntriesBulk(c echo.Context) error {
	var results []models.LootdropEntry

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

	err := e.db.QueryContext(models.LootdropEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getLootdropEntriesCount godoc
// @Id getLootdropEntriesCount
// @Summary Counts LootdropEntries
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entries/count [get]
func (e *LootdropEntryController) getLootdropEntriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.LootdropEntry{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}