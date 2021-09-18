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

type AdventureTemplateEntryFlavorController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAdventureTemplateEntryFlavorController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AdventureTemplateEntryFlavorController {
	return &AdventureTemplateEntryFlavorController {
		db:     db,
		logger: logger,
	}
}

func (e *AdventureTemplateEntryFlavorController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "adventure_template_entry_flavor/:adventure_template_entry_flavor", e.deleteAdventureTemplateEntryFlavor, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_template_entry_flavor/:adventure_template_entry_flavor", e.getAdventureTemplateEntryFlavor, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_template_entry_flavors", e.listAdventureTemplateEntryFlavors, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_template_entry_flavor/:adventure_template_entry_flavor", e.updateAdventureTemplateEntryFlavor, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_template_entry_flavor", e.createAdventureTemplateEntryFlavor, nil),
	}
}

// listAdventureTemplateEntryFlavors godoc
// @Id listAdventureTemplateEntryFlavors
// @Summary Lists AdventureTemplateEntryFlavors
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntryFlavor
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplateEntryFlavor
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entry_flavors [get]
func (e *AdventureTemplateEntryFlavorController) listAdventureTemplateEntryFlavors(c echo.Context) error {
	var results []models.AdventureTemplateEntryFlavor
	err := e.db.QueryContext(models.AdventureTemplateEntryFlavor{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureTemplateEntryFlavor godoc
// @Id getAdventureTemplateEntryFlavor
// @Summary Gets AdventureTemplateEntryFlavor
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntryFlavor
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplateEntryFlavor
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entry_flavor/{id} [get]
func (e *AdventureTemplateEntryFlavorController) getAdventureTemplateEntryFlavor(c echo.Context) error {
	adventureTemplateEntryFlavorId, err := strconv.Atoi(c.Param("adventure_template_entry_flavor"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.AdventureTemplateEntryFlavor
	err = e.db.QueryContext(models.AdventureTemplateEntryFlavor{}, c).First(&result, adventureTemplateEntryFlavorId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAdventureTemplateEntryFlavor godoc
// @Id updateAdventureTemplateEntryFlavor
// @Summary Updates AdventureTemplateEntryFlavor
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntryFlavor
// @Param id path int true "Id"
// @Param adventure_template_entry_flavor body models.AdventureTemplateEntryFlavor true "AdventureTemplateEntryFlavor"
// @Success 200 {array} models.AdventureTemplateEntryFlavor
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_template_entry_flavor/{id} [patch]
func (e *AdventureTemplateEntryFlavorController) updateAdventureTemplateEntryFlavor(c echo.Context) error {
	adventureTemplateEntryFlavor := new(models.AdventureTemplateEntryFlavor)
	if err := c.Bind(adventureTemplateEntryFlavor); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.AdventureTemplateEntryFlavor{}, c).Model(&models.AdventureTemplateEntryFlavor{}).First(&models.AdventureTemplateEntryFlavor{}, adventureTemplateEntryFlavor.ID).Error
	if err != nil || adventureTemplateEntryFlavor.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AdventureTemplateEntryFlavor{}, c).Model(&models.AdventureTemplateEntryFlavor{}).Update(&adventureTemplateEntryFlavor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, adventureTemplateEntryFlavor)
}

// createAdventureTemplateEntryFlavor godoc
// @Id createAdventureTemplateEntryFlavor
// @Summary Creates AdventureTemplateEntryFlavor
// @Accept json
// @Produce json
// @Param adventure_template_entry_flavor body models.AdventureTemplateEntryFlavor true "AdventureTemplateEntryFlavor"
// @Tags AdventureTemplateEntryFlavor
// @Success 200 {array} models.AdventureTemplateEntryFlavor
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_template_entry_flavor [put]
func (e *AdventureTemplateEntryFlavorController) createAdventureTemplateEntryFlavor(c echo.Context) error {
	adventureTemplateEntryFlavor := new(models.AdventureTemplateEntryFlavor)
	if err := c.Bind(adventureTemplateEntryFlavor); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.AdventureTemplateEntryFlavor{}, c).Model(&models.AdventureTemplateEntryFlavor{}).Create(&adventureTemplateEntryFlavor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, adventureTemplateEntryFlavor)
}

// deleteAdventureTemplateEntryFlavor godoc
// @Id deleteAdventureTemplateEntryFlavor
// @Summary Deletes AdventureTemplateEntryFlavor
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntryFlavor
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_template_entry_flavor/{id} [delete]
func (e *AdventureTemplateEntryFlavorController) deleteAdventureTemplateEntryFlavor(c echo.Context) error {
	adventureTemplateEntryFlavorId, err := strconv.Atoi(c.Param("adventure_template_entry_flavor"))
	if err != nil {
		e.logger.Error(err)
	}

	adventureTemplateEntryFlavor := new(models.AdventureTemplateEntryFlavor)
	err = e.db.Get(models.AdventureTemplateEntryFlavor{}, c).Model(&models.AdventureTemplateEntryFlavor{}).First(&adventureTemplateEntryFlavor, adventureTemplateEntryFlavorId).Error
	if err != nil || adventureTemplateEntryFlavor.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AdventureTemplateEntryFlavor{}, c).Model(&models.AdventureTemplateEntryFlavor{}).Delete(&adventureTemplateEntryFlavor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
