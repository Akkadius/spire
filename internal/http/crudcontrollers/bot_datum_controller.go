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

type BotDatumController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewBotDatumController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *BotDatumController {
	return &BotDatumController{
		db:	 db,
		logger: logger,
	}
}

func (e *BotDatumController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_datum/:botId", e.getBotDatum, nil),
		routes.RegisterRoute(http.MethodGet, "bot_data", e.listBotData, nil),
		routes.RegisterRoute(http.MethodPut, "bot_datum", e.createBotDatum, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_datum/:botId", e.deleteBotDatum, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_datum/:botId", e.updateBotDatum, nil),
		routes.RegisterRoute(http.MethodPost, "bot_data/bulk", e.getBotDataBulk, nil),
	}
}

// listBotData godoc
// @Id listBotData
// @Summary Lists BotData
// @Accept json
// @Produce json
// @Tags BotDatum
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotDatum
// @Failure 500 {string} string "Bad query request"
// @Router /bot_data [get]
func (e *BotDatumController) listBotData(c echo.Context) error {
	var results []models.BotDatum
	err := e.db.QueryContext(models.BotDatum{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotDatum godoc
// @Id getBotDatum
// @Summary Gets BotDatum
// @Accept json
// @Produce json
// @Tags BotDatum
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotDatum
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_datum/{id} [get]
func (e *BotDatumController) getBotDatum(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [BotId]"})
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// query builder
	var result models.BotDatum
	query := e.db.QueryContext(models.BotDatum{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.BotId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotDatum godoc
// @Id updateBotDatum
// @Summary Updates BotDatum
// @Accept json
// @Produce json
// @Tags BotDatum
// @Param id path int true "Id"
// @Param bot_datum body models.BotDatum true "BotDatum"
// @Success 200 {array} models.BotDatum
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_datum/{id} [patch]
func (e *BotDatumController) updateBotDatum(c echo.Context) error {
	request := new(models.BotDatum)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [BotId]"})
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// query builder
	var result models.BotDatum
	query := e.db.QueryContext(models.BotDatum{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.BotDatum{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createBotDatum godoc
// @Id createBotDatum
// @Summary Creates BotDatum
// @Accept json
// @Produce json
// @Param bot_datum body models.BotDatum true "BotDatum"
// @Tags BotDatum
// @Success 200 {array} models.BotDatum
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_datum [put]
func (e *BotDatumController) createBotDatum(c echo.Context) error {
	botDatum := new(models.BotDatum)
	if err := c.Bind(botDatum); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BotDatum{}, c).Model(&models.BotDatum{}).Create(&botDatum).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, botDatum)
}

// deleteBotDatum godoc
// @Id deleteBotDatum
// @Summary Deletes BotDatum
// @Accept json
// @Produce json
// @Tags BotDatum
// @Param id path int true "botId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_datum/{id} [delete]
func (e *BotDatumController) deleteBotDatum(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	botId, err := strconv.Atoi(c.Param("botId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, botId)
	keys = append(keys, "bot_id = ?")

	// query builder
	var result models.BotDatum
	query := e.db.QueryContext(models.BotDatum{}, c)
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

// getBotDataBulk godoc
// @Id getBotDataBulk
// @Summary Gets BotData in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotDatum
// @Success 200 {array} models.BotDatum
// @Failure 500 {string} string "Bad query request"
// @Router /bot_data/bulk [post]
func (e *BotDatumController) getBotDataBulk(c echo.Context) error {
	var results []models.BotDatum

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

	err := e.db.QueryContext(models.BotDatum{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
