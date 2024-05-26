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

type GuildTributeController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewGuildTributeController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *GuildTributeController {
	return &GuildTributeController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *GuildTributeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "guild_tribute/:guildId", e.getGuildTribute, nil),
		routes.RegisterRoute(http.MethodGet, "guild_tributes", e.listGuildTributes, nil),
		routes.RegisterRoute(http.MethodGet, "guild_tributes/count", e.getGuildTributesCount, nil),
		routes.RegisterRoute(http.MethodPut, "guild_tribute", e.createGuildTribute, nil),
		routes.RegisterRoute(http.MethodDelete, "guild_tribute/:guildId", e.deleteGuildTribute, nil),
		routes.RegisterRoute(http.MethodPatch, "guild_tribute/:guildId", e.updateGuildTribute, nil),
		routes.RegisterRoute(http.MethodPost, "guild_tributes/bulk", e.getGuildTributesBulk, nil),
	}
}

// listGuildTributes godoc
// @Id listGuildTributes
// @Summary Lists GuildTributes
// @Accept json
// @Produce json
// @Tags GuildTribute
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildTribute
// @Failure 500 {string} string "Bad query request"
// @Router /guild_tributes [get]
func (e *GuildTributeController) listGuildTributes(c echo.Context) error {
	var results []models.GuildTribute
	err := e.db.QueryContext(models.GuildTribute{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildTribute godoc
// @Id getGuildTribute
// @Summary Gets GuildTribute
// @Accept json
// @Produce json
// @Tags GuildTribute
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildTribute
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /guild_tribute/{id} [get]
func (e *GuildTributeController) getGuildTribute(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	guildId, err := strconv.Atoi(c.Param("guildId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GuildId]"})
	}
	params = append(params, guildId)
	keys = append(keys, "guild_id = ?")

	// query builder
	var result models.GuildTribute
	query := e.db.QueryContext(models.GuildTribute{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.GuildId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGuildTribute godoc
// @Id updateGuildTribute
// @Summary Updates GuildTribute
// @Accept json
// @Produce json
// @Tags GuildTribute
// @Param id path int true "Id"
// @Param guild_tribute body models.GuildTribute true "GuildTribute"
// @Success 200 {array} models.GuildTribute
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /guild_tribute/{id} [patch]
func (e *GuildTributeController) updateGuildTribute(c echo.Context) error {
	request := new(models.GuildTribute)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	guildId, err := strconv.Atoi(c.Param("guildId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GuildId]"})
	}
	params = append(params, guildId)
	keys = append(keys, "guild_id = ?")

	// query builder
	var result models.GuildTribute
	query := e.db.QueryContext(models.GuildTribute{}, c)
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
		event := fmt.Sprintf("Updated [GuildTribute] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createGuildTribute godoc
// @Id createGuildTribute
// @Summary Creates GuildTribute
// @Accept json
// @Produce json
// @Param guild_tribute body models.GuildTribute true "GuildTribute"
// @Tags GuildTribute
// @Success 200 {array} models.GuildTribute
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /guild_tribute [put]
func (e *GuildTributeController) createGuildTribute(c echo.Context) error {
	guildTribute := new(models.GuildTribute)
	if err := c.Bind(guildTribute); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.GuildTribute{}, c).Model(&models.GuildTribute{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&guildTribute).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GuildTribute{}, guildTribute)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GuildTribute] [%v] data [%v]", guildTribute.GuildId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, guildTribute)
}

// deleteGuildTribute godoc
// @Id deleteGuildTribute
// @Summary Deletes GuildTribute
// @Accept json
// @Produce json
// @Tags GuildTribute
// @Param id path int true "guildId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /guild_tribute/{id} [delete]
func (e *GuildTributeController) deleteGuildTribute(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	guildId, err := strconv.Atoi(c.Param("guildId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, guildId)
	keys = append(keys, "guild_id = ?")

	// query builder
	var result models.GuildTribute
	query := e.db.QueryContext(models.GuildTribute{}, c)
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
		event := fmt.Sprintf("Deleted [GuildTribute] [%v] keys [%v]", result.GuildId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGuildTributesBulk godoc
// @Id getGuildTributesBulk
// @Summary Gets GuildTributes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GuildTribute
// @Success 200 {array} models.GuildTribute
// @Failure 500 {string} string "Bad query request"
// @Router /guild_tributes/bulk [post]
func (e *GuildTributeController) getGuildTributesBulk(c echo.Context) error {
	var results []models.GuildTribute

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

	err := e.db.QueryContext(models.GuildTribute{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildTributesCount godoc
// @Id getGuildTributesCount
// @Summary Counts GuildTributes
// @Accept json
// @Produce json
// @Tags GuildTribute
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildTribute
// @Failure 500 {string} string "Bad query request"
// @Router /guild_tributes/count [get]
func (e *GuildTributeController) getGuildTributesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.GuildTribute{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}