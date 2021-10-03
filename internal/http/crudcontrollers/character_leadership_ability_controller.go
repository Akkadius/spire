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

type CharacterLeadershipAbilityController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterLeadershipAbilityController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterLeadershipAbilityController {
	return &CharacterLeadershipAbilityController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterLeadershipAbilityController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_leadership_ability/:character_leadership_ability", e.deleteCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodGet, "character_leadership_ability/:character_leadership_ability", e.getCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodGet, "character_leadership_abilities", e.listCharacterLeadershipAbilities, nil),
		routes.RegisterRoute(http.MethodPost, "character_leadership_abilities/bulk", e.getCharacterLeadershipAbilitiesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_leadership_ability/:character_leadership_ability", e.updateCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodPut, "character_leadership_ability", e.createCharacterLeadershipAbility, nil),
	}
}

// listCharacterLeadershipAbilities godoc
// @Id listCharacterLeadershipAbilities
// @Summary Lists CharacterLeadershipAbilities
// @Accept json
// @Produce json
// @Tags CharacterLeadershipAbility
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterLeadershipAbility
// @Failure 500 {string} string "Bad query request"
// @Router /character_leadership_abilities [get]
func (e *CharacterLeadershipAbilityController) listCharacterLeadershipAbilities(c echo.Context) error {
	var results []models.CharacterLeadershipAbility
	err := e.db.QueryContext(models.CharacterLeadershipAbility{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterLeadershipAbility godoc
// @Id getCharacterLeadershipAbility
// @Summary Gets CharacterLeadershipAbility
// @Accept json
// @Produce json
// @Tags CharacterLeadershipAbility
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterLeadershipAbility
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_leadership_ability/{id} [get]
func (e *CharacterLeadershipAbilityController) getCharacterLeadershipAbility(c echo.Context) error {
	characterLeadershipAbilityId, err := strconv.Atoi(c.Param("character_leadership_ability"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterLeadershipAbility
	err = e.db.QueryContext(models.CharacterLeadershipAbility{}, c).First(&result, characterLeadershipAbilityId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterLeadershipAbility godoc
// @Id updateCharacterLeadershipAbility
// @Summary Updates CharacterLeadershipAbility
// @Accept json
// @Produce json
// @Tags CharacterLeadershipAbility
// @Param id path int true "Id"
// @Param character_leadership_ability body models.CharacterLeadershipAbility true "CharacterLeadershipAbility"
// @Success 200 {array} models.CharacterLeadershipAbility
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_leadership_ability/{id} [patch]
func (e *CharacterLeadershipAbilityController) updateCharacterLeadershipAbility(c echo.Context) error {
	characterLeadershipAbility := new(models.CharacterLeadershipAbility)
	if err := c.Bind(characterLeadershipAbility); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.CharacterLeadershipAbility{}
	err := e.db.Get(models.CharacterLeadershipAbility{}, c).Model(&models.CharacterLeadershipAbility{}).First(&entity, characterLeadershipAbility.ID).Error
	if err != nil || characterLeadershipAbility.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterLeadershipAbility{}, c).Model(&entity).Updates(&characterLeadershipAbility).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterLeadershipAbility)
}

// createCharacterLeadershipAbility godoc
// @Id createCharacterLeadershipAbility
// @Summary Creates CharacterLeadershipAbility
// @Accept json
// @Produce json
// @Param character_leadership_ability body models.CharacterLeadershipAbility true "CharacterLeadershipAbility"
// @Tags CharacterLeadershipAbility
// @Success 200 {array} models.CharacterLeadershipAbility
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_leadership_ability [put]
func (e *CharacterLeadershipAbilityController) createCharacterLeadershipAbility(c echo.Context) error {
	characterLeadershipAbility := new(models.CharacterLeadershipAbility)
	if err := c.Bind(characterLeadershipAbility); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterLeadershipAbility{}, c).Model(&models.CharacterLeadershipAbility{}).Create(&characterLeadershipAbility).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, characterLeadershipAbility)
}

// deleteCharacterLeadershipAbility godoc
// @Id deleteCharacterLeadershipAbility
// @Summary Deletes CharacterLeadershipAbility
// @Accept json
// @Produce json
// @Tags CharacterLeadershipAbility
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_leadership_ability/{id} [delete]
func (e *CharacterLeadershipAbilityController) deleteCharacterLeadershipAbility(c echo.Context) error {
	characterLeadershipAbilityId, err := strconv.Atoi(c.Param("character_leadership_ability"))
	if err != nil {
		e.logger.Error(err)
	}

	characterLeadershipAbility := new(models.CharacterLeadershipAbility)
	err = e.db.Get(models.CharacterLeadershipAbility{}, c).Model(&models.CharacterLeadershipAbility{}).First(&characterLeadershipAbility, characterLeadershipAbilityId).Error
	if err != nil || characterLeadershipAbility.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterLeadershipAbility{}, c).Model(&models.CharacterLeadershipAbility{}).Delete(&characterLeadershipAbility).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterLeadershipAbilitiesBulk godoc
// @Id getCharacterLeadershipAbilitiesBulk
// @Summary Gets CharacterLeadershipAbilities in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterLeadershipAbility
// @Success 200 {array} models.CharacterLeadershipAbility
// @Failure 500 {string} string "Bad query request"
// @Router /character_leadership_abilities/bulk [post]
func (e *CharacterLeadershipAbilityController) getCharacterLeadershipAbilitiesBulk(c echo.Context) error {
	var results []models.CharacterLeadershipAbility

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

	err := e.db.QueryContext(models.CharacterLeadershipAbility{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
