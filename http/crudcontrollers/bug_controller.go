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

type BugController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewBugController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *BugController {
	return &BugController {
		db:     db,
		logger: logger,
	}
}

func (e *BugController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "bug/:bug", e.deleteBug, nil),
		routes.RegisterRoute(http.MethodGet, "bug/:bug", e.getBug, nil),
		routes.RegisterRoute(http.MethodGet, "bugs", e.listBugs, nil),
		routes.RegisterRoute(http.MethodPatch, "bug/:bug", e.updateBug, nil),
		routes.RegisterRoute(http.MethodPut, "bug", e.createBug, nil),
	}
}

// listBugs godoc
// @Id listBugs
// @Summary Lists Bugs
// @Accept json
// @Produce json
// @Tags Bug
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Bug
// @Failure 500 {string} string "Bad query request"
// @Router /bugs [get]
func (e *BugController) listBugs(c echo.Context) error {
	var results []models.Bug
	err := e.db.QueryContext(models.Bug{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBug godoc
// @Id getBug
// @Summary Gets Bug
// @Accept json
// @Produce json
// @Tags Bug
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Bug
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bug/{id} [get]
func (e *BugController) getBug(c echo.Context) error {
	bugId, err := strconv.Atoi(c.Param("bug"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Bug
	err = e.db.QueryContext(models.Bug{}, c).First(&result, bugId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBug godoc
// @Id updateBug
// @Summary Updates Bug
// @Accept json
// @Produce json
// @Tags Bug
// @Param id path int true "Id"
// @Param bug body models.Bug true "Bug"
// @Success 200 {array} models.Bug
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bug/{id} [patch]
func (e *BugController) updateBug(c echo.Context) error {
	bug := new(models.Bug)
	if err := c.Bind(bug); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Bug{}, c).Model(&models.Bug{}).First(&models.Bug{}, bug.ID).Error
	if err != nil || bug.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Bug{}, c).Model(&models.Bug{}).Update(&bug).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, bug)
}

// createBug godoc
// @Id createBug
// @Summary Creates Bug
// @Accept json
// @Produce json
// @Param bug body models.Bug true "Bug"
// @Tags Bug
// @Success 200 {array} models.Bug
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bug [put]
func (e *BugController) createBug(c echo.Context) error {
	bug := new(models.Bug)
	if err := c.Bind(bug); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Bug{}, c).Model(&models.Bug{}).Create(&bug).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, bug)
}

// deleteBug godoc
// @Id deleteBug
// @Summary Deletes Bug
// @Accept json
// @Produce json
// @Tags Bug
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bug/{id} [delete]
func (e *BugController) deleteBug(c echo.Context) error {
	bugId, err := strconv.Atoi(c.Param("bug"))
	if err != nil {
		e.logger.Error(err)
	}

	bug := new(models.Bug)
	err = e.db.Get(models.Bug{}, c).Model(&models.Bug{}).First(&bug, bugId).Error
	if err != nil || bug.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Bug{}, c).Model(&models.Bug{}).Delete(&bug).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
