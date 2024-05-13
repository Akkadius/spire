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

type LdonTrapTemplateController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewLdonTrapTemplateController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *LdonTrapTemplateController {
	return &LdonTrapTemplateController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *LdonTrapTemplateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "ldon_trap_template/:id", e.getLdonTrapTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "ldon_trap_templates", e.listLdonTrapTemplates, nil),
		routes.RegisterRoute(http.MethodGet, "ldon_trap_templates/count", e.getLdonTrapTemplatesCount, nil),
		routes.RegisterRoute(http.MethodPut, "ldon_trap_template", e.createLdonTrapTemplate, nil),
		routes.RegisterRoute(http.MethodDelete, "ldon_trap_template/:id", e.deleteLdonTrapTemplate, nil),
		routes.RegisterRoute(http.MethodPatch, "ldon_trap_template/:id", e.updateLdonTrapTemplate, nil),
		routes.RegisterRoute(http.MethodPost, "ldon_trap_templates/bulk", e.getLdonTrapTemplatesBulk, nil),
	}
}

// listLdonTrapTemplates godoc
// @Id listLdonTrapTemplates
// @Summary Lists LdonTrapTemplates
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_templates [get]
func (e *LdonTrapTemplateController) listLdonTrapTemplates(c echo.Context) error {
	var results []models.LdonTrapTemplate
	err := e.db.QueryContext(models.LdonTrapTemplate{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getLdonTrapTemplate godoc
// @Id getLdonTrapTemplate
// @Summary Gets LdonTrapTemplate
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_template/{id} [get]
func (e *LdonTrapTemplateController) getLdonTrapTemplate(c echo.Context) error {
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
	var result models.LdonTrapTemplate
	query := e.db.QueryContext(models.LdonTrapTemplate{}, c)
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

// updateLdonTrapTemplate godoc
// @Id updateLdonTrapTemplate
// @Summary Updates LdonTrapTemplate
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param id path int true "Id"
// @Param ldon_trap_template body models.LdonTrapTemplate true "LdonTrapTemplate"
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /ldon_trap_template/{id} [patch]
func (e *LdonTrapTemplateController) updateLdonTrapTemplate(c echo.Context) error {
	request := new(models.LdonTrapTemplate)
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
	var result models.LdonTrapTemplate
	query := e.db.QueryContext(models.LdonTrapTemplate{}, c)
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
		event := fmt.Sprintf("Updated [LdonTrapTemplate] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createLdonTrapTemplate godoc
// @Id createLdonTrapTemplate
// @Summary Creates LdonTrapTemplate
// @Accept json
// @Produce json
// @Param ldon_trap_template body models.LdonTrapTemplate true "LdonTrapTemplate"
// @Tags LdonTrapTemplate
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /ldon_trap_template [put]
func (e *LdonTrapTemplateController) createLdonTrapTemplate(c echo.Context) error {
	ldonTrapTemplate := new(models.LdonTrapTemplate)
	if err := c.Bind(ldonTrapTemplate); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.LdonTrapTemplate{}, c).Model(&models.LdonTrapTemplate{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&ldonTrapTemplate).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.LdonTrapTemplate{}, ldonTrapTemplate)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [LdonTrapTemplate] [%v] data [%v]", ldonTrapTemplate.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, ldonTrapTemplate)
}

// deleteLdonTrapTemplate godoc
// @Id deleteLdonTrapTemplate
// @Summary Deletes LdonTrapTemplate
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /ldon_trap_template/{id} [delete]
func (e *LdonTrapTemplateController) deleteLdonTrapTemplate(c echo.Context) error {
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
	var result models.LdonTrapTemplate
	query := e.db.QueryContext(models.LdonTrapTemplate{}, c)
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
		event := fmt.Sprintf("Deleted [LdonTrapTemplate] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLdonTrapTemplatesBulk godoc
// @Id getLdonTrapTemplatesBulk
// @Summary Gets LdonTrapTemplates in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LdonTrapTemplate
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_templates/bulk [post]
func (e *LdonTrapTemplateController) getLdonTrapTemplatesBulk(c echo.Context) error {
	var results []models.LdonTrapTemplate

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

	err := e.db.QueryContext(models.LdonTrapTemplate{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getLdonTrapTemplatesCount godoc
// @Id getLdonTrapTemplatesCount
// @Summary Counts LdonTrapTemplates
// @Accept json
// @Produce json
// @Tags LdonTrapTemplate
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_templates/count [get]
func (e *LdonTrapTemplateController) getLdonTrapTemplatesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.LdonTrapTemplate{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}