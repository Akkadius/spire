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

type CompletedSharedTaskController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCompletedSharedTaskController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CompletedSharedTaskController {
	return &CompletedSharedTaskController{
		db:	 db,
		logger: logger,
	}
}

func (e *CompletedSharedTaskController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "completed_shared_task/:id", e.getCompletedSharedTask, nil),
		routes.RegisterRoute(http.MethodGet, "completed_shared_tasks", e.listCompletedSharedTasks, nil),
		routes.RegisterRoute(http.MethodPut, "completed_shared_task", e.createCompletedSharedTask, nil),
		routes.RegisterRoute(http.MethodDelete, "completed_shared_task/:id", e.deleteCompletedSharedTask, nil),
		routes.RegisterRoute(http.MethodPatch, "completed_shared_task/:id", e.updateCompletedSharedTask, nil),
		routes.RegisterRoute(http.MethodPost, "completed_shared_tasks/bulk", e.getCompletedSharedTasksBulk, nil),
	}
}

// listCompletedSharedTasks godoc
// @Id listCompletedSharedTasks
// @Summary Lists CompletedSharedTasks
// @Accept json
// @Produce json
// @Tags CompletedSharedTask
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedSharedTask
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_tasks [get]
func (e *CompletedSharedTaskController) listCompletedSharedTasks(c echo.Context) error {
	var results []models.CompletedSharedTask
	err := e.db.QueryContext(models.CompletedSharedTask{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCompletedSharedTask godoc
// @Id getCompletedSharedTask
// @Summary Gets CompletedSharedTask
// @Accept json
// @Produce json
// @Tags CompletedSharedTask
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedSharedTask
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_task/{id} [get]
func (e *CompletedSharedTaskController) getCompletedSharedTask(c echo.Context) error {
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
	var result models.CompletedSharedTask
	query := e.db.QueryContext(models.CompletedSharedTask{}, c)
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

// updateCompletedSharedTask godoc
// @Id updateCompletedSharedTask
// @Summary Updates CompletedSharedTask
// @Accept json
// @Produce json
// @Tags CompletedSharedTask
// @Param id path int true "Id"
// @Param completed_shared_task body models.CompletedSharedTask true "CompletedSharedTask"
// @Success 200 {array} models.CompletedSharedTask
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /completed_shared_task/{id} [patch]
func (e *CompletedSharedTaskController) updateCompletedSharedTask(c echo.Context) error {
	request := new(models.CompletedSharedTask)
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
	var result models.CompletedSharedTask
	query := e.db.QueryContext(models.CompletedSharedTask{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCompletedSharedTask godoc
// @Id createCompletedSharedTask
// @Summary Creates CompletedSharedTask
// @Accept json
// @Produce json
// @Param completed_shared_task body models.CompletedSharedTask true "CompletedSharedTask"
// @Tags CompletedSharedTask
// @Success 200 {array} models.CompletedSharedTask
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /completed_shared_task [put]
func (e *CompletedSharedTaskController) createCompletedSharedTask(c echo.Context) error {
	completedSharedTask := new(models.CompletedSharedTask)
	if err := c.Bind(completedSharedTask); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CompletedSharedTask{}, c).Model(&models.CompletedSharedTask{}).Create(&completedSharedTask).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, completedSharedTask)
}

// deleteCompletedSharedTask godoc
// @Id deleteCompletedSharedTask
// @Summary Deletes CompletedSharedTask
// @Accept json
// @Produce json
// @Tags CompletedSharedTask
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /completed_shared_task/{id} [delete]
func (e *CompletedSharedTaskController) deleteCompletedSharedTask(c echo.Context) error {
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
	var result models.CompletedSharedTask
	query := e.db.QueryContext(models.CompletedSharedTask{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.CompletedSharedTask{}, c).Model(&models.CompletedSharedTask{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCompletedSharedTasksBulk godoc
// @Id getCompletedSharedTasksBulk
// @Summary Gets CompletedSharedTasks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CompletedSharedTask
// @Success 200 {array} models.CompletedSharedTask
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_tasks/bulk [post]
func (e *CompletedSharedTaskController) getCompletedSharedTasksBulk(c echo.Context) error {
	var results []models.CompletedSharedTask

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

	err := e.db.QueryContext(models.CompletedSharedTask{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
