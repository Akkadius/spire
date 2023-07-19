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
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type DamageshieldtypeController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewDamageshieldtypeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *DamageshieldtypeController {
	return &DamageshieldtypeController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *DamageshieldtypeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "damageshieldtype/:spellid", e.getDamageshieldtype, nil),
		routes.RegisterRoute(http.MethodGet, "damageshieldtypes", e.listDamageshieldtypes, nil),
		routes.RegisterRoute(http.MethodGet, "damageshieldtypes/count", e.getDamageshieldtypesCount, nil),
		routes.RegisterRoute(http.MethodPut, "damageshieldtype", e.createDamageshieldtype, nil),
		routes.RegisterRoute(http.MethodDelete, "damageshieldtype/:spellid", e.deleteDamageshieldtype, nil),
		routes.RegisterRoute(http.MethodPatch, "damageshieldtype/:spellid", e.updateDamageshieldtype, nil),
		routes.RegisterRoute(http.MethodPost, "damageshieldtypes/bulk", e.getDamageshieldtypesBulk, nil),
	}
}

// listDamageshieldtypes godoc
// @Id listDamageshieldtypes
// @Summary Lists Damageshieldtypes
// @Accept json
// @Produce json
// @Tags Damageshieldtype
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Damageshieldtype
// @Failure 500 {string} string "Bad query request"
// @Router /damageshieldtypes [get]
func (e *DamageshieldtypeController) listDamageshieldtypes(c echo.Context) error {
	var results []models.Damageshieldtype
	err := e.db.QueryContext(models.Damageshieldtype{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDamageshieldtype godoc
// @Id getDamageshieldtype
// @Summary Gets Damageshieldtype
// @Accept json
// @Produce json
// @Tags Damageshieldtype
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Damageshieldtype
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /damageshieldtype/{id} [get]
func (e *DamageshieldtypeController) getDamageshieldtype(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	spellid, err := strconv.Atoi(c.Param("spellid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Spellid]"})
	}
	params = append(params, spellid)
	keys = append(keys, "spellid = ?")

	// query builder
	var result models.Damageshieldtype
	query := e.db.QueryContext(models.Damageshieldtype{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Spellid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateDamageshieldtype godoc
// @Id updateDamageshieldtype
// @Summary Updates Damageshieldtype
// @Accept json
// @Produce json
// @Tags Damageshieldtype
// @Param id path int true "Id"
// @Param damageshieldtype body models.Damageshieldtype true "Damageshieldtype"
// @Success 200 {array} models.Damageshieldtype
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /damageshieldtype/{id} [patch]
func (e *DamageshieldtypeController) updateDamageshieldtype(c echo.Context) error {
	request := new(models.Damageshieldtype)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	spellid, err := strconv.Atoi(c.Param("spellid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Spellid]"})
	}
	params = append(params, spellid)
	keys = append(keys, "spellid = ?")

	// query builder
	var result models.Damageshieldtype
	query := e.db.QueryContext(models.Damageshieldtype{}, c)
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
		event := fmt.Sprintf("Updated [Damageshieldtype] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createDamageshieldtype godoc
// @Id createDamageshieldtype
// @Summary Creates Damageshieldtype
// @Accept json
// @Produce json
// @Param damageshieldtype body models.Damageshieldtype true "Damageshieldtype"
// @Tags Damageshieldtype
// @Success 200 {array} models.Damageshieldtype
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /damageshieldtype [put]
func (e *DamageshieldtypeController) createDamageshieldtype(c echo.Context) error {
	damageshieldtype := new(models.Damageshieldtype)
	if err := c.Bind(damageshieldtype); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.Damageshieldtype{}, c).Model(&models.Damageshieldtype{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&damageshieldtype).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.Damageshieldtype{}, damageshieldtype)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [Damageshieldtype] [%v] data [%v]", damageshieldtype.Spellid, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, damageshieldtype)
}

// deleteDamageshieldtype godoc
// @Id deleteDamageshieldtype
// @Summary Deletes Damageshieldtype
// @Accept json
// @Produce json
// @Tags Damageshieldtype
// @Param id path int true "spellid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /damageshieldtype/{id} [delete]
func (e *DamageshieldtypeController) deleteDamageshieldtype(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	spellid, err := strconv.Atoi(c.Param("spellid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, spellid)
	keys = append(keys, "spellid = ?")

	// query builder
	var result models.Damageshieldtype
	query := e.db.QueryContext(models.Damageshieldtype{}, c)
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
		event := fmt.Sprintf("Deleted [Damageshieldtype] [%v] keys [%v]", result.Spellid, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDamageshieldtypesBulk godoc
// @Id getDamageshieldtypesBulk
// @Summary Gets Damageshieldtypes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Damageshieldtype
// @Success 200 {array} models.Damageshieldtype
// @Failure 500 {string} string "Bad query request"
// @Router /damageshieldtypes/bulk [post]
func (e *DamageshieldtypeController) getDamageshieldtypesBulk(c echo.Context) error {
	var results []models.Damageshieldtype

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

	err := e.db.QueryContext(models.Damageshieldtype{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getDamageshieldtypesCount godoc
// @Id getDamageshieldtypesCount
// @Summary Counts Damageshieldtypes
// @Accept json
// @Produce json
// @Tags Damageshieldtype
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Damageshieldtype
// @Failure 500 {string} string "Bad query request"
// @Router /damageshieldtypes/count [get]
func (e *DamageshieldtypeController) getDamageshieldtypesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.Damageshieldtype{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}