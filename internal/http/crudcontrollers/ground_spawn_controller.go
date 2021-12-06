package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GroundSpawnController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGroundSpawnController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GroundSpawnController {
	return &GroundSpawnController{
		db:     db,
		logger: logger,
	}
}

func (e *GroundSpawnController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "ground_spawn/:ground_spawn", e.deleteGroundSpawn, nil),
		routes.RegisterRoute(http.MethodGet, "ground_spawn/:ground_spawn", e.getGroundSpawn, nil),
		routes.RegisterRoute(http.MethodGet, "ground_spawns", e.listGroundSpawns, nil),
		routes.RegisterRoute(http.MethodPost, "ground_spawns/bulk", e.getGroundSpawnsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "ground_spawn/:ground_spawn", e.updateGroundSpawn, nil),
		routes.RegisterRoute(http.MethodPut, "ground_spawn", e.createGroundSpawn, nil),
	}
}

// listGroundSpawns godoc
// @Id listGroundSpawns
// @Summary Lists GroundSpawns
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroundSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /ground_spawns [get]
func (e *GroundSpawnController) listGroundSpawns(c echo.Context) error {
	var results []models.GroundSpawn
	err := e.db.QueryContext(models.GroundSpawn{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGroundSpawn godoc
// @Id getGroundSpawn
// @Summary Gets GroundSpawn
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroundSpawn
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /ground_spawn/{id} [get]
func (e *GroundSpawnController) getGroundSpawn(c echo.Context) error {
	groundSpawnId, err := strconv.Atoi(c.Param("ground_spawn"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.GroundSpawn
	err = e.db.QueryContext(models.GroundSpawn{}, c).First(&result, groundSpawnId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGroundSpawn godoc
// @Id updateGroundSpawn
// @Summary Updates GroundSpawn
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param id path int true "Id"
// @Param ground_spawn body models.GroundSpawn true "GroundSpawn"
// @Success 200 {array} models.GroundSpawn
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /ground_spawn/{id} [patch]
func (e *GroundSpawnController) updateGroundSpawn(c echo.Context) error {
	groundSpawn := new(models.GroundSpawn)
	if err := c.Bind(groundSpawn); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.GroundSpawn{}
	err := e.db.Get(models.GroundSpawn{}, c).Model(&models.GroundSpawn{}).First(&entity, groundSpawn.ID).Error
	if err != nil || groundSpawn.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.GroundSpawn{}, c).Model(&entity).Select("*").Updates(&groundSpawn).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, groundSpawn)
}

// createGroundSpawn godoc
// @Id createGroundSpawn
// @Summary Creates GroundSpawn
// @Accept json
// @Produce json
// @Param ground_spawn body models.GroundSpawn true "GroundSpawn"
// @Tags GroundSpawn
// @Success 200 {array} models.GroundSpawn
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /ground_spawn [put]
func (e *GroundSpawnController) createGroundSpawn(c echo.Context) error {
	groundSpawn := new(models.GroundSpawn)
	if err := c.Bind(groundSpawn); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.GroundSpawn{}, c).Model(&models.GroundSpawn{}).Create(&groundSpawn).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, groundSpawn)
}

// deleteGroundSpawn godoc
// @Id deleteGroundSpawn
// @Summary Deletes GroundSpawn
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /ground_spawn/{id} [delete]
func (e *GroundSpawnController) deleteGroundSpawn(c echo.Context) error {
	groundSpawnId, err := strconv.Atoi(c.Param("ground_spawn"))
	if err != nil {
		e.logger.Error(err)
	}

	groundSpawn := new(models.GroundSpawn)
	err = e.db.Get(models.GroundSpawn{}, c).Model(&models.GroundSpawn{}).First(&groundSpawn, groundSpawnId).Error
	if err != nil || groundSpawn.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.GroundSpawn{}, c).Model(&models.GroundSpawn{}).Delete(&groundSpawn).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGroundSpawnsBulk godoc
// @Id getGroundSpawnsBulk
// @Summary Gets GroundSpawns in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GroundSpawn
// @Success 200 {array} models.GroundSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /ground_spawns/bulk [post]
func (e *GroundSpawnController) getGroundSpawnsBulk(c echo.Context) error {
	var results []models.GroundSpawn

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

	err := e.db.QueryContext(models.GroundSpawn{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
