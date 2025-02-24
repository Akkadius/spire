package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type ItemsEvolvingDetailController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewItemsEvolvingDetailController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *ItemsEvolvingDetailController {
	return &ItemsEvolvingDetailController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *ItemsEvolvingDetailController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "items_evolving_detail/:id", e.getItemsEvolvingDetail, nil),
		routes.RegisterRoute(http.MethodGet, "items_evolving_details", e.listItemsEvolvingDetails, nil),
		routes.RegisterRoute(http.MethodGet, "items_evolving_details/count", e.getItemsEvolvingDetailsCount, nil),
		routes.RegisterRoute(http.MethodPut, "items_evolving_detail", e.createItemsEvolvingDetail, nil),
		routes.RegisterRoute(http.MethodDelete, "items_evolving_detail/:id", e.deleteItemsEvolvingDetail, nil),
		routes.RegisterRoute(http.MethodPatch, "items_evolving_detail/:id", e.updateItemsEvolvingDetail, nil),
		routes.RegisterRoute(http.MethodPost, "items_evolving_details/bulk", e.getItemsEvolvingDetailsBulk, nil),
	}
}

// listItemsEvolvingDetails godoc
// @Id listItemsEvolvingDetails
// @Summary Lists ItemsEvolvingDetails
// @Accept json
// @Produce json
// @Tags ItemsEvolvingDetail
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ItemsEvolvingDetail
// @Failure 500 {string} string "Bad query request"
// @Router /items_evolving_details [get]
func (e *ItemsEvolvingDetailController) listItemsEvolvingDetails(c echo.Context) error {
	var results []models.ItemsEvolvingDetail
	err := e.db.QueryContext(models.ItemsEvolvingDetail{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getItemsEvolvingDetail godoc
// @Id getItemsEvolvingDetail
// @Summary Gets ItemsEvolvingDetail
// @Accept json
// @Produce json
// @Tags ItemsEvolvingDetail
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ItemsEvolvingDetail
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /items_evolving_detail/{id} [get]
func (e *ItemsEvolvingDetailController) getItemsEvolvingDetail(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.ItemsEvolvingDetail
	query := e.db.QueryContext(models.ItemsEvolvingDetail{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateItemsEvolvingDetail godoc
// @Id updateItemsEvolvingDetail
// @Summary Updates ItemsEvolvingDetail
// @Accept json
// @Produce json
// @Tags ItemsEvolvingDetail
// @Param id path int true "Id"
// @Param items_evolving_detail body models.ItemsEvolvingDetail true "ItemsEvolvingDetail"
// @Success 200 {array} models.ItemsEvolvingDetail
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /items_evolving_detail/{id} [patch]
func (e *ItemsEvolvingDetailController) updateItemsEvolvingDetail(c echo.Context) error {
	request := new(models.ItemsEvolvingDetail)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.ItemsEvolvingDetail
	query := e.db.QueryContext(models.ItemsEvolvingDetail{}, c)
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
		event := fmt.Sprintf("Updated [ItemsEvolvingDetail] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createItemsEvolvingDetail godoc
// @Id createItemsEvolvingDetail
// @Summary Creates ItemsEvolvingDetail
// @Accept json
// @Produce json
// @Param items_evolving_detail body models.ItemsEvolvingDetail true "ItemsEvolvingDetail"
// @Tags ItemsEvolvingDetail
// @Success 200 {array} models.ItemsEvolvingDetail
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /items_evolving_detail [put]
func (e *ItemsEvolvingDetailController) createItemsEvolvingDetail(c echo.Context) error {
	itemsEvolvingDetail := new(models.ItemsEvolvingDetail)
	if err := c.Bind(itemsEvolvingDetail); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.ItemsEvolvingDetail{}, c).Model(&models.ItemsEvolvingDetail{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&itemsEvolvingDetail).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.ItemsEvolvingDetail{}, itemsEvolvingDetail)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [ItemsEvolvingDetail] [%v] data [%v]", itemsEvolvingDetail.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, itemsEvolvingDetail)
}

// deleteItemsEvolvingDetail godoc
// @Id deleteItemsEvolvingDetail
// @Summary Deletes ItemsEvolvingDetail
// @Accept json
// @Produce json
// @Tags ItemsEvolvingDetail
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /items_evolving_detail/{id} [delete]
func (e *ItemsEvolvingDetailController) deleteItemsEvolvingDetail(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.ItemsEvolvingDetail
	query := e.db.QueryContext(models.ItemsEvolvingDetail{}, c)
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
		event := fmt.Sprintf("Deleted [ItemsEvolvingDetail] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getItemsEvolvingDetailsBulk godoc
// @Id getItemsEvolvingDetailsBulk
// @Summary Gets ItemsEvolvingDetails in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ItemsEvolvingDetail
// @Success 200 {array} models.ItemsEvolvingDetail
// @Failure 500 {string} string "Bad query request"
// @Router /items_evolving_details/bulk [post]
func (e *ItemsEvolvingDetailController) getItemsEvolvingDetailsBulk(c echo.Context) error {
	var results []models.ItemsEvolvingDetail

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

	err := e.db.QueryContext(models.ItemsEvolvingDetail{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getItemsEvolvingDetailsCount godoc
// @Id getItemsEvolvingDetailsCount
// @Summary Counts ItemsEvolvingDetails
// @Accept json
// @Produce json
// @Tags ItemsEvolvingDetail
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ItemsEvolvingDetail
// @Failure 500 {string} string "Bad query request"
// @Router /items_evolving_details/count [get]
func (e *ItemsEvolvingDetailController) getItemsEvolvingDetailsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.ItemsEvolvingDetail{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}