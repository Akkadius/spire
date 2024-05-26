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

type Spawn2DisabledController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewSpawn2DisabledController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *Spawn2DisabledController {
	return &Spawn2DisabledController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *Spawn2DisabledController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spawn_2_disabled/:id", e.getSpawn2Disabled, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_2_disableds", e.listSpawn2Disableds, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_2_disableds/count", e.getSpawn2DisabledsCount, nil),
		routes.RegisterRoute(http.MethodPut, "spawn_2_disabled", e.createSpawn2Disabled, nil),
		routes.RegisterRoute(http.MethodDelete, "spawn_2_disabled/:id", e.deleteSpawn2Disabled, nil),
		routes.RegisterRoute(http.MethodPatch, "spawn_2_disabled/:id", e.updateSpawn2Disabled, nil),
		routes.RegisterRoute(http.MethodPost, "spawn_2_disableds/bulk", e.getSpawn2DisabledsBulk, nil),
	}
}

// listSpawn2Disableds godoc
// @Id listSpawn2Disableds
// @Summary Lists Spawn2Disableds
// @Accept json
// @Produce json
// @Tags Spawn2Disabled
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawn2Disabled
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2_disableds [get]
func (e *Spawn2DisabledController) listSpawn2Disableds(c echo.Context) error {
	var results []models.Spawn2Disabled
	err := e.db.QueryContext(models.Spawn2Disabled{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawn2Disabled godoc
// @Id getSpawn2Disabled
// @Summary Gets Spawn2Disabled
// @Accept json
// @Produce json
// @Tags Spawn2Disabled
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawn2Disabled
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2_disabled/{id} [get]
func (e *Spawn2DisabledController) getSpawn2Disabled(c echo.Context) error {
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
	var result models.Spawn2Disabled
	query := e.db.QueryContext(models.Spawn2Disabled{}, c)
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

// updateSpawn2Disabled godoc
// @Id updateSpawn2Disabled
// @Summary Updates Spawn2Disabled
// @Accept json
// @Produce json
// @Tags Spawn2Disabled
// @Param id path int true "Id"
// @Param spawn_2_disabled body models.Spawn2Disabled true "Spawn2Disabled"
// @Success 200 {array} models.Spawn2Disabled
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawn_2_disabled/{id} [patch]
func (e *Spawn2DisabledController) updateSpawn2Disabled(c echo.Context) error {
	request := new(models.Spawn2Disabled)
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
	var result models.Spawn2Disabled
	query := e.db.QueryContext(models.Spawn2Disabled{}, c)
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
		event := fmt.Sprintf("Updated [Spawn2Disabled] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createSpawn2Disabled godoc
// @Id createSpawn2Disabled
// @Summary Creates Spawn2Disabled
// @Accept json
// @Produce json
// @Param spawn_2_disabled body models.Spawn2Disabled true "Spawn2Disabled"
// @Tags Spawn2Disabled
// @Success 200 {array} models.Spawn2Disabled
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawn_2_disabled [put]
func (e *Spawn2DisabledController) createSpawn2Disabled(c echo.Context) error {
	spawn2Disabled := new(models.Spawn2Disabled)
	if err := c.Bind(spawn2Disabled); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.Spawn2Disabled{}, c).Model(&models.Spawn2Disabled{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&spawn2Disabled).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.Spawn2Disabled{}, spawn2Disabled)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [Spawn2Disabled] [%v] data [%v]", spawn2Disabled.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, spawn2Disabled)
}

// deleteSpawn2Disabled godoc
// @Id deleteSpawn2Disabled
// @Summary Deletes Spawn2Disabled
// @Accept json
// @Produce json
// @Tags Spawn2Disabled
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawn_2_disabled/{id} [delete]
func (e *Spawn2DisabledController) deleteSpawn2Disabled(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.Spawn2Disabled
	query := e.db.QueryContext(models.Spawn2Disabled{}, c)
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
		event := fmt.Sprintf("Deleted [Spawn2Disabled] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpawn2DisabledsBulk godoc
// @Id getSpawn2DisabledsBulk
// @Summary Gets Spawn2Disableds in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Spawn2Disabled
// @Success 200 {array} models.Spawn2Disabled
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2_disableds/bulk [post]
func (e *Spawn2DisabledController) getSpawn2DisabledsBulk(c echo.Context) error {
	var results []models.Spawn2Disabled

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

	err := e.db.QueryContext(models.Spawn2Disabled{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawn2DisabledsCount godoc
// @Id getSpawn2DisabledsCount
// @Summary Counts Spawn2Disableds
// @Accept json
// @Produce json
// @Tags Spawn2Disabled
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawn2Disabled
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2_disableds/count [get]
func (e *Spawn2DisabledController) getSpawn2DisabledsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.Spawn2Disabled{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}