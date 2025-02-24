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

type PlayerEventTradeController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewPlayerEventTradeController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *PlayerEventTradeController {
	return &PlayerEventTradeController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *PlayerEventTradeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "player_event_trade/:id", e.getPlayerEventTrade, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_trades", e.listPlayerEventTrades, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_trades/count", e.getPlayerEventTradesCount, nil),
		routes.RegisterRoute(http.MethodPut, "player_event_trade", e.createPlayerEventTrade, nil),
		routes.RegisterRoute(http.MethodDelete, "player_event_trade/:id", e.deletePlayerEventTrade, nil),
		routes.RegisterRoute(http.MethodPatch, "player_event_trade/:id", e.updatePlayerEventTrade, nil),
		routes.RegisterRoute(http.MethodPost, "player_event_trades/bulk", e.getPlayerEventTradesBulk, nil),
	}
}

// listPlayerEventTrades godoc
// @Id listPlayerEventTrades
// @Summary Lists PlayerEventTrades
// @Accept json
// @Produce json
// @Tags PlayerEventTrade
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventTrade
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_trades [get]
func (e *PlayerEventTradeController) listPlayerEventTrades(c echo.Context) error {
	var results []models.PlayerEventTrade
	err := e.db.QueryContext(models.PlayerEventTrade{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventTrade godoc
// @Id getPlayerEventTrade
// @Summary Gets PlayerEventTrade
// @Accept json
// @Produce json
// @Tags PlayerEventTrade
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventTrade
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_trade/{id} [get]
func (e *PlayerEventTradeController) getPlayerEventTrade(c echo.Context) error {
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
	var result models.PlayerEventTrade
	query := e.db.QueryContext(models.PlayerEventTrade{}, c)
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

// updatePlayerEventTrade godoc
// @Id updatePlayerEventTrade
// @Summary Updates PlayerEventTrade
// @Accept json
// @Produce json
// @Tags PlayerEventTrade
// @Param id path int true "Id"
// @Param player_event_trade body models.PlayerEventTrade true "PlayerEventTrade"
// @Success 200 {array} models.PlayerEventTrade
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /player_event_trade/{id} [patch]
func (e *PlayerEventTradeController) updatePlayerEventTrade(c echo.Context) error {
	request := new(models.PlayerEventTrade)
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
	var result models.PlayerEventTrade
	query := e.db.QueryContext(models.PlayerEventTrade{}, c)
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
		event := fmt.Sprintf("Updated [PlayerEventTrade] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createPlayerEventTrade godoc
// @Id createPlayerEventTrade
// @Summary Creates PlayerEventTrade
// @Accept json
// @Produce json
// @Param player_event_trade body models.PlayerEventTrade true "PlayerEventTrade"
// @Tags PlayerEventTrade
// @Success 200 {array} models.PlayerEventTrade
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /player_event_trade [put]
func (e *PlayerEventTradeController) createPlayerEventTrade(c echo.Context) error {
	playerEventTrade := new(models.PlayerEventTrade)
	if err := c.Bind(playerEventTrade); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.PlayerEventTrade{}, c).Model(&models.PlayerEventTrade{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&playerEventTrade).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.PlayerEventTrade{}, playerEventTrade)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [PlayerEventTrade] [%v] data [%v]", playerEventTrade.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, playerEventTrade)
}

// deletePlayerEventTrade godoc
// @Id deletePlayerEventTrade
// @Summary Deletes PlayerEventTrade
// @Accept json
// @Produce json
// @Tags PlayerEventTrade
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /player_event_trade/{id} [delete]
func (e *PlayerEventTradeController) deletePlayerEventTrade(c echo.Context) error {
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
	var result models.PlayerEventTrade
	query := e.db.QueryContext(models.PlayerEventTrade{}, c)
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
		event := fmt.Sprintf("Deleted [PlayerEventTrade] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getPlayerEventTradesBulk godoc
// @Id getPlayerEventTradesBulk
// @Summary Gets PlayerEventTrades in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PlayerEventTrade
// @Success 200 {array} models.PlayerEventTrade
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_trades/bulk [post]
func (e *PlayerEventTradeController) getPlayerEventTradesBulk(c echo.Context) error {
	var results []models.PlayerEventTrade

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

	err := e.db.QueryContext(models.PlayerEventTrade{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventTradesCount godoc
// @Id getPlayerEventTradesCount
// @Summary Counts PlayerEventTrades
// @Accept json
// @Produce json
// @Tags PlayerEventTrade
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventTrade
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_trades/count [get]
func (e *PlayerEventTradeController) getPlayerEventTradesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.PlayerEventTrade{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}