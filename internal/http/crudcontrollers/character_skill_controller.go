package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type CharacterSkillController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterSkillController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterSkillController {
	return &CharacterSkillController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterSkillController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_skill/:id", e.getCharacterSkill, nil),
		routes.RegisterRoute(http.MethodGet, "character_skills", e.listCharacterSkills, nil),
		routes.RegisterRoute(http.MethodGet, "character_skills/count", e.getCharacterSkillsCount, nil),
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
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
// @Param id path int true "Id"
// @Param character_skill body models.CharacterSkill true "CharacterSkill"
// @Success 200 {array} models.CharacterSkill
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_skill/{id} [patch]
func (e *CharacterSkillController) updateCharacterSkill(c echo.Context) error {
	request := new(models.CharacterSkill)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [CharacterSkill] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
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

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterSkill{}, characterSkill)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterSkill] [%v] data [%v]", characterSkill.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
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

	err = query.Limit(10000).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [CharacterSkill] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
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

// getCharacterSkillsCount godoc
// @Id getCharacterSkillsCount
// @Summary Counts CharacterSkills
// @Accept json
// @Produce json
// @Tags CharacterSkill
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterSkill
// @Failure 500 {string} string "Bad query request"
// @Router /character_skills/count [get]
func (e *CharacterSkillController) getCharacterSkillsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterSkill{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}