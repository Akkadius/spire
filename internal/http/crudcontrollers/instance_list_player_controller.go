package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type InstanceListPlayerController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewInstanceListPlayerController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *InstanceListPlayerController {
	return &InstanceListPlayerController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *InstanceListPlayerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "instance_list_player/:id", e.getInstanceListPlayer, nil),
		routes.RegisterRoute(http.MethodGet, "instance_list_players", e.listInstanceListPlayers, nil),
		routes.RegisterRoute(http.MethodGet, "instance_list_players/count", e.getInstanceListPlayersCount, nil),
		routes.RegisterRoute(http.MethodPut, "instance_list_player", e.createInstanceListPlayer, nil),
		routes.RegisterRoute(http.MethodDelete, "instance_list_player/:id", e.deleteInstanceListPlayer, nil),
		routes.RegisterRoute(http.MethodPatch, "instance_list_player/:id", e.updateInstanceListPlayer, nil),
		routes.RegisterRoute(http.MethodPost, "instance_list_players/bulk", e.getInstanceListPlayersBulk, nil),
	}
}

// listInstanceListPlayers godoc
// @Id listInstanceListPlayers
// @Summary Lists InstanceListPlayers
// @Accept json
// @Produce json
// @Tags InstanceListPlayer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InstanceListPlayer
// @Failure 500 {string} string "Bad query request"
// @Router /instance_list_players [get]
func (e *InstanceListPlayerController) listInstanceListPlayers(c echo.Context) error {
	var results []models.InstanceListPlayer
	err := e.db.QueryContext(models.InstanceListPlayer{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getInstanceListPlayer godoc
// @Id getInstanceListPlayer
// @Summary Gets InstanceListPlayer
// @Accept json
// @Produce json
// @Tags InstanceListPlayer
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InstanceListPlayer
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /instance_list_player/{id} [get]
func (e *InstanceListPlayerController) getInstanceListPlayer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// query builder
	var result models.InstanceListPlayer
	query := e.db.QueryContext(models.InstanceListPlayer{}, c)
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

// updateInstanceListPlayer godoc
// @Id updateInstanceListPlayer
// @Summary Updates InstanceListPlayer
// @Accept json
// @Produce json
// @Tags InstanceListPlayer
// @Param id path int true "Id"
// @Param instance_list_player body models.InstanceListPlayer true "InstanceListPlayer"
// @Success 200 {array} models.InstanceListPlayer
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /instance_list_player/{id} [patch]
func (e *InstanceListPlayerController) updateInstanceListPlayer(c echo.Context) error {
	request := new(models.InstanceListPlayer)
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

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// query builder
	var result models.InstanceListPlayer
	query := e.db.QueryContext(models.InstanceListPlayer{}, c)
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
		event := fmt.Sprintf("Updated [InstanceListPlayer] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createInstanceListPlayer godoc
// @Id createInstanceListPlayer
// @Summary Creates InstanceListPlayer
// @Accept json
// @Produce json
// @Param instance_list_player body models.InstanceListPlayer true "InstanceListPlayer"
// @Tags InstanceListPlayer
// @Success 200 {array} models.InstanceListPlayer
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /instance_list_player [put]
func (e *InstanceListPlayerController) createInstanceListPlayer(c echo.Context) error {
	instanceListPlayer := new(models.InstanceListPlayer)
	if err := c.Bind(instanceListPlayer); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.InstanceListPlayer{}, c).Model(&models.InstanceListPlayer{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&instanceListPlayer).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.InstanceListPlayer{}, instanceListPlayer)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [InstanceListPlayer] [%v] data [%v]", instanceListPlayer.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, instanceListPlayer)
}

// deleteInstanceListPlayer godoc
// @Id deleteInstanceListPlayer
// @Summary Deletes InstanceListPlayer
// @Accept json
// @Produce json
// @Tags InstanceListPlayer
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /instance_list_player/{id} [delete]
func (e *InstanceListPlayerController) deleteInstanceListPlayer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [charid] position [2] type [int]
	if len(c.QueryParam("charid")) > 0 {
		charidParam, err := strconv.Atoi(c.QueryParam("charid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [charid] err [%s]", err.Error())})
		}

		params = append(params, charidParam)
		keys = append(keys, "charid = ?")
	}

	// query builder
	var result models.InstanceListPlayer
	query := e.db.QueryContext(models.InstanceListPlayer{}, c)
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
		event := fmt.Sprintf("Deleted [InstanceListPlayer] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getInstanceListPlayersBulk godoc
// @Id getInstanceListPlayersBulk
// @Summary Gets InstanceListPlayers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags InstanceListPlayer
// @Success 200 {array} models.InstanceListPlayer
// @Failure 500 {string} string "Bad query request"
// @Router /instance_list_players/bulk [post]
func (e *InstanceListPlayerController) getInstanceListPlayersBulk(c echo.Context) error {
	var results []models.InstanceListPlayer

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

	err := e.db.QueryContext(models.InstanceListPlayer{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getInstanceListPlayersCount godoc
// @Id getInstanceListPlayersCount
// @Summary Counts InstanceListPlayers
// @Accept json
// @Produce json
// @Tags InstanceListPlayer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.InstanceListPlayer
// @Failure 500 {string} string "Bad query request"
// @Router /instance_list_players/count [get]
func (e *InstanceListPlayerController) getInstanceListPlayersCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.InstanceListPlayer{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}