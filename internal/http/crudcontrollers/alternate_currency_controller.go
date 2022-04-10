package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type AlternateCurrencyController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAlternateCurrencyController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AlternateCurrencyController {
	return &AlternateCurrencyController{
		db:	 db,
		logger: logger,
	}
}

func (e *AlternateCurrencyController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "alternate_currency/:id", e.getAlternateCurrency, nil),
		routes.RegisterRoute(http.MethodGet, "alternate_currencies", e.listAlternateCurrencies, nil),
		routes.RegisterRoute(http.MethodPut, "alternate_currency", e.createAlternateCurrency, nil),
		routes.RegisterRoute(http.MethodDelete, "alternate_currency/:id", e.deleteAlternateCurrency, nil),
		routes.RegisterRoute(http.MethodPatch, "alternate_currency/:id", e.updateAlternateCurrency, nil),
		routes.RegisterRoute(http.MethodPost, "alternate_currencies/bulk", e.getAlternateCurrenciesBulk, nil),
	}
}

// listAlternateCurrencies godoc
// @Id listAlternateCurrencies
// @Summary Lists AlternateCurrencies
// @Accept json
// @Produce json
// @Tags AlternateCurrency
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AlternateCurrency
// @Failure 500 {string} string "Bad query request"
// @Router /alternate_currencies [get]
func (e *AlternateCurrencyController) listAlternateCurrencies(c echo.Context) error {
	var results []models.AlternateCurrency
	err := e.db.QueryContext(models.AlternateCurrency{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAlternateCurrency godoc
// @Id getAlternateCurrency
// @Summary Gets AlternateCurrency
// @Accept json
// @Produce json
// @Tags AlternateCurrency
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AlternateCurrency
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /alternate_currency/{id} [get]
func (e *AlternateCurrencyController) getAlternateCurrency(c echo.Context) error {
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
	var result models.AlternateCurrency
	query := e.db.QueryContext(models.AlternateCurrency{}, c)
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

// updateAlternateCurrency godoc
// @Id updateAlternateCurrency
// @Summary Updates AlternateCurrency
// @Accept json
// @Produce json
// @Tags AlternateCurrency
// @Param id path int true "Id"
// @Param alternate_currency body models.AlternateCurrency true "AlternateCurrency"
// @Success 200 {array} models.AlternateCurrency
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /alternate_currency/{id} [patch]
func (e *AlternateCurrencyController) updateAlternateCurrency(c echo.Context) error {
	request := new(models.AlternateCurrency)
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
	var result models.AlternateCurrency
	query := e.db.QueryContext(models.AlternateCurrency{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAlternateCurrency godoc
// @Id createAlternateCurrency
// @Summary Creates AlternateCurrency
// @Accept json
// @Produce json
// @Param alternate_currency body models.AlternateCurrency true "AlternateCurrency"
// @Tags AlternateCurrency
// @Success 200 {array} models.AlternateCurrency
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /alternate_currency [put]
func (e *AlternateCurrencyController) createAlternateCurrency(c echo.Context) error {
	alternateCurrency := new(models.AlternateCurrency)
	if err := c.Bind(alternateCurrency); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AlternateCurrency{}, c).Model(&models.AlternateCurrency{}).Create(&alternateCurrency).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, alternateCurrency)
}

// deleteAlternateCurrency godoc
// @Id deleteAlternateCurrency
// @Summary Deletes AlternateCurrency
// @Accept json
// @Produce json
// @Tags AlternateCurrency
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /alternate_currency/{id} [delete]
func (e *AlternateCurrencyController) deleteAlternateCurrency(c echo.Context) error {
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
	var result models.AlternateCurrency
	query := e.db.QueryContext(models.AlternateCurrency{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.AlternateCurrency{}, c).Model(&models.AlternateCurrency{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getAlternateCurrenciesBulk godoc
// @Id getAlternateCurrenciesBulk
// @Summary Gets AlternateCurrencies in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AlternateCurrency
// @Success 200 {array} models.AlternateCurrency
// @Failure 500 {string} string "Bad query request"
// @Router /alternate_currencies/bulk [post]
func (e *AlternateCurrencyController) getAlternateCurrenciesBulk(c echo.Context) error {
	var results []models.AlternateCurrency

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

	err := e.db.QueryContext(models.AlternateCurrency{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
