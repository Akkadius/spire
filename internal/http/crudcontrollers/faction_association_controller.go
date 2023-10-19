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

type FactionAssociationController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewFactionAssociationController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *FactionAssociationController {
	return &FactionAssociationController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *FactionAssociationController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "faction_association/:id", e.getFactionAssociation, nil),
		routes.RegisterRoute(http.MethodGet, "faction_associations", e.listFactionAssociations, nil),
		routes.RegisterRoute(http.MethodGet, "faction_associations/count", e.getFactionAssociationsCount, nil),
		routes.RegisterRoute(http.MethodPut, "faction_association", e.createFactionAssociation, nil),
		routes.RegisterRoute(http.MethodDelete, "faction_association/:id", e.deleteFactionAssociation, nil),
		routes.RegisterRoute(http.MethodPatch, "faction_association/:id", e.updateFactionAssociation, nil),
		routes.RegisterRoute(http.MethodPost, "faction_associations/bulk", e.getFactionAssociationsBulk, nil),
	}
}

// listFactionAssociations godoc
// @Id listFactionAssociations
// @Summary Lists FactionAssociations
// @Accept json
// @Produce json
// @Tags FactionAssociation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionAssociation
// @Failure 500 {string} string "Bad query request"
// @Router /faction_associations [get]
func (e *FactionAssociationController) listFactionAssociations(c echo.Context) error {
	var results []models.FactionAssociation
	err := e.db.QueryContext(models.FactionAssociation{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getFactionAssociation godoc
// @Id getFactionAssociation
// @Summary Gets FactionAssociation
// @Accept json
// @Produce json
// @Tags FactionAssociation
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionAssociation
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /faction_association/{id} [get]
func (e *FactionAssociationController) getFactionAssociation(c echo.Context) error {
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
	var result models.FactionAssociation
	query := e.db.QueryContext(models.FactionAssociation{}, c)
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

// updateFactionAssociation godoc
// @Id updateFactionAssociation
// @Summary Updates FactionAssociation
// @Accept json
// @Produce json
// @Tags FactionAssociation
// @Param id path int true "Id"
// @Param faction_association body models.FactionAssociation true "FactionAssociation"
// @Success 200 {array} models.FactionAssociation
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /faction_association/{id} [patch]
func (e *FactionAssociationController) updateFactionAssociation(c echo.Context) error {
	request := new(models.FactionAssociation)
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
	var result models.FactionAssociation
	query := e.db.QueryContext(models.FactionAssociation{}, c)
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
		event := fmt.Sprintf("Updated [FactionAssociation] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createFactionAssociation godoc
// @Id createFactionAssociation
// @Summary Creates FactionAssociation
// @Accept json
// @Produce json
// @Param faction_association body models.FactionAssociation true "FactionAssociation"
// @Tags FactionAssociation
// @Success 200 {array} models.FactionAssociation
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /faction_association [put]
func (e *FactionAssociationController) createFactionAssociation(c echo.Context) error {
	factionAssociation := new(models.FactionAssociation)
	if err := c.Bind(factionAssociation); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.FactionAssociation{}, c).Model(&models.FactionAssociation{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&factionAssociation).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.FactionAssociation{}, factionAssociation)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [FactionAssociation] [%v] data [%v]", factionAssociation.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, factionAssociation)
}

// deleteFactionAssociation godoc
// @Id deleteFactionAssociation
// @Summary Deletes FactionAssociation
// @Accept json
// @Produce json
// @Tags FactionAssociation
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /faction_association/{id} [delete]
func (e *FactionAssociationController) deleteFactionAssociation(c echo.Context) error {
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
	var result models.FactionAssociation
	query := e.db.QueryContext(models.FactionAssociation{}, c)
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
		event := fmt.Sprintf("Deleted [FactionAssociation] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getFactionAssociationsBulk godoc
// @Id getFactionAssociationsBulk
// @Summary Gets FactionAssociations in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags FactionAssociation
// @Success 200 {array} models.FactionAssociation
// @Failure 500 {string} string "Bad query request"
// @Router /faction_associations/bulk [post]
func (e *FactionAssociationController) getFactionAssociationsBulk(c echo.Context) error {
	var results []models.FactionAssociation

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

	err := e.db.QueryContext(models.FactionAssociation{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getFactionAssociationsCount godoc
// @Id getFactionAssociationsCount
// @Summary Counts FactionAssociations
// @Accept json
// @Produce json
// @Tags FactionAssociation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionAssociation
// @Failure 500 {string} string "Bad query request"
// @Router /faction_associations/count [get]
func (e *FactionAssociationController) getFactionAssociationsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.FactionAssociation{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}