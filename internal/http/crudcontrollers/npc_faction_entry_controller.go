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
	"net/http"
	"strconv"
	"strings"
)

type NpcFactionEntryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewNpcFactionEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *NpcFactionEntryController {
	return &NpcFactionEntryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *NpcFactionEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_faction_entry/:npcFactionId", e.getNpcFactionEntry, nil),
		routes.RegisterRoute(http.MethodGet, "npc_faction_entries", e.listNpcFactionEntries, nil),
		routes.RegisterRoute(http.MethodGet, "npc_faction_entries/count", e.getNpcFactionEntriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "npc_faction_entry", e.createNpcFactionEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_faction_entry/:npcFactionId", e.deleteNpcFactionEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_faction_entry/:npcFactionId", e.updateNpcFactionEntry, nil),
		routes.RegisterRoute(http.MethodPost, "npc_faction_entries/bulk", e.getNpcFactionEntriesBulk, nil),
	}
}

// listNpcFactionEntries godoc
// @Id listNpcFactionEntries
// @Summary Lists NpcFactionEntries
// @Accept json
// @Produce json
// @Tags NpcFactionEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcFactionEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_faction_entries [get]
func (e *NpcFactionEntryController) listNpcFactionEntries(c echo.Context) error {
	var results []models.NpcFactionEntry
	err := e.db.QueryContext(models.NpcFactionEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcFactionEntry godoc
// @Id getNpcFactionEntry
// @Summary Gets NpcFactionEntry
// @Accept json
// @Produce json
// @Tags NpcFactionEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcFactionEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_faction_entry/{id} [get]
func (e *NpcFactionEntryController) getNpcFactionEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	npcFactionId, err := strconv.Atoi(c.Param("npcFactionId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [NpcFactionId]"})
	}
	params = append(params, npcFactionId)
	keys = append(keys, "npc_faction_id = ?")

	// key param [faction_id] position [2] type [int]
	if len(c.QueryParam("faction_id")) > 0 {
		factionIdParam, err := strconv.Atoi(c.QueryParam("faction_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [faction_id] err [%s]", err.Error())})
		}

		params = append(params, factionIdParam)
		keys = append(keys, "faction_id = ?")
	}

	// query builder
	var result models.NpcFactionEntry
	query := e.db.QueryContext(models.NpcFactionEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.NpcFactionId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateNpcFactionEntry godoc
// @Id updateNpcFactionEntry
// @Summary Updates NpcFactionEntry
// @Accept json
// @Produce json
// @Tags NpcFactionEntry
// @Param id path int true "Id"
// @Param npc_faction_entry body models.NpcFactionEntry true "NpcFactionEntry"
// @Success 200 {array} models.NpcFactionEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_faction_entry/{id} [patch]
func (e *NpcFactionEntryController) updateNpcFactionEntry(c echo.Context) error {
	request := new(models.NpcFactionEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	npcFactionId, err := strconv.Atoi(c.Param("npcFactionId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [NpcFactionId]"})
	}
	params = append(params, npcFactionId)
	keys = append(keys, "npc_faction_id = ?")

	// key param [faction_id] position [2] type [int]
	if len(c.QueryParam("faction_id")) > 0 {
		factionIdParam, err := strconv.Atoi(c.QueryParam("faction_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [faction_id] err [%s]", err.Error())})
		}

		params = append(params, factionIdParam)
		keys = append(keys, "faction_id = ?")
	}

	// query builder
	var result models.NpcFactionEntry
	query := e.db.QueryContext(models.NpcFactionEntry{}, c)
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
		event := fmt.Sprintf("Updated [NpcFactionEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcFactionEntry godoc
// @Id createNpcFactionEntry
// @Summary Creates NpcFactionEntry
// @Accept json
// @Produce json
// @Param npc_faction_entry body models.NpcFactionEntry true "NpcFactionEntry"
// @Tags NpcFactionEntry
// @Success 200 {array} models.NpcFactionEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_faction_entry [put]
func (e *NpcFactionEntryController) createNpcFactionEntry(c echo.Context) error {
	npcFactionEntry := new(models.NpcFactionEntry)
	if err := c.Bind(npcFactionEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcFactionEntry{}, c).Model(&models.NpcFactionEntry{}).Create(&npcFactionEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.NpcFactionEntry{}, npcFactionEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [NpcFactionEntry] [%v] data [%v]", npcFactionEntry.NpcFactionId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, npcFactionEntry)
}

// deleteNpcFactionEntry godoc
// @Id deleteNpcFactionEntry
// @Summary Deletes NpcFactionEntry
// @Accept json
// @Produce json
// @Tags NpcFactionEntry
// @Param id path int true "npcFactionId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_faction_entry/{id} [delete]
func (e *NpcFactionEntryController) deleteNpcFactionEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	npcFactionId, err := strconv.Atoi(c.Param("npcFactionId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, npcFactionId)
	keys = append(keys, "npc_faction_id = ?")

	// key param [faction_id] position [2] type [int]
	if len(c.QueryParam("faction_id")) > 0 {
		factionIdParam, err := strconv.Atoi(c.QueryParam("faction_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [faction_id] err [%s]", err.Error())})
		}

		params = append(params, factionIdParam)
		keys = append(keys, "faction_id = ?")
	}

	// query builder
	var result models.NpcFactionEntry
	query := e.db.QueryContext(models.NpcFactionEntry{}, c)
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
		event := fmt.Sprintf("Deleted [NpcFactionEntry] [%v] keys [%v]", result.NpcFactionId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcFactionEntriesBulk godoc
// @Id getNpcFactionEntriesBulk
// @Summary Gets NpcFactionEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcFactionEntry
// @Success 200 {array} models.NpcFactionEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_faction_entries/bulk [post]
func (e *NpcFactionEntryController) getNpcFactionEntriesBulk(c echo.Context) error {
	var results []models.NpcFactionEntry

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

	err := e.db.QueryContext(models.NpcFactionEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcFactionEntriesCount godoc
// @Id getNpcFactionEntriesCount
// @Summary Counts NpcFactionEntries
// @Accept json
// @Produce json
// @Tags NpcFactionEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcFactionEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_faction_entries/count [get]
func (e *NpcFactionEntryController) getNpcFactionEntriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.NpcFactionEntry{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}