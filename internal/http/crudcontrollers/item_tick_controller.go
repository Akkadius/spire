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

type ItemTickController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewItemTickController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *ItemTickController {
	return &ItemTickController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *ItemTickController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "item_tick/:itId", e.getItemTick, nil),
		routes.RegisterRoute(http.MethodGet, "item_ticks", e.listItemTicks, nil),
		routes.RegisterRoute(http.MethodGet, "item_ticks/count", e.getItemTicksCount, nil),
		routes.RegisterRoute(http.MethodPut, "item_tick", e.createItemTick, nil),
		routes.RegisterRoute(http.MethodDelete, "item_tick/:itId", e.deleteItemTick, nil),
		routes.RegisterRoute(http.MethodPatch, "item_tick/:itId", e.updateItemTick, nil),
		routes.RegisterRoute(http.MethodPost, "item_ticks/bulk", e.getItemTicksBulk, nil),
	}
}

// listItemTicks godoc
// @Id listItemTicks
// @Summary Lists ItemTicks
// @Accept json
// @Produce json
// @Tags ItemTick
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ItemTick
// @Failure 500 {string} string "Bad query request"
// @Router /item_ticks [get]
func (e *ItemTickController) listItemTicks(c echo.Context) error {
	var results []models.ItemTick
	err := e.db.QueryContext(models.ItemTick{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getItemTick godoc
// @Id getItemTick
// @Summary Gets ItemTick
// @Accept json
// @Produce json
// @Tags ItemTick
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ItemTick
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /item_tick/{id} [get]
func (e *ItemTickController) getItemTick(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	itId, err := strconv.Atoi(c.Param("itId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ItId]"})
	}
	params = append(params, itId)
	keys = append(keys, "it_id = ?")

	// query builder
	var result models.ItemTick
	query := e.db.QueryContext(models.ItemTick{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ItId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateItemTick godoc
// @Id updateItemTick
// @Summary Updates ItemTick
// @Accept json
// @Produce json
// @Tags ItemTick
// @Param id path int true "Id"
// @Param item_tick body models.ItemTick true "ItemTick"
// @Success 200 {array} models.ItemTick
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /item_tick/{id} [patch]
func (e *ItemTickController) updateItemTick(c echo.Context) error {
	request := new(models.ItemTick)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	itId, err := strconv.Atoi(c.Param("itId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ItId]"})
	}
	params = append(params, itId)
	keys = append(keys, "it_id = ?")

	// query builder
	var result models.ItemTick
	query := e.db.QueryContext(models.ItemTick{}, c)
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
		event := fmt.Sprintf("Updated [ItemTick] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createItemTick godoc
// @Id createItemTick
// @Summary Creates ItemTick
// @Accept json
// @Produce json
// @Param item_tick body models.ItemTick true "ItemTick"
// @Tags ItemTick
// @Success 200 {array} models.ItemTick
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /item_tick [put]
func (e *ItemTickController) createItemTick(c echo.Context) error {
	itemTick := new(models.ItemTick)
	if err := c.Bind(itemTick); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.ItemTick{}, c).Model(&models.ItemTick{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&itemTick).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.ItemTick{}, itemTick)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [ItemTick] [%v] data [%v]", itemTick.ItId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, itemTick)
}

// deleteItemTick godoc
// @Id deleteItemTick
// @Summary Deletes ItemTick
// @Accept json
// @Produce json
// @Tags ItemTick
// @Param id path int true "itId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /item_tick/{id} [delete]
func (e *ItemTickController) deleteItemTick(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	itId, err := strconv.Atoi(c.Param("itId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, itId)
	keys = append(keys, "it_id = ?")

	// query builder
	var result models.ItemTick
	query := e.db.QueryContext(models.ItemTick{}, c)
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
		event := fmt.Sprintf("Deleted [ItemTick] [%v] keys [%v]", result.ItId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getItemTicksBulk godoc
// @Id getItemTicksBulk
// @Summary Gets ItemTicks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ItemTick
// @Success 200 {array} models.ItemTick
// @Failure 500 {string} string "Bad query request"
// @Router /item_ticks/bulk [post]
func (e *ItemTickController) getItemTicksBulk(c echo.Context) error {
	var results []models.ItemTick

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

	err := e.db.QueryContext(models.ItemTick{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getItemTicksCount godoc
// @Id getItemTicksCount
// @Summary Counts ItemTicks
// @Accept json
// @Produce json
// @Tags ItemTick
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ItemTick
// @Failure 500 {string} string "Bad query request"
// @Router /item_ticks/count [get]
func (e *ItemTickController) getItemTicksCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.ItemTick{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}