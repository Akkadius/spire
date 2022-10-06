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

type AaRankPrereqController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAaRankPrereqController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AaRankPrereqController {
	return &AaRankPrereqController{
		db:	    db,
		logger: logger,
	}
}

func (e *AaRankPrereqController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "aa_rank_prereq/:rankId", e.getAaRankPrereq, nil),
		routes.RegisterRoute(http.MethodGet, "aa_rank_prereqs", e.listAaRankPrereqs, nil),
		routes.RegisterRoute(http.MethodPut, "aa_rank_prereq", e.createAaRankPrereq, nil),
		routes.RegisterRoute(http.MethodDelete, "aa_rank_prereq/:rankId", e.deleteAaRankPrereq, nil),
		routes.RegisterRoute(http.MethodPatch, "aa_rank_prereq/:rankId", e.updateAaRankPrereq, nil),
		routes.RegisterRoute(http.MethodPost, "aa_rank_prereqs/bulk", e.getAaRankPrereqsBulk, nil),
	}
}

// listAaRankPrereqs godoc
// @Id listAaRankPrereqs
// @Summary Lists AaRankPrereqs
// @Accept json
// @Produce json
// @Tags AaRankPrereq
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaRankPrereq
// @Failure 500 {string} string "Bad query request"
// @Router /aa_rank_prereqs [get]
func (e *AaRankPrereqController) listAaRankPrereqs(c echo.Context) error {
	var results []models.AaRankPrereq
	err := e.db.QueryContext(models.AaRankPrereq{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAaRankPrereq godoc
// @Id getAaRankPrereq
// @Summary Gets AaRankPrereq
// @Accept json
// @Produce json
// @Tags AaRankPrereq
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaRankPrereq
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /aa_rank_prereq/{id} [get]
func (e *AaRankPrereqController) getAaRankPrereq(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rankId, err := strconv.Atoi(c.Param("rankId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [RankId]"})
	}
	params = append(params, rankId)
	keys = append(keys, "rank_id = ?")

	// key param [aa_id] position [2] type [int]
	if len(c.QueryParam("aa_id")) > 0 {
		aaIdParam, err := strconv.Atoi(c.QueryParam("aa_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [aa_id] err [%s]", err.Error())})
		}

		params = append(params, aaIdParam)
		keys = append(keys, "aa_id = ?")
	}

	// query builder
	var result models.AaRankPrereq
	query := e.db.QueryContext(models.AaRankPrereq{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.RankId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAaRankPrereq godoc
// @Id updateAaRankPrereq
// @Summary Updates AaRankPrereq
// @Accept json
// @Produce json
// @Tags AaRankPrereq
// @Param id path int true "Id"
// @Param aa_rank_prereq body models.AaRankPrereq true "AaRankPrereq"
// @Success 200 {array} models.AaRankPrereq
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /aa_rank_prereq/{id} [patch]
func (e *AaRankPrereqController) updateAaRankPrereq(c echo.Context) error {
	request := new(models.AaRankPrereq)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	rankId, err := strconv.Atoi(c.Param("rankId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [RankId]"})
	}
	params = append(params, rankId)
	keys = append(keys, "rank_id = ?")

	// key param [aa_id] position [2] type [int]
	if len(c.QueryParam("aa_id")) > 0 {
		aaIdParam, err := strconv.Atoi(c.QueryParam("aa_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [aa_id] err [%s]", err.Error())})
		}

		params = append(params, aaIdParam)
		keys = append(keys, "aa_id = ?")
	}

	// query builder
	var result models.AaRankPrereq
	query := e.db.QueryContext(models.AaRankPrereq{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAaRankPrereq godoc
// @Id createAaRankPrereq
// @Summary Creates AaRankPrereq
// @Accept json
// @Produce json
// @Param aa_rank_prereq body models.AaRankPrereq true "AaRankPrereq"
// @Tags AaRankPrereq
// @Success 200 {array} models.AaRankPrereq
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /aa_rank_prereq [put]
func (e *AaRankPrereqController) createAaRankPrereq(c echo.Context) error {
	aaRankPrereq := new(models.AaRankPrereq)
	if err := c.Bind(aaRankPrereq); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AaRankPrereq{}, c).Model(&models.AaRankPrereq{}).Create(&aaRankPrereq).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, aaRankPrereq)
}

// deleteAaRankPrereq godoc
// @Id deleteAaRankPrereq
// @Summary Deletes AaRankPrereq
// @Accept json
// @Produce json
// @Tags AaRankPrereq
// @Param id path int true "rankId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /aa_rank_prereq/{id} [delete]
func (e *AaRankPrereqController) deleteAaRankPrereq(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rankId, err := strconv.Atoi(c.Param("rankId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, rankId)
	keys = append(keys, "rank_id = ?")

	// key param [aa_id] position [2] type [int]
	if len(c.QueryParam("aa_id")) > 0 {
		aaIdParam, err := strconv.Atoi(c.QueryParam("aa_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [aa_id] err [%s]", err.Error())})
		}

		params = append(params, aaIdParam)
		keys = append(keys, "aa_id = ?")
	}

	// query builder
	var result models.AaRankPrereq
	query := e.db.QueryContext(models.AaRankPrereq{}, c)
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

// getAaRankPrereqsBulk godoc
// @Id getAaRankPrereqsBulk
// @Summary Gets AaRankPrereqs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AaRankPrereq
// @Success 200 {array} models.AaRankPrereq
// @Failure 500 {string} string "Bad query request"
// @Router /aa_rank_prereqs/bulk [post]
func (e *AaRankPrereqController) getAaRankPrereqsBulk(c echo.Context) error {
	var results []models.AaRankPrereq

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

	err := e.db.QueryContext(models.AaRankPrereq{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
