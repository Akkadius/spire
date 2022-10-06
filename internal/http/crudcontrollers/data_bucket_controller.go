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

type DataBucketController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewDataBucketController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *DataBucketController {
	return &DataBucketController{
		db:	    db,
		logger: logger,
	}
}

func (e *DataBucketController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "data_bucket/:id", e.getDataBucket, nil),
		routes.RegisterRoute(http.MethodGet, "data_buckets", e.listDataBuckets, nil),
		routes.RegisterRoute(http.MethodPut, "data_bucket", e.createDataBucket, nil),
		routes.RegisterRoute(http.MethodDelete, "data_bucket/:id", e.deleteDataBucket, nil),
		routes.RegisterRoute(http.MethodPatch, "data_bucket/:id", e.updateDataBucket, nil),
		routes.RegisterRoute(http.MethodPost, "data_buckets/bulk", e.getDataBucketsBulk, nil),
	}
}

// listDataBuckets godoc
// @Id listDataBuckets
// @Summary Lists DataBuckets
// @Accept json
// @Produce json
// @Tags DataBucket
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DataBucket
// @Failure 500 {string} string "Bad query request"
// @Router /data_buckets [get]
func (e *DataBucketController) listDataBuckets(c echo.Context) error {
	var results []models.DataBucket
	err := e.db.QueryContext(models.DataBucket{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDataBucket godoc
// @Id getDataBucket
// @Summary Gets DataBucket
// @Accept json
// @Produce json
// @Tags DataBucket
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DataBucket
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /data_bucket/{id} [get]
func (e *DataBucketController) getDataBucket(c echo.Context) error {
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
	var result models.DataBucket
	query := e.db.QueryContext(models.DataBucket{}, c)
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

// updateDataBucket godoc
// @Id updateDataBucket
// @Summary Updates DataBucket
// @Accept json
// @Produce json
// @Tags DataBucket
// @Param id path int true "Id"
// @Param data_bucket body models.DataBucket true "DataBucket"
// @Success 200 {array} models.DataBucket
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /data_bucket/{id} [patch]
func (e *DataBucketController) updateDataBucket(c echo.Context) error {
	request := new(models.DataBucket)
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
	var result models.DataBucket
	query := e.db.QueryContext(models.DataBucket{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createDataBucket godoc
// @Id createDataBucket
// @Summary Creates DataBucket
// @Accept json
// @Produce json
// @Param data_bucket body models.DataBucket true "DataBucket"
// @Tags DataBucket
// @Success 200 {array} models.DataBucket
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /data_bucket [put]
func (e *DataBucketController) createDataBucket(c echo.Context) error {
	dataBucket := new(models.DataBucket)
	if err := c.Bind(dataBucket); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.DataBucket{}, c).Model(&models.DataBucket{}).Create(&dataBucket).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, dataBucket)
}

// deleteDataBucket godoc
// @Id deleteDataBucket
// @Summary Deletes DataBucket
// @Accept json
// @Produce json
// @Tags DataBucket
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /data_bucket/{id} [delete]
func (e *DataBucketController) deleteDataBucket(c echo.Context) error {
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
	var result models.DataBucket
	query := e.db.QueryContext(models.DataBucket{}, c)
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

// getDataBucketsBulk godoc
// @Id getDataBucketsBulk
// @Summary Gets DataBuckets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DataBucket
// @Success 200 {array} models.DataBucket
// @Failure 500 {string} string "Bad query request"
// @Router /data_buckets/bulk [post]
func (e *DataBucketController) getDataBucketsBulk(c echo.Context) error {
	var results []models.DataBucket

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

	err := e.db.QueryContext(models.DataBucket{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
