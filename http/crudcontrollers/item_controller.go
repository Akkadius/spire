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

type ItemController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewItemController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ItemController {
	return &ItemController{
		db:     db,
		logger: logger,
	}
}

func (e *ItemController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "item/:item", e.deleteItem, nil),
		routes.RegisterRoute(http.MethodGet, "item/:item", e.getItem, nil),
		routes.RegisterRoute(http.MethodGet, "items", e.listItems, nil),
		routes.RegisterRoute(http.MethodPost, "items/bulk", e.getItemsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "item/:item", e.updateItem, nil),
		routes.RegisterRoute(http.MethodPut, "item", e.createItem, nil),
	}
}

// listItems godoc
// @Id listItems
// @Summary Lists Items
// @Accept json
// @Produce json
// @Tags Item
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>DiscoveredItems"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Item
// @Failure 500 {string} string "Bad query request"
// @Router /items [get]
func (e *ItemController) listItems(c echo.Context) error {
	var results []models.Item
	err := e.db.QueryContext(models.Item{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getItem godoc
// @Id getItem
// @Summary Gets Item
// @Accept json
// @Produce json
// @Tags Item
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>DiscoveredItems"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Item
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /item/{id} [get]
func (e *ItemController) getItem(c echo.Context) error {
	itemId, err := strconv.Atoi(c.Param("item"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Item
	err = e.db.QueryContext(models.Item{}, c).First(&result, itemId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateItem godoc
// @Id updateItem
// @Summary Updates Item
// @Accept json
// @Produce json
// @Tags Item
// @Param id path int true "Id"
// @Param item body models.Item true "Item"
// @Success 200 {array} models.Item
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /item/{id} [patch]
func (e *ItemController) updateItem(c echo.Context) error {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Item{}
	err := e.db.Get(models.Item{}, c).Model(&models.Item{}).First(&entity, item.ID).Error
	if err != nil || item.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Item{}, c).Model(&entity).Updates(&item).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, item)
}

// createItem godoc
// @Id createItem
// @Summary Creates Item
// @Accept json
// @Produce json
// @Param item body models.Item true "Item"
// @Tags Item
// @Success 200 {array} models.Item
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /item [put]
func (e *ItemController) createItem(c echo.Context) error {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Item{}, c).Model(&models.Item{}).Create(&item).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, item)
}

// deleteItem godoc
// @Id deleteItem
// @Summary Deletes Item
// @Accept json
// @Produce json
// @Tags Item
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /item/{id} [delete]
func (e *ItemController) deleteItem(c echo.Context) error {
	itemId, err := strconv.Atoi(c.Param("item"))
	if err != nil {
		e.logger.Error(err)
	}

	item := new(models.Item)
	err = e.db.Get(models.Item{}, c).Model(&models.Item{}).First(&item, itemId).Error
	if err != nil || item.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Item{}, c).Model(&models.Item{}).Delete(&item).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getItemsBulk godoc
// @Id getItemsBulk
// @Summary Gets Items in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Item
// @Success 200 {array} models.Item
// @Failure 500 {string} string "Bad query request"
// @Router /items/bulk [post]
func (e *ItemController) getItemsBulk(c echo.Context) error {
	var results []models.Item

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

	err := e.db.QueryContext(models.Item{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
