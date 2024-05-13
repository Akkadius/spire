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

type GridEntryController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewGridEntryController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *GridEntryController {
	return &GridEntryController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *GridEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "grid_entry/:gridid", e.getGridEntry, nil),
		routes.RegisterRoute(http.MethodGet, "grid_entries", e.listGridEntries, nil),
		routes.RegisterRoute(http.MethodGet, "grid_entries/count", e.getGridEntriesCount, nil),
		routes.RegisterRoute(http.MethodPut, "grid_entry", e.createGridEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "grid_entry/:gridid", e.deleteGridEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "grid_entry/:gridid", e.updateGridEntry, nil),
		routes.RegisterRoute(http.MethodPost, "grid_entries/bulk", e.getGridEntriesBulk, nil),
	}
}

// listGridEntries godoc
// @Id listGridEntries
// @Summary Lists GridEntries
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GridEntry
// @Failure 500 {string} string "Bad query request"
// @Router /grid_entries [get]
func (e *GridEntryController) listGridEntries(c echo.Context) error {
	var results []models.GridEntry
	err := e.db.QueryContext(models.GridEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGridEntry godoc
// @Id getGridEntry
// @Summary Gets GridEntry
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GridEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /grid_entry/{id} [get]
func (e *GridEntryController) getGridEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	gridid, err := strconv.Atoi(c.Param("gridid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Gridid]"})
	}
	params = append(params, gridid)
	keys = append(keys, "gridid = ?")

	// key param [zoneid] position [2] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [number] position [3] type [int]
	if len(c.QueryParam("number")) > 0 {
		numberParam, err := strconv.Atoi(c.QueryParam("number"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [number] err [%s]", err.Error())})
		}

		params = append(params, numberParam)
		keys = append(keys, "number = ?")
	}

	// query builder
	var result models.GridEntry
	query := e.db.QueryContext(models.GridEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Gridid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGridEntry godoc
// @Id updateGridEntry
// @Summary Updates GridEntry
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param id path int true "Id"
// @Param grid_entry body models.GridEntry true "GridEntry"
// @Success 200 {array} models.GridEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /grid_entry/{id} [patch]
func (e *GridEntryController) updateGridEntry(c echo.Context) error {
	request := new(models.GridEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	gridid, err := strconv.Atoi(c.Param("gridid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Gridid]"})
	}
	params = append(params, gridid)
	keys = append(keys, "gridid = ?")

	// key param [zoneid] position [2] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [number] position [3] type [int]
	if len(c.QueryParam("number")) > 0 {
		numberParam, err := strconv.Atoi(c.QueryParam("number"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [number] err [%s]", err.Error())})
		}

		params = append(params, numberParam)
		keys = append(keys, "number = ?")
	}

	// query builder
	var result models.GridEntry
	query := e.db.QueryContext(models.GridEntry{}, c)
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
		event := fmt.Sprintf("Updated [GridEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createGridEntry godoc
// @Id createGridEntry
// @Summary Creates GridEntry
// @Accept json
// @Produce json
// @Param grid_entry body models.GridEntry true "GridEntry"
// @Tags GridEntry
// @Success 200 {array} models.GridEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /grid_entry [put]
func (e *GridEntryController) createGridEntry(c echo.Context) error {
	gridEntry := new(models.GridEntry)
	if err := c.Bind(gridEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.GridEntry{}, c).Model(&models.GridEntry{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&gridEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GridEntry{}, gridEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GridEntry] [%v] data [%v]", gridEntry.Gridid, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, gridEntry)
}

// deleteGridEntry godoc
// @Id deleteGridEntry
// @Summary Deletes GridEntry
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param id path int true "gridid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /grid_entry/{id} [delete]
func (e *GridEntryController) deleteGridEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	gridid, err := strconv.Atoi(c.Param("gridid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, gridid)
	keys = append(keys, "gridid = ?")

	// key param [zoneid] position [2] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [number] position [3] type [int]
	if len(c.QueryParam("number")) > 0 {
		numberParam, err := strconv.Atoi(c.QueryParam("number"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [number] err [%s]", err.Error())})
		}

		params = append(params, numberParam)
		keys = append(keys, "number = ?")
	}

	// query builder
	var result models.GridEntry
	query := e.db.QueryContext(models.GridEntry{}, c)
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
		event := fmt.Sprintf("Deleted [GridEntry] [%v] keys [%v]", result.Gridid, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGridEntriesBulk godoc
// @Id getGridEntriesBulk
// @Summary Gets GridEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GridEntry
// @Success 200 {array} models.GridEntry
// @Failure 500 {string} string "Bad query request"
// @Router /grid_entries/bulk [post]
func (e *GridEntryController) getGridEntriesBulk(c echo.Context) error {
	var results []models.GridEntry

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

	err := e.db.QueryContext(models.GridEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGridEntriesCount godoc
// @Id getGridEntriesCount
// @Summary Counts GridEntries
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GridEntry
// @Failure 500 {string} string "Bad query request"
// @Router /grid_entries/count [get]
func (e *GridEntryController) getGridEntriesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.GridEntry{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}