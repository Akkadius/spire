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

type SpellsNewController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewSpellsNewController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *SpellsNewController {
	return &SpellsNewController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *SpellsNewController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spells_new/:id", e.getSpellsNew, nil),
		routes.RegisterRoute(http.MethodGet, "spells_news", e.listSpellsNews, nil),
		routes.RegisterRoute(http.MethodGet, "spells_news/count", e.getSpellsNewsCount, nil),
		routes.RegisterRoute(http.MethodPut, "spells_new", e.createSpellsNew, nil),
		routes.RegisterRoute(http.MethodDelete, "spells_new/:id", e.deleteSpellsNew, nil),
		routes.RegisterRoute(http.MethodPatch, "spells_new/:id", e.updateSpellsNew, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getSpellsNewsBulk, nil),
	}
}

// listSpellsNews godoc
// @Id listSpellsNews
// @Summary Lists SpellsNews
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Bad query request"
// @Router /spells_news [get]
func (e *SpellsNewController) listSpellsNews(c echo.Context) error {
	var results []models.SpellsNew
	err := e.db.QueryContext(models.SpellsNew{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpellsNew godoc
// @Id getSpellsNew
// @Summary Gets SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellsNew
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spells_new/{id} [get]
func (e *SpellsNewController) getSpellsNew(c echo.Context) error {
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
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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

// updateSpellsNew godoc
// @Id updateSpellsNew
// @Summary Updates SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "Id"
// @Param spells_new body models.SpellsNew true "SpellsNew"
// @Success 200 {array} models.SpellsNew
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spells_new/{id} [patch]
func (e *SpellsNewController) updateSpellsNew(c echo.Context) error {
	request := new(models.SpellsNew)
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
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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
		event := fmt.Sprintf("Updated [SpellsNew] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createSpellsNew godoc
// @Id createSpellsNew
// @Summary Creates SpellsNew
// @Accept json
// @Produce json
// @Param spells_new body models.SpellsNew true "SpellsNew"
// @Tags SpellsNew
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spells_new [put]
func (e *SpellsNewController) createSpellsNew(c echo.Context) error {
	spellsNew := new(models.SpellsNew)
	if err := c.Bind(spellsNew); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.SpellsNew{}, c).Model(&models.SpellsNew{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&spellsNew).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.SpellsNew{}, spellsNew)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [SpellsNew] [%v] data [%v]", spellsNew.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, spellsNew)
}

// deleteSpellsNew godoc
// @Id deleteSpellsNew
// @Summary Deletes SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spells_new/{id} [delete]
func (e *SpellsNewController) deleteSpellsNew(c echo.Context) error {
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
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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
		event := fmt.Sprintf("Deleted [SpellsNew] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpellsNewsBulk godoc
// @Id getSpellsNewsBulk
// @Summary Gets SpellsNews in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpellsNew
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Bad query request"
// @Router /spells_news/bulk [post]
func (e *SpellsNewController) getSpellsNewsBulk(c echo.Context) error {
	var results []models.SpellsNew

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

	err := e.db.QueryContext(models.SpellsNew{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpellsNewsCount godoc
// @Id getSpellsNewsCount
// @Summary Counts SpellsNews
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Bad query request"
// @Router /spells_news/count [get]
func (e *SpellsNewController) getSpellsNewsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.SpellsNew{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}