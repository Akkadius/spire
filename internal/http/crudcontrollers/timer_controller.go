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

type TimerController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTimerController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TimerController {
	return &TimerController{
		db:	 db,
		logger: logger,
	}
}

func (e *TimerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "timer/:charId", e.getTimer, nil),
		routes.RegisterRoute(http.MethodGet, "timers", e.listTimers, nil),
		routes.RegisterRoute(http.MethodPut, "timer", e.createTimer, nil),
		routes.RegisterRoute(http.MethodDelete, "timer/:charId", e.deleteTimer, nil),
		routes.RegisterRoute(http.MethodPatch, "timer/:charId", e.updateTimer, nil),
		routes.RegisterRoute(http.MethodPost, "timers/bulk", e.getTimersBulk, nil),
	}
}

// listTimers godoc
// @Id listTimers
// @Summary Lists Timers
// @Accept json
// @Produce json
// @Tags Timer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Timer
// @Failure 500 {string} string "Bad query request"
// @Router /timers [get]
func (e *TimerController) listTimers(c echo.Context) error {
	var results []models.Timer
	err := e.db.QueryContext(models.Timer{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTimer godoc
// @Id getTimer
// @Summary Gets Timer
// @Accept json
// @Produce json
// @Tags Timer
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Timer
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /timer/{id} [get]
func (e *TimerController) getTimer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [typeId] position [2] type [mediumint]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// query builder
	var result models.Timer
	query := e.db.QueryContext(models.Timer{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTimer godoc
// @Id updateTimer
// @Summary Updates Timer
// @Accept json
// @Produce json
// @Tags Timer
// @Param id path int true "Id"
// @Param timer body models.Timer true "Timer"
// @Success 200 {array} models.Timer
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /timer/{id} [patch]
func (e *TimerController) updateTimer(c echo.Context) error {
	request := new(models.Timer)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [typeId] position [2] type [mediumint]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// query builder
	var result models.Timer
	query := e.db.QueryContext(models.Timer{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Timer{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createTimer godoc
// @Id createTimer
// @Summary Creates Timer
// @Accept json
// @Produce json
// @Param timer body models.Timer true "Timer"
// @Tags Timer
// @Success 200 {array} models.Timer
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /timer [put]
func (e *TimerController) createTimer(c echo.Context) error {
	timer := new(models.Timer)
	if err := c.Bind(timer); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Timer{}, c).Model(&models.Timer{}).Create(&timer).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, timer)
}

// deleteTimer godoc
// @Id deleteTimer
// @Summary Deletes Timer
// @Accept json
// @Produce json
// @Tags Timer
// @Param id path int true "charId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /timer/{id} [delete]
func (e *TimerController) deleteTimer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [typeId] position [2] type [mediumint]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// query builder
	var result models.Timer
	query := e.db.QueryContext(models.Timer{}, c)
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

// getTimersBulk godoc
// @Id getTimersBulk
// @Summary Gets Timers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Timer
// @Success 200 {array} models.Timer
// @Failure 500 {string} string "Bad query request"
// @Router /timers/bulk [post]
func (e *TimerController) getTimersBulk(c echo.Context) error {
	var results []models.Timer

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

	err := e.db.QueryContext(models.Timer{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
