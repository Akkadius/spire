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

type ZoneStateSpawnController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewZoneStateSpawnController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *ZoneStateSpawnController {
	return &ZoneStateSpawnController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *ZoneStateSpawnController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "zone_state_spawn/:id", e.getZoneStateSpawn, nil),
		routes.RegisterRoute(http.MethodGet, "zone_state_spawns", e.listZoneStateSpawns, nil),
		routes.RegisterRoute(http.MethodGet, "zone_state_spawns/count", e.getZoneStateSpawnsCount, nil),
		routes.RegisterRoute(http.MethodPut, "zone_state_spawn", e.createZoneStateSpawn, nil),
		routes.RegisterRoute(http.MethodDelete, "zone_state_spawn/:id", e.deleteZoneStateSpawn, nil),
		routes.RegisterRoute(http.MethodPatch, "zone_state_spawn/:id", e.updateZoneStateSpawn, nil),
		routes.RegisterRoute(http.MethodPost, "zone_state_spawns/bulk", e.getZoneStateSpawnsBulk, nil),
	}
}

// listZoneStateSpawns godoc
// @Id listZoneStateSpawns
// @Summary Lists ZoneStateSpawns
// @Accept json
// @Produce json
// @Tags ZoneStateSpawn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZoneStateSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /zone_state_spawns [get]
func (e *ZoneStateSpawnController) listZoneStateSpawns(c echo.Context) error {
	var results []models.ZoneStateSpawn
	err := e.db.QueryContext(models.ZoneStateSpawn{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getZoneStateSpawn godoc
// @Id getZoneStateSpawn
// @Summary Gets ZoneStateSpawn
// @Accept json
// @Produce json
// @Tags ZoneStateSpawn
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZoneStateSpawn
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /zone_state_spawn/{id} [get]
func (e *ZoneStateSpawnController) getZoneStateSpawn(c echo.Context) error {
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
	var result models.ZoneStateSpawn
	query := e.db.QueryContext(models.ZoneStateSpawn{}, c)
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

// updateZoneStateSpawn godoc
// @Id updateZoneStateSpawn
// @Summary Updates ZoneStateSpawn
// @Accept json
// @Produce json
// @Tags ZoneStateSpawn
// @Param id path int true "Id"
// @Param zone_state_spawn body models.ZoneStateSpawn true "ZoneStateSpawn"
// @Success 200 {array} models.ZoneStateSpawn
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /zone_state_spawn/{id} [patch]
func (e *ZoneStateSpawnController) updateZoneStateSpawn(c echo.Context) error {
	request := new(models.ZoneStateSpawn)
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
	var result models.ZoneStateSpawn
	query := e.db.QueryContext(models.ZoneStateSpawn{}, c)
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
		event := fmt.Sprintf("Updated [ZoneStateSpawn] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createZoneStateSpawn godoc
// @Id createZoneStateSpawn
// @Summary Creates ZoneStateSpawn
// @Accept json
// @Produce json
// @Param zone_state_spawn body models.ZoneStateSpawn true "ZoneStateSpawn"
// @Tags ZoneStateSpawn
// @Success 200 {array} models.ZoneStateSpawn
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /zone_state_spawn [put]
func (e *ZoneStateSpawnController) createZoneStateSpawn(c echo.Context) error {
	zoneStateSpawn := new(models.ZoneStateSpawn)
	if err := c.Bind(zoneStateSpawn); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.ZoneStateSpawn{}, c).Model(&models.ZoneStateSpawn{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&zoneStateSpawn).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.ZoneStateSpawn{}, zoneStateSpawn)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [ZoneStateSpawn] [%v] data [%v]", zoneStateSpawn.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, zoneStateSpawn)
}

// deleteZoneStateSpawn godoc
// @Id deleteZoneStateSpawn
// @Summary Deletes ZoneStateSpawn
// @Accept json
// @Produce json
// @Tags ZoneStateSpawn
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /zone_state_spawn/{id} [delete]
func (e *ZoneStateSpawnController) deleteZoneStateSpawn(c echo.Context) error {
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
	var result models.ZoneStateSpawn
	query := e.db.QueryContext(models.ZoneStateSpawn{}, c)
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
		event := fmt.Sprintf("Deleted [ZoneStateSpawn] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getZoneStateSpawnsBulk godoc
// @Id getZoneStateSpawnsBulk
// @Summary Gets ZoneStateSpawns in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ZoneStateSpawn
// @Success 200 {array} models.ZoneStateSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /zone_state_spawns/bulk [post]
func (e *ZoneStateSpawnController) getZoneStateSpawnsBulk(c echo.Context) error {
	var results []models.ZoneStateSpawn

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

	err := e.db.QueryContext(models.ZoneStateSpawn{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getZoneStateSpawnsCount godoc
// @Id getZoneStateSpawnsCount
// @Summary Counts ZoneStateSpawns
// @Accept json
// @Produce json
// @Tags ZoneStateSpawn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZoneStateSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /zone_state_spawns/count [get]
func (e *ZoneStateSpawnController) getZoneStateSpawnsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.ZoneStateSpawn{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}