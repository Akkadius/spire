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

type TradeskillRecipeEntryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewTradeskillRecipeEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *TradeskillRecipeEntryController {
	return &TradeskillRecipeEntryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *TradeskillRecipeEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipe_entry/:id", e.getTradeskillRecipeEntry, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipe_entries", e.listTradeskillRecipeEntries, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipe_entries/count", e.getTradeskillRecipeEntriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "tradeskill_recipe_entry", e.createTradeskillRecipeEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "tradeskill_recipe_entry/:id", e.deleteTradeskillRecipeEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "tradeskill_recipe_entry/:id", e.updateTradeskillRecipeEntry, nil),
		routes.RegisterRoute(http.MethodPost, "tradeskill_recipe_entries/bulk", e.getTradeskillRecipeEntriesBulk, nil),
	}
}

// listTradeskillRecipeEntries godoc
// @Id listTradeskillRecipeEntries
// @Summary Lists TradeskillRecipeEntries
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe_entries [get]
func (e *TradeskillRecipeEntryController) listTradeskillRecipeEntries(c echo.Context) error {
	var results []models.TradeskillRecipeEntry
	err := e.db.QueryContext(models.TradeskillRecipeEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTradeskillRecipeEntry godoc
// @Id getTradeskillRecipeEntry
// @Summary Gets TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe_entry/{id} [get]
func (e *TradeskillRecipeEntryController) getTradeskillRecipeEntry(c echo.Context) error {
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
	var result models.TradeskillRecipeEntry
	query := e.db.QueryContext(models.TradeskillRecipeEntry{}, c)
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

// updateTradeskillRecipeEntry godoc
// @Id updateTradeskillRecipeEntry
// @Summary Updates TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param id path int true "Id"
// @Param tradeskill_recipe_entry body models.TradeskillRecipeEntry true "TradeskillRecipeEntry"
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /tradeskill_recipe_entry/{id} [patch]
func (e *TradeskillRecipeEntryController) updateTradeskillRecipeEntry(c echo.Context) error {
	request := new(models.TradeskillRecipeEntry)
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
	var result models.TradeskillRecipeEntry
	query := e.db.QueryContext(models.TradeskillRecipeEntry{}, c)
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
		event := fmt.Sprintf("Updated [TradeskillRecipeEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createTradeskillRecipeEntry godoc
// @Id createTradeskillRecipeEntry
// @Summary Creates TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Param tradeskill_recipe_entry body models.TradeskillRecipeEntry true "TradeskillRecipeEntry"
// @Tags TradeskillRecipeEntry
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /tradeskill_recipe_entry [put]
func (e *TradeskillRecipeEntryController) createTradeskillRecipeEntry(c echo.Context) error {
	tradeskillRecipeEntry := new(models.TradeskillRecipeEntry)
	if err := c.Bind(tradeskillRecipeEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.TradeskillRecipeEntry{}, c).Model(&models.TradeskillRecipeEntry{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
        db = db.Omit(clause.Associations)
    }

	err := db.Create(&tradeskillRecipeEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.TradeskillRecipeEntry{}, tradeskillRecipeEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [TradeskillRecipeEntry] [%v] data [%v]", tradeskillRecipeEntry.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, tradeskillRecipeEntry)
}

// deleteTradeskillRecipeEntry godoc
// @Id deleteTradeskillRecipeEntry
// @Summary Deletes TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tradeskill_recipe_entry/{id} [delete]
func (e *TradeskillRecipeEntryController) deleteTradeskillRecipeEntry(c echo.Context) error {
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
	var result models.TradeskillRecipeEntry
	query := e.db.QueryContext(models.TradeskillRecipeEntry{}, c)
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
		event := fmt.Sprintf("Deleted [TradeskillRecipeEntry] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getTradeskillRecipeEntriesBulk godoc
// @Id getTradeskillRecipeEntriesBulk
// @Summary Gets TradeskillRecipeEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags TradeskillRecipeEntry
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe_entries/bulk [post]
func (e *TradeskillRecipeEntryController) getTradeskillRecipeEntriesBulk(c echo.Context) error {
	var results []models.TradeskillRecipeEntry

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

	err := e.db.QueryContext(models.TradeskillRecipeEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getTradeskillRecipeEntriesCount godoc
// @Id getTradeskillRecipeEntriesCount
// @Summary Counts TradeskillRecipeEntries
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe_entries/count [get]
func (e *TradeskillRecipeEntryController) getTradeskillRecipeEntriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.TradeskillRecipeEntry{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}