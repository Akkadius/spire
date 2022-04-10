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

type HackerController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewHackerController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *HackerController {
	return &HackerController{
		db:	 db,
		logger: logger,
	}
}

func (e *HackerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "hacker/:id", e.getHacker, nil),
		routes.RegisterRoute(http.MethodGet, "hackers", e.listHackers, nil),
		routes.RegisterRoute(http.MethodPut, "hacker", e.createHacker, nil),
		routes.RegisterRoute(http.MethodDelete, "hacker/:id", e.deleteHacker, nil),
		routes.RegisterRoute(http.MethodPatch, "hacker/:id", e.updateHacker, nil),
		routes.RegisterRoute(http.MethodPost, "hackers/bulk", e.getHackersBulk, nil),
	}
}

// listHackers godoc
// @Id listHackers
// @Summary Lists Hackers
// @Accept json
// @Produce json
// @Tags Hacker
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Hacker
// @Failure 500 {string} string "Bad query request"
// @Router /hackers [get]
func (e *HackerController) listHackers(c echo.Context) error {
	var results []models.Hacker
	err := e.db.QueryContext(models.Hacker{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getHacker godoc
// @Id getHacker
// @Summary Gets Hacker
// @Accept json
// @Produce json
// @Tags Hacker
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Hacker
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /hacker/{id} [get]
func (e *HackerController) getHacker(c echo.Context) error {
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
	var result models.Hacker
	query := e.db.QueryContext(models.Hacker{}, c)
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

// updateHacker godoc
// @Id updateHacker
// @Summary Updates Hacker
// @Accept json
// @Produce json
// @Tags Hacker
// @Param id path int true "Id"
// @Param hacker body models.Hacker true "Hacker"
// @Success 200 {array} models.Hacker
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /hacker/{id} [patch]
func (e *HackerController) updateHacker(c echo.Context) error {
	request := new(models.Hacker)
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
	var result models.Hacker
	query := e.db.QueryContext(models.Hacker{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Hacker{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createHacker godoc
// @Id createHacker
// @Summary Creates Hacker
// @Accept json
// @Produce json
// @Param hacker body models.Hacker true "Hacker"
// @Tags Hacker
// @Success 200 {array} models.Hacker
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /hacker [put]
func (e *HackerController) createHacker(c echo.Context) error {
	hacker := new(models.Hacker)
	if err := c.Bind(hacker); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Hacker{}, c).Model(&models.Hacker{}).Create(&hacker).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, hacker)
}

// deleteHacker godoc
// @Id deleteHacker
// @Summary Deletes Hacker
// @Accept json
// @Produce json
// @Tags Hacker
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /hacker/{id} [delete]
func (e *HackerController) deleteHacker(c echo.Context) error {
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
	var result models.Hacker
	query := e.db.QueryContext(models.Hacker{}, c)
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

// getHackersBulk godoc
// @Id getHackersBulk
// @Summary Gets Hackers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Hacker
// @Success 200 {array} models.Hacker
// @Failure 500 {string} string "Bad query request"
// @Router /hackers/bulk [post]
func (e *HackerController) getHackersBulk(c echo.Context) error {
	var results []models.Hacker

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

	err := e.db.QueryContext(models.Hacker{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
