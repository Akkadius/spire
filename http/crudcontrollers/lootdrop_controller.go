package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type LootdropController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLootdropController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LootdropController {
	return &LootdropController{
		db:     db,
		logger: logger,
	}
}

func (e *LootdropController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "lootdrop/:lootdrop", e.deleteLootdrop, nil),
		routes.RegisterRoute(http.MethodGet, "lootdrop/:lootdrop", e.getLootdrop, nil),
		routes.RegisterRoute(http.MethodGet, "lootdrops", e.listLootdrops, nil),
		routes.RegisterRoute(http.MethodPost, "lootdrops/bulk", e.getLootdropsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "lootdrop/:lootdrop", e.updateLootdrop, nil),
		routes.RegisterRoute(http.MethodPut, "lootdrop", e.createLootdrop, nil),
	}
}

// listLootdrops godoc
// @Id listLootdrops
// @Summary Lists Lootdrops
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Item.DiscoveredItems"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Lootdrop
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrops [get]
func (e *LootdropController) listLootdrops(c echo.Context) error {
	var results []models.Lootdrop
	err := e.db.QueryContext(models.Lootdrop{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLootdrop godoc
// @Id getLootdrop
// @Summary Gets Lootdrop
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Item.DiscoveredItems"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Lootdrop
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop/{id} [get]
func (e *LootdropController) getLootdrop(c echo.Context) error {
	lootdropId, err := strconv.Atoi(c.Param("lootdrop"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Lootdrop
	err = e.db.QueryContext(models.Lootdrop{}, c).First(&result, lootdropId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLootdrop godoc
// @Id updateLootdrop
// @Summary Updates Lootdrop
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param id path int true "Id"
// @Param lootdrop body models.Lootdrop true "Lootdrop"
// @Success 200 {array} models.Lootdrop
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /lootdrop/{id} [patch]
func (e *LootdropController) updateLootdrop(c echo.Context) error {
	lootdrop := new(models.Lootdrop)
	if err := c.Bind(lootdrop); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Lootdrop{}
	err := e.db.Get(models.Lootdrop{}, c).Model(&models.Lootdrop{}).First(&entity, lootdrop.ID).Error
	if err != nil || lootdrop.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Lootdrop{}, c).Model(&entity).Updates(&lootdrop).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, lootdrop)
}

// createLootdrop godoc
// @Id createLootdrop
// @Summary Creates Lootdrop
// @Accept json
// @Produce json
// @Param lootdrop body models.Lootdrop true "Lootdrop"
// @Tags Lootdrop
// @Success 200 {array} models.Lootdrop
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /lootdrop [put]
func (e *LootdropController) createLootdrop(c echo.Context) error {
	lootdrop := new(models.Lootdrop)
	if err := c.Bind(lootdrop); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Lootdrop{}, c).Model(&models.Lootdrop{}).Create(&lootdrop).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, lootdrop)
}

// deleteLootdrop godoc
// @Id deleteLootdrop
// @Summary Deletes Lootdrop
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /lootdrop/{id} [delete]
func (e *LootdropController) deleteLootdrop(c echo.Context) error {
	lootdropId, err := strconv.Atoi(c.Param("lootdrop"))
	if err != nil {
		e.logger.Error(err)
	}

	lootdrop := new(models.Lootdrop)
	err = e.db.Get(models.Lootdrop{}, c).Model(&models.Lootdrop{}).First(&lootdrop, lootdropId).Error
	if err != nil || lootdrop.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Lootdrop{}, c).Model(&models.Lootdrop{}).Delete(&lootdrop).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLootdropsBulk godoc
// @Id getLootdropsBulk
// @Summary Gets Lootdrops in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Lootdrop
// @Success 200 {array} models.Lootdrop
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrops/bulk [post]
func (e *LootdropController) getLootdropsBulk(c echo.Context) error {
	var results []models.Lootdrop

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

	err := e.db.QueryContext(models.Lootdrop{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
