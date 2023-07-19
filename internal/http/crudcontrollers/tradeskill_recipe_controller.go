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

type TradeskillRecipeController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewTradeskillRecipeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *TradeskillRecipeController {
	return &TradeskillRecipeController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *TradeskillRecipeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipe/:id", e.getTradeskillRecipe, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipes", e.listTradeskillRecipes, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipes/count", e.getTradeskillRecipesCount, nil),
		routes.RegisterRoute(http.MethodPut, "tradeskill_recipe", e.createTradeskillRecipe, nil),
		routes.RegisterRoute(http.MethodDelete, "tradeskill_recipe/:id", e.deleteTradeskillRecipe, nil),
		routes.RegisterRoute(http.MethodPatch, "tradeskill_recipe/:id", e.updateTradeskillRecipe, nil),
		routes.RegisterRoute(http.MethodPost, "tradeskill_recipes/bulk", e.getTradeskillRecipesBulk, nil),
	}
}

// listTradeskillRecipes godoc
// @Id listTradeskillRecipes
// @Summary Lists TradeskillRecipes
// @Accept json
// @Produce json
// @Tags TradeskillRecipe
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipe
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipes [get]
func (e *TradeskillRecipeController) listTradeskillRecipes(c echo.Context) error {
	var results []models.TradeskillRecipe
	err := e.db.QueryContext(models.TradeskillRecipe{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTradeskillRecipe godoc
// @Id getTradeskillRecipe
// @Summary Gets TradeskillRecipe
// @Accept json
// @Produce json
// @Tags TradeskillRecipe
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipe
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe/{id} [get]
func (e *TradeskillRecipeController) getTradeskillRecipe(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.TradeskillRecipe
	query := e.db.QueryContext(models.TradeskillRecipe{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTradeskillRecipe godoc
// @Id updateTradeskillRecipe
// @Summary Updates TradeskillRecipe
// @Accept json
// @Produce json
// @Tags TradeskillRecipe
// @Param id path int true "Id"
// @Param tradeskill_recipe body models.TradeskillRecipe true "TradeskillRecipe"
// @Success 200 {array} models.TradeskillRecipe
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /tradeskill_recipe/{id} [patch]
func (e *TradeskillRecipeController) updateTradeskillRecipe(c echo.Context) error {
	request := new(models.TradeskillRecipe)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.TradeskillRecipe
	query := e.db.QueryContext(models.TradeskillRecipe{}, c)
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
		event := fmt.Sprintf("Updated [TradeskillRecipe] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createTradeskillRecipe godoc
// @Id createTradeskillRecipe
// @Summary Creates TradeskillRecipe
// @Accept json
// @Produce json
// @Param tradeskill_recipe body models.TradeskillRecipe true "TradeskillRecipe"
// @Tags TradeskillRecipe
// @Success 200 {array} models.TradeskillRecipe
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /tradeskill_recipe [put]
func (e *TradeskillRecipeController) createTradeskillRecipe(c echo.Context) error {
	tradeskillRecipe := new(models.TradeskillRecipe)
	if err := c.Bind(tradeskillRecipe); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.TradeskillRecipe{}, c).Model(&models.TradeskillRecipe{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&tradeskillRecipe).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.TradeskillRecipe{}, tradeskillRecipe)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [TradeskillRecipe] [%v] data [%v]", tradeskillRecipe.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, tradeskillRecipe)
}

// deleteTradeskillRecipe godoc
// @Id deleteTradeskillRecipe
// @Summary Deletes TradeskillRecipe
// @Accept json
// @Produce json
// @Tags TradeskillRecipe
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tradeskill_recipe/{id} [delete]
func (e *TradeskillRecipeController) deleteTradeskillRecipe(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.TradeskillRecipe
	query := e.db.QueryContext(models.TradeskillRecipe{}, c)
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
		event := fmt.Sprintf("Deleted [TradeskillRecipe] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getTradeskillRecipesBulk godoc
// @Id getTradeskillRecipesBulk
// @Summary Gets TradeskillRecipes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags TradeskillRecipe
// @Success 200 {array} models.TradeskillRecipe
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipes/bulk [post]
func (e *TradeskillRecipeController) getTradeskillRecipesBulk(c echo.Context) error {
	var results []models.TradeskillRecipe

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

	err := e.db.QueryContext(models.TradeskillRecipe{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getTradeskillRecipesCount godoc
// @Id getTradeskillRecipesCount
// @Summary Counts TradeskillRecipes
// @Accept json
// @Produce json
// @Tags TradeskillRecipe
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipe
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipes/count [get]
func (e *TradeskillRecipeController) getTradeskillRecipesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.TradeskillRecipe{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}