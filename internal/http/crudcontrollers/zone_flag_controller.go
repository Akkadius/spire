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

type ZoneFlagController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewZoneFlagController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *ZoneFlagController {
	return &ZoneFlagController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *ZoneFlagController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "zone_flag/:charID", e.getZoneFlag, nil),
		routes.RegisterRoute(http.MethodGet, "zone_flags", e.listZoneFlags, nil),
		routes.RegisterRoute(http.MethodGet, "zone_flags/count", e.getZoneFlagsCount, nil),
		routes.RegisterRoute(http.MethodPut, "zone_flag", e.createZoneFlag, nil),
		routes.RegisterRoute(http.MethodDelete, "zone_flag/:charID", e.deleteZoneFlag, nil),
		routes.RegisterRoute(http.MethodPatch, "zone_flag/:charID", e.updateZoneFlag, nil),
		routes.RegisterRoute(http.MethodPost, "zone_flags/bulk", e.getZoneFlagsBulk, nil),
	}
}

// listZoneFlags godoc
// @Id listZoneFlags
// @Summary Lists ZoneFlags
// @Accept json
// @Produce json
// @Tags ZoneFlag
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZoneFlag
// @Failure 500 {string} string "Bad query request"
// @Router /zone_flags [get]
func (e *ZoneFlagController) listZoneFlags(c echo.Context) error {
	var results []models.ZoneFlag
	err := e.db.QueryContext(models.ZoneFlag{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getZoneFlag godoc
// @Id getZoneFlag
// @Summary Gets ZoneFlag
// @Accept json
// @Produce json
// @Tags ZoneFlag
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZoneFlag
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /zone_flag/{id} [get]
func (e *ZoneFlagController) getZoneFlag(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charID, err := strconv.Atoi(c.Param("charID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharID]"})
	}
	params = append(params, charID)
	keys = append(keys, "charID = ?")

	// key param [zoneID] position [2] type [int]
	if len(c.QueryParam("zoneID")) > 0 {
		zoneIDParam, err := strconv.Atoi(c.QueryParam("zoneID"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneID] err [%s]", err.Error())})
		}

		params = append(params, zoneIDParam)
		keys = append(keys, "zoneID = ?")
	}

	// query builder
	var result models.ZoneFlag
	query := e.db.QueryContext(models.ZoneFlag{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateZoneFlag godoc
// @Id updateZoneFlag
// @Summary Updates ZoneFlag
// @Accept json
// @Produce json
// @Tags ZoneFlag
// @Param id path int true "Id"
// @Param zone_flag body models.ZoneFlag true "ZoneFlag"
// @Success 200 {array} models.ZoneFlag
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /zone_flag/{id} [patch]
func (e *ZoneFlagController) updateZoneFlag(c echo.Context) error {
	request := new(models.ZoneFlag)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charID, err := strconv.Atoi(c.Param("charID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharID]"})
	}
	params = append(params, charID)
	keys = append(keys, "charID = ?")

	// key param [zoneID] position [2] type [int]
	if len(c.QueryParam("zoneID")) > 0 {
		zoneIDParam, err := strconv.Atoi(c.QueryParam("zoneID"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneID] err [%s]", err.Error())})
		}

		params = append(params, zoneIDParam)
		keys = append(keys, "zoneID = ?")
	}

	// query builder
	var result models.ZoneFlag
	query := e.db.QueryContext(models.ZoneFlag{}, c)
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
		event := fmt.Sprintf("Updated [ZoneFlag] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createZoneFlag godoc
// @Id createZoneFlag
// @Summary Creates ZoneFlag
// @Accept json
// @Produce json
// @Param zone_flag body models.ZoneFlag true "ZoneFlag"
// @Tags ZoneFlag
// @Success 200 {array} models.ZoneFlag
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /zone_flag [put]
func (e *ZoneFlagController) createZoneFlag(c echo.Context) error {
	zoneFlag := new(models.ZoneFlag)
	if err := c.Bind(zoneFlag); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.ZoneFlag{}, c).Model(&models.ZoneFlag{}).Create(&zoneFlag).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.ZoneFlag{}, zoneFlag)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [ZoneFlag] [%v] data [%v]", zoneFlag.CharID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, zoneFlag)
}

// deleteZoneFlag godoc
// @Id deleteZoneFlag
// @Summary Deletes ZoneFlag
// @Accept json
// @Produce json
// @Tags ZoneFlag
// @Param id path int true "charID"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /zone_flag/{id} [delete]
func (e *ZoneFlagController) deleteZoneFlag(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charID, err := strconv.Atoi(c.Param("charID"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charID)
	keys = append(keys, "charID = ?")

	// key param [zoneID] position [2] type [int]
	if len(c.QueryParam("zoneID")) > 0 {
		zoneIDParam, err := strconv.Atoi(c.QueryParam("zoneID"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneID] err [%s]", err.Error())})
		}

		params = append(params, zoneIDParam)
		keys = append(keys, "zoneID = ?")
	}

	// query builder
	var result models.ZoneFlag
	query := e.db.QueryContext(models.ZoneFlag{}, c)
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
		event := fmt.Sprintf("Deleted [ZoneFlag] [%v] keys [%v]", result.CharID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getZoneFlagsBulk godoc
// @Id getZoneFlagsBulk
// @Summary Gets ZoneFlags in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ZoneFlag
// @Success 200 {array} models.ZoneFlag
// @Failure 500 {string} string "Bad query request"
// @Router /zone_flags/bulk [post]
func (e *ZoneFlagController) getZoneFlagsBulk(c echo.Context) error {
	var results []models.ZoneFlag

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

	err := e.db.QueryContext(models.ZoneFlag{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getZoneFlagsCount godoc
// @Id getZoneFlagsCount
// @Summary Counts ZoneFlags
// @Accept json
// @Produce json
// @Tags ZoneFlag
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ZoneFlag
// @Failure 500 {string} string "Bad query request"
// @Router /zone_flags/count [get]
func (e *ZoneFlagController) getZoneFlagsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.ZoneFlag{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}