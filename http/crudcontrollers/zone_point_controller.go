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

type ZonePointController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewZonePointController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ZonePointController {
	return &ZonePointController{
		db:     db,
		logger: logger,
	}
}

func (e *ZonePointController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "zone_point/:zone_point", e.deleteZonePoint, nil),
		routes.RegisterRoute(http.MethodGet, "zone_point/:zone_point", e.getZonePoint, nil),
		routes.RegisterRoute(http.MethodGet, "zone_points", e.listZonePoints, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getZonePointsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "zone_point/:zone_point", e.updateZonePoint, nil),
		routes.RegisterRoute(http.MethodPut, "zone_point", e.createZonePoint, nil),
	}
}

// listZonePoints godoc
// @Id listZonePoints
// @Summary Lists ZonePoints
// @Accept json
// @Produce json
// @Tags ZonePoint
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZonePoint
// @Failure 500 {string} string "Bad query request"
// @Router /zone_points [get]
func (e *ZonePointController) listZonePoints(c echo.Context) error {
	var results []models.ZonePoint
	err := e.db.QueryContext(models.ZonePoint{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getZonePoint godoc
// @Id getZonePoint
// @Summary Gets ZonePoint
// @Accept json
// @Produce json
// @Tags ZonePoint
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZonePoint
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /zone_point/{id} [get]
func (e *ZonePointController) getZonePoint(c echo.Context) error {
	zonePointId, err := strconv.Atoi(c.Param("zone_point"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.ZonePoint
	err = e.db.QueryContext(models.ZonePoint{}, c).First(&result, zonePointId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateZonePoint godoc
// @Id updateZonePoint
// @Summary Updates ZonePoint
// @Accept json
// @Produce json
// @Tags ZonePoint
// @Param id path int true "Id"
// @Param zone_point body models.ZonePoint true "ZonePoint"
// @Success 200 {array} models.ZonePoint
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /zone_point/{id} [patch]
func (e *ZonePointController) updateZonePoint(c echo.Context) error {
	zonePoint := new(models.ZonePoint)
	if err := c.Bind(zonePoint); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.ZonePoint{}, c).Model(&models.ZonePoint{}).First(&models.ZonePoint{}, zonePoint.ID).Error
	if err != nil || zonePoint.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ZonePoint{}, c).Model(&models.ZonePoint{}).Updates(&zonePoint).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, zonePoint)
}

// createZonePoint godoc
// @Id createZonePoint
// @Summary Creates ZonePoint
// @Accept json
// @Produce json
// @Param zone_point body models.ZonePoint true "ZonePoint"
// @Tags ZonePoint
// @Success 200 {array} models.ZonePoint
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /zone_point [put]
func (e *ZonePointController) createZonePoint(c echo.Context) error {
	zonePoint := new(models.ZonePoint)
	if err := c.Bind(zonePoint); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.ZonePoint{}, c).Model(&models.ZonePoint{}).Create(&zonePoint).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, zonePoint)
}

// deleteZonePoint godoc
// @Id deleteZonePoint
// @Summary Deletes ZonePoint
// @Accept json
// @Produce json
// @Tags ZonePoint
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /zone_point/{id} [delete]
func (e *ZonePointController) deleteZonePoint(c echo.Context) error {
	zonePointId, err := strconv.Atoi(c.Param("zone_point"))
	if err != nil {
		e.logger.Error(err)
	}

	zonePoint := new(models.ZonePoint)
	err = e.db.Get(models.ZonePoint{}, c).Model(&models.ZonePoint{}).First(&zonePoint, zonePointId).Error
	if err != nil || zonePoint.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ZonePoint{}, c).Model(&models.ZonePoint{}).Delete(&zonePoint).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getZonePointsBulk godoc
// @Id getZonePointsBulk
// @Summary Gets ZonePoints in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ZonePoint
// @Success 200 {array} models.ZonePoint
// @Failure 500 {string} string "Bad query request"
// @Router /zone_points/bulk [post]
func (e *ZonePointController) getZonePointsBulk(c echo.Context) error {
	var results []models.ZonePoint

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

	err := e.db.QueryContext(models.ZonePoint{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
