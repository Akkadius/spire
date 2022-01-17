package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TradeskillRecipeController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTradeskillRecipeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TradeskillRecipeController {
	return &TradeskillRecipeController{
		db:     db,
		logger: logger,
	}
}

func (e *TradeskillRecipeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "tradeskill_recipe/:tradeskill_recipe", e.deleteTradeskillRecipe, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipe/:tradeskill_recipe", e.getTradeskillRecipe, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipes", e.listTradeskillRecipes, nil),
		routes.RegisterRoute(http.MethodPost, "tradeskill_recipes/bulk", e.getTradeskillRecipesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "tradeskill_recipe/:tradeskill_recipe", e.updateTradeskillRecipe, nil),
		routes.RegisterRoute(http.MethodPut, "tradeskill_recipe", e.createTradeskillRecipe, nil),
	}
}

// listTradeskillRecipes godoc
// @Id listTradeskillRecipes
// @Summary Lists TradeskillRecipes
// @Accept json
// @Produce json
// @Tags TradeskillRecipe
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipe
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe/{id} [get]
func (e *TradeskillRecipeController) getTradeskillRecipe(c echo.Context) error {
	tradeskillRecipeId, err := strconv.Atoi(c.Param("tradeskill_recipe"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.TradeskillRecipe
	err = e.db.QueryContext(models.TradeskillRecipe{}, c).First(&result, tradeskillRecipeId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

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
	tradeskillRecipe := new(models.TradeskillRecipe)
	if err := c.Bind(tradeskillRecipe); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.TradeskillRecipe{}
	err := e.db.Get(models.TradeskillRecipe{}, c).Model(&models.TradeskillRecipe{}).First(&entity, tradeskillRecipe.ID).Error
	if err != nil || tradeskillRecipe.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.TradeskillRecipe{}, c).Model(&entity).Select("*").Updates(&tradeskillRecipe).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, tradeskillRecipe)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.TradeskillRecipe{}, c).Model(&models.TradeskillRecipe{}).Create(&tradeskillRecipe).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, tradeskillRecipe)
}

// deleteTradeskillRecipe godoc
// @Id deleteTradeskillRecipe
// @Summary Deletes TradeskillRecipe
// @Accept json
// @Produce json
// @Tags TradeskillRecipe
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tradeskill_recipe/{id} [delete]
func (e *TradeskillRecipeController) deleteTradeskillRecipe(c echo.Context) error {
	tradeskillRecipeId, err := strconv.Atoi(c.Param("tradeskill_recipe"))
	if err != nil {
		e.logger.Error(err)
	}

	tradeskillRecipe := new(models.TradeskillRecipe)
	err = e.db.Get(models.TradeskillRecipe{}, c).Model(&models.TradeskillRecipe{}).First(&tradeskillRecipe, tradeskillRecipeId).Error
	if err != nil || tradeskillRecipe.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.TradeskillRecipe{}, c).Model(&models.TradeskillRecipe{}).Delete(&tradeskillRecipe).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
