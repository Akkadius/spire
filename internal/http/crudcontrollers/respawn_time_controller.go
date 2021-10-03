package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RespawnTimeController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewRespawnTimeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *RespawnTimeController {
	return &RespawnTimeController{
		db:     db,
		logger: logger,
	}
}

func (e *RespawnTimeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "respawn_time/:respawn_time", e.deleteRespawnTime, nil),
		routes.RegisterRoute(http.MethodGet, "respawn_time/:respawn_time", e.getRespawnTime, nil),
		routes.RegisterRoute(http.MethodGet, "respawn_times", e.listRespawnTimes, nil),
		routes.RegisterRoute(http.MethodPost, "respawn_times/bulk", e.getRespawnTimesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "respawn_time/:respawn_time", e.updateRespawnTime, nil),
		routes.RegisterRoute(http.MethodPut, "respawn_time", e.createRespawnTime, nil),
	}
}

// listRespawnTimes godoc
// @Id listRespawnTimes
// @Summary Lists RespawnTimes
// @Accept json
// @Produce json
// @Tags RespawnTime
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RespawnTime
// @Failure 500 {string} string "Bad query request"
// @Router /respawn_times [get]
func (e *RespawnTimeController) listRespawnTimes(c echo.Context) error {
	var results []models.RespawnTime
	err := e.db.QueryContext(models.RespawnTime{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getRespawnTime godoc
// @Id getRespawnTime
// @Summary Gets RespawnTime
// @Accept json
// @Produce json
// @Tags RespawnTime
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RespawnTime
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /respawn_time/{id} [get]
func (e *RespawnTimeController) getRespawnTime(c echo.Context) error {
	respawnTimeId, err := strconv.Atoi(c.Param("respawn_time"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.RespawnTime
	err = e.db.QueryContext(models.RespawnTime{}, c).First(&result, respawnTimeId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateRespawnTime godoc
// @Id updateRespawnTime
// @Summary Updates RespawnTime
// @Accept json
// @Produce json
// @Tags RespawnTime
// @Param id path int true "Id"
// @Param respawn_time body models.RespawnTime true "RespawnTime"
// @Success 200 {array} models.RespawnTime
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /respawn_time/{id} [patch]
func (e *RespawnTimeController) updateRespawnTime(c echo.Context) error {
	respawnTime := new(models.RespawnTime)
	if err := c.Bind(respawnTime); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.RespawnTime{}
	err := e.db.Get(models.RespawnTime{}, c).Model(&models.RespawnTime{}).First(&entity, respawnTime.ID).Error
	if err != nil || respawnTime.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.RespawnTime{}, c).Model(&entity).Updates(&respawnTime).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, respawnTime)
}

// createRespawnTime godoc
// @Id createRespawnTime
// @Summary Creates RespawnTime
// @Accept json
// @Produce json
// @Param respawn_time body models.RespawnTime true "RespawnTime"
// @Tags RespawnTime
// @Success 200 {array} models.RespawnTime
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /respawn_time [put]
func (e *RespawnTimeController) createRespawnTime(c echo.Context) error {
	respawnTime := new(models.RespawnTime)
	if err := c.Bind(respawnTime); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.RespawnTime{}, c).Model(&models.RespawnTime{}).Create(&respawnTime).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, respawnTime)
}

// deleteRespawnTime godoc
// @Id deleteRespawnTime
// @Summary Deletes RespawnTime
// @Accept json
// @Produce json
// @Tags RespawnTime
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /respawn_time/{id} [delete]
func (e *RespawnTimeController) deleteRespawnTime(c echo.Context) error {
	respawnTimeId, err := strconv.Atoi(c.Param("respawn_time"))
	if err != nil {
		e.logger.Error(err)
	}

	respawnTime := new(models.RespawnTime)
	err = e.db.Get(models.RespawnTime{}, c).Model(&models.RespawnTime{}).First(&respawnTime, respawnTimeId).Error
	if err != nil || respawnTime.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.RespawnTime{}, c).Model(&models.RespawnTime{}).Delete(&respawnTime).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getRespawnTimesBulk godoc
// @Id getRespawnTimesBulk
// @Summary Gets RespawnTimes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags RespawnTime
// @Success 200 {array} models.RespawnTime
// @Failure 500 {string} string "Bad query request"
// @Router /respawn_times/bulk [post]
func (e *RespawnTimeController) getRespawnTimesBulk(c echo.Context) error {
	var results []models.RespawnTime

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.RespawnTime{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
