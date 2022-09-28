package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TaskActivityController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTaskActivityController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TaskActivityController {
	return &TaskActivityController{
		db:	 db,
		logger: logger,
	}
}

func (e *TaskActivityController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "task_activity/:taskid", e.getTaskActivity, nil),
		routes.RegisterRoute(http.MethodGet, "task_activities", e.listTaskActivities, nil),
		routes.RegisterRoute(http.MethodPut, "task_activity", e.createTaskActivity, nil),
		routes.RegisterRoute(http.MethodDelete, "task_activity/:taskid", e.deleteTaskActivity, nil),
		routes.RegisterRoute(http.MethodPatch, "task_activity/:taskid", e.updateTaskActivity, nil),
		routes.RegisterRoute(http.MethodPost, "task_activities/bulk", e.getTaskActivitiesBulk, nil),
	}
}

// listTaskActivities godoc
// @Id listTaskActivities
// @Summary Lists TaskActivities
// @Accept json
// @Produce json
// @Tags TaskActivity
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TaskActivity
// @Failure 500 {string} string "Bad query request"
// @Router /task_activities [get]
func (e *TaskActivityController) listTaskActivities(c echo.Context) error {
	var results []models.TaskActivity
	err := e.db.QueryContext(models.TaskActivity{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTaskActivity godoc
// @Id getTaskActivity
// @Summary Gets TaskActivity
// @Accept json
// @Produce json
// @Tags TaskActivity
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TaskActivity
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /task_activity/{id} [get]
func (e *TaskActivityController) getTaskActivity(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	taskid, err := strconv.Atoi(c.Param("taskid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Taskid]"})
	}
	params = append(params, taskid)
	keys = append(keys, "taskid = ?")

	// key param [activityid] position [2] type [int]
	if len(c.QueryParam("activityid")) > 0 {
		activityidParam, err := strconv.Atoi(c.QueryParam("activityid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activityid] err [%s]", err.Error())})
		}

		params = append(params, activityidParam)
		keys = append(keys, "activityid = ?")
	}

	// query builder
	var result models.TaskActivity
	query := e.db.QueryContext(models.TaskActivity{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Taskid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTaskActivity godoc
// @Id updateTaskActivity
// @Summary Updates TaskActivity
// @Accept json
// @Produce json
// @Tags TaskActivity
// @Param id path int true "Id"
// @Param task_activity body models.TaskActivity true "TaskActivity"
// @Success 200 {array} models.TaskActivity
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /task_activity/{id} [patch]
func (e *TaskActivityController) updateTaskActivity(c echo.Context) error {
	request := new(models.TaskActivity)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	taskid, err := strconv.Atoi(c.Param("taskid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Taskid]"})
	}
	params = append(params, taskid)
	keys = append(keys, "taskid = ?")

	// key param [activityid] position [2] type [int]
	if len(c.QueryParam("activityid")) > 0 {
		activityidParam, err := strconv.Atoi(c.QueryParam("activityid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activityid] err [%s]", err.Error())})
		}

		params = append(params, activityidParam)
		keys = append(keys, "activityid = ?")
	}

	// query builder
	var result models.TaskActivity
	query := e.db.QueryContext(models.TaskActivity{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.TaskActivity{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createTaskActivity godoc
// @Id createTaskActivity
// @Summary Creates TaskActivity
// @Accept json
// @Produce json
// @Param task_activity body models.TaskActivity true "TaskActivity"
// @Tags TaskActivity
// @Success 200 {array} models.TaskActivity
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /task_activity [put]
func (e *TaskActivityController) createTaskActivity(c echo.Context) error {
	taskActivity := new(models.TaskActivity)
	if err := c.Bind(taskActivity); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.TaskActivity{}, c).Model(&models.TaskActivity{}).Create(&taskActivity).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, taskActivity)
}

// deleteTaskActivity godoc
// @Id deleteTaskActivity
// @Summary Deletes TaskActivity
// @Accept json
// @Produce json
// @Tags TaskActivity
// @Param id path int true "taskid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /task_activity/{id} [delete]
func (e *TaskActivityController) deleteTaskActivity(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	taskid, err := strconv.Atoi(c.Param("taskid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, taskid)
	keys = append(keys, "taskid = ?")

	// key param [activityid] position [2] type [int]
	if len(c.QueryParam("activityid")) > 0 {
		activityidParam, err := strconv.Atoi(c.QueryParam("activityid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activityid] err [%s]", err.Error())})
		}

		params = append(params, activityidParam)
		keys = append(keys, "activityid = ?")
	}

	// query builder
	var result models.TaskActivity
	query := e.db.QueryContext(models.TaskActivity{}, c)
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

// getTaskActivitiesBulk godoc
// @Id getTaskActivitiesBulk
// @Summary Gets TaskActivities in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags TaskActivity
// @Success 200 {array} models.TaskActivity
// @Failure 500 {string} string "Bad query request"
// @Router /task_activities/bulk [post]
func (e *TaskActivityController) getTaskActivitiesBulk(c echo.Context) error {
	var results []models.TaskActivity

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

	err := e.db.QueryContext(models.TaskActivity{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
