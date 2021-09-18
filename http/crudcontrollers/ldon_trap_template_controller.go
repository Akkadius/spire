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

type LdonTrapTemplateController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLdonTrapTemplateController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LdonTrapTemplateController {
	return &LdonTrapTemplateController {
		db:     db,
		logger: logger,
	}
}

func (e *LdonTrapTemplateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "ldon_trap_template/:ldon_trap_template", e.deleteLdonTrapTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "ldon_trap_template/:ldon_trap_template", e.getLdonTrapTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "ldon_trap_templates", e.listLdonTrapTemplates, nil),
		routes.RegisterRoute(http.MethodPatch, "ldon_trap_template/:ldon_trap_template", e.updateLdonTrapTemplate, nil),
		routes.RegisterRoute(http.MethodPut, "ldon_trap_template", e.createLdonTrapTemplate, nil),
	}
}

// listLdonTrapTemplates godoc
// @Id listLdonTrapTemplates
// @Summary Lists LdonTrapTemplates
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_templates [get]
func (e *LdonTrapTemplateController) listLdonTrapTemplates(c echo.Context) error {
	var results []models.LdonTrapTemplate
	err := e.db.QueryContext(models.LdonTrapTemplate{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLdonTrapTemplate godoc
// @Id getLdonTrapTemplate
// @Summary Gets LdonTrapTemplate
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_template/{id} [get]
func (e *LdonTrapTemplateController) getLdonTrapTemplate(c echo.Context) error {
	ldonTrapTemplateId, err := strconv.Atoi(c.Param("ldon_trap_template"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.LdonTrapTemplate
	err = e.db.QueryContext(models.LdonTrapTemplate{}, c).First(&result, ldonTrapTemplateId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLdonTrapTemplate godoc
// @Id updateLdonTrapTemplate
// @Summary Updates LdonTrapTemplate
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param id path int true "Id"
// @Param ldon_trap_template body models.LdonTrapTemplate true "LdonTrapTemplate"
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /ldon_trap_template/{id} [patch]
func (e *LdonTrapTemplateController) updateLdonTrapTemplate(c echo.Context) error {
	ldonTrapTemplate := new(models.LdonTrapTemplate)
	if err := c.Bind(ldonTrapTemplate); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.LdonTrapTemplate{}, c).Model(&models.LdonTrapTemplate{}).First(&models.LdonTrapTemplate{}, ldonTrapTemplate.ID).Error
	if err != nil || ldonTrapTemplate.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LdonTrapTemplate{}, c).Model(&models.LdonTrapTemplate{}).Updates(&ldonTrapTemplate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, ldonTrapTemplate)
}

// createLdonTrapTemplate godoc
// @Id createLdonTrapTemplate
// @Summary Creates LdonTrapTemplate
// @Accept json
// @Produce json
// @Param ldon_trap_template body models.LdonTrapTemplate true "LdonTrapTemplate"
// @Tags LdonTrapTemplate
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /ldon_trap_template [put]
func (e *LdonTrapTemplateController) createLdonTrapTemplate(c echo.Context) error {
	ldonTrapTemplate := new(models.LdonTrapTemplate)
	if err := c.Bind(ldonTrapTemplate); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.LdonTrapTemplate{}, c).Model(&models.LdonTrapTemplate{}).Create(&ldonTrapTemplate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, ldonTrapTemplate)
}

// deleteLdonTrapTemplate godoc
// @Id deleteLdonTrapTemplate
// @Summary Deletes LdonTrapTemplate
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /ldon_trap_template/{id} [delete]
func (e *LdonTrapTemplateController) deleteLdonTrapTemplate(c echo.Context) error {
	ldonTrapTemplateId, err := strconv.Atoi(c.Param("ldon_trap_template"))
	if err != nil {
		e.logger.Error(err)
	}

	ldonTrapTemplate := new(models.LdonTrapTemplate)
	err = e.db.Get(models.LdonTrapTemplate{}, c).Model(&models.LdonTrapTemplate{}).First(&ldonTrapTemplate, ldonTrapTemplateId).Error
	if err != nil || ldonTrapTemplate.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LdonTrapTemplate{}, c).Model(&models.LdonTrapTemplate{}).Delete(&ldonTrapTemplate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
