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

type LoottableController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoottableController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoottableController {
	return &LoottableController{
		db:	 db,
		logger: logger,
	}
}

func (e *LoottableController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "loottable/:id", e.getLoottable, nil),
		routes.RegisterRoute(http.MethodGet, "loottables", e.listLoottables, nil),
		routes.RegisterRoute(http.MethodPut, "loottable", e.createLoottable, nil),
		routes.RegisterRoute(http.MethodDelete, "loottable/:id", e.deleteLoottable, nil),
		routes.RegisterRoute(http.MethodPatch, "loottable/:id", e.updateLoottable, nil),
		routes.RegisterRoute(http.MethodPost, "loottables/bulk", e.getLoottablesBulk, nil),
	}
}

// listLoottables godoc
// @Id listLoottables
// @Summary Lists Loottables
// @Accept json
// @Produce json
// @Tags Loottable
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LoottableEntries<br>LoottableEntries.LootdropEntries<br>LoottableEntries.LootdropEntries.Item<br>LoottableEntries.LootdropEntries.Item.DiscoveredItems"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Loottable
// @Failure 500 {string} string "Bad query request"
// @Router /loottables [get]
func (e *LoottableController) listLoottables(c echo.Context) error {
	var results []models.Loottable
	err := e.db.QueryContext(models.Loottable{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoottable godoc
// @Id getLoottable
// @Summary Gets Loottable
// @Accept json
// @Produce json
// @Tags Loottable
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LoottableEntries<br>LoottableEntries.LootdropEntries<br>LoottableEntries.LootdropEntries.Item<br>LoottableEntries.LootdropEntries.Item.DiscoveredItems"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Loottable
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /loottable/{id} [get]
func (e *LoottableController) getLoottable(c echo.Context) error {
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
	var result models.Loottable
	query := e.db.QueryContext(models.Loottable{}, c)
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

// updateLoottable godoc
// @Id updateLoottable
// @Summary Updates Loottable
// @Accept json
// @Produce json
// @Tags Loottable
// @Param id path int true "Id"
// @Param loottable body models.Loottable true "Loottable"
// @Success 200 {array} models.Loottable
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /loottable/{id} [patch]
func (e *LoottableController) updateLoottable(c echo.Context) error {
	request := new(models.Loottable)
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
	var result models.Loottable
	query := e.db.QueryContext(models.Loottable{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createLoottable godoc
// @Id createLoottable
// @Summary Creates Loottable
// @Accept json
// @Produce json
// @Param loottable body models.Loottable true "Loottable"
// @Tags Loottable
// @Success 200 {array} models.Loottable
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /loottable [put]
func (e *LoottableController) createLoottable(c echo.Context) error {
	loottable := new(models.Loottable)
	if err := c.Bind(loottable); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Loottable{}, c).Model(&models.Loottable{}).Create(&loottable).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, loottable)
}

// deleteLoottable godoc
// @Id deleteLoottable
// @Summary Deletes Loottable
// @Accept json
// @Produce json
// @Tags Loottable
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /loottable/{id} [delete]
func (e *LoottableController) deleteLoottable(c echo.Context) error {
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
	var result models.Loottable
	query := e.db.QueryContext(models.Loottable{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.Loottable{}, c).Model(&models.Loottable{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLoottablesBulk godoc
// @Id getLoottablesBulk
// @Summary Gets Loottables in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Loottable
// @Success 200 {array} models.Loottable
// @Failure 500 {string} string "Bad query request"
// @Router /loottables/bulk [post]
func (e *LoottableController) getLoottablesBulk(c echo.Context) error {
	var results []models.Loottable

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

	err := e.db.QueryContext(models.Loottable{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
