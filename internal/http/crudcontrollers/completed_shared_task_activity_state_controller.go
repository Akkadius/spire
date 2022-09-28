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

type CompletedSharedTaskActivityStateController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCompletedSharedTaskActivityStateController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CompletedSharedTaskActivityStateController {
	return &CompletedSharedTaskActivityStateController{
		db:	 db,
		logger: logger,
	}
}

func (e *CompletedSharedTaskActivityStateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "completed_shared_task_activity_state/:sharedTaskId", e.getCompletedSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodGet, "completed_shared_task_activity_states", e.listCompletedSharedTaskActivityStates, nil),
		routes.RegisterRoute(http.MethodPut, "completed_shared_task_activity_state", e.createCompletedSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodDelete, "completed_shared_task_activity_state/:sharedTaskId", e.deleteCompletedSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodPatch, "completed_shared_task_activity_state/:sharedTaskId", e.updateCompletedSharedTaskActivityState, nil),
		routes.RegisterRoute(http.MethodPost, "completed_shared_task_activity_states/bulk", e.getCompletedSharedTaskActivityStatesBulk, nil),
	}
}

// listCompletedSharedTaskActivityStates godoc
// @Id listCompletedSharedTaskActivityStates
// @Summary Lists CompletedSharedTaskActivityStates
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskActivityState
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedSharedTaskActivityState
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_task_activity_states [get]
func (e *CompletedSharedTaskActivityStateController) listCompletedSharedTaskActivityStates(c echo.Context) error {
	var results []models.CompletedSharedTaskActivityState
	err := e.db.QueryContext(models.CompletedSharedTaskActivityState{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCompletedSharedTaskActivityState godoc
// @Id getCompletedSharedTaskActivityState
// @Summary Gets CompletedSharedTaskActivityState
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskActivityState
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedSharedTaskActivityState
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_task_activity_state/{id} [get]
func (e *CompletedSharedTaskActivityStateController) getCompletedSharedTaskActivityState(c echo.Context) error {
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
	var result models.CompletedSharedTaskActivityState
	query := e.db.QueryContext(models.CompletedSharedTaskActivityState{}, c)
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

// updateCompletedSharedTaskActivityState godoc
// @Id updateCompletedSharedTaskActivityState
// @Summary Updates CompletedSharedTaskActivityState
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskActivityState
// @Param id path int true "Id"
// @Param completed_shared_task_activity_state body models.CompletedSharedTaskActivityState true "CompletedSharedTaskActivityState"
// @Success 200 {array} models.CompletedSharedTaskActivityState
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /completed_shared_task_activity_state/{id} [patch]
func (e *CompletedSharedTaskActivityStateController) updateCompletedSharedTaskActivityState(c echo.Context) error {
	request := new(models.CompletedSharedTaskActivityState)
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
	var result models.CompletedSharedTaskActivityState
	query := e.db.QueryContext(models.CompletedSharedTaskActivityState{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CompletedSharedTaskActivityState{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCompletedSharedTaskActivityState godoc
// @Id createCompletedSharedTaskActivityState
// @Summary Creates CompletedSharedTaskActivityState
// @Accept json
// @Produce json
// @Param completed_shared_task_activity_state body models.CompletedSharedTaskActivityState true "CompletedSharedTaskActivityState"
// @Tags CompletedSharedTaskActivityState
// @Success 200 {array} models.CompletedSharedTaskActivityState
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /completed_shared_task_activity_state [put]
func (e *CompletedSharedTaskActivityStateController) createCompletedSharedTaskActivityState(c echo.Context) error {
	completedSharedTaskActivityState := new(models.CompletedSharedTaskActivityState)
	if err := c.Bind(completedSharedTaskActivityState); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CompletedSharedTaskActivityState{}, c).Model(&models.CompletedSharedTaskActivityState{}).Create(&completedSharedTaskActivityState).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, completedSharedTaskActivityState)
}

// deleteCompletedSharedTaskActivityState godoc
// @Id deleteCompletedSharedTaskActivityState
// @Summary Deletes CompletedSharedTaskActivityState
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskActivityState
// @Param id path int true "sharedTaskId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /completed_shared_task_activity_state/{id} [delete]
func (e *CompletedSharedTaskActivityStateController) deleteCompletedSharedTaskActivityState(c echo.Context) error {
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
	var result models.CompletedSharedTaskActivityState
	query := e.db.QueryContext(models.CompletedSharedTaskActivityState{}, c)
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

// getCompletedSharedTaskActivityStatesBulk godoc
// @Id getCompletedSharedTaskActivityStatesBulk
// @Summary Gets CompletedSharedTaskActivityStates in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CompletedSharedTaskActivityState
// @Success 200 {array} models.CompletedSharedTaskActivityState
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_task_activity_states/bulk [post]
func (e *CompletedSharedTaskActivityStateController) getCompletedSharedTaskActivityStatesBulk(c echo.Context) error {
	var results []models.CompletedSharedTaskActivityState

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

	err := e.db.QueryContext(models.CompletedSharedTaskActivityState{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
