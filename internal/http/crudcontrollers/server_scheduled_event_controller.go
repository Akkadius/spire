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

type ServerScheduledEventController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewServerScheduledEventController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ServerScheduledEventController {
	return &ServerScheduledEventController{
		db:     db,
		logger: logger,
	}
}

func (e *ServerScheduledEventController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "server_scheduled_event/:server_scheduled_event", e.deleteServerScheduledEvent, nil),
		routes.RegisterRoute(http.MethodGet, "server_scheduled_event/:server_scheduled_event", e.getServerScheduledEvent, nil),
		routes.RegisterRoute(http.MethodGet, "server_scheduled_events", e.listServerScheduledEvents, nil),
		routes.RegisterRoute(http.MethodPost, "server_scheduled_events/bulk", e.getServerScheduledEventsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "server_scheduled_event/:server_scheduled_event", e.updateServerScheduledEvent, nil),
		routes.RegisterRoute(http.MethodPut, "server_scheduled_event", e.createServerScheduledEvent, nil),
	}
}

// listServerScheduledEvents godoc
// @Id listServerScheduledEvents
// @Summary Lists ServerScheduledEvents
// @Accept json
// @Produce json
// @Tags ServerScheduledEvent
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ServerScheduledEvent
// @Failure 500 {string} string "Bad query request"
// @Router /server_scheduled_events [get]
func (e *ServerScheduledEventController) listServerScheduledEvents(c echo.Context) error {
	var results []models.ServerScheduledEvent
	err := e.db.QueryContext(models.ServerScheduledEvent{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getServerScheduledEvent godoc
// @Id getServerScheduledEvent
// @Summary Gets ServerScheduledEvent
// @Accept json
// @Produce json
// @Tags ServerScheduledEvent
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ServerScheduledEvent
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /server_scheduled_event/{id} [get]
func (e *ServerScheduledEventController) getServerScheduledEvent(c echo.Context) error {
	serverScheduledEventId, err := strconv.Atoi(c.Param("server_scheduled_event"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.ServerScheduledEvent
	err = e.db.QueryContext(models.ServerScheduledEvent{}, c).First(&result, serverScheduledEventId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateServerScheduledEvent godoc
// @Id updateServerScheduledEvent
// @Summary Updates ServerScheduledEvent
// @Accept json
// @Produce json
// @Tags ServerScheduledEvent
// @Param id path int true "Id"
// @Param server_scheduled_event body models.ServerScheduledEvent true "ServerScheduledEvent"
// @Success 200 {array} models.ServerScheduledEvent
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /server_scheduled_event/{id} [patch]
func (e *ServerScheduledEventController) updateServerScheduledEvent(c echo.Context) error {
	serverScheduledEvent := new(models.ServerScheduledEvent)
	if err := c.Bind(serverScheduledEvent); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.ServerScheduledEvent{}
	err := e.db.Get(models.ServerScheduledEvent{}, c).Model(&models.ServerScheduledEvent{}).First(&entity, serverScheduledEvent.ID).Error
	if err != nil || serverScheduledEvent.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ServerScheduledEvent{}, c).Model(&entity).Select("*").Updates(&serverScheduledEvent).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, serverScheduledEvent)
}

// createServerScheduledEvent godoc
// @Id createServerScheduledEvent
// @Summary Creates ServerScheduledEvent
// @Accept json
// @Produce json
// @Param server_scheduled_event body models.ServerScheduledEvent true "ServerScheduledEvent"
// @Tags ServerScheduledEvent
// @Success 200 {array} models.ServerScheduledEvent
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /server_scheduled_event [put]
func (e *ServerScheduledEventController) createServerScheduledEvent(c echo.Context) error {
	serverScheduledEvent := new(models.ServerScheduledEvent)
	if err := c.Bind(serverScheduledEvent); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.ServerScheduledEvent{}, c).Model(&models.ServerScheduledEvent{}).Create(&serverScheduledEvent).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, serverScheduledEvent)
}

// deleteServerScheduledEvent godoc
// @Id deleteServerScheduledEvent
// @Summary Deletes ServerScheduledEvent
// @Accept json
// @Produce json
// @Tags ServerScheduledEvent
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /server_scheduled_event/{id} [delete]
func (e *ServerScheduledEventController) deleteServerScheduledEvent(c echo.Context) error {
	serverScheduledEventId, err := strconv.Atoi(c.Param("server_scheduled_event"))
	if err != nil {
		e.logger.Error(err)
	}

	serverScheduledEvent := new(models.ServerScheduledEvent)
	err = e.db.Get(models.ServerScheduledEvent{}, c).Model(&models.ServerScheduledEvent{}).First(&serverScheduledEvent, serverScheduledEventId).Error
	if err != nil || serverScheduledEvent.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ServerScheduledEvent{}, c).Model(&models.ServerScheduledEvent{}).Delete(&serverScheduledEvent).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getServerScheduledEventsBulk godoc
// @Id getServerScheduledEventsBulk
// @Summary Gets ServerScheduledEvents in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ServerScheduledEvent
// @Success 200 {array} models.ServerScheduledEvent
// @Failure 500 {string} string "Bad query request"
// @Router /server_scheduled_events/bulk [post]
func (e *ServerScheduledEventController) getServerScheduledEventsBulk(c echo.Context) error {
	var results []models.ServerScheduledEvent

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

	err := e.db.QueryContext(models.ServerScheduledEvent{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
