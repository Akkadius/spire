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

type BotTimerController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotTimerController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotTimerController {
	return &BotTimerController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotTimerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_timer/:botId", e.getBotTimer, nil),
		routes.RegisterRoute(http.MethodGet, "bot_timers", e.listBotTimers, nil),
		routes.RegisterRoute(http.MethodGet, "bot_timers/count", e.getBotTimersCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_timer", e.createBotTimer, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_timer/:botId", e.deleteBotTimer, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_timer/:botId", e.updateBotTimer, nil),
		routes.RegisterRoute(http.MethodPost, "bot_timers/bulk", e.getBotTimersBulk, nil),
	}
}

// listBotTimers godoc
// @Id listBotTimers
// @Summary Lists BotTimers
// @Accept json
// @Produce json
// @Tags BotTimer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotTimer
// @Failure 500 {string} string "Bad query request"
// @Router /bot_timers [get]
func (e *BotTimerController) listBotTimers(c echo.Context) error {
	var results []models.BotTimer
	err := e.db.QueryContext(models.BotTimer{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotTimer godoc
// @Id getBotTimer
// @Summary Gets BotTimer
// @Accept json
// @Produce json
// @Tags BotTimer
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotTimer
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_timer/{id} [get]
func (e *BotTimerController) getBotTimer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [BotId]"})
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// query builder
	var result models.BotTimer
	query := e.db.QueryContext(models.BotTimer{}, c)
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

// updateBotTimer godoc
// @Id updateBotTimer
// @Summary Updates BotTimer
// @Accept json
// @Produce json
// @Tags BotTimer
// @Param id path int true "Id"
// @Param bot_timer body models.BotTimer true "BotTimer"
// @Success 200 {array} models.BotTimer
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_timer/{id} [patch]
func (e *BotTimerController) updateBotTimer(c echo.Context) error {
	request := new(models.BotTimer)
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

	// query builder
	var result models.BotTimer
	query := e.db.QueryContext(models.BotTimer{}, c)
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
		event := fmt.Sprintf("Updated [BotTimer] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotTimer godoc
// @Id createBotTimer
// @Summary Creates BotTimer
// @Accept json
// @Produce json
// @Param bot_timer body models.BotTimer true "BotTimer"
// @Tags BotTimer
// @Success 200 {array} models.BotTimer
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_timer [put]
func (e *BotTimerController) createBotTimer(c echo.Context) error {
	botTimer := new(models.BotTimer)
	if err := c.Bind(botTimer); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotTimer{}, c).Model(&models.BotTimer{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&botTimer).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotTimer{}, botTimer)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotTimer] [%v] data [%v]", botTimer.BotId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botTimer)
}

// deleteBotTimer godoc
// @Id deleteBotTimer
// @Summary Deletes BotTimer
// @Accept json
// @Produce json
// @Tags BotTimer
// @Param id path int true "botId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_timer/{id} [delete]
func (e *BotTimerController) deleteBotTimer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// query builder
	var result models.BotTimer
	query := e.db.QueryContext(models.BotTimer{}, c)
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
		event := fmt.Sprintf("Deleted [BotTimer] [%v] keys [%v]", result.BotId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotTimersBulk godoc
// @Id getBotTimersBulk
// @Summary Gets BotTimers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotTimer
// @Success 200 {array} models.BotTimer
// @Failure 500 {string} string "Bad query request"
// @Router /bot_timers/bulk [post]
func (e *BotTimerController) getBotTimersBulk(c echo.Context) error {
	var results []models.BotTimer

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

	err := e.db.QueryContext(models.BotTimer{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotTimersCount godoc
// @Id getBotTimersCount
// @Summary Counts BotTimers
// @Accept json
// @Produce json
// @Tags BotTimer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotTimer
// @Failure 500 {string} string "Bad query request"
// @Router /bot_timers/count [get]
func (e *BotTimerController) getBotTimersCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotTimer{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}