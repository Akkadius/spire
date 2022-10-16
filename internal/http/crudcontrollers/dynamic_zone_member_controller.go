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

type DynamicZoneMemberController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewDynamicZoneMemberController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *DynamicZoneMemberController {
	return &DynamicZoneMemberController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *DynamicZoneMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "dynamic_zone_member/:id", e.getDynamicZoneMember, nil),
		routes.RegisterRoute(http.MethodGet, "dynamic_zone_members", e.listDynamicZoneMembers, nil),
		routes.RegisterRoute(http.MethodPut, "dynamic_zone_member", e.createDynamicZoneMember, nil),
		routes.RegisterRoute(http.MethodDelete, "dynamic_zone_member/:id", e.deleteDynamicZoneMember, nil),
		routes.RegisterRoute(http.MethodPatch, "dynamic_zone_member/:id", e.updateDynamicZoneMember, nil),
		routes.RegisterRoute(http.MethodPost, "dynamic_zone_members/bulk", e.getDynamicZoneMembersBulk, nil),
	}
}

// listDynamicZoneMembers godoc
// @Id listDynamicZoneMembers
// @Summary Lists DynamicZoneMembers
// @Accept json
// @Produce json
// @Tags DynamicZoneMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZoneMember
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_members [get]
func (e *DynamicZoneMemberController) listDynamicZoneMembers(c echo.Context) error {
	var results []models.DynamicZoneMember
	err := e.db.QueryContext(models.DynamicZoneMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDynamicZoneMember godoc
// @Id getDynamicZoneMember
// @Summary Gets DynamicZoneMember
// @Accept json
// @Produce json
// @Tags DynamicZoneMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DynamicZoneMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_member/{id} [get]
func (e *DynamicZoneMemberController) getDynamicZoneMember(c echo.Context) error {
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
	var result models.DynamicZoneMember
	query := e.db.QueryContext(models.DynamicZoneMember{}, c)
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

// updateDynamicZoneMember godoc
// @Id updateDynamicZoneMember
// @Summary Updates DynamicZoneMember
// @Accept json
// @Produce json
// @Tags DynamicZoneMember
// @Param id path int true "Id"
// @Param dynamic_zone_member body models.DynamicZoneMember true "DynamicZoneMember"
// @Success 200 {array} models.DynamicZoneMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /dynamic_zone_member/{id} [patch]
func (e *DynamicZoneMemberController) updateDynamicZoneMember(c echo.Context) error {
	request := new(models.DynamicZoneMember)
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
	var result models.DynamicZoneMember
	query := e.db.QueryContext(models.DynamicZoneMember{}, c)
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
		event := fmt.Sprintf("Updated [DynamicZoneMember] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createDynamicZoneMember godoc
// @Id createDynamicZoneMember
// @Summary Creates DynamicZoneMember
// @Accept json
// @Produce json
// @Param dynamic_zone_member body models.DynamicZoneMember true "DynamicZoneMember"
// @Tags DynamicZoneMember
// @Success 200 {array} models.DynamicZoneMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /dynamic_zone_member [put]
func (e *DynamicZoneMemberController) createDynamicZoneMember(c echo.Context) error {
	dynamicZoneMember := new(models.DynamicZoneMember)
	if err := c.Bind(dynamicZoneMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.DynamicZoneMember{}, c).Model(&models.DynamicZoneMember{}).Create(&dynamicZoneMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.DynamicZoneMember{}, dynamicZoneMember)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [DynamicZoneMember] [%v] data [%v]", dynamicZoneMember.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, dynamicZoneMember)
}

// deleteDynamicZoneMember godoc
// @Id deleteDynamicZoneMember
// @Summary Deletes DynamicZoneMember
// @Accept json
// @Produce json
// @Tags DynamicZoneMember
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /dynamic_zone_member/{id} [delete]
func (e *DynamicZoneMemberController) deleteDynamicZoneMember(c echo.Context) error {
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
	var result models.DynamicZoneMember
	query := e.db.QueryContext(models.DynamicZoneMember{}, c)
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
		event := fmt.Sprintf("Deleted [DynamicZoneMember] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDynamicZoneMembersBulk godoc
// @Id getDynamicZoneMembersBulk
// @Summary Gets DynamicZoneMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DynamicZoneMember
// @Success 200 {array} models.DynamicZoneMember
// @Failure 500 {string} string "Bad query request"
// @Router /dynamic_zone_members/bulk [post]
func (e *DynamicZoneMemberController) getDynamicZoneMembersBulk(c echo.Context) error {
	var results []models.DynamicZoneMember

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

	err := e.db.QueryContext(models.DynamicZoneMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
