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

type LdonTrapEntryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLdonTrapEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LdonTrapEntryController {
	return &LdonTrapEntryController{
		db:	 db,
		logger: logger,
	}
}

func (e *LdonTrapEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "ldon_trap_entry/:id", e.getLdonTrapEntry, nil),
		routes.RegisterRoute(http.MethodGet, "ldon_trap_entries", e.listLdonTrapEntries, nil),
		routes.RegisterRoute(http.MethodPut, "ldon_trap_entry", e.createLdonTrapEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "ldon_trap_entry/:id", e.deleteLdonTrapEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "ldon_trap_entry/:id", e.updateLdonTrapEntry, nil),
		routes.RegisterRoute(http.MethodPost, "ldon_trap_entries/bulk", e.getLdonTrapEntriesBulk, nil),
	}
}

// listLdonTrapEntries godoc
// @Id listLdonTrapEntries
// @Summary Lists LdonTrapEntries
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapEntry
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_entries [get]
func (e *LdonTrapEntryController) listLdonTrapEntries(c echo.Context) error {
	var results []models.LdonTrapEntry
	err := e.db.QueryContext(models.LdonTrapEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLdonTrapEntry godoc
// @Id getLdonTrapEntry
// @Summary Gets LdonTrapEntry
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LdonTrapEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_entry/{id} [get]
func (e *LdonTrapEntryController) getLdonTrapEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [trap_id] position [2] type [int]
	if len(c.QueryParam("trap_id")) > 0 {
		trapIdParam, err := strconv.Atoi(c.QueryParam("trap_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [trap_id] err [%s]", err.Error())})
		}

		params = append(params, trapIdParam)
		keys = append(keys, "trap_id = ?")
	}

	// query builder
	var result models.LdonTrapEntry
	query := e.db.QueryContext(models.LdonTrapEntry{}, c)
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

// updateLdonTrapEntry godoc
// @Id updateLdonTrapEntry
// @Summary Updates LdonTrapEntry
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param id path int true "Id"
// @Param ldon_trap_entry body models.LdonTrapEntry true "LdonTrapEntry"
// @Success 200 {array} models.LdonTrapEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /ldon_trap_entry/{id} [patch]
func (e *LdonTrapEntryController) updateLdonTrapEntry(c echo.Context) error {
	request := new(models.LdonTrapEntry)
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

	// key param [trap_id] position [2] type [int]
	if len(c.QueryParam("trap_id")) > 0 {
		trapIdParam, err := strconv.Atoi(c.QueryParam("trap_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [trap_id] err [%s]", err.Error())})
		}

		params = append(params, trapIdParam)
		keys = append(keys, "trap_id = ?")
	}

	// query builder
	var result models.LdonTrapEntry
	query := e.db.QueryContext(models.LdonTrapEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.LdonTrapEntry{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createLdonTrapEntry godoc
// @Id createLdonTrapEntry
// @Summary Creates LdonTrapEntry
// @Accept json
// @Produce json
// @Param ldon_trap_entry body models.LdonTrapEntry true "LdonTrapEntry"
// @Tags LdonTrapEntry
// @Success 200 {array} models.LdonTrapEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /ldon_trap_entry [put]
func (e *LdonTrapEntryController) createLdonTrapEntry(c echo.Context) error {
	ldonTrapEntry := new(models.LdonTrapEntry)
	if err := c.Bind(ldonTrapEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LdonTrapEntry{}, c).Model(&models.LdonTrapEntry{}).Create(&ldonTrapEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, ldonTrapEntry)
}

// deleteLdonTrapEntry godoc
// @Id deleteLdonTrapEntry
// @Summary Deletes LdonTrapEntry
// @Accept json
// @Produce json
// @Tags LdonTrapEntry
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /ldon_trap_entry/{id} [delete]
func (e *LdonTrapEntryController) deleteLdonTrapEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [trap_id] position [2] type [int]
	if len(c.QueryParam("trap_id")) > 0 {
		trapIdParam, err := strconv.Atoi(c.QueryParam("trap_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [trap_id] err [%s]", err.Error())})
		}

		params = append(params, trapIdParam)
		keys = append(keys, "trap_id = ?")
	}

	// query builder
	var result models.LdonTrapEntry
	query := e.db.QueryContext(models.LdonTrapEntry{}, c)
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

// getLdonTrapEntriesBulk godoc
// @Id getLdonTrapEntriesBulk
// @Summary Gets LdonTrapEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LdonTrapEntry
// @Success 200 {array} models.LdonTrapEntry
// @Failure 500 {string} string "Bad query request"
// @Router /ldon_trap_entries/bulk [post]
func (e *LdonTrapEntryController) getLdonTrapEntriesBulk(c echo.Context) error {
	var results []models.LdonTrapEntry

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

	err := e.db.QueryContext(models.LdonTrapEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
