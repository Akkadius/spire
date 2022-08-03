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

type NpcEmoteController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcEmoteController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcEmoteController {
	return &NpcEmoteController{
		db:	 db,
		logger: logger,
	}
}

func (e *NpcEmoteController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_emote/:id", e.getNpcEmote, nil),
		routes.RegisterRoute(http.MethodGet, "npc_emotes", e.listNpcEmotes, nil),
		routes.RegisterRoute(http.MethodPut, "npc_emote", e.createNpcEmote, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_emote/:id", e.deleteNpcEmote, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_emote/:id", e.updateNpcEmote, nil),
		routes.RegisterRoute(http.MethodPost, "npc_emotes/bulk", e.getNpcEmotesBulk, nil),
	}
}

// listNpcEmotes godoc
// @Id listNpcEmotes
// @Summary Lists NpcEmotes
// @Accept json
// @Produce json
// @Tags NpcEmote
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcEmote
// @Failure 500 {string} string "Bad query request"
// @Router /npc_emotes [get]
func (e *NpcEmoteController) listNpcEmotes(c echo.Context) error {
	var results []models.NpcEmote
	err := e.db.QueryContext(models.NpcEmote{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcEmote godoc
// @Id getNpcEmote
// @Summary Gets NpcEmote
// @Accept json
// @Produce json
// @Tags NpcEmote
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcEmote
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_emote/{id} [get]
func (e *NpcEmoteController) getNpcEmote(c echo.Context) error {
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
	var result models.NpcEmote
	query := e.db.QueryContext(models.NpcEmote{}, c)
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

// updateNpcEmote godoc
// @Id updateNpcEmote
// @Summary Updates NpcEmote
// @Accept json
// @Produce json
// @Tags NpcEmote
// @Param id path int true "Id"
// @Param npc_emote body models.NpcEmote true "NpcEmote"
// @Success 200 {array} models.NpcEmote
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_emote/{id} [patch]
func (e *NpcEmoteController) updateNpcEmote(c echo.Context) error {
	request := new(models.NpcEmote)
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
	var result models.NpcEmote
	query := e.db.QueryContext(models.NpcEmote{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.NpcEmote{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcEmote godoc
// @Id createNpcEmote
// @Summary Creates NpcEmote
// @Accept json
// @Produce json
// @Param npc_emote body models.NpcEmote true "NpcEmote"
// @Tags NpcEmote
// @Success 200 {array} models.NpcEmote
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_emote [put]
func (e *NpcEmoteController) createNpcEmote(c echo.Context) error {
	npcEmote := new(models.NpcEmote)
	if err := c.Bind(npcEmote); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcEmote{}, c).Model(&models.NpcEmote{}).Create(&npcEmote).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcEmote)
}

// deleteNpcEmote godoc
// @Id deleteNpcEmote
// @Summary Deletes NpcEmote
// @Accept json
// @Produce json
// @Tags NpcEmote
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_emote/{id} [delete]
func (e *NpcEmoteController) deleteNpcEmote(c echo.Context) error {
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
	var result models.NpcEmote
	query := e.db.QueryContext(models.NpcEmote{}, c)
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

// getNpcEmotesBulk godoc
// @Id getNpcEmotesBulk
// @Summary Gets NpcEmotes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcEmote
// @Success 200 {array} models.NpcEmote
// @Failure 500 {string} string "Bad query request"
// @Router /npc_emotes/bulk [post]
func (e *NpcEmoteController) getNpcEmotesBulk(c echo.Context) error {
	var results []models.NpcEmote

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

	err := e.db.QueryContext(models.NpcEmote{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
