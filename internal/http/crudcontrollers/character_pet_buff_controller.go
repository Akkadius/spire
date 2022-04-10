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

type CharacterPetBuffController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterPetBuffController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterPetBuffController {
	return &CharacterPetBuffController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterPetBuffController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_pet_buff/:charId", e.getCharacterPetBuff, nil),
		routes.RegisterRoute(http.MethodGet, "character_pet_buffs", e.listCharacterPetBuffs, nil),
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
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

	err = e.db.QueryContext(models.CharacterPetBuff{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
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

	err := e.db.Get(models.CharacterPetBuff{}, c).Model(&models.CharacterPetBuff{}).Create(&characterPetBuff).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
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
