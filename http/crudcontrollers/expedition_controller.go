package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ExpeditionController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewExpeditionController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ExpeditionController {
	return &ExpeditionController{
		db:     db,
		logger: logger,
	}
}

func (e *ExpeditionController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "expedition/:expedition", e.deleteExpedition, nil),
		routes.RegisterRoute(http.MethodGet, "expedition/:expedition", e.getExpedition, nil),
		routes.RegisterRoute(http.MethodGet, "expeditions", e.listExpeditions, nil),
		routes.RegisterRoute(http.MethodPost, "expeditions/bulk", e.getExpeditionsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "expedition/:expedition", e.updateExpedition, nil),
		routes.RegisterRoute(http.MethodPut, "expedition", e.createExpedition, nil),
	}
}

// listExpeditions godoc
// @Id listExpeditions
// @Summary Lists Expeditions
// @Accept json
// @Produce json
// @Tags Expedition
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Expedition
// @Failure 500 {string} string "Bad query request"
// @Router /expeditions [get]
func (e *ExpeditionController) listExpeditions(c echo.Context) error {
	var results []models.Expedition
	err := e.db.QueryContext(models.Expedition{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getExpedition godoc
// @Id getExpedition
// @Summary Gets Expedition
// @Accept json
// @Produce json
// @Tags Expedition
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Expedition
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /expedition/{id} [get]
func (e *ExpeditionController) getExpedition(c echo.Context) error {
	expeditionId, err := strconv.Atoi(c.Param("expedition"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Expedition
	err = e.db.QueryContext(models.Expedition{}, c).First(&result, expeditionId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateExpedition godoc
// @Id updateExpedition
// @Summary Updates Expedition
// @Accept json
// @Produce json
// @Tags Expedition
// @Param id path int true "Id"
// @Param expedition body models.Expedition true "Expedition"
// @Success 200 {array} models.Expedition
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /expedition/{id} [patch]
func (e *ExpeditionController) updateExpedition(c echo.Context) error {
	expedition := new(models.Expedition)
	if err := c.Bind(expedition); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Expedition{}
	err := e.db.Get(models.Expedition{}, c).Model(&models.Expedition{}).First(&entity, expedition.ID).Error
	if err != nil || expedition.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Expedition{}, c).Model(&entity).Updates(&expedition).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, expedition)
}

// createExpedition godoc
// @Id createExpedition
// @Summary Creates Expedition
// @Accept json
// @Produce json
// @Param expedition body models.Expedition true "Expedition"
// @Tags Expedition
// @Success 200 {array} models.Expedition
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /expedition [put]
func (e *ExpeditionController) createExpedition(c echo.Context) error {
	expedition := new(models.Expedition)
	if err := c.Bind(expedition); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Expedition{}, c).Model(&models.Expedition{}).Create(&expedition).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, expedition)
}

// deleteExpedition godoc
// @Id deleteExpedition
// @Summary Deletes Expedition
// @Accept json
// @Produce json
// @Tags Expedition
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /expedition/{id} [delete]
func (e *ExpeditionController) deleteExpedition(c echo.Context) error {
	expeditionId, err := strconv.Atoi(c.Param("expedition"))
	if err != nil {
		e.logger.Error(err)
	}

	expedition := new(models.Expedition)
	err = e.db.Get(models.Expedition{}, c).Model(&models.Expedition{}).First(&expedition, expeditionId).Error
	if err != nil || expedition.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Expedition{}, c).Model(&models.Expedition{}).Delete(&expedition).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getExpeditionsBulk godoc
// @Id getExpeditionsBulk
// @Summary Gets Expeditions in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Expedition
// @Success 200 {array} models.Expedition
// @Failure 500 {string} string "Bad query request"
// @Router /expeditions/bulk [post]
func (e *ExpeditionController) getExpeditionsBulk(c echo.Context) error {
	var results []models.Expedition

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

	err := e.db.QueryContext(models.Expedition{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
