package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type ServerScheduledEventController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewServerScheduledEventController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *ServerScheduledEventController {
	return &ServerScheduledEventController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *ServerScheduledEventController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "server_scheduled_event/:id", e.getServerScheduledEvent, nil),
		routes.RegisterRoute(http.MethodGet, "server_scheduled_events", e.listServerScheduledEvents, nil),
		routes.RegisterRoute(http.MethodGet, "server_scheduled_events/count", e.getServerScheduledEventsCount, nil),
		routes.RegisterRoute(http.MethodPut, "server_scheduled_event", e.createServerScheduledEvent, nil),
		routes.RegisterRoute(http.MethodDelete, "server_scheduled_event/:id", e.deleteServerScheduledEvent, nil),
		routes.RegisterRoute(http.MethodPatch, "server_scheduled_event/:id", e.updateServerScheduledEvent, nil),
		routes.RegisterRoute(http.MethodPost, "server_scheduled_events/bulk", e.getServerScheduledEventsBulk, nil),
	}
}

// listServerScheduledEvents godoc
// @Id listServerScheduledEvents
// @Summary Lists ServerScheduledEvents
// @Accept json
// @Produce json
// @Tags ServerScheduledEvent
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ServerScheduledEvent
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /server_scheduled_event/{id} [get]
func (e *ServerScheduledEventController) getServerScheduledEvent(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.ServerScheduledEvent
	query := e.db.QueryContext(models.ServerScheduledEvent{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
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
	request := new(models.ServerScheduledEvent)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.ServerScheduledEvent
	query := e.db.QueryContext(models.ServerScheduledEvent{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [ServerScheduledEvent] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.ServerScheduledEvent{}, c).Model(&models.ServerScheduledEvent{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&serverScheduledEvent).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.ServerScheduledEvent{}, serverScheduledEvent)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [ServerScheduledEvent] [%v] data [%v]", serverScheduledEvent.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, serverScheduledEvent)
}

// deleteServerScheduledEvent godoc
// @Id deleteServerScheduledEvent
// @Summary Deletes ServerScheduledEvent
// @Accept json
// @Produce json
// @Tags ServerScheduledEvent
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /server_scheduled_event/{id} [delete]
func (e *ServerScheduledEventController) deleteServerScheduledEvent(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.ServerScheduledEvent
	query := e.db.QueryContext(models.ServerScheduledEvent{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [ServerScheduledEvent] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getServerScheduledEventsCount godoc
// @Id getServerScheduledEventsCount
// @Summary Counts ServerScheduledEvents
// @Accept json
// @Produce json
// @Tags ServerScheduledEvent
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ServerScheduledEvent
// @Failure 500 {string} string "Bad query request"
// @Router /server_scheduled_events/count [get]
func (e *ServerScheduledEventController) getServerScheduledEventsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.ServerScheduledEvent{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}