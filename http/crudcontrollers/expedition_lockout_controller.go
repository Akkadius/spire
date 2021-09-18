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

type ExpeditionLockoutController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewExpeditionLockoutController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ExpeditionLockoutController {
	return &ExpeditionLockoutController {
		db:     db,
		logger: logger,
	}
}

func (e *ExpeditionLockoutController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "expedition_lockout/:expedition_lockout", e.deleteExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodGet, "expedition_lockout/:expedition_lockout", e.getExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodGet, "expedition_lockouts", e.listExpeditionLockouts, nil),
		routes.RegisterRoute(http.MethodPatch, "expedition_lockout/:expedition_lockout", e.updateExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodPut, "expedition_lockout", e.createExpeditionLockout, nil),
	}
}

// listExpeditionLockouts godoc
// @Id listExpeditionLockouts
// @Summary Lists ExpeditionLockouts
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ExpeditionLockout
// @Failure 500 {string} string "Bad query request"
// @Router /expedition_lockouts [get]
func (e *ExpeditionLockoutController) listExpeditionLockouts(c echo.Context) error {
	var results []models.ExpeditionLockout
	err := e.db.QueryContext(models.ExpeditionLockout{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getExpeditionLockout godoc
// @Id getExpeditionLockout
// @Summary Gets ExpeditionLockout
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ExpeditionLockout
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /expedition_lockout/{id} [get]
func (e *ExpeditionLockoutController) getExpeditionLockout(c echo.Context) error {
	expeditionLockoutId, err := strconv.Atoi(c.Param("expedition_lockout"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.ExpeditionLockout
	err = e.db.QueryContext(models.ExpeditionLockout{}, c).First(&result, expeditionLockoutId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateExpeditionLockout godoc
// @Id updateExpeditionLockout
// @Summary Updates ExpeditionLockout
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param id path int true "Id"
// @Param expedition_lockout body models.ExpeditionLockout true "ExpeditionLockout"
// @Success 200 {array} models.ExpeditionLockout
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /expedition_lockout/{id} [patch]
func (e *ExpeditionLockoutController) updateExpeditionLockout(c echo.Context) error {
	expeditionLockout := new(models.ExpeditionLockout)
	if err := c.Bind(expeditionLockout); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.ExpeditionLockout{}, c).Model(&models.ExpeditionLockout{}).First(&models.ExpeditionLockout{}, expeditionLockout.ID).Error
	if err != nil || expeditionLockout.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ExpeditionLockout{}, c).Model(&models.ExpeditionLockout{}).Update(&expeditionLockout).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, expeditionLockout)
}

// createExpeditionLockout godoc
// @Id createExpeditionLockout
// @Summary Creates ExpeditionLockout
// @Accept json
// @Produce json
// @Param expedition_lockout body models.ExpeditionLockout true "ExpeditionLockout"
// @Tags ExpeditionLockout
// @Success 200 {array} models.ExpeditionLockout
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /expedition_lockout [put]
func (e *ExpeditionLockoutController) createExpeditionLockout(c echo.Context) error {
	expeditionLockout := new(models.ExpeditionLockout)
	if err := c.Bind(expeditionLockout); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.ExpeditionLockout{}, c).Model(&models.ExpeditionLockout{}).Create(&expeditionLockout).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, expeditionLockout)
}

// deleteExpeditionLockout godoc
// @Id deleteExpeditionLockout
// @Summary Deletes ExpeditionLockout
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /expedition_lockout/{id} [delete]
func (e *ExpeditionLockoutController) deleteExpeditionLockout(c echo.Context) error {
	expeditionLockoutId, err := strconv.Atoi(c.Param("expedition_lockout"))
	if err != nil {
		e.logger.Error(err)
	}

	expeditionLockout := new(models.ExpeditionLockout)
	err = e.db.Get(models.ExpeditionLockout{}, c).Model(&models.ExpeditionLockout{}).First(&expeditionLockout, expeditionLockoutId).Error
	if err != nil || expeditionLockout.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ExpeditionLockout{}, c).Model(&models.ExpeditionLockout{}).Delete(&expeditionLockout).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
