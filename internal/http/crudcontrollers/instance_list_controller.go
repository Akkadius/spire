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

type InstanceListController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewInstanceListController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *InstanceListController {
	return &InstanceListController{
		db:	 db,
		logger: logger,
	}
}

func (e *InstanceListController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "instance_list/:id", e.getInstanceList, nil),
		routes.RegisterRoute(http.MethodGet, "instance_lists", e.listInstanceLists, nil),
		routes.RegisterRoute(http.MethodPut, "instance_list", e.createInstanceList, nil),
		routes.RegisterRoute(http.MethodDelete, "instance_list/:id", e.deleteInstanceList, nil),
		routes.RegisterRoute(http.MethodPatch, "instance_list/:id", e.updateInstanceList, nil),
		routes.RegisterRoute(http.MethodPost, "instance_lists/bulk", e.getInstanceListsBulk, nil),
	}
}

// listInstanceLists godoc
// @Id listInstanceLists
// @Summary Lists InstanceLists
// @Accept json
// @Produce json
// @Tags InstanceList
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>InstanceListPlayers<br>Zones"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InstanceList
// @Failure 500 {string} string "Bad query request"
// @Router /instance_lists [get]
func (e *InstanceListController) listInstanceLists(c echo.Context) error {
	var results []models.InstanceList
	err := e.db.QueryContext(models.InstanceList{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getInstanceList godoc
// @Id getInstanceList
// @Summary Gets InstanceList
// @Accept json
// @Produce json
// @Tags InstanceList
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>InstanceListPlayers<br>Zones"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InstanceList
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /instance_list/{id} [get]
func (e *InstanceListController) getInstanceList(c echo.Context) error {
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
	var result models.InstanceList
	query := e.db.QueryContext(models.InstanceList{}, c)
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

// updateInstanceList godoc
// @Id updateInstanceList
// @Summary Updates InstanceList
// @Accept json
// @Produce json
// @Tags InstanceList
// @Param id path int true "Id"
// @Param instance_list body models.InstanceList true "InstanceList"
// @Success 200 {array} models.InstanceList
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /instance_list/{id} [patch]
func (e *InstanceListController) updateInstanceList(c echo.Context) error {
	request := new(models.InstanceList)
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
	var result models.InstanceList
	query := e.db.QueryContext(models.InstanceList{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.InstanceList{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createInstanceList godoc
// @Id createInstanceList
// @Summary Creates InstanceList
// @Accept json
// @Produce json
// @Param instance_list body models.InstanceList true "InstanceList"
// @Tags InstanceList
// @Success 200 {array} models.InstanceList
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /instance_list [put]
func (e *InstanceListController) createInstanceList(c echo.Context) error {
	instanceList := new(models.InstanceList)
	if err := c.Bind(instanceList); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.InstanceList{}, c).Model(&models.InstanceList{}).Create(&instanceList).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, instanceList)
}

// deleteInstanceList godoc
// @Id deleteInstanceList
// @Summary Deletes InstanceList
// @Accept json
// @Produce json
// @Tags InstanceList
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /instance_list/{id} [delete]
func (e *InstanceListController) deleteInstanceList(c echo.Context) error {
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
	var result models.InstanceList
	query := e.db.QueryContext(models.InstanceList{}, c)
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

// getInstanceListsBulk godoc
// @Id getInstanceListsBulk
// @Summary Gets InstanceLists in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags InstanceList
// @Success 200 {array} models.InstanceList
// @Failure 500 {string} string "Bad query request"
// @Router /instance_lists/bulk [post]
func (e *InstanceListController) getInstanceListsBulk(c echo.Context) error {
	var results []models.InstanceList

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

	err := e.db.QueryContext(models.InstanceList{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
