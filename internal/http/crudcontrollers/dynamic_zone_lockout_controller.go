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

type DynamicZoneLockoutController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewDynamicZoneLockoutController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *DynamicZoneLockoutController {
	return &DynamicZoneLockoutController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *DynamicZoneLockoutController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "dynamic_zone_lockout/:id", e.getDynamicZoneLockout, nil),
		routes.RegisterRoute(http.MethodGet, "dynamic_zone_lockouts", e.listDynamicZoneLockouts, nil),
		routes.RegisterRoute(http.MethodGet, "dynamic_zone_lockouts/count", e.getDynamicZoneLockoutsCount, nil),
		routes.RegisterRoute(http.MethodPut, "dynamic_zone_lockout", e.createDynamicZoneLockout, nil),
		routes.RegisterRoute(http.MethodDelete, "dynamic_zone_lockout/:id", e.deleteDynamicZoneLockout, nil),
		routes.RegisterRoute(http.MethodPatch, "dynamic_zone_lockout/:id", e.updateDynamicZoneLockout, nil),
		routes.RegisterRoute(http.MethodPost, "dynamic_zone_lockouts/bulk", e.getDynamicZoneLockoutsBulk, nil),
	}
}

// listDynamicZoneLockouts godoc
// @Id listDynamicZoneLockouts
// @Summary Lists DynamicZoneLockouts
// @Accept json
// @Produce json
// @Tags DynamicZoneLockout
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZoneLockout
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_lockouts [get]
func (e *DynamicZoneLockoutController) listDynamicZoneLockouts(c echo.Context) error {
	var results []models.DynamicZoneLockout
	err := e.db.QueryContext(models.DynamicZoneLockout{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getDynamicZoneLockout godoc
// @Id getDynamicZoneLockout
// @Summary Gets DynamicZoneLockout
// @Accept json
// @Produce json
// @Tags DynamicZoneLockout
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZoneLockout
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_lockout/{id} [get]
func (e *DynamicZoneLockoutController) getDynamicZoneLockout(c echo.Context) error {
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
	var result models.DynamicZoneLockout
	query := e.db.QueryContext(models.DynamicZoneLockout{}, c)
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

// updateDynamicZoneLockout godoc
// @Id updateDynamicZoneLockout
// @Summary Updates DynamicZoneLockout
// @Accept json
// @Produce json
// @Tags DynamicZoneLockout
// @Param id path int true "Id"
// @Param dynamic_zone_lockout body models.DynamicZoneLockout true "DynamicZoneLockout"
// @Success 200 {array} models.DynamicZoneLockout
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /dynamic_zone_lockout/{id} [patch]
func (e *DynamicZoneLockoutController) updateDynamicZoneLockout(c echo.Context) error {
	request := new(models.DynamicZoneLockout)
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
	var result models.DynamicZoneLockout
	query := e.db.QueryContext(models.DynamicZoneLockout{}, c)
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
		event := fmt.Sprintf("Updated [DynamicZoneLockout] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createDynamicZoneLockout godoc
// @Id createDynamicZoneLockout
// @Summary Creates DynamicZoneLockout
// @Accept json
// @Produce json
// @Param dynamic_zone_lockout body models.DynamicZoneLockout true "DynamicZoneLockout"
// @Tags DynamicZoneLockout
// @Success 200 {array} models.DynamicZoneLockout
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /dynamic_zone_lockout [put]
func (e *DynamicZoneLockoutController) createDynamicZoneLockout(c echo.Context) error {
	dynamicZoneLockout := new(models.DynamicZoneLockout)
	if err := c.Bind(dynamicZoneLockout); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.DynamicZoneLockout{}, c).Model(&models.DynamicZoneLockout{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&dynamicZoneLockout).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.DynamicZoneLockout{}, dynamicZoneLockout)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [DynamicZoneLockout] [%v] data [%v]", dynamicZoneLockout.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, dynamicZoneLockout)
}

// deleteDynamicZoneLockout godoc
// @Id deleteDynamicZoneLockout
// @Summary Deletes DynamicZoneLockout
// @Accept json
// @Produce json
// @Tags DynamicZoneLockout
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /dynamic_zone_lockout/{id} [delete]
func (e *DynamicZoneLockoutController) deleteDynamicZoneLockout(c echo.Context) error {
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
	var result models.DynamicZoneLockout
	query := e.db.QueryContext(models.DynamicZoneLockout{}, c)
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
		event := fmt.Sprintf("Deleted [DynamicZoneLockout] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDynamicZoneLockoutsBulk godoc
// @Id getDynamicZoneLockoutsBulk
// @Summary Gets DynamicZoneLockouts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DynamicZoneLockout
// @Success 200 {array} models.DynamicZoneLockout
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_lockouts/bulk [post]
func (e *DynamicZoneLockoutController) getDynamicZoneLockoutsBulk(c echo.Context) error {
	var results []models.DynamicZoneLockout

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

	err := e.db.QueryContext(models.DynamicZoneLockout{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getDynamicZoneLockoutsCount godoc
// @Id getDynamicZoneLockoutsCount
// @Summary Counts DynamicZoneLockouts
// @Accept json
// @Produce json
// @Tags DynamicZoneLockout
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZoneLockout
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_lockouts/count [get]
func (e *DynamicZoneLockoutController) getDynamicZoneLockoutsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.DynamicZoneLockout{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}