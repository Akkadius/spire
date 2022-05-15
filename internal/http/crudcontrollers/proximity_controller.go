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

type ProximityController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewProximityController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ProximityController {
	return &ProximityController{
		db:	 db,
		logger: logger,
	}
}

func (e *ProximityController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "proximity/:zoneid", e.getProximity, nil),
		routes.RegisterRoute(http.MethodGet, "proximities", e.listProximities, nil),
		routes.RegisterRoute(http.MethodPut, "proximity", e.createProximity, nil),
		routes.RegisterRoute(http.MethodDelete, "proximity/:zoneid", e.deleteProximity, nil),
		routes.RegisterRoute(http.MethodPatch, "proximity/:zoneid", e.updateProximity, nil),
		routes.RegisterRoute(http.MethodPost, "proximities/bulk", e.getProximitiesBulk, nil),
	}
}

// listProximities godoc
// @Id listProximities
// @Summary Lists Proximities
// @Accept json
// @Produce json
// @Tags Proximity
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Proximity
// @Failure 500 {string} string "Bad query request"
// @Router /proximities [get]
func (e *ProximityController) listProximities(c echo.Context) error {
	var results []models.Proximity
	err := e.db.QueryContext(models.Proximity{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getProximity godoc
// @Id getProximity
// @Summary Gets Proximity
// @Accept json
// @Produce json
// @Tags Proximity
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Proximity
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /proximity/{id} [get]
func (e *ProximityController) getProximity(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	zoneid, err := strconv.Atoi(c.Param("zoneid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Zoneid]"})
	}
	params = append(params, zoneid)
	keys = append(keys, "zoneid = ?")

	// key param [exploreid] position [2] type [int]
	if len(c.QueryParam("exploreid")) > 0 {
		exploreidParam, err := strconv.Atoi(c.QueryParam("exploreid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [exploreid] err [%s]", err.Error())})
		}

		params = append(params, exploreidParam)
		keys = append(keys, "exploreid = ?")
	}

	// query builder
	var result models.Proximity
	query := e.db.QueryContext(models.Proximity{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Zoneid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateProximity godoc
// @Id updateProximity
// @Summary Updates Proximity
// @Accept json
// @Produce json
// @Tags Proximity
// @Param id path int true "Id"
// @Param proximity body models.Proximity true "Proximity"
// @Success 200 {array} models.Proximity
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /proximity/{id} [patch]
func (e *ProximityController) updateProximity(c echo.Context) error {
	request := new(models.Proximity)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	zoneid, err := strconv.Atoi(c.Param("zoneid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Zoneid]"})
	}
	params = append(params, zoneid)
	keys = append(keys, "zoneid = ?")

	// key param [exploreid] position [2] type [int]
	if len(c.QueryParam("exploreid")) > 0 {
		exploreidParam, err := strconv.Atoi(c.QueryParam("exploreid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [exploreid] err [%s]", err.Error())})
		}

		params = append(params, exploreidParam)
		keys = append(keys, "exploreid = ?")
	}

	// query builder
	var result models.Proximity
	query := e.db.QueryContext(models.Proximity{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Proximity{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createProximity godoc
// @Id createProximity
// @Summary Creates Proximity
// @Accept json
// @Produce json
// @Param proximity body models.Proximity true "Proximity"
// @Tags Proximity
// @Success 200 {array} models.Proximity
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /proximity [put]
func (e *ProximityController) createProximity(c echo.Context) error {
	proximity := new(models.Proximity)
	if err := c.Bind(proximity); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Proximity{}, c).Model(&models.Proximity{}).Create(&proximity).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, proximity)
}

// deleteProximity godoc
// @Id deleteProximity
// @Summary Deletes Proximity
// @Accept json
// @Produce json
// @Tags Proximity
// @Param id path int true "zoneid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /proximity/{id} [delete]
func (e *ProximityController) deleteProximity(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	zoneid, err := strconv.Atoi(c.Param("zoneid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, zoneid)
	keys = append(keys, "zoneid = ?")

	// key param [exploreid] position [2] type [int]
	if len(c.QueryParam("exploreid")) > 0 {
		exploreidParam, err := strconv.Atoi(c.QueryParam("exploreid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [exploreid] err [%s]", err.Error())})
		}

		params = append(params, exploreidParam)
		keys = append(keys, "exploreid = ?")
	}

	// query builder
	var result models.Proximity
	query := e.db.QueryContext(models.Proximity{}, c)
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

// getProximitiesBulk godoc
// @Id getProximitiesBulk
// @Summary Gets Proximities in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Proximity
// @Success 200 {array} models.Proximity
// @Failure 500 {string} string "Bad query request"
// @Router /proximities/bulk [post]
func (e *ProximityController) getProximitiesBulk(c echo.Context) error {
	var results []models.Proximity

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

	err := e.db.QueryContext(models.Proximity{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
