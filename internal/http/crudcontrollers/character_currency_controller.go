package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CharacterCurrencyController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterCurrencyController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterCurrencyController {
	return &CharacterCurrencyController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterCurrencyController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_currency/:id", e.getCharacterCurrency, nil),
		routes.RegisterRoute(http.MethodGet, "character_currencies", e.listCharacterCurrencies, nil),
		routes.RegisterRoute(http.MethodPut, "character_currency", e.createCharacterCurrency, nil),
		routes.RegisterRoute(http.MethodDelete, "character_currency/:id", e.deleteCharacterCurrency, nil),
		routes.RegisterRoute(http.MethodPatch, "character_currency/:id", e.updateCharacterCurrency, nil),
		routes.RegisterRoute(http.MethodPost, "character_currencies/bulk", e.getCharacterCurrenciesBulk, nil),
	}
}

// listCharacterCurrencies godoc
// @Id listCharacterCurrencies
// @Summary Lists CharacterCurrencies
// @Accept json
// @Produce json
// @Tags CharacterCurrency
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterCurrency
// @Failure 500 {string} string "Bad query request"
// @Router /character_currencies [get]
func (e *CharacterCurrencyController) listCharacterCurrencies(c echo.Context) error {
	var results []models.CharacterCurrency
	err := e.db.QueryContext(models.CharacterCurrency{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterCurrency godoc
// @Id getCharacterCurrency
// @Summary Gets CharacterCurrency
// @Accept json
// @Produce json
// @Tags CharacterCurrency
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterCurrency
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_currency/{id} [get]
func (e *CharacterCurrencyController) getCharacterCurrency(c echo.Context) error {
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
	var result models.CharacterCurrency
	query := e.db.QueryContext(models.CharacterCurrency{}, c)
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

// updateCharacterCurrency godoc
// @Id updateCharacterCurrency
// @Summary Updates CharacterCurrency
// @Accept json
// @Produce json
// @Tags CharacterCurrency
// @Param id path int true "Id"
// @Param character_currency body models.CharacterCurrency true "CharacterCurrency"
// @Success 200 {array} models.CharacterCurrency
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_currency/{id} [patch]
func (e *CharacterCurrencyController) updateCharacterCurrency(c echo.Context) error {
	request := new(models.CharacterCurrency)
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
	var result models.CharacterCurrency
	query := e.db.QueryContext(models.CharacterCurrency{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterCurrency{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterCurrency godoc
// @Id createCharacterCurrency
// @Summary Creates CharacterCurrency
// @Accept json
// @Produce json
// @Param character_currency body models.CharacterCurrency true "CharacterCurrency"
// @Tags CharacterCurrency
// @Success 200 {array} models.CharacterCurrency
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_currency [put]
func (e *CharacterCurrencyController) createCharacterCurrency(c echo.Context) error {
	characterCurrency := new(models.CharacterCurrency)
	if err := c.Bind(characterCurrency); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterCurrency{}, c).Model(&models.CharacterCurrency{}).Create(&characterCurrency).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterCurrency)
}

// deleteCharacterCurrency godoc
// @Id deleteCharacterCurrency
// @Summary Deletes CharacterCurrency
// @Accept json
// @Produce json
// @Tags CharacterCurrency
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_currency/{id} [delete]
func (e *CharacterCurrencyController) deleteCharacterCurrency(c echo.Context) error {
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
	var result models.CharacterCurrency
	query := e.db.QueryContext(models.CharacterCurrency{}, c)
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

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterCurrenciesBulk godoc
// @Id getCharacterCurrenciesBulk
// @Summary Gets CharacterCurrencies in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterCurrency
// @Success 200 {array} models.CharacterCurrency
// @Failure 500 {string} string "Bad query request"
// @Router /character_currencies/bulk [post]
func (e *CharacterCurrencyController) getCharacterCurrenciesBulk(c echo.Context) error {
	var results []models.CharacterCurrency

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

	err := e.db.QueryContext(models.CharacterCurrency{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
