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

type TradeskillRecipeEntryController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTradeskillRecipeEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TradeskillRecipeEntryController {
	return &TradeskillRecipeEntryController{
		db:     db,
		logger: logger,
	}
}

func (e *TradeskillRecipeEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "tradeskill_recipe_entry/:tradeskill_recipe_entry", e.deleteTradeskillRecipeEntry, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipe_entry/:tradeskill_recipe_entry", e.getTradeskillRecipeEntry, nil),
		routes.RegisterRoute(http.MethodGet, "tradeskill_recipe_entries", e.listTradeskillRecipeEntries, nil),
		routes.RegisterRoute(http.MethodPost, "tradeskill_recipe_entries/bulk", e.getTradeskillRecipeEntriesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "tradeskill_recipe_entry/:tradeskill_recipe_entry", e.updateTradeskillRecipeEntry, nil),
		routes.RegisterRoute(http.MethodPut, "tradeskill_recipe_entry", e.createTradeskillRecipeEntry, nil),
	}
}

// listTradeskillRecipeEntries godoc
// @Id listTradeskillRecipeEntries
// @Summary Lists TradeskillRecipeEntries
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe_entries [get]
func (e *TradeskillRecipeEntryController) listTradeskillRecipeEntries(c echo.Context) error {
	var results []models.TradeskillRecipeEntry
	err := e.db.QueryContext(models.TradeskillRecipeEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTradeskillRecipeEntry godoc
// @Id getTradeskillRecipeEntry
// @Summary Gets TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe_entry/{id} [get]
func (e *TradeskillRecipeEntryController) getTradeskillRecipeEntry(c echo.Context) error {
	tradeskillRecipeEntryId, err := strconv.Atoi(c.Param("tradeskill_recipe_entry"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.TradeskillRecipeEntry
	err = e.db.QueryContext(models.TradeskillRecipeEntry{}, c).First(&result, tradeskillRecipeEntryId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTradeskillRecipeEntry godoc
// @Id updateTradeskillRecipeEntry
// @Summary Updates TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param id path int true "Id"
// @Param tradeskill_recipe_entry body models.TradeskillRecipeEntry true "TradeskillRecipeEntry"
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /tradeskill_recipe_entry/{id} [patch]
func (e *TradeskillRecipeEntryController) updateTradeskillRecipeEntry(c echo.Context) error {
	tradeskillRecipeEntry := new(models.TradeskillRecipeEntry)
	if err := c.Bind(tradeskillRecipeEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.TradeskillRecipeEntry{}
	err := e.db.Get(models.TradeskillRecipeEntry{}, c).Model(&models.TradeskillRecipeEntry{}).First(&entity, tradeskillRecipeEntry.ID).Error
	if err != nil || tradeskillRecipeEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.TradeskillRecipeEntry{}, c).Model(&entity).Select("*").Updates(&tradeskillRecipeEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, tradeskillRecipeEntry)
}

// createTradeskillRecipeEntry godoc
// @Id createTradeskillRecipeEntry
// @Summary Creates TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Param tradeskill_recipe_entry body models.TradeskillRecipeEntry true "TradeskillRecipeEntry"
// @Tags TradeskillRecipeEntry
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /tradeskill_recipe_entry [put]
func (e *TradeskillRecipeEntryController) createTradeskillRecipeEntry(c echo.Context) error {
	tradeskillRecipeEntry := new(models.TradeskillRecipeEntry)
	if err := c.Bind(tradeskillRecipeEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.TradeskillRecipeEntry{}, c).Model(&models.TradeskillRecipeEntry{}).Create(&tradeskillRecipeEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, tradeskillRecipeEntry)
}

// deleteTradeskillRecipeEntry godoc
// @Id deleteTradeskillRecipeEntry
// @Summary Deletes TradeskillRecipeEntry
// @Accept json
// @Produce json
// @Tags TradeskillRecipeEntry
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tradeskill_recipe_entry/{id} [delete]
func (e *TradeskillRecipeEntryController) deleteTradeskillRecipeEntry(c echo.Context) error {
	tradeskillRecipeEntryId, err := strconv.Atoi(c.Param("tradeskill_recipe_entry"))
	if err != nil {
		e.logger.Error(err)
	}

	tradeskillRecipeEntry := new(models.TradeskillRecipeEntry)
	err = e.db.Get(models.TradeskillRecipeEntry{}, c).Model(&models.TradeskillRecipeEntry{}).First(&tradeskillRecipeEntry, tradeskillRecipeEntryId).Error
	if err != nil || tradeskillRecipeEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.TradeskillRecipeEntry{}, c).Model(&models.TradeskillRecipeEntry{}).Delete(&tradeskillRecipeEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getTradeskillRecipeEntriesBulk godoc
// @Id getTradeskillRecipeEntriesBulk
// @Summary Gets TradeskillRecipeEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags TradeskillRecipeEntry
// @Success 200 {array} models.TradeskillRecipeEntry
// @Failure 500 {string} string "Bad query request"
// @Router /tradeskill_recipe_entries/bulk [post]
func (e *TradeskillRecipeEntryController) getTradeskillRecipeEntriesBulk(c echo.Context) error {
	var results []models.TradeskillRecipeEntry

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

	err := e.db.QueryContext(models.TradeskillRecipeEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
