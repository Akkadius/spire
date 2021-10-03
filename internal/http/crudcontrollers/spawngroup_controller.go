package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type SpawngroupController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawngroupController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpawngroupController {
	return &SpawngroupController{
		db:     db,
		logger: logger,
	}
}

func (e *SpawngroupController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "spawngroup/:spawngroup", e.deleteSpawngroup, nil),
		routes.RegisterRoute(http.MethodGet, "spawngroup/:spawngroup", e.getSpawngroup, nil),
		routes.RegisterRoute(http.MethodGet, "spawngroups", e.listSpawngroups, nil),
		routes.RegisterRoute(http.MethodPost, "spawngroups/bulk", e.getSpawngroupsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "spawngroup/:spawngroup", e.updateSpawngroup, nil),
		routes.RegisterRoute(http.MethodPut, "spawngroup", e.createSpawngroup, nil),
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
	spawngroupId, err := strconv.Atoi(c.Param("spawngroup"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Spawngroup
	err = e.db.QueryContext(models.Spawngroup{}, c).First(&result, spawngroupId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

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
	spawngroup := new(models.Spawngroup)
	if err := c.Bind(spawngroup); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Spawngroup{}
	err := e.db.Get(models.Spawngroup{}, c).Model(&models.Spawngroup{}).First(&entity, spawngroup.ID).Error
	if err != nil || spawngroup.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Spawngroup{}, c).Model(&entity).Updates(&spawngroup).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, spawngroup)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Spawngroup{}, c).Model(&models.Spawngroup{}).Create(&spawngroup).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
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
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawngroup/{id} [delete]
func (e *SpawngroupController) deleteSpawngroup(c echo.Context) error {
	spawngroupId, err := strconv.Atoi(c.Param("spawngroup"))
	if err != nil {
		e.logger.Error(err)
	}

	spawngroup := new(models.Spawngroup)
	err = e.db.Get(models.Spawngroup{}, c).Model(&models.Spawngroup{}).First(&spawngroup, spawngroupId).Error
	if err != nil || spawngroup.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Spawngroup{}, c).Model(&models.Spawngroup{}).Delete(&spawngroup).Error
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
