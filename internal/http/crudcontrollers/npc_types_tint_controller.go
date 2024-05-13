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

type NpcTypesTintController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewNpcTypesTintController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *NpcTypesTintController {
	return &NpcTypesTintController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *NpcTypesTintController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_types_tint/:id", e.getNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodGet, "npc_types_tints", e.listNpcTypesTints, nil),
		routes.RegisterRoute(http.MethodGet, "npc_types_tints/count", e.getNpcTypesTintsCount, nil),
		routes.RegisterRoute(http.MethodPut, "npc_types_tint", e.createNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_types_tint/:id", e.deleteNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_types_tint/:id", e.updateNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodPost, "npc_types_tints/bulk", e.getNpcTypesTintsBulk, nil),
	}
}

// listNpcTypesTints godoc
// @Id listNpcTypesTints
// @Summary Lists NpcTypesTints
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcTypesTint
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types_tints [get]
func (e *NpcTypesTintController) listNpcTypesTints(c echo.Context) error {
	var results []models.NpcTypesTint
	err := e.db.QueryContext(models.NpcTypesTint{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcTypesTint godoc
// @Id getNpcTypesTint
// @Summary Gets NpcTypesTint
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcTypesTint
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types_tint/{id} [get]
func (e *NpcTypesTintController) getNpcTypesTint(c echo.Context) error {
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
	var result models.NpcTypesTint
	query := e.db.QueryContext(models.NpcTypesTint{}, c)
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

// updateNpcTypesTint godoc
// @Id updateNpcTypesTint
// @Summary Updates NpcTypesTint
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param id path int true "Id"
// @Param npc_types_tint body models.NpcTypesTint true "NpcTypesTint"
// @Success 200 {array} models.NpcTypesTint
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_types_tint/{id} [patch]
func (e *NpcTypesTintController) updateNpcTypesTint(c echo.Context) error {
	request := new(models.NpcTypesTint)
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
	var result models.NpcTypesTint
	query := e.db.QueryContext(models.NpcTypesTint{}, c)
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
		event := fmt.Sprintf("Updated [NpcTypesTint] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcTypesTint godoc
// @Id createNpcTypesTint
// @Summary Creates NpcTypesTint
// @Accept json
// @Produce json
// @Param npc_types_tint body models.NpcTypesTint true "NpcTypesTint"
// @Tags NpcTypesTint
// @Success 200 {array} models.NpcTypesTint
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_types_tint [put]
func (e *NpcTypesTintController) createNpcTypesTint(c echo.Context) error {
	npcTypesTint := new(models.NpcTypesTint)
	if err := c.Bind(npcTypesTint); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.NpcTypesTint{}, c).Model(&models.NpcTypesTint{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&npcTypesTint).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.NpcTypesTint{}, npcTypesTint)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [NpcTypesTint] [%v] data [%v]", npcTypesTint.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, npcTypesTint)
}

// deleteNpcTypesTint godoc
// @Id deleteNpcTypesTint
// @Summary Deletes NpcTypesTint
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_types_tint/{id} [delete]
func (e *NpcTypesTintController) deleteNpcTypesTint(c echo.Context) error {
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
	var result models.NpcTypesTint
	query := e.db.QueryContext(models.NpcTypesTint{}, c)
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
		event := fmt.Sprintf("Deleted [NpcTypesTint] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcTypesTintsBulk godoc
// @Id getNpcTypesTintsBulk
// @Summary Gets NpcTypesTints in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcTypesTint
// @Success 200 {array} models.NpcTypesTint
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types_tints/bulk [post]
func (e *NpcTypesTintController) getNpcTypesTintsBulk(c echo.Context) error {
	var results []models.NpcTypesTint

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

	err := e.db.QueryContext(models.NpcTypesTint{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcTypesTintsCount godoc
// @Id getNpcTypesTintsCount
// @Summary Counts NpcTypesTints
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcTypesTint
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types_tints/count [get]
func (e *NpcTypesTintController) getNpcTypesTintsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.NpcTypesTint{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}