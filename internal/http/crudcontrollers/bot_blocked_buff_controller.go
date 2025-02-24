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

type BotBlockedBuffController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewBotBlockedBuffController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *BotBlockedBuffController {
	return &BotBlockedBuffController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *BotBlockedBuffController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_blocked_buff/:botId", e.getBotBlockedBuff, nil),
		routes.RegisterRoute(http.MethodGet, "bot_blocked_buffs", e.listBotBlockedBuffs, nil),
		routes.RegisterRoute(http.MethodGet, "bot_blocked_buffs/count", e.getBotBlockedBuffsCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_blocked_buff", e.createBotBlockedBuff, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_blocked_buff/:botId", e.deleteBotBlockedBuff, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_blocked_buff/:botId", e.updateBotBlockedBuff, nil),
		routes.RegisterRoute(http.MethodPost, "bot_blocked_buffs/bulk", e.getBotBlockedBuffsBulk, nil),
	}
}

// listBotBlockedBuffs godoc
// @Id listBotBlockedBuffs
// @Summary Lists BotBlockedBuffs
// @Accept json
// @Produce json
// @Tags BotBlockedBuff
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotBlockedBuff
// @Failure 500 {string} string "Bad query request"
// @Router /bot_blocked_buffs [get]
func (e *BotBlockedBuffController) listBotBlockedBuffs(c echo.Context) error {
	var results []models.BotBlockedBuff
	err := e.db.QueryContext(models.BotBlockedBuff{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotBlockedBuff godoc
// @Id getBotBlockedBuff
// @Summary Gets BotBlockedBuff
// @Accept json
// @Produce json
// @Tags BotBlockedBuff
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotBlockedBuff
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_blocked_buff/{id} [get]
func (e *BotBlockedBuffController) getBotBlockedBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [BotId]"})
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// key param [spell_id] position [2] type [int]
	if len(c.QueryParam("spell_id")) > 0 {
		spellIdParam, err := strconv.Atoi(c.QueryParam("spell_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [spell_id] err [%s]", err.Error())})
		}

		params = append(params, spellIdParam)
		keys = append(keys, "spell_id = ?")
	}

	// query builder
	var result models.BotBlockedBuff
	query := e.db.QueryContext(models.BotBlockedBuff{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.BotId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotBlockedBuff godoc
// @Id updateBotBlockedBuff
// @Summary Updates BotBlockedBuff
// @Accept json
// @Produce json
// @Tags BotBlockedBuff
// @Param id path int true "Id"
// @Param bot_blocked_buff body models.BotBlockedBuff true "BotBlockedBuff"
// @Success 200 {array} models.BotBlockedBuff
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_blocked_buff/{id} [patch]
func (e *BotBlockedBuffController) updateBotBlockedBuff(c echo.Context) error {
	request := new(models.BotBlockedBuff)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [BotId]"})
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// key param [spell_id] position [2] type [int]
	if len(c.QueryParam("spell_id")) > 0 {
		spellIdParam, err := strconv.Atoi(c.QueryParam("spell_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [spell_id] err [%s]", err.Error())})
		}

		params = append(params, spellIdParam)
		keys = append(keys, "spell_id = ?")
	}

	// query builder
	var result models.BotBlockedBuff
	query := e.db.QueryContext(models.BotBlockedBuff{}, c)
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
		event := fmt.Sprintf("Updated [BotBlockedBuff] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotBlockedBuff godoc
// @Id createBotBlockedBuff
// @Summary Creates BotBlockedBuff
// @Accept json
// @Produce json
// @Param bot_blocked_buff body models.BotBlockedBuff true "BotBlockedBuff"
// @Tags BotBlockedBuff
// @Success 200 {array} models.BotBlockedBuff
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_blocked_buff [put]
func (e *BotBlockedBuffController) createBotBlockedBuff(c echo.Context) error {
	botBlockedBuff := new(models.BotBlockedBuff)
	if err := c.Bind(botBlockedBuff); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotBlockedBuff{}, c).Model(&models.BotBlockedBuff{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&botBlockedBuff).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotBlockedBuff{}, botBlockedBuff)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotBlockedBuff] [%v] data [%v]", botBlockedBuff.BotId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botBlockedBuff)
}

// deleteBotBlockedBuff godoc
// @Id deleteBotBlockedBuff
// @Summary Deletes BotBlockedBuff
// @Accept json
// @Produce json
// @Tags BotBlockedBuff
// @Param id path int true "botId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_blocked_buff/{id} [delete]
func (e *BotBlockedBuffController) deleteBotBlockedBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// key param [spell_id] position [2] type [int]
	if len(c.QueryParam("spell_id")) > 0 {
		spellIdParam, err := strconv.Atoi(c.QueryParam("spell_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [spell_id] err [%s]", err.Error())})
		}

		params = append(params, spellIdParam)
		keys = append(keys, "spell_id = ?")
	}

	// query builder
	var result models.BotBlockedBuff
	query := e.db.QueryContext(models.BotBlockedBuff{}, c)
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
		event := fmt.Sprintf("Deleted [BotBlockedBuff] [%v] keys [%v]", result.BotId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotBlockedBuffsBulk godoc
// @Id getBotBlockedBuffsBulk
// @Summary Gets BotBlockedBuffs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotBlockedBuff
// @Success 200 {array} models.BotBlockedBuff
// @Failure 500 {string} string "Bad query request"
// @Router /bot_blocked_buffs/bulk [post]
func (e *BotBlockedBuffController) getBotBlockedBuffsBulk(c echo.Context) error {
	var results []models.BotBlockedBuff

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

	err := e.db.QueryContext(models.BotBlockedBuff{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotBlockedBuffsCount godoc
// @Id getBotBlockedBuffsCount
// @Summary Counts BotBlockedBuffs
// @Accept json
// @Produce json
// @Tags BotBlockedBuff
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotBlockedBuff
// @Failure 500 {string} string "Bad query request"
// @Router /bot_blocked_buffs/count [get]
func (e *BotBlockedBuffController) getBotBlockedBuffsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotBlockedBuff{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}