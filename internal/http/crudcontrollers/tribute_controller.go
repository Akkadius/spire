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

type TributeController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTributeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TributeController {
	return &TributeController{
		db:     db,
		logger: logger,
	}
}

func (e *TributeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "tribute/:tribute", e.deleteTribute, nil),
		routes.RegisterRoute(http.MethodGet, "tribute/:tribute", e.getTribute, nil),
		routes.RegisterRoute(http.MethodGet, "tributes", e.listTributes, nil),
		routes.RegisterRoute(http.MethodPost, "tributes/bulk", e.getTributesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "tribute/:tribute", e.updateTribute, nil),
		routes.RegisterRoute(http.MethodPut, "tribute", e.createTribute, nil),
	}
}

// listTributes godoc
// @Id listTributes
// @Summary Lists Tributes
// @Accept json
// @Produce json
// @Tags Tribute
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Tribute
// @Failure 500 {string} string "Bad query request"
// @Router /tributes [get]
func (e *TributeController) listTributes(c echo.Context) error {
	var results []models.Tribute
	err := e.db.QueryContext(models.Tribute{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTribute godoc
// @Id getTribute
// @Summary Gets Tribute
// @Accept json
// @Produce json
// @Tags Tribute
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Tribute
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tribute/{id} [get]
func (e *TributeController) getTribute(c echo.Context) error {
	tributeId, err := strconv.Atoi(c.Param("tribute"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Tribute
	err = e.db.QueryContext(models.Tribute{}, c).First(&result, tributeId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTribute godoc
// @Id updateTribute
// @Summary Updates Tribute
// @Accept json
// @Produce json
// @Tags Tribute
// @Param id path int true "Id"
// @Param tribute body models.Tribute true "Tribute"
// @Success 200 {array} models.Tribute
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /tribute/{id} [patch]
func (e *TributeController) updateTribute(c echo.Context) error {
	tribute := new(models.Tribute)
	if err := c.Bind(tribute); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Tribute{}
	err := e.db.Get(models.Tribute{}, c).Model(&models.Tribute{}).First(&entity, tribute.ID).Error
	if err != nil || tribute.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Tribute{}, c).Model(&entity).Select("*").Updates(&tribute).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, tribute)
}

// createTribute godoc
// @Id createTribute
// @Summary Creates Tribute
// @Accept json
// @Produce json
// @Param tribute body models.Tribute true "Tribute"
// @Tags Tribute
// @Success 200 {array} models.Tribute
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /tribute [put]
func (e *TributeController) createTribute(c echo.Context) error {
	tribute := new(models.Tribute)
	if err := c.Bind(tribute); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Tribute{}, c).Model(&models.Tribute{}).Create(&tribute).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, tribute)
}

// deleteTribute godoc
// @Id deleteTribute
// @Summary Deletes Tribute
// @Accept json
// @Produce json
// @Tags Tribute
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tribute/{id} [delete]
func (e *TributeController) deleteTribute(c echo.Context) error {
	tributeId, err := strconv.Atoi(c.Param("tribute"))
	if err != nil {
		e.logger.Error(err)
	}

	tribute := new(models.Tribute)
	err = e.db.Get(models.Tribute{}, c).Model(&models.Tribute{}).First(&tribute, tributeId).Error
	if err != nil || tribute.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Tribute{}, c).Model(&models.Tribute{}).Delete(&tribute).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getTributesBulk godoc
// @Id getTributesBulk
// @Summary Gets Tributes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Tribute
// @Success 200 {array} models.Tribute
// @Failure 500 {string} string "Bad query request"
// @Router /tributes/bulk [post]
func (e *TributeController) getTributesBulk(c echo.Context) error {
	var results []models.Tribute

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

	err := e.db.QueryContext(models.Tribute{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
