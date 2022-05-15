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

type NpcSpellsEffectController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcSpellsEffectController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcSpellsEffectController {
	return &NpcSpellsEffectController{
		db:	 db,
		logger: logger,
	}
}

func (e *NpcSpellsEffectController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_spells_effect/:id", e.getNpcSpellsEffect, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells_effects", e.listNpcSpellsEffects, nil),
		routes.RegisterRoute(http.MethodPut, "npc_spells_effect", e.createNpcSpellsEffect, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_spells_effect/:id", e.deleteNpcSpellsEffect, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_spells_effect/:id", e.updateNpcSpellsEffect, nil),
		routes.RegisterRoute(http.MethodPost, "npc_spells_effects/bulk", e.getNpcSpellsEffectsBulk, nil),
	}
}

// listNpcSpellsEffects godoc
// @Id listNpcSpellsEffects
// @Summary Lists NpcSpellsEffects
// @Accept json
// @Produce json
// @Tags NpcSpellsEffect
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcSpellsEffectsEntries"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEffect
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_effects [get]
func (e *NpcSpellsEffectController) listNpcSpellsEffects(c echo.Context) error {
	var results []models.NpcSpellsEffect
	err := e.db.QueryContext(models.NpcSpellsEffect{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcSpellsEffect godoc
// @Id getNpcSpellsEffect
// @Summary Gets NpcSpellsEffect
// @Accept json
// @Produce json
// @Tags NpcSpellsEffect
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcSpellsEffectsEntries"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEffect
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_effect/{id} [get]
func (e *NpcSpellsEffectController) getNpcSpellsEffect(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.NpcSpellsEffect
	query := e.db.QueryContext(models.NpcSpellsEffect{}, c)
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

// updateNpcSpellsEffect godoc
// @Id updateNpcSpellsEffect
// @Summary Updates NpcSpellsEffect
// @Accept json
// @Produce json
// @Tags NpcSpellsEffect
// @Param id path int true "Id"
// @Param npc_spells_effect body models.NpcSpellsEffect true "NpcSpellsEffect"
// @Success 200 {array} models.NpcSpellsEffect
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_spells_effect/{id} [patch]
func (e *NpcSpellsEffectController) updateNpcSpellsEffect(c echo.Context) error {
	request := new(models.NpcSpellsEffect)
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

	// query builder
	var result models.NpcSpellsEffect
	query := e.db.QueryContext(models.NpcSpellsEffect{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.NpcSpellsEffect{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcSpellsEffect godoc
// @Id createNpcSpellsEffect
// @Summary Creates NpcSpellsEffect
// @Accept json
// @Produce json
// @Param npc_spells_effect body models.NpcSpellsEffect true "NpcSpellsEffect"
// @Tags NpcSpellsEffect
// @Success 200 {array} models.NpcSpellsEffect
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_spells_effect [put]
func (e *NpcSpellsEffectController) createNpcSpellsEffect(c echo.Context) error {
	npcSpellsEffect := new(models.NpcSpellsEffect)
	if err := c.Bind(npcSpellsEffect); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcSpellsEffect{}, c).Model(&models.NpcSpellsEffect{}).Create(&npcSpellsEffect).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcSpellsEffect)
}

// deleteNpcSpellsEffect godoc
// @Id deleteNpcSpellsEffect
// @Summary Deletes NpcSpellsEffect
// @Accept json
// @Produce json
// @Tags NpcSpellsEffect
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_spells_effect/{id} [delete]
func (e *NpcSpellsEffectController) deleteNpcSpellsEffect(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.NpcSpellsEffect
	query := e.db.QueryContext(models.NpcSpellsEffect{}, c)
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

// getNpcSpellsEffectsBulk godoc
// @Id getNpcSpellsEffectsBulk
// @Summary Gets NpcSpellsEffects in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcSpellsEffect
// @Success 200 {array} models.NpcSpellsEffect
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_effects/bulk [post]
func (e *NpcSpellsEffectController) getNpcSpellsEffectsBulk(c echo.Context) error {
	var results []models.NpcSpellsEffect

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

	err := e.db.QueryContext(models.NpcSpellsEffect{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
