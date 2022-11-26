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

type BotHealRotationTargetController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotHealRotationTargetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotHealRotationTargetController {
	return &BotHealRotationTargetController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotHealRotationTargetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotation_target/:targetIndex", e.getBotHealRotationTarget, nil),
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotation_targets", e.listBotHealRotationTargets, nil),
		routes.RegisterRoute(http.MethodPut, "bot_heal_rotation_target", e.createBotHealRotationTarget, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_heal_rotation_target/:targetIndex", e.deleteBotHealRotationTarget, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_heal_rotation_target/:targetIndex", e.updateBotHealRotationTarget, nil),
		routes.RegisterRoute(http.MethodPost, "bot_heal_rotation_targets/bulk", e.getBotHealRotationTargetsBulk, nil),
	}
}

// listBotHealRotationTargets godoc
// @Id listBotHealRotationTargets
// @Summary Lists BotHealRotationTargets
// @Accept json
// @Produce json
// @Tags BotHealRotationTarget
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotationTarget
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation_targets [get]
func (e *BotHealRotationTargetController) listBotHealRotationTargets(c echo.Context) error {
	var results []models.BotHealRotationTarget
	err := e.db.QueryContext(models.BotHealRotationTarget{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotHealRotationTarget godoc
// @Id getBotHealRotationTarget
// @Summary Gets BotHealRotationTarget
// @Accept json
// @Produce json
// @Tags BotHealRotationTarget
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotationTarget
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation_target/{id} [get]
func (e *BotHealRotationTargetController) getBotHealRotationTarget(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	targetIndex, err := strconv.Atoi(c.Param("targetIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TargetIndex]"})
	}
	params = append(params, targetIndex)
	keys = append(keys, "target_index = ?")

	// query builder
	var result models.BotHealRotationTarget
	query := e.db.QueryContext(models.BotHealRotationTarget{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.TargetIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotHealRotationTarget godoc
// @Id updateBotHealRotationTarget
// @Summary Updates BotHealRotationTarget
// @Accept json
// @Produce json
// @Tags BotHealRotationTarget
// @Param id path int true "Id"
// @Param bot_heal_rotation_target body models.BotHealRotationTarget true "BotHealRotationTarget"
// @Success 200 {array} models.BotHealRotationTarget
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_heal_rotation_target/{id} [patch]
func (e *BotHealRotationTargetController) updateBotHealRotationTarget(c echo.Context) error {
	request := new(models.BotHealRotationTarget)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	targetIndex, err := strconv.Atoi(c.Param("targetIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TargetIndex]"})
	}
	params = append(params, targetIndex)
	keys = append(keys, "target_index = ?")

	// query builder
	var result models.BotHealRotationTarget
	query := e.db.QueryContext(models.BotHealRotationTarget{}, c)
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
		event := fmt.Sprintf("Updated [BotHealRotationTarget] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotHealRotationTarget godoc
// @Id createBotHealRotationTarget
// @Summary Creates BotHealRotationTarget
// @Accept json
// @Produce json
// @Param bot_heal_rotation_target body models.BotHealRotationTarget true "BotHealRotationTarget"
// @Tags BotHealRotationTarget
// @Success 200 {array} models.BotHealRotationTarget
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_heal_rotation_target [put]
func (e *BotHealRotationTargetController) createBotHealRotationTarget(c echo.Context) error {
	botHealRotationTarget := new(models.BotHealRotationTarget)
	if err := c.Bind(botHealRotationTarget); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotHealRotationTarget{}, c).Model(&models.BotHealRotationTarget{}).Create(&botHealRotationTarget).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotHealRotationTarget{}, botHealRotationTarget)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotHealRotationTarget] [%v] data [%v]", botHealRotationTarget.TargetIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botHealRotationTarget)
}

// deleteBotHealRotationTarget godoc
// @Id deleteBotHealRotationTarget
// @Summary Deletes BotHealRotationTarget
// @Accept json
// @Produce json
// @Tags BotHealRotationTarget
// @Param id path int true "targetIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_heal_rotation_target/{id} [delete]
func (e *BotHealRotationTargetController) deleteBotHealRotationTarget(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	targetIndex, err := strconv.Atoi(c.Param("targetIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, targetIndex)
	keys = append(keys, "target_index = ?")

	// query builder
	var result models.BotHealRotationTarget
	query := e.db.QueryContext(models.BotHealRotationTarget{}, c)
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
		event := fmt.Sprintf("Deleted [BotHealRotationTarget] [%v] keys [%v]", result.TargetIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotHealRotationTargetsBulk godoc
// @Id getBotHealRotationTargetsBulk
// @Summary Gets BotHealRotationTargets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotHealRotationTarget
// @Success 200 {array} models.BotHealRotationTarget
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation_targets/bulk [post]
func (e *BotHealRotationTargetController) getBotHealRotationTargetsBulk(c echo.Context) error {
	var results []models.BotHealRotationTarget

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

	err := e.db.QueryContext(models.BotHealRotationTarget{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
