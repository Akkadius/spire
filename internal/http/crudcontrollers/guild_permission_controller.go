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

type GuildPermissionController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewGuildPermissionController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *GuildPermissionController {
	return &GuildPermissionController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *GuildPermissionController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "guild_permission/:id", e.getGuildPermission, nil),
		routes.RegisterRoute(http.MethodGet, "guild_permissions", e.listGuildPermissions, nil),
		routes.RegisterRoute(http.MethodGet, "guild_permissions/count", e.getGuildPermissionsCount, nil),
		routes.RegisterRoute(http.MethodPut, "guild_permission", e.createGuildPermission, nil),
		routes.RegisterRoute(http.MethodDelete, "guild_permission/:id", e.deleteGuildPermission, nil),
		routes.RegisterRoute(http.MethodPatch, "guild_permission/:id", e.updateGuildPermission, nil),
		routes.RegisterRoute(http.MethodPost, "guild_permissions/bulk", e.getGuildPermissionsBulk, nil),
	}
}

// listGuildPermissions godoc
// @Id listGuildPermissions
// @Summary Lists GuildPermissions
// @Accept json
// @Produce json
// @Tags GuildPermission
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildPermission
// @Failure 500 {string} string "Bad query request"
// @Router /guild_permissions [get]
func (e *GuildPermissionController) listGuildPermissions(c echo.Context) error {
	var results []models.GuildPermission
	err := e.db.QueryContext(models.GuildPermission{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildPermission godoc
// @Id getGuildPermission
// @Summary Gets GuildPermission
// @Accept json
// @Produce json
// @Tags GuildPermission
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildPermission
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /guild_permission/{id} [get]
func (e *GuildPermissionController) getGuildPermission(c echo.Context) error {
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
	var result models.GuildPermission
	query := e.db.QueryContext(models.GuildPermission{}, c)
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

// updateGuildPermission godoc
// @Id updateGuildPermission
// @Summary Updates GuildPermission
// @Accept json
// @Produce json
// @Tags GuildPermission
// @Param id path int true "Id"
// @Param guild_permission body models.GuildPermission true "GuildPermission"
// @Success 200 {array} models.GuildPermission
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /guild_permission/{id} [patch]
func (e *GuildPermissionController) updateGuildPermission(c echo.Context) error {
	request := new(models.GuildPermission)
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
	var result models.GuildPermission
	query := e.db.QueryContext(models.GuildPermission{}, c)
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
		event := fmt.Sprintf("Updated [GuildPermission] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createGuildPermission godoc
// @Id createGuildPermission
// @Summary Creates GuildPermission
// @Accept json
// @Produce json
// @Param guild_permission body models.GuildPermission true "GuildPermission"
// @Tags GuildPermission
// @Success 200 {array} models.GuildPermission
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /guild_permission [put]
func (e *GuildPermissionController) createGuildPermission(c echo.Context) error {
	guildPermission := new(models.GuildPermission)
	if err := c.Bind(guildPermission); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.GuildPermission{}, c).Model(&models.GuildPermission{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&guildPermission).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GuildPermission{}, guildPermission)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GuildPermission] [%v] data [%v]", guildPermission.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, guildPermission)
}

// deleteGuildPermission godoc
// @Id deleteGuildPermission
// @Summary Deletes GuildPermission
// @Accept json
// @Produce json
// @Tags GuildPermission
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /guild_permission/{id} [delete]
func (e *GuildPermissionController) deleteGuildPermission(c echo.Context) error {
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
	var result models.GuildPermission
	query := e.db.QueryContext(models.GuildPermission{}, c)
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
		event := fmt.Sprintf("Deleted [GuildPermission] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGuildPermissionsBulk godoc
// @Id getGuildPermissionsBulk
// @Summary Gets GuildPermissions in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GuildPermission
// @Success 200 {array} models.GuildPermission
// @Failure 500 {string} string "Bad query request"
// @Router /guild_permissions/bulk [post]
func (e *GuildPermissionController) getGuildPermissionsBulk(c echo.Context) error {
	var results []models.GuildPermission

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

	err := e.db.QueryContext(models.GuildPermission{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildPermissionsCount godoc
// @Id getGuildPermissionsCount
// @Summary Counts GuildPermissions
// @Accept json
// @Produce json
// @Tags GuildPermission
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildPermission
// @Failure 500 {string} string "Bad query request"
// @Router /guild_permissions/count [get]
func (e *GuildPermissionController) getGuildPermissionsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.GuildPermission{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}