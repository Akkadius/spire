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

type ZoneController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewZoneController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ZoneController {
	return &ZoneController{
		db:     db,
		logger: logger,
	}
}

func (e *ZoneController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "zone/:zone", e.deleteZone, nil),
		routes.RegisterRoute(http.MethodGet, "zone/:zone", e.getZone, nil),
		routes.RegisterRoute(http.MethodGet, "zones", e.listZones, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getZonesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "zone/:zone", e.updateZone, nil),
		routes.RegisterRoute(http.MethodPut, "zone", e.createZone, nil),
	}
}

// listZones godoc
// @Id listZones
// @Summary Lists Zones
// @Accept json
// @Produce json
// @Tags Zone
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Zone
// @Failure 500 {string} string "Bad query request"
// @Router /zones [get]
func (e *ZoneController) listZones(c echo.Context) error {
	var results []models.Zone
	err := e.db.QueryContext(models.Zone{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getZone godoc
// @Id getZone
// @Summary Gets Zone
// @Accept json
// @Produce json
// @Tags Zone
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Zone
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /zone/{id} [get]
func (e *ZoneController) getZone(c echo.Context) error {
	zoneId, err := strconv.Atoi(c.Param("zone"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Zone
	err = e.db.QueryContext(models.Zone{}, c).First(&result, zoneId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateZone godoc
// @Id updateZone
// @Summary Updates Zone
// @Accept json
// @Produce json
// @Tags Zone
// @Param id path int true "Id"
// @Param zone body models.Zone true "Zone"
// @Success 200 {array} models.Zone
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /zone/{id} [patch]
func (e *ZoneController) updateZone(c echo.Context) error {
	zone := new(models.Zone)
	if err := c.Bind(zone); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Zone{}, c).Model(&models.Zone{}).First(&models.Zone{}, zone.ID).Error
	if err != nil || zone.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Zone{}, c).Model(&models.Zone{}).Updates(&zone).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, zone)
}

// createZone godoc
// @Id createZone
// @Summary Creates Zone
// @Accept json
// @Produce json
// @Param zone body models.Zone true "Zone"
// @Tags Zone
// @Success 200 {array} models.Zone
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /zone [put]
func (e *ZoneController) createZone(c echo.Context) error {
	zone := new(models.Zone)
	if err := c.Bind(zone); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Zone{}, c).Model(&models.Zone{}).Create(&zone).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, zone)
}

// deleteZone godoc
// @Id deleteZone
// @Summary Deletes Zone
// @Accept json
// @Produce json
// @Tags Zone
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /zone/{id} [delete]
func (e *ZoneController) deleteZone(c echo.Context) error {
	zoneId, err := strconv.Atoi(c.Param("zone"))
	if err != nil {
		e.logger.Error(err)
	}

	zone := new(models.Zone)
	err = e.db.Get(models.Zone{}, c).Model(&models.Zone{}).First(&zone, zoneId).Error
	if err != nil || zone.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Zone{}, c).Model(&models.Zone{}).Delete(&zone).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getZonesBulk godoc
// @Id getZonesBulk
// @Summary Gets Zones in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Zone
// @Success 200 {array} models.Zone
// @Failure 500 {string} string "Bad query request"
// @Router /zones/bulk [post]
func (e *ZoneController) getZonesBulk(c echo.Context) error {
	var results []models.Zone

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.Zone{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
