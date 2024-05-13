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

type BotGroupMemberController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewBotGroupMemberController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *BotGroupMemberController {
	return &BotGroupMemberController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *BotGroupMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_group_member/:groupMembersIndex", e.getBotGroupMember, nil),
		routes.RegisterRoute(http.MethodGet, "bot_group_members", e.listBotGroupMembers, nil),
		routes.RegisterRoute(http.MethodGet, "bot_group_members/count", e.getBotGroupMembersCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_group_member", e.createBotGroupMember, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_group_member/:groupMembersIndex", e.deleteBotGroupMember, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_group_member/:groupMembersIndex", e.updateBotGroupMember, nil),
		routes.RegisterRoute(http.MethodPost, "bot_group_members/bulk", e.getBotGroupMembersBulk, nil),
	}
}

// listBotGroupMembers godoc
// @Id listBotGroupMembers
// @Summary Lists BotGroupMembers
// @Accept json
// @Produce json
// @Tags BotGroupMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGroupMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_group_members [get]
func (e *BotGroupMemberController) listBotGroupMembers(c echo.Context) error {
	var results []models.BotGroupMember
	err := e.db.QueryContext(models.BotGroupMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotGroupMember godoc
// @Id getBotGroupMember
// @Summary Gets BotGroupMember
// @Accept json
// @Produce json
// @Tags BotGroupMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGroupMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_group_member/{id} [get]
func (e *BotGroupMemberController) getBotGroupMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	groupMembersIndex, err := strconv.Atoi(c.Param("groupMembersIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GroupMembersIndex]"})
	}
	params = append(params, groupMembersIndex)
	keys = append(keys, "group_members_index = ?")

	// query builder
	var result models.BotGroupMember
	query := e.db.QueryContext(models.BotGroupMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.GroupMembersIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotGroupMember godoc
// @Id updateBotGroupMember
// @Summary Updates BotGroupMember
// @Accept json
// @Produce json
// @Tags BotGroupMember
// @Param id path int true "Id"
// @Param bot_group_member body models.BotGroupMember true "BotGroupMember"
// @Success 200 {array} models.BotGroupMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_group_member/{id} [patch]
func (e *BotGroupMemberController) updateBotGroupMember(c echo.Context) error {
	request := new(models.BotGroupMember)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	groupMembersIndex, err := strconv.Atoi(c.Param("groupMembersIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GroupMembersIndex]"})
	}
	params = append(params, groupMembersIndex)
	keys = append(keys, "group_members_index = ?")

	// query builder
	var result models.BotGroupMember
	query := e.db.QueryContext(models.BotGroupMember{}, c)
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
		event := fmt.Sprintf("Updated [BotGroupMember] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotGroupMember godoc
// @Id createBotGroupMember
// @Summary Creates BotGroupMember
// @Accept json
// @Produce json
// @Param bot_group_member body models.BotGroupMember true "BotGroupMember"
// @Tags BotGroupMember
// @Success 200 {array} models.BotGroupMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_group_member [put]
func (e *BotGroupMemberController) createBotGroupMember(c echo.Context) error {
	botGroupMember := new(models.BotGroupMember)
	if err := c.Bind(botGroupMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotGroupMember{}, c).Model(&models.BotGroupMember{}).Create(&botGroupMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotGroupMember{}, botGroupMember)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotGroupMember] [%v] data [%v]", botGroupMember.GroupMembersIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botGroupMember)
}

// deleteBotGroupMember godoc
// @Id deleteBotGroupMember
// @Summary Deletes BotGroupMember
// @Accept json
// @Produce json
// @Tags BotGroupMember
// @Param id path int true "groupMembersIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_group_member/{id} [delete]
func (e *BotGroupMemberController) deleteBotGroupMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	groupMembersIndex, err := strconv.Atoi(c.Param("groupMembersIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GroupMembersIndex]"})
	}
	params = append(params, groupMembersIndex)
	keys = append(keys, "group_members_index = ?")

	// query builder
	var result models.BotGroupMember
	query := e.db.QueryContext(models.BotGroupMember{}, c)
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
		event := fmt.Sprintf("Deleted [BotGroupMember] [%v] keys [%v]", result.GroupMembersIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotGroupMembersBulk godoc
// @Id getBotGroupMembersBulk
// @Summary Gets BotGroupMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotGroupMember
// @Success 200 {array} models.BotGroupMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_group_members/bulk [post]
func (e *BotGroupMemberController) getBotGroupMembersBulk(c echo.Context) error {
	var results []models.BotGroupMember

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

	err := e.db.QueryContext(models.BotGroupMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotGroupMembersCount godoc
// @Id getBotGroupMembersCount
// @Summary Counts BotGroupMembers
// @Accept json
// @Produce json
// @Tags BotGroupMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGroupMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_group_members/count [get]
func (e *BotGroupMemberController) getBotGroupMembersCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotGroupMember{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}
