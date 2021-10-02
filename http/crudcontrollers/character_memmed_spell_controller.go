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

type CharacterMemmedSpellController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterMemmedSpellController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterMemmedSpellController {
	return &CharacterMemmedSpellController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterMemmedSpellController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_memmed_spell/:character_memmed_spell", e.deleteCharacterMemmedSpell, nil),
		routes.RegisterRoute(http.MethodGet, "character_memmed_spell/:character_memmed_spell", e.getCharacterMemmedSpell, nil),
		routes.RegisterRoute(http.MethodGet, "character_memmed_spells", e.listCharacterMemmedSpells, nil),
		routes.RegisterRoute(http.MethodPost, "character_memmed_spells/bulk", e.getCharacterMemmedSpellsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_memmed_spell/:character_memmed_spell", e.updateCharacterMemmedSpell, nil),
		routes.RegisterRoute(http.MethodPut, "character_memmed_spell", e.createCharacterMemmedSpell, nil),
	}
}

// listCharacterMemmedSpells godoc
// @Id listCharacterMemmedSpells
// @Summary Lists CharacterMemmedSpells
// @Accept json
// @Produce json
// @Tags CharacterMemmedSpell
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterMemmedSpell
// @Failure 500 {string} string "Bad query request"
// @Router /character_memmed_spells [get]
func (e *CharacterMemmedSpellController) listCharacterMemmedSpells(c echo.Context) error {
	var results []models.CharacterMemmedSpell
	err := e.db.QueryContext(models.CharacterMemmedSpell{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterMemmedSpell godoc
// @Id getCharacterMemmedSpell
// @Summary Gets CharacterMemmedSpell
// @Accept json
// @Produce json
// @Tags CharacterMemmedSpell
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterMemmedSpell
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_memmed_spell/{id} [get]
func (e *CharacterMemmedSpellController) getCharacterMemmedSpell(c echo.Context) error {
	characterMemmedSpellId, err := strconv.Atoi(c.Param("character_memmed_spell"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterMemmedSpell
	err = e.db.QueryContext(models.CharacterMemmedSpell{}, c).First(&result, characterMemmedSpellId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterMemmedSpell godoc
// @Id updateCharacterMemmedSpell
// @Summary Updates CharacterMemmedSpell
// @Accept json
// @Produce json
// @Tags CharacterMemmedSpell
// @Param id path int true "Id"
// @Param character_memmed_spell body models.CharacterMemmedSpell true "CharacterMemmedSpell"
// @Success 200 {array} models.CharacterMemmedSpell
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_memmed_spell/{id} [patch]
func (e *CharacterMemmedSpellController) updateCharacterMemmedSpell(c echo.Context) error {
	characterMemmedSpell := new(models.CharacterMemmedSpell)
	if err := c.Bind(characterMemmedSpell); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.CharacterMemmedSpell{}
	err := e.db.Get(models.CharacterMemmedSpell{}, c).Model(&models.CharacterMemmedSpell{}).First(&entity, characterMemmedSpell.ID).Error
	if err != nil || characterMemmedSpell.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterMemmedSpell{}, c).Model(&entity).Updates(&characterMemmedSpell).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterMemmedSpell)
}

// createCharacterMemmedSpell godoc
// @Id createCharacterMemmedSpell
// @Summary Creates CharacterMemmedSpell
// @Accept json
// @Produce json
// @Param character_memmed_spell body models.CharacterMemmedSpell true "CharacterMemmedSpell"
// @Tags CharacterMemmedSpell
// @Success 200 {array} models.CharacterMemmedSpell
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_memmed_spell [put]
func (e *CharacterMemmedSpellController) createCharacterMemmedSpell(c echo.Context) error {
	characterMemmedSpell := new(models.CharacterMemmedSpell)
	if err := c.Bind(characterMemmedSpell); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterMemmedSpell{}, c).Model(&models.CharacterMemmedSpell{}).Create(&characterMemmedSpell).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, characterMemmedSpell)
}

// deleteCharacterMemmedSpell godoc
// @Id deleteCharacterMemmedSpell
// @Summary Deletes CharacterMemmedSpell
// @Accept json
// @Produce json
// @Tags CharacterMemmedSpell
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_memmed_spell/{id} [delete]
func (e *CharacterMemmedSpellController) deleteCharacterMemmedSpell(c echo.Context) error {
	characterMemmedSpellId, err := strconv.Atoi(c.Param("character_memmed_spell"))
	if err != nil {
		e.logger.Error(err)
	}

	characterMemmedSpell := new(models.CharacterMemmedSpell)
	err = e.db.Get(models.CharacterMemmedSpell{}, c).Model(&models.CharacterMemmedSpell{}).First(&characterMemmedSpell, characterMemmedSpellId).Error
	if err != nil || characterMemmedSpell.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterMemmedSpell{}, c).Model(&models.CharacterMemmedSpell{}).Delete(&characterMemmedSpell).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterMemmedSpellsBulk godoc
// @Id getCharacterMemmedSpellsBulk
// @Summary Gets CharacterMemmedSpells in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterMemmedSpell
// @Success 200 {array} models.CharacterMemmedSpell
// @Failure 500 {string} string "Bad query request"
// @Router /character_memmed_spells/bulk [post]
func (e *CharacterMemmedSpellController) getCharacterMemmedSpellsBulk(c echo.Context) error {
	var results []models.CharacterMemmedSpell

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

	err := e.db.QueryContext(models.CharacterMemmedSpell{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
