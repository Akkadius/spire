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

type NpcTypeController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcTypeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcTypeController {
	return &NpcTypeController{
		db:	 db,
		logger: logger,
	}
}

func (e *NpcTypeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_type/:id", e.getNpcType, nil),
		routes.RegisterRoute(http.MethodGet, "npc_types", e.listNpcTypes, nil),
		routes.RegisterRoute(http.MethodPut, "npc_type", e.createNpcType, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_type/:id", e.deleteNpcType, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_type/:id", e.updateNpcType, nil),
		routes.RegisterRoute(http.MethodPost, "npc_types/bulk", e.getNpcTypesBulk, nil),
	}
}

// listNpcTypes godoc
// @Id listNpcTypes
// @Summary Lists NpcTypes
// @Accept json
// @Produce json
// @Tags NpcType
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrency<br>Merchantlists<br>NpcEmotes<br>NpcFactions<br>NpcFactions.NpcFactionEntries<br>NpcSpells<br>NpcSpells.NpcSpellsEntries<br>NpcTypesTint<br>Spawnentries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcType
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types [get]
func (e *NpcTypeController) listNpcTypes(c echo.Context) error {
	var results []models.NpcType
	err := e.db.QueryContext(models.NpcType{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcType godoc
// @Id getNpcType
// @Summary Gets NpcType
// @Accept json
// @Produce json
// @Tags NpcType
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrency<br>Merchantlists<br>NpcEmotes<br>NpcFactions<br>NpcFactions.NpcFactionEntries<br>NpcSpells<br>NpcSpells.NpcSpellsEntries<br>NpcTypesTint<br>Spawnentries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcType
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_type/{id} [get]
func (e *NpcTypeController) getNpcType(c echo.Context) error {
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
	var result models.NpcType
	query := e.db.QueryContext(models.NpcType{}, c)
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

// updateNpcType godoc
// @Id updateNpcType
// @Summary Updates NpcType
// @Accept json
// @Produce json
// @Tags NpcType
// @Param id path int true "Id"
// @Param npc_type body models.NpcType true "NpcType"
// @Success 200 {array} models.NpcType
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_type/{id} [patch]
func (e *NpcTypeController) updateNpcType(c echo.Context) error {
	request := new(models.NpcType)
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
	var result models.NpcType
	query := e.db.QueryContext(models.NpcType{}, c)
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

// createNpcType godoc
// @Id createNpcType
// @Summary Creates NpcType
// @Accept json
// @Produce json
// @Param npc_type body models.NpcType true "NpcType"
// @Tags NpcType
// @Success 200 {array} models.NpcType
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_type [put]
func (e *NpcTypeController) createNpcType(c echo.Context) error {
	npcType := new(models.NpcType)
	if err := c.Bind(npcType); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcType{}, c).Model(&models.NpcType{}).Create(&npcType).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcType)
}

// deleteNpcType godoc
// @Id deleteNpcType
// @Summary Deletes NpcType
// @Accept json
// @Produce json
// @Tags NpcType
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_type/{id} [delete]
func (e *NpcTypeController) deleteNpcType(c echo.Context) error {
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
	var result models.NpcType
	query := e.db.QueryContext(models.NpcType{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.NpcType{}, c).Model(&models.NpcType{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcTypesBulk godoc
// @Id getNpcTypesBulk
// @Summary Gets NpcTypes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcType
// @Success 200 {array} models.NpcType
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types/bulk [post]
func (e *NpcTypeController) getNpcTypesBulk(c echo.Context) error {
	var results []models.NpcType

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

	err := e.db.QueryContext(models.NpcType{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
