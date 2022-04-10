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

type TaskController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTaskController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TaskController {
	return &TaskController{
		db:	 db,
		logger: logger,
	}
}

func (e *TaskController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "task/:id", e.getTask, nil),
		routes.RegisterRoute(http.MethodGet, "tasks", e.listTasks, nil),
		routes.RegisterRoute(http.MethodPut, "task", e.createTask, nil),
		routes.RegisterRoute(http.MethodDelete, "task/:id", e.deleteTask, nil),
		routes.RegisterRoute(http.MethodPatch, "task/:id", e.updateTask, nil),
		routes.RegisterRoute(http.MethodPost, "tasks/bulk", e.getTasksBulk, nil),
	}
}

// listTasks godoc
// @Id listTasks
// @Summary Lists Tasks
// @Accept json
// @Produce json
// @Tags Task
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>TaskActivities<br>TaskActivities.Goallists<br>TaskActivities.NpcType<br>TaskActivities.NpcType.AlternateCurrency<br>TaskActivities.NpcType.Merchantlists<br>TaskActivities.NpcType.NpcEmotes<br>TaskActivities.NpcType.NpcFactions<br>TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>TaskActivities.NpcType.NpcSpells<br>TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>TaskActivities.NpcType.NpcTypesTint<br>TaskActivities.NpcType.Spawnentries<br>TaskActivities.NpcType.Spawnentries.Spawngroup<br>TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Tasksets"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Bad query request"
// @Router /tasks [get]
func (e *TaskController) listTasks(c echo.Context) error {
	var results []models.Task
	err := e.db.QueryContext(models.Task{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTask godoc
// @Id getTask
// @Summary Gets Task
// @Accept json
// @Produce json
// @Tags Task
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>TaskActivities<br>TaskActivities.Goallists<br>TaskActivities.NpcType<br>TaskActivities.NpcType.AlternateCurrency<br>TaskActivities.NpcType.Merchantlists<br>TaskActivities.NpcType.NpcEmotes<br>TaskActivities.NpcType.NpcFactions<br>TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>TaskActivities.NpcType.NpcSpells<br>TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>TaskActivities.NpcType.NpcTypesTint<br>TaskActivities.NpcType.Spawnentries<br>TaskActivities.NpcType.Spawnentries.Spawngroup<br>TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Tasksets"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Task
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /task/{id} [get]
func (e *TaskController) getTask(c echo.Context) error {
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
	var result models.Task
	query := e.db.QueryContext(models.Task{}, c)
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

// updateTask godoc
// @Id updateTask
// @Summary Updates Task
// @Accept json
// @Produce json
// @Tags Task
// @Param id path int true "Id"
// @Param task body models.Task true "Task"
// @Success 200 {array} models.Task
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /task/{id} [patch]
func (e *TaskController) updateTask(c echo.Context) error {
	request := new(models.Task)
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
	var result models.Task
	query := e.db.QueryContext(models.Task{}, c)
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

// createTask godoc
// @Id createTask
// @Summary Creates Task
// @Accept json
// @Produce json
// @Param task body models.Task true "Task"
// @Tags Task
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /task [put]
func (e *TaskController) createTask(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Task{}, c).Model(&models.Task{}).Create(&task).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, task)
}

// deleteTask godoc
// @Id deleteTask
// @Summary Deletes Task
// @Accept json
// @Produce json
// @Tags Task
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /task/{id} [delete]
func (e *TaskController) deleteTask(c echo.Context) error {
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
	var result models.Task
	query := e.db.QueryContext(models.Task{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.Task{}, c).Model(&models.Task{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getTasksBulk godoc
// @Id getTasksBulk
// @Summary Gets Tasks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Task
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Bad query request"
// @Router /tasks/bulk [post]
func (e *TaskController) getTasksBulk(c echo.Context) error {
	var results []models.Task

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

	err := e.db.QueryContext(models.Task{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
