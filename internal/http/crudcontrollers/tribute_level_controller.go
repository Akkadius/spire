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

type TributeLevelController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTributeLevelController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TributeLevelController {
	return &TributeLevelController{
		db:	 db,
		logger: logger,
	}
}

func (e *TributeLevelController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "tribute_level/:tributeId", e.getTributeLevel, nil),
		routes.RegisterRoute(http.MethodGet, "tribute_levels", e.listTributeLevels, nil),
		routes.RegisterRoute(http.MethodPut, "tribute_level", e.createTributeLevel, nil),
		routes.RegisterRoute(http.MethodDelete, "tribute_level/:tributeId", e.deleteTributeLevel, nil),
		routes.RegisterRoute(http.MethodPatch, "tribute_level/:tributeId", e.updateTributeLevel, nil),
		routes.RegisterRoute(http.MethodPost, "tribute_levels/bulk", e.getTributeLevelsBulk, nil),
	}
}

// listTributeLevels godoc
// @Id listTributeLevels
// @Summary Lists TributeLevels
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TributeLevel
// @Failure 500 {string} string "Bad query request"
// @Router /tribute_levels [get]
func (e *TributeLevelController) listTributeLevels(c echo.Context) error {
	var results []models.TributeLevel
	err := e.db.QueryContext(models.TributeLevel{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTributeLevel godoc
// @Id getTributeLevel
// @Summary Gets TributeLevel
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TributeLevel
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tribute_level/{id} [get]
func (e *TributeLevelController) getTributeLevel(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	tributeId, err := strconv.Atoi(c.Param("tributeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TributeId]"})
	}
	params = append(params, tributeId)
	keys = append(keys, "tribute_id = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// query builder
	var result models.TributeLevel
	query := e.db.QueryContext(models.TributeLevel{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.TributeId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTributeLevel godoc
// @Id updateTributeLevel
// @Summary Updates TributeLevel
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param id path int true "Id"
// @Param tribute_level body models.TributeLevel true "TributeLevel"
// @Success 200 {array} models.TributeLevel
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /tribute_level/{id} [patch]
func (e *TributeLevelController) updateTributeLevel(c echo.Context) error {
	request := new(models.TributeLevel)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	tributeId, err := strconv.Atoi(c.Param("tributeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TributeId]"})
	}
	params = append(params, tributeId)
	keys = append(keys, "tribute_id = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// query builder
	var result models.TributeLevel
	query := e.db.QueryContext(models.TributeLevel{}, c)
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

// createTributeLevel godoc
// @Id createTributeLevel
// @Summary Creates TributeLevel
// @Accept json
// @Produce json
// @Param tribute_level body models.TributeLevel true "TributeLevel"
// @Tags TributeLevel
// @Success 200 {array} models.TributeLevel
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /tribute_level [put]
func (e *TributeLevelController) createTributeLevel(c echo.Context) error {
	tributeLevel := new(models.TributeLevel)
	if err := c.Bind(tributeLevel); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.TributeLevel{}, c).Model(&models.TributeLevel{}).Create(&tributeLevel).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, tributeLevel)
}

// deleteTributeLevel godoc
// @Id deleteTributeLevel
// @Summary Deletes TributeLevel
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param id path int true "tributeId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tribute_level/{id} [delete]
func (e *TributeLevelController) deleteTributeLevel(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	tributeId, err := strconv.Atoi(c.Param("tributeId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, tributeId)
	keys = append(keys, "tribute_id = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// query builder
	var result models.TributeLevel
	query := e.db.QueryContext(models.TributeLevel{}, c)
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

// getTributeLevelsBulk godoc
// @Id getTributeLevelsBulk
// @Summary Gets TributeLevels in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags TributeLevel
// @Success 200 {array} models.TributeLevel
// @Failure 500 {string} string "Bad query request"
// @Router /tribute_levels/bulk [post]
func (e *TributeLevelController) getTributeLevelsBulk(c echo.Context) error {
	var results []models.TributeLevel

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

	err := e.db.QueryContext(models.TributeLevel{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
