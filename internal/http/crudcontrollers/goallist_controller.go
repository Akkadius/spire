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

type GoallistController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGoallistController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GoallistController {
	return &GoallistController{
		db:	 db,
		logger: logger,
	}
}

func (e *GoallistController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "goallist/:listid", e.getGoallist, nil),
		routes.RegisterRoute(http.MethodGet, "goallists", e.listGoallists, nil),
		routes.RegisterRoute(http.MethodPut, "goallist", e.createGoallist, nil),
		routes.RegisterRoute(http.MethodDelete, "goallist/:listid", e.deleteGoallist, nil),
		routes.RegisterRoute(http.MethodPatch, "goallist/:listid", e.updateGoallist, nil),
		routes.RegisterRoute(http.MethodPost, "goallists/bulk", e.getGoallistsBulk, nil),
	}
}

// listGoallists godoc
// @Id listGoallists
// @Summary Lists Goallists
// @Accept json
// @Produce json
// @Tags Goallist
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Goallist
// @Failure 500 {string} string "Bad query request"
// @Router /goallists [get]
func (e *GoallistController) listGoallists(c echo.Context) error {
	var results []models.Goallist
	err := e.db.QueryContext(models.Goallist{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGoallist godoc
// @Id getGoallist
// @Summary Gets Goallist
// @Accept json
// @Produce json
// @Tags Goallist
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Goallist
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /goallist/{id} [get]
func (e *GoallistController) getGoallist(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	listid, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Listid]"})
	}
	params = append(params, listid)
	keys = append(keys, "listid = ?")

	// key param [entry] position [2] type [int]
	if len(c.QueryParam("entry")) > 0 {
		entryParam, err := strconv.Atoi(c.QueryParam("entry"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [entry] err [%s]", err.Error())})
		}

		params = append(params, entryParam)
		keys = append(keys, "entry = ?")
	}

	// query builder
	var result models.Goallist
	query := e.db.QueryContext(models.Goallist{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Listid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGoallist godoc
// @Id updateGoallist
// @Summary Updates Goallist
// @Accept json
// @Produce json
// @Tags Goallist
// @Param id path int true "Id"
// @Param goallist body models.Goallist true "Goallist"
// @Success 200 {array} models.Goallist
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /goallist/{id} [patch]
func (e *GoallistController) updateGoallist(c echo.Context) error {
	request := new(models.Goallist)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	listid, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Listid]"})
	}
	params = append(params, listid)
	keys = append(keys, "listid = ?")

	// key param [entry] position [2] type [int]
	if len(c.QueryParam("entry")) > 0 {
		entryParam, err := strconv.Atoi(c.QueryParam("entry"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [entry] err [%s]", err.Error())})
		}

		params = append(params, entryParam)
		keys = append(keys, "entry = ?")
	}

	// query builder
	var result models.Goallist
	query := e.db.QueryContext(models.Goallist{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Goallist{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createGoallist godoc
// @Id createGoallist
// @Summary Creates Goallist
// @Accept json
// @Produce json
// @Param goallist body models.Goallist true "Goallist"
// @Tags Goallist
// @Success 200 {array} models.Goallist
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /goallist [put]
func (e *GoallistController) createGoallist(c echo.Context) error {
	goallist := new(models.Goallist)
	if err := c.Bind(goallist); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Goallist{}, c).Model(&models.Goallist{}).Create(&goallist).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, goallist)
}

// deleteGoallist godoc
// @Id deleteGoallist
// @Summary Deletes Goallist
// @Accept json
// @Produce json
// @Tags Goallist
// @Param id path int true "listid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /goallist/{id} [delete]
func (e *GoallistController) deleteGoallist(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	listid, err := strconv.Atoi(c.Param("listid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, listid)
	keys = append(keys, "listid = ?")

	// key param [entry] position [2] type [int]
	if len(c.QueryParam("entry")) > 0 {
		entryParam, err := strconv.Atoi(c.QueryParam("entry"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [entry] err [%s]", err.Error())})
		}

		params = append(params, entryParam)
		keys = append(keys, "entry = ?")
	}

	// query builder
	var result models.Goallist
	query := e.db.QueryContext(models.Goallist{}, c)
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

// getGoallistsBulk godoc
// @Id getGoallistsBulk
// @Summary Gets Goallists in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Goallist
// @Success 200 {array} models.Goallist
// @Failure 500 {string} string "Bad query request"
// @Router /goallists/bulk [post]
func (e *GoallistController) getGoallistsBulk(c echo.Context) error {
	var results []models.Goallist

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

	err := e.db.QueryContext(models.Goallist{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
