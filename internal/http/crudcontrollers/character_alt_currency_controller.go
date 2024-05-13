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

type CharacterAltCurrencyController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewCharacterAltCurrencyController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *CharacterAltCurrencyController {
	return &CharacterAltCurrencyController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *CharacterAltCurrencyController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_alt_currency/:charId", e.getCharacterAltCurrency, nil),
		routes.RegisterRoute(http.MethodGet, "character_alt_currencies", e.listCharacterAltCurrencies, nil),
		routes.RegisterRoute(http.MethodGet, "character_alt_currencies/count", e.getCharacterAltCurrenciesCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_alt_currency", e.createCharacterAltCurrency, nil),
		routes.RegisterRoute(http.MethodDelete, "character_alt_currency/:charId", e.deleteCharacterAltCurrency, nil),
		routes.RegisterRoute(http.MethodPatch, "character_alt_currency/:charId", e.updateCharacterAltCurrency, nil),
		routes.RegisterRoute(http.MethodPost, "character_alt_currencies/bulk", e.getCharacterAltCurrenciesBulk, nil),
	}
}

// listCharacterAltCurrencies godoc
// @Id listCharacterAltCurrencies
// @Summary Lists CharacterAltCurrencies
// @Accept json
// @Produce json
// @Tags CharacterAltCurrency
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterAltCurrency
// @Failure 500 {string} string "Bad query request"
// @Router /character_alt_currencies [get]
func (e *CharacterAltCurrencyController) listCharacterAltCurrencies(c echo.Context) error {
	var results []models.CharacterAltCurrency
	err := e.db.QueryContext(models.CharacterAltCurrency{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterAltCurrency godoc
// @Id getCharacterAltCurrency
// @Summary Gets CharacterAltCurrency
// @Accept json
// @Produce json
// @Tags CharacterAltCurrency
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterAltCurrency
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_alt_currency/{id} [get]
func (e *CharacterAltCurrencyController) getCharacterAltCurrency(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [currency_id] position [2] type [int]
	if len(c.QueryParam("currency_id")) > 0 {
		currencyIdParam, err := strconv.Atoi(c.QueryParam("currency_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [currency_id] err [%s]", err.Error())})
		}

		params = append(params, currencyIdParam)
		keys = append(keys, "currency_id = ?")
	}

	// query builder
	var result models.CharacterAltCurrency
	query := e.db.QueryContext(models.CharacterAltCurrency{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterAltCurrency godoc
// @Id updateCharacterAltCurrency
// @Summary Updates CharacterAltCurrency
// @Accept json
// @Produce json
// @Tags CharacterAltCurrency
// @Param id path int true "Id"
// @Param character_alt_currency body models.CharacterAltCurrency true "CharacterAltCurrency"
// @Success 200 {array} models.CharacterAltCurrency
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_alt_currency/{id} [patch]
func (e *CharacterAltCurrencyController) updateCharacterAltCurrency(c echo.Context) error {
	request := new(models.CharacterAltCurrency)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [currency_id] position [2] type [int]
	if len(c.QueryParam("currency_id")) > 0 {
		currencyIdParam, err := strconv.Atoi(c.QueryParam("currency_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [currency_id] err [%s]", err.Error())})
		}

		params = append(params, currencyIdParam)
		keys = append(keys, "currency_id = ?")
	}

	// query builder
	var result models.CharacterAltCurrency
	query := e.db.QueryContext(models.CharacterAltCurrency{}, c)
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
		event := fmt.Sprintf("Updated [CharacterAltCurrency] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterAltCurrency godoc
// @Id createCharacterAltCurrency
// @Summary Creates CharacterAltCurrency
// @Accept json
// @Produce json
// @Param character_alt_currency body models.CharacterAltCurrency true "CharacterAltCurrency"
// @Tags CharacterAltCurrency
// @Success 200 {array} models.CharacterAltCurrency
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_alt_currency [put]
func (e *CharacterAltCurrencyController) createCharacterAltCurrency(c echo.Context) error {
	characterAltCurrency := new(models.CharacterAltCurrency)
	if err := c.Bind(characterAltCurrency); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterAltCurrency{}, c).Model(&models.CharacterAltCurrency{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterAltCurrency).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterAltCurrency{}, characterAltCurrency)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterAltCurrency] [%v] data [%v]", characterAltCurrency.CharId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterAltCurrency)
}

// deleteCharacterAltCurrency godoc
// @Id deleteCharacterAltCurrency
// @Summary Deletes CharacterAltCurrency
// @Accept json
// @Produce json
// @Tags CharacterAltCurrency
// @Param id path int true "charId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_alt_currency/{id} [delete]
func (e *CharacterAltCurrencyController) deleteCharacterAltCurrency(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [currency_id] position [2] type [int]
	if len(c.QueryParam("currency_id")) > 0 {
		currencyIdParam, err := strconv.Atoi(c.QueryParam("currency_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [currency_id] err [%s]", err.Error())})
		}

		params = append(params, currencyIdParam)
		keys = append(keys, "currency_id = ?")
	}

	// query builder
	var result models.CharacterAltCurrency
	query := e.db.QueryContext(models.CharacterAltCurrency{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterAltCurrency] [%v] keys [%v]", result.CharId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterAltCurrenciesBulk godoc
// @Id getCharacterAltCurrenciesBulk
// @Summary Gets CharacterAltCurrencies in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterAltCurrency
// @Success 200 {array} models.CharacterAltCurrency
// @Failure 500 {string} string "Bad query request"
// @Router /character_alt_currencies/bulk [post]
func (e *CharacterAltCurrencyController) getCharacterAltCurrenciesBulk(c echo.Context) error {
	var results []models.CharacterAltCurrency

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

	err := e.db.QueryContext(models.CharacterAltCurrency{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterAltCurrenciesCount godoc
// @Id getCharacterAltCurrenciesCount
// @Summary Counts CharacterAltCurrencies
// @Accept json
// @Produce json
// @Tags CharacterAltCurrency
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterAltCurrency
// @Failure 500 {string} string "Bad query request"
// @Router /character_alt_currencies/count [get]
func (e *CharacterAltCurrencyController) getCharacterAltCurrenciesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterAltCurrency{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}