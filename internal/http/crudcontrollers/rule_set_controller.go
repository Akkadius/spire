package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RuleSetController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewRuleSetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *RuleSetController {
	return &RuleSetController{
		db:	    db,
		logger: logger,
	}
}

func (e *RuleSetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "rule_set/:rulesetId", e.getRuleSet, nil),
		routes.RegisterRoute(http.MethodGet, "rule_sets", e.listRuleSets, nil),
		routes.RegisterRoute(http.MethodPut, "rule_set", e.createRuleSet, nil),
		routes.RegisterRoute(http.MethodDelete, "rule_set/:rulesetId", e.deleteRuleSet, nil),
		routes.RegisterRoute(http.MethodPatch, "rule_set/:rulesetId", e.updateRuleSet, nil),
		routes.RegisterRoute(http.MethodPost, "rule_sets/bulk", e.getRuleSetsBulk, nil),
	}
}

// listRuleSets godoc
// @Id listRuleSets
// @Summary Lists RuleSets
// @Accept json
// @Produce json
// @Tags RuleSet
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>RuleValues"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RuleSet
// @Failure 500 {string} string "Bad query request"
// @Router /rule_sets [get]
func (e *RuleSetController) listRuleSets(c echo.Context) error {
	var results []models.RuleSet
	err := e.db.QueryContext(models.RuleSet{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getRuleSet godoc
// @Id getRuleSet
// @Summary Gets RuleSet
// @Accept json
// @Produce json
// @Tags RuleSet
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>RuleValues"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RuleSet
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /rule_set/{id} [get]
func (e *RuleSetController) getRuleSet(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rulesetId, err := strconv.Atoi(c.Param("rulesetId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [RulesetId]"})
	}
	params = append(params, rulesetId)
	keys = append(keys, "ruleset_id = ?")

	// query builder
	var result models.RuleSet
	query := e.db.QueryContext(models.RuleSet{}, c)
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

// updateRuleSet godoc
// @Id updateRuleSet
// @Summary Updates RuleSet
// @Accept json
// @Produce json
// @Tags RuleSet
// @Param id path int true "Id"
// @Param rule_set body models.RuleSet true "RuleSet"
// @Success 200 {array} models.RuleSet
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /rule_set/{id} [patch]
func (e *RuleSetController) updateRuleSet(c echo.Context) error {
	request := new(models.RuleSet)
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

	// query builder
	var result models.RuleSet
	query := e.db.QueryContext(models.RuleSet{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createRuleSet godoc
// @Id createRuleSet
// @Summary Creates RuleSet
// @Accept json
// @Produce json
// @Param rule_set body models.RuleSet true "RuleSet"
// @Tags RuleSet
// @Success 200 {array} models.RuleSet
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /rule_set [put]
func (e *RuleSetController) createRuleSet(c echo.Context) error {
	ruleSet := new(models.RuleSet)
	if err := c.Bind(ruleSet); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.RuleSet{}, c).Model(&models.RuleSet{}).Create(&ruleSet).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, ruleSet)
}

// deleteRuleSet godoc
// @Id deleteRuleSet
// @Summary Deletes RuleSet
// @Accept json
// @Produce json
// @Tags RuleSet
// @Param id path int true "rulesetId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /rule_set/{id} [delete]
func (e *RuleSetController) deleteRuleSet(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rulesetId, err := strconv.Atoi(c.Param("rulesetId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, rulesetId)
	keys = append(keys, "ruleset_id = ?")

	// query builder
	var result models.RuleSet
	query := e.db.QueryContext(models.RuleSet{}, c)
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

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getRuleSetsBulk godoc
// @Id getRuleSetsBulk
// @Summary Gets RuleSets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags RuleSet
// @Success 200 {array} models.RuleSet
// @Failure 500 {string} string "Bad query request"
// @Router /rule_sets/bulk [post]
func (e *RuleSetController) getRuleSetsBulk(c echo.Context) error {
	var results []models.RuleSet

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

	err := e.db.QueryContext(models.RuleSet{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
