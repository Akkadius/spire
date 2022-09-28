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

type ObjectContentController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewObjectContentController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ObjectContentController {
	return &ObjectContentController{
		db:	 db,
		logger: logger,
	}
}

func (e *ObjectContentController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "object_content/:parentid", e.getObjectContent, nil),
		routes.RegisterRoute(http.MethodGet, "object_contents", e.listObjectContents, nil),
		routes.RegisterRoute(http.MethodPut, "object_content", e.createObjectContent, nil),
		routes.RegisterRoute(http.MethodDelete, "object_content/:parentid", e.deleteObjectContent, nil),
		routes.RegisterRoute(http.MethodPatch, "object_content/:parentid", e.updateObjectContent, nil),
		routes.RegisterRoute(http.MethodPost, "object_contents/bulk", e.getObjectContentsBulk, nil),
	}
}

// listObjectContents godoc
// @Id listObjectContents
// @Summary Lists ObjectContents
// @Accept json
// @Produce json
// @Tags ObjectContent
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ObjectContent
// @Failure 500 {string} string "Bad query request"
// @Router /object_contents [get]
func (e *ObjectContentController) listObjectContents(c echo.Context) error {
	var results []models.ObjectContent
	err := e.db.QueryContext(models.ObjectContent{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getObjectContent godoc
// @Id getObjectContent
// @Summary Gets ObjectContent
// @Accept json
// @Produce json
// @Tags ObjectContent
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ObjectContent
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /object_content/{id} [get]
func (e *ObjectContentController) getObjectContent(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	parentid, err := strconv.Atoi(c.Param("parentid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Parentid]"})
	}
	params = append(params, parentid)
	keys = append(keys, "parentid = ?")

	// key param [bagidx] position [3] type [int]
	if len(c.QueryParam("bagidx")) > 0 {
		bagidxParam, err := strconv.Atoi(c.QueryParam("bagidx"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bagidx] err [%s]", err.Error())})
		}

		params = append(params, bagidxParam)
		keys = append(keys, "bagidx = ?")
	}

	// query builder
	var result models.ObjectContent
	query := e.db.QueryContext(models.ObjectContent{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Parentid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateObjectContent godoc
// @Id updateObjectContent
// @Summary Updates ObjectContent
// @Accept json
// @Produce json
// @Tags ObjectContent
// @Param id path int true "Id"
// @Param object_content body models.ObjectContent true "ObjectContent"
// @Success 200 {array} models.ObjectContent
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /object_content/{id} [patch]
func (e *ObjectContentController) updateObjectContent(c echo.Context) error {
	request := new(models.ObjectContent)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	parentid, err := strconv.Atoi(c.Param("parentid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Parentid]"})
	}
	params = append(params, parentid)
	keys = append(keys, "parentid = ?")

	// key param [bagidx] position [3] type [int]
	if len(c.QueryParam("bagidx")) > 0 {
		bagidxParam, err := strconv.Atoi(c.QueryParam("bagidx"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bagidx] err [%s]", err.Error())})
		}

		params = append(params, bagidxParam)
		keys = append(keys, "bagidx = ?")
	}

	// query builder
	var result models.ObjectContent
	query := e.db.QueryContext(models.ObjectContent{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.ObjectContent{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createObjectContent godoc
// @Id createObjectContent
// @Summary Creates ObjectContent
// @Accept json
// @Produce json
// @Param object_content body models.ObjectContent true "ObjectContent"
// @Tags ObjectContent
// @Success 200 {array} models.ObjectContent
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /object_content [put]
func (e *ObjectContentController) createObjectContent(c echo.Context) error {
	objectContent := new(models.ObjectContent)
	if err := c.Bind(objectContent); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.ObjectContent{}, c).Model(&models.ObjectContent{}).Create(&objectContent).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, objectContent)
}

// deleteObjectContent godoc
// @Id deleteObjectContent
// @Summary Deletes ObjectContent
// @Accept json
// @Produce json
// @Tags ObjectContent
// @Param id path int true "parentid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /object_content/{id} [delete]
func (e *ObjectContentController) deleteObjectContent(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	parentid, err := strconv.Atoi(c.Param("parentid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, parentid)
	keys = append(keys, "parentid = ?")

	// key param [bagidx] position [3] type [int]
	if len(c.QueryParam("bagidx")) > 0 {
		bagidxParam, err := strconv.Atoi(c.QueryParam("bagidx"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bagidx] err [%s]", err.Error())})
		}

		params = append(params, bagidxParam)
		keys = append(keys, "bagidx = ?")
	}

	// query builder
	var result models.ObjectContent
	query := e.db.QueryContext(models.ObjectContent{}, c)
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

// getObjectContentsBulk godoc
// @Id getObjectContentsBulk
// @Summary Gets ObjectContents in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ObjectContent
// @Success 200 {array} models.ObjectContent
// @Failure 500 {string} string "Bad query request"
// @Router /object_contents/bulk [post]
func (e *ObjectContentController) getObjectContentsBulk(c echo.Context) error {
	var results []models.ObjectContent

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

	err := e.db.QueryContext(models.ObjectContent{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
