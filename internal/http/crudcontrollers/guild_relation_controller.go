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

type GuildRelationController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGuildRelationController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GuildRelationController {
	return &GuildRelationController{
		db:	 db,
		logger: logger,
	}
}

func (e *GuildRelationController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "guild_relation/:guild1", e.getGuildRelation, nil),
		routes.RegisterRoute(http.MethodGet, "guild_relations", e.listGuildRelations, nil),
		routes.RegisterRoute(http.MethodPut, "guild_relation", e.createGuildRelation, nil),
		routes.RegisterRoute(http.MethodDelete, "guild_relation/:guild1", e.deleteGuildRelation, nil),
		routes.RegisterRoute(http.MethodPatch, "guild_relation/:guild1", e.updateGuildRelation, nil),
		routes.RegisterRoute(http.MethodPost, "guild_relations/bulk", e.getGuildRelationsBulk, nil),
	}
}

// listGuildRelations godoc
// @Id listGuildRelations
// @Summary Lists GuildRelations
// @Accept json
// @Produce json
// @Tags GuildRelation
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildRelation
// @Failure 500 {string} string "Bad query request"
// @Router /guild_relations [get]
func (e *GuildRelationController) listGuildRelations(c echo.Context) error {
	var results []models.GuildRelation
	err := e.db.QueryContext(models.GuildRelation{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildRelation godoc
// @Id getGuildRelation
// @Summary Gets GuildRelation
// @Accept json
// @Produce json
// @Tags GuildRelation
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildRelation
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /guild_relation/{id} [get]
func (e *GuildRelationController) getGuildRelation(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	guild1, err := strconv.Atoi(c.Param("guild1"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Guild1]"})
	}
	params = append(params, guild1)
	keys = append(keys, "guild1 = ?")

	// key param [guild2] position [2] type [mediumint]
	if len(c.QueryParam("guild2")) > 0 {
		guild2Param, err := strconv.Atoi(c.QueryParam("guild2"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [guild2] err [%s]", err.Error())})
		}

		params = append(params, guild2Param)
		keys = append(keys, "guild2 = ?")
	}

	// query builder
	var result models.GuildRelation
	query := e.db.QueryContext(models.GuildRelation{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Guild1 == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGuildRelation godoc
// @Id updateGuildRelation
// @Summary Updates GuildRelation
// @Accept json
// @Produce json
// @Tags GuildRelation
// @Param id path int true "Id"
// @Param guild_relation body models.GuildRelation true "GuildRelation"
// @Success 200 {array} models.GuildRelation
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /guild_relation/{id} [patch]
func (e *GuildRelationController) updateGuildRelation(c echo.Context) error {
	request := new(models.GuildRelation)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	guild1, err := strconv.Atoi(c.Param("guild1"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Guild1]"})
	}
	params = append(params, guild1)
	keys = append(keys, "guild1 = ?")

	// key param [guild2] position [2] type [mediumint]
	if len(c.QueryParam("guild2")) > 0 {
		guild2Param, err := strconv.Atoi(c.QueryParam("guild2"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [guild2] err [%s]", err.Error())})
		}

		params = append(params, guild2Param)
		keys = append(keys, "guild2 = ?")
	}

	// query builder
	var result models.GuildRelation
	query := e.db.QueryContext(models.GuildRelation{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.GuildRelation{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createGuildRelation godoc
// @Id createGuildRelation
// @Summary Creates GuildRelation
// @Accept json
// @Produce json
// @Param guild_relation body models.GuildRelation true "GuildRelation"
// @Tags GuildRelation
// @Success 200 {array} models.GuildRelation
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /guild_relation [put]
func (e *GuildRelationController) createGuildRelation(c echo.Context) error {
	guildRelation := new(models.GuildRelation)
	if err := c.Bind(guildRelation); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.GuildRelation{}, c).Model(&models.GuildRelation{}).Create(&guildRelation).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, guildRelation)
}

// deleteGuildRelation godoc
// @Id deleteGuildRelation
// @Summary Deletes GuildRelation
// @Accept json
// @Produce json
// @Tags GuildRelation
// @Param id path int true "guild1"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /guild_relation/{id} [delete]
func (e *GuildRelationController) deleteGuildRelation(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	guild1, err := strconv.Atoi(c.Param("guild1"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, guild1)
	keys = append(keys, "guild1 = ?")

	// key param [guild2] position [2] type [mediumint]
	if len(c.QueryParam("guild2")) > 0 {
		guild2Param, err := strconv.Atoi(c.QueryParam("guild2"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [guild2] err [%s]", err.Error())})
		}

		params = append(params, guild2Param)
		keys = append(keys, "guild2 = ?")
	}

	// query builder
	var result models.GuildRelation
	query := e.db.QueryContext(models.GuildRelation{}, c)
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

// getGuildRelationsBulk godoc
// @Id getGuildRelationsBulk
// @Summary Gets GuildRelations in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GuildRelation
// @Success 200 {array} models.GuildRelation
// @Failure 500 {string} string "Bad query request"
// @Router /guild_relations/bulk [post]
func (e *GuildRelationController) getGuildRelationsBulk(c echo.Context) error {
	var results []models.GuildRelation

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

	err := e.db.QueryContext(models.GuildRelation{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
