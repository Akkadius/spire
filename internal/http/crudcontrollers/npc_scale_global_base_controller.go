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

type NpcScaleGlobalBaseController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewNpcScaleGlobalBaseController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *NpcScaleGlobalBaseController {
	return &NpcScaleGlobalBaseController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *NpcScaleGlobalBaseController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_scale_global_base/:typeId", e.getNpcScaleGlobalBase, nil),
		routes.RegisterRoute(http.MethodGet, "npc_scale_global_bases", e.listNpcScaleGlobalBases, nil),
		routes.RegisterRoute(http.MethodGet, "npc_scale_global_bases/count", e.getNpcScaleGlobalBasesCount, nil),
		routes.RegisterRoute(http.MethodPut, "npc_scale_global_base", e.createNpcScaleGlobalBase, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_scale_global_base/:typeId", e.deleteNpcScaleGlobalBase, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_scale_global_base/:typeId", e.updateNpcScaleGlobalBase, nil),
		routes.RegisterRoute(http.MethodPost, "npc_scale_global_bases/bulk", e.getNpcScaleGlobalBasesBulk, nil),
	}
}

// listNpcScaleGlobalBases godoc
// @Id listNpcScaleGlobalBases
// @Summary Lists NpcScaleGlobalBases
// @Accept json
// @Produce json
// @Tags NpcScaleGlobalBase
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcScaleGlobalBase
// @Failure 500 {string} string "Bad query request"
// @Router /npc_scale_global_bases [get]
func (e *NpcScaleGlobalBaseController) listNpcScaleGlobalBases(c echo.Context) error {
	var results []models.NpcScaleGlobalBase
	err := e.db.QueryContext(models.NpcScaleGlobalBase{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcScaleGlobalBase godoc
// @Id getNpcScaleGlobalBase
// @Summary Gets NpcScaleGlobalBase
// @Accept json
// @Produce json
// @Tags NpcScaleGlobalBase
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcScaleGlobalBase
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_scale_global_base/{id} [get]
func (e *NpcScaleGlobalBaseController) getNpcScaleGlobalBase(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	typeId, err := strconv.Atoi(c.Param("typeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TypeId]"})
	}
	params = append(params, typeId)
	keys = append(keys, "type = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// key param [zone_id] position [3] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// key param [instance_version] position [4] type [int]
	if len(c.QueryParam("instance_version")) > 0 {
		instanceVersionParam, err := strconv.Atoi(c.QueryParam("instance_version"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [instance_version] err [%s]", err.Error())})
		}

		params = append(params, instanceVersionParam)
		keys = append(keys, "instance_version = ?")
	}

	// query builder
	var result models.NpcScaleGlobalBase
	query := e.db.QueryContext(models.NpcScaleGlobalBase{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Type == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateNpcScaleGlobalBase godoc
// @Id updateNpcScaleGlobalBase
// @Summary Updates NpcScaleGlobalBase
// @Accept json
// @Produce json
// @Tags NpcScaleGlobalBase
// @Param id path int true "Id"
// @Param npc_scale_global_base body models.NpcScaleGlobalBase true "NpcScaleGlobalBase"
// @Success 200 {array} models.NpcScaleGlobalBase
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_scale_global_base/{id} [patch]
func (e *NpcScaleGlobalBaseController) updateNpcScaleGlobalBase(c echo.Context) error {
	request := new(models.NpcScaleGlobalBase)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	typeId, err := strconv.Atoi(c.Param("typeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TypeId]"})
	}
	params = append(params, typeId)
	keys = append(keys, "type = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// key param [zone_id] position [3] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// key param [instance_version] position [4] type [int]
	if len(c.QueryParam("instance_version")) > 0 {
		instanceVersionParam, err := strconv.Atoi(c.QueryParam("instance_version"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [instance_version] err [%s]", err.Error())})
		}

		params = append(params, instanceVersionParam)
		keys = append(keys, "instance_version = ?")
	}

	// query builder
	var result models.NpcScaleGlobalBase
	query := e.db.QueryContext(models.NpcScaleGlobalBase{}, c)
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
		event := fmt.Sprintf("Updated [NpcScaleGlobalBase] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcScaleGlobalBase godoc
// @Id createNpcScaleGlobalBase
// @Summary Creates NpcScaleGlobalBase
// @Accept json
// @Produce json
// @Param npc_scale_global_base body models.NpcScaleGlobalBase true "NpcScaleGlobalBase"
// @Tags NpcScaleGlobalBase
// @Success 200 {array} models.NpcScaleGlobalBase
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_scale_global_base [put]
func (e *NpcScaleGlobalBaseController) createNpcScaleGlobalBase(c echo.Context) error {
	npcScaleGlobalBase := new(models.NpcScaleGlobalBase)
	if err := c.Bind(npcScaleGlobalBase); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.NpcScaleGlobalBase{}, c).Model(&models.NpcScaleGlobalBase{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&npcScaleGlobalBase).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.NpcScaleGlobalBase{}, npcScaleGlobalBase)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [NpcScaleGlobalBase] [%v] data [%v]", npcScaleGlobalBase.Type, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, npcScaleGlobalBase)
}

// deleteNpcScaleGlobalBase godoc
// @Id deleteNpcScaleGlobalBase
// @Summary Deletes NpcScaleGlobalBase
// @Accept json
// @Produce json
// @Tags NpcScaleGlobalBase
// @Param id path int true "typeId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_scale_global_base/{id} [delete]
func (e *NpcScaleGlobalBaseController) deleteNpcScaleGlobalBase(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	typeId, err := strconv.Atoi(c.Param("typeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, typeId)
	keys = append(keys, "type = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// key param [zone_id] position [3] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// key param [instance_version] position [4] type [int]
	if len(c.QueryParam("instance_version")) > 0 {
		instanceVersionParam, err := strconv.Atoi(c.QueryParam("instance_version"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [instance_version] err [%s]", err.Error())})
		}

		params = append(params, instanceVersionParam)
		keys = append(keys, "instance_version = ?")
	}

	// query builder
	var result models.NpcScaleGlobalBase
	query := e.db.QueryContext(models.NpcScaleGlobalBase{}, c)
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
		event := fmt.Sprintf("Deleted [NpcScaleGlobalBase] [%v] keys [%v]", result.Type, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getNpcScaleGlobalBasesBulk godoc
// @Id getNpcScaleGlobalBasesBulk
// @Summary Gets NpcScaleGlobalBases in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcScaleGlobalBase
// @Success 200 {array} models.NpcScaleGlobalBase
// @Failure 500 {string} string "Bad query request"
// @Router /npc_scale_global_bases/bulk [post]
func (e *NpcScaleGlobalBaseController) getNpcScaleGlobalBasesBulk(c echo.Context) error {
	var results []models.NpcScaleGlobalBase

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

	err := e.db.QueryContext(models.NpcScaleGlobalBase{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcScaleGlobalBasesCount godoc
// @Id getNpcScaleGlobalBasesCount
// @Summary Counts NpcScaleGlobalBases
// @Accept json
// @Produce json
// @Tags NpcScaleGlobalBase
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcScaleGlobalBase
// @Failure 500 {string} string "Bad query request"
// @Router /npc_scale_global_bases/count [get]
func (e *NpcScaleGlobalBaseController) getNpcScaleGlobalBasesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.NpcScaleGlobalBase{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}