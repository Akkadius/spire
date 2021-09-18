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

type AdventureTemplateController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAdventureTemplateController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AdventureTemplateController {
	return &AdventureTemplateController {
		db:     db,
		logger: logger,
	}
}

func (e *AdventureTemplateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "adventure_template/:adventure_template", e.deleteAdventureTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_template/:adventure_template", e.getAdventureTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_templates", e.listAdventureTemplates, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_template/:adventure_template", e.updateAdventureTemplate, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_template", e.createAdventureTemplate, nil),
	}
}

// listAdventureTemplates godoc
// @Id listAdventureTemplates
// @Summary Lists AdventureTemplates
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_templates [get]
func (e *AdventureTemplateController) listAdventureTemplates(c echo.Context) error {
	var results []models.AdventureTemplate
	err := e.db.QueryContext(models.AdventureTemplate{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureTemplate godoc
// @Id getAdventureTemplate
// @Summary Gets AdventureTemplate
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplate
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template/{id} [get]
func (e *AdventureTemplateController) getAdventureTemplate(c echo.Context) error {
	adventureTemplateId, err := strconv.Atoi(c.Param("adventure_template"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.AdventureTemplate
	err = e.db.QueryContext(models.AdventureTemplate{}, c).First(&result, adventureTemplateId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAdventureTemplate godoc
// @Id updateAdventureTemplate
// @Summary Updates AdventureTemplate
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param id path int true "Id"
// @Param adventure_template body models.AdventureTemplate true "AdventureTemplate"
// @Success 200 {array} models.AdventureTemplate
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_template/{id} [patch]
func (e *AdventureTemplateController) updateAdventureTemplate(c echo.Context) error {
	adventureTemplate := new(models.AdventureTemplate)
	if err := c.Bind(adventureTemplate); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.AdventureTemplate{}, c).Model(&models.AdventureTemplate{}).First(&models.AdventureTemplate{}, adventureTemplate.ID).Error
	if err != nil || adventureTemplate.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AdventureTemplate{}, c).Model(&models.AdventureTemplate{}).Update(&adventureTemplate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, adventureTemplate)
}

// createAdventureTemplate godoc
// @Id createAdventureTemplate
// @Summary Creates AdventureTemplate
// @Accept json
// @Produce json
// @Param adventure_template body models.AdventureTemplate true "AdventureTemplate"
// @Tags AdventureTemplate
// @Success 200 {array} models.AdventureTemplate
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_template [put]
func (e *AdventureTemplateController) createAdventureTemplate(c echo.Context) error {
	adventureTemplate := new(models.AdventureTemplate)
	if err := c.Bind(adventureTemplate); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.AdventureTemplate{}, c).Model(&models.AdventureTemplate{}).Create(&adventureTemplate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, adventureTemplate)
}

// deleteAdventureTemplate godoc
// @Id deleteAdventureTemplate
// @Summary Deletes AdventureTemplate
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_template/{id} [delete]
func (e *AdventureTemplateController) deleteAdventureTemplate(c echo.Context) error {
	adventureTemplateId, err := strconv.Atoi(c.Param("adventure_template"))
	if err != nil {
		e.logger.Error(err)
	}

	adventureTemplate := new(models.AdventureTemplate)
	err = e.db.Get(models.AdventureTemplate{}, c).Model(&models.AdventureTemplate{}).First(&adventureTemplate, adventureTemplateId).Error
	if err != nil || adventureTemplate.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AdventureTemplate{}, c).Model(&models.AdventureTemplate{}).Delete(&adventureTemplate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
