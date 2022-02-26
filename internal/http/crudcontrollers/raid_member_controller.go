package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RaidMemberController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewRaidMemberController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *RaidMemberController {
	return &RaidMemberController{
		db:	 db,
		logger: logger,
	}
}

func (e *RaidMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "raid_member/:charid", e.getRaidMember, nil),
		routes.RegisterRoute(http.MethodGet, "raid_members", e.listRaidMembers, nil),
		routes.RegisterRoute(http.MethodPut, "raid_member", e.createRaidMember, nil),
		routes.RegisterRoute(http.MethodDelete, "raid_member/:charid", e.deleteRaidMember, nil),
		routes.RegisterRoute(http.MethodPatch, "raid_member/:charid", e.updateRaidMember, nil),
		routes.RegisterRoute(http.MethodPost, "raid_members/bulk", e.getRaidMembersBulk, nil),
	}
}

// listRaidMembers godoc
// @Id listRaidMembers
// @Summary Lists RaidMembers
// @Accept json
// @Produce json
// @Tags RaidMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RaidMember
// @Failure 500 {string} string "Bad query request"
// @Router /raid_members [get]
func (e *RaidMemberController) listRaidMembers(c echo.Context) error {
	var results []models.RaidMember
	err := e.db.QueryContext(models.RaidMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getRaidMember godoc
// @Id getRaidMember
// @Summary Gets RaidMember
// @Accept json
// @Produce json
// @Tags RaidMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RaidMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /raid_member/{id} [get]
func (e *RaidMemberController) getRaidMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// query builder
	var result models.RaidMember
	query := e.db.QueryContext(models.RaidMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Charid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateRaidMember godoc
// @Id updateRaidMember
// @Summary Updates RaidMember
// @Accept json
// @Produce json
// @Tags RaidMember
// @Param id path int true "Id"
// @Param raid_member body models.RaidMember true "RaidMember"
// @Success 200 {array} models.RaidMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /raid_member/{id} [patch]
func (e *RaidMemberController) updateRaidMember(c echo.Context) error {
	request := new(models.RaidMember)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// query builder
	var result models.RaidMember
	query := e.db.QueryContext(models.RaidMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createRaidMember godoc
// @Id createRaidMember
// @Summary Creates RaidMember
// @Accept json
// @Produce json
// @Param raid_member body models.RaidMember true "RaidMember"
// @Tags RaidMember
// @Success 200 {array} models.RaidMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /raid_member [put]
func (e *RaidMemberController) createRaidMember(c echo.Context) error {
	raidMember := new(models.RaidMember)
	if err := c.Bind(raidMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.RaidMember{}, c).Model(&models.RaidMember{}).Create(&raidMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, raidMember)
}

// deleteRaidMember godoc
// @Id deleteRaidMember
// @Summary Deletes RaidMember
// @Accept json
// @Produce json
// @Tags RaidMember
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /raid_member/{id} [delete]
func (e *RaidMemberController) deleteRaidMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// query builder
	var result models.RaidMember
	query := e.db.QueryContext(models.RaidMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.RaidMember{}, c).Model(&models.RaidMember{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getRaidMembersBulk godoc
// @Id getRaidMembersBulk
// @Summary Gets RaidMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags RaidMember
// @Success 200 {array} models.RaidMember
// @Failure 500 {string} string "Bad query request"
// @Router /raid_members/bulk [post]
func (e *RaidMemberController) getRaidMembersBulk(c echo.Context) error {
	var results []models.RaidMember

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

	err := e.db.QueryContext(models.RaidMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
