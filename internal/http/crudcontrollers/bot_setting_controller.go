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

type BotSettingController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewBotSettingController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *BotSettingController {
	return &BotSettingController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *BotSettingController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_setting/:characterId", e.getBotSetting, nil),
		routes.RegisterRoute(http.MethodGet, "bot_settings", e.listBotSettings, nil),
		routes.RegisterRoute(http.MethodGet, "bot_settings/count", e.getBotSettingsCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_setting", e.createBotSetting, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_setting/:characterId", e.deleteBotSetting, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_setting/:characterId", e.updateBotSetting, nil),
		routes.RegisterRoute(http.MethodPost, "bot_settings/bulk", e.getBotSettingsBulk, nil),
	}
}

// listBotSettings godoc
// @Id listBotSettings
// @Summary Lists BotSettings
// @Accept json
// @Produce json
// @Tags BotSetting
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotSetting
// @Failure 500 {string} string "Bad query request"
// @Router /bot_settings [get]
func (e *BotSettingController) listBotSettings(c echo.Context) error {
	var results []models.BotSetting
	err := e.db.QueryContext(models.BotSetting{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotSetting godoc
// @Id getBotSetting
// @Summary Gets BotSetting
// @Accept json
// @Produce json
// @Tags BotSetting
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotSetting
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_setting/{id} [get]
func (e *BotSettingController) getBotSetting(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharacterId]"})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [bot_id] position [2] type [int]
	if len(c.QueryParam("bot_id")) > 0 {
		botIdParam, err := strconv.Atoi(c.QueryParam("bot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bot_id] err [%s]", err.Error())})
		}

		params = append(params, botIdParam)
		keys = append(keys, "bot_id = ?")
	}

	// key param [stance] position [3] type [tinyint]
	if len(c.QueryParam("stance")) > 0 {
		stanceParam, err := strconv.Atoi(c.QueryParam("stance"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [stance] err [%s]", err.Error())})
		}

		params = append(params, stanceParam)
		keys = append(keys, "stance = ?")
	}

	// key param [setting_id] position [4] type [smallint]
	if len(c.QueryParam("setting_id")) > 0 {
		settingIdParam, err := strconv.Atoi(c.QueryParam("setting_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [setting_id] err [%s]", err.Error())})
		}

		params = append(params, settingIdParam)
		keys = append(keys, "setting_id = ?")
	}

	// key param [setting_type] position [5] type [tinyint]
	if len(c.QueryParam("setting_type")) > 0 {
		settingTypeParam, err := strconv.Atoi(c.QueryParam("setting_type"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [setting_type] err [%s]", err.Error())})
		}

		params = append(params, settingTypeParam)
		keys = append(keys, "setting_type = ?")
	}

	// query builder
	var result models.BotSetting
	query := e.db.QueryContext(models.BotSetting{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharacterId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotSetting godoc
// @Id updateBotSetting
// @Summary Updates BotSetting
// @Accept json
// @Produce json
// @Tags BotSetting
// @Param id path int true "Id"
// @Param bot_setting body models.BotSetting true "BotSetting"
// @Success 200 {array} models.BotSetting
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_setting/{id} [patch]
func (e *BotSettingController) updateBotSetting(c echo.Context) error {
	request := new(models.BotSetting)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharacterId]"})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [bot_id] position [2] type [int]
	if len(c.QueryParam("bot_id")) > 0 {
		botIdParam, err := strconv.Atoi(c.QueryParam("bot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bot_id] err [%s]", err.Error())})
		}

		params = append(params, botIdParam)
		keys = append(keys, "bot_id = ?")
	}

	// key param [stance] position [3] type [tinyint]
	if len(c.QueryParam("stance")) > 0 {
		stanceParam, err := strconv.Atoi(c.QueryParam("stance"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [stance] err [%s]", err.Error())})
		}

		params = append(params, stanceParam)
		keys = append(keys, "stance = ?")
	}

	// key param [setting_id] position [4] type [smallint]
	if len(c.QueryParam("setting_id")) > 0 {
		settingIdParam, err := strconv.Atoi(c.QueryParam("setting_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [setting_id] err [%s]", err.Error())})
		}

		params = append(params, settingIdParam)
		keys = append(keys, "setting_id = ?")
	}

	// key param [setting_type] position [5] type [tinyint]
	if len(c.QueryParam("setting_type")) > 0 {
		settingTypeParam, err := strconv.Atoi(c.QueryParam("setting_type"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [setting_type] err [%s]", err.Error())})
		}

		params = append(params, settingTypeParam)
		keys = append(keys, "setting_type = ?")
	}

	// query builder
	var result models.BotSetting
	query := e.db.QueryContext(models.BotSetting{}, c)
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
		event := fmt.Sprintf("Updated [BotSetting] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotSetting godoc
// @Id createBotSetting
// @Summary Creates BotSetting
// @Accept json
// @Produce json
// @Param bot_setting body models.BotSetting true "BotSetting"
// @Tags BotSetting
// @Success 200 {array} models.BotSetting
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_setting [put]
func (e *BotSettingController) createBotSetting(c echo.Context) error {
	botSetting := new(models.BotSetting)
	if err := c.Bind(botSetting); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotSetting{}, c).Model(&models.BotSetting{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&botSetting).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotSetting{}, botSetting)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotSetting] [%v] data [%v]", botSetting.CharacterId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botSetting)
}

