package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PlayerTitlesetController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewPlayerTitlesetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *PlayerTitlesetController {
	return &PlayerTitlesetController{
		db:     db,
		logger: logger,
	}
}

func (e *PlayerTitlesetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "player_titleset/:player_titleset", e.deletePlayerTitleset, nil),
		routes.RegisterRoute(http.MethodGet, "player_titleset/:player_titleset", e.getPlayerTitleset, nil),
		routes.RegisterRoute(http.MethodGet, "player_titlesets", e.listPlayerTitlesets, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getPlayerTitlesetsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "player_titleset/:player_titleset", e.updatePlayerTitleset, nil),
		routes.RegisterRoute(http.MethodPut, "player_titleset", e.createPlayerTitleset, nil),
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
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
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
	playerTitlesetId, err := strconv.Atoi(c.Param("player_titleset"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.PlayerTitleset
	err = e.db.QueryContext(models.PlayerTitleset{}, c).First(&result, playerTitlesetId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

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
	playerTitleset := new(models.PlayerTitleset)
	if err := c.Bind(playerTitleset); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.PlayerTitleset{}, c).Model(&models.PlayerTitleset{}).First(&models.PlayerTitleset{}, playerTitleset.ID).Error
	if err != nil || playerTitleset.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.PlayerTitleset{}, c).Model(&models.PlayerTitleset{}).Updates(&playerTitleset).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, playerTitleset)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.PlayerTitleset{}, c).Model(&models.PlayerTitleset{}).Create(&playerTitleset).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
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
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /player_titleset/{id} [delete]
func (e *PlayerTitlesetController) deletePlayerTitleset(c echo.Context) error {
	playerTitlesetId, err := strconv.Atoi(c.Param("player_titleset"))
	if err != nil {
		e.logger.Error(err)
	}

	playerTitleset := new(models.PlayerTitleset)
	err = e.db.Get(models.PlayerTitleset{}, c).Model(&models.PlayerTitleset{}).First(&playerTitleset, playerTitlesetId).Error
	if err != nil || playerTitleset.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.PlayerTitleset{}, c).Model(&models.PlayerTitleset{}).Delete(&playerTitleset).Error
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.PlayerTitleset{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
