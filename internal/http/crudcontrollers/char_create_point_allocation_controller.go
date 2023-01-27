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

type CharCreatePointAllocationController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharCreatePointAllocationController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharCreatePointAllocationController {
	return &CharCreatePointAllocationController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharCreatePointAllocationController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "char_create_point_allocation/:id", e.getCharCreatePointAllocation, nil),
		routes.RegisterRoute(http.MethodGet, "char_create_point_allocations", e.listCharCreatePointAllocations, nil),
		routes.RegisterRoute(http.MethodGet, "char_create_point_allocations/count", e.getCharCreatePointAllocationsCount, nil),
		routes.RegisterRoute(http.MethodPut, "char_create_point_allocation", e.createCharCreatePointAllocation, nil),
		routes.RegisterRoute(http.MethodDelete, "char_create_point_allocation/:id", e.deleteCharCreatePointAllocation, nil),
		routes.RegisterRoute(http.MethodPatch, "char_create_point_allocation/:id", e.updateCharCreatePointAllocation, nil),
		routes.RegisterRoute(http.MethodPost, "char_create_point_allocations/bulk", e.getCharCreatePointAllocationsBulk, nil),
	}
}

// listCharCreatePointAllocations godoc
// @Id listCharCreatePointAllocations
// @Summary Lists CharCreatePointAllocations
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_point_allocations [get]
func (e *CharCreatePointAllocationController) listCharCreatePointAllocations(c echo.Context) error {
	var results []models.CharCreatePointAllocation
	err := e.db.QueryContext(models.CharCreatePointAllocation{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharCreatePointAllocation godoc
// @Id getCharCreatePointAllocation
// @Summary Gets CharCreatePointAllocation
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_point_allocation/{id} [get]
func (e *CharCreatePointAllocationController) getCharCreatePointAllocation(c echo.Context) error {
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
	var result models.CharCreatePointAllocation
	query := e.db.QueryContext(models.CharCreatePointAllocation{}, c)
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

// updateCharCreatePointAllocation godoc
// @Id updateCharCreatePointAllocation
// @Summary Updates CharCreatePointAllocation
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param id path int true "Id"
// @Param char_create_point_allocation body models.CharCreatePointAllocation true "CharCreatePointAllocation"
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /char_create_point_allocation/{id} [patch]
func (e *CharCreatePointAllocationController) updateCharCreatePointAllocation(c echo.Context) error {
	request := new(models.CharCreatePointAllocation)
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
	var result models.CharCreatePointAllocation
	query := e.db.QueryContext(models.CharCreatePointAllocation{}, c)
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
		event := fmt.Sprintf("Updated [CharCreatePointAllocation] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharCreatePointAllocation godoc
// @Id createCharCreatePointAllocation
// @Summary Creates CharCreatePointAllocation
// @Accept json
// @Produce json
// @Param char_create_point_allocation body models.CharCreatePointAllocation true "CharCreatePointAllocation"
// @Tags CharCreatePointAllocation
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /char_create_point_allocation [put]
func (e *CharCreatePointAllocationController) createCharCreatePointAllocation(c echo.Context) error {
	charCreatePointAllocation := new(models.CharCreatePointAllocation)
	if err := c.Bind(charCreatePointAllocation); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharCreatePointAllocation{}, c).Model(&models.CharCreatePointAllocation{}).Create(&charCreatePointAllocation).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharCreatePointAllocation{}, charCreatePointAllocation)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharCreatePointAllocation] [%v] data [%v]", charCreatePointAllocation.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, charCreatePointAllocation)
}

// deleteCharCreatePointAllocation godoc
// @Id deleteCharCreatePointAllocation
// @Summary Deletes CharCreatePointAllocation
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /char_create_point_allocation/{id} [delete]
func (e *CharCreatePointAllocationController) deleteCharCreatePointAllocation(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.CharCreatePointAllocation
	query := e.db.QueryContext(models.CharCreatePointAllocation{}, c)
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
		event := fmt.Sprintf("Deleted [CharCreatePointAllocation] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharCreatePointAllocationsBulk godoc
// @Id getCharCreatePointAllocationsBulk
// @Summary Gets CharCreatePointAllocations in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharCreatePointAllocation
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_point_allocations/bulk [post]
func (e *CharCreatePointAllocationController) getCharCreatePointAllocationsBulk(c echo.Context) error {
	var results []models.CharCreatePointAllocation

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

	err := e.db.QueryContext(models.CharCreatePointAllocation{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharCreatePointAllocationsCount godoc
// @Id getCharCreatePointAllocationsCount
// @Summary Counts CharCreatePointAllocations
// @Accept json
// @Produce json
// @Tags CharCreatePointAllocation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharCreatePointAllocation
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_point_allocations/count [get]
func (e *CharCreatePointAllocationController) getCharCreatePointAllocationsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharCreatePointAllocation{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}