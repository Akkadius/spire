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

type BotCreateCombinationController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotCreateCombinationController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotCreateCombinationController {
	return &BotCreateCombinationController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotCreateCombinationController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_create_combination/:race", e.getBotCreateCombination, nil),
		routes.RegisterRoute(http.MethodGet, "bot_create_combinations", e.listBotCreateCombinations, nil),
		routes.RegisterRoute(http.MethodGet, "bot_create_combinations/count", e.getBotCreateCombinationsCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_create_combination", e.createBotCreateCombination, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_create_combination/:race", e.deleteBotCreateCombination, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_create_combination/:race", e.updateBotCreateCombination, nil),
		routes.RegisterRoute(http.MethodPost, "bot_create_combinations/bulk", e.getBotCreateCombinationsBulk, nil),
	}
}

// listBotCreateCombinations godoc
// @Id listBotCreateCombinations
// @Summary Lists BotCreateCombinations
// @Accept json
// @Produce json
// @Tags BotCreateCombination
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotCreateCombination
// @Failure 500 {string} string "Bad query request"
// @Router /bot_create_combinations [get]
func (e *BotCreateCombinationController) listBotCreateCombinations(c echo.Context) error {
	var results []models.BotCreateCombination
	err := e.db.QueryContext(models.BotCreateCombination{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotCreateCombination godoc
// @Id getBotCreateCombination
// @Summary Gets BotCreateCombination
// @Accept json
// @Produce json
// @Tags BotCreateCombination
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotCreateCombination
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_create_combination/{id} [get]
func (e *BotCreateCombinationController) getBotCreateCombination(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	race, err := strconv.Atoi(c.Param("race"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Race]"})
	}
	params = append(params, race)
	keys = append(keys, "race = ?")

	// query builder
	var result models.BotCreateCombination
	query := e.db.QueryContext(models.BotCreateCombination{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Race == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotCreateCombination godoc
// @Id updateBotCreateCombination
// @Summary Updates BotCreateCombination
// @Accept json
// @Produce json
// @Tags BotCreateCombination
// @Param id path int true "Id"
// @Param bot_create_combination body models.BotCreateCombination true "BotCreateCombination"
// @Success 200 {array} models.BotCreateCombination
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_create_combination/{id} [patch]
func (e *BotCreateCombinationController) updateBotCreateCombination(c echo.Context) error {
	request := new(models.BotCreateCombination)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	race, err := strconv.Atoi(c.Param("race"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Race]"})
	}
	params = append(params, race)
	keys = append(keys, "race = ?")

	// query builder
	var result models.BotCreateCombination
	query := e.db.QueryContext(models.BotCreateCombination{}, c)
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
		event := fmt.Sprintf("Updated [BotCreateCombination] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotCreateCombination godoc
// @Id createBotCreateCombination
// @Summary Creates BotCreateCombination
// @Accept json
// @Produce json
// @Param bot_create_combination body models.BotCreateCombination true "BotCreateCombination"
// @Tags BotCreateCombination
// @Success 200 {array} models.BotCreateCombination
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_create_combination [put]
func (e *BotCreateCombinationController) createBotCreateCombination(c echo.Context) error {
	botCreateCombination := new(models.BotCreateCombination)
	if err := c.Bind(botCreateCombination); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotCreateCombination{}, c).Model(&models.BotCreateCombination{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&botCreateCombination).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotCreateCombination{}, botCreateCombination)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotCreateCombination] [%v] data [%v]", botCreateCombination.Race, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botCreateCombination)
}

// deleteBotCreateCombination godoc
// @Id deleteBotCreateCombination
// @Summary Deletes BotCreateCombination
// @Accept json
// @Produce json
// @Tags BotCreateCombination
// @Param id path int true "race"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_create_combination/{id} [delete]
func (e *BotCreateCombinationController) deleteBotCreateCombination(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	race, err := strconv.Atoi(c.Param("race"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, race)
	keys = append(keys, "race = ?")

	// query builder
	var result models.BotCreateCombination
	query := e.db.QueryContext(models.BotCreateCombination{}, c)
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
		event := fmt.Sprintf("Deleted [BotCreateCombination] [%v] keys [%v]", result.Race, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotCreateCombinationsBulk godoc
// @Id getBotCreateCombinationsBulk
// @Summary Gets BotCreateCombinations in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotCreateCombination
// @Success 200 {array} models.BotCreateCombination
// @Failure 500 {string} string "Bad query request"
// @Router /bot_create_combinations/bulk [post]
func (e *BotCreateCombinationController) getBotCreateCombinationsBulk(c echo.Context) error {
	var results []models.BotCreateCombination

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

	err := e.db.QueryContext(models.BotCreateCombination{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotCreateCombinationsCount godoc
// @Id getBotCreateCombinationsCount
// @Summary Counts BotCreateCombinations
// @Accept json
// @Produce json
// @Tags BotCreateCombination
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotCreateCombination
// @Failure 500 {string} string "Bad query request"
// @Router /bot_create_combinations/count [get]
func (e *BotCreateCombinationController) getBotCreateCombinationsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotCreateCombination{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}