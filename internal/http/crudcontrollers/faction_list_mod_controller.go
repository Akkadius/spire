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

type FactionListModController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewFactionListModController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *FactionListModController {
	return &FactionListModController{
		db:     db,
		logger: logger,
	}
}

func (e *FactionListModController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "faction_list_mod/:faction_list_mod", e.deleteFactionListMod, nil),
		routes.RegisterRoute(http.MethodGet, "faction_list_mod/:faction_list_mod", e.getFactionListMod, nil),
		routes.RegisterRoute(http.MethodGet, "faction_list_mods", e.listFactionListMods, nil),
		routes.RegisterRoute(http.MethodPost, "faction_list_mods/bulk", e.getFactionListModsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "faction_list_mod/:faction_list_mod", e.updateFactionListMod, nil),
		routes.RegisterRoute(http.MethodPut, "faction_list_mod", e.createFactionListMod, nil),
	}
}

// listFactionListMods godoc
// @Id listFactionListMods
// @Summary Lists FactionListMods
// @Accept json
// @Produce json
// @Tags FactionListMod
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionListMod
// @Failure 500 {string} string "Bad query request"
// @Router /faction_list_mods [get]
func (e *FactionListModController) listFactionListMods(c echo.Context) error {
	var results []models.FactionListMod
	err := e.db.QueryContext(models.FactionListMod{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getFactionListMod godoc
// @Id getFactionListMod
// @Summary Gets FactionListMod
// @Accept json
// @Produce json
// @Tags FactionListMod
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionListMod
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /faction_list_mod/{id} [get]
func (e *FactionListModController) getFactionListMod(c echo.Context) error {
	factionListModId, err := strconv.Atoi(c.Param("faction_list_mod"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.FactionListMod
	err = e.db.QueryContext(models.FactionListMod{}, c).First(&result, factionListModId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateFactionListMod godoc
// @Id updateFactionListMod
// @Summary Updates FactionListMod
// @Accept json
// @Produce json
// @Tags FactionListMod
// @Param id path int true "Id"
// @Param faction_list_mod body models.FactionListMod true "FactionListMod"
// @Success 200 {array} models.FactionListMod
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /faction_list_mod/{id} [patch]
func (e *FactionListModController) updateFactionListMod(c echo.Context) error {
	factionListMod := new(models.FactionListMod)
	if err := c.Bind(factionListMod); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.FactionListMod{}
	err := e.db.Get(models.FactionListMod{}, c).Model(&models.FactionListMod{}).First(&entity, factionListMod.ID).Error
	if err != nil || factionListMod.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.FactionListMod{}, c).Model(&entity).Select("*").Updates(&factionListMod).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, factionListMod)
}

// createFactionListMod godoc
// @Id createFactionListMod
// @Summary Creates FactionListMod
// @Accept json
// @Produce json
// @Param faction_list_mod body models.FactionListMod true "FactionListMod"
// @Tags FactionListMod
// @Success 200 {array} models.FactionListMod
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /faction_list_mod [put]
func (e *FactionListModController) createFactionListMod(c echo.Context) error {
	factionListMod := new(models.FactionListMod)
	if err := c.Bind(factionListMod); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.FactionListMod{}, c).Model(&models.FactionListMod{}).Create(&factionListMod).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, factionListMod)
}

// deleteFactionListMod godoc
// @Id deleteFactionListMod
// @Summary Deletes FactionListMod
// @Accept json
// @Produce json
// @Tags FactionListMod
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /faction_list_mod/{id} [delete]
func (e *FactionListModController) deleteFactionListMod(c echo.Context) error {
	factionListModId, err := strconv.Atoi(c.Param("faction_list_mod"))
	if err != nil {
		e.logger.Error(err)
	}

	factionListMod := new(models.FactionListMod)
	err = e.db.Get(models.FactionListMod{}, c).Model(&models.FactionListMod{}).First(&factionListMod, factionListModId).Error
	if err != nil || factionListMod.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.FactionListMod{}, c).Model(&models.FactionListMod{}).Delete(&factionListMod).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getFactionListModsBulk godoc
// @Id getFactionListModsBulk
// @Summary Gets FactionListMods in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags FactionListMod
// @Success 200 {array} models.FactionListMod
// @Failure 500 {string} string "Bad query request"
// @Router /faction_list_mods/bulk [post]
func (e *FactionListModController) getFactionListModsBulk(c echo.Context) error {
	var results []models.FactionListMod

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

	err := e.db.QueryContext(models.FactionListMod{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
