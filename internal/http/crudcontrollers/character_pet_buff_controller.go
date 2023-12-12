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

type CharacterPetBuffController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterPetBuffController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterPetBuffController {
	return &CharacterPetBuffController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterPetBuffController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_pet_buff/:charId", e.getCharacterPetBuff, nil),
		routes.RegisterRoute(http.MethodGet, "character_pet_buffs", e.listCharacterPetBuffs, nil),
		routes.RegisterRoute(http.MethodGet, "character_pet_buffs/count", e.getCharacterPetBuffsCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_pet_buff", e.createCharacterPetBuff, nil),
		routes.RegisterRoute(http.MethodDelete, "character_pet_buff/:charId", e.deleteCharacterPetBuff, nil),
		routes.RegisterRoute(http.MethodPatch, "character_pet_buff/:charId", e.updateCharacterPetBuff, nil),
		routes.RegisterRoute(http.MethodPost, "character_pet_buffs/bulk", e.getCharacterPetBuffsBulk, nil),
	}
}

// listCharacterPetBuffs godoc
// @Id listCharacterPetBuffs
// @Summary Lists CharacterPetBuffs
// @Accept json
// @Produce json
// @Tags CharacterPetBuff
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetBuff
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_buffs [get]
func (e *CharacterPetBuffController) listCharacterPetBuffs(c echo.Context) error {
	var results []models.CharacterPetBuff
	err := e.db.QueryContext(models.CharacterPetBuff{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPetBuff godoc
// @Id getCharacterPetBuff
// @Summary Gets CharacterPetBuff
// @Accept json
// @Produce json
// @Tags CharacterPetBuff
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetBuff
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_buff/{id} [get]
func (e *CharacterPetBuffController) getCharacterPetBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [pet] position [2] type [int]
	if len(c.QueryParam("pet")) > 0 {
		petParam, err := strconv.Atoi(c.QueryParam("pet"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [pet] err [%s]", err.Error())})
		}

		params = append(params, petParam)
		keys = append(keys, "pet = ?")
	}

	// key param [slot] position [3] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterPetBuff
	query := e.db.QueryContext(models.CharacterPetBuff{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterPetBuff godoc
// @Id updateCharacterPetBuff
// @Summary Updates CharacterPetBuff
// @Accept json
// @Produce json
// @Tags CharacterPetBuff
// @Param id path int true "Id"
// @Param character_pet_buff body models.CharacterPetBuff true "CharacterPetBuff"
// @Success 200 {array} models.CharacterPetBuff
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_pet_buff/{id} [patch]
func (e *CharacterPetBuffController) updateCharacterPetBuff(c echo.Context) error {
	request := new(models.CharacterPetBuff)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [pet] position [2] type [int]
	if len(c.QueryParam("pet")) > 0 {
		petParam, err := strconv.Atoi(c.QueryParam("pet"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [pet] err [%s]", err.Error())})
		}

		params = append(params, petParam)
		keys = append(keys, "pet = ?")
	}

	// key param [slot] position [3] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterPetBuff
	query := e.db.QueryContext(models.CharacterPetBuff{}, c)
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
		event := fmt.Sprintf("Updated [CharacterPetBuff] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterPetBuff godoc
// @Id createCharacterPetBuff
// @Summary Creates CharacterPetBuff
// @Accept json
// @Produce json
// @Param character_pet_buff body models.CharacterPetBuff true "CharacterPetBuff"
// @Tags CharacterPetBuff
// @Success 200 {array} models.CharacterPetBuff
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_pet_buff [put]
func (e *CharacterPetBuffController) createCharacterPetBuff(c echo.Context) error {
	characterPetBuff := new(models.CharacterPetBuff)
	if err := c.Bind(characterPetBuff); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterPetBuff{}, c).Model(&models.CharacterPetBuff{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterPetBuff).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterPetBuff{}, characterPetBuff)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterPetBuff] [%v] data [%v]", characterPetBuff.CharId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterPetBuff)
}

// deleteCharacterPetBuff godoc
// @Id deleteCharacterPetBuff
// @Summary Deletes CharacterPetBuff
// @Accept json
// @Produce json
// @Tags CharacterPetBuff
// @Param id path int true "charId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_pet_buff/{id} [delete]
func (e *CharacterPetBuffController) deleteCharacterPetBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [pet] position [2] type [int]
	if len(c.QueryParam("pet")) > 0 {
		petParam, err := strconv.Atoi(c.QueryParam("pet"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [pet] err [%s]", err.Error())})
		}

		params = append(params, petParam)
		keys = append(keys, "pet = ?")
	}

	// key param [slot] position [3] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterPetBuff
	query := e.db.QueryContext(models.CharacterPetBuff{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterPetBuff] [%v] keys [%v]", result.CharId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterPetBuffsBulk godoc
// @Id getCharacterPetBuffsBulk
// @Summary Gets CharacterPetBuffs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterPetBuff
// @Success 200 {array} models.CharacterPetBuff
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_buffs/bulk [post]
func (e *CharacterPetBuffController) getCharacterPetBuffsBulk(c echo.Context) error {
	var results []models.CharacterPetBuff

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

	err := e.db.QueryContext(models.CharacterPetBuff{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPetBuffsCount godoc
// @Id getCharacterPetBuffsCount
// @Summary Counts CharacterPetBuffs
// @Accept json
// @Produce json
// @Tags CharacterPetBuff
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetBuff
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_buffs/count [get]
func (e *CharacterPetBuffController) getCharacterPetBuffsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterPetBuff{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}