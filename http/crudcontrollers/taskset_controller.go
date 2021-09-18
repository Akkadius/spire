package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TasksetController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTasksetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TasksetController {
	return &TasksetController {
		db:     db,
		logger: logger,
	}
}

func (e *TasksetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "taskset/:taskset", e.deleteTaskset, nil),
		routes.RegisterRoute(http.MethodGet, "taskset/:taskset", e.getTaskset, nil),
		routes.RegisterRoute(http.MethodGet, "tasksets", e.listTasksets, nil),
		routes.RegisterRoute(http.MethodPatch, "taskset/:taskset", e.updateTaskset, nil),
		routes.RegisterRoute(http.MethodPut, "taskset", e.createTaskset, nil),
	}
}

// listTasksets godoc
// @Id listTasksets
// @Summary Lists Tasksets
// @Accept json
// @Produce json
// @Tags Taskset
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Taskset
// @Failure 500 {string} string "Bad query request"
// @Router /tasksets [get]
func (e *TasksetController) listTasksets(c echo.Context) error {
	var results []models.Taskset
	err := e.db.QueryContext(models.Taskset{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTaskset godoc
// @Id getTaskset
// @Summary Gets Taskset
// @Accept json
// @Produce json
// @Tags Taskset
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Taskset
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /taskset/{id} [get]
func (e *TasksetController) getTaskset(c echo.Context) error {
	tasksetId, err := strconv.Atoi(c.Param("taskset"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Taskset
	err = e.db.QueryContext(models.Taskset{}, c).First(&result, tasksetId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTaskset godoc
// @Id updateTaskset
// @Summary Updates Taskset
// @Accept json
// @Produce json
// @Tags Taskset
// @Param id path int true "Id"
// @Param taskset body models.Taskset true "Taskset"
// @Success 200 {array} models.Taskset
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /taskset/{id} [patch]
func (e *TasksetController) updateTaskset(c echo.Context) error {
	taskset := new(models.Taskset)
	if err := c.Bind(taskset); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Taskset{}, c).Model(&models.Taskset{}).First(&models.Taskset{}, taskset.ID).Error
	if err != nil || taskset.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Taskset{}, c).Model(&models.Taskset{}).Updates(&taskset).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, taskset)
}

// createTaskset godoc
// @Id createTaskset
// @Summary Creates Taskset
// @Accept json
// @Produce json
// @Param taskset body models.Taskset true "Taskset"
// @Tags Taskset
// @Success 200 {array} models.Taskset
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /taskset [put]
func (e *TasksetController) createTaskset(c echo.Context) error {
	taskset := new(models.Taskset)
	if err := c.Bind(taskset); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Taskset{}, c).Model(&models.Taskset{}).Create(&taskset).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, taskset)
}

// deleteTaskset godoc
// @Id deleteTaskset
// @Summary Deletes Taskset
// @Accept json
// @Produce json
// @Tags Taskset
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /taskset/{id} [delete]
func (e *TasksetController) deleteTaskset(c echo.Context) error {
	tasksetId, err := strconv.Atoi(c.Param("taskset"))
	if err != nil {
		e.logger.Error(err)
	}

	taskset := new(models.Taskset)
	err = e.db.Get(models.Taskset{}, c).Model(&models.Taskset{}).First(&taskset, tasksetId).Error
	if err != nil || taskset.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Taskset{}, c).Model(&models.Taskset{}).Delete(&taskset).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
