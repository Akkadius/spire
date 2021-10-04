package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ObjectController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewObjectController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ObjectController {
	return &ObjectController{
		db:     db,
		logger: logger,
	}
}

func (e *ObjectController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "object/:object", e.deleteObject, nil),
		routes.RegisterRoute(http.MethodGet, "object/:object", e.getObject, nil),
		routes.RegisterRoute(http.MethodGet, "objects", e.listObjects, nil),
		routes.RegisterRoute(http.MethodPost, "objects/bulk", e.getObjectsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "object/:object", e.updateObject, nil),
		routes.RegisterRoute(http.MethodPut, "object", e.createObject, nil),
	}
}

// listObjects godoc
// @Id listObjects
// @Summary Lists Objects
// @Accept json
// @Produce json
// @Tags Object
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Object
// @Failure 500 {string} string "Bad query request"
// @Router /objects [get]
func (e *ObjectController) listObjects(c echo.Context) error {
	var results []models.Object
	err := e.db.QueryContext(models.Object{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getObject godoc
// @Id getObject
// @Summary Gets Object
// @Accept json
// @Produce json
// @Tags Object
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Object
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /object/{id} [get]
func (e *ObjectController) getObject(c echo.Context) error {
	objectId, err := strconv.Atoi(c.Param("object"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Object
	err = e.db.QueryContext(models.Object{}, c).First(&result, objectId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateObject godoc
// @Id updateObject
// @Summary Updates Object
// @Accept json
// @Produce json
// @Tags Object
// @Param id path int true "Id"
// @Param object body models.Object true "Object"
// @Success 200 {array} models.Object
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /object/{id} [patch]
func (e *ObjectController) updateObject(c echo.Context) error {
	object := new(models.Object)
	if err := c.Bind(object); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Object{}
	err := e.db.Get(models.Object{}, c).Model(&models.Object{}).First(&entity, object.ID).Error
	if err != nil || object.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Object{}, c).Model(&entity).Updates(&object).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, object)
}

// createObject godoc
// @Id createObject
// @Summary Creates Object
// @Accept json
// @Produce json
// @Param object body models.Object true "Object"
// @Tags Object
// @Success 200 {array} models.Object
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /object [put]
func (e *ObjectController) createObject(c echo.Context) error {
	object := new(models.Object)
	if err := c.Bind(object); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Object{}, c).Model(&models.Object{}).Create(&object).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, object)
}

// deleteObject godoc
// @Id deleteObject
// @Summary Deletes Object
// @Accept json
// @Produce json
// @Tags Object
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /object/{id} [delete]
func (e *ObjectController) deleteObject(c echo.Context) error {
	objectId, err := strconv.Atoi(c.Param("object"))
	if err != nil {
		e.logger.Error(err)
	}

	object := new(models.Object)
	err = e.db.Get(models.Object{}, c).Model(&models.Object{}).First(&object, objectId).Error
	if err != nil || object.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Object{}, c).Model(&models.Object{}).Delete(&object).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getObjectsBulk godoc
// @Id getObjectsBulk
// @Summary Gets Objects in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Object
// @Success 200 {array} models.Object
// @Failure 500 {string} string "Bad query request"
// @Router /objects/bulk [post]
func (e *ObjectController) getObjectsBulk(c echo.Context) error {
	var results []models.Object

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

	err := e.db.QueryContext(models.Object{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
