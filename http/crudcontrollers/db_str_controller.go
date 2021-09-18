package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DbStrController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewDbStrController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *DbStrController {
	return &DbStrController {
		db:     db,
		logger: logger,
	}
}

func (e *DbStrController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "db_str/:db_str", e.deleteDbStr, nil),
		routes.RegisterRoute(http.MethodGet, "db_str/:db_str", e.getDbStr, nil),
		routes.RegisterRoute(http.MethodGet, "db_strs", e.listDbStrs, nil),
		routes.RegisterRoute(http.MethodPatch, "db_str/:db_str", e.updateDbStr, nil),
		routes.RegisterRoute(http.MethodPut, "db_str", e.createDbStr, nil),
	}
}

// listDbStrs godoc
// @Id listDbStrs
// @Summary Lists DbStrs
// @Accept json
// @Produce json
// @Tags DbStr
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DbStr
// @Failure 500 {string} string "Bad query request"
// @Router /db_strs [get]
func (e *DbStrController) listDbStrs(c echo.Context) error {
	var results []models.DbStr
	err := e.db.QueryContext(models.DbStr{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDbStr godoc
// @Id getDbStr
// @Summary Gets DbStr
// @Accept json
// @Produce json
// @Tags DbStr
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DbStr
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /db_str/{id} [get]
func (e *DbStrController) getDbStr(c echo.Context) error {
	dbStrId, err := strconv.Atoi(c.Param("db_str"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.DbStr
	err = e.db.QueryContext(models.DbStr{}, c).First(&result, dbStrId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateDbStr godoc
// @Id updateDbStr
// @Summary Updates DbStr
// @Accept json
// @Produce json
// @Tags DbStr
// @Param id path int true "Id"
// @Param db_str body models.DbStr true "DbStr"
// @Success 200 {array} models.DbStr
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /db_str/{id} [patch]
func (e *DbStrController) updateDbStr(c echo.Context) error {
	dbStr := new(models.DbStr)
	if err := c.Bind(dbStr); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.DbStr{}, c).Model(&models.DbStr{}).First(&models.DbStr{}, dbStr.ID).Error
	if err != nil || dbStr.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.DbStr{}, c).Model(&models.DbStr{}).Update(&dbStr).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, dbStr)
}

// createDbStr godoc
// @Id createDbStr
// @Summary Creates DbStr
// @Accept json
// @Produce json
// @Param db_str body models.DbStr true "DbStr"
// @Tags DbStr
// @Success 200 {array} models.DbStr
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /db_str [put]
func (e *DbStrController) createDbStr(c echo.Context) error {
	dbStr := new(models.DbStr)
	if err := c.Bind(dbStr); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.DbStr{}, c).Model(&models.DbStr{}).Create(&dbStr).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, dbStr)
}

// deleteDbStr godoc
// @Id deleteDbStr
// @Summary Deletes DbStr
// @Accept json
// @Produce json
// @Tags DbStr
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /db_str/{id} [delete]
func (e *DbStrController) deleteDbStr(c echo.Context) error {
	dbStrId, err := strconv.Atoi(c.Param("db_str"))
	if err != nil {
		e.logger.Error(err)
	}

	dbStr := new(models.DbStr)
	err = e.db.Get(models.DbStr{}, c).Model(&models.DbStr{}).First(&dbStr, dbStrId).Error
	if err != nil || dbStr.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.DbStr{}, c).Model(&models.DbStr{}).Delete(&dbStr).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
