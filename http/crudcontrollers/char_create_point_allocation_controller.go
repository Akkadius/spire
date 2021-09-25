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

type CharCreatePointAllocationController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharCreatePointAllocationController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharCreatePointAllocationController {
	return &CharCreatePointAllocationController{
		db:     db,
		logger: logger,
	}
}

func (e *CharCreatePointAllocationController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "char_create_point_allocation/:char_create_point_allocation", e.deleteCharCreatePointAllocation, nil),
		routes.RegisterRoute(http.MethodGet, "char_create_point_allocation/:char_create_point_allocation", e.getCharCreatePointAllocation, nil),
		routes.RegisterRoute(http.MethodGet, "char_create_point_allocations", e.listCharCreatePointAllocations, nil),
		routes.RegisterRoute(http.MethodPost, "char_create_point_allocations/bulk", e.getCharCreatePointAllocationsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "char_create_point_allocation/:char_create_point_allocation", e.updateCharCreatePointAllocation, nil),
		routes.RegisterRoute(http.MethodPut, "char_create_point_allocation", e.createCharCreatePointAllocation, nil),
	}
}

// listCharCreatePointAllocations godoc
// @Id listCharCreatePointAllocations
// @Summary Lists CharCreatePointAllocations
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_point_allocations [get]
func (e *CharCreatePointAllocationController) listCharCreatePointAllocations(c echo.Context) error {
	var results []models.CharCreatePointAllocation
	err := e.db.QueryContext(models.CharCreatePointAllocation{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharCreatePointAllocation godoc
// @Id getCharCreatePointAllocation
// @Summary Gets CharCreatePointAllocation
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_point_allocation/{id} [get]
func (e *CharCreatePointAllocationController) getCharCreatePointAllocation(c echo.Context) error {
	charCreatePointAllocationId, err := strconv.Atoi(c.Param("char_create_point_allocation"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharCreatePointAllocation
	err = e.db.QueryContext(models.CharCreatePointAllocation{}, c).First(&result, charCreatePointAllocationId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharCreatePointAllocation godoc
// @Id updateCharCreatePointAllocation
// @Summary Updates CharCreatePointAllocation
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param id path int true "Id"
// @Param char_create_point_allocation body models.CharCreatePointAllocation true "CharCreatePointAllocation"
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /char_create_point_allocation/{id} [patch]
func (e *CharCreatePointAllocationController) updateCharCreatePointAllocation(c echo.Context) error {
	charCreatePointAllocation := new(models.CharCreatePointAllocation)
	if err := c.Bind(charCreatePointAllocation); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharCreatePointAllocation{}, c).Model(&models.CharCreatePointAllocation{}).First(&models.CharCreatePointAllocation{}, charCreatePointAllocation.ID).Error
	if err != nil || charCreatePointAllocation.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharCreatePointAllocation{}, c).Model(&models.CharCreatePointAllocation{}).Updates(&charCreatePointAllocation).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, charCreatePointAllocation)
}

// createCharCreatePointAllocation godoc
// @Id createCharCreatePointAllocation
// @Summary Creates CharCreatePointAllocation
// @Accept json
// @Produce json
// @Param char_create_point_allocation body models.CharCreatePointAllocation true "CharCreatePointAllocation"
// @Tags CharCreatePointAllocation
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /char_create_point_allocation [put]
func (e *CharCreatePointAllocationController) createCharCreatePointAllocation(c echo.Context) error {
	charCreatePointAllocation := new(models.CharCreatePointAllocation)
	if err := c.Bind(charCreatePointAllocation); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharCreatePointAllocation{}, c).Model(&models.CharCreatePointAllocation{}).Create(&charCreatePointAllocation).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, charCreatePointAllocation)
}

// deleteCharCreatePointAllocation godoc
// @Id deleteCharCreatePointAllocation
// @Summary Deletes CharCreatePointAllocation
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /char_create_point_allocation/{id} [delete]
func (e *CharCreatePointAllocationController) deleteCharCreatePointAllocation(c echo.Context) error {
	charCreatePointAllocationId, err := strconv.Atoi(c.Param("char_create_point_allocation"))
	if err != nil {
		e.logger.Error(err)
	}

	charCreatePointAllocation := new(models.CharCreatePointAllocation)
	err = e.db.Get(models.CharCreatePointAllocation{}, c).Model(&models.CharCreatePointAllocation{}).First(&charCreatePointAllocation, charCreatePointAllocationId).Error
	if err != nil || charCreatePointAllocation.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharCreatePointAllocation{}, c).Model(&models.CharCreatePointAllocation{}).Delete(&charCreatePointAllocation).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharCreatePointAllocationsBulk godoc
// @Id getCharCreatePointAllocationsBulk
// @Summary Gets CharCreatePointAllocations in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharCreatePointAllocation
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_point_allocations/bulk [post]
func (e *CharCreatePointAllocationController) getCharCreatePointAllocationsBulk(c echo.Context) error {
	var results []models.CharCreatePointAllocation

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

	err := e.db.QueryContext(models.CharCreatePointAllocation{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
