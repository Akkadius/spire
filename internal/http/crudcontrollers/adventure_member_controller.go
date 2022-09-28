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

type AdventureMemberController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAdventureMemberController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AdventureMemberController {
	return &AdventureMemberController{
		db:	 db,
		logger: logger,
	}
}

func (e *AdventureMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "adventure_member/:charid", e.getAdventureMember, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_members", e.listAdventureMembers, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_member", e.createAdventureMember, nil),
		routes.RegisterRoute(http.MethodDelete, "adventure_member/:charid", e.deleteAdventureMember, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_member/:charid", e.updateAdventureMember, nil),
		routes.RegisterRoute(http.MethodPost, "adventure_members/bulk", e.getAdventureMembersBulk, nil),
	}
}

// listAdventureMembers godoc
// @Id listAdventureMembers
// @Summary Lists AdventureMembers
// @Accept json
// @Produce json
// @Tags AdventureMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureMember
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_members [get]
func (e *AdventureMemberController) listAdventureMembers(c echo.Context) error {
	var results []models.AdventureMember
	err := e.db.QueryContext(models.AdventureMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureMember godoc
// @Id getAdventureMember
// @Summary Gets AdventureMember
// @Accept json
// @Produce json
// @Tags AdventureMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_member/{id} [get]
func (e *AdventureMemberController) getAdventureMember(c echo.Context) error {
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
	var result models.AdventureMember
	query := e.db.QueryContext(models.AdventureMember{}, c)
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

// updateAdventureMember godoc
// @Id updateAdventureMember
// @Summary Updates AdventureMember
// @Accept json
// @Produce json
// @Tags AdventureMember
// @Param id path int true "Id"
// @Param adventure_member body models.AdventureMember true "AdventureMember"
// @Success 200 {array} models.AdventureMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_member/{id} [patch]
func (e *AdventureMemberController) updateAdventureMember(c echo.Context) error {
	request := new(models.AdventureMember)
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
	var result models.AdventureMember
	query := e.db.QueryContext(models.AdventureMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AdventureMember{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAdventureMember godoc
// @Id createAdventureMember
// @Summary Creates AdventureMember
// @Accept json
// @Produce json
// @Param adventure_member body models.AdventureMember true "AdventureMember"
// @Tags AdventureMember
// @Success 200 {array} models.AdventureMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_member [put]
func (e *AdventureMemberController) createAdventureMember(c echo.Context) error {
	adventureMember := new(models.AdventureMember)
	if err := c.Bind(adventureMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AdventureMember{}, c).Model(&models.AdventureMember{}).Create(&adventureMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, adventureMember)
}

// deleteAdventureMember godoc
// @Id deleteAdventureMember
// @Summary Deletes AdventureMember
// @Accept json
// @Produce json
// @Tags AdventureMember
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_member/{id} [delete]
func (e *AdventureMemberController) deleteAdventureMember(c echo.Context) error {
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
	var result models.AdventureMember
	query := e.db.QueryContext(models.AdventureMember{}, c)
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

// getAdventureMembersBulk godoc
// @Id getAdventureMembersBulk
// @Summary Gets AdventureMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AdventureMember
// @Success 200 {array} models.AdventureMember
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_members/bulk [post]
func (e *AdventureMemberController) getAdventureMembersBulk(c echo.Context) error {
	var results []models.AdventureMember

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

	err := e.db.QueryContext(models.AdventureMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
