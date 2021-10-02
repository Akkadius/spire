package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterCurrencyController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterCurrencyController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterCurrencyController {
	return &CharacterCurrencyController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterCurrencyController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_currency/:character_currency", e.deleteCharacterCurrency, nil),
		routes.RegisterRoute(http.MethodGet, "character_currency/:character_currency", e.getCharacterCurrency, nil),
		routes.RegisterRoute(http.MethodGet, "character_currencies", e.listCharacterCurrencies, nil),
		routes.RegisterRoute(http.MethodPost, "character_currencies/bulk", e.getCharacterCurrenciesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_currency/:character_currency", e.updateCharacterCurrency, nil),
		routes.RegisterRoute(http.MethodPut, "character_currency", e.createCharacterCurrency, nil),
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
	characterCurrencyId, err := strconv.Atoi(c.Param("character_currency"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterCurrency
	err = e.db.QueryContext(models.CharacterCurrency{}, c).First(&result, characterCurrencyId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

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
	characterCurrency := new(models.CharacterCurrency)
	if err := c.Bind(characterCurrency); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.CharacterCurrency{}
	err := e.db.Get(models.CharacterCurrency{}, c).Model(&models.CharacterCurrency{}).First(&entity, characterCurrency.ID).Error
	if err != nil || characterCurrency.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterCurrency{}, c).Model(&entity).Updates(&characterCurrency).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterCurrency)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterCurrency{}, c).Model(&models.CharacterCurrency{}).Create(&characterCurrency).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
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
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_currency/{id} [delete]
func (e *CharacterCurrencyController) deleteCharacterCurrency(c echo.Context) error {
	characterCurrencyId, err := strconv.Atoi(c.Param("character_currency"))
	if err != nil {
		e.logger.Error(err)
	}

	characterCurrency := new(models.CharacterCurrency)
	err = e.db.Get(models.CharacterCurrency{}, c).Model(&models.CharacterCurrency{}).First(&characterCurrency, characterCurrencyId).Error
	if err != nil || characterCurrency.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterCurrency{}, c).Model(&models.CharacterCurrency{}).Delete(&characterCurrency).Error
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
