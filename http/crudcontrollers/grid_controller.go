package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GridController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGridController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GridController {
	return &GridController{
		db:     db,
		logger: logger,
	}
}

func (e *GridController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "grid/:grid", e.deleteGrid, nil),
		routes.RegisterRoute(http.MethodGet, "grid/:grid", e.getGrid, nil),
		routes.RegisterRoute(http.MethodGet, "grids", e.listGrids, nil),
		routes.RegisterRoute(http.MethodPost, "grids/bulk", e.getGridsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "grid/:grid", e.updateGrid, nil),
		routes.RegisterRoute(http.MethodPut, "grid", e.createGrid, nil),
	}
}

// listGrids godoc
// @Id listGrids
// @Summary Lists Grids
// @Accept json
// @Produce json
// @Tags Grid
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>GridEntries"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Grid
// @Failure 500 {string} string "Bad query request"
// @Router /grids [get]
func (e *GridController) listGrids(c echo.Context) error {
	var results []models.Grid
	err := e.db.QueryContext(models.Grid{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGrid godoc
// @Id getGrid
// @Summary Gets Grid
// @Accept json
// @Produce json
// @Tags Grid
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>GridEntries"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Grid
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /grid/{id} [get]
func (e *GridController) getGrid(c echo.Context) error {
	gridId, err := strconv.Atoi(c.Param("grid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Grid
	err = e.db.QueryContext(models.Grid{}, c).First(&result, gridId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGrid godoc
// @Id updateGrid
// @Summary Updates Grid
// @Accept json
// @Produce json
// @Tags Grid
// @Param id path int true "Id"
// @Param grid body models.Grid true "Grid"
// @Success 200 {array} models.Grid
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /grid/{id} [patch]
func (e *GridController) updateGrid(c echo.Context) error {
	grid := new(models.Grid)
	if err := c.Bind(grid); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Grid{}
	err := e.db.Get(models.Grid{}, c).Model(&models.Grid{}).First(&entity, grid.ID).Error
	if err != nil || grid.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Grid{}, c).Model(&entity).Updates(&grid).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, grid)
}

// createGrid godoc
// @Id createGrid
// @Summary Creates Grid
// @Accept json
// @Produce json
// @Param grid body models.Grid true "Grid"
// @Tags Grid
// @Success 200 {array} models.Grid
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /grid [put]
func (e *GridController) createGrid(c echo.Context) error {
	grid := new(models.Grid)
	if err := c.Bind(grid); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Grid{}, c).Model(&models.Grid{}).Create(&grid).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, grid)
}

// deleteGrid godoc
// @Id deleteGrid
// @Summary Deletes Grid
// @Accept json
// @Produce json
// @Tags Grid
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /grid/{id} [delete]
func (e *GridController) deleteGrid(c echo.Context) error {
	gridId, err := strconv.Atoi(c.Param("grid"))
	if err != nil {
		e.logger.Error(err)
	}

	grid := new(models.Grid)
	err = e.db.Get(models.Grid{}, c).Model(&models.Grid{}).First(&grid, gridId).Error
	if err != nil || grid.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Grid{}, c).Model(&models.Grid{}).Delete(&grid).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGridsBulk godoc
// @Id getGridsBulk
// @Summary Gets Grids in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Grid
// @Success 200 {array} models.Grid
// @Failure 500 {string} string "Bad query request"
// @Router /grids/bulk [post]
func (e *GridController) getGridsBulk(c echo.Context) error {
	var results []models.Grid

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

	err := e.db.QueryContext(models.Grid{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
