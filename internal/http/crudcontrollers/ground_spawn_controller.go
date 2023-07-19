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

type GroundSpawnController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewGroundSpawnController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *GroundSpawnController {
	return &GroundSpawnController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *GroundSpawnController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "ground_spawn/:id", e.getGroundSpawn, nil),
		routes.RegisterRoute(http.MethodGet, "ground_spawns", e.listGroundSpawns, nil),
		routes.RegisterRoute(http.MethodGet, "ground_spawns/count", e.getGroundSpawnsCount, nil),
		routes.RegisterRoute(http.MethodPut, "ground_spawn", e.createGroundSpawn, nil),
		routes.RegisterRoute(http.MethodDelete, "ground_spawn/:id", e.deleteGroundSpawn, nil),
		routes.RegisterRoute(http.MethodPatch, "ground_spawn/:id", e.updateGroundSpawn, nil),
		routes.RegisterRoute(http.MethodPost, "ground_spawns/bulk", e.getGroundSpawnsBulk, nil),
	}
}

// listGroundSpawns godoc
// @Id listGroundSpawns
// @Summary Lists GroundSpawns
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroundSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /ground_spawns [get]
func (e *GroundSpawnController) listGroundSpawns(c echo.Context) error {
	var results []models.GroundSpawn
	err := e.db.QueryContext(models.GroundSpawn{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGroundSpawn godoc
// @Id getGroundSpawn
// @Summary Gets GroundSpawn
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroundSpawn
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /ground_spawn/{id} [get]
func (e *GroundSpawnController) getGroundSpawn(c echo.Context) error {
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
	var result models.GroundSpawn
	query := e.db.QueryContext(models.GroundSpawn{}, c)
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

// updateGroundSpawn godoc
// @Id updateGroundSpawn
// @Summary Updates GroundSpawn
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param id path int true "Id"
// @Param ground_spawn body models.GroundSpawn true "GroundSpawn"
// @Success 200 {array} models.GroundSpawn
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /ground_spawn/{id} [patch]
func (e *GroundSpawnController) updateGroundSpawn(c echo.Context) error {
	request := new(models.GroundSpawn)
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
	var result models.GroundSpawn
	query := e.db.QueryContext(models.GroundSpawn{}, c)
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
		event := fmt.Sprintf("Updated [GroundSpawn] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createGroundSpawn godoc
// @Id createGroundSpawn
// @Summary Creates GroundSpawn
// @Accept json
// @Produce json
// @Param ground_spawn body models.GroundSpawn true "GroundSpawn"
// @Tags GroundSpawn
// @Success 200 {array} models.GroundSpawn
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /ground_spawn [put]
func (e *GroundSpawnController) createGroundSpawn(c echo.Context) error {
	groundSpawn := new(models.GroundSpawn)
	if err := c.Bind(groundSpawn); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.GroundSpawn{}, c).Model(&models.GroundSpawn{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
        db = db.Omit(clause.Associations)
    }

	err := db.Create(&groundSpawn).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GroundSpawn{}, groundSpawn)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GroundSpawn] [%v] data [%v]", groundSpawn.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, groundSpawn)
}

// deleteGroundSpawn godoc
// @Id deleteGroundSpawn
// @Summary Deletes GroundSpawn
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /ground_spawn/{id} [delete]
func (e *GroundSpawnController) deleteGroundSpawn(c echo.Context) error {
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
	var result models.GroundSpawn
	query := e.db.QueryContext(models.GroundSpawn{}, c)
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
		event := fmt.Sprintf("Deleted [GroundSpawn] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGroundSpawnsBulk godoc
// @Id getGroundSpawnsBulk
// @Summary Gets GroundSpawns in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GroundSpawn
// @Success 200 {array} models.GroundSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /ground_spawns/bulk [post]
func (e *GroundSpawnController) getGroundSpawnsBulk(c echo.Context) error {
	var results []models.GroundSpawn

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

	err := e.db.QueryContext(models.GroundSpawn{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGroundSpawnsCount godoc
// @Id getGroundSpawnsCount
// @Summary Counts GroundSpawns
// @Accept json
// @Produce json
// @Tags GroundSpawn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GroundSpawn
// @Failure 500 {string} string "Bad query request"
// @Router /ground_spawns/count [get]
func (e *GroundSpawnController) getGroundSpawnsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.GroundSpawn{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}