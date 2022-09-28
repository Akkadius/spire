package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterSpellController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterSpellController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterSpellController {
	return &CharacterSpellController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterSpellController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_spell/:id", e.getCharacterSpell, nil),
		routes.RegisterRoute(http.MethodGet, "character_spells", e.listCharacterSpells, nil),
		routes.RegisterRoute(http.MethodPut, "character_spell", e.createCharacterSpell, nil),
		routes.RegisterRoute(http.MethodDelete, "character_spell/:id", e.deleteCharacterSpell, nil),
		routes.RegisterRoute(http.MethodPatch, "character_spell/:id", e.updateCharacterSpell, nil),
		routes.RegisterRoute(http.MethodPost, "character_spells/bulk", e.getCharacterSpellsBulk, nil),
	}
}

// listCharacterSpells godoc
// @Id listCharacterSpells
// @Summary Lists CharacterSpells
// @Accept json
// @Produce json
// @Tags CharacterSpell
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterSpell
// @Failure 500 {string} string "Bad query request"
// @Router /character_spells [get]
func (e *CharacterSpellController) listCharacterSpells(c echo.Context) error {
	var results []models.CharacterSpell
	err := e.db.QueryContext(models.CharacterSpell{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterSpell godoc
// @Id getCharacterSpell
// @Summary Gets CharacterSpell
// @Accept json
// @Produce json
// @Tags CharacterSpell
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterSpell
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_spell/{id} [get]
func (e *CharacterSpellController) getCharacterSpell(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot_id] position [2] type [smallint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterSpell
	query := e.db.QueryContext(models.CharacterSpell{}, c)
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

// updateCharacterSpell godoc
// @Id updateCharacterSpell
// @Summary Updates CharacterSpell
// @Accept json
// @Produce json
// @Tags CharacterSpell
// @Param id path int true "Id"
// @Param character_spell body models.CharacterSpell true "CharacterSpell"
// @Success 200 {array} models.CharacterSpell
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_spell/{id} [patch]
func (e *CharacterSpellController) updateCharacterSpell(c echo.Context) error {
	request := new(models.CharacterSpell)
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

	// key param [slot_id] position [2] type [smallint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterSpell
	query := e.db.QueryContext(models.CharacterSpell{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterSpell{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterSpell godoc
// @Id createCharacterSpell
// @Summary Creates CharacterSpell
// @Accept json
// @Produce json
// @Param character_spell body models.CharacterSpell true "CharacterSpell"
// @Tags CharacterSpell
// @Success 200 {array} models.CharacterSpell
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_spell [put]
func (e *CharacterSpellController) createCharacterSpell(c echo.Context) error {
	characterSpell := new(models.CharacterSpell)
	if err := c.Bind(characterSpell); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterSpell{}, c).Model(&models.CharacterSpell{}).Create(&characterSpell).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterSpell)
}

// deleteCharacterSpell godoc
// @Id deleteCharacterSpell
// @Summary Deletes CharacterSpell
// @Accept json
// @Produce json
// @Tags CharacterSpell
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_spell/{id} [delete]
func (e *CharacterSpellController) deleteCharacterSpell(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot_id] position [2] type [smallint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterSpell
	query := e.db.QueryContext(models.CharacterSpell{}, c)
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

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterSpellsBulk godoc
// @Id getCharacterSpellsBulk
// @Summary Gets CharacterSpells in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterSpell
// @Success 200 {array} models.CharacterSpell
// @Failure 500 {string} string "Bad query request"
// @Router /character_spells/bulk [post]
func (e *CharacterSpellController) getCharacterSpellsBulk(c echo.Context) error {
	var results []models.CharacterSpell

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

	err := e.db.QueryContext(models.CharacterSpell{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
