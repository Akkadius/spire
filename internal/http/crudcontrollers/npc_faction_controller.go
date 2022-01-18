package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type NpcFactionController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcFactionController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcFactionController {
	return &NpcFactionController{
		db:     db,
		logger: logger,
	}
}

func (e *NpcFactionController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "npc_faction/:npc_faction", e.deleteNpcFaction, nil),
		routes.RegisterRoute(http.MethodGet, "npc_faction/:npc_faction", e.getNpcFaction, nil),
		routes.RegisterRoute(http.MethodGet, "npc_factions", e.listNpcFactions, nil),
		routes.RegisterRoute(http.MethodPost, "npc_factions/bulk", e.getNpcFactionsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_faction/:npc_faction", e.updateNpcFaction, nil),
		routes.RegisterRoute(http.MethodPut, "npc_faction", e.createNpcFaction, nil),
	}
}

// listNpcFactions godoc
// @Id listNpcFactions
// @Summary Lists NpcFactions
// @Accept json
// @Produce json
// @Tags NpcFaction
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcFactionEntries"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcFaction
// @Failure 500 {string} string "Bad query request"
// @Router /npc_factions [get]
func (e *NpcFactionController) listNpcFactions(c echo.Context) error {
	var results []models.NpcFaction
	err := e.db.QueryContext(models.NpcFaction{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcFaction godoc
// @Id getNpcFaction
// @Summary Gets NpcFaction
// @Accept json
// @Produce json
// @Tags NpcFaction
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcFactionEntries"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcFaction
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_faction/{id} [get]
func (e *NpcFactionController) getNpcFaction(c echo.Context) error {
	npcFactionId, err := strconv.Atoi(c.Param("npc_faction"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.NpcFaction
	err = e.db.QueryContext(models.NpcFaction{}, c).First(&result, npcFactionId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateNpcFaction godoc
// @Id updateNpcFaction
// @Summary Updates NpcFaction
// @Accept json
// @Produce json
// @Tags NpcFaction
// @Param id path int true "Id"
// @Param npc_faction body models.NpcFaction true "NpcFaction"
// @Success 200 {array} models.NpcFaction
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_faction/{id} [patch]
func (e *NpcFactionController) updateNpcFaction(c echo.Context) error {
	npcFaction := new(models.NpcFaction)
	if err := c.Bind(npcFaction); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.NpcFaction{}
	err := e.db.Get(models.NpcFaction{}, c).Model(&models.NpcFaction{}).First(&entity, npcFaction.ID).Error
	if err != nil || npcFaction.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.NpcFaction{}, c).Model(&entity).Select("*").Updates(&npcFaction).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, npcFaction)
}

// createNpcFaction godoc
// @Id createNpcFaction
// @Summary Creates NpcFaction
// @Accept json
// @Produce json
// @Param npc_faction body models.NpcFaction true "NpcFaction"
// @Tags NpcFaction
// @Success 200 {array} models.NpcFaction
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_faction [put]
func (e *NpcFactionController) createNpcFaction(c echo.Context) error {
	npcFaction := new(models.NpcFaction)
	if err := c.Bind(npcFaction); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.NpcFaction{}, c).Model(&models.NpcFaction{}).Create(&npcFaction).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, npcFaction)
}

// deleteNpcFaction godoc
// @Id deleteNpcFaction
// @Summary Deletes NpcFaction
// @Accept json
// @Produce json
// @Tags NpcFaction
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_faction/{id} [delete]
func (e *NpcFactionController) deleteNpcFaction(c echo.Context) error {
	npcFactionId, err := strconv.Atoi(c.Param("npc_faction"))
	if err != nil {
		e.logger.Error(err)
	}

	npcFaction := new(models.NpcFaction)
	err = e.db.Get(models.NpcFaction{}, c).Model(&models.NpcFaction{}).First(&npcFaction, npcFactionId).Error
	if err != nil || npcFaction.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.NpcFaction{}, c).Model(&models.NpcFaction{}).Delete(&npcFaction).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcFactionsBulk godoc
// @Id getNpcFactionsBulk
// @Summary Gets NpcFactions in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcFaction
// @Success 200 {array} models.NpcFaction
// @Failure 500 {string} string "Bad query request"
// @Router /npc_factions/bulk [post]
func (e *NpcFactionController) getNpcFactionsBulk(c echo.Context) error {
	var results []models.NpcFaction

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.NpcFaction{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
