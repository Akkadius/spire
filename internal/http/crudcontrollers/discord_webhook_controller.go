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

type DiscordWebhookController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewDiscordWebhookController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *DiscordWebhookController {
	return &DiscordWebhookController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *DiscordWebhookController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "discord_webhook/:id", e.getDiscordWebhook, nil),
		routes.RegisterRoute(http.MethodGet, "discord_webhooks", e.listDiscordWebhooks, nil),
		routes.RegisterRoute(http.MethodPut, "discord_webhook", e.createDiscordWebhook, nil),
		routes.RegisterRoute(http.MethodDelete, "discord_webhook/:id", e.deleteDiscordWebhook, nil),
		routes.RegisterRoute(http.MethodPatch, "discord_webhook/:id", e.updateDiscordWebhook, nil),
		routes.RegisterRoute(http.MethodPost, "discord_webhooks/bulk", e.getDiscordWebhooksBulk, nil),
	}
}

// listDiscordWebhooks godoc
// @Id listDiscordWebhooks
// @Summary Lists DiscordWebhooks
// @Accept json
// @Produce json
// @Tags DiscordWebhook
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DiscordWebhook
// @Failure 500 {string} string "Bad query request"
// @Router /discord_webhooks [get]
func (e *DiscordWebhookController) listDiscordWebhooks(c echo.Context) error {
	var results []models.DiscordWebhook
	err := e.db.QueryContext(models.DiscordWebhook{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDiscordWebhook godoc
// @Id getDiscordWebhook
// @Summary Gets DiscordWebhook
// @Accept json
// @Produce json
// @Tags DiscordWebhook
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DiscordWebhook
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /discord_webhook/{id} [get]
func (e *DiscordWebhookController) getDiscordWebhook(c echo.Context) error {
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
	var result models.DiscordWebhook
	query := e.db.QueryContext(models.DiscordWebhook{}, c)
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

// updateDiscordWebhook godoc
// @Id updateDiscordWebhook
// @Summary Updates DiscordWebhook
// @Accept json
// @Produce json
// @Tags DiscordWebhook
// @Param id path int true "Id"
// @Param discord_webhook body models.DiscordWebhook true "DiscordWebhook"
// @Success 200 {array} models.DiscordWebhook
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /discord_webhook/{id} [patch]
func (e *DiscordWebhookController) updateDiscordWebhook(c echo.Context) error {
	request := new(models.DiscordWebhook)
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
	var result models.DiscordWebhook
	query := e.db.QueryContext(models.DiscordWebhook{}, c)
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
		event := fmt.Sprintf("Updated [DiscordWebhook] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createDiscordWebhook godoc
// @Id createDiscordWebhook
// @Summary Creates DiscordWebhook
// @Accept json
// @Produce json
// @Param discord_webhook body models.DiscordWebhook true "DiscordWebhook"
// @Tags DiscordWebhook
// @Success 200 {array} models.DiscordWebhook
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /discord_webhook [put]
func (e *DiscordWebhookController) createDiscordWebhook(c echo.Context) error {
	discordWebhook := new(models.DiscordWebhook)
	if err := c.Bind(discordWebhook); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.DiscordWebhook{}, c).Model(&models.DiscordWebhook{}).Create(&discordWebhook).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.DiscordWebhook{}, discordWebhook)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [DiscordWebhook] [%v] data [%v]", discordWebhook.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, discordWebhook)
}

// deleteDiscordWebhook godoc
// @Id deleteDiscordWebhook
// @Summary Deletes DiscordWebhook
// @Accept json
// @Produce json
// @Tags DiscordWebhook
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /discord_webhook/{id} [delete]
func (e *DiscordWebhookController) deleteDiscordWebhook(c echo.Context) error {
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
	var result models.DiscordWebhook
	query := e.db.QueryContext(models.DiscordWebhook{}, c)
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
		event := fmt.Sprintf("Deleted [DiscordWebhook] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDiscordWebhooksBulk godoc
// @Id getDiscordWebhooksBulk
// @Summary Gets DiscordWebhooks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DiscordWebhook
// @Success 200 {array} models.DiscordWebhook
// @Failure 500 {string} string "Bad query request"
// @Router /discord_webhooks/bulk [post]
func (e *DiscordWebhookController) getDiscordWebhooksBulk(c echo.Context) error {
	var results []models.DiscordWebhook

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

	err := e.db.QueryContext(models.DiscordWebhook{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
