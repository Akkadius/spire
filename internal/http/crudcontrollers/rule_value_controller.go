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

type RuleValueController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewRuleValueController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *RuleValueController {
	return &RuleValueController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *RuleValueController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "rule_value/:rulesetId", e.getRuleValue, nil),
		routes.RegisterRoute(http.MethodGet, "rule_values", e.listRuleValues, nil),
		routes.RegisterRoute(http.MethodGet, "rule_values/count", e.getRuleValuesCount, nil),
		routes.RegisterRoute(http.MethodPut, "rule_value", e.createRuleValue, nil),
		routes.RegisterRoute(http.MethodDelete, "rule_value/:rulesetId", e.deleteRuleValue, nil),
		routes.RegisterRoute(http.MethodPatch, "rule_value/:rulesetId", e.updateRuleValue, nil),
		routes.RegisterRoute(http.MethodPost, "rule_values/bulk", e.getRuleValuesBulk, nil),
	}
}

// listRuleValues godoc
// @Id listRuleValues
// @Summary Lists RuleValues
// @Accept json
// @Produce json
// @Tags RuleValue
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RuleValue
// @Failure 500 {string} string "Bad query request"
// @Router /rule_values [get]
func (e *RuleValueController) listRuleValues(c echo.Context) error {
	var results []models.RuleValue
	err := e.db.QueryContext(models.RuleValue{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getRuleValue godoc
// @Id getRuleValue
// @Summary Gets RuleValue
// @Accept json
// @Produce json
// @Tags RuleValue
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RuleValue
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /rule_value/{id} [get]
func (e *RuleValueController) getRuleValue(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rulesetId, err := strconv.Atoi(c.Param("rulesetId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [RulesetId]"})
	}
	params = append(params, rulesetId)
	keys = append(keys, "ruleset_id = ?")

	// key param [rule_name] position [2] type [varchar]
	if len(c.QueryParam("rule_name")) > 0 {
		ruleNameParam, err := strconv.Atoi(c.QueryParam("rule_name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [rule_name] err [%s]", err.Error())})
		}

		params = append(params, ruleNameParam)
		keys = append(keys, "rule_name = ?")
	}

	// query builder
	var result models.RuleValue
	query := e.db.QueryContext(models.RuleValue{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.RulesetId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateRuleValue godoc
// @Id updateRuleValue
// @Summary Updates RuleValue
// @Accept json
// @Produce json
// @Tags RuleValue
// @Param id path int true "Id"
// @Param rule_value body models.RuleValue true "RuleValue"
// @Success 200 {array} models.RuleValue
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /rule_value/{id} [patch]
func (e *RuleValueController) updateRuleValue(c echo.Context) error {
	request := new(models.RuleValue)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	rulesetId, err := strconv.Atoi(c.Param("rulesetId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [RulesetId]"})
	}
	params = append(params, rulesetId)
	keys = append(keys, "ruleset_id = ?")

	// key param [rule_name] position [2] type [varchar]
	if len(c.QueryParam("rule_name")) > 0 {
		ruleNameParam, err := strconv.Atoi(c.QueryParam("rule_name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [rule_name] err [%s]", err.Error())})
		}

		params = append(params, ruleNameParam)
		keys = append(keys, "rule_name = ?")
	}

	// query builder
	var result models.RuleValue
	query := e.db.QueryContext(models.RuleValue{}, c)
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
		event := fmt.Sprintf("Updated [RuleValue] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createRuleValue godoc
// @Id createRuleValue
// @Summary Creates RuleValue
// @Accept json
// @Produce json
// @Param rule_value body models.RuleValue true "RuleValue"
// @Tags RuleValue
// @Success 200 {array} models.RuleValue
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /rule_value [put]
func (e *RuleValueController) createRuleValue(c echo.Context) error {
	ruleValue := new(models.RuleValue)
	if err := c.Bind(ruleValue); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.RuleValue{}, c).Model(&models.RuleValue{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&ruleValue).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.RuleValue{}, ruleValue)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [RuleValue] [%v] data [%v]", ruleValue.RulesetId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, ruleValue)
}

// deleteRuleValue godoc
// @Id deleteRuleValue
// @Summary Deletes RuleValue
// @Accept json
// @Produce json
// @Tags RuleValue
// @Param id path int true "rulesetId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /rule_value/{id} [delete]
func (e *RuleValueController) deleteRuleValue(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rulesetId, err := strconv.Atoi(c.Param("rulesetId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, rulesetId)
	keys = append(keys, "ruleset_id = ?")

	// key param [rule_name] position [2] type [varchar]
	if len(c.QueryParam("rule_name")) > 0 {
		ruleNameParam, err := strconv.Atoi(c.QueryParam("rule_name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [rule_name] err [%s]", err.Error())})
		}

		params = append(params, ruleNameParam)
		keys = append(keys, "rule_name = ?")
	}

	// query builder
	var result models.RuleValue
	query := e.db.QueryContext(models.RuleValue{}, c)
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
		event := fmt.Sprintf("Deleted [RuleValue] [%v] keys [%v]", result.RulesetId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getRuleValuesBulk godoc
// @Id getRuleValuesBulk
// @Summary Gets RuleValues in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags RuleValue
// @Success 200 {array} models.RuleValue
// @Failure 500 {string} string "Bad query request"
// @Router /rule_values/bulk [post]
func (e *RuleValueController) getRuleValuesBulk(c echo.Context) error {
	var results []models.RuleValue

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

	err := e.db.QueryContext(models.RuleValue{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getRuleValuesCount godoc
// @Id getRuleValuesCount
// @Summary Counts RuleValues
// @Accept json
// @Produce json
// @Tags RuleValue
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RuleValue
// @Failure 500 {string} string "Bad query request"
// @Router /rule_values/count [get]
func (e *RuleValueController) getRuleValuesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.RuleValue{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}