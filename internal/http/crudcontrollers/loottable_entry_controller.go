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

type LoottableEntryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoottableEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoottableEntryController {
	return &LoottableEntryController{
		db:	 db,
		logger: logger,
	}
}

func (e *LoottableEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "loottable_entry/:loottableId", e.getLoottableEntry, nil),
		routes.RegisterRoute(http.MethodGet, "loottable_entries", e.listLoottableEntries, nil),
		routes.RegisterRoute(http.MethodPut, "loottable_entry", e.createLoottableEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "loottable_entry/:loottableId", e.deleteLoottableEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "loottable_entry/:loottableId", e.updateLoottableEntry, nil),
		routes.RegisterRoute(http.MethodPost, "loottable_entries/bulk", e.getLoottableEntriesBulk, nil),
	}
}

// listLoottableEntries godoc
// @Id listLoottableEntries
// @Summary Lists LoottableEntries
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Item.DiscoveredItems"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entries [get]
func (e *LoottableEntryController) listLoottableEntries(c echo.Context) error {
	var results []models.LoottableEntry
	err := e.db.QueryContext(models.LoottableEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoottableEntry godoc
// @Id getLoottableEntry
// @Summary Gets LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Item.DiscoveredItems"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoottableEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entry/{id} [get]
func (e *LoottableEntryController) getLoottableEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LoottableId]"})
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.LoottableId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLoottableEntry godoc
// @Id updateLoottableEntry
// @Summary Updates LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "Id"
// @Param loottable_entry body models.LoottableEntry true "LoottableEntry"
// @Success 200 {array} models.LoottableEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /loottable_entry/{id} [patch]
func (e *LoottableEntryController) updateLoottableEntry(c echo.Context) error {
	request := new(models.LoottableEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LoottableId]"})
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
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

// createLoottableEntry godoc
// @Id createLoottableEntry
// @Summary Creates LoottableEntry
// @Accept json
// @Produce json
// @Param loottable_entry body models.LoottableEntry true "LoottableEntry"
// @Tags LoottableEntry
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /loottable_entry [put]
func (e *LoottableEntryController) createLoottableEntry(c echo.Context) error {
	loottableEntry := new(models.LoottableEntry)
	if err := c.Bind(loottableEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LoottableEntry{}, c).Model(&models.LoottableEntry{}).Create(&loottableEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, loottableEntry)
}

// deleteLoottableEntry godoc
// @Id deleteLoottableEntry
// @Summary Deletes LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "loottableId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /loottable_entry/{id} [delete]
func (e *LoottableEntryController) deleteLoottableEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.LoottableEntry{}, c).Model(&models.LoottableEntry{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLoottableEntriesBulk godoc
// @Id getLoottableEntriesBulk
// @Summary Gets LoottableEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoottableEntry
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entries/bulk [post]
func (e *LoottableEntryController) getLoottableEntriesBulk(c echo.Context) error {
	var results []models.LoottableEntry

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

	err := e.db.QueryContext(models.LoottableEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
