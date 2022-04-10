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

type GridEntryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGridEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GridEntryController {
	return &GridEntryController{
		db:	 db,
		logger: logger,
	}
}

func (e *GridEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "grid_entry/:gridid", e.getGridEntry, nil),
		routes.RegisterRoute(http.MethodGet, "grid_entries", e.listGridEntries, nil),
		routes.RegisterRoute(http.MethodPut, "grid_entry", e.createGridEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "grid_entry/:gridid", e.deleteGridEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "grid_entry/:gridid", e.updateGridEntry, nil),
		routes.RegisterRoute(http.MethodPost, "grid_entries/bulk", e.getGridEntriesBulk, nil),
	}
}

// listGridEntries godoc
// @Id listGridEntries
// @Summary Lists GridEntries
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GridEntry
// @Failure 500 {string} string "Bad query request"
// @Router /grid_entries [get]
func (e *GridEntryController) listGridEntries(c echo.Context) error {
	var results []models.GridEntry
	err := e.db.QueryContext(models.GridEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGridEntry godoc
// @Id getGridEntry
// @Summary Gets GridEntry
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GridEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /grid_entry/{id} [get]
func (e *GridEntryController) getGridEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	gridid, err := strconv.Atoi(c.Param("gridid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Gridid]"})
	}
	params = append(params, gridid)
	keys = append(keys, "gridid = ?")

	// key param [zoneid] position [2] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [number] position [3] type [int]
	if len(c.QueryParam("number")) > 0 {
		numberParam, err := strconv.Atoi(c.QueryParam("number"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [number] err [%s]", err.Error())})
		}

		params = append(params, numberParam)
		keys = append(keys, "number = ?")
	}

	// query builder
	var result models.GridEntry
	query := e.db.QueryContext(models.GridEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Gridid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGridEntry godoc
// @Id updateGridEntry
// @Summary Updates GridEntry
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param id path int true "Id"
// @Param grid_entry body models.GridEntry true "GridEntry"
// @Success 200 {array} models.GridEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /grid_entry/{id} [patch]
func (e *GridEntryController) updateGridEntry(c echo.Context) error {
	request := new(models.GridEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	gridid, err := strconv.Atoi(c.Param("gridid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Gridid]"})
	}
	params = append(params, gridid)
	keys = append(keys, "gridid = ?")

	// key param [zoneid] position [2] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [number] position [3] type [int]
	if len(c.QueryParam("number")) > 0 {
		numberParam, err := strconv.Atoi(c.QueryParam("number"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [number] err [%s]", err.Error())})
		}

		params = append(params, numberParam)
		keys = append(keys, "number = ?")
	}

	// query builder
	var result models.GridEntry
	query := e.db.QueryContext(models.GridEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.GridEntry{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createGridEntry godoc
// @Id createGridEntry
// @Summary Creates GridEntry
// @Accept json
// @Produce json
// @Param grid_entry body models.GridEntry true "GridEntry"
// @Tags GridEntry
// @Success 200 {array} models.GridEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /grid_entry [put]
func (e *GridEntryController) createGridEntry(c echo.Context) error {
	gridEntry := new(models.GridEntry)
	if err := c.Bind(gridEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.GridEntry{}, c).Model(&models.GridEntry{}).Create(&gridEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, gridEntry)
}

// deleteGridEntry godoc
// @Id deleteGridEntry
// @Summary Deletes GridEntry
// @Accept json
// @Produce json
// @Tags GridEntry
// @Param id path int true "gridid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /grid_entry/{id} [delete]
func (e *GridEntryController) deleteGridEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	gridid, err := strconv.Atoi(c.Param("gridid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, gridid)
	keys = append(keys, "gridid = ?")

	// key param [zoneid] position [2] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [number] position [3] type [int]
	if len(c.QueryParam("number")) > 0 {
		numberParam, err := strconv.Atoi(c.QueryParam("number"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [number] err [%s]", err.Error())})
		}

		params = append(params, numberParam)
		keys = append(keys, "number = ?")
	}

	// query builder
	var result models.GridEntry
	query := e.db.QueryContext(models.GridEntry{}, c)
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

// getGridEntriesBulk godoc
// @Id getGridEntriesBulk
// @Summary Gets GridEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GridEntry
// @Success 200 {array} models.GridEntry
// @Failure 500 {string} string "Bad query request"
// @Router /grid_entries/bulk [post]
func (e *GridEntryController) getGridEntriesBulk(c echo.Context) error {
	var results []models.GridEntry

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

	err := e.db.QueryContext(models.GridEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
