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

type AaRankController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAaRankController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AaRankController {
	return &AaRankController{
		db:	 db,
		logger: logger,
	}
}

func (e *AaRankController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "aa_rank/:id", e.getAaRank, nil),
		routes.RegisterRoute(http.MethodGet, "aa_ranks", e.listAaRanks, nil),
		routes.RegisterRoute(http.MethodPut, "aa_rank", e.createAaRank, nil),
		routes.RegisterRoute(http.MethodDelete, "aa_rank/:id", e.deleteAaRank, nil),
		routes.RegisterRoute(http.MethodPatch, "aa_rank/:id", e.updateAaRank, nil),
		routes.RegisterRoute(http.MethodPost, "aa_ranks/bulk", e.getAaRanksBulk, nil),
	}
}

// listAaRanks godoc
// @Id listAaRanks
// @Summary Lists AaRanks
// @Accept json
// @Produce json
// @Tags AaRank
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaRank
// @Failure 500 {string} string "Bad query request"
// @Router /aa_ranks [get]
func (e *AaRankController) listAaRanks(c echo.Context) error {
	var results []models.AaRank
	err := e.db.QueryContext(models.AaRank{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAaRank godoc
// @Id getAaRank
// @Summary Gets AaRank
// @Accept json
// @Produce json
// @Tags AaRank
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaRank
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /aa_rank/{id} [get]
func (e *AaRankController) getAaRank(c echo.Context) error {
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
	var result models.AaRank
	query := e.db.QueryContext(models.AaRank{}, c)
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

// updateAaRank godoc
// @Id updateAaRank
// @Summary Updates AaRank
// @Accept json
// @Produce json
// @Tags AaRank
// @Param id path int true "Id"
// @Param aa_rank body models.AaRank true "AaRank"
// @Success 200 {array} models.AaRank
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /aa_rank/{id} [patch]
func (e *AaRankController) updateAaRank(c echo.Context) error {
	request := new(models.AaRank)
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
	var result models.AaRank
	query := e.db.QueryContext(models.AaRank{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AaRank{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAaRank godoc
// @Id createAaRank
// @Summary Creates AaRank
// @Accept json
// @Produce json
// @Param aa_rank body models.AaRank true "AaRank"
// @Tags AaRank
// @Success 200 {array} models.AaRank
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /aa_rank [put]
func (e *AaRankController) createAaRank(c echo.Context) error {
	aaRank := new(models.AaRank)
	if err := c.Bind(aaRank); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AaRank{}, c).Model(&models.AaRank{}).Create(&aaRank).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, aaRank)
}

// deleteAaRank godoc
// @Id deleteAaRank
// @Summary Deletes AaRank
// @Accept json
// @Produce json
// @Tags AaRank
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /aa_rank/{id} [delete]
func (e *AaRankController) deleteAaRank(c echo.Context) error {
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
	var result models.AaRank
	query := e.db.QueryContext(models.AaRank{}, c)
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

// getAaRanksBulk godoc
// @Id getAaRanksBulk
// @Summary Gets AaRanks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AaRank
// @Success 200 {array} models.AaRank
// @Failure 500 {string} string "Bad query request"
// @Router /aa_ranks/bulk [post]
func (e *AaRankController) getAaRanksBulk(c echo.Context) error {
	var results []models.AaRank

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

	err := e.db.QueryContext(models.AaRank{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
