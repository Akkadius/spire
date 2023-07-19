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
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type DiscoveredItemController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewDiscoveredItemController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *DiscoveredItemController {
	return &DiscoveredItemController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *DiscoveredItemController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "discovered_item/:itemId", e.getDiscoveredItem, nil),
		routes.RegisterRoute(http.MethodGet, "discovered_items", e.listDiscoveredItems, nil),
		routes.RegisterRoute(http.MethodGet, "discovered_items/count", e.getDiscoveredItemsCount, nil),
		routes.RegisterRoute(http.MethodPut, "discovered_item", e.createDiscoveredItem, nil),
		routes.RegisterRoute(http.MethodDelete, "discovered_item/:itemId", e.deleteDiscoveredItem, nil),
		routes.RegisterRoute(http.MethodPatch, "discovered_item/:itemId", e.updateDiscoveredItem, nil),
		routes.RegisterRoute(http.MethodPost, "discovered_items/bulk", e.getDiscoveredItemsBulk, nil),
	}
}

// listDiscoveredItems godoc
// @Id listDiscoveredItems
// @Summary Lists DiscoveredItems
// @Accept json
// @Produce json
// @Tags DiscoveredItem
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DiscoveredItem
// @Failure 500 {string} string "Bad query request"
// @Router /discovered_items [get]
func (e *DiscoveredItemController) listDiscoveredItems(c echo.Context) error {
	var results []models.DiscoveredItem
	err := e.db.QueryContext(models.DiscoveredItem{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDiscoveredItem godoc
// @Id getDiscoveredItem
// @Summary Gets DiscoveredItem
// @Accept json
// @Produce json
// @Tags DiscoveredItem
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DiscoveredItem
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /discovered_item/{id} [get]
func (e *DiscoveredItemController) getDiscoveredItem(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ItemId]"})
	}
	params = append(params, itemId)
	keys = append(keys, "item_id = ?")

	// query builder
	var result models.DiscoveredItem
	query := e.db.QueryContext(models.DiscoveredItem{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ItemId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateDiscoveredItem godoc
// @Id updateDiscoveredItem
// @Summary Updates DiscoveredItem
// @Accept json
// @Produce json
// @Tags DiscoveredItem
// @Param id path int true "Id"
// @Param discovered_item body models.DiscoveredItem true "DiscoveredItem"
// @Success 200 {array} models.DiscoveredItem
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /discovered_item/{id} [patch]
func (e *DiscoveredItemController) updateDiscoveredItem(c echo.Context) error {
	request := new(models.DiscoveredItem)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ItemId]"})
	}
	params = append(params, itemId)
	keys = append(keys, "item_id = ?")

	// query builder
	var result models.DiscoveredItem
	query := e.db.QueryContext(models.DiscoveredItem{}, c)
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
		event := fmt.Sprintf("Updated [DiscoveredItem] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createDiscoveredItem godoc
// @Id createDiscoveredItem
// @Summary Creates DiscoveredItem
// @Accept json
// @Produce json
// @Param discovered_item body models.DiscoveredItem true "DiscoveredItem"
// @Tags DiscoveredItem
// @Success 200 {array} models.DiscoveredItem
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /discovered_item [put]
func (e *DiscoveredItemController) createDiscoveredItem(c echo.Context) error {
	discoveredItem := new(models.DiscoveredItem)
	if err := c.Bind(discoveredItem); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.DiscoveredItem{}, c).Model(&models.DiscoveredItem{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
        db = db.Omit(clause.Associations)
    }

	err := db.Create(&discoveredItem).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.DiscoveredItem{}, discoveredItem)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [DiscoveredItem] [%v] data [%v]", discoveredItem.ItemId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, discoveredItem)
}

// deleteDiscoveredItem godoc
// @Id deleteDiscoveredItem
// @Summary Deletes DiscoveredItem
// @Accept json
// @Produce json
// @Tags DiscoveredItem
// @Param id path int true "itemId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /discovered_item/{id} [delete]
func (e *DiscoveredItemController) deleteDiscoveredItem(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, itemId)
	keys = append(keys, "item_id = ?")

	// query builder
	var result models.DiscoveredItem
	query := e.db.QueryContext(models.DiscoveredItem{}, c)
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
		event := fmt.Sprintf("Deleted [DiscoveredItem] [%v] keys [%v]", result.ItemId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDiscoveredItemsBulk godoc
// @Id getDiscoveredItemsBulk
// @Summary Gets DiscoveredItems in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DiscoveredItem
// @Success 200 {array} models.DiscoveredItem
// @Failure 500 {string} string "Bad query request"
// @Router /discovered_items/bulk [post]
func (e *DiscoveredItemController) getDiscoveredItemsBulk(c echo.Context) error {
	var results []models.DiscoveredItem

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

	err := e.db.QueryContext(models.DiscoveredItem{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getDiscoveredItemsCount godoc
// @Id getDiscoveredItemsCount
// @Summary Counts DiscoveredItems
// @Accept json
// @Produce json
// @Tags DiscoveredItem
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DiscoveredItem
// @Failure 500 {string} string "Bad query request"
// @Router /discovered_items/count [get]
func (e *DiscoveredItemController) getDiscoveredItemsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.DiscoveredItem{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}