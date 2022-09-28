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

type SharedTaskController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSharedTaskController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SharedTaskController {
	return &SharedTaskController{
		db:	 db,
		logger: logger,
	}
}

func (e *SharedTaskController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "shared_task/:id", e.getSharedTask, nil),
		routes.RegisterRoute(http.MethodGet, "shared_tasks", e.listSharedTasks, nil),
		routes.RegisterRoute(http.MethodPut, "shared_task", e.createSharedTask, nil),
		routes.RegisterRoute(http.MethodDelete, "shared_task/:id", e.deleteSharedTask, nil),
		routes.RegisterRoute(http.MethodPatch, "shared_task/:id", e.updateSharedTask, nil),
		routes.RegisterRoute(http.MethodPost, "shared_tasks/bulk", e.getSharedTasksBulk, nil),
	}
}

// listSharedTasks godoc
// @Id listSharedTasks
// @Summary Lists SharedTasks
// @Accept json
// @Produce json
// @Tags SharedTask
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTask
// @Failure 500 {string} string "Bad query request"
// @Router /shared_tasks [get]
func (e *SharedTaskController) listSharedTasks(c echo.Context) error {
	var results []models.SharedTask
	err := e.db.QueryContext(models.SharedTask{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSharedTask godoc
// @Id getSharedTask
// @Summary Gets SharedTask
// @Accept json
// @Produce json
// @Tags SharedTask
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTask
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task/{id} [get]
func (e *SharedTaskController) getSharedTask(c echo.Context) error {
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
	var result models.SharedTask
	query := e.db.QueryContext(models.SharedTask{}, c)
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

// updateSharedTask godoc
// @Id updateSharedTask
// @Summary Updates SharedTask
// @Accept json
// @Produce json
// @Tags SharedTask
// @Param id path int true "Id"
// @Param shared_task body models.SharedTask true "SharedTask"
// @Success 200 {array} models.SharedTask
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /shared_task/{id} [patch]
func (e *SharedTaskController) updateSharedTask(c echo.Context) error {
	request := new(models.SharedTask)
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
	var result models.SharedTask
	query := e.db.QueryContext(models.SharedTask{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.SharedTask{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createSharedTask godoc
// @Id createSharedTask
// @Summary Creates SharedTask
// @Accept json
// @Produce json
// @Param shared_task body models.SharedTask true "SharedTask"
// @Tags SharedTask
// @Success 200 {array} models.SharedTask
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /shared_task [put]
func (e *SharedTaskController) createSharedTask(c echo.Context) error {
	sharedTask := new(models.SharedTask)
	if err := c.Bind(sharedTask); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SharedTask{}, c).Model(&models.SharedTask{}).Create(&sharedTask).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, sharedTask)
}

// deleteSharedTask godoc
// @Id deleteSharedTask
// @Summary Deletes SharedTask
// @Accept json
// @Produce json
// @Tags SharedTask
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /shared_task/{id} [delete]
func (e *SharedTaskController) deleteSharedTask(c echo.Context) error {
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
	var result models.SharedTask
	query := e.db.QueryContext(models.SharedTask{}, c)
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

// getSharedTasksBulk godoc
// @Id getSharedTasksBulk
// @Summary Gets SharedTasks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SharedTask
// @Success 200 {array} models.SharedTask
// @Failure 500 {string} string "Bad query request"
// @Router /shared_tasks/bulk [post]
func (e *SharedTaskController) getSharedTasksBulk(c echo.Context) error {
	var results []models.SharedTask

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

	err := e.db.QueryContext(models.SharedTask{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
