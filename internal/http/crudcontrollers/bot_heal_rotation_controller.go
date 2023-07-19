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

type BotHealRotationController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotHealRotationController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotHealRotationController {
	return &BotHealRotationController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotHealRotationController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotation/:healRotationIndex", e.getBotHealRotation, nil),
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotations", e.listBotHealRotations, nil),
		routes.RegisterRoute(http.MethodGet, "bot_heal_rotations/count", e.getBotHealRotationsCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_heal_rotation", e.createBotHealRotation, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_heal_rotation/:healRotationIndex", e.deleteBotHealRotation, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_heal_rotation/:healRotationIndex", e.updateBotHealRotation, nil),
		routes.RegisterRoute(http.MethodPost, "bot_heal_rotations/bulk", e.getBotHealRotationsBulk, nil),
	}
}

// listBotHealRotations godoc
// @Id listBotHealRotations
// @Summary Lists BotHealRotations
// @Accept json
// @Produce json
// @Tags BotHealRotation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotation
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotations [get]
func (e *BotHealRotationController) listBotHealRotations(c echo.Context) error {
	var results []models.BotHealRotation
	err := e.db.QueryContext(models.BotHealRotation{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotHealRotation godoc
// @Id getBotHealRotation
// @Summary Gets BotHealRotation
// @Accept json
// @Produce json
// @Tags BotHealRotation
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotation
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotation/{id} [get]
func (e *BotHealRotationController) getBotHealRotation(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	healRotationIndex, err := strconv.Atoi(c.Param("healRotationIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [HealRotationIndex]"})
	}
	params = append(params, healRotationIndex)
	keys = append(keys, "heal_rotation_index = ?")

	// query builder
	var result models.BotHealRotation
	query := e.db.QueryContext(models.BotHealRotation{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.HealRotationIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotHealRotation godoc
// @Id updateBotHealRotation
// @Summary Updates BotHealRotation
// @Accept json
// @Produce json
// @Tags BotHealRotation
// @Param id path int true "Id"
// @Param bot_heal_rotation body models.BotHealRotation true "BotHealRotation"
// @Success 200 {array} models.BotHealRotation
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_heal_rotation/{id} [patch]
func (e *BotHealRotationController) updateBotHealRotation(c echo.Context) error {
	request := new(models.BotHealRotation)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	healRotationIndex, err := strconv.Atoi(c.Param("healRotationIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [HealRotationIndex]"})
	}
	params = append(params, healRotationIndex)
	keys = append(keys, "heal_rotation_index = ?")

	// query builder
	var result models.BotHealRotation
	query := e.db.QueryContext(models.BotHealRotation{}, c)
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
		event := fmt.Sprintf("Updated [BotHealRotation] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotHealRotation godoc
// @Id createBotHealRotation
// @Summary Creates BotHealRotation
// @Accept json
// @Produce json
// @Param bot_heal_rotation body models.BotHealRotation true "BotHealRotation"
// @Tags BotHealRotation
// @Success 200 {array} models.BotHealRotation
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_heal_rotation [put]
func (e *BotHealRotationController) createBotHealRotation(c echo.Context) error {
	botHealRotation := new(models.BotHealRotation)
	if err := c.Bind(botHealRotation); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotHealRotation{}, c).Model(&models.BotHealRotation{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&botHealRotation).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotHealRotation{}, botHealRotation)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotHealRotation] [%v] data [%v]", botHealRotation.HealRotationIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botHealRotation)
}

// deleteBotHealRotation godoc
// @Id deleteBotHealRotation
// @Summary Deletes BotHealRotation
// @Accept json
// @Produce json
// @Tags BotHealRotation
// @Param id path int true "healRotationIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_heal_rotation/{id} [delete]
func (e *BotHealRotationController) deleteBotHealRotation(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	healRotationIndex, err := strconv.Atoi(c.Param("healRotationIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, healRotationIndex)
	keys = append(keys, "heal_rotation_index = ?")

	// query builder
	var result models.BotHealRotation
	query := e.db.QueryContext(models.BotHealRotation{}, c)
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
		event := fmt.Sprintf("Deleted [BotHealRotation] [%v] keys [%v]", result.HealRotationIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotHealRotationsBulk godoc
// @Id getBotHealRotationsBulk
// @Summary Gets BotHealRotations in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotHealRotation
// @Success 200 {array} models.BotHealRotation
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotations/bulk [post]
func (e *BotHealRotationController) getBotHealRotationsBulk(c echo.Context) error {
	var results []models.BotHealRotation

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

	err := e.db.QueryContext(models.BotHealRotation{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotHealRotationsCount godoc
// @Id getBotHealRotationsCount
// @Summary Counts BotHealRotations
// @Accept json
// @Produce json
// @Tags BotHealRotation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotHealRotation
// @Failure 500 {string} string "Bad query request"
// @Router /bot_heal_rotations/count [get]
func (e *BotHealRotationController) getBotHealRotationsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotHealRotation{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}