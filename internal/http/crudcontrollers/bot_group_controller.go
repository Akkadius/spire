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

type BotGroupController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotGroupController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotGroupController {
	return &BotGroupController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotGroupController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_group/:groupsIndex", e.getBotGroup, nil),
		routes.RegisterRoute(http.MethodGet, "bot_groups", e.listBotGroups, nil),
		routes.RegisterRoute(http.MethodGet, "bot_groups/count", e.getBotGroupsCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_group", e.createBotGroup, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_group/:groupsIndex", e.deleteBotGroup, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_group/:groupsIndex", e.updateBotGroup, nil),
		routes.RegisterRoute(http.MethodPost, "bot_groups/bulk", e.getBotGroupsBulk, nil),
	}
}

// listBotGroups godoc
// @Id listBotGroups
// @Summary Lists BotGroups
// @Accept json
// @Produce json
// @Tags BotGroup
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGroup
// @Failure 500 {string} string "Bad query request"
// @Router /bot_groups [get]
func (e *BotGroupController) listBotGroups(c echo.Context) error {
	var results []models.BotGroup
	err := e.db.QueryContext(models.BotGroup{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotGroup godoc
// @Id getBotGroup
// @Summary Gets BotGroup
// @Accept json
// @Produce json
// @Tags BotGroup
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGroup
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_group/{id} [get]
func (e *BotGroupController) getBotGroup(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	groupsIndex, err := strconv.Atoi(c.Param("groupsIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GroupsIndex]"})
	}
	params = append(params, groupsIndex)
	keys = append(keys, "groups_index = ?")

	// query builder
	var result models.BotGroup
	query := e.db.QueryContext(models.BotGroup{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.GroupsIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotGroup godoc
// @Id updateBotGroup
// @Summary Updates BotGroup
// @Accept json
// @Produce json
// @Tags BotGroup
// @Param id path int true "Id"
// @Param bot_group body models.BotGroup true "BotGroup"
// @Success 200 {array} models.BotGroup
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_group/{id} [patch]
func (e *BotGroupController) updateBotGroup(c echo.Context) error {
	request := new(models.BotGroup)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	groupsIndex, err := strconv.Atoi(c.Param("groupsIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GroupsIndex]"})
	}
	params = append(params, groupsIndex)
	keys = append(keys, "groups_index = ?")

	// query builder
	var result models.BotGroup
	query := e.db.QueryContext(models.BotGroup{}, c)
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
		event := fmt.Sprintf("Updated [BotGroup] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotGroup godoc
// @Id createBotGroup
// @Summary Creates BotGroup
// @Accept json
// @Produce json
// @Param bot_group body models.BotGroup true "BotGroup"
// @Tags BotGroup
// @Success 200 {array} models.BotGroup
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_group [put]
func (e *BotGroupController) createBotGroup(c echo.Context) error {
	botGroup := new(models.BotGroup)
	if err := c.Bind(botGroup); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotGroup{}, c).Model(&models.BotGroup{}).Create(&botGroup).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotGroup{}, botGroup)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotGroup] [%v] data [%v]", botGroup.GroupsIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botGroup)
}

// deleteBotGroup godoc
// @Id deleteBotGroup
// @Summary Deletes BotGroup
// @Accept json
// @Produce json
// @Tags BotGroup
// @Param id path int true "groupsIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_group/{id} [delete]
func (e *BotGroupController) deleteBotGroup(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	groupsIndex, err := strconv.Atoi(c.Param("groupsIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, groupsIndex)
	keys = append(keys, "groups_index = ?")

	// query builder
	var result models.BotGroup
	query := e.db.QueryContext(models.BotGroup{}, c)
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
		event := fmt.Sprintf("Deleted [BotGroup] [%v] keys [%v]", result.GroupsIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotGroupsBulk godoc
// @Id getBotGroupsBulk
// @Summary Gets BotGroups in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotGroup
// @Success 200 {array} models.BotGroup
// @Failure 500 {string} string "Bad query request"
// @Router /bot_groups/bulk [post]
func (e *BotGroupController) getBotGroupsBulk(c echo.Context) error {
	var results []models.BotGroup

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

	err := e.db.QueryContext(models.BotGroup{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotGroupsCount godoc
// @Id getBotGroupsCount
// @Summary Counts BotGroups
// @Accept json
// @Produce json
// @Tags BotGroup
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotGroup
// @Failure 500 {string} string "Bad query request"
// @Router /bot_groups/count [get]
func (e *BotGroupController) getBotGroupsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotGroup{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}