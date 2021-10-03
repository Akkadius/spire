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

type TrapController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTrapController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TrapController {
	return &TrapController{
		db:     db,
		logger: logger,
	}
}

func (e *TrapController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "trap/:trap", e.deleteTrap, nil),
		routes.RegisterRoute(http.MethodGet, "trap/:trap", e.getTrap, nil),
		routes.RegisterRoute(http.MethodGet, "traps", e.listTraps, nil),
		routes.RegisterRoute(http.MethodPost, "traps/bulk", e.getTrapsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "trap/:trap", e.updateTrap, nil),
		routes.RegisterRoute(http.MethodPut, "trap", e.createTrap, nil),
	}
}

// listTraps godoc
// @Id listTraps
// @Summary Lists Traps
// @Accept json
// @Produce json
// @Tags Trap
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Trap
// @Failure 500 {string} string "Bad query request"
// @Router /traps [get]
func (e *TrapController) listTraps(c echo.Context) error {
	var results []models.Trap
	err := e.db.QueryContext(models.Trap{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTrap godoc
// @Id getTrap
// @Summary Gets Trap
// @Accept json
// @Produce json
// @Tags Trap
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Trap
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /trap/{id} [get]
func (e *TrapController) getTrap(c echo.Context) error {
	trapId, err := strconv.Atoi(c.Param("trap"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Trap
	err = e.db.QueryContext(models.Trap{}, c).First(&result, trapId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTrap godoc
// @Id updateTrap
// @Summary Updates Trap
// @Accept json
// @Produce json
// @Tags Trap
// @Param id path int true "Id"
// @Param trap body models.Trap true "Trap"
// @Success 200 {array} models.Trap
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /trap/{id} [patch]
func (e *TrapController) updateTrap(c echo.Context) error {
	trap := new(models.Trap)
	if err := c.Bind(trap); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Trap{}
	err := e.db.Get(models.Trap{}, c).Model(&models.Trap{}).First(&entity, trap.ID).Error
	if err != nil || trap.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Trap{}, c).Model(&entity).Updates(&trap).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, trap)
}

// createTrap godoc
// @Id createTrap
// @Summary Creates Trap
// @Accept json
// @Produce json
// @Param trap body models.Trap true "Trap"
// @Tags Trap
// @Success 200 {array} models.Trap
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /trap [put]
func (e *TrapController) createTrap(c echo.Context) error {
	trap := new(models.Trap)
	if err := c.Bind(trap); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Trap{}, c).Model(&models.Trap{}).Create(&trap).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, trap)
}

// deleteTrap godoc
// @Id deleteTrap
// @Summary Deletes Trap
// @Accept json
// @Produce json
// @Tags Trap
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /trap/{id} [delete]
func (e *TrapController) deleteTrap(c echo.Context) error {
	trapId, err := strconv.Atoi(c.Param("trap"))
	if err != nil {
		e.logger.Error(err)
	}

	trap := new(models.Trap)
	err = e.db.Get(models.Trap{}, c).Model(&models.Trap{}).First(&trap, trapId).Error
	if err != nil || trap.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Trap{}, c).Model(&models.Trap{}).Delete(&trap).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getTrapsBulk godoc
// @Id getTrapsBulk
// @Summary Gets Traps in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Trap
// @Success 200 {array} models.Trap
// @Failure 500 {string} string "Bad query request"
// @Router /traps/bulk [post]
func (e *TrapController) getTrapsBulk(c echo.Context) error {
	var results []models.Trap

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

	err := e.db.QueryContext(models.Trap{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
