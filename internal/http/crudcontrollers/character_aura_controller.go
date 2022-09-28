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

type CharacterAuraController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterAuraController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterAuraController {
	return &CharacterAuraController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterAuraController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_aura/:id", e.getCharacterAura, nil),
		routes.RegisterRoute(http.MethodGet, "character_auras", e.listCharacterAuras, nil),
		routes.RegisterRoute(http.MethodPut, "character_aura", e.createCharacterAura, nil),
		routes.RegisterRoute(http.MethodDelete, "character_aura/:id", e.deleteCharacterAura, nil),
		routes.RegisterRoute(http.MethodPatch, "character_aura/:id", e.updateCharacterAura, nil),
		routes.RegisterRoute(http.MethodPost, "character_auras/bulk", e.getCharacterAurasBulk, nil),
	}
}

// listCharacterAuras godoc
// @Id listCharacterAuras
// @Summary Lists CharacterAuras
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterAura
// @Failure 500 {string} string "Bad query request"
// @Router /character_auras [get]
func (e *CharacterAuraController) listCharacterAuras(c echo.Context) error {
	var results []models.CharacterAura
	err := e.db.QueryContext(models.CharacterAura{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterAura godoc
// @Id getCharacterAura
// @Summary Gets CharacterAura
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterAura
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_aura/{id} [get]
func (e *CharacterAuraController) getCharacterAura(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterAura
	query := e.db.QueryContext(models.CharacterAura{}, c)
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

// updateCharacterAura godoc
// @Id updateCharacterAura
// @Summary Updates CharacterAura
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param id path int true "Id"
// @Param character_aura body models.CharacterAura true "CharacterAura"
// @Success 200 {array} models.CharacterAura
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_aura/{id} [patch]
func (e *CharacterAuraController) updateCharacterAura(c echo.Context) error {
	request := new(models.CharacterAura)
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

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterAura
	query := e.db.QueryContext(models.CharacterAura{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterAura{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterAura godoc
// @Id createCharacterAura
// @Summary Creates CharacterAura
// @Accept json
// @Produce json
// @Param character_aura body models.CharacterAura true "CharacterAura"
// @Tags CharacterAura
// @Success 200 {array} models.CharacterAura
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_aura [put]
func (e *CharacterAuraController) createCharacterAura(c echo.Context) error {
	characterAura := new(models.CharacterAura)
	if err := c.Bind(characterAura); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterAura{}, c).Model(&models.CharacterAura{}).Create(&characterAura).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterAura)
}

// deleteCharacterAura godoc
// @Id deleteCharacterAura
// @Summary Deletes CharacterAura
// @Accept json
// @Produce json
// @Tags CharacterAura
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_aura/{id} [delete]
func (e *CharacterAuraController) deleteCharacterAura(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterAura
	query := e.db.QueryContext(models.CharacterAura{}, c)
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

// getCharacterAurasBulk godoc
// @Id getCharacterAurasBulk
// @Summary Gets CharacterAuras in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterAura
// @Success 200 {array} models.CharacterAura
// @Failure 500 {string} string "Bad query request"
// @Router /character_auras/bulk [post]
func (e *CharacterAuraController) getCharacterAurasBulk(c echo.Context) error {
	var results []models.CharacterAura

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

	err := e.db.QueryContext(models.CharacterAura{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
