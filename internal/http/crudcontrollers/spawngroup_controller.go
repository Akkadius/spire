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

type SpawngroupController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawngroupController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpawngroupController {
	return &SpawngroupController{
		db:	 db,
		logger: logger,
	}
}

func (e *SpawngroupController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spawngroup/:id", e.getSpawngroup, nil),
		routes.RegisterRoute(http.MethodGet, "spawngroups", e.listSpawngroups, nil),
		routes.RegisterRoute(http.MethodPut, "spawngroup", e.createSpawngroup, nil),
		routes.RegisterRoute(http.MethodDelete, "spawngroup/:id", e.deleteSpawngroup, nil),
		routes.RegisterRoute(http.MethodPatch, "spawngroup/:id", e.updateSpawngroup, nil),
		routes.RegisterRoute(http.MethodPost, "spawngroups/bulk", e.getSpawngroupsBulk, nil),
	}
}

// listSpawngroups godoc
// @Id listSpawngroups
// @Summary Lists Spawngroups
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawn2"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawngroup
// @Failure 500 {string} string "Bad query request"
// @Router /spawngroups [get]
func (e *SpawngroupController) listSpawngroups(c echo.Context) error {
	var results []models.Spawngroup
	err := e.db.QueryContext(models.Spawngroup{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawngroup godoc
// @Id getSpawngroup
// @Summary Gets Spawngroup
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawn2"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawngroup
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawngroup/{id} [get]
func (e *SpawngroupController) getSpawngroup(c echo.Context) error {
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
	var result models.Spawngroup
	query := e.db.QueryContext(models.Spawngroup{}, c)
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

// updateSpawngroup godoc
// @Id updateSpawngroup
// @Summary Updates Spawngroup
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param id path int true "Id"
// @Param spawngroup body models.Spawngroup true "Spawngroup"
// @Success 200 {array} models.Spawngroup
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawngroup/{id} [patch]
func (e *SpawngroupController) updateSpawngroup(c echo.Context) error {
	request := new(models.Spawngroup)
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
	var result models.Spawngroup
	query := e.db.QueryContext(models.Spawngroup{}, c)
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

// createSpawngroup godoc
// @Id createSpawngroup
// @Summary Creates Spawngroup
// @Accept json
// @Produce json
// @Param spawngroup body models.Spawngroup true "Spawngroup"
// @Tags Spawngroup
// @Success 200 {array} models.Spawngroup
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawngroup [put]
func (e *SpawngroupController) createSpawngroup(c echo.Context) error {
	spawngroup := new(models.Spawngroup)
	if err := c.Bind(spawngroup); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Spawngroup{}, c).Model(&models.Spawngroup{}).Create(&spawngroup).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spawngroup)
}

// deleteSpawngroup godoc
// @Id deleteSpawngroup
// @Summary Deletes Spawngroup
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawngroup/{id} [delete]
func (e *SpawngroupController) deleteSpawngroup(c echo.Context) error {
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
	var result models.Spawngroup
	query := e.db.QueryContext(models.Spawngroup{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.Spawngroup{}, c).Model(&models.Spawngroup{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpawngroupsBulk godoc
// @Id getSpawngroupsBulk
// @Summary Gets Spawngroups in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Spawngroup
// @Success 200 {array} models.Spawngroup
// @Failure 500 {string} string "Bad query request"
// @Router /spawngroups/bulk [post]
func (e *SpawngroupController) getSpawngroupsBulk(c echo.Context) error {
	var results []models.Spawngroup

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

	err := e.db.QueryContext(models.Spawngroup{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
