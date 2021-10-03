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

type GraveyardController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGraveyardController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GraveyardController {
	return &GraveyardController{
		db:     db,
		logger: logger,
	}
}

func (e *GraveyardController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "graveyard/:graveyard", e.deleteGraveyard, nil),
		routes.RegisterRoute(http.MethodGet, "graveyard/:graveyard", e.getGraveyard, nil),
		routes.RegisterRoute(http.MethodGet, "graveyards", e.listGraveyards, nil),
		routes.RegisterRoute(http.MethodPost, "graveyards/bulk", e.getGraveyardsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "graveyard/:graveyard", e.updateGraveyard, nil),
		routes.RegisterRoute(http.MethodPut, "graveyard", e.createGraveyard, nil),
	}
}

// listGraveyards godoc
// @Id listGraveyards
// @Summary Lists Graveyards
// @Accept json
// @Produce json
// @Tags Graveyard
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Graveyard
// @Failure 500 {string} string "Bad query request"
// @Router /graveyards [get]
func (e *GraveyardController) listGraveyards(c echo.Context) error {
	var results []models.Graveyard
	err := e.db.QueryContext(models.Graveyard{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGraveyard godoc
// @Id getGraveyard
// @Summary Gets Graveyard
// @Accept json
// @Produce json
// @Tags Graveyard
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Graveyard
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /graveyard/{id} [get]
func (e *GraveyardController) getGraveyard(c echo.Context) error {
	graveyardId, err := strconv.Atoi(c.Param("graveyard"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Graveyard
	err = e.db.QueryContext(models.Graveyard{}, c).First(&result, graveyardId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGraveyard godoc
// @Id updateGraveyard
// @Summary Updates Graveyard
// @Accept json
// @Produce json
// @Tags Graveyard
// @Param id path int true "Id"
// @Param graveyard body models.Graveyard true "Graveyard"
// @Success 200 {array} models.Graveyard
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /graveyard/{id} [patch]
func (e *GraveyardController) updateGraveyard(c echo.Context) error {
	graveyard := new(models.Graveyard)
	if err := c.Bind(graveyard); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Graveyard{}
	err := e.db.Get(models.Graveyard{}, c).Model(&models.Graveyard{}).First(&entity, graveyard.ID).Error
	if err != nil || graveyard.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Graveyard{}, c).Model(&entity).Updates(&graveyard).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, graveyard)
}

// createGraveyard godoc
// @Id createGraveyard
// @Summary Creates Graveyard
// @Accept json
// @Produce json
// @Param graveyard body models.Graveyard true "Graveyard"
// @Tags Graveyard
// @Success 200 {array} models.Graveyard
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /graveyard [put]
func (e *GraveyardController) createGraveyard(c echo.Context) error {
	graveyard := new(models.Graveyard)
	if err := c.Bind(graveyard); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Graveyard{}, c).Model(&models.Graveyard{}).Create(&graveyard).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, graveyard)
}

// deleteGraveyard godoc
// @Id deleteGraveyard
// @Summary Deletes Graveyard
// @Accept json
// @Produce json
// @Tags Graveyard
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /graveyard/{id} [delete]
func (e *GraveyardController) deleteGraveyard(c echo.Context) error {
	graveyardId, err := strconv.Atoi(c.Param("graveyard"))
	if err != nil {
		e.logger.Error(err)
	}

	graveyard := new(models.Graveyard)
	err = e.db.Get(models.Graveyard{}, c).Model(&models.Graveyard{}).First(&graveyard, graveyardId).Error
	if err != nil || graveyard.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Graveyard{}, c).Model(&models.Graveyard{}).Delete(&graveyard).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGraveyardsBulk godoc
// @Id getGraveyardsBulk
// @Summary Gets Graveyards in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Graveyard
// @Success 200 {array} models.Graveyard
// @Failure 500 {string} string "Bad query request"
// @Router /graveyards/bulk [post]
func (e *GraveyardController) getGraveyardsBulk(c echo.Context) error {
	var results []models.Graveyard

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

	err := e.db.QueryContext(models.Graveyard{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
