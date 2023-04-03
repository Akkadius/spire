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
	"net/http"
	"strconv"
	"strings"
)

type ChatchannelReservedNameController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewChatchannelReservedNameController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *ChatchannelReservedNameController {
	return &ChatchannelReservedNameController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *ChatchannelReservedNameController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "chatchannel_reserved_name/:id", e.getChatchannelReservedName, nil),
		routes.RegisterRoute(http.MethodGet, "chatchannel_reserved_names", e.listChatchannelReservedNames, nil),
		routes.RegisterRoute(http.MethodGet, "chatchannel_reserved_names/count", e.getChatchannelReservedNamesCount, nil),
		routes.RegisterRoute(http.MethodPut, "chatchannel_reserved_name", e.createChatchannelReservedName, nil),
		routes.RegisterRoute(http.MethodDelete, "chatchannel_reserved_name/:id", e.deleteChatchannelReservedName, nil),
		routes.RegisterRoute(http.MethodPatch, "chatchannel_reserved_name/:id", e.updateChatchannelReservedName, nil),
		routes.RegisterRoute(http.MethodPost, "chatchannel_reserved_names/bulk", e.getChatchannelReservedNamesBulk, nil),
	}
}

// listChatchannelReservedNames godoc
// @Id listChatchannelReservedNames
// @Summary Lists ChatchannelReservedNames
// @Accept json
// @Produce json
// @Tags ChatchannelReservedName
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ChatchannelReservedName
// @Failure 500 {string} string "Bad query request"
// @Router /chatchannel_reserved_names [get]
func (e *ChatchannelReservedNameController) listChatchannelReservedNames(c echo.Context) error {
	var results []models.ChatchannelReservedName
	err := e.db.QueryContext(models.ChatchannelReservedName{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getChatchannelReservedName godoc
// @Id getChatchannelReservedName
// @Summary Gets ChatchannelReservedName
// @Accept json
// @Produce json
// @Tags ChatchannelReservedName
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ChatchannelReservedName
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /chatchannel_reserved_name/{id} [get]
func (e *ChatchannelReservedNameController) getChatchannelReservedName(c echo.Context) error {
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
	var result models.ChatchannelReservedName
	query := e.db.QueryContext(models.ChatchannelReservedName{}, c)
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

// updateChatchannelReservedName godoc
// @Id updateChatchannelReservedName
// @Summary Updates ChatchannelReservedName
// @Accept json
// @Produce json
// @Tags ChatchannelReservedName
// @Param id path int true "Id"
// @Param chatchannel_reserved_name body models.ChatchannelReservedName true "ChatchannelReservedName"
// @Success 200 {array} models.ChatchannelReservedName
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /chatchannel_reserved_name/{id} [patch]
func (e *ChatchannelReservedNameController) updateChatchannelReservedName(c echo.Context) error {
	request := new(models.ChatchannelReservedName)
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
	var result models.ChatchannelReservedName
	query := e.db.QueryContext(models.ChatchannelReservedName{}, c)
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
		event := fmt.Sprintf("Updated [ChatchannelReservedName] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createChatchannelReservedName godoc
// @Id createChatchannelReservedName
// @Summary Creates ChatchannelReservedName
// @Accept json
// @Produce json
// @Param chatchannel_reserved_name body models.ChatchannelReservedName true "ChatchannelReservedName"
// @Tags ChatchannelReservedName
// @Success 200 {array} models.ChatchannelReservedName
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /chatchannel_reserved_name [put]
func (e *ChatchannelReservedNameController) createChatchannelReservedName(c echo.Context) error {
	chatchannelReservedName := new(models.ChatchannelReservedName)
	if err := c.Bind(chatchannelReservedName); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.ChatchannelReservedName{}, c).Model(&models.ChatchannelReservedName{}).Create(&chatchannelReservedName).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.ChatchannelReservedName{}, chatchannelReservedName)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [ChatchannelReservedName] [%v] data [%v]", chatchannelReservedName.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, chatchannelReservedName)
}

// deleteChatchannelReservedName godoc
// @Id deleteChatchannelReservedName
// @Summary Deletes ChatchannelReservedName
// @Accept json
// @Produce json
// @Tags ChatchannelReservedName
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /chatchannel_reserved_name/{id} [delete]
func (e *ChatchannelReservedNameController) deleteChatchannelReservedName(c echo.Context) error {
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
	var result models.ChatchannelReservedName
	query := e.db.QueryContext(models.ChatchannelReservedName{}, c)
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
		event := fmt.Sprintf("Deleted [ChatchannelReservedName] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getChatchannelReservedNamesBulk godoc
// @Id getChatchannelReservedNamesBulk
// @Summary Gets ChatchannelReservedNames in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ChatchannelReservedName
// @Success 200 {array} models.ChatchannelReservedName
// @Failure 500 {string} string "Bad query request"
// @Router /chatchannel_reserved_names/bulk [post]
func (e *ChatchannelReservedNameController) getChatchannelReservedNamesBulk(c echo.Context) error {
	var results []models.ChatchannelReservedName

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

	err := e.db.QueryContext(models.ChatchannelReservedName{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getChatchannelReservedNamesCount godoc
// @Id getChatchannelReservedNamesCount
// @Summary Counts ChatchannelReservedNames
// @Accept json
// @Produce json
// @Tags ChatchannelReservedName
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ChatchannelReservedName
// @Failure 500 {string} string "Bad query request"
// @Router /chatchannel_reserved_names/count [get]
func (e *ChatchannelReservedNameController) getChatchannelReservedNamesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.ChatchannelReservedName{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}