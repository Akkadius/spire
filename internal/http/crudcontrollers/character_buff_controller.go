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

type CharacterBuffController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterBuffController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterBuffController {
	return &CharacterBuffController{
		db:	    db,
		logger: logger,
	}
}

func (e *CharacterBuffController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_buff/:characterId", e.getCharacterBuff, nil),
		routes.RegisterRoute(http.MethodGet, "character_buffs", e.listCharacterBuffs, nil),
		routes.RegisterRoute(http.MethodPut, "character_buff", e.createCharacterBuff, nil),
		routes.RegisterRoute(http.MethodDelete, "character_buff/:characterId", e.deleteCharacterBuff, nil),
		routes.RegisterRoute(http.MethodPatch, "character_buff/:characterId", e.updateCharacterBuff, nil),
		routes.RegisterRoute(http.MethodPost, "character_buffs/bulk", e.getCharacterBuffsBulk, nil),
	}
}

// listCharacterBuffs godoc
// @Id listCharacterBuffs
// @Summary Lists CharacterBuffs
// @Accept json
// @Produce json
// @Tags CharacterBuff
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterBuff
// @Failure 500 {string} string "Bad query request"
// @Router /character_buffs [get]
func (e *CharacterBuffController) listCharacterBuffs(c echo.Context) error {
	var results []models.CharacterBuff
	err := e.db.QueryContext(models.CharacterBuff{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterBuff godoc
// @Id getCharacterBuff
// @Summary Gets CharacterBuff
// @Accept json
// @Produce json
// @Tags CharacterBuff
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterBuff
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_buff/{id} [get]
func (e *CharacterBuffController) getCharacterBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharacterId]"})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [slot_id] position [2] type [tinyint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterBuff
	query := e.db.QueryContext(models.CharacterBuff{}, c)
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

// updateCharacterBuff godoc
// @Id updateCharacterBuff
// @Summary Updates CharacterBuff
// @Accept json
// @Produce json
// @Tags CharacterBuff
// @Param id path int true "Id"
// @Param character_buff body models.CharacterBuff true "CharacterBuff"
// @Success 200 {array} models.CharacterBuff
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_buff/{id} [patch]
func (e *CharacterBuffController) updateCharacterBuff(c echo.Context) error {
	request := new(models.CharacterBuff)
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

	// key param [slot_id] position [2] type [tinyint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterBuff
	query := e.db.QueryContext(models.CharacterBuff{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterBuff godoc
// @Id createCharacterBuff
// @Summary Creates CharacterBuff
// @Accept json
// @Produce json
// @Param character_buff body models.CharacterBuff true "CharacterBuff"
// @Tags CharacterBuff
// @Success 200 {array} models.CharacterBuff
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_buff [put]
func (e *CharacterBuffController) createCharacterBuff(c echo.Context) error {
	characterBuff := new(models.CharacterBuff)
	if err := c.Bind(characterBuff); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterBuff{}, c).Model(&models.CharacterBuff{}).Create(&characterBuff).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterBuff)
}

// deleteCharacterBuff godoc
// @Id deleteCharacterBuff
// @Summary Deletes CharacterBuff
// @Accept json
// @Produce json
// @Tags CharacterBuff
// @Param id path int true "characterId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_buff/{id} [delete]
func (e *CharacterBuffController) deleteCharacterBuff(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [slot_id] position [2] type [tinyint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterBuff
	query := e.db.QueryContext(models.CharacterBuff{}, c)
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

// getCharacterBuffsBulk godoc
// @Id getCharacterBuffsBulk
// @Summary Gets CharacterBuffs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterBuff
// @Success 200 {array} models.CharacterBuff
// @Failure 500 {string} string "Bad query request"
// @Router /character_buffs/bulk [post]
func (e *CharacterBuffController) getCharacterBuffsBulk(c echo.Context) error {
	var results []models.CharacterBuff

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

	err := e.db.QueryContext(models.CharacterBuff{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
