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

type CharacterSkillController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterSkillController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterSkillController {
	return &CharacterSkillController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterSkillController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_skill/:id", e.getCharacterSkill, nil),
		routes.RegisterRoute(http.MethodGet, "character_skills", e.listCharacterSkills, nil),
		routes.RegisterRoute(http.MethodPut, "character_skill", e.createCharacterSkill, nil),
		routes.RegisterRoute(http.MethodDelete, "character_skill/:id", e.deleteCharacterSkill, nil),
		routes.RegisterRoute(http.MethodPatch, "character_skill/:id", e.updateCharacterSkill, nil),
		routes.RegisterRoute(http.MethodPost, "character_skills/bulk", e.getCharacterSkillsBulk, nil),
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
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
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
// @Param id path int true "id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterSkill
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_skill/{id} [get]
func (e *CharacterSkillController) getCharacterSkill(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ID]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [skill_id] position [2] type [smallint]
	if len(c.QueryParam("skill_id")) > 0 {
		skillIdParam, err := strconv.Atoi(c.QueryParam("skill_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [skill_id] err [%s]", err.Error())})
		}

		params = append(params, skillIdParam)
		keys = append(keys, "skill_id = ?")
	}

	// query builder
	var result models.CharacterSkill
	query := e.db.QueryContext(models.CharacterSkill{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
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
// @Param ID path int true "ID"
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	entity := models.CharacterSkill{}
	err := e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).First(&entity, characterSkill.ID).Error
	if err != nil || characterSkill.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterSkill{}, c).Model(&entity).Select("*").Updates(&characterSkill).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).Create(&characterSkill).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
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
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_skill/{id} [delete]
func (e *CharacterSkillController) deleteCharacterSkill(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [skill_id] position [2] type [smallint]
	if len(c.QueryParam("skill_id")) > 0 {
		skillIdParam, err := strconv.Atoi(c.QueryParam("skill_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [skill_id] err [%s]", err.Error())})
		}

		params = append(params, skillIdParam)
		keys = append(keys, "skill_id = ?")
	}

	// query builder
	var result models.CharacterSkill
	query := e.db.QueryContext(models.CharacterSkill{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.CharacterSkill{}, c).Model(&models.CharacterSkill{}).Delete(&result).Error
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.CharacterSkill{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
