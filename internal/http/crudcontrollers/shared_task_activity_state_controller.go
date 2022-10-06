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

type SharedTaskActivityStateController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSharedTaskActivityStateController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SharedTaskActivityStateController {
	return &SharedTaskActivityStateController{
		db:	    db,
		logger: logger,
	}
}

func (e *SharedTaskActivityStateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "shared_task_activity_state/:sharedTaskId", e.getSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodGet, "shared_task_activity_states", e.listSharedTaskActivityStates, nil),
		routes.RegisterRoute(http.MethodPut, "shared_task_activity_state", e.createSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodDelete, "shared_task_activity_state/:sharedTaskId", e.deleteSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodPatch, "shared_task_activity_state/:sharedTaskId", e.updateSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodPost, "shared_task_activity_states/bulk", e.getSharedTaskActivityStatesBulk, nil),
	}
}

// listSharedTaskActivityStates godoc
// @Id listSharedTaskActivityStates
// @Summary Lists SharedTaskActivityStates
// @Accept json
// @Produce json
// @Tags SharedTaskActivityState
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTaskActivityState
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_activity_states [get]
func (e *SharedTaskActivityStateController) listSharedTaskActivityStates(c echo.Context) error {
	var results []models.SharedTaskActivityState
	err := e.db.QueryContext(models.SharedTaskActivityState{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSharedTaskActivityState godoc
// @Id getSharedTaskActivityState
// @Summary Gets SharedTaskActivityState
// @Accept json
// @Produce json
// @Tags SharedTaskActivityState
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTaskActivityState
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_activity_state/{id} [get]
func (e *SharedTaskActivityStateController) getSharedTaskActivityState(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SharedTaskId]"})
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [activity_id] position [2] type [int]
	if len(c.QueryParam("activity_id")) > 0 {
		activityIdParam, err := strconv.Atoi(c.QueryParam("activity_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activity_id] err [%s]", err.Error())})
		}

		params = append(params, activityIdParam)
		keys = append(keys, "activity_id = ?")
	}

	// query builder
	var result models.SharedTaskActivityState
	query := e.db.QueryContext(models.SharedTaskActivityState{}, c)
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

// updateSharedTaskActivityState godoc
// @Id updateSharedTaskActivityState
// @Summary Updates SharedTaskActivityState
// @Accept json
// @Produce json
// @Tags SharedTaskActivityState
// @Param id path int true "Id"
// @Param shared_task_activity_state body models.SharedTaskActivityState true "SharedTaskActivityState"
// @Success 200 {array} models.SharedTaskActivityState
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /shared_task_activity_state/{id} [patch]
func (e *SharedTaskActivityStateController) updateSharedTaskActivityState(c echo.Context) error {
	request := new(models.SharedTaskActivityState)
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

	// key param [activity_id] position [2] type [int]
	if len(c.QueryParam("activity_id")) > 0 {
		activityIdParam, err := strconv.Atoi(c.QueryParam("activity_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activity_id] err [%s]", err.Error())})
		}

		params = append(params, activityIdParam)
		keys = append(keys, "activity_id = ?")
	}

	// query builder
	var result models.SharedTaskActivityState
	query := e.db.QueryContext(models.SharedTaskActivityState{}, c)
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

// createSharedTaskActivityState godoc
// @Id createSharedTaskActivityState
// @Summary Creates SharedTaskActivityState
// @Accept json
// @Produce json
// @Param shared_task_activity_state body models.SharedTaskActivityState true "SharedTaskActivityState"
// @Tags SharedTaskActivityState
// @Success 200 {array} models.SharedTaskActivityState
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /shared_task_activity_state [put]
func (e *SharedTaskActivityStateController) createSharedTaskActivityState(c echo.Context) error {
	sharedTaskActivityState := new(models.SharedTaskActivityState)
	if err := c.Bind(sharedTaskActivityState); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SharedTaskActivityState{}, c).Model(&models.SharedTaskActivityState{}).Create(&sharedTaskActivityState).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, sharedTaskActivityState)
}

// deleteSharedTaskActivityState godoc
// @Id deleteSharedTaskActivityState
// @Summary Deletes SharedTaskActivityState
// @Accept json
// @Produce json
// @Tags SharedTaskActivityState
// @Param id path int true "sharedTaskId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /shared_task_activity_state/{id} [delete]
func (e *SharedTaskActivityStateController) deleteSharedTaskActivityState(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [activity_id] position [2] type [int]
	if len(c.QueryParam("activity_id")) > 0 {
		activityIdParam, err := strconv.Atoi(c.QueryParam("activity_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activity_id] err [%s]", err.Error())})
		}

		params = append(params, activityIdParam)
		keys = append(keys, "activity_id = ?")
	}

	// query builder
	var result models.SharedTaskActivityState
	query := e.db.QueryContext(models.SharedTaskActivityState{}, c)
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

// getSharedTaskActivityStatesBulk godoc
// @Id getSharedTaskActivityStatesBulk
// @Summary Gets SharedTaskActivityStates in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SharedTaskActivityState
// @Success 200 {array} models.SharedTaskActivityState
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_activity_states/bulk [post]
func (e *SharedTaskActivityStateController) getSharedTaskActivityStatesBulk(c echo.Context) error {
	var results []models.SharedTaskActivityState

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

	err := e.db.QueryContext(models.SharedTaskActivityState{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
