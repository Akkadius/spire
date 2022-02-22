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

type SpellGlobalController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpellGlobalController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpellGlobalController {
	return &SpellGlobalController{
		db:	 db,
		logger: logger,
	}
}

func (e *SpellGlobalController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spell_global/:spellid", e.getSpellGlobal, nil),
		routes.RegisterRoute(http.MethodGet, "spell_globals", e.listSpellGlobals, nil),
		routes.RegisterRoute(http.MethodPut, "spell_global", e.createSpellGlobal, nil),
		routes.RegisterRoute(http.MethodDelete, "spell_global/:spellid", e.deleteSpellGlobal, nil),
		routes.RegisterRoute(http.MethodPatch, "spell_global/:spellid", e.updateSpellGlobal, nil),
		routes.RegisterRoute(http.MethodPost, "spell_globals/bulk", e.getSpellGlobalsBulk, nil),
	}
}

// listSpellGlobals godoc
// @Id listSpellGlobals
// @Summary Lists SpellGlobals
// @Accept json
// @Produce json
// @Tags SpellGlobal
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellGlobal
// @Failure 500 {string} string "Bad query request"
// @Router /spell_globals [get]
func (e *SpellGlobalController) listSpellGlobals(c echo.Context) error {
	var results []models.SpellGlobal
	err := e.db.QueryContext(models.SpellGlobal{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpellGlobal godoc
// @Id getSpellGlobal
// @Summary Gets SpellGlobal
// @Accept json
// @Produce json
// @Tags SpellGlobal
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellGlobal
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spell_global/{id} [get]
func (e *SpellGlobalController) getSpellGlobal(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	spellid, err := strconv.Atoi(c.Param("spellid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Spellid]"})
	}
	params = append(params, spellid)
	keys = append(keys, "spellid = ?")

	// query builder
	var result models.SpellGlobal
	query := e.db.QueryContext(models.SpellGlobal{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Spellid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSpellGlobal godoc
// @Id updateSpellGlobal
// @Summary Updates SpellGlobal
// @Accept json
// @Produce json
// @Tags SpellGlobal
// @Param id path int true "Id"
// @Param spell_global body models.SpellGlobal true "SpellGlobal"
// @Success 200 {array} models.SpellGlobal
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spell_global/{id} [patch]
func (e *SpellGlobalController) updateSpellGlobal(c echo.Context) error {
	request := new(models.SpellGlobal)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	spellid, err := strconv.Atoi(c.Param("spellid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Spellid]"})
	}
	params = append(params, spellid)
	keys = append(keys, "spellid = ?")

	// query builder
	var result models.SpellGlobal
	query := e.db.QueryContext(models.SpellGlobal{}, c)
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

// createSpellGlobal godoc
// @Id createSpellGlobal
// @Summary Creates SpellGlobal
// @Accept json
// @Produce json
// @Param spell_global body models.SpellGlobal true "SpellGlobal"
// @Tags SpellGlobal
// @Success 200 {array} models.SpellGlobal
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spell_global [put]
func (e *SpellGlobalController) createSpellGlobal(c echo.Context) error {
	spellGlobal := new(models.SpellGlobal)
	if err := c.Bind(spellGlobal); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SpellGlobal{}, c).Model(&models.SpellGlobal{}).Create(&spellGlobal).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spellGlobal)
}

// deleteSpellGlobal godoc
// @Id deleteSpellGlobal
// @Summary Deletes SpellGlobal
// @Accept json
// @Produce json
// @Tags SpellGlobal
// @Param id path int true "spellid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spell_global/{id} [delete]
func (e *SpellGlobalController) deleteSpellGlobal(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	spellid, err := strconv.Atoi(c.Param("spellid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, spellid)
	keys = append(keys, "spellid = ?")

	// query builder
	var result models.SpellGlobal
	query := e.db.QueryContext(models.SpellGlobal{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.SpellGlobal{}, c).Model(&models.SpellGlobal{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpellGlobalsBulk godoc
// @Id getSpellGlobalsBulk
// @Summary Gets SpellGlobals in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpellGlobal
// @Success 200 {array} models.SpellGlobal
// @Failure 500 {string} string "Bad query request"
// @Router /spell_globals/bulk [post]
func (e *SpellGlobalController) getSpellGlobalsBulk(c echo.Context) error {
	var results []models.SpellGlobal

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

	err := e.db.QueryContext(models.SpellGlobal{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
