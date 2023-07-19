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

type BotInspectMessageController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotInspectMessageController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotInspectMessageController {
	return &BotInspectMessageController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotInspectMessageController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_inspect_message/:botId", e.getBotInspectMessage, nil),
		routes.RegisterRoute(http.MethodGet, "bot_inspect_messages", e.listBotInspectMessages, nil),
		routes.RegisterRoute(http.MethodGet, "bot_inspect_messages/count", e.getBotInspectMessagesCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_inspect_message", e.createBotInspectMessage, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_inspect_message/:botId", e.deleteBotInspectMessage, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_inspect_message/:botId", e.updateBotInspectMessage, nil),
		routes.RegisterRoute(http.MethodPost, "bot_inspect_messages/bulk", e.getBotInspectMessagesBulk, nil),
	}
}

// listBotInspectMessages godoc
// @Id listBotInspectMessages
// @Summary Lists BotInspectMessages
// @Accept json
// @Produce json
// @Tags BotInspectMessage
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotInspectMessage
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inspect_messages [get]
func (e *BotInspectMessageController) listBotInspectMessages(c echo.Context) error {
	var results []models.BotInspectMessage
	err := e.db.QueryContext(models.BotInspectMessage{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotInspectMessage godoc
// @Id getBotInspectMessage
// @Summary Gets BotInspectMessage
// @Accept json
// @Produce json
// @Tags BotInspectMessage
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotInspectMessage
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inspect_message/{id} [get]
func (e *BotInspectMessageController) getBotInspectMessage(c echo.Context) error {
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
	var result models.BotInspectMessage
	query := e.db.QueryContext(models.BotInspectMessage{}, c)
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

// updateBotInspectMessage godoc
// @Id updateBotInspectMessage
// @Summary Updates BotInspectMessage
// @Accept json
// @Produce json
// @Tags BotInspectMessage
// @Param id path int true "Id"
// @Param bot_inspect_message body models.BotInspectMessage true "BotInspectMessage"
// @Success 200 {array} models.BotInspectMessage
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_inspect_message/{id} [patch]
func (e *BotInspectMessageController) updateBotInspectMessage(c echo.Context) error {
	request := new(models.BotInspectMessage)
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
	var result models.BotInspectMessage
	query := e.db.QueryContext(models.BotInspectMessage{}, c)
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
		event := fmt.Sprintf("Updated [BotInspectMessage] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotInspectMessage godoc
// @Id createBotInspectMessage
// @Summary Creates BotInspectMessage
// @Accept json
// @Produce json
// @Param bot_inspect_message body models.BotInspectMessage true "BotInspectMessage"
// @Tags BotInspectMessage
// @Success 200 {array} models.BotInspectMessage
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_inspect_message [put]
func (e *BotInspectMessageController) createBotInspectMessage(c echo.Context) error {
	botInspectMessage := new(models.BotInspectMessage)
	if err := c.Bind(botInspectMessage); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotInspectMessage{}, c).Model(&models.BotInspectMessage{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
        db = db.Omit(clause.Associations)
    }

	err := db.Create(&botInspectMessage).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotInspectMessage{}, botInspectMessage)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotInspectMessage] [%v] data [%v]", botInspectMessage.BotId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botInspectMessage)
}

// deleteBotInspectMessage godoc
// @Id deleteBotInspectMessage
// @Summary Deletes BotInspectMessage
// @Accept json
// @Produce json
// @Tags BotInspectMessage
// @Param id path int true "botId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_inspect_message/{id} [delete]
func (e *BotInspectMessageController) deleteBotInspectMessage(c echo.Context) error {
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
	var result models.BotInspectMessage
	query := e.db.QueryContext(models.BotInspectMessage{}, c)
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
		event := fmt.Sprintf("Deleted [BotInspectMessage] [%v] keys [%v]", result.BotId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotInspectMessagesBulk godoc
// @Id getBotInspectMessagesBulk
// @Summary Gets BotInspectMessages in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotInspectMessage
// @Success 200 {array} models.BotInspectMessage
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inspect_messages/bulk [post]
func (e *BotInspectMessageController) getBotInspectMessagesBulk(c echo.Context) error {
	var results []models.BotInspectMessage

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

	err := e.db.QueryContext(models.BotInspectMessage{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotInspectMessagesCount godoc
// @Id getBotInspectMessagesCount
// @Summary Counts BotInspectMessages
// @Accept json
// @Produce json
// @Tags BotInspectMessage
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotInspectMessage
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inspect_messages/count [get]
func (e *BotInspectMessageController) getBotInspectMessagesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotInspectMessage{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}