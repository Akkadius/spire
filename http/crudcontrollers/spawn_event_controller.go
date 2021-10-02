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

type SpawnEventController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawnEventController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpawnEventController {
	return &SpawnEventController{
		db:     db,
		logger: logger,
	}
}

func (e *SpawnEventController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "spawn_event/:spawn_event", e.deleteSpawnEvent, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_event/:spawn_event", e.getSpawnEvent, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_events", e.listSpawnEvents, nil),
		routes.RegisterRoute(http.MethodPost, "spawn_events/bulk", e.getSpawnEventsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "spawn_event/:spawn_event", e.updateSpawnEvent, nil),
		routes.RegisterRoute(http.MethodPut, "spawn_event", e.createSpawnEvent, nil),
	}
}

// listSpawnEvents godoc
// @Id listSpawnEvents
// @Summary Lists SpawnEvents
// @Accept json
// @Produce json
// @Tags SpawnEvent
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpawnEvent
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_events [get]
func (e *SpawnEventController) listSpawnEvents(c echo.Context) error {
	var results []models.SpawnEvent
	err := e.db.QueryContext(models.SpawnEvent{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawnEvent godoc
// @Id getSpawnEvent
// @Summary Gets SpawnEvent
// @Accept json
// @Produce json
// @Tags SpawnEvent
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpawnEvent
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_event/{id} [get]
func (e *SpawnEventController) getSpawnEvent(c echo.Context) error {
	spawnEventId, err := strconv.Atoi(c.Param("spawn_event"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.SpawnEvent
	err = e.db.QueryContext(models.SpawnEvent{}, c).First(&result, spawnEventId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSpawnEvent godoc
// @Id updateSpawnEvent
// @Summary Updates SpawnEvent
// @Accept json
// @Produce json
// @Tags SpawnEvent
// @Param id path int true "Id"
// @Param spawn_event body models.SpawnEvent true "SpawnEvent"
// @Success 200 {array} models.SpawnEvent
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawn_event/{id} [patch]
func (e *SpawnEventController) updateSpawnEvent(c echo.Context) error {
	spawnEvent := new(models.SpawnEvent)
	if err := c.Bind(spawnEvent); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.SpawnEvent{}
	err := e.db.Get(models.SpawnEvent{}, c).Model(&models.SpawnEvent{}).First(&entity, spawnEvent.ID).Error
	if err != nil || spawnEvent.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.SpawnEvent{}, c).Model(&entity).Updates(&spawnEvent).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, spawnEvent)
}

// createSpawnEvent godoc
// @Id createSpawnEvent
// @Summary Creates SpawnEvent
// @Accept json
// @Produce json
// @Param spawn_event body models.SpawnEvent true "SpawnEvent"
// @Tags SpawnEvent
// @Success 200 {array} models.SpawnEvent
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawn_event [put]
func (e *SpawnEventController) createSpawnEvent(c echo.Context) error {
	spawnEvent := new(models.SpawnEvent)
	if err := c.Bind(spawnEvent); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.SpawnEvent{}, c).Model(&models.SpawnEvent{}).Create(&spawnEvent).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, spawnEvent)
}

// deleteSpawnEvent godoc
// @Id deleteSpawnEvent
// @Summary Deletes SpawnEvent
// @Accept json
// @Produce json
// @Tags SpawnEvent
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawn_event/{id} [delete]
func (e *SpawnEventController) deleteSpawnEvent(c echo.Context) error {
	spawnEventId, err := strconv.Atoi(c.Param("spawn_event"))
	if err != nil {
		e.logger.Error(err)
	}

	spawnEvent := new(models.SpawnEvent)
	err = e.db.Get(models.SpawnEvent{}, c).Model(&models.SpawnEvent{}).First(&spawnEvent, spawnEventId).Error
	if err != nil || spawnEvent.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.SpawnEvent{}, c).Model(&models.SpawnEvent{}).Delete(&spawnEvent).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpawnEventsBulk godoc
// @Id getSpawnEventsBulk
// @Summary Gets SpawnEvents in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpawnEvent
// @Success 200 {array} models.SpawnEvent
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_events/bulk [post]
func (e *SpawnEventController) getSpawnEventsBulk(c echo.Context) error {
	var results []models.SpawnEvent

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

	err := e.db.QueryContext(models.SpawnEvent{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
