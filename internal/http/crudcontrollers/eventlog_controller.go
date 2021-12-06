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

type EventlogController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewEventlogController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *EventlogController {
	return &EventlogController{
		db:     db,
		logger: logger,
	}
}

func (e *EventlogController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "eventlog/:eventlog", e.deleteEventlog, nil),
		routes.RegisterRoute(http.MethodGet, "eventlog/:eventlog", e.getEventlog, nil),
		routes.RegisterRoute(http.MethodGet, "eventlogs", e.listEventlogs, nil),
		routes.RegisterRoute(http.MethodPost, "eventlogs/bulk", e.getEventlogsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "eventlog/:eventlog", e.updateEventlog, nil),
		routes.RegisterRoute(http.MethodPut, "eventlog", e.createEventlog, nil),
	}
}

// listEventlogs godoc
// @Id listEventlogs
// @Summary Lists Eventlogs
// @Accept json
// @Produce json
// @Tags Eventlog
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Eventlog
// @Failure 500 {string} string "Bad query request"
// @Router /eventlogs [get]
func (e *EventlogController) listEventlogs(c echo.Context) error {
	var results []models.Eventlog
	err := e.db.QueryContext(models.Eventlog{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getEventlog godoc
// @Id getEventlog
// @Summary Gets Eventlog
// @Accept json
// @Produce json
// @Tags Eventlog
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Eventlog
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /eventlog/{id} [get]
func (e *EventlogController) getEventlog(c echo.Context) error {
	eventlogId, err := strconv.Atoi(c.Param("eventlog"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Eventlog
	err = e.db.QueryContext(models.Eventlog{}, c).First(&result, eventlogId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateEventlog godoc
// @Id updateEventlog
// @Summary Updates Eventlog
// @Accept json
// @Produce json
// @Tags Eventlog
// @Param id path int true "Id"
// @Param eventlog body models.Eventlog true "Eventlog"
// @Success 200 {array} models.Eventlog
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /eventlog/{id} [patch]
func (e *EventlogController) updateEventlog(c echo.Context) error {
	eventlog := new(models.Eventlog)
	if err := c.Bind(eventlog); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Eventlog{}
	err := e.db.Get(models.Eventlog{}, c).Model(&models.Eventlog{}).First(&entity, eventlog.ID).Error
	if err != nil || eventlog.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Eventlog{}, c).Model(&entity).Select("*").Updates(&eventlog).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, eventlog)
}

// createEventlog godoc
// @Id createEventlog
// @Summary Creates Eventlog
// @Accept json
// @Produce json
// @Param eventlog body models.Eventlog true "Eventlog"
// @Tags Eventlog
// @Success 200 {array} models.Eventlog
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /eventlog [put]
func (e *EventlogController) createEventlog(c echo.Context) error {
	eventlog := new(models.Eventlog)
	if err := c.Bind(eventlog); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Eventlog{}, c).Model(&models.Eventlog{}).Create(&eventlog).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, eventlog)
}

// deleteEventlog godoc
// @Id deleteEventlog
// @Summary Deletes Eventlog
// @Accept json
// @Produce json
// @Tags Eventlog
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /eventlog/{id} [delete]
func (e *EventlogController) deleteEventlog(c echo.Context) error {
	eventlogId, err := strconv.Atoi(c.Param("eventlog"))
	if err != nil {
		e.logger.Error(err)
	}

	eventlog := new(models.Eventlog)
	err = e.db.Get(models.Eventlog{}, c).Model(&models.Eventlog{}).First(&eventlog, eventlogId).Error
	if err != nil || eventlog.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Eventlog{}, c).Model(&models.Eventlog{}).Delete(&eventlog).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getEventlogsBulk godoc
// @Id getEventlogsBulk
// @Summary Gets Eventlogs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Eventlog
// @Success 200 {array} models.Eventlog
// @Failure 500 {string} string "Bad query request"
// @Router /eventlogs/bulk [post]
func (e *EventlogController) getEventlogsBulk(c echo.Context) error {
	var results []models.Eventlog

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

	err := e.db.QueryContext(models.Eventlog{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
