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

type SpellsNewController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpellsNewController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpellsNewController {
	return &SpellsNewController{
		db:	 db,
		logger: logger,
	}
}

func (e *SpellsNewController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spells_new/:id", e.getSpellsNew, nil),
		routes.RegisterRoute(http.MethodGet, "spells_news", e.listSpellsNews, nil),
		routes.RegisterRoute(http.MethodPut, "spells_new", e.createSpellsNew, nil),
		routes.RegisterRoute(http.MethodDelete, "spells_new/:id", e.deleteSpellsNew, nil),
		routes.RegisterRoute(http.MethodPatch, "spells_new/:id", e.updateSpellsNew, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getSpellsNewsBulk, nil),
	}
}

// listSpellsNews godoc
// @Id listSpellsNews
// @Summary Lists SpellsNews
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Aura<br>BlockedSpells<br>Damageshieldtypes<br>SpellBuckets<br>SpellGlobals"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Bad query request"
// @Router /spells_news [get]
func (e *SpellsNewController) listSpellsNews(c echo.Context) error {
	var results []models.SpellsNew
	err := e.db.QueryContext(models.SpellsNew{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpellsNew godoc
// @Id getSpellsNew
// @Summary Gets SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Aura<br>BlockedSpells<br>Damageshieldtypes<br>SpellBuckets<br>SpellGlobals"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellsNew
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spells_new/{id} [get]
func (e *SpellsNewController) getSpellsNew(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ID]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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

// updateSpellsNew godoc
// @Id updateSpellsNew
// @Summary Updates SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param ID path int true "ID"
// @Param spells_new body models.SpellsNew true "SpellsNew"
// @Success 200 {array} models.SpellsNew
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spells_new/{id} [patch]
func (e *SpellsNewController) updateSpellsNew(c echo.Context) error {
	request := new(models.SpellsNew)
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ID]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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

// createSpellsNew godoc
// @Id createSpellsNew
// @Summary Creates SpellsNew
// @Accept json
// @Produce json
// @Param spells_new body models.SpellsNew true "SpellsNew"
// @Tags SpellsNew
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spells_new [put]
func (e *SpellsNewController) createSpellsNew(c echo.Context) error {
	spellsNew := new(models.SpellsNew)
	if err := c.Bind(spellsNew); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SpellsNew{}, c).Model(&models.SpellsNew{}).Create(&spellsNew).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spellsNew)
}

// deleteSpellsNew godoc
// @Id deleteSpellsNew
// @Summary Deletes SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spells_new/{id} [delete]
func (e *SpellsNewController) deleteSpellsNew(c echo.Context) error {
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
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.SpellsNew{}, c).Model(&models.SpellsNew{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpellsNewsBulk godoc
// @Id getSpellsNewsBulk
// @Summary Gets SpellsNews in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpellsNew
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Bad query request"
// @Router /spells_news/bulk [post]
func (e *SpellsNewController) getSpellsNewsBulk(c echo.Context) error {
	var results []models.SpellsNew

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

	err := e.db.QueryContext(models.SpellsNew{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
