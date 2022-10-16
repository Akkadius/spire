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

type CharacterExpModifierController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterExpModifierController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterExpModifierController {
	return &CharacterExpModifierController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterExpModifierController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_exp_modifier/:characterId", e.getCharacterExpModifier, nil),
		routes.RegisterRoute(http.MethodGet, "character_exp_modifiers", e.listCharacterExpModifiers, nil),
		routes.RegisterRoute(http.MethodPut, "character_exp_modifier", e.createCharacterExpModifier, nil),
		routes.RegisterRoute(http.MethodDelete, "character_exp_modifier/:characterId", e.deleteCharacterExpModifier, nil),
		routes.RegisterRoute(http.MethodPatch, "character_exp_modifier/:characterId", e.updateCharacterExpModifier, nil),
		routes.RegisterRoute(http.MethodPost, "character_exp_modifiers/bulk", e.getCharacterExpModifiersBulk, nil),
	}
}

// listCharacterExpModifiers godoc
// @Id listCharacterExpModifiers
// @Summary Lists CharacterExpModifiers
// @Accept json
// @Produce json
// @Tags CharacterExpModifier
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterExpModifier
// @Failure 500 {string} string "Bad query request"
// @Router /character_exp_modifiers [get]
func (e *CharacterExpModifierController) listCharacterExpModifiers(c echo.Context) error {
	var results []models.CharacterExpModifier
	err := e.db.QueryContext(models.CharacterExpModifier{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterExpModifier godoc
// @Id getCharacterExpModifier
// @Summary Gets CharacterExpModifier
// @Accept json
// @Produce json
// @Tags CharacterExpModifier
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterExpModifier
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_exp_modifier/{id} [get]
func (e *CharacterExpModifierController) getCharacterExpModifier(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharacterId]"})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [zone_id] position [2] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// query builder
	var result models.CharacterExpModifier
	query := e.db.QueryContext(models.CharacterExpModifier{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharacterId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterExpModifier godoc
// @Id updateCharacterExpModifier
// @Summary Updates CharacterExpModifier
// @Accept json
// @Produce json
// @Tags CharacterExpModifier
// @Param id path int true "Id"
// @Param character_exp_modifier body models.CharacterExpModifier true "CharacterExpModifier"
// @Success 200 {array} models.CharacterExpModifier
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_exp_modifier/{id} [patch]
func (e *CharacterExpModifierController) updateCharacterExpModifier(c echo.Context) error {
	request := new(models.CharacterExpModifier)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharacterId]"})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [zone_id] position [2] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// query builder
	var result models.CharacterExpModifier
	query := e.db.QueryContext(models.CharacterExpModifier{}, c)
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
		event := fmt.Sprintf("Updated [CharacterExpModifier] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterExpModifier godoc
// @Id createCharacterExpModifier
// @Summary Creates CharacterExpModifier
// @Accept json
// @Produce json
// @Param character_exp_modifier body models.CharacterExpModifier true "CharacterExpModifier"
// @Tags CharacterExpModifier
// @Success 200 {array} models.CharacterExpModifier
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_exp_modifier [put]
func (e *CharacterExpModifierController) createCharacterExpModifier(c echo.Context) error {
	characterExpModifier := new(models.CharacterExpModifier)
	if err := c.Bind(characterExpModifier); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterExpModifier{}, c).Model(&models.CharacterExpModifier{}).Create(&characterExpModifier).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterExpModifier{}, characterExpModifier)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterExpModifier] [%v] data [%v]", characterExpModifier.CharacterId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterExpModifier)
}

// deleteCharacterExpModifier godoc
// @Id deleteCharacterExpModifier
// @Summary Deletes CharacterExpModifier
// @Accept json
// @Produce json
// @Tags CharacterExpModifier
// @Param id path int true "characterId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_exp_modifier/{id} [delete]
func (e *CharacterExpModifierController) deleteCharacterExpModifier(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [zone_id] position [2] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// query builder
	var result models.CharacterExpModifier
	query := e.db.QueryContext(models.CharacterExpModifier{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterExpModifier] [%v] keys [%v]", result.CharacterId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterExpModifiersBulk godoc
// @Id getCharacterExpModifiersBulk
// @Summary Gets CharacterExpModifiers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterExpModifier
// @Success 200 {array} models.CharacterExpModifier
// @Failure 500 {string} string "Bad query request"
// @Router /character_exp_modifiers/bulk [post]
func (e *CharacterExpModifierController) getCharacterExpModifiersBulk(c echo.Context) error {
	var results []models.CharacterExpModifier

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

	err := e.db.QueryContext(models.CharacterExpModifier{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
