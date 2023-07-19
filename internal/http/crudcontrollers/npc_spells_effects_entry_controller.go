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

type NpcSpellsEffectsEntryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewNpcSpellsEffectsEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *NpcSpellsEffectsEntryController {
	return &NpcSpellsEffectsEntryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *NpcSpellsEffectsEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_spells_effects_entry/:id", e.getNpcSpellsEffectsEntry, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells_effects_entries", e.listNpcSpellsEffectsEntries, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells_effects_entries/count", e.getNpcSpellsEffectsEntriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "npc_spells_effects_entry", e.createNpcSpellsEffectsEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_spells_effects_entry/:id", e.deleteNpcSpellsEffectsEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_spells_effects_entry/:id", e.updateNpcSpellsEffectsEntry, nil),
		routes.RegisterRoute(http.MethodPost, "npc_spells_effects_entries/bulk", e.getNpcSpellsEffectsEntriesBulk, nil),
	}
}

// listNpcSpellsEffectsEntries godoc
// @Id listNpcSpellsEffectsEntries
// @Summary Lists NpcSpellsEffectsEntries
// @Accept json
// @Produce json
// @Tags NpcSpellsEffectsEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEffectsEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_effects_entries [get]
func (e *NpcSpellsEffectsEntryController) listNpcSpellsEffectsEntries(c echo.Context) error {
	var results []models.NpcSpellsEffectsEntry
	err := e.db.QueryContext(models.NpcSpellsEffectsEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcSpellsEffectsEntry godoc
// @Id getNpcSpellsEffectsEntry
// @Summary Gets NpcSpellsEffectsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEffectsEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEffectsEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_effects_entry/{id} [get]
func (e *NpcSpellsEffectsEntryController) getNpcSpellsEffectsEntry(c echo.Context) error {
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
	var result models.NpcSpellsEffectsEntry
	query := e.db.QueryContext(models.NpcSpellsEffectsEntry{}, c)
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

// updateNpcSpellsEffectsEntry godoc
// @Id updateNpcSpellsEffectsEntry
// @Summary Updates NpcSpellsEffectsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEffectsEntry
// @Param id path int true "Id"
// @Param npc_spells_effects_entry body models.NpcSpellsEffectsEntry true "NpcSpellsEffectsEntry"
// @Success 200 {array} models.NpcSpellsEffectsEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_spells_effects_entry/{id} [patch]
func (e *NpcSpellsEffectsEntryController) updateNpcSpellsEffectsEntry(c echo.Context) error {
	request := new(models.NpcSpellsEffectsEntry)
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
	var result models.NpcSpellsEffectsEntry
	query := e.db.QueryContext(models.NpcSpellsEffectsEntry{}, c)
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
		event := fmt.Sprintf("Updated [NpcSpellsEffectsEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcSpellsEffectsEntry godoc
// @Id createNpcSpellsEffectsEntry
// @Summary Creates NpcSpellsEffectsEntry
// @Accept json
// @Produce json
// @Param npc_spells_effects_entry body models.NpcSpellsEffectsEntry true "NpcSpellsEffectsEntry"
// @Tags NpcSpellsEffectsEntry
// @Success 200 {array} models.NpcSpellsEffectsEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_spells_effects_entry [put]
func (e *NpcSpellsEffectsEntryController) createNpcSpellsEffectsEntry(c echo.Context) error {
	npcSpellsEffectsEntry := new(models.NpcSpellsEffectsEntry)
	if err := c.Bind(npcSpellsEffectsEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.NpcSpellsEffectsEntry{}, c).Model(&models.NpcSpellsEffectsEntry{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&npcSpellsEffectsEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.NpcSpellsEffectsEntry{}, npcSpellsEffectsEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [NpcSpellsEffectsEntry] [%v] data [%v]", npcSpellsEffectsEntry.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, npcSpellsEffectsEntry)
}

// deleteNpcSpellsEffectsEntry godoc
// @Id deleteNpcSpellsEffectsEntry
// @Summary Deletes NpcSpellsEffectsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEffectsEntry
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_spells_effects_entry/{id} [delete]
func (e *NpcSpellsEffectsEntryController) deleteNpcSpellsEffectsEntry(c echo.Context) error {
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
	var result models.NpcSpellsEffectsEntry
	query := e.db.QueryContext(models.NpcSpellsEffectsEntry{}, c)
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
		event := fmt.Sprintf("Deleted [NpcSpellsEffectsEntry] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcSpellsEffectsEntriesBulk godoc
// @Id getNpcSpellsEffectsEntriesBulk
// @Summary Gets NpcSpellsEffectsEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcSpellsEffectsEntry
// @Success 200 {array} models.NpcSpellsEffectsEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_effects_entries/bulk [post]
func (e *NpcSpellsEffectsEntryController) getNpcSpellsEffectsEntriesBulk(c echo.Context) error {
	var results []models.NpcSpellsEffectsEntry

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

	err := e.db.QueryContext(models.NpcSpellsEffectsEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcSpellsEffectsEntriesCount godoc
// @Id getNpcSpellsEffectsEntriesCount
// @Summary Counts NpcSpellsEffectsEntries
// @Accept json
// @Produce json
// @Tags NpcSpellsEffectsEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEffectsEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_effects_entries/count [get]
func (e *NpcSpellsEffectsEntryController) getNpcSpellsEffectsEntriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.NpcSpellsEffectsEntry{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}