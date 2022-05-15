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

type RaidDetailController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewRaidDetailController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *RaidDetailController {
	return &RaidDetailController{
		db:	 db,
		logger: logger,
	}
}

func (e *RaidDetailController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "raid_detail/:raidid", e.getRaidDetail, nil),
		routes.RegisterRoute(http.MethodGet, "raid_details", e.listRaidDetails, nil),
		routes.RegisterRoute(http.MethodPut, "raid_detail", e.createRaidDetail, nil),
		routes.RegisterRoute(http.MethodDelete, "raid_detail/:raidid", e.deleteRaidDetail, nil),
		routes.RegisterRoute(http.MethodPatch, "raid_detail/:raidid", e.updateRaidDetail, nil),
		routes.RegisterRoute(http.MethodPost, "raid_details/bulk", e.getRaidDetailsBulk, nil),
	}
}

// listRaidDetails godoc
// @Id listRaidDetails
// @Summary Lists RaidDetails
// @Accept json
// @Produce json
// @Tags RaidDetail
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>RaidLeaders<br>RaidMembers"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RaidDetail
// @Failure 500 {string} string "Bad query request"
// @Router /raid_details [get]
func (e *RaidDetailController) listRaidDetails(c echo.Context) error {
	var results []models.RaidDetail
	err := e.db.QueryContext(models.RaidDetail{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getRaidDetail godoc
// @Id getRaidDetail
// @Summary Gets RaidDetail
// @Accept json
// @Produce json
// @Tags RaidDetail
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>RaidLeaders<br>RaidMembers"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.RaidDetail
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /raid_detail/{id} [get]
func (e *RaidDetailController) getRaidDetail(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	raidid, err := strconv.Atoi(c.Param("raidid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Raidid]"})
	}
	params = append(params, raidid)
	keys = append(keys, "raidid = ?")

	// query builder
	var result models.RaidDetail
	query := e.db.QueryContext(models.RaidDetail{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Raidid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateRaidDetail godoc
// @Id updateRaidDetail
// @Summary Updates RaidDetail
// @Accept json
// @Produce json
// @Tags RaidDetail
// @Param id path int true "Id"
// @Param raid_detail body models.RaidDetail true "RaidDetail"
// @Success 200 {array} models.RaidDetail
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /raid_detail/{id} [patch]
func (e *RaidDetailController) updateRaidDetail(c echo.Context) error {
	request := new(models.RaidDetail)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	raidid, err := strconv.Atoi(c.Param("raidid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Raidid]"})
	}
	params = append(params, raidid)
	keys = append(keys, "raidid = ?")

	// query builder
	var result models.RaidDetail
	query := e.db.QueryContext(models.RaidDetail{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.RaidDetail{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createRaidDetail godoc
// @Id createRaidDetail
// @Summary Creates RaidDetail
// @Accept json
// @Produce json
// @Param raid_detail body models.RaidDetail true "RaidDetail"
// @Tags RaidDetail
// @Success 200 {array} models.RaidDetail
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /raid_detail [put]
func (e *RaidDetailController) createRaidDetail(c echo.Context) error {
	raidDetail := new(models.RaidDetail)
	if err := c.Bind(raidDetail); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.RaidDetail{}, c).Model(&models.RaidDetail{}).Create(&raidDetail).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, raidDetail)
}

// deleteRaidDetail godoc
// @Id deleteRaidDetail
// @Summary Deletes RaidDetail
// @Accept json
// @Produce json
// @Tags RaidDetail
// @Param id path int true "raidid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /raid_detail/{id} [delete]
func (e *RaidDetailController) deleteRaidDetail(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	raidid, err := strconv.Atoi(c.Param("raidid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, raidid)
	keys = append(keys, "raidid = ?")

	// query builder
	var result models.RaidDetail
	query := e.db.QueryContext(models.RaidDetail{}, c)
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

// getRaidDetailsBulk godoc
// @Id getRaidDetailsBulk
// @Summary Gets RaidDetails in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags RaidDetail
// @Success 200 {array} models.RaidDetail
// @Failure 500 {string} string "Bad query request"
// @Router /raid_details/bulk [post]
func (e *RaidDetailController) getRaidDetailsBulk(c echo.Context) error {
	var results []models.RaidDetail

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

	err := e.db.QueryContext(models.RaidDetail{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
