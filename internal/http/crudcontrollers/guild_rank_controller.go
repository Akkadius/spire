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

type GuildRankController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewGuildRankController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *GuildRankController {
	return &GuildRankController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *GuildRankController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "guild_rank/:guildId", e.getGuildRank, nil),
		routes.RegisterRoute(http.MethodGet, "guild_ranks", e.listGuildRanks, nil),
		routes.RegisterRoute(http.MethodGet, "guild_ranks/count", e.getGuildRanksCount, nil),
		routes.RegisterRoute(http.MethodPut, "guild_rank", e.createGuildRank, nil),
		routes.RegisterRoute(http.MethodDelete, "guild_rank/:guildId", e.deleteGuildRank, nil),
		routes.RegisterRoute(http.MethodPatch, "guild_rank/:guildId", e.updateGuildRank, nil),
		routes.RegisterRoute(http.MethodPost, "guild_ranks/bulk", e.getGuildRanksBulk, nil),
	}
}

// listGuildRanks godoc
// @Id listGuildRanks
// @Summary Lists GuildRanks
// @Accept json
// @Produce json
// @Tags GuildRank
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildRank
// @Failure 500 {string} string "Bad query request"
// @Router /guild_ranks [get]
func (e *GuildRankController) listGuildRanks(c echo.Context) error {
	var results []models.GuildRank
	err := e.db.QueryContext(models.GuildRank{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildRank godoc
// @Id getGuildRank
// @Summary Gets GuildRank
// @Accept json
// @Produce json
// @Tags GuildRank
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildRank
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /guild_rank/{id} [get]
func (e *GuildRankController) getGuildRank(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	guildId, err := strconv.Atoi(c.Param("guildId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GuildId]"})
	}
	params = append(params, guildId)
	keys = append(keys, "guild_id = ?")

	// key param [rank] position [2] type [tinyint]
	if len(c.QueryParam("rank")) > 0 {
		rankParam, err := strconv.Atoi(c.QueryParam("rank"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [rank] err [%s]", err.Error())})
		}

		params = append(params, rankParam)
		keys = append(keys, "rank = ?")
	}

	// query builder
	var result models.GuildRank
	query := e.db.QueryContext(models.GuildRank{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.GuildId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGuildRank godoc
// @Id updateGuildRank
// @Summary Updates GuildRank
// @Accept json
// @Produce json
// @Tags GuildRank
// @Param id path int true "Id"
// @Param guild_rank body models.GuildRank true "GuildRank"
// @Success 200 {array} models.GuildRank
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /guild_rank/{id} [patch]
func (e *GuildRankController) updateGuildRank(c echo.Context) error {
	request := new(models.GuildRank)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	guildId, err := strconv.Atoi(c.Param("guildId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [GuildId]"})
	}
	params = append(params, guildId)
	keys = append(keys, "guild_id = ?")

	// key param [rank] position [2] type [tinyint]
	if len(c.QueryParam("rank")) > 0 {
		rankParam, err := strconv.Atoi(c.QueryParam("rank"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [rank] err [%s]", err.Error())})
		}

		params = append(params, rankParam)
		keys = append(keys, "rank = ?")
	}

	// query builder
	var result models.GuildRank
	query := e.db.QueryContext(models.GuildRank{}, c)
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
		event := fmt.Sprintf("Updated [GuildRank] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createGuildRank godoc
// @Id createGuildRank
// @Summary Creates GuildRank
// @Accept json
// @Produce json
// @Param guild_rank body models.GuildRank true "GuildRank"
// @Tags GuildRank
// @Success 200 {array} models.GuildRank
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /guild_rank [put]
func (e *GuildRankController) createGuildRank(c echo.Context) error {
	guildRank := new(models.GuildRank)
	if err := c.Bind(guildRank); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.GuildRank{}, c).Model(&models.GuildRank{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&guildRank).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GuildRank{}, guildRank)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GuildRank] [%v] data [%v]", guildRank.GuildId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, guildRank)
}

// deleteGuildRank godoc
// @Id deleteGuildRank
// @Summary Deletes GuildRank
// @Accept json
// @Produce json
// @Tags GuildRank
// @Param id path int true "guildId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /guild_rank/{id} [delete]
func (e *GuildRankController) deleteGuildRank(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	guildId, err := strconv.Atoi(c.Param("guildId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, guildId)
	keys = append(keys, "guild_id = ?")

	// key param [rank] position [2] type [tinyint]
	if len(c.QueryParam("rank")) > 0 {
		rankParam, err := strconv.Atoi(c.QueryParam("rank"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [rank] err [%s]", err.Error())})
		}

		params = append(params, rankParam)
		keys = append(keys, "rank = ?")
	}

	// query builder
	var result models.GuildRank
	query := e.db.QueryContext(models.GuildRank{}, c)
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
		event := fmt.Sprintf("Deleted [GuildRank] [%v] keys [%v]", result.GuildId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGuildRanksBulk godoc
// @Id getGuildRanksBulk
// @Summary Gets GuildRanks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GuildRank
// @Success 200 {array} models.GuildRank
// @Failure 500 {string} string "Bad query request"
// @Router /guild_ranks/bulk [post]
func (e *GuildRankController) getGuildRanksBulk(c echo.Context) error {
	var results []models.GuildRank

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

	err := e.db.QueryContext(models.GuildRank{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildRanksCount godoc
// @Id getGuildRanksCount
// @Summary Counts GuildRanks
// @Accept json
// @Produce json
// @Tags GuildRank
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildRank
// @Failure 500 {string} string "Bad query request"
// @Router /guild_ranks/count [get]
func (e *GuildRankController) getGuildRanksCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.GuildRank{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}