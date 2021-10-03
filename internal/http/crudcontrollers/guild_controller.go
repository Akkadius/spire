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

type GuildController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGuildController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GuildController {
	return &GuildController{
		db:     db,
		logger: logger,
	}
}

func (e *GuildController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "guild/:guild", e.deleteGuild, nil),
		routes.RegisterRoute(http.MethodGet, "guild/:guild", e.getGuild, nil),
		routes.RegisterRoute(http.MethodGet, "guilds", e.listGuilds, nil),
		routes.RegisterRoute(http.MethodPost, "guilds/bulk", e.getGuildsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "guild/:guild", e.updateGuild, nil),
		routes.RegisterRoute(http.MethodPut, "guild", e.createGuild, nil),
	}
}

// listGuilds godoc
// @Id listGuilds
// @Summary Lists Guilds
// @Accept json
// @Produce json
// @Tags Guild
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>GuildBanks<br>GuildMembers<br>GuildRanks"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Guild
// @Failure 500 {string} string "Bad query request"
// @Router /guilds [get]
func (e *GuildController) listGuilds(c echo.Context) error {
	var results []models.Guild
	err := e.db.QueryContext(models.Guild{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuild godoc
// @Id getGuild
// @Summary Gets Guild
// @Accept json
// @Produce json
// @Tags Guild
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>GuildBanks<br>GuildMembers<br>GuildRanks"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Guild
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /guild/{id} [get]
func (e *GuildController) getGuild(c echo.Context) error {
	guildId, err := strconv.Atoi(c.Param("guild"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Guild
	err = e.db.QueryContext(models.Guild{}, c).First(&result, guildId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGuild godoc
// @Id updateGuild
// @Summary Updates Guild
// @Accept json
// @Produce json
// @Tags Guild
// @Param id path int true "Id"
// @Param guild body models.Guild true "Guild"
// @Success 200 {array} models.Guild
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /guild/{id} [patch]
func (e *GuildController) updateGuild(c echo.Context) error {
	guild := new(models.Guild)
	if err := c.Bind(guild); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Guild{}
	err := e.db.Get(models.Guild{}, c).Model(&models.Guild{}).First(&entity, guild.ID).Error
	if err != nil || guild.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Guild{}, c).Model(&entity).Updates(&guild).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, guild)
}

// createGuild godoc
// @Id createGuild
// @Summary Creates Guild
// @Accept json
// @Produce json
// @Param guild body models.Guild true "Guild"
// @Tags Guild
// @Success 200 {array} models.Guild
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /guild [put]
func (e *GuildController) createGuild(c echo.Context) error {
	guild := new(models.Guild)
	if err := c.Bind(guild); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Guild{}, c).Model(&models.Guild{}).Create(&guild).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, guild)
}

// deleteGuild godoc
// @Id deleteGuild
// @Summary Deletes Guild
// @Accept json
// @Produce json
// @Tags Guild
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /guild/{id} [delete]
func (e *GuildController) deleteGuild(c echo.Context) error {
	guildId, err := strconv.Atoi(c.Param("guild"))
	if err != nil {
		e.logger.Error(err)
	}

	guild := new(models.Guild)
	err = e.db.Get(models.Guild{}, c).Model(&models.Guild{}).First(&guild, guildId).Error
	if err != nil || guild.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Guild{}, c).Model(&models.Guild{}).Delete(&guild).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getGuildsBulk godoc
// @Id getGuildsBulk
// @Summary Gets Guilds in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Guild
// @Success 200 {array} models.Guild
// @Failure 500 {string} string "Bad query request"
// @Router /guilds/bulk [post]
func (e *GuildController) getGuildsBulk(c echo.Context) error {
	var results []models.Guild

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

	err := e.db.QueryContext(models.Guild{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
