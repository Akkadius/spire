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

type GuildMemberController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGuildMemberController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GuildMemberController {
	return &GuildMemberController{
		db:	 db,
		logger: logger,
	}
}

func (e *GuildMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "guild_member/:charId", e.getGuildMember, nil),
		routes.RegisterRoute(http.MethodGet, "guild_members", e.listGuildMembers, nil),
		routes.RegisterRoute(http.MethodPut, "guild_member", e.createGuildMember, nil),
		routes.RegisterRoute(http.MethodDelete, "guild_member/:charId", e.deleteGuildMember, nil),
		routes.RegisterRoute(http.MethodPatch, "guild_member/:charId", e.updateGuildMember, nil),
		routes.RegisterRoute(http.MethodPost, "guild_members/bulk", e.getGuildMembersBulk, nil),
	}
}

// listGuildMembers godoc
// @Id listGuildMembers
// @Summary Lists GuildMembers
// @Accept json
// @Produce json
// @Tags GuildMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildMember
// @Failure 500 {string} string "Bad query request"
// @Router /guild_members [get]
func (e *GuildMemberController) listGuildMembers(c echo.Context) error {
	var results []models.GuildMember
	err := e.db.QueryContext(models.GuildMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGuildMember godoc
// @Id getGuildMember
// @Summary Gets GuildMember
// @Accept json
// @Produce json
// @Tags GuildMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GuildMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /guild_member/{id} [get]
func (e *GuildMemberController) getGuildMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// query builder
	var result models.GuildMember
	query := e.db.QueryContext(models.GuildMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateGuildMember godoc
// @Id updateGuildMember
// @Summary Updates GuildMember
// @Accept json
// @Produce json
// @Tags GuildMember
// @Param id path int true "Id"
// @Param guild_member body models.GuildMember true "GuildMember"
// @Success 200 {array} models.GuildMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /guild_member/{id} [patch]
func (e *GuildMemberController) updateGuildMember(c echo.Context) error {
	request := new(models.GuildMember)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// query builder
	var result models.GuildMember
	query := e.db.QueryContext(models.GuildMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.GuildMember{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createGuildMember godoc
// @Id createGuildMember
// @Summary Creates GuildMember
// @Accept json
// @Produce json
// @Param guild_member body models.GuildMember true "GuildMember"
// @Tags GuildMember
// @Success 200 {array} models.GuildMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /guild_member [put]
func (e *GuildMemberController) createGuildMember(c echo.Context) error {
	guildMember := new(models.GuildMember)
	if err := c.Bind(guildMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.GuildMember{}, c).Model(&models.GuildMember{}).Create(&guildMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, guildMember)
}

// deleteGuildMember godoc
// @Id deleteGuildMember
// @Summary Deletes GuildMember
// @Accept json
// @Produce json
// @Tags GuildMember
// @Param id path int true "charId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /guild_member/{id} [delete]
func (e *GuildMemberController) deleteGuildMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// query builder
	var result models.GuildMember
	query := e.db.QueryContext(models.GuildMember{}, c)
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

// getGuildMembersBulk godoc
// @Id getGuildMembersBulk
// @Summary Gets GuildMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GuildMember
// @Success 200 {array} models.GuildMember
// @Failure 500 {string} string "Bad query request"
// @Router /guild_members/bulk [post]
func (e *GuildMemberController) getGuildMembersBulk(c echo.Context) error {
	var results []models.GuildMember

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

	err := e.db.QueryContext(models.GuildMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
