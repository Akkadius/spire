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
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type CharacterLeadershipAbilityController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterLeadershipAbilityController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterLeadershipAbilityController {
	return &CharacterLeadershipAbilityController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterLeadershipAbilityController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_leadership_ability/:id", e.getCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodGet, "character_leadership_abilities", e.listCharacterLeadershipAbilities, nil),
		routes.RegisterRoute(http.MethodGet, "character_leadership_abilities/count", e.getCharacterLeadershipAbilitiesCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_leadership_ability", e.createCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodDelete, "character_leadership_ability/:id", e.deleteCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodPatch, "character_leadership_ability/:id", e.updateCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodPost, "character_leadership_abilities/bulk", e.getCharacterLeadershipAbilitiesBulk, nil),
	}
}

// listCharacterLeadershipAbilities godoc
// @Id listCharacterLeadershipAbilities
// @Summary Lists CharacterLeadershipAbilities
// @Accept json
// @Produce json
// @Tags CharacterLeadershipAbility
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterLeadershipAbility
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_leadership_ability/{id} [get]
func (e *CharacterLeadershipAbilityController) getCharacterLeadershipAbility(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot] position [2] type [smallint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterLeadershipAbility
	query := e.db.QueryContext(models.CharacterLeadershipAbility{}, c)
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
	request := new(models.CharacterLeadershipAbility)
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

	// key param [slot] position [2] type [smallint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterLeadershipAbility
	query := e.db.QueryContext(models.CharacterLeadershipAbility{}, c)
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
		event := fmt.Sprintf("Updated [CharacterLeadershipAbility] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterLeadershipAbility{}, c).Model(&models.CharacterLeadershipAbility{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
        db = db.Omit(clause.Associations)
    }

	err := db.Create(&characterLeadershipAbility).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterLeadershipAbility{}, characterLeadershipAbility)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterLeadershipAbility] [%v] data [%v]", characterLeadershipAbility.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterLeadershipAbility)
}

// deleteCharacterLeadershipAbility godoc
// @Id deleteCharacterLeadershipAbility
// @Summary Deletes CharacterLeadershipAbility
// @Accept json
// @Produce json
// @Tags CharacterLeadershipAbility
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_leadership_ability/{id} [delete]
func (e *CharacterLeadershipAbilityController) deleteCharacterLeadershipAbility(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot] position [2] type [smallint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterLeadershipAbility
	query := e.db.QueryContext(models.CharacterLeadershipAbility{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterLeadershipAbility] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterLeadershipAbilitiesCount godoc
// @Id getCharacterLeadershipAbilitiesCount
// @Summary Counts CharacterLeadershipAbilities
// @Accept json
// @Produce json
// @Tags CharacterLeadershipAbility
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterLeadershipAbility
// @Failure 500 {string} string "Bad query request"
// @Router /character_leadership_abilities/count [get]
func (e *CharacterLeadershipAbilityController) getCharacterLeadershipAbilitiesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterLeadershipAbility{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}