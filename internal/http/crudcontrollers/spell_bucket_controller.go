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

type SpellBucketController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpellBucketController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpellBucketController {
	return &SpellBucketController{
		db:	 db,
		logger: logger,
	}
}

func (e *SpellBucketController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spell_bucket/:spellid", e.getSpellBucket, nil),
		routes.RegisterRoute(http.MethodGet, "spell_buckets", e.listSpellBuckets, nil),
		routes.RegisterRoute(http.MethodPut, "spell_bucket", e.createSpellBucket, nil),
		routes.RegisterRoute(http.MethodDelete, "spell_bucket/:spellid", e.deleteSpellBucket, nil),
		routes.RegisterRoute(http.MethodPatch, "spell_bucket/:spellid", e.updateSpellBucket, nil),
		routes.RegisterRoute(http.MethodPost, "spell_buckets/bulk", e.getSpellBucketsBulk, nil),
	}
}

// listSpellBuckets godoc
// @Id listSpellBuckets
// @Summary Lists SpellBuckets
// @Accept json
// @Produce json
// @Tags SpellBucket
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellBucket
// @Failure 500 {string} string "Bad query request"
// @Router /spell_buckets [get]
func (e *SpellBucketController) listSpellBuckets(c echo.Context) error {
	var results []models.SpellBucket
	err := e.db.QueryContext(models.SpellBucket{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpellBucket godoc
// @Id getSpellBucket
// @Summary Gets SpellBucket
// @Accept json
// @Produce json
// @Tags SpellBucket
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellBucket
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spell_bucket/{id} [get]
func (e *SpellBucketController) getSpellBucket(c echo.Context) error {
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
	var result models.SpellBucket
	query := e.db.QueryContext(models.SpellBucket{}, c)
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

// updateSpellBucket godoc
// @Id updateSpellBucket
// @Summary Updates SpellBucket
// @Accept json
// @Produce json
// @Tags SpellBucket
// @Param id path int true "Id"
// @Param spell_bucket body models.SpellBucket true "SpellBucket"
// @Success 200 {array} models.SpellBucket
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spell_bucket/{id} [patch]
func (e *SpellBucketController) updateSpellBucket(c echo.Context) error {
	request := new(models.SpellBucket)
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
	var result models.SpellBucket
	query := e.db.QueryContext(models.SpellBucket{}, c)
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

// createSpellBucket godoc
// @Id createSpellBucket
// @Summary Creates SpellBucket
// @Accept json
// @Produce json
// @Param spell_bucket body models.SpellBucket true "SpellBucket"
// @Tags SpellBucket
// @Success 200 {array} models.SpellBucket
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spell_bucket [put]
func (e *SpellBucketController) createSpellBucket(c echo.Context) error {
	spellBucket := new(models.SpellBucket)
	if err := c.Bind(spellBucket); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SpellBucket{}, c).Model(&models.SpellBucket{}).Create(&spellBucket).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spellBucket)
}

// deleteSpellBucket godoc
// @Id deleteSpellBucket
// @Summary Deletes SpellBucket
// @Accept json
// @Produce json
// @Tags SpellBucket
// @Param id path int true "spellid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spell_bucket/{id} [delete]
func (e *SpellBucketController) deleteSpellBucket(c echo.Context) error {
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
	var result models.SpellBucket
	query := e.db.QueryContext(models.SpellBucket{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.SpellBucket{}, c).Model(&models.SpellBucket{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpellBucketsBulk godoc
// @Id getSpellBucketsBulk
// @Summary Gets SpellBuckets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpellBucket
// @Success 200 {array} models.SpellBucket
// @Failure 500 {string} string "Bad query request"
// @Router /spell_buckets/bulk [post]
func (e *SpellBucketController) getSpellBucketsBulk(c echo.Context) error {
	var results []models.SpellBucket

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

	err := e.db.QueryContext(models.SpellBucket{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
