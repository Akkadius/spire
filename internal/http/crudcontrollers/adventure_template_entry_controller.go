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

type AdventureTemplateEntryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewAdventureTemplateEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *AdventureTemplateEntryController {
	return &AdventureTemplateEntryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *AdventureTemplateEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "adventure_template_entry/:id", e.getAdventureTemplateEntry, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_template_entries", e.listAdventureTemplateEntries, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_template_entries/count", e.getAdventureTemplateEntriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_template_entry", e.createAdventureTemplateEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "adventure_template_entry/:id", e.deleteAdventureTemplateEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_template_entry/:id", e.updateAdventureTemplateEntry, nil),
		routes.RegisterRoute(http.MethodPost, "adventure_template_entries/bulk", e.getAdventureTemplateEntriesBulk, nil),
	}
}

// listAdventureTemplateEntries godoc
// @Id listAdventureTemplateEntries
// @Summary Lists AdventureTemplateEntries
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entries [get]
func (e *AdventureTemplateEntryController) listAdventureTemplateEntries(c echo.Context) error {
	var results []models.AdventureTemplateEntry
	err := e.db.QueryContext(models.AdventureTemplateEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureTemplateEntry godoc
// @Id getAdventureTemplateEntry
// @Summary Gets AdventureTemplateEntry
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entry/{id} [get]
func (e *AdventureTemplateEntryController) getAdventureTemplateEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [template_id] position [2] type [int]
	if len(c.QueryParam("template_id")) > 0 {
		templateIdParam, err := strconv.Atoi(c.QueryParam("template_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [template_id] err [%s]", err.Error())})
		}

		params = append(params, templateIdParam)
		keys = append(keys, "template_id = ?")
	}

	// query builder
	var result models.AdventureTemplateEntry
	query := e.db.QueryContext(models.AdventureTemplateEntry{}, c)
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

// updateAdventureTemplateEntry godoc
// @Id updateAdventureTemplateEntry
// @Summary Updates AdventureTemplateEntry
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param id path int true "Id"
// @Param adventure_template_entry body models.AdventureTemplateEntry true "AdventureTemplateEntry"
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_template_entry/{id} [patch]
func (e *AdventureTemplateEntryController) updateAdventureTemplateEntry(c echo.Context) error {
	request := new(models.AdventureTemplateEntry)
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

	// key param [template_id] position [2] type [int]
	if len(c.QueryParam("template_id")) > 0 {
		templateIdParam, err := strconv.Atoi(c.QueryParam("template_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [template_id] err [%s]", err.Error())})
		}

		params = append(params, templateIdParam)
		keys = append(keys, "template_id = ?")
	}

	// query builder
	var result models.AdventureTemplateEntry
	query := e.db.QueryContext(models.AdventureTemplateEntry{}, c)
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
		event := fmt.Sprintf("Updated [AdventureTemplateEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createAdventureTemplateEntry godoc
// @Id createAdventureTemplateEntry
// @Summary Creates AdventureTemplateEntry
// @Accept json
// @Produce json
// @Param adventure_template_entry body models.AdventureTemplateEntry true "AdventureTemplateEntry"
// @Tags AdventureTemplateEntry
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_template_entry [put]
func (e *AdventureTemplateEntryController) createAdventureTemplateEntry(c echo.Context) error {
	adventureTemplateEntry := new(models.AdventureTemplateEntry)
	if err := c.Bind(adventureTemplateEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.AdventureTemplateEntry{}, c).Model(&models.AdventureTemplateEntry{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&adventureTemplateEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.AdventureTemplateEntry{}, adventureTemplateEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [AdventureTemplateEntry] [%v] data [%v]", adventureTemplateEntry.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, adventureTemplateEntry)
}

// deleteAdventureTemplateEntry godoc
// @Id deleteAdventureTemplateEntry
// @Summary Deletes AdventureTemplateEntry
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_template_entry/{id} [delete]
func (e *AdventureTemplateEntryController) deleteAdventureTemplateEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [template_id] position [2] type [int]
	if len(c.QueryParam("template_id")) > 0 {
		templateIdParam, err := strconv.Atoi(c.QueryParam("template_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [template_id] err [%s]", err.Error())})
		}

		params = append(params, templateIdParam)
		keys = append(keys, "template_id = ?")
	}

	// query builder
	var result models.AdventureTemplateEntry
	query := e.db.QueryContext(models.AdventureTemplateEntry{}, c)
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
		event := fmt.Sprintf("Deleted [AdventureTemplateEntry] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getAdventureTemplateEntriesBulk godoc
// @Id getAdventureTemplateEntriesBulk
// @Summary Gets AdventureTemplateEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AdventureTemplateEntry
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entries/bulk [post]
func (e *AdventureTemplateEntryController) getAdventureTemplateEntriesBulk(c echo.Context) error {
	var results []models.AdventureTemplateEntry

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

	err := e.db.QueryContext(models.AdventureTemplateEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureTemplateEntriesCount godoc
// @Id getAdventureTemplateEntriesCount
// @Summary Counts AdventureTemplateEntries
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entries/count [get]
func (e *AdventureTemplateEntryController) getAdventureTemplateEntriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.AdventureTemplateEntry{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}