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

type GroupIdController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewGroupIdController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *GroupIdController {
	return &GroupIdController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *GroupIdController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "group_id/:groupid", e.getGroupId, nil),
		routes.RegisterRoute(http.MethodGet, "group_ids", e.listGroupIds, nil),
		routes.RegisterRoute(http.MethodGet, "group_ids/count", e.getGroupIdsCount, nil),
		routes.RegisterRoute(http.MethodPut, "group_id", e.createGroupId, nil),
		routes.RegisterRoute(http.MethodDelete, "group_id/:groupid", e.deleteGroupId, nil),
		routes.RegisterRoute(http.MethodPatch, "group_id/:groupid", e.updateGroupId, nil),
		routes.RegisterRoute(http.MethodPost, "group_ids/bulk", e.getGroupIdsBulk, nil),
	}
}

// listGroupIds godoc
// @Id listGroupIds
// @Summary Lists GroupIds
// @Accept json
// @Produce json
// @Tags GroupId
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroupId
// @Failure 500 {string} string "Bad query request"
// @Router /group_ids [get]
func (e *GroupIdController) listGroupIds(c echo.Context) error {
	var results []models.GroupId
	err := e.db.QueryContext(models.GroupId{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGroupId godoc
// @Id getGroupId
// @Summary Gets GroupId
// @Accept json
// @Produce json
// @Tags GroupId
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroupId
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /group_id/{id} [get]
func (e *GroupIdController) getGroupId(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	groupid, err := strconv.Atoi(c.Param("groupid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Groupid]"})
	}
	params = append(params, groupid)
	keys = append(keys, "groupid = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [ismerc] position [4] type [tinyint]
	if len(c.QueryParam("ismerc")) > 0 {
		ismercParam, err := strconv.Atoi(c.QueryParam("ismerc"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ismerc] err [%s]", err.Error())})
		}

		params = append(params, ismercParam)
		keys = append(keys, "ismerc = ?")
	}

	// query builder
	var result models.GroupId
	query := e.db.QueryContext(models.GroupId{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Groupid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGroupId godoc
// @Id updateGroupId
// @Summary Updates GroupId
// @Accept json
// @Produce json
// @Tags GroupId
// @Param id path int true "Id"
// @Param group_id body models.GroupId true "GroupId"
// @Success 200 {array} models.GroupId
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /group_id/{id} [patch]
func (e *GroupIdController) updateGroupId(c echo.Context) error {
	request := new(models.GroupId)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	groupid, err := strconv.Atoi(c.Param("groupid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Groupid]"})
	}
	params = append(params, groupid)
	keys = append(keys, "groupid = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [ismerc] position [4] type [tinyint]
	if len(c.QueryParam("ismerc")) > 0 {
		ismercParam, err := strconv.Atoi(c.QueryParam("ismerc"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ismerc] err [%s]", err.Error())})
		}

		params = append(params, ismercParam)
		keys = append(keys, "ismerc = ?")
	}

	// query builder
	var result models.GroupId
	query := e.db.QueryContext(models.GroupId{}, c)
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
		event := fmt.Sprintf("Updated [GroupId] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createGroupId godoc
// @Id createGroupId
// @Summary Creates GroupId
// @Accept json
// @Produce json
// @Param group_id body models.GroupId true "GroupId"
// @Tags GroupId
// @Success 200 {array} models.GroupId
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /group_id [put]
func (e *GroupIdController) createGroupId(c echo.Context) error {
	groupId := new(models.GroupId)
	if err := c.Bind(groupId); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.GroupId{}, c).Model(&models.GroupId{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&groupId).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GroupId{}, groupId)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GroupId] [%v] data [%v]", groupId.Groupid, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, groupId)
}

// deleteGroupId godoc
// @Id deleteGroupId
// @Summary Deletes GroupId
// @Accept json
// @Produce json
// @Tags GroupId
// @Param id path int true "groupid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /group_id/{id} [delete]
func (e *GroupIdController) deleteGroupId(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	groupid, err := strconv.Atoi(c.Param("groupid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, groupid)
	keys = append(keys, "groupid = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// key param [ismerc] position [4] type [tinyint]
	if len(c.QueryParam("ismerc")) > 0 {
		ismercParam, err := strconv.Atoi(c.QueryParam("ismerc"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [ismerc] err [%s]", err.Error())})
		}

		params = append(params, ismercParam)
		keys = append(keys, "ismerc = ?")
	}

	// query builder
	var result models.GroupId
	query := e.db.QueryContext(models.GroupId{}, c)
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
		event := fmt.Sprintf("Deleted [GroupId] [%v] keys [%v]", result.Groupid, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGroupIdsBulk godoc
// @Id getGroupIdsBulk
// @Summary Gets GroupIds in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GroupId
// @Success 200 {array} models.GroupId
// @Failure 500 {string} string "Bad query request"
// @Router /group_ids/bulk [post]
func (e *GroupIdController) getGroupIdsBulk(c echo.Context) error {
	var results []models.GroupId

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

	err := e.db.QueryContext(models.GroupId{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGroupIdsCount godoc
// @Id getGroupIdsCount
// @Summary Counts GroupIds
// @Accept json
// @Produce json
// @Tags GroupId
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroupId
// @Failure 500 {string} string "Bad query request"
// @Router /group_ids/count [get]
func (e *GroupIdController) getGroupIdsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.GroupId{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}