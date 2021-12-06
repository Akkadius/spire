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

type ForageController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewForageController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ForageController {
	return &ForageController{
		db:     db,
		logger: logger,
	}
}

func (e *ForageController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "forage/:forage", e.deleteForage, nil),
		routes.RegisterRoute(http.MethodGet, "forage/:forage", e.getForage, nil),
		routes.RegisterRoute(http.MethodGet, "forages", e.listForages, nil),
		routes.RegisterRoute(http.MethodPost, "forages/bulk", e.getForagesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "forage/:forage", e.updateForage, nil),
		routes.RegisterRoute(http.MethodPut, "forage", e.createForage, nil),
	}
}

// listForages godoc
// @Id listForages
// @Summary Lists Forages
// @Accept json
// @Produce json
// @Tags Forage
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Forage
// @Failure 500 {string} string "Bad query request"
// @Router /forages [get]
func (e *ForageController) listForages(c echo.Context) error {
	var results []models.Forage
	err := e.db.QueryContext(models.Forage{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getForage godoc
// @Id getForage
// @Summary Gets Forage
// @Accept json
// @Produce json
// @Tags Forage
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Forage
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /forage/{id} [get]
func (e *ForageController) getForage(c echo.Context) error {
	forageId, err := strconv.Atoi(c.Param("forage"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Forage
	err = e.db.QueryContext(models.Forage{}, c).First(&result, forageId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateForage godoc
// @Id updateForage
// @Summary Updates Forage
// @Accept json
// @Produce json
// @Tags Forage
// @Param id path int true "Id"
// @Param forage body models.Forage true "Forage"
// @Success 200 {array} models.Forage
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /forage/{id} [patch]
func (e *ForageController) updateForage(c echo.Context) error {
	forage := new(models.Forage)
	if err := c.Bind(forage); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Forage{}
	err := e.db.Get(models.Forage{}, c).Model(&models.Forage{}).First(&entity, forage.ID).Error
	if err != nil || forage.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Forage{}, c).Model(&entity).Select("*").Updates(&forage).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, forage)
}

// createForage godoc
// @Id createForage
// @Summary Creates Forage
// @Accept json
// @Produce json
// @Param forage body models.Forage true "Forage"
// @Tags Forage
// @Success 200 {array} models.Forage
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /forage [put]
func (e *ForageController) createForage(c echo.Context) error {
	forage := new(models.Forage)
	if err := c.Bind(forage); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Forage{}, c).Model(&models.Forage{}).Create(&forage).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, forage)
}

// deleteForage godoc
// @Id deleteForage
// @Summary Deletes Forage
// @Accept json
// @Produce json
// @Tags Forage
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /forage/{id} [delete]
func (e *ForageController) deleteForage(c echo.Context) error {
	forageId, err := strconv.Atoi(c.Param("forage"))
	if err != nil {
		e.logger.Error(err)
	}

	forage := new(models.Forage)
	err = e.db.Get(models.Forage{}, c).Model(&models.Forage{}).First(&forage, forageId).Error
	if err != nil || forage.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Forage{}, c).Model(&models.Forage{}).Delete(&forage).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getForagesBulk godoc
// @Id getForagesBulk
// @Summary Gets Forages in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Forage
// @Success 200 {array} models.Forage
// @Failure 500 {string} string "Bad query request"
// @Router /forages/bulk [post]
func (e *ForageController) getForagesBulk(c echo.Context) error {
	var results []models.Forage

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

	err := e.db.QueryContext(models.Forage{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
