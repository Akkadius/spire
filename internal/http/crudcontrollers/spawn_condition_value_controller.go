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

type SpawnConditionValueController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawnConditionValueController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpawnConditionValueController {
	return &SpawnConditionValueController{
		db:     db,
		logger: logger,
	}
}

func (e *SpawnConditionValueController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "spawn_condition_value/:spawn_condition_value", e.deleteSpawnConditionValue, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_condition_value/:spawn_condition_value", e.getSpawnConditionValue, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_condition_values", e.listSpawnConditionValues, nil),
		routes.RegisterRoute(http.MethodPost, "spawn_condition_values/bulk", e.getSpawnConditionValuesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "spawn_condition_value/:spawn_condition_value", e.updateSpawnConditionValue, nil),
		routes.RegisterRoute(http.MethodPut, "spawn_condition_value", e.createSpawnConditionValue, nil),
	}
}

// listSpawnConditionValues godoc
// @Id listSpawnConditionValues
// @Summary Lists SpawnConditionValues
// @Accept json
// @Produce json
// @Tags SpawnConditionValue
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpawnConditionValue
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_condition_values [get]
func (e *SpawnConditionValueController) listSpawnConditionValues(c echo.Context) error {
	var results []models.SpawnConditionValue
	err := e.db.QueryContext(models.SpawnConditionValue{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawnConditionValue godoc
// @Id getSpawnConditionValue
// @Summary Gets SpawnConditionValue
// @Accept json
// @Produce json
// @Tags SpawnConditionValue
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpawnConditionValue
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_condition_value/{id} [get]
func (e *SpawnConditionValueController) getSpawnConditionValue(c echo.Context) error {
	spawnConditionValueId, err := strconv.Atoi(c.Param("spawn_condition_value"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.SpawnConditionValue
	err = e.db.QueryContext(models.SpawnConditionValue{}, c).First(&result, spawnConditionValueId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSpawnConditionValue godoc
// @Id updateSpawnConditionValue
// @Summary Updates SpawnConditionValue
// @Accept json
// @Produce json
// @Tags SpawnConditionValue
// @Param id path int true "Id"
// @Param spawn_condition_value body models.SpawnConditionValue true "SpawnConditionValue"
// @Success 200 {array} models.SpawnConditionValue
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawn_condition_value/{id} [patch]
func (e *SpawnConditionValueController) updateSpawnConditionValue(c echo.Context) error {
	spawnConditionValue := new(models.SpawnConditionValue)
	if err := c.Bind(spawnConditionValue); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.SpawnConditionValue{}
	err := e.db.Get(models.SpawnConditionValue{}, c).Model(&models.SpawnConditionValue{}).First(&entity, spawnConditionValue.ID).Error
	if err != nil || spawnConditionValue.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.SpawnConditionValue{}, c).Model(&entity).Updates(&spawnConditionValue).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, spawnConditionValue)
}

// createSpawnConditionValue godoc
// @Id createSpawnConditionValue
// @Summary Creates SpawnConditionValue
// @Accept json
// @Produce json
// @Param spawn_condition_value body models.SpawnConditionValue true "SpawnConditionValue"
// @Tags SpawnConditionValue
// @Success 200 {array} models.SpawnConditionValue
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawn_condition_value [put]
func (e *SpawnConditionValueController) createSpawnConditionValue(c echo.Context) error {
	spawnConditionValue := new(models.SpawnConditionValue)
	if err := c.Bind(spawnConditionValue); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.SpawnConditionValue{}, c).Model(&models.SpawnConditionValue{}).Create(&spawnConditionValue).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, spawnConditionValue)
}

// deleteSpawnConditionValue godoc
// @Id deleteSpawnConditionValue
// @Summary Deletes SpawnConditionValue
// @Accept json
// @Produce json
// @Tags SpawnConditionValue
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawn_condition_value/{id} [delete]
func (e *SpawnConditionValueController) deleteSpawnConditionValue(c echo.Context) error {
	spawnConditionValueId, err := strconv.Atoi(c.Param("spawn_condition_value"))
	if err != nil {
		e.logger.Error(err)
	}

	spawnConditionValue := new(models.SpawnConditionValue)
	err = e.db.Get(models.SpawnConditionValue{}, c).Model(&models.SpawnConditionValue{}).First(&spawnConditionValue, spawnConditionValueId).Error
	if err != nil || spawnConditionValue.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.SpawnConditionValue{}, c).Model(&models.SpawnConditionValue{}).Delete(&spawnConditionValue).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpawnConditionValuesBulk godoc
// @Id getSpawnConditionValuesBulk
// @Summary Gets SpawnConditionValues in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpawnConditionValue
// @Success 200 {array} models.SpawnConditionValue
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_condition_values/bulk [post]
func (e *SpawnConditionValueController) getSpawnConditionValuesBulk(c echo.Context) error {
	var results []models.SpawnConditionValue

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

	err := e.db.QueryContext(models.SpawnConditionValue{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
