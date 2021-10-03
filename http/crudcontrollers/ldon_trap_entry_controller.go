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

type LdonTrapEntryController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLdonTrapEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LdonTrapEntryController {
	return &LdonTrapEntryController{
		db:     db,
		logger: logger,
	}
}

func (e *LdonTrapEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "ldon_trap_entry/:ldon_trap_entry", e.deleteLdonTrapEntry, nil),
		routes.RegisterRoute(http.MethodGet, "ldon_trap_entry/:ldon_trap_entry", e.getLdonTrapEntry, nil),
		routes.RegisterRoute(http.MethodGet, "ldon_trap_entries", e.listLdonTrapEntries, nil),
		routes.RegisterRoute(http.MethodPost, "ldon_trap_entries/bulk", e.getLdonTrapEntriesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "ldon_trap_entry/:ldon_trap_entry", e.updateLdonTrapEntry, nil),
		routes.RegisterRoute(http.MethodPut, "ldon_trap_entry", e.createLdonTrapEntry, nil),
	}
}

// listLdonTrapEntries godoc
// @Id listLdonTrapEntries
// @Summary Lists LdonTrapEntries
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapEntry
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_entries [get]
func (e *LdonTrapEntryController) listLdonTrapEntries(c echo.Context) error {
	var results []models.LdonTrapEntry
	err := e.db.QueryContext(models.LdonTrapEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLdonTrapEntry godoc
// @Id getLdonTrapEntry
// @Summary Gets LdonTrapEntry
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_entry/{id} [get]
func (e *LdonTrapEntryController) getLdonTrapEntry(c echo.Context) error {
	ldonTrapEntryId, err := strconv.Atoi(c.Param("ldon_trap_entry"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.LdonTrapEntry
	err = e.db.QueryContext(models.LdonTrapEntry{}, c).First(&result, ldonTrapEntryId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLdonTrapEntry godoc
// @Id updateLdonTrapEntry
// @Summary Updates LdonTrapEntry
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param id path int true "Id"
// @Param ldon_trap_entry body models.LdonTrapEntry true "LdonTrapEntry"
// @Success 200 {array} models.LdonTrapEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /ldon_trap_entry/{id} [patch]
func (e *LdonTrapEntryController) updateLdonTrapEntry(c echo.Context) error {
	ldonTrapEntry := new(models.LdonTrapEntry)
	if err := c.Bind(ldonTrapEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.LdonTrapEntry{}
	err := e.db.Get(models.LdonTrapEntry{}, c).Model(&models.LdonTrapEntry{}).First(&entity, ldonTrapEntry.ID).Error
	if err != nil || ldonTrapEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LdonTrapEntry{}, c).Model(&entity).Updates(&ldonTrapEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, ldonTrapEntry)
}

// createLdonTrapEntry godoc
// @Id createLdonTrapEntry
// @Summary Creates LdonTrapEntry
// @Accept json
// @Produce json
// @Param ldon_trap_entry body models.LdonTrapEntry true "LdonTrapEntry"
// @Tags LdonTrapEntry
// @Success 200 {array} models.LdonTrapEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /ldon_trap_entry [put]
func (e *LdonTrapEntryController) createLdonTrapEntry(c echo.Context) error {
	ldonTrapEntry := new(models.LdonTrapEntry)
	if err := c.Bind(ldonTrapEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.LdonTrapEntry{}, c).Model(&models.LdonTrapEntry{}).Create(&ldonTrapEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, ldonTrapEntry)
}

// deleteLdonTrapEntry godoc
// @Id deleteLdonTrapEntry
// @Summary Deletes LdonTrapEntry
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /ldon_trap_entry/{id} [delete]
func (e *LdonTrapEntryController) deleteLdonTrapEntry(c echo.Context) error {
	ldonTrapEntryId, err := strconv.Atoi(c.Param("ldon_trap_entry"))
	if err != nil {
		e.logger.Error(err)
	}

	ldonTrapEntry := new(models.LdonTrapEntry)
	err = e.db.Get(models.LdonTrapEntry{}, c).Model(&models.LdonTrapEntry{}).First(&ldonTrapEntry, ldonTrapEntryId).Error
	if err != nil || ldonTrapEntry.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LdonTrapEntry{}, c).Model(&models.LdonTrapEntry{}).Delete(&ldonTrapEntry).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLdonTrapEntriesBulk godoc
// @Id getLdonTrapEntriesBulk
// @Summary Gets LdonTrapEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LdonTrapEntry
// @Success 200 {array} models.LdonTrapEntry
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_entries/bulk [post]
func (e *LdonTrapEntryController) getLdonTrapEntriesBulk(c echo.Context) error {
	var results []models.LdonTrapEntry

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

	err := e.db.QueryContext(models.LdonTrapEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
