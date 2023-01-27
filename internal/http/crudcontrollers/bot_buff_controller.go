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

type BotBuffController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotBuffController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotBuffController {
	return &BotBuffController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotBuffController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_buff/:buffsIndex", e.getBotBuff, nil),
		routes.RegisterRoute(http.MethodGet, "bot_buffs", e.listBotBuffs, nil),
		routes.RegisterRoute(http.MethodGet, "bot_buffs/count", e.getBotBuffsCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_buff", e.createBotBuff, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_buff/:buffsIndex", e.deleteBotBuff, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_buff/:buffsIndex", e.updateBotBuff, nil),
		routes.RegisterRoute(http.MethodPost, "bot_buffs/bulk", e.getBotBuffsBulk, nil),
	}
}

// listBotBuffs godoc
// @Id listBotBuffs
// @Summary Lists BotBuffs
// @Accept json
// @Produce json
// @Tags BotBuff
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotBuff
// @Failure 500 {string} string "Bad query request"
// @Router /bot_buffs [get]
func (e *BotBuffController) listBotBuffs(c echo.Context) error {
	var results []models.BotBuff
	err := e.db.QueryContext(models.BotBuff{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotBuff godoc
// @Id getBotBuff
// @Summary Gets BotBuff
// @Accept json
// @Produce json
// @Tags BotBuff
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotBuff
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_buff/{id} [get]
func (e *BotBuffController) getBotBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	buffsIndex, err := strconv.Atoi(c.Param("buffsIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [BuffsIndex]"})
	}
	params = append(params, buffsIndex)
	keys = append(keys, "buffs_index = ?")

	// query builder
	var result models.BotBuff
	query := e.db.QueryContext(models.BotBuff{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.BuffsIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotBuff godoc
// @Id updateBotBuff
// @Summary Updates BotBuff
// @Accept json
// @Produce json
// @Tags BotBuff
// @Param id path int true "Id"
// @Param bot_buff body models.BotBuff true "BotBuff"
// @Success 200 {array} models.BotBuff
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_buff/{id} [patch]
func (e *BotBuffController) updateBotBuff(c echo.Context) error {
	request := new(models.BotBuff)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	buffsIndex, err := strconv.Atoi(c.Param("buffsIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [BuffsIndex]"})
	}
	params = append(params, buffsIndex)
	keys = append(keys, "buffs_index = ?")

	// query builder
	var result models.BotBuff
	query := e.db.QueryContext(models.BotBuff{}, c)
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
		event := fmt.Sprintf("Updated [BotBuff] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotBuff godoc
// @Id createBotBuff
// @Summary Creates BotBuff
// @Accept json
// @Produce json
// @Param bot_buff body models.BotBuff true "BotBuff"
// @Tags BotBuff
// @Success 200 {array} models.BotBuff
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_buff [put]
func (e *BotBuffController) createBotBuff(c echo.Context) error {
	botBuff := new(models.BotBuff)
	if err := c.Bind(botBuff); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotBuff{}, c).Model(&models.BotBuff{}).Create(&botBuff).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotBuff{}, botBuff)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotBuff] [%v] data [%v]", botBuff.BuffsIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botBuff)
}

// deleteBotBuff godoc
// @Id deleteBotBuff
// @Summary Deletes BotBuff
// @Accept json
// @Produce json
// @Tags BotBuff
// @Param id path int true "buffsIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_buff/{id} [delete]
func (e *BotBuffController) deleteBotBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	buffsIndex, err := strconv.Atoi(c.Param("buffsIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, buffsIndex)
	keys = append(keys, "buffs_index = ?")

	// query builder
	var result models.BotBuff
	query := e.db.QueryContext(models.BotBuff{}, c)
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
		event := fmt.Sprintf("Deleted [BotBuff] [%v] keys [%v]", result.BuffsIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotBuffsBulk godoc
// @Id getBotBuffsBulk
// @Summary Gets BotBuffs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotBuff
// @Success 200 {array} models.BotBuff
// @Failure 500 {string} string "Bad query request"
// @Router /bot_buffs/bulk [post]
func (e *BotBuffController) getBotBuffsBulk(c echo.Context) error {
	var results []models.BotBuff

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

	err := e.db.QueryContext(models.BotBuff{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotBuffsCount godoc
// @Id getBotBuffsCount
// @Summary Counts BotBuffs
// @Accept json
// @Produce json
// @Tags BotBuff
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotBuff
// @Failure 500 {string} string "Bad query request"
// @Router /bot_buffs/count [get]
func (e *BotBuffController) getBotBuffsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotBuff{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}