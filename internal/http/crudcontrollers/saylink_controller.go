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

type SaylinkController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSaylinkController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SaylinkController {
	return &SaylinkController{
		db:	 db,
		logger: logger,
	}
}

func (e *SaylinkController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "saylink/:id", e.getSaylink, nil),
		routes.RegisterRoute(http.MethodGet, "saylinks", e.listSaylinks, nil),
		routes.RegisterRoute(http.MethodPut, "saylink", e.createSaylink, nil),
		routes.RegisterRoute(http.MethodDelete, "saylink/:id", e.deleteSaylink, nil),
		routes.RegisterRoute(http.MethodPatch, "saylink/:id", e.updateSaylink, nil),
		routes.RegisterRoute(http.MethodPost, "saylinks/bulk", e.getSaylinksBulk, nil),
	}
}

// listSaylinks godoc
// @Id listSaylinks
// @Summary Lists Saylinks
// @Accept json
// @Produce json
// @Tags Saylink
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Saylink
// @Failure 500 {string} string "Bad query request"
// @Router /saylinks [get]
func (e *SaylinkController) listSaylinks(c echo.Context) error {
	var results []models.Saylink
	err := e.db.QueryContext(models.Saylink{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSaylink godoc
// @Id getSaylink
// @Summary Gets Saylink
// @Accept json
// @Produce json
// @Tags Saylink
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Saylink
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /saylink/{id} [get]
func (e *SaylinkController) getSaylink(c echo.Context) error {
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
	var result models.Saylink
	query := e.db.QueryContext(models.Saylink{}, c)
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

// updateSaylink godoc
// @Id updateSaylink
// @Summary Updates Saylink
// @Accept json
// @Produce json
// @Tags Saylink
// @Param id path int true "Id"
// @Param saylink body models.Saylink true "Saylink"
// @Success 200 {array} models.Saylink
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /saylink/{id} [patch]
func (e *SaylinkController) updateSaylink(c echo.Context) error {
	request := new(models.Saylink)
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
	var result models.Saylink
	query := e.db.QueryContext(models.Saylink{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Saylink{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createSaylink godoc
// @Id createSaylink
// @Summary Creates Saylink
// @Accept json
// @Produce json
// @Param saylink body models.Saylink true "Saylink"
// @Tags Saylink
// @Success 200 {array} models.Saylink
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /saylink [put]
func (e *SaylinkController) createSaylink(c echo.Context) error {
	saylink := new(models.Saylink)
	if err := c.Bind(saylink); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Saylink{}, c).Model(&models.Saylink{}).Create(&saylink).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, saylink)
}

// deleteSaylink godoc
// @Id deleteSaylink
// @Summary Deletes Saylink
// @Accept json
// @Produce json
// @Tags Saylink
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /saylink/{id} [delete]
func (e *SaylinkController) deleteSaylink(c echo.Context) error {
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
	var result models.Saylink
	query := e.db.QueryContext(models.Saylink{}, c)
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

// getSaylinksBulk godoc
// @Id getSaylinksBulk
// @Summary Gets Saylinks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Saylink
// @Success 200 {array} models.Saylink
// @Failure 500 {string} string "Bad query request"
// @Router /saylinks/bulk [post]
func (e *SaylinkController) getSaylinksBulk(c echo.Context) error {
	var results []models.Saylink

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

	err := e.db.QueryContext(models.Saylink{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
