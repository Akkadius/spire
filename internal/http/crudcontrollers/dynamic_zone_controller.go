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

type DynamicZoneController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewDynamicZoneController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *DynamicZoneController {
	return &DynamicZoneController{
		db:     db,
		logger: logger,
	}
}

func (e *DynamicZoneController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "dynamic_zone/:dynamic_zone", e.deleteDynamicZone, nil),
		routes.RegisterRoute(http.MethodGet, "dynamic_zone/:dynamic_zone", e.getDynamicZone, nil),
		routes.RegisterRoute(http.MethodGet, "dynamic_zones", e.listDynamicZones, nil),
		routes.RegisterRoute(http.MethodPost, "dynamic_zones/bulk", e.getDynamicZonesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "dynamic_zone/:dynamic_zone", e.updateDynamicZone, nil),
		routes.RegisterRoute(http.MethodPut, "dynamic_zone", e.createDynamicZone, nil),
	}
}

// listDynamicZones godoc
// @Id listDynamicZones
// @Summary Lists DynamicZones
// @Accept json
// @Produce json
// @Tags DynamicZone
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZone
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zones [get]
func (e *DynamicZoneController) listDynamicZones(c echo.Context) error {
	var results []models.DynamicZone
	err := e.db.QueryContext(models.DynamicZone{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDynamicZone godoc
// @Id getDynamicZone
// @Summary Gets DynamicZone
// @Accept json
// @Produce json
// @Tags DynamicZone
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZone
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone/{id} [get]
func (e *DynamicZoneController) getDynamicZone(c echo.Context) error {
	dynamicZoneId, err := strconv.Atoi(c.Param("dynamic_zone"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.DynamicZone
	err = e.db.QueryContext(models.DynamicZone{}, c).First(&result, dynamicZoneId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateDynamicZone godoc
// @Id updateDynamicZone
// @Summary Updates DynamicZone
// @Accept json
// @Produce json
// @Tags DynamicZone
// @Param id path int true "Id"
// @Param dynamic_zone body models.DynamicZone true "DynamicZone"
// @Success 200 {array} models.DynamicZone
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /dynamic_zone/{id} [patch]
func (e *DynamicZoneController) updateDynamicZone(c echo.Context) error {
	dynamicZone := new(models.DynamicZone)
	if err := c.Bind(dynamicZone); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.DynamicZone{}
	err := e.db.Get(models.DynamicZone{}, c).Model(&models.DynamicZone{}).First(&entity, dynamicZone.ID).Error
	if err != nil || dynamicZone.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.DynamicZone{}, c).Model(&entity).Select("*").Updates(&dynamicZone).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, dynamicZone)
}

// createDynamicZone godoc
// @Id createDynamicZone
// @Summary Creates DynamicZone
// @Accept json
// @Produce json
// @Param dynamic_zone body models.DynamicZone true "DynamicZone"
// @Tags DynamicZone
// @Success 200 {array} models.DynamicZone
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /dynamic_zone [put]
func (e *DynamicZoneController) createDynamicZone(c echo.Context) error {
	dynamicZone := new(models.DynamicZone)
	if err := c.Bind(dynamicZone); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.DynamicZone{}, c).Model(&models.DynamicZone{}).Create(&dynamicZone).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, dynamicZone)
}

// deleteDynamicZone godoc
// @Id deleteDynamicZone
// @Summary Deletes DynamicZone
// @Accept json
// @Produce json
// @Tags DynamicZone
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /dynamic_zone/{id} [delete]
func (e *DynamicZoneController) deleteDynamicZone(c echo.Context) error {
	dynamicZoneId, err := strconv.Atoi(c.Param("dynamic_zone"))
	if err != nil {
		e.logger.Error(err)
	}

	dynamicZone := new(models.DynamicZone)
	err = e.db.Get(models.DynamicZone{}, c).Model(&models.DynamicZone{}).First(&dynamicZone, dynamicZoneId).Error
	if err != nil || dynamicZone.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.DynamicZone{}, c).Model(&models.DynamicZone{}).Delete(&dynamicZone).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDynamicZonesBulk godoc
// @Id getDynamicZonesBulk
// @Summary Gets DynamicZones in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DynamicZone
// @Success 200 {array} models.DynamicZone
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zones/bulk [post]
func (e *DynamicZoneController) getDynamicZonesBulk(c echo.Context) error {
	var results []models.DynamicZone

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

	err := e.db.QueryContext(models.DynamicZone{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
