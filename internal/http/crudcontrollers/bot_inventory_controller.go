package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type BotInventoryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotInventoryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotInventoryController {
	return &BotInventoryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotInventoryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_inventory/:inventoriesIndex", e.getBotInventory, nil),
		routes.RegisterRoute(http.MethodGet, "bot_inventories", e.listBotInventories, nil),
		routes.RegisterRoute(http.MethodGet, "bot_inventories/count", e.getBotInventoriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_inventory", e.createBotInventory, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_inventory/:inventoriesIndex", e.deleteBotInventory, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_inventory/:inventoriesIndex", e.updateBotInventory, nil),
		routes.RegisterRoute(http.MethodPost, "bot_inventories/bulk", e.getBotInventoriesBulk, nil),
	}
}

// listBotInventories godoc
// @Id listBotInventories
// @Summary Lists BotInventories
// @Accept json
// @Produce json
// @Tags BotInventory
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotInventory
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inventories [get]
func (e *BotInventoryController) listBotInventories(c echo.Context) error {
	var results []models.BotInventory
	err := e.db.QueryContext(models.BotInventory{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotInventory godoc
// @Id getBotInventory
// @Summary Gets BotInventory
// @Accept json
// @Produce json
// @Tags BotInventory
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotInventory
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inventory/{id} [get]
func (e *BotInventoryController) getBotInventory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	inventoriesIndex, err := strconv.Atoi(c.Param("inventoriesIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [InventoriesIndex]"})
	}
	params = append(params, inventoriesIndex)
	keys = append(keys, "inventories_index = ?")

	// query builder
	var result models.BotInventory
	query := e.db.QueryContext(models.BotInventory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.InventoriesIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotInventory godoc
// @Id updateBotInventory
// @Summary Updates BotInventory
// @Accept json
// @Produce json
// @Tags BotInventory
// @Param id path int true "Id"
// @Param bot_inventory body models.BotInventory true "BotInventory"
// @Success 200 {array} models.BotInventory
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_inventory/{id} [patch]
func (e *BotInventoryController) updateBotInventory(c echo.Context) error {
	request := new(models.BotInventory)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	inventoriesIndex, err := strconv.Atoi(c.Param("inventoriesIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [InventoriesIndex]"})
	}
	params = append(params, inventoriesIndex)
	keys = append(keys, "inventories_index = ?")

	// query builder
	var result models.BotInventory
	query := e.db.QueryContext(models.BotInventory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [BotInventory] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotInventory godoc
// @Id createBotInventory
// @Summary Creates BotInventory
// @Accept json
// @Produce json
// @Param bot_inventory body models.BotInventory true "BotInventory"
// @Tags BotInventory
// @Success 200 {array} models.BotInventory
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_inventory [put]
func (e *BotInventoryController) createBotInventory(c echo.Context) error {
	botInventory := new(models.BotInventory)
	if err := c.Bind(botInventory); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotInventory{}, c).Model(&models.BotInventory{}).Create(&botInventory).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotInventory{}, botInventory)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotInventory] [%v] data [%v]", botInventory.InventoriesIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botInventory)
}

// deleteBotInventory godoc
// @Id deleteBotInventory
// @Summary Deletes BotInventory
// @Accept json
// @Produce json
// @Tags BotInventory
// @Param id path int true "inventoriesIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_inventory/{id} [delete]
func (e *BotInventoryController) deleteBotInventory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	inventoriesIndex, err := strconv.Atoi(c.Param("inventoriesIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, inventoriesIndex)
	keys = append(keys, "inventories_index = ?")

	// query builder
	var result models.BotInventory
	query := e.db.QueryContext(models.BotInventory{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [BotInventory] [%v] keys [%v]", result.InventoriesIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotInventoriesBulk godoc
// @Id getBotInventoriesBulk
// @Summary Gets BotInventories in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotInventory
// @Success 200 {array} models.BotInventory
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inventories/bulk [post]
func (e *BotInventoryController) getBotInventoriesBulk(c echo.Context) error {
	var results []models.BotInventory

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

	err := e.db.QueryContext(models.BotInventory{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotInventoriesCount godoc
// @Id getBotInventoriesCount
// @Summary Counts BotInventories
// @Accept json
// @Produce json
// @Tags BotInventory
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotInventory
// @Failure 500 {string} string "Bad query request"
// @Router /bot_inventories/count [get]
func (e *BotInventoryController) getBotInventoriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotInventory{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}