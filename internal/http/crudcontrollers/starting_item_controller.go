package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type StartingItemController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewStartingItemController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *StartingItemController {
	return &StartingItemController{
		db:     db,
		logger: logger,
	}
}

func (e *StartingItemController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "starting_item/:starting_item", e.deleteStartingItem, nil),
		routes.RegisterRoute(http.MethodGet, "starting_item/:starting_item", e.getStartingItem, nil),
		routes.RegisterRoute(http.MethodGet, "starting_items", e.listStartingItems, nil),
		routes.RegisterRoute(http.MethodPost, "starting_items/bulk", e.getStartingItemsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "starting_item/:starting_item", e.updateStartingItem, nil),
		routes.RegisterRoute(http.MethodPut, "starting_item", e.createStartingItem, nil),
	}
}

// listStartingItems godoc
// @Id listStartingItems
// @Summary Lists StartingItems
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.StartingItem
// @Failure 500 {string} string "Bad query request"
// @Router /starting_items [get]
func (e *StartingItemController) listStartingItems(c echo.Context) error {
	var results []models.StartingItem
	err := e.db.QueryContext(models.StartingItem{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getStartingItem godoc
// @Id getStartingItem
// @Summary Gets StartingItem
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.StartingItem
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /starting_item/{id} [get]
func (e *StartingItemController) getStartingItem(c echo.Context) error {
	startingItemId, err := strconv.Atoi(c.Param("starting_item"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.StartingItem
	err = e.db.QueryContext(models.StartingItem{}, c).First(&result, startingItemId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateStartingItem godoc
// @Id updateStartingItem
// @Summary Updates StartingItem
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param id path int true "Id"
// @Param starting_item body models.StartingItem true "StartingItem"
// @Success 200 {array} models.StartingItem
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /starting_item/{id} [patch]
func (e *StartingItemController) updateStartingItem(c echo.Context) error {
	startingItem := new(models.StartingItem)
	if err := c.Bind(startingItem); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.StartingItem{}
	err := e.db.Get(models.StartingItem{}, c).Model(&models.StartingItem{}).First(&entity, startingItem.ID).Error
	if err != nil || startingItem.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.StartingItem{}, c).Model(&entity).Updates(&startingItem).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, startingItem)
}

// createStartingItem godoc
// @Id createStartingItem
// @Summary Creates StartingItem
// @Accept json
// @Produce json
// @Param starting_item body models.StartingItem true "StartingItem"
// @Tags StartingItem
// @Success 200 {array} models.StartingItem
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /starting_item [put]
func (e *StartingItemController) createStartingItem(c echo.Context) error {
	startingItem := new(models.StartingItem)
	if err := c.Bind(startingItem); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.StartingItem{}, c).Model(&models.StartingItem{}).Create(&startingItem).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, startingItem)
}

// deleteStartingItem godoc
// @Id deleteStartingItem
// @Summary Deletes StartingItem
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /starting_item/{id} [delete]
func (e *StartingItemController) deleteStartingItem(c echo.Context) error {
	startingItemId, err := strconv.Atoi(c.Param("starting_item"))
	if err != nil {
		e.logger.Error(err)
	}

	startingItem := new(models.StartingItem)
	err = e.db.Get(models.StartingItem{}, c).Model(&models.StartingItem{}).First(&startingItem, startingItemId).Error
	if err != nil || startingItem.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.StartingItem{}, c).Model(&models.StartingItem{}).Delete(&startingItem).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getStartingItemsBulk godoc
// @Id getStartingItemsBulk
// @Summary Gets StartingItems in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags StartingItem
// @Success 200 {array} models.StartingItem
// @Failure 500 {string} string "Bad query request"
// @Router /starting_items/bulk [post]
func (e *StartingItemController) getStartingItemsBulk(c echo.Context) error {
	var results []models.StartingItem

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

	err := e.db.QueryContext(models.StartingItem{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
