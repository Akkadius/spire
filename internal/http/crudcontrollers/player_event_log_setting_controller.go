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

type PlayerEventLogSettingController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewPlayerEventLogSettingController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *PlayerEventLogSettingController {
	return &PlayerEventLogSettingController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *PlayerEventLogSettingController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "player_event_log_setting/:id", e.getPlayerEventLogSetting, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_log_settings", e.listPlayerEventLogSettings, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_log_settings/count", e.getPlayerEventLogSettingsCount, nil),
		routes.RegisterRoute(http.MethodPut, "player_event_log_setting", e.createPlayerEventLogSetting, nil),
		routes.RegisterRoute(http.MethodDelete, "player_event_log_setting/:id", e.deletePlayerEventLogSetting, nil),
		routes.RegisterRoute(http.MethodPatch, "player_event_log_setting/:id", e.updatePlayerEventLogSetting, nil),
		routes.RegisterRoute(http.MethodPost, "player_event_log_settings/bulk", e.getPlayerEventLogSettingsBulk, nil),
	}
}

// listPlayerEventLogSettings godoc
// @Id listPlayerEventLogSettings
// @Summary Lists PlayerEventLogSettings
// @Accept json
// @Produce json
// @Tags PlayerEventLogSetting
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventLogSetting
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_log_settings [get]
func (e *PlayerEventLogSettingController) listPlayerEventLogSettings(c echo.Context) error {
	var results []models.PlayerEventLogSetting
	err := e.db.QueryContext(models.PlayerEventLogSetting{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventLogSetting godoc
// @Id getPlayerEventLogSetting
// @Summary Gets PlayerEventLogSetting
// @Accept json
// @Produce json
// @Tags PlayerEventLogSetting
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventLogSetting
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_log_setting/{id} [get]
func (e *PlayerEventLogSettingController) getPlayerEventLogSetting(c echo.Context) error {
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
	var result models.PlayerEventLogSetting
	query := e.db.QueryContext(models.PlayerEventLogSetting{}, c)
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

// updatePlayerEventLogSetting godoc
// @Id updatePlayerEventLogSetting
// @Summary Updates PlayerEventLogSetting
// @Accept json
// @Produce json
// @Tags PlayerEventLogSetting
// @Param id path int true "Id"
// @Param player_event_log_setting body models.PlayerEventLogSetting true "PlayerEventLogSetting"
// @Success 200 {array} models.PlayerEventLogSetting
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /player_event_log_setting/{id} [patch]
func (e *PlayerEventLogSettingController) updatePlayerEventLogSetting(c echo.Context) error {
	request := new(models.PlayerEventLogSetting)
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
	var result models.PlayerEventLogSetting
	query := e.db.QueryContext(models.PlayerEventLogSetting{}, c)
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
		event := fmt.Sprintf("Updated [PlayerEventLogSetting] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createPlayerEventLogSetting godoc
// @Id createPlayerEventLogSetting
// @Summary Creates PlayerEventLogSetting
// @Accept json
// @Produce json
// @Param player_event_log_setting body models.PlayerEventLogSetting true "PlayerEventLogSetting"
// @Tags PlayerEventLogSetting
// @Success 200 {array} models.PlayerEventLogSetting
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /player_event_log_setting [put]
func (e *PlayerEventLogSettingController) createPlayerEventLogSetting(c echo.Context) error {
	playerEventLogSetting := new(models.PlayerEventLogSetting)
	if err := c.Bind(playerEventLogSetting); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.PlayerEventLogSetting{}, c).Model(&models.PlayerEventLogSetting{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&playerEventLogSetting).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.PlayerEventLogSetting{}, playerEventLogSetting)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [PlayerEventLogSetting] [%v] data [%v]", playerEventLogSetting.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, playerEventLogSetting)
}

// deletePlayerEventLogSetting godoc
// @Id deletePlayerEventLogSetting
// @Summary Deletes PlayerEventLogSetting
// @Accept json
// @Produce json
// @Tags PlayerEventLogSetting
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /player_event_log_setting/{id} [delete]
func (e *PlayerEventLogSettingController) deletePlayerEventLogSetting(c echo.Context) error {
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
	var result models.PlayerEventLogSetting
	query := e.db.QueryContext(models.PlayerEventLogSetting{}, c)
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
		event := fmt.Sprintf("Deleted [PlayerEventLogSetting] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getPlayerEventLogSettingsBulk godoc
// @Id getPlayerEventLogSettingsBulk
// @Summary Gets PlayerEventLogSettings in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PlayerEventLogSetting
// @Success 200 {array} models.PlayerEventLogSetting
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_log_settings/bulk [post]
func (e *PlayerEventLogSettingController) getPlayerEventLogSettingsBulk(c echo.Context) error {
	var results []models.PlayerEventLogSetting

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

	err := e.db.QueryContext(models.PlayerEventLogSetting{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventLogSettingsCount godoc
// @Id getPlayerEventLogSettingsCount
// @Summary Counts PlayerEventLogSettings
// @Accept json
// @Produce json
// @Tags PlayerEventLogSetting
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventLogSetting
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_log_settings/count [get]
func (e *PlayerEventLogSettingController) getPlayerEventLogSettingsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.PlayerEventLogSetting{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}