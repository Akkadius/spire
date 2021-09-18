package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GlobalLootController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGlobalLootController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GlobalLootController {
	return &GlobalLootController {
		db:     db,
		logger: logger,
	}
}

func (e *GlobalLootController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "global_loot/:global_loot", e.deleteGlobalLoot, nil),
		routes.RegisterRoute(http.MethodGet, "global_loot/:global_loot", e.getGlobalLoot, nil),
		routes.RegisterRoute(http.MethodGet, "global_loots", e.listGlobalLoots, nil),
		routes.RegisterRoute(http.MethodPatch, "global_loot/:global_loot", e.updateGlobalLoot, nil),
		routes.RegisterRoute(http.MethodPut, "global_loot", e.createGlobalLoot, nil),
	}
}

// listGlobalLoots godoc
// @Id listGlobalLoots
// @Summary Lists GlobalLoots
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GlobalLoot
// @Failure 500 {string} string "Bad query request"
// @Router /global_loots [get]
func (e *GlobalLootController) listGlobalLoots(c echo.Context) error {
	var results []models.GlobalLoot
	err := e.db.QueryContext(models.GlobalLoot{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGlobalLoot godoc
// @Id getGlobalLoot
// @Summary Gets GlobalLoot
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GlobalLoot
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /global_loot/{id} [get]
func (e *GlobalLootController) getGlobalLoot(c echo.Context) error {
	globalLootId, err := strconv.Atoi(c.Param("global_loot"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.GlobalLoot
	err = e.db.QueryContext(models.GlobalLoot{}, c).First(&result, globalLootId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGlobalLoot godoc
// @Id updateGlobalLoot
// @Summary Updates GlobalLoot
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param id path int true "Id"
// @Param global_loot body models.GlobalLoot true "GlobalLoot"
// @Success 200 {array} models.GlobalLoot
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /global_loot/{id} [patch]
func (e *GlobalLootController) updateGlobalLoot(c echo.Context) error {
	globalLoot := new(models.GlobalLoot)
	if err := c.Bind(globalLoot); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.GlobalLoot{}, c).Model(&models.GlobalLoot{}).First(&models.GlobalLoot{}, globalLoot.ID).Error
	if err != nil || globalLoot.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.GlobalLoot{}, c).Model(&models.GlobalLoot{}).Update(&globalLoot).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, globalLoot)
}

// createGlobalLoot godoc
// @Id createGlobalLoot
// @Summary Creates GlobalLoot
// @Accept json
// @Produce json
// @Param global_loot body models.GlobalLoot true "GlobalLoot"
// @Tags GlobalLoot
// @Success 200 {array} models.GlobalLoot
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /global_loot [put]
func (e *GlobalLootController) createGlobalLoot(c echo.Context) error {
	globalLoot := new(models.GlobalLoot)
	if err := c.Bind(globalLoot); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.GlobalLoot{}, c).Model(&models.GlobalLoot{}).Create(&globalLoot).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, globalLoot)
}

// deleteGlobalLoot godoc
// @Id deleteGlobalLoot
// @Summary Deletes GlobalLoot
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /global_loot/{id} [delete]
func (e *GlobalLootController) deleteGlobalLoot(c echo.Context) error {
	globalLootId, err := strconv.Atoi(c.Param("global_loot"))
	if err != nil {
		e.logger.Error(err)
	}

	globalLoot := new(models.GlobalLoot)
	err = e.db.Get(models.GlobalLoot{}, c).Model(&models.GlobalLoot{}).First(&globalLoot, globalLootId).Error
	if err != nil || globalLoot.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.GlobalLoot{}, c).Model(&models.GlobalLoot{}).Delete(&globalLoot).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
