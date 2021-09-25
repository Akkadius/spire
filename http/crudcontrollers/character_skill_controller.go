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

type CharacterSkillController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterSkillController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterSkillController {
	return &CharacterSkillController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterSkillController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_skill/:character_skill", e.deleteCharacterSkill, nil),
		routes.RegisterRoute(http.MethodGet, "character_skill/:character_skill", e.getCharacterSkill, nil),
		routes.RegisterRoute(http.MethodGet, "character_skills", e.listCharacterSkills, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getCharacterSkillsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_skill/:character_skill", e.updateCharacterSkill, nil),
		routes.RegisterRoute(http.MethodPut, "character_skill", e.createCharacterSkill, nil),
	}
}

// listCharacterSkills godoc
// @Id listCharacterSkills
// @Summary Lists CharacterSkills
// @Accept json
// @Produce json
// @Tags CharacterSkill
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterSkill
// @Failure 500 {string} string "Bad query request"
// @Router /character_skills [get]
func (e *CharacterSkillController) listCharacterSkills(c echo.Context) error {
	var results []models.CharacterSkill
	err := e.db.QueryContext(models.CharacterSkill{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterSkill godoc
// @Id getCharacterSkill
// @Summary Gets CharacterSkill
// @Accept json
// @Produce json
// @Tags CharacterSkill
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterSkill
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_skill/{id} [get]
func (e *CharacterSkillController) getCharacterSkill(c echo.Context) error {
	characterSkillId, err := strconv.Atoi(c.Param("character_skill"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterSkill
	err = e.db.QueryContext(models.CharacterSkill{}, c).First(&result, characterSkillId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterSkill godoc
// @Id updateCharacterSkill
// @Summary Updates CharacterSkill
// @Accept json
// @Produce json
// @Tags CharacterSkill
// @Param id path int true "Id"
// @Param character_skill body models.CharacterSkill true "CharacterSkill"
// @Success 200 {array} models.CharacterSkill
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_skill/{id} [patch]
func (e *CharacterSkillController) updateCharacterSkill(c echo.Context) error {
	characterSkill := new(models.CharacterSkill)
	if err := c.Bind(characterSkill); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).First(&models.CharacterSkill{}, characterSkill.ID).Error
	if err != nil || characterSkill.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).Updates(&characterSkill).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterSkill)
}

// createCharacterSkill godoc
// @Id createCharacterSkill
// @Summary Creates CharacterSkill
// @Accept json
// @Produce json
// @Param character_skill body models.CharacterSkill true "CharacterSkill"
// @Tags CharacterSkill
// @Success 200 {array} models.CharacterSkill
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_skill [put]
func (e *CharacterSkillController) createCharacterSkill(c echo.Context) error {
	characterSkill := new(models.CharacterSkill)
	if err := c.Bind(characterSkill); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).Create(&characterSkill).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, characterSkill)
}

// deleteCharacterSkill godoc
// @Id deleteCharacterSkill
// @Summary Deletes CharacterSkill
// @Accept json
// @Produce json
// @Tags CharacterSkill
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_skill/{id} [delete]
func (e *CharacterSkillController) deleteCharacterSkill(c echo.Context) error {
	characterSkillId, err := strconv.Atoi(c.Param("character_skill"))
	if err != nil {
		e.logger.Error(err)
	}

	characterSkill := new(models.CharacterSkill)
	err = e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).First(&characterSkill, characterSkillId).Error
	if err != nil || characterSkill.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).Delete(&characterSkill).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterSkillsBulk godoc
// @Id getCharacterSkillsBulk
// @Summary Gets CharacterSkills in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterSkill
// @Success 200 {array} models.CharacterSkill
// @Failure 500 {string} string "Bad query request"
// @Router /character_skills/bulk [post]
func (e *CharacterSkillController) getCharacterSkillsBulk(c echo.Context) error {
	var results []models.CharacterSkill

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.CharacterSkill{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
