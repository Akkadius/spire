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

type InventoryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewInventoryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *InventoryController {
	return &InventoryController{
		db:	 db,
		logger: logger,
	}
}

func (e *InventoryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "inventory/:charid", e.getInventory, nil),
		routes.RegisterRoute(http.MethodGet, "inventories", e.listInventories, nil),
		routes.RegisterRoute(http.MethodPut, "inventory", e.createInventory, nil),
		routes.RegisterRoute(http.MethodDelete, "inventory/:charid", e.deleteInventory, nil),
		routes.RegisterRoute(http.MethodPatch, "inventory/:charid", e.updateInventory, nil),
		routes.RegisterRoute(http.MethodPost, "inventories/bulk", e.getInventoriesBulk, nil),
	}
}

// listInventories godoc
// @Id listInventories
// @Summary Lists Inventories
// @Accept json
// @Produce json
// @Tags Inventory
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.DiscoveredItems"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Inventory
// @Failure 500 {string} string "Bad query request"
// @Router /inventories [get]
func (e *InventoryController) listInventories(c echo.Context) error {
	var results []models.Inventory
	err := e.db.QueryContext(models.Inventory{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getInventory godoc
// @Id getInventory
// @Summary Gets Inventory
// @Accept json
// @Produce json
// @Tags Inventory
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.DiscoveredItems"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Inventory
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /inventory/{id} [get]
func (e *InventoryController) getInventory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [slotid] position [2] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.Inventory
	query := e.db.QueryContext(models.Inventory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Charid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateInventory godoc
// @Id updateInventory
// @Summary Updates Inventory
// @Accept json
// @Produce json
// @Tags Inventory
// @Param id path int true "Id"
// @Param inventory body models.Inventory true "Inventory"
// @Success 200 {array} models.Inventory
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /inventory/{id} [patch]
func (e *InventoryController) updateInventory(c echo.Context) error {
	request := new(models.Inventory)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [slotid] position [2] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.Inventory
	query := e.db.QueryContext(models.Inventory{}, c)
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

// createInventory godoc
// @Id createInventory
// @Summary Creates Inventory
// @Accept json
// @Produce json
// @Param inventory body models.Inventory true "Inventory"
// @Tags Inventory
// @Success 200 {array} models.Inventory
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /inventory [put]
func (e *InventoryController) createInventory(c echo.Context) error {
	inventory := new(models.Inventory)
	if err := c.Bind(inventory); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Inventory{}, c).Model(&models.Inventory{}).Create(&inventory).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, inventory)
}

// deleteInventory godoc
// @Id deleteInventory
// @Summary Deletes Inventory
// @Accept json
// @Produce json
// @Tags Inventory
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /inventory/{id} [delete]
func (e *InventoryController) deleteInventory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [slotid] position [2] type [mediumint]
	if len(c.QueryParam("slotid")) > 0 {
		slotidParam, err := strconv.Atoi(c.QueryParam("slotid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slotid] err [%s]", err.Error())})
		}

		params = append(params, slotidParam)
		keys = append(keys, "slotid = ?")
	}

	// query builder
	var result models.Inventory
	query := e.db.QueryContext(models.Inventory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.Inventory{}, c).Model(&models.Inventory{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getInventoriesBulk godoc
// @Id getInventoriesBulk
// @Summary Gets Inventories in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Inventory
// @Success 200 {array} models.Inventory
// @Failure 500 {string} string "Bad query request"
// @Router /inventories/bulk [post]
func (e *InventoryController) getInventoriesBulk(c echo.Context) error {
	var results []models.Inventory

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

	err := e.db.QueryContext(models.Inventory{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
