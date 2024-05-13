package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type BotGuildMemberController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewBotGuildMemberController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *BotGuildMemberController {
	return &BotGuildMemberController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *BotGuildMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_guild_member/:botId", e.getBotGuildMember, nil),
		routes.RegisterRoute(http.MethodGet, "bot_guild_members", e.listBotGuildMembers, nil),
		routes.RegisterRoute(http.MethodGet, "bot_guild_members/count", e.getBotGuildMembersCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_guild_member", e.createBotGuildMember, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_guild_member/:botId", e.deleteBotGuildMember, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_guild_member/:botId", e.updateBotGuildMember, nil),
		routes.RegisterRoute(http.MethodPost, "bot_guild_members/bulk", e.getBotGuildMembersBulk, nil),
	}
}

// listBotGuildMembers godoc
// @Id listBotGuildMembers
// @Summary Lists BotGuildMembers
// @Accept json
// @Produce json
// @Tags BotGuildMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGuildMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_guild_members [get]
func (e *BotGuildMemberController) listBotGuildMembers(c echo.Context) error {
	var results []models.BotGuildMember
	err := e.db.QueryContext(models.BotGuildMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotGuildMember godoc
// @Id getBotGuildMember
// @Summary Gets BotGuildMember
// @Accept json
// @Produce json
// @Tags BotGuildMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGuildMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_guild_member/{id} [get]
func (e *BotGuildMemberController) getBotGuildMember(c echo.Context) error {
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
	var result models.BotGuildMember
	query := e.db.QueryContext(models.BotGuildMember{}, c)
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

// updateBotGuildMember godoc
// @Id updateBotGuildMember
// @Summary Updates BotGuildMember
// @Accept json
// @Produce json
// @Tags BotGuildMember
// @Param id path int true "Id"
// @Param bot_guild_member body models.BotGuildMember true "BotGuildMember"
// @Success 200 {array} models.BotGuildMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_guild_member/{id} [patch]
func (e *BotGuildMemberController) updateBotGuildMember(c echo.Context) error {
	request := new(models.BotGuildMember)
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
	var result models.BotGuildMember
	query := e.db.QueryContext(models.BotGuildMember{}, c)
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
		event := fmt.Sprintf("Updated [BotGuildMember] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotGuildMember godoc
// @Id createBotGuildMember
// @Summary Creates BotGuildMember
// @Accept json
// @Produce json
// @Param bot_guild_member body models.BotGuildMember true "BotGuildMember"
// @Tags BotGuildMember
// @Success 200 {array} models.BotGuildMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_guild_member [put]
func (e *BotGuildMemberController) createBotGuildMember(c echo.Context) error {
	botGuildMember := new(models.BotGuildMember)
	if err := c.Bind(botGuildMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotGuildMember{}, c).Model(&models.BotGuildMember{}).Create(&botGuildMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotGuildMember{}, botGuildMember)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotGuildMember] [%v] data [%v]", botGuildMember.BotId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botGuildMember)
}

// deleteBotGuildMember godoc
// @Id deleteBotGuildMember
// @Summary Deletes BotGuildMember
// @Accept json
// @Produce json
// @Tags BotGuildMember
// @Param id path int true "botId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_guild_member/{id} [delete]
func (e *BotGuildMemberController) deleteBotGuildMember(c echo.Context) error {
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
	var result models.BotGuildMember
	query := e.db.QueryContext(models.BotGuildMember{}, c)
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
		event := fmt.Sprintf("Deleted [BotGuildMember] [%v] keys [%v]", result.BotId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotGuildMembersBulk godoc
// @Id getBotGuildMembersBulk
// @Summary Gets BotGuildMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotGuildMember
// @Success 200 {array} models.BotGuildMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_guild_members/bulk [post]
func (e *BotGuildMemberController) getBotGuildMembersBulk(c echo.Context) error {
	var results []models.BotGuildMember

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

	err := e.db.QueryContext(models.BotGuildMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotGuildMembersCount godoc
// @Id getBotGuildMembersCount
// @Summary Counts BotGuildMembers
// @Accept json
// @Produce json
// @Tags BotGuildMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGuildMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_guild_members/count [get]
func (e *BotGuildMemberController) getBotGuildMembersCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotGuildMember{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}
