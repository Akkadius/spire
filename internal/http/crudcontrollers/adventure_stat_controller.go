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

type AdventureStatController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAdventureStatController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AdventureStatController {
	return &AdventureStatController{
		db:	 db,
		logger: logger,
	}
}

func (e *AdventureStatController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "adventure_stat/:playerId", e.getAdventureStat, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_stats", e.listAdventureStats, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_stat", e.createAdventureStat, nil),
		routes.RegisterRoute(http.MethodDelete, "adventure_stat/:playerId", e.deleteAdventureStat, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_stat/:playerId", e.updateAdventureStat, nil),
		routes.RegisterRoute(http.MethodPost, "adventure_stats/bulk", e.getAdventureStatsBulk, nil),
	}
}

// listAdventureStats godoc
// @Id listAdventureStats
// @Summary Lists AdventureStats
// @Accept json
// @Produce json
// @Tags AdventureStat
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureStat
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_stats [get]
func (e *AdventureStatController) listAdventureStats(c echo.Context) error {
	var results []models.AdventureStat
	err := e.db.QueryContext(models.AdventureStat{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureStat godoc
// @Id getAdventureStat
// @Summary Gets AdventureStat
// @Accept json
// @Produce json
// @Tags AdventureStat
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureStat
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_stat/{id} [get]
func (e *AdventureStatController) getAdventureStat(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PlayerId]"})
	}
	params = append(params, playerId)
	keys = append(keys, "player_id = ?")

	// query builder
	var result models.AdventureStat
	query := e.db.QueryContext(models.AdventureStat{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.PlayerId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAdventureStat godoc
// @Id updateAdventureStat
// @Summary Updates AdventureStat
// @Accept json
// @Produce json
// @Tags AdventureStat
// @Param id path int true "Id"
// @Param adventure_stat body models.AdventureStat true "AdventureStat"
// @Success 200 {array} models.AdventureStat
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_stat/{id} [patch]
func (e *AdventureStatController) updateAdventureStat(c echo.Context) error {
	request := new(models.AdventureStat)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PlayerId]"})
	}
	params = append(params, playerId)
	keys = append(keys, "player_id = ?")

	// query builder
	var result models.AdventureStat
	query := e.db.QueryContext(models.AdventureStat{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AdventureStat{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAdventureStat godoc
// @Id createAdventureStat
// @Summary Creates AdventureStat
// @Accept json
// @Produce json
// @Param adventure_stat body models.AdventureStat true "AdventureStat"
// @Tags AdventureStat
// @Success 200 {array} models.AdventureStat
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_stat [put]
func (e *AdventureStatController) createAdventureStat(c echo.Context) error {
	adventureStat := new(models.AdventureStat)
	if err := c.Bind(adventureStat); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AdventureStat{}, c).Model(&models.AdventureStat{}).Create(&adventureStat).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, adventureStat)
}

// deleteAdventureStat godoc
// @Id deleteAdventureStat
// @Summary Deletes AdventureStat
// @Accept json
// @Produce json
// @Tags AdventureStat
// @Param id path int true "playerId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_stat/{id} [delete]
func (e *AdventureStatController) deleteAdventureStat(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	playerId, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, playerId)
	keys = append(keys, "player_id = ?")

	// query builder
	var result models.AdventureStat
	query := e.db.QueryContext(models.AdventureStat{}, c)
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

// getAdventureStatsBulk godoc
// @Id getAdventureStatsBulk
// @Summary Gets AdventureStats in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AdventureStat
// @Success 200 {array} models.AdventureStat
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_stats/bulk [post]
func (e *AdventureStatController) getAdventureStatsBulk(c echo.Context) error {
	var results []models.AdventureStat

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

	err := e.db.QueryContext(models.AdventureStat{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
