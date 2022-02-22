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

type NpcSpellController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcSpellController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcSpellController {
	return &NpcSpellController{
		db:	 db,
		logger: logger,
	}
}

func (e *NpcSpellController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_spell/:id", e.getNpcSpell, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells", e.listNpcSpells, nil),
		routes.RegisterRoute(http.MethodPut, "npc_spell", e.createNpcSpell, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_spell/:id", e.deleteNpcSpell, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_spell/:id", e.updateNpcSpell, nil),
		routes.RegisterRoute(http.MethodPost, "npc_spells/bulk", e.getNpcSpellsBulk, nil),
	}
}

// listNpcSpells godoc
// @Id listNpcSpells
// @Summary Lists NpcSpells
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcSpellsEntries"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpell
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells [get]
func (e *NpcSpellController) listNpcSpells(c echo.Context) error {
	var results []models.NpcSpell
	err := e.db.QueryContext(models.NpcSpell{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcSpell godoc
// @Id getNpcSpell
// @Summary Gets NpcSpell
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcSpellsEntries"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpell
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spell/{id} [get]
func (e *NpcSpellController) getNpcSpell(c echo.Context) error {
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
	var result models.NpcSpell
	query := e.db.QueryContext(models.NpcSpell{}, c)
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

// updateNpcSpell godoc
// @Id updateNpcSpell
// @Summary Updates NpcSpell
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param id path int true "Id"
// @Param npc_spell body models.NpcSpell true "NpcSpell"
// @Success 200 {array} models.NpcSpell
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_spell/{id} [patch]
func (e *NpcSpellController) updateNpcSpell(c echo.Context) error {
	request := new(models.NpcSpell)
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
	var result models.NpcSpell
	query := e.db.QueryContext(models.NpcSpell{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcSpell godoc
// @Id createNpcSpell
// @Summary Creates NpcSpell
// @Accept json
// @Produce json
// @Param npc_spell body models.NpcSpell true "NpcSpell"
// @Tags NpcSpell
// @Success 200 {array} models.NpcSpell
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_spell [put]
func (e *NpcSpellController) createNpcSpell(c echo.Context) error {
	npcSpell := new(models.NpcSpell)
	if err := c.Bind(npcSpell); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcSpell{}, c).Model(&models.NpcSpell{}).Create(&npcSpell).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcSpell)
}

// deleteNpcSpell godoc
// @Id deleteNpcSpell
// @Summary Deletes NpcSpell
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_spell/{id} [delete]
func (e *NpcSpellController) deleteNpcSpell(c echo.Context) error {
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
	var result models.NpcSpell
	query := e.db.QueryContext(models.NpcSpell{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.NpcSpell{}, c).Model(&models.NpcSpell{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcSpellsBulk godoc
// @Id getNpcSpellsBulk
// @Summary Gets NpcSpells in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcSpell
// @Success 200 {array} models.NpcSpell
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells/bulk [post]
func (e *NpcSpellController) getNpcSpellsBulk(c echo.Context) error {
	var results []models.NpcSpell

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

	err := e.db.QueryContext(models.NpcSpell{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
