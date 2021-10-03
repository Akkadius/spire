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

type AaAbilityController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAaAbilityController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AaAbilityController {
	return &AaAbilityController{
		db:     db,
		logger: logger,
	}
}

func (e *AaAbilityController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "aa_ability/:aa_ability", e.deleteAaAbility, nil),
		routes.RegisterRoute(http.MethodGet, "aa_ability/:aa_ability", e.getAaAbility, nil),
		routes.RegisterRoute(http.MethodGet, "aa_abilities", e.listAaAbilities, nil),
		routes.RegisterRoute(http.MethodPost, "aa_abilities/bulk", e.getAaAbilitiesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "aa_ability/:aa_ability", e.updateAaAbility, nil),
		routes.RegisterRoute(http.MethodPut, "aa_ability", e.createAaAbility, nil),
	}
}

// listAaAbilities godoc
// @Id listAaAbilities
// @Summary Lists AaAbilities
// @Accept json
// @Produce json
// @Tags AaAbility
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaAbility
// @Failure 500 {string} string "Bad query request"
// @Router /aa_abilities [get]
func (e *AaAbilityController) listAaAbilities(c echo.Context) error {
	var results []models.AaAbility
	err := e.db.QueryContext(models.AaAbility{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAaAbility godoc
// @Id getAaAbility
// @Summary Gets AaAbility
// @Accept json
// @Produce json
// @Tags AaAbility
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaAbility
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /aa_ability/{id} [get]
func (e *AaAbilityController) getAaAbility(c echo.Context) error {
	aaAbilityId, err := strconv.Atoi(c.Param("aa_ability"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.AaAbility
	err = e.db.QueryContext(models.AaAbility{}, c).First(&result, aaAbilityId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAaAbility godoc
// @Id updateAaAbility
// @Summary Updates AaAbility
// @Accept json
// @Produce json
// @Tags AaAbility
// @Param id path int true "Id"
// @Param aa_ability body models.AaAbility true "AaAbility"
// @Success 200 {array} models.AaAbility
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /aa_ability/{id} [patch]
func (e *AaAbilityController) updateAaAbility(c echo.Context) error {
	aaAbility := new(models.AaAbility)
	if err := c.Bind(aaAbility); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.AaAbility{}
	err := e.db.Get(models.AaAbility{}, c).Model(&models.AaAbility{}).First(&entity, aaAbility.ID).Error
	if err != nil || aaAbility.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AaAbility{}, c).Model(&entity).Updates(&aaAbility).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, aaAbility)
}

// createAaAbility godoc
// @Id createAaAbility
// @Summary Creates AaAbility
// @Accept json
// @Produce json
// @Param aa_ability body models.AaAbility true "AaAbility"
// @Tags AaAbility
// @Success 200 {array} models.AaAbility
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /aa_ability [put]
func (e *AaAbilityController) createAaAbility(c echo.Context) error {
	aaAbility := new(models.AaAbility)
	if err := c.Bind(aaAbility); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.AaAbility{}, c).Model(&models.AaAbility{}).Create(&aaAbility).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, aaAbility)
}

// deleteAaAbility godoc
// @Id deleteAaAbility
// @Summary Deletes AaAbility
// @Accept json
// @Produce json
// @Tags AaAbility
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /aa_ability/{id} [delete]
func (e *AaAbilityController) deleteAaAbility(c echo.Context) error {
	aaAbilityId, err := strconv.Atoi(c.Param("aa_ability"))
	if err != nil {
		e.logger.Error(err)
	}

	aaAbility := new(models.AaAbility)
	err = e.db.Get(models.AaAbility{}, c).Model(&models.AaAbility{}).First(&aaAbility, aaAbilityId).Error
	if err != nil || aaAbility.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.AaAbility{}, c).Model(&models.AaAbility{}).Delete(&aaAbility).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getAaAbilitiesBulk godoc
// @Id getAaAbilitiesBulk
// @Summary Gets AaAbilities in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AaAbility
// @Success 200 {array} models.AaAbility
// @Failure 500 {string} string "Bad query request"
// @Router /aa_abilities/bulk [post]
func (e *AaAbilityController) getAaAbilitiesBulk(c echo.Context) error {
	var results []models.AaAbility

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

	err := e.db.QueryContext(models.AaAbility{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
