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

type AdventureTemplateEntryController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAdventureTemplateEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AdventureTemplateEntryController {
	return &AdventureTemplateEntryController {
		db:     db,
		logger: logger,
	}
}

func (e *AdventureTemplateEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "adventure_template_entry/:adventure_template_entry", e.deleteAdventureTemplateEntry, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_template_entry/:adventure_template_entry", e.getAdventureTemplateEntry, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_template_entries", e.listAdventureTemplateEntries, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_template_entry/:adventure_template_entry", e.updateAdventureTemplateEntry, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_template_entry", e.createAdventureTemplateEntry, nil),
	}
}

// listAdventureTemplateEntries godoc
// @Id listAdventureTemplateEntries
// @Summary Lists AdventureTemplateEntries
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entries [get]
func (e *AdventureTemplateEntryController) listAdventureTemplateEntries(c echo.Context) error {
	var results []models.AdventureTemplateEntry
	err := e.db.QueryContext(models.AdventureTemplateEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureTemplateEntry godoc
// @Id getAdventureTemplateEntry
// @Summary Gets AdventureTemplateEntry
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template_entry/{id} [get]
func (e *AdventureTemplateEntryController) getAdventureTemplateEntry(c echo.Context) error {
	adventureTemplateEntryId, err := strconv.Atoi(c.Param("adventure_template_entry"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.AdventureTemplateEntry
	err = e.db.QueryContext(models.AdventureTemplateEntry{}, c).First(&result, adventureTemplateEntryId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAdventureTemplateEntry godoc
// @Id updateAdventureTemplateEntry
// @Summary Updates AdventureTemplateEntry
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param id path int true "Id"
// @Param adventure_template_entry body models.AdventureTemplateEntry true "AdventureTemplateEntry"
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_template_entry/{id} [patch]
func (e *AdventureTemplateEntryController) updateAdventureTemplateEntry(c echo.Context) error {
	adventureTemplateEntry := new(models.AdventureTemplateEntry)
	if err := c.Bind(adventureTemplateEntry); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.AdventureTemplateEntry{}, c).Model(&models.AdventureTemplateEntry{}).First(&models.AdventureTemplateEntry{}, adventureTemplateEntry.ID).Error
	if err != nil || adventureTemplateEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AdventureTemplateEntry{}, c).Model(&models.AdventureTemplateEntry{}).Update(&adventureTemplateEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, adventureTemplateEntry)
}

// createAdventureTemplateEntry godoc
// @Id createAdventureTemplateEntry
// @Summary Creates AdventureTemplateEntry
// @Accept json
// @Produce json
// @Param adventure_template_entry body models.AdventureTemplateEntry true "AdventureTemplateEntry"
// @Tags AdventureTemplateEntry
// @Success 200 {array} models.AdventureTemplateEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_template_entry [put]
func (e *AdventureTemplateEntryController) createAdventureTemplateEntry(c echo.Context) error {
	adventureTemplateEntry := new(models.AdventureTemplateEntry)
	if err := c.Bind(adventureTemplateEntry); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.AdventureTemplateEntry{}, c).Model(&models.AdventureTemplateEntry{}).Create(&adventureTemplateEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, adventureTemplateEntry)
}

// deleteAdventureTemplateEntry godoc
// @Id deleteAdventureTemplateEntry
// @Summary Deletes AdventureTemplateEntry
// @Accept json
// @Produce json
// @Tags AdventureTemplateEntry
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_template_entry/{id} [delete]
func (e *AdventureTemplateEntryController) deleteAdventureTemplateEntry(c echo.Context) error {
	adventureTemplateEntryId, err := strconv.Atoi(c.Param("adventure_template_entry"))
	if err != nil {
		e.logger.Error(err)
	}

	adventureTemplateEntry := new(models.AdventureTemplateEntry)
	err = e.db.Get(models.AdventureTemplateEntry{}, c).Model(&models.AdventureTemplateEntry{}).First(&adventureTemplateEntry, adventureTemplateEntryId).Error
	if err != nil || adventureTemplateEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AdventureTemplateEntry{}, c).Model(&models.AdventureTemplateEntry{}).Delete(&adventureTemplateEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
