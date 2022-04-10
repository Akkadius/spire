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

type AdventureDetailController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAdventureDetailController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AdventureDetailController {
	return &AdventureDetailController{
		db:	 db,
		logger: logger,
	}
}

func (e *AdventureDetailController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "adventure_detail/:id", e.getAdventureDetail, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_details", e.listAdventureDetails, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_detail", e.createAdventureDetail, nil),
		routes.RegisterRoute(http.MethodDelete, "adventure_detail/:id", e.deleteAdventureDetail, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_detail/:id", e.updateAdventureDetail, nil),
		routes.RegisterRoute(http.MethodPost, "adventure_details/bulk", e.getAdventureDetailsBulk, nil),
	}
}

// listAdventureDetails godoc
// @Id listAdventureDetails
// @Summary Lists AdventureDetails
// @Accept json
// @Produce json
// @Tags AdventureDetail
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureDetail
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_details [get]
func (e *AdventureDetailController) listAdventureDetails(c echo.Context) error {
	var results []models.AdventureDetail
	err := e.db.QueryContext(models.AdventureDetail{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureDetail godoc
// @Id getAdventureDetail
// @Summary Gets AdventureDetail
// @Accept json
// @Produce json
// @Tags AdventureDetail
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureDetail
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_detail/{id} [get]
func (e *AdventureDetailController) getAdventureDetail(c echo.Context) error {
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
	var result models.AdventureDetail
	query := e.db.QueryContext(models.AdventureDetail{}, c)
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

// updateAdventureDetail godoc
// @Id updateAdventureDetail
// @Summary Updates AdventureDetail
// @Accept json
// @Produce json
// @Tags AdventureDetail
// @Param id path int true "Id"
// @Param adventure_detail body models.AdventureDetail true "AdventureDetail"
// @Success 200 {array} models.AdventureDetail
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_detail/{id} [patch]
func (e *AdventureDetailController) updateAdventureDetail(c echo.Context) error {
	request := new(models.AdventureDetail)
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
	var result models.AdventureDetail
	query := e.db.QueryContext(models.AdventureDetail{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AdventureDetail{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAdventureDetail godoc
// @Id createAdventureDetail
// @Summary Creates AdventureDetail
// @Accept json
// @Produce json
// @Param adventure_detail body models.AdventureDetail true "AdventureDetail"
// @Tags AdventureDetail
// @Success 200 {array} models.AdventureDetail
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_detail [put]
func (e *AdventureDetailController) createAdventureDetail(c echo.Context) error {
	adventureDetail := new(models.AdventureDetail)
	if err := c.Bind(adventureDetail); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AdventureDetail{}, c).Model(&models.AdventureDetail{}).Create(&adventureDetail).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, adventureDetail)
}

// deleteAdventureDetail godoc
// @Id deleteAdventureDetail
// @Summary Deletes AdventureDetail
// @Accept json
// @Produce json
// @Tags AdventureDetail
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_detail/{id} [delete]
func (e *AdventureDetailController) deleteAdventureDetail(c echo.Context) error {
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
	var result models.AdventureDetail
	query := e.db.QueryContext(models.AdventureDetail{}, c)
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

// getAdventureDetailsBulk godoc
// @Id getAdventureDetailsBulk
// @Summary Gets AdventureDetails in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AdventureDetail
// @Success 200 {array} models.AdventureDetail
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_details/bulk [post]
func (e *AdventureDetailController) getAdventureDetailsBulk(c echo.Context) error {
	var results []models.AdventureDetail

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

	err := e.db.QueryContext(models.AdventureDetail{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
