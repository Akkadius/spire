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

type SpawnConditionController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawnConditionController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpawnConditionController {
	return &SpawnConditionController{
		db:	 db,
		logger: logger,
	}
}

func (e *SpawnConditionController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spawn_condition/:id", e.getSpawnCondition, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_conditions", e.listSpawnConditions, nil),
		routes.RegisterRoute(http.MethodPut, "spawn_condition", e.createSpawnCondition, nil),
		routes.RegisterRoute(http.MethodDelete, "spawn_condition/:id", e.deleteSpawnCondition, nil),
		routes.RegisterRoute(http.MethodPatch, "spawn_condition/:id", e.updateSpawnCondition, nil),
		routes.RegisterRoute(http.MethodPost, "spawn_conditions/bulk", e.getSpawnConditionsBulk, nil),
	}
}

// listSpawnConditions godoc
// @Id listSpawnConditions
// @Summary Lists SpawnConditions
// @Accept json
// @Produce json
// @Tags SpawnCondition
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpawnCondition
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_conditions [get]
func (e *SpawnConditionController) listSpawnConditions(c echo.Context) error {
	var results []models.SpawnCondition
	err := e.db.QueryContext(models.SpawnCondition{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawnCondition godoc
// @Id getSpawnCondition
// @Summary Gets SpawnCondition
// @Accept json
// @Produce json
// @Tags SpawnCondition
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpawnCondition
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_condition/{id} [get]
func (e *SpawnConditionController) getSpawnCondition(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [id] position [2] type [mediumint]
	if len(c.QueryParam("id")) > 0 {
		idParam, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [id] err [%s]", err.Error())})
		}

		params = append(params, idParam)
		keys = append(keys, "id = ?")
	}

	// query builder
	var result models.SpawnCondition
	query := e.db.QueryContext(models.SpawnCondition{}, c)
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

// updateSpawnCondition godoc
// @Id updateSpawnCondition
// @Summary Updates SpawnCondition
// @Accept json
// @Produce json
// @Tags SpawnCondition
// @Param id path int true "Id"
// @Param spawn_condition body models.SpawnCondition true "SpawnCondition"
// @Success 200 {array} models.SpawnCondition
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawn_condition/{id} [patch]
func (e *SpawnConditionController) updateSpawnCondition(c echo.Context) error {
	request := new(models.SpawnCondition)
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

	// key param [id] position [2] type [mediumint]
	if len(c.QueryParam("id")) > 0 {
		idParam, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [id] err [%s]", err.Error())})
		}

		params = append(params, idParam)
		keys = append(keys, "id = ?")
	}

	// query builder
	var result models.SpawnCondition
	query := e.db.QueryContext(models.SpawnCondition{}, c)
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

// createSpawnCondition godoc
// @Id createSpawnCondition
// @Summary Creates SpawnCondition
// @Accept json
// @Produce json
// @Param spawn_condition body models.SpawnCondition true "SpawnCondition"
// @Tags SpawnCondition
// @Success 200 {array} models.SpawnCondition
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawn_condition [put]
func (e *SpawnConditionController) createSpawnCondition(c echo.Context) error {
	spawnCondition := new(models.SpawnCondition)
	if err := c.Bind(spawnCondition); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SpawnCondition{}, c).Model(&models.SpawnCondition{}).Create(&spawnCondition).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spawnCondition)
}

// deleteSpawnCondition godoc
// @Id deleteSpawnCondition
// @Summary Deletes SpawnCondition
// @Accept json
// @Produce json
// @Tags SpawnCondition
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawn_condition/{id} [delete]
func (e *SpawnConditionController) deleteSpawnCondition(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [id] position [2] type [mediumint]
	if len(c.QueryParam("id")) > 0 {
		idParam, err := strconv.Atoi(c.QueryParam("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [id] err [%s]", err.Error())})
		}

		params = append(params, idParam)
		keys = append(keys, "id = ?")
	}

	// query builder
	var result models.SpawnCondition
	query := e.db.QueryContext(models.SpawnCondition{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.SpawnCondition{}, c).Model(&models.SpawnCondition{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpawnConditionsBulk godoc
// @Id getSpawnConditionsBulk
// @Summary Gets SpawnConditions in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpawnCondition
// @Success 200 {array} models.SpawnCondition
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_conditions/bulk [post]
func (e *SpawnConditionController) getSpawnConditionsBulk(c echo.Context) error {
	var results []models.SpawnCondition

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

	err := e.db.QueryContext(models.SpawnCondition{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
