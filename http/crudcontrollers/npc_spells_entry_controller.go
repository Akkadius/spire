package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type NpcSpellsEntryController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcSpellsEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcSpellsEntryController {
	return &NpcSpellsEntryController{
		db:     db,
		logger: logger,
	}
}

func (e *NpcSpellsEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "npc_spells_entry/:npc_spells_entry", e.deleteNpcSpellsEntry, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells_entry/:npc_spells_entry", e.getNpcSpellsEntry, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells_entries", e.listNpcSpellsEntries, nil),
		routes.RegisterRoute(http.MethodPost, "npc_spells_entries/bulk", e.getNpcSpellsEntriesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_spells_entry/:npc_spells_entry", e.updateNpcSpellsEntry, nil),
		routes.RegisterRoute(http.MethodPut, "npc_spells_entry", e.createNpcSpellsEntry, nil),
	}
}

// listNpcSpellsEntries godoc
// @Id listNpcSpellsEntries
// @Summary Lists NpcSpellsEntries
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_entries [get]
func (e *NpcSpellsEntryController) listNpcSpellsEntries(c echo.Context) error {
	var results []models.NpcSpellsEntry
	err := e.db.QueryContext(models.NpcSpellsEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcSpellsEntry godoc
// @Id getNpcSpellsEntry
// @Summary Gets NpcSpellsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_entry/{id} [get]
func (e *NpcSpellsEntryController) getNpcSpellsEntry(c echo.Context) error {
	npcSpellsEntryId, err := strconv.Atoi(c.Param("npc_spells_entry"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.NpcSpellsEntry
	err = e.db.QueryContext(models.NpcSpellsEntry{}, c).First(&result, npcSpellsEntryId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateNpcSpellsEntry godoc
// @Id updateNpcSpellsEntry
// @Summary Updates NpcSpellsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param id path int true "Id"
// @Param npc_spells_entry body models.NpcSpellsEntry true "NpcSpellsEntry"
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_spells_entry/{id} [patch]
func (e *NpcSpellsEntryController) updateNpcSpellsEntry(c echo.Context) error {
	npcSpellsEntry := new(models.NpcSpellsEntry)
	if err := c.Bind(npcSpellsEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.NpcSpellsEntry{}
	err := e.db.Get(models.NpcSpellsEntry{}, c).Model(&models.NpcSpellsEntry{}).First(&entity, npcSpellsEntry.ID).Error
	if err != nil || npcSpellsEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.NpcSpellsEntry{}, c).Model(&entity).Updates(&npcSpellsEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, npcSpellsEntry)
}

// createNpcSpellsEntry godoc
// @Id createNpcSpellsEntry
// @Summary Creates NpcSpellsEntry
// @Accept json
// @Produce json
// @Param npc_spells_entry body models.NpcSpellsEntry true "NpcSpellsEntry"
// @Tags NpcSpellsEntry
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_spells_entry [put]
func (e *NpcSpellsEntryController) createNpcSpellsEntry(c echo.Context) error {
	npcSpellsEntry := new(models.NpcSpellsEntry)
	if err := c.Bind(npcSpellsEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.NpcSpellsEntry{}, c).Model(&models.NpcSpellsEntry{}).Create(&npcSpellsEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, npcSpellsEntry)
}

// deleteNpcSpellsEntry godoc
// @Id deleteNpcSpellsEntry
// @Summary Deletes NpcSpellsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_spells_entry/{id} [delete]
func (e *NpcSpellsEntryController) deleteNpcSpellsEntry(c echo.Context) error {
	npcSpellsEntryId, err := strconv.Atoi(c.Param("npc_spells_entry"))
	if err != nil {
		e.logger.Error(err)
	}

	npcSpellsEntry := new(models.NpcSpellsEntry)
	err = e.db.Get(models.NpcSpellsEntry{}, c).Model(&models.NpcSpellsEntry{}).First(&npcSpellsEntry, npcSpellsEntryId).Error
	if err != nil || npcSpellsEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.NpcSpellsEntry{}, c).Model(&models.NpcSpellsEntry{}).Delete(&npcSpellsEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcSpellsEntriesBulk godoc
// @Id getNpcSpellsEntriesBulk
// @Summary Gets NpcSpellsEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcSpellsEntry
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_entries/bulk [post]
func (e *NpcSpellsEntryController) getNpcSpellsEntriesBulk(c echo.Context) error {
	var results []models.NpcSpellsEntry

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

	err := e.db.QueryContext(models.NpcSpellsEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