// deleteBotSetting godoc
// @Id deleteBotSetting
// @Summary Deletes BotSetting
// @Accept json
// @Produce json
// @Tags BotSetting
// @Param id path int true "characterId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_setting/{id} [delete]
func (e *BotSettingController) deleteBotSetting(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// key param [bot_id] position [2] type [int]
	if len(c.QueryParam("bot_id")) > 0 {
		botIdParam, err := strconv.Atoi(c.QueryParam("bot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bot_id] err [%s]", err.Error())})
		}

		params = append(params, botIdParam)
		keys = append(keys, "bot_id = ?")
	}

	// key param [stance] position [3] type [tinyint]
	if len(c.QueryParam("stance")) > 0 {
		stanceParam, err := strconv.Atoi(c.QueryParam("stance"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [stance] err [%s]", err.Error())})
		}

		params = append(params, stanceParam)
		keys = append(keys, "stance = ?")
	}

	// key param [setting_id] position [4] type [smallint]
	if len(c.QueryParam("setting_id")) > 0 {
		settingIdParam, err := strconv.Atoi(c.QueryParam("setting_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [setting_id] err [%s]", err.Error())})
		}

		params = append(params, settingIdParam)
		keys = append(keys, "setting_id = ?")
	}

	// key param [setting_type] position [5] type [tinyint]
	if len(c.QueryParam("setting_type")) > 0 {
		settingTypeParam, err := strconv.Atoi(c.QueryParam("setting_type"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [setting_type] err [%s]", err.Error())})
		}

		params = append(params, settingTypeParam)
		keys = append(keys, "setting_type = ?")
	}

	// query builder
	var result models.BotSetting
	query := e.db.QueryContext(models.BotSetting{}, c)
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
		event := fmt.Sprintf("Deleted [BotSetting] [%v] keys [%v]", result.CharacterId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotSettingsBulk godoc
// @Id getBotSettingsBulk
// @Summary Gets BotSettings in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotSetting
// @Success 200 {array} models.BotSetting
// @Failure 500 {string} string "Bad query request"
// @Router /bot_settings/bulk [post]
func (e *BotSettingController) getBotSettingsBulk(c echo.Context) error {
	var results []models.BotSetting

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

	err := e.db.QueryContext(models.BotSetting{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotSettingsCount godoc
// @Id getBotSettingsCount
// @Summary Counts BotSettings
// @Accept json
// @Produce json
// @Tags BotSetting
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotSetting
// @Failure 500 {string} string "Bad query request"
// @Router /bot_settings/count [get]
func (e *BotSettingController) getBotSettingsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotSetting{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}