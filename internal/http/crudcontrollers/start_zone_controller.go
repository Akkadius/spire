package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type StartZoneController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewStartZoneController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *StartZoneController {
	return &StartZoneController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *StartZoneController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "start_zone/:playerChoice", e.getStartZone, nil),
		routes.RegisterRoute(http.MethodGet, "start_zones", e.listStartZones, nil),
		routes.RegisterRoute(http.MethodGet, "start_zones/count", e.getStartZonesCount, nil),
		routes.RegisterRoute(http.MethodPut, "start_zone", e.createStartZone, nil),
		routes.RegisterRoute(http.MethodDelete, "start_zone/:playerChoice", e.deleteStartZone, nil),
		routes.RegisterRoute(http.MethodPatch, "start_zone/:playerChoice", e.updateStartZone, nil),
		routes.RegisterRoute(http.MethodPost, "start_zones/bulk", e.getStartZonesBulk, nil),
	}
}

// listStartZones godoc
// @Id listStartZones
// @Summary Lists StartZones
// @Accept json
// @Produce json
// @Tags StartZone
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.StartZone
// @Failure 500 {string} string "Bad query request"
// @Router /start_zones [get]
func (e *StartZoneController) listStartZones(c echo.Context) error {
	var results []models.StartZone
	err := e.db.QueryContext(models.StartZone{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getStartZone godoc
// @Id getStartZone
// @Summary Gets StartZone
// @Accept json
// @Produce json
// @Tags StartZone
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.StartZone
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /start_zone/{id} [get]
func (e *StartZoneController) getStartZone(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	playerChoice, err := strconv.Atoi(c.Param("playerChoice"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PlayerChoice]"})
	}
	params = append(params, playerChoice)
	keys = append(keys, "player_choice = ?")

	// key param [player_class] position [8] type [int]
	if len(c.QueryParam("player_class")) > 0 {
		playerClassParam, err := strconv.Atoi(c.QueryParam("player_class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_class] err [%s]", err.Error())})
		}

		params = append(params, playerClassParam)
		keys = append(keys, "player_class = ?")
	}

	// key param [player_deity] position [9] type [int]
	if len(c.QueryParam("player_deity")) > 0 {
		playerDeityParam, err := strconv.Atoi(c.QueryParam("player_deity"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_deity] err [%s]", err.Error())})
		}

		params = append(params, playerDeityParam)
		keys = append(keys, "player_deity = ?")
	}

	// key param [player_race] position [10] type [int]
	if len(c.QueryParam("player_race")) > 0 {
		playerRaceParam, err := strconv.Atoi(c.QueryParam("player_race"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_race] err [%s]", err.Error())})
		}

		params = append(params, playerRaceParam)
		keys = append(keys, "player_race = ?")
	}

	// query builder
	var result models.StartZone
	query := e.db.QueryContext(models.StartZone{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.PlayerChoice == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateStartZone godoc
// @Id updateStartZone
// @Summary Updates StartZone
// @Accept json
// @Produce json
// @Tags StartZone
// @Param id path int true "Id"
// @Param start_zone body models.StartZone true "StartZone"
// @Success 200 {array} models.StartZone
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /start_zone/{id} [patch]
func (e *StartZoneController) updateStartZone(c echo.Context) error {
	request := new(models.StartZone)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	playerChoice, err := strconv.Atoi(c.Param("playerChoice"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PlayerChoice]"})
	}
	params = append(params, playerChoice)
	keys = append(keys, "player_choice = ?")

	// key param [player_class] position [8] type [int]
	if len(c.QueryParam("player_class")) > 0 {
		playerClassParam, err := strconv.Atoi(c.QueryParam("player_class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_class] err [%s]", err.Error())})
		}

		params = append(params, playerClassParam)
		keys = append(keys, "player_class = ?")
	}

	// key param [player_deity] position [9] type [int]
	if len(c.QueryParam("player_deity")) > 0 {
		playerDeityParam, err := strconv.Atoi(c.QueryParam("player_deity"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_deity] err [%s]", err.Error())})
		}

		params = append(params, playerDeityParam)
		keys = append(keys, "player_deity = ?")
	}

	// key param [player_race] position [10] type [int]
	if len(c.QueryParam("player_race")) > 0 {
		playerRaceParam, err := strconv.Atoi(c.QueryParam("player_race"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_race] err [%s]", err.Error())})
		}

		params = append(params, playerRaceParam)
		keys = append(keys, "player_race = ?")
	}

	// query builder
	var result models.StartZone
	query := e.db.QueryContext(models.StartZone{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [StartZone] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createStartZone godoc
// @Id createStartZone
// @Summary Creates StartZone
// @Accept json
// @Produce json
// @Param start_zone body models.StartZone true "StartZone"
// @Tags StartZone
// @Success 200 {array} models.StartZone
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /start_zone [put]
func (e *StartZoneController) createStartZone(c echo.Context) error {
	startZone := new(models.StartZone)
	if err := c.Bind(startZone); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.StartZone{}, c).Model(&models.StartZone{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
        db = db.Omit(clause.Associations)
    }

	err := db.Create(&startZone).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.StartZone{}, startZone)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [StartZone] [%v] data [%v]", startZone.PlayerChoice, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, startZone)
}

// deleteStartZone godoc
// @Id deleteStartZone
// @Summary Deletes StartZone
// @Accept json
// @Produce json
// @Tags StartZone
// @Param id path int true "playerChoice"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /start_zone/{id} [delete]
func (e *StartZoneController) deleteStartZone(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	playerChoice, err := strconv.Atoi(c.Param("playerChoice"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, playerChoice)
	keys = append(keys, "player_choice = ?")

	// key param [player_class] position [8] type [int]
	if len(c.QueryParam("player_class")) > 0 {
		playerClassParam, err := strconv.Atoi(c.QueryParam("player_class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_class] err [%s]", err.Error())})
		}

		params = append(params, playerClassParam)
		keys = append(keys, "player_class = ?")
	}

	// key param [player_deity] position [9] type [int]
	if len(c.QueryParam("player_deity")) > 0 {
		playerDeityParam, err := strconv.Atoi(c.QueryParam("player_deity"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_deity] err [%s]", err.Error())})
		}

		params = append(params, playerDeityParam)
		keys = append(keys, "player_deity = ?")
	}

	// key param [player_race] position [10] type [int]
	if len(c.QueryParam("player_race")) > 0 {
		playerRaceParam, err := strconv.Atoi(c.QueryParam("player_race"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [player_race] err [%s]", err.Error())})
		}

		params = append(params, playerRaceParam)
		keys = append(keys, "player_race = ?")
	}

	// query builder
	var result models.StartZone
	query := e.db.QueryContext(models.StartZone{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [StartZone] [%v] keys [%v]", result.PlayerChoice, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getStartZonesBulk godoc
// @Id getStartZonesBulk
// @Summary Gets StartZones in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags StartZone
// @Success 200 {array} models.StartZone
// @Failure 500 {string} string "Bad query request"
// @Router /start_zones/bulk [post]
func (e *StartZoneController) getStartZonesBulk(c echo.Context) error {
	var results []models.StartZone

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

	err := e.db.QueryContext(models.StartZone{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getStartZonesCount godoc
// @Id getStartZonesCount
// @Summary Counts StartZones
// @Accept json
// @Produce json
// @Tags StartZone
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.StartZone
// @Failure 500 {string} string "Bad query request"
// @Router /start_zones/count [get]
func (e *StartZoneController) getStartZonesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.StartZone{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}