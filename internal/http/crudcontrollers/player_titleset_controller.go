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

type PlayerTitlesetController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewPlayerTitlesetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *PlayerTitlesetController {
	return &PlayerTitlesetController{
		db:	 db,
		logger: logger,
	}
}

func (e *PlayerTitlesetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "player_titleset/:id", e.getPlayerTitleset, nil),
		routes.RegisterRoute(http.MethodGet, "player_titlesets", e.listPlayerTitlesets, nil),
		routes.RegisterRoute(http.MethodPut, "player_titleset", e.createPlayerTitleset, nil),
		routes.RegisterRoute(http.MethodDelete, "player_titleset/:id", e.deletePlayerTitleset, nil),
		routes.RegisterRoute(http.MethodPatch, "player_titleset/:id", e.updatePlayerTitleset, nil),
		routes.RegisterRoute(http.MethodPost, "player_titlesets/bulk", e.getPlayerTitlesetsBulk, nil),
	}
}

// listPlayerTitlesets godoc
// @Id listPlayerTitlesets
// @Summary Lists PlayerTitlesets
// @Accept json
// @Produce json
// @Tags PlayerTitleset
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerTitleset
// @Failure 500 {string} string "Bad query request"
// @Router /player_titlesets [get]
func (e *PlayerTitlesetController) listPlayerTitlesets(c echo.Context) error {
	var results []models.PlayerTitleset
	err := e.db.QueryContext(models.PlayerTitleset{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getPlayerTitleset godoc
// @Id getPlayerTitleset
// @Summary Gets PlayerTitleset
// @Accept json
// @Produce json
// @Tags PlayerTitleset
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PlayerTitleset
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /player_titleset/{id} [get]
func (e *PlayerTitlesetController) getPlayerTitleset(c echo.Context) error {
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
	var result models.PlayerTitleset
	query := e.db.QueryContext(models.PlayerTitleset{}, c)
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

// updatePlayerTitleset godoc
// @Id updatePlayerTitleset
// @Summary Updates PlayerTitleset
// @Accept json
// @Produce json
// @Tags PlayerTitleset
// @Param id path int true "Id"
// @Param player_titleset body models.PlayerTitleset true "PlayerTitleset"
// @Success 200 {array} models.PlayerTitleset
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /player_titleset/{id} [patch]
func (e *PlayerTitlesetController) updatePlayerTitleset(c echo.Context) error {
	request := new(models.PlayerTitleset)
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
	var result models.PlayerTitleset
	query := e.db.QueryContext(models.PlayerTitleset{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.PlayerTitleset{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createPlayerTitleset godoc
// @Id createPlayerTitleset
// @Summary Creates PlayerTitleset
// @Accept json
// @Produce json
// @Param player_titleset body models.PlayerTitleset true "PlayerTitleset"
// @Tags PlayerTitleset
// @Success 200 {array} models.PlayerTitleset
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /player_titleset [put]
func (e *PlayerTitlesetController) createPlayerTitleset(c echo.Context) error {
	playerTitleset := new(models.PlayerTitleset)
	if err := c.Bind(playerTitleset); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.PlayerTitleset{}, c).Model(&models.PlayerTitleset{}).Create(&playerTitleset).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, playerTitleset)
}

// deletePlayerTitleset godoc
// @Id deletePlayerTitleset
// @Summary Deletes PlayerTitleset
// @Accept json
// @Produce json
// @Tags PlayerTitleset
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /player_titleset/{id} [delete]
func (e *PlayerTitlesetController) deletePlayerTitleset(c echo.Context) error {
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
	var result models.PlayerTitleset
	query := e.db.QueryContext(models.PlayerTitleset{}, c)
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

// getPlayerTitlesetsBulk godoc
// @Id getPlayerTitlesetsBulk
// @Summary Gets PlayerTitlesets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PlayerTitleset
// @Success 200 {array} models.PlayerTitleset
// @Failure 500 {string} string "Bad query request"
// @Router /player_titlesets/bulk [post]
func (e *PlayerTitlesetController) getPlayerTitlesetsBulk(c echo.Context) error {
	var results []models.PlayerTitleset

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

	err := e.db.QueryContext(models.PlayerTitleset{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
