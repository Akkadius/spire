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

type CompletedTaskController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCompletedTaskController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CompletedTaskController {
	return &CompletedTaskController{
		db:	 db,
		logger: logger,
	}
}

func (e *CompletedTaskController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "completed_task/:charid", e.getCompletedTask, nil),
		routes.RegisterRoute(http.MethodGet, "completed_tasks", e.listCompletedTasks, nil),
		routes.RegisterRoute(http.MethodPut, "completed_task", e.createCompletedTask, nil),
		routes.RegisterRoute(http.MethodDelete, "completed_task/:charid", e.deleteCompletedTask, nil),
		routes.RegisterRoute(http.MethodPatch, "completed_task/:charid", e.updateCompletedTask, nil),
		routes.RegisterRoute(http.MethodPost, "completed_tasks/bulk", e.getCompletedTasksBulk, nil),
	}
}

// listCompletedTasks godoc
// @Id listCompletedTasks
// @Summary Lists CompletedTasks
// @Accept json
// @Produce json
// @Tags CompletedTask
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedTask
// @Failure 500 {string} string "Bad query request"
// @Router /completed_tasks [get]
func (e *CompletedTaskController) listCompletedTasks(c echo.Context) error {
	var results []models.CompletedTask
	err := e.db.QueryContext(models.CompletedTask{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCompletedTask godoc
// @Id getCompletedTask
// @Summary Gets CompletedTask
// @Accept json
// @Produce json
// @Tags CompletedTask
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedTask
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /completed_task/{id} [get]
func (e *CompletedTaskController) getCompletedTask(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [completedtime] position [2] type [int]
	if len(c.QueryParam("completedtime")) > 0 {
		completedtimeParam, err := strconv.Atoi(c.QueryParam("completedtime"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [completedtime] err [%s]", err.Error())})
		}

		params = append(params, completedtimeParam)
		keys = append(keys, "completedtime = ?")
	}

	// key param [taskid] position [3] type [int]
	if len(c.QueryParam("taskid")) > 0 {
		taskidParam, err := strconv.Atoi(c.QueryParam("taskid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [taskid] err [%s]", err.Error())})
		}

		params = append(params, taskidParam)
		keys = append(keys, "taskid = ?")
	}

	// key param [activityid] position [4] type [int]
	if len(c.QueryParam("activityid")) > 0 {
		activityidParam, err := strconv.Atoi(c.QueryParam("activityid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activityid] err [%s]", err.Error())})
		}

		params = append(params, activityidParam)
		keys = append(keys, "activityid = ?")
	}

	// query builder
	var result models.CompletedTask
	query := e.db.QueryContext(models.CompletedTask{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Charid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCompletedTask godoc
// @Id updateCompletedTask
// @Summary Updates CompletedTask
// @Accept json
// @Produce json
// @Tags CompletedTask
// @Param id path int true "Id"
// @Param completed_task body models.CompletedTask true "CompletedTask"
// @Success 200 {array} models.CompletedTask
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /completed_task/{id} [patch]
func (e *CompletedTaskController) updateCompletedTask(c echo.Context) error {
	request := new(models.CompletedTask)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [completedtime] position [2] type [int]
	if len(c.QueryParam("completedtime")) > 0 {
		completedtimeParam, err := strconv.Atoi(c.QueryParam("completedtime"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [completedtime] err [%s]", err.Error())})
		}

		params = append(params, completedtimeParam)
		keys = append(keys, "completedtime = ?")
	}

	// key param [taskid] position [3] type [int]
	if len(c.QueryParam("taskid")) > 0 {
		taskidParam, err := strconv.Atoi(c.QueryParam("taskid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [taskid] err [%s]", err.Error())})
		}

		params = append(params, taskidParam)
		keys = append(keys, "taskid = ?")
	}

	// key param [activityid] position [4] type [int]
	if len(c.QueryParam("activityid")) > 0 {
		activityidParam, err := strconv.Atoi(c.QueryParam("activityid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activityid] err [%s]", err.Error())})
		}

		params = append(params, activityidParam)
		keys = append(keys, "activityid = ?")
	}

	// query builder
	var result models.CompletedTask
	query := e.db.QueryContext(models.CompletedTask{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CompletedTask{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCompletedTask godoc
// @Id createCompletedTask
// @Summary Creates CompletedTask
// @Accept json
// @Produce json
// @Param completed_task body models.CompletedTask true "CompletedTask"
// @Tags CompletedTask
// @Success 200 {array} models.CompletedTask
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /completed_task [put]
func (e *CompletedTaskController) createCompletedTask(c echo.Context) error {
	completedTask := new(models.CompletedTask)
	if err := c.Bind(completedTask); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CompletedTask{}, c).Model(&models.CompletedTask{}).Create(&completedTask).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, completedTask)
}

// deleteCompletedTask godoc
// @Id deleteCompletedTask
// @Summary Deletes CompletedTask
// @Accept json
// @Produce json
// @Tags CompletedTask
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /completed_task/{id} [delete]
func (e *CompletedTaskController) deleteCompletedTask(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [completedtime] position [2] type [int]
	if len(c.QueryParam("completedtime")) > 0 {
		completedtimeParam, err := strconv.Atoi(c.QueryParam("completedtime"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [completedtime] err [%s]", err.Error())})
		}

		params = append(params, completedtimeParam)
		keys = append(keys, "completedtime = ?")
	}

	// key param [taskid] position [3] type [int]
	if len(c.QueryParam("taskid")) > 0 {
		taskidParam, err := strconv.Atoi(c.QueryParam("taskid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [taskid] err [%s]", err.Error())})
		}

		params = append(params, taskidParam)
		keys = append(keys, "taskid = ?")
	}

	// key param [activityid] position [4] type [int]
	if len(c.QueryParam("activityid")) > 0 {
		activityidParam, err := strconv.Atoi(c.QueryParam("activityid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [activityid] err [%s]", err.Error())})
		}

		params = append(params, activityidParam)
		keys = append(keys, "activityid = ?")
	}

	// query builder
	var result models.CompletedTask
	query := e.db.QueryContext(models.CompletedTask{}, c)
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

// getCompletedTasksBulk godoc
// @Id getCompletedTasksBulk
// @Summary Gets CompletedTasks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CompletedTask
// @Success 200 {array} models.CompletedTask
// @Failure 500 {string} string "Bad query request"
// @Router /completed_tasks/bulk [post]
func (e *CompletedTaskController) getCompletedTasksBulk(c echo.Context) error {
	var results []models.CompletedTask

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

	err := e.db.QueryContext(models.CompletedTask{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
