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

type BotHealRotationMemberController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotHealRotationMemberController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotHealRotationMemberController {
	return &BotHealRotationMemberController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotHealRotationMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotation_member/:memberIndex", e.getBotHealRotationMember, nil),
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotation_members", e.listBotHealRotationMembers, nil),
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotation_members/count", e.getBotHealRotationMembersCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_heal_rotation_member", e.createBotHealRotationMember, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_heal_rotation_member/:memberIndex", e.deleteBotHealRotationMember, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_heal_rotation_member/:memberIndex", e.updateBotHealRotationMember, nil),
		routes.RegisterRoute(http.MethodPost, "bot_heal_rotation_members/bulk", e.getBotHealRotationMembersBulk, nil),
	}
}

// listBotHealRotationMembers godoc
// @Id listBotHealRotationMembers
// @Summary Lists BotHealRotationMembers
// @Accept json
// @Produce json
// @Tags BotHealRotationMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotationMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation_members [get]
func (e *BotHealRotationMemberController) listBotHealRotationMembers(c echo.Context) error {
	var results []models.BotHealRotationMember
	err := e.db.QueryContext(models.BotHealRotationMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotHealRotationMember godoc
// @Id getBotHealRotationMember
// @Summary Gets BotHealRotationMember
// @Accept json
// @Produce json
// @Tags BotHealRotationMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotationMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation_member/{id} [get]
func (e *BotHealRotationMemberController) getBotHealRotationMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	memberIndex, err := strconv.Atoi(c.Param("memberIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [MemberIndex]"})
	}
	params = append(params, memberIndex)
	keys = append(keys, "member_index = ?")

	// query builder
	var result models.BotHealRotationMember
	query := e.db.QueryContext(models.BotHealRotationMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.MemberIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotHealRotationMember godoc
// @Id updateBotHealRotationMember
// @Summary Updates BotHealRotationMember
// @Accept json
// @Produce json
// @Tags BotHealRotationMember
// @Param id path int true "Id"
// @Param bot_heal_rotation_member body models.BotHealRotationMember true "BotHealRotationMember"
// @Success 200 {array} models.BotHealRotationMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_heal_rotation_member/{id} [patch]
func (e *BotHealRotationMemberController) updateBotHealRotationMember(c echo.Context) error {
	request := new(models.BotHealRotationMember)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	memberIndex, err := strconv.Atoi(c.Param("memberIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [MemberIndex]"})
	}
	params = append(params, memberIndex)
	keys = append(keys, "member_index = ?")

	// query builder
	var result models.BotHealRotationMember
	query := e.db.QueryContext(models.BotHealRotationMember{}, c)
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
		event := fmt.Sprintf("Updated [BotHealRotationMember] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotHealRotationMember godoc
// @Id createBotHealRotationMember
// @Summary Creates BotHealRotationMember
// @Accept json
// @Produce json
// @Param bot_heal_rotation_member body models.BotHealRotationMember true "BotHealRotationMember"
// @Tags BotHealRotationMember
// @Success 200 {array} models.BotHealRotationMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_heal_rotation_member [put]
func (e *BotHealRotationMemberController) createBotHealRotationMember(c echo.Context) error {
	botHealRotationMember := new(models.BotHealRotationMember)
	if err := c.Bind(botHealRotationMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotHealRotationMember{}, c).Model(&models.BotHealRotationMember{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&botHealRotationMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotHealRotationMember{}, botHealRotationMember)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotHealRotationMember] [%v] data [%v]", botHealRotationMember.MemberIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botHealRotationMember)
}

// deleteBotHealRotationMember godoc
// @Id deleteBotHealRotationMember
// @Summary Deletes BotHealRotationMember
// @Accept json
// @Produce json
// @Tags BotHealRotationMember
// @Param id path int true "memberIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_heal_rotation_member/{id} [delete]
func (e *BotHealRotationMemberController) deleteBotHealRotationMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	memberIndex, err := strconv.Atoi(c.Param("memberIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, memberIndex)
	keys = append(keys, "member_index = ?")

	// query builder
	var result models.BotHealRotationMember
	query := e.db.QueryContext(models.BotHealRotationMember{}, c)
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
		event := fmt.Sprintf("Deleted [BotHealRotationMember] [%v] keys [%v]", result.MemberIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotHealRotationMembersBulk godoc
// @Id getBotHealRotationMembersBulk
// @Summary Gets BotHealRotationMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotHealRotationMember
// @Success 200 {array} models.BotHealRotationMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation_members/bulk [post]
func (e *BotHealRotationMemberController) getBotHealRotationMembersBulk(c echo.Context) error {
	var results []models.BotHealRotationMember

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

	err := e.db.QueryContext(models.BotHealRotationMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotHealRotationMembersCount godoc
// @Id getBotHealRotationMembersCount
// @Summary Counts BotHealRotationMembers
// @Accept json
// @Produce json
// @Tags BotHealRotationMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotationMember
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation_members/count [get]
func (e *BotHealRotationMemberController) getBotHealRotationMembersCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotHealRotationMember{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}