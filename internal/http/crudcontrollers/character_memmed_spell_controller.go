package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CharacterMemmedSpellController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterMemmedSpellController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterMemmedSpellController {
	return &CharacterMemmedSpellController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterMemmedSpellController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_memmed_spell/:id", e.getCharacterMemmedSpell, nil),
		routes.RegisterRoute(http.MethodGet, "character_memmed_spells", e.listCharacterMemmedSpells, nil),
		routes.RegisterRoute(http.MethodPut, "character_memmed_spell", e.createCharacterMemmedSpell, nil),
		routes.RegisterRoute(http.MethodDelete, "character_memmed_spell/:id", e.deleteCharacterMemmedSpell, nil),
		routes.RegisterRoute(http.MethodPatch, "character_memmed_spell/:id", e.updateCharacterMemmedSpell, nil),
		routes.RegisterRoute(http.MethodPost, "character_memmed_spells/bulk", e.getCharacterMemmedSpellsBulk, nil),
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
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
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
	var result models.CharacterMemmedSpell
	query := e.db.QueryContext(models.CharacterMemmedSpell{}, c)
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
	request := new(models.CharacterMemmedSpell)
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
	var result models.CharacterMemmedSpell
	query := e.db.QueryContext(models.CharacterMemmedSpell{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterMemmedSpell{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterMemmedSpell{}, c).Model(&models.CharacterMemmedSpell{}).Create(&characterMemmedSpell).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
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
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_memmed_spell/{id} [delete]
func (e *CharacterMemmedSpellController) deleteCharacterMemmedSpell(c echo.Context) error {
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
	var result models.CharacterMemmedSpell
	query := e.db.QueryContext(models.CharacterMemmedSpell{}, c)
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
