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

type PlayerEventKilledRaidNpcController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewPlayerEventKilledRaidNpcController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *PlayerEventKilledRaidNpcController {
	return &PlayerEventKilledRaidNpcController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *PlayerEventKilledRaidNpcController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "player_event_killed_raid_npc/:id", e.getPlayerEventKilledRaidNpc, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_killed_raid_npcs", e.listPlayerEventKilledRaidNpcs, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_killed_raid_npcs/count", e.getPlayerEventKilledRaidNpcsCount, nil),
		routes.RegisterRoute(http.MethodPut, "player_event_killed_raid_npc", e.createPlayerEventKilledRaidNpc, nil),
		routes.RegisterRoute(http.MethodDelete, "player_event_killed_raid_npc/:id", e.deletePlayerEventKilledRaidNpc, nil),
		routes.RegisterRoute(http.MethodPatch, "player_event_killed_raid_npc/:id", e.updatePlayerEventKilledRaidNpc, nil),
		routes.RegisterRoute(http.MethodPost, "player_event_killed_raid_npcs/bulk", e.getPlayerEventKilledRaidNpcsBulk, nil),
	}
}

// listPlayerEventKilledRaidNpcs godoc
// @Id listPlayerEventKilledRaidNpcs
// @Summary Lists PlayerEventKilledRaidNpcs
// @Accept json
// @Produce json
// @Tags PlayerEventKilledRaidNpc
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventKilledRaidNpc
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_killed_raid_npcs [get]
func (e *PlayerEventKilledRaidNpcController) listPlayerEventKilledRaidNpcs(c echo.Context) error {
	var results []models.PlayerEventKilledRaidNpc
	err := e.db.QueryContext(models.PlayerEventKilledRaidNpc{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventKilledRaidNpc godoc
// @Id getPlayerEventKilledRaidNpc
// @Summary Gets PlayerEventKilledRaidNpc
// @Accept json
// @Produce json
// @Tags PlayerEventKilledRaidNpc
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventKilledRaidNpc
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_killed_raid_npc/{id} [get]
func (e *PlayerEventKilledRaidNpcController) getPlayerEventKilledRaidNpc(c echo.Context) error {
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
	var result models.PlayerEventKilledRaidNpc
	query := e.db.QueryContext(models.PlayerEventKilledRaidNpc{}, c)
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

// updatePlayerEventKilledRaidNpc godoc
// @Id updatePlayerEventKilledRaidNpc
// @Summary Updates PlayerEventKilledRaidNpc
// @Accept json
// @Produce json
// @Tags PlayerEventKilledRaidNpc
// @Param id path int true "Id"
// @Param player_event_killed_raid_npc body models.PlayerEventKilledRaidNpc true "PlayerEventKilledRaidNpc"
// @Success 200 {array} models.PlayerEventKilledRaidNpc
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /player_event_killed_raid_npc/{id} [patch]
func (e *PlayerEventKilledRaidNpcController) updatePlayerEventKilledRaidNpc(c echo.Context) error {
	request := new(models.PlayerEventKilledRaidNpc)
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
	var result models.PlayerEventKilledRaidNpc
	query := e.db.QueryContext(models.PlayerEventKilledRaidNpc{}, c)
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
		event := fmt.Sprintf("Updated [PlayerEventKilledRaidNpc] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createPlayerEventKilledRaidNpc godoc
// @Id createPlayerEventKilledRaidNpc
// @Summary Creates PlayerEventKilledRaidNpc
// @Accept json
// @Produce json
// @Param player_event_killed_raid_npc body models.PlayerEventKilledRaidNpc true "PlayerEventKilledRaidNpc"
// @Tags PlayerEventKilledRaidNpc
// @Success 200 {array} models.PlayerEventKilledRaidNpc
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /player_event_killed_raid_npc [put]
func (e *PlayerEventKilledRaidNpcController) createPlayerEventKilledRaidNpc(c echo.Context) error {
	playerEventKilledRaidNpc := new(models.PlayerEventKilledRaidNpc)
	if err := c.Bind(playerEventKilledRaidNpc); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.PlayerEventKilledRaidNpc{}, c).Model(&models.PlayerEventKilledRaidNpc{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&playerEventKilledRaidNpc).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.PlayerEventKilledRaidNpc{}, playerEventKilledRaidNpc)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [PlayerEventKilledRaidNpc] [%v] data [%v]", playerEventKilledRaidNpc.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, playerEventKilledRaidNpc)
}

// deletePlayerEventKilledRaidNpc godoc
// @Id deletePlayerEventKilledRaidNpc
// @Summary Deletes PlayerEventKilledRaidNpc
// @Accept json
// @Produce json
// @Tags PlayerEventKilledRaidNpc
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /player_event_killed_raid_npc/{id} [delete]
func (e *PlayerEventKilledRaidNpcController) deletePlayerEventKilledRaidNpc(c echo.Context) error {
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
	var result models.PlayerEventKilledRaidNpc
	query := e.db.QueryContext(models.PlayerEventKilledRaidNpc{}, c)
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
		event := fmt.Sprintf("Deleted [PlayerEventKilledRaidNpc] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getPlayerEventKilledRaidNpcsBulk godoc
// @Id getPlayerEventKilledRaidNpcsBulk
// @Summary Gets PlayerEventKilledRaidNpcs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PlayerEventKilledRaidNpc
// @Success 200 {array} models.PlayerEventKilledRaidNpc
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_killed_raid_npcs/bulk [post]
func (e *PlayerEventKilledRaidNpcController) getPlayerEventKilledRaidNpcsBulk(c echo.Context) error {
	var results []models.PlayerEventKilledRaidNpc

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

	err := e.db.QueryContext(models.PlayerEventKilledRaidNpc{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventKilledRaidNpcsCount godoc
// @Id getPlayerEventKilledRaidNpcsCount
// @Summary Counts PlayerEventKilledRaidNpcs
// @Accept json
// @Produce json
// @Tags PlayerEventKilledRaidNpc
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventKilledRaidNpc
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_killed_raid_npcs/count [get]
func (e *PlayerEventKilledRaidNpcController) getPlayerEventKilledRaidNpcsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.PlayerEventKilledRaidNpc{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}