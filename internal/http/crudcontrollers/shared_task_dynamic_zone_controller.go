package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type SharedTaskDynamicZoneController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSharedTaskDynamicZoneController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SharedTaskDynamicZoneController {
	return &SharedTaskDynamicZoneController{
		db:	    db,
		logger: logger,
	}
}

func (e *SharedTaskDynamicZoneController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "shared_task_dynamic_zone/:sharedTaskId", e.getSharedTaskDynamicZone, nil),
		routes.RegisterRoute(http.MethodGet, "shared_task_dynamic_zones", e.listSharedTaskDynamicZones, nil),
		routes.RegisterRoute(http.MethodPut, "shared_task_dynamic_zone", e.createSharedTaskDynamicZone, nil),
		routes.RegisterRoute(http.MethodDelete, "shared_task_dynamic_zone/:sharedTaskId", e.deleteSharedTaskDynamicZone, nil),
		routes.RegisterRoute(http.MethodPatch, "shared_task_dynamic_zone/:sharedTaskId", e.updateSharedTaskDynamicZone, nil),
		routes.RegisterRoute(http.MethodPost, "shared_task_dynamic_zones/bulk", e.getSharedTaskDynamicZonesBulk, nil),
	}
}

// listSharedTaskDynamicZones godoc
// @Id listSharedTaskDynamicZones
// @Summary Lists SharedTaskDynamicZones
// @Accept json
// @Produce json
// @Tags SharedTaskDynamicZone
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTaskDynamicZone
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_dynamic_zones [get]
func (e *SharedTaskDynamicZoneController) listSharedTaskDynamicZones(c echo.Context) error {
	var results []models.SharedTaskDynamicZone
	err := e.db.QueryContext(models.SharedTaskDynamicZone{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSharedTaskDynamicZone godoc
// @Id getSharedTaskDynamicZone
// @Summary Gets SharedTaskDynamicZone
// @Accept json
// @Produce json
// @Tags SharedTaskDynamicZone
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTaskDynamicZone
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_dynamic_zone/{id} [get]
func (e *SharedTaskDynamicZoneController) getSharedTaskDynamicZone(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SharedTaskId]"})
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [dynamic_zone_id] position [2] type [int]
	if len(c.QueryParam("dynamic_zone_id")) > 0 {
		dynamicZoneIdParam, err := strconv.Atoi(c.QueryParam("dynamic_zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [dynamic_zone_id] err [%s]", err.Error())})
		}

		params = append(params, dynamicZoneIdParam)
		keys = append(keys, "dynamic_zone_id = ?")
	}

	// query builder
	var result models.SharedTaskDynamicZone
	query := e.db.QueryContext(models.SharedTaskDynamicZone{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.SharedTaskId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSharedTaskDynamicZone godoc
// @Id updateSharedTaskDynamicZone
// @Summary Updates SharedTaskDynamicZone
// @Accept json
// @Produce json
// @Tags SharedTaskDynamicZone
// @Param id path int true "Id"
// @Param shared_task_dynamic_zone body models.SharedTaskDynamicZone true "SharedTaskDynamicZone"
// @Success 200 {array} models.SharedTaskDynamicZone
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /shared_task_dynamic_zone/{id} [patch]
func (e *SharedTaskDynamicZoneController) updateSharedTaskDynamicZone(c echo.Context) error {
	request := new(models.SharedTaskDynamicZone)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SharedTaskId]"})
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [dynamic_zone_id] position [2] type [int]
	if len(c.QueryParam("dynamic_zone_id")) > 0 {
		dynamicZoneIdParam, err := strconv.Atoi(c.QueryParam("dynamic_zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [dynamic_zone_id] err [%s]", err.Error())})
		}

		params = append(params, dynamicZoneIdParam)
		keys = append(keys, "dynamic_zone_id = ?")
	}

	// query builder
	var result models.SharedTaskDynamicZone
	query := e.db.QueryContext(models.SharedTaskDynamicZone{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createSharedTaskDynamicZone godoc
// @Id createSharedTaskDynamicZone
// @Summary Creates SharedTaskDynamicZone
// @Accept json
// @Produce json
// @Param shared_task_dynamic_zone body models.SharedTaskDynamicZone true "SharedTaskDynamicZone"
// @Tags SharedTaskDynamicZone
// @Success 200 {array} models.SharedTaskDynamicZone
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /shared_task_dynamic_zone [put]
func (e *SharedTaskDynamicZoneController) createSharedTaskDynamicZone(c echo.Context) error {
	sharedTaskDynamicZone := new(models.SharedTaskDynamicZone)
	if err := c.Bind(sharedTaskDynamicZone); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SharedTaskDynamicZone{}, c).Model(&models.SharedTaskDynamicZone{}).Create(&sharedTaskDynamicZone).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, sharedTaskDynamicZone)
}

// deleteSharedTaskDynamicZone godoc
// @Id deleteSharedTaskDynamicZone
// @Summary Deletes SharedTaskDynamicZone
// @Accept json
// @Produce json
// @Tags SharedTaskDynamicZone
// @Param id path int true "sharedTaskId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /shared_task_dynamic_zone/{id} [delete]
func (e *SharedTaskDynamicZoneController) deleteSharedTaskDynamicZone(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [dynamic_zone_id] position [2] type [int]
	if len(c.QueryParam("dynamic_zone_id")) > 0 {
		dynamicZoneIdParam, err := strconv.Atoi(c.QueryParam("dynamic_zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [dynamic_zone_id] err [%s]", err.Error())})
		}

		params = append(params, dynamicZoneIdParam)
		keys = append(keys, "dynamic_zone_id = ?")
	}

	// query builder
	var result models.SharedTaskDynamicZone
	query := e.db.QueryContext(models.SharedTaskDynamicZone{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = query.Limit(10000).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSharedTaskDynamicZonesBulk godoc
// @Id getSharedTaskDynamicZonesBulk
// @Summary Gets SharedTaskDynamicZones in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SharedTaskDynamicZone
// @Success 200 {array} models.SharedTaskDynamicZone
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_dynamic_zones/bulk [post]
func (e *SharedTaskDynamicZoneController) getSharedTaskDynamicZonesBulk(c echo.Context) error {
	var results []models.SharedTaskDynamicZone

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.SharedTaskDynamicZone{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
