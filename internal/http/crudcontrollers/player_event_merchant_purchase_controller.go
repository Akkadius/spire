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

type PlayerEventMerchantPurchaseController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewPlayerEventMerchantPurchaseController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *PlayerEventMerchantPurchaseController {
	return &PlayerEventMerchantPurchaseController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *PlayerEventMerchantPurchaseController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "player_event_merchant_purchase/:id", e.getPlayerEventMerchantPurchase, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_merchant_purchases", e.listPlayerEventMerchantPurchases, nil),
		routes.RegisterRoute(http.MethodGet, "player_event_merchant_purchases/count", e.getPlayerEventMerchantPurchasesCount, nil),
		routes.RegisterRoute(http.MethodPut, "player_event_merchant_purchase", e.createPlayerEventMerchantPurchase, nil),
		routes.RegisterRoute(http.MethodDelete, "player_event_merchant_purchase/:id", e.deletePlayerEventMerchantPurchase, nil),
		routes.RegisterRoute(http.MethodPatch, "player_event_merchant_purchase/:id", e.updatePlayerEventMerchantPurchase, nil),
		routes.RegisterRoute(http.MethodPost, "player_event_merchant_purchases/bulk", e.getPlayerEventMerchantPurchasesBulk, nil),
	}
}

// listPlayerEventMerchantPurchases godoc
// @Id listPlayerEventMerchantPurchases
// @Summary Lists PlayerEventMerchantPurchases
// @Accept json
// @Produce json
// @Tags PlayerEventMerchantPurchase
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventMerchantPurchase
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_merchant_purchases [get]
func (e *PlayerEventMerchantPurchaseController) listPlayerEventMerchantPurchases(c echo.Context) error {
	var results []models.PlayerEventMerchantPurchase
	err := e.db.QueryContext(models.PlayerEventMerchantPurchase{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventMerchantPurchase godoc
// @Id getPlayerEventMerchantPurchase
// @Summary Gets PlayerEventMerchantPurchase
// @Accept json
// @Produce json
// @Tags PlayerEventMerchantPurchase
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventMerchantPurchase
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_merchant_purchase/{id} [get]
func (e *PlayerEventMerchantPurchaseController) getPlayerEventMerchantPurchase(c echo.Context) error {
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
	var result models.PlayerEventMerchantPurchase
	query := e.db.QueryContext(models.PlayerEventMerchantPurchase{}, c)
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

// updatePlayerEventMerchantPurchase godoc
// @Id updatePlayerEventMerchantPurchase
// @Summary Updates PlayerEventMerchantPurchase
// @Accept json
// @Produce json
// @Tags PlayerEventMerchantPurchase
// @Param id path int true "Id"
// @Param player_event_merchant_purchase body models.PlayerEventMerchantPurchase true "PlayerEventMerchantPurchase"
// @Success 200 {array} models.PlayerEventMerchantPurchase
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /player_event_merchant_purchase/{id} [patch]
func (e *PlayerEventMerchantPurchaseController) updatePlayerEventMerchantPurchase(c echo.Context) error {
	request := new(models.PlayerEventMerchantPurchase)
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
	var result models.PlayerEventMerchantPurchase
	query := e.db.QueryContext(models.PlayerEventMerchantPurchase{}, c)
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
		event := fmt.Sprintf("Updated [PlayerEventMerchantPurchase] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createPlayerEventMerchantPurchase godoc
// @Id createPlayerEventMerchantPurchase
// @Summary Creates PlayerEventMerchantPurchase
// @Accept json
// @Produce json
// @Param player_event_merchant_purchase body models.PlayerEventMerchantPurchase true "PlayerEventMerchantPurchase"
// @Tags PlayerEventMerchantPurchase
// @Success 200 {array} models.PlayerEventMerchantPurchase
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /player_event_merchant_purchase [put]
func (e *PlayerEventMerchantPurchaseController) createPlayerEventMerchantPurchase(c echo.Context) error {
	playerEventMerchantPurchase := new(models.PlayerEventMerchantPurchase)
	if err := c.Bind(playerEventMerchantPurchase); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.PlayerEventMerchantPurchase{}, c).Model(&models.PlayerEventMerchantPurchase{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&playerEventMerchantPurchase).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.PlayerEventMerchantPurchase{}, playerEventMerchantPurchase)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [PlayerEventMerchantPurchase] [%v] data [%v]", playerEventMerchantPurchase.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, playerEventMerchantPurchase)
}

// deletePlayerEventMerchantPurchase godoc
// @Id deletePlayerEventMerchantPurchase
// @Summary Deletes PlayerEventMerchantPurchase
// @Accept json
// @Produce json
// @Tags PlayerEventMerchantPurchase
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /player_event_merchant_purchase/{id} [delete]
func (e *PlayerEventMerchantPurchaseController) deletePlayerEventMerchantPurchase(c echo.Context) error {
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
	var result models.PlayerEventMerchantPurchase
	query := e.db.QueryContext(models.PlayerEventMerchantPurchase{}, c)
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
		event := fmt.Sprintf("Deleted [PlayerEventMerchantPurchase] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getPlayerEventMerchantPurchasesBulk godoc
// @Id getPlayerEventMerchantPurchasesBulk
// @Summary Gets PlayerEventMerchantPurchases in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PlayerEventMerchantPurchase
// @Success 200 {array} models.PlayerEventMerchantPurchase
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_merchant_purchases/bulk [post]
func (e *PlayerEventMerchantPurchaseController) getPlayerEventMerchantPurchasesBulk(c echo.Context) error {
	var results []models.PlayerEventMerchantPurchase

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

	err := e.db.QueryContext(models.PlayerEventMerchantPurchase{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerEventMerchantPurchasesCount godoc
// @Id getPlayerEventMerchantPurchasesCount
// @Summary Counts PlayerEventMerchantPurchases
// @Accept json
// @Produce json
// @Tags PlayerEventMerchantPurchase
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerEventMerchantPurchase
// @Failure 500 {string} string "Bad query request"
// @Router /player_event_merchant_purchases/count [get]
func (e *PlayerEventMerchantPurchaseController) getPlayerEventMerchantPurchasesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.PlayerEventMerchantPurchase{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}