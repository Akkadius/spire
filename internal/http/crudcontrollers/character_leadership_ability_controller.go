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

type CharacterLeadershipAbilityController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterLeadershipAbilityController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterLeadershipAbilityController {
	return &CharacterLeadershipAbilityController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterLeadershipAbilityController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_leadership_ability/:id", e.getCharacterLeadershipAbility, nil),
		routes.RegisterRoute(http.MethodGet, "character_leadership_abilities", e.listCharacterLeadershipAbilities, nil),
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
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

	err = e.db.QueryContext(models.CharacterLeadershipAbility{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
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

	err := e.db.Get(models.CharacterLeadershipAbility{}, c).Model(&models.CharacterLeadershipAbility{}).Create(&characterLeadershipAbility).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
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
