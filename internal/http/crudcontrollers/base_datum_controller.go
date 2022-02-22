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

type BaseDatumController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewBaseDatumController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *BaseDatumController {
	return &BaseDatumController{
		db:	 db,
		logger: logger,
	}
}

func (e *BaseDatumController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "base_datum/:level", e.getBaseDatum, nil),
		routes.RegisterRoute(http.MethodGet, "base_data", e.listBaseData, nil),
		routes.RegisterRoute(http.MethodPut, "base_datum", e.createBaseDatum, nil),
		routes.RegisterRoute(http.MethodDelete, "base_datum/:level", e.deleteBaseDatum, nil),
		routes.RegisterRoute(http.MethodPatch, "base_datum/:level", e.updateBaseDatum, nil),
		routes.RegisterRoute(http.MethodPost, "base_data/bulk", e.getBaseDataBulk, nil),
	}
}

// listBaseData godoc
// @Id listBaseData
// @Summary Lists BaseData
// @Accept json
// @Produce json
// @Tags BaseDatum
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BaseDatum
// @Failure 500 {string} string "Bad query request"
// @Router /base_data [get]
func (e *BaseDatumController) listBaseData(c echo.Context) error {
	var results []models.BaseDatum
	err := e.db.QueryContext(models.BaseDatum{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBaseDatum godoc
// @Id getBaseDatum
// @Summary Gets BaseDatum
// @Accept json
// @Produce json
// @Tags BaseDatum
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BaseDatum
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /base_datum/{id} [get]
func (e *BaseDatumController) getBaseDatum(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Level]"})
	}
	params = append(params, level)
	keys = append(keys, "level = ?")

	// key param [class] position [2] type [int]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// query builder
	var result models.BaseDatum
	query := e.db.QueryContext(models.BaseDatum{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Level == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBaseDatum godoc
// @Id updateBaseDatum
// @Summary Updates BaseDatum
// @Accept json
// @Produce json
// @Tags BaseDatum
// @Param id path int true "Id"
// @Param base_datum body models.BaseDatum true "BaseDatum"
// @Success 200 {array} models.BaseDatum
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /base_datum/{id} [patch]
func (e *BaseDatumController) updateBaseDatum(c echo.Context) error {
	request := new(models.BaseDatum)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Level]"})
	}
	params = append(params, level)
	keys = append(keys, "level = ?")

	// key param [class] position [2] type [int]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// query builder
	var result models.BaseDatum
	query := e.db.QueryContext(models.BaseDatum{}, c)
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

// createBaseDatum godoc
// @Id createBaseDatum
// @Summary Creates BaseDatum
// @Accept json
// @Produce json
// @Param base_datum body models.BaseDatum true "BaseDatum"
// @Tags BaseDatum
// @Success 200 {array} models.BaseDatum
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /base_datum [put]
func (e *BaseDatumController) createBaseDatum(c echo.Context) error {
	baseDatum := new(models.BaseDatum)
	if err := c.Bind(baseDatum); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BaseDatum{}, c).Model(&models.BaseDatum{}).Create(&baseDatum).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, baseDatum)
}

// deleteBaseDatum godoc
// @Id deleteBaseDatum
// @Summary Deletes BaseDatum
// @Accept json
// @Produce json
// @Tags BaseDatum
// @Param id path int true "level"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /base_datum/{id} [delete]
func (e *BaseDatumController) deleteBaseDatum(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, level)
	keys = append(keys, "level = ?")

	// key param [class] position [2] type [int]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// query builder
	var result models.BaseDatum
	query := e.db.QueryContext(models.BaseDatum{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.BaseDatum{}, c).Model(&models.BaseDatum{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBaseDataBulk godoc
// @Id getBaseDataBulk
// @Summary Gets BaseData in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BaseDatum
// @Success 200 {array} models.BaseDatum
// @Failure 500 {string} string "Bad query request"
// @Router /base_data/bulk [post]
func (e *BaseDatumController) getBaseDataBulk(c echo.Context) error {
	var results []models.BaseDatum

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

	err := e.db.QueryContext(models.BaseDatum{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
