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

type BotPetInventoryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotPetInventoryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotPetInventoryController {
	return &BotPetInventoryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotPetInventoryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_pet_inventory/:petInventoriesIndex", e.getBotPetInventory, nil),
		routes.RegisterRoute(http.MethodGet, "bot_pet_inventories", e.listBotPetInventories, nil),
		routes.RegisterRoute(http.MethodGet, "bot_pet_inventories/count", e.getBotPetInventoriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_pet_inventory", e.createBotPetInventory, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_pet_inventory/:petInventoriesIndex", e.deleteBotPetInventory, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_pet_inventory/:petInventoriesIndex", e.updateBotPetInventory, nil),
		routes.RegisterRoute(http.MethodPost, "bot_pet_inventories/bulk", e.getBotPetInventoriesBulk, nil),
	}
}

// listBotPetInventories godoc
// @Id listBotPetInventories
// @Summary Lists BotPetInventories
// @Accept json
// @Produce json
// @Tags BotPetInventory
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotPetInventory
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pet_inventories [get]
func (e *BotPetInventoryController) listBotPetInventories(c echo.Context) error {
	var results []models.BotPetInventory
	err := e.db.QueryContext(models.BotPetInventory{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotPetInventory godoc
// @Id getBotPetInventory
// @Summary Gets BotPetInventory
// @Accept json
// @Produce json
// @Tags BotPetInventory
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotPetInventory
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pet_inventory/{id} [get]
func (e *BotPetInventoryController) getBotPetInventory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	petInventoriesIndex, err := strconv.Atoi(c.Param("petInventoriesIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PetInventoriesIndex]"})
	}
	params = append(params, petInventoriesIndex)
	keys = append(keys, "pet_inventories_index = ?")

	// query builder
	var result models.BotPetInventory
	query := e.db.QueryContext(models.BotPetInventory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.PetInventoriesIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotPetInventory godoc
// @Id updateBotPetInventory
// @Summary Updates BotPetInventory
// @Accept json
// @Produce json
// @Tags BotPetInventory
// @Param id path int true "Id"
// @Param bot_pet_inventory body models.BotPetInventory true "BotPetInventory"
// @Success 200 {array} models.BotPetInventory
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_pet_inventory/{id} [patch]
func (e *BotPetInventoryController) updateBotPetInventory(c echo.Context) error {
	request := new(models.BotPetInventory)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	petInventoriesIndex, err := strconv.Atoi(c.Param("petInventoriesIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PetInventoriesIndex]"})
	}
	params = append(params, petInventoriesIndex)
	keys = append(keys, "pet_inventories_index = ?")

	// query builder
	var result models.BotPetInventory
	query := e.db.QueryContext(models.BotPetInventory{}, c)
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
		event := fmt.Sprintf("Updated [BotPetInventory] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotPetInventory godoc
// @Id createBotPetInventory
// @Summary Creates BotPetInventory
// @Accept json
// @Produce json
// @Param bot_pet_inventory body models.BotPetInventory true "BotPetInventory"
// @Tags BotPetInventory
// @Success 200 {array} models.BotPetInventory
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_pet_inventory [put]
func (e *BotPetInventoryController) createBotPetInventory(c echo.Context) error {
	botPetInventory := new(models.BotPetInventory)
	if err := c.Bind(botPetInventory); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotPetInventory{}, c).Model(&models.BotPetInventory{}).Create(&botPetInventory).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotPetInventory{}, botPetInventory)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotPetInventory] [%v] data [%v]", botPetInventory.PetInventoriesIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botPetInventory)
}

// deleteBotPetInventory godoc
// @Id deleteBotPetInventory
// @Summary Deletes BotPetInventory
// @Accept json
// @Produce json
// @Tags BotPetInventory
// @Param id path int true "petInventoriesIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_pet_inventory/{id} [delete]
func (e *BotPetInventoryController) deleteBotPetInventory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	petInventoriesIndex, err := strconv.Atoi(c.Param("petInventoriesIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, petInventoriesIndex)
	keys = append(keys, "pet_inventories_index = ?")

	// query builder
	var result models.BotPetInventory
	query := e.db.QueryContext(models.BotPetInventory{}, c)
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
		event := fmt.Sprintf("Deleted [BotPetInventory] [%v] keys [%v]", result.PetInventoriesIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotPetInventoriesBulk godoc
// @Id getBotPetInventoriesBulk
// @Summary Gets BotPetInventories in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotPetInventory
// @Success 200 {array} models.BotPetInventory
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pet_inventories/bulk [post]
func (e *BotPetInventoryController) getBotPetInventoriesBulk(c echo.Context) error {
	var results []models.BotPetInventory

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

	err := e.db.QueryContext(models.BotPetInventory{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotPetInventoriesCount godoc
// @Id getBotPetInventoriesCount
// @Summary Counts BotPetInventories
// @Accept json
// @Produce json
// @Tags BotPetInventory
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotPetInventory
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pet_inventories/count [get]
func (e *BotPetInventoryController) getBotPetInventoriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotPetInventory{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}