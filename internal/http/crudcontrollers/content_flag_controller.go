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

type ContentFlagController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewContentFlagController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ContentFlagController {
	return &ContentFlagController{
		db:	 db,
		logger: logger,
	}
}

func (e *ContentFlagController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "content_flag/:id", e.getContentFlag, nil),
		routes.RegisterRoute(http.MethodGet, "content_flags", e.listContentFlags, nil),
		routes.RegisterRoute(http.MethodPut, "content_flag", e.createContentFlag, nil),
		routes.RegisterRoute(http.MethodDelete, "content_flag/:id", e.deleteContentFlag, nil),
		routes.RegisterRoute(http.MethodPatch, "content_flag/:id", e.updateContentFlag, nil),
		routes.RegisterRoute(http.MethodPost, "content_flags/bulk", e.getContentFlagsBulk, nil),
	}
}

// listContentFlags godoc
// @Id listContentFlags
// @Summary Lists ContentFlags
// @Accept json
// @Produce json
// @Tags ContentFlag
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ContentFlag
// @Failure 500 {string} string "Bad query request"
// @Router /content_flags [get]
func (e *ContentFlagController) listContentFlags(c echo.Context) error {
	var results []models.ContentFlag
	err := e.db.QueryContext(models.ContentFlag{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getContentFlag godoc
// @Id getContentFlag
// @Summary Gets ContentFlag
// @Accept json
// @Produce json
// @Tags ContentFlag
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ContentFlag
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /content_flag/{id} [get]
func (e *ContentFlagController) getContentFlag(c echo.Context) error {
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
	var result models.ContentFlag
	query := e.db.QueryContext(models.ContentFlag{}, c)
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

// updateContentFlag godoc
// @Id updateContentFlag
// @Summary Updates ContentFlag
// @Accept json
// @Produce json
// @Tags ContentFlag
// @Param id path int true "Id"
// @Param content_flag body models.ContentFlag true "ContentFlag"
// @Success 200 {array} models.ContentFlag
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /content_flag/{id} [patch]
func (e *ContentFlagController) updateContentFlag(c echo.Context) error {
	request := new(models.ContentFlag)
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
	var result models.ContentFlag
	query := e.db.QueryContext(models.ContentFlag{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.ContentFlag{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createContentFlag godoc
// @Id createContentFlag
// @Summary Creates ContentFlag
// @Accept json
// @Produce json
// @Param content_flag body models.ContentFlag true "ContentFlag"
// @Tags ContentFlag
// @Success 200 {array} models.ContentFlag
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /content_flag [put]
func (e *ContentFlagController) createContentFlag(c echo.Context) error {
	contentFlag := new(models.ContentFlag)
	if err := c.Bind(contentFlag); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.ContentFlag{}, c).Model(&models.ContentFlag{}).Create(&contentFlag).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, contentFlag)
}

// deleteContentFlag godoc
// @Id deleteContentFlag
// @Summary Deletes ContentFlag
// @Accept json
// @Produce json
// @Tags ContentFlag
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /content_flag/{id} [delete]
func (e *ContentFlagController) deleteContentFlag(c echo.Context) error {
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
	var result models.ContentFlag
	query := e.db.QueryContext(models.ContentFlag{}, c)
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

// getContentFlagsBulk godoc
// @Id getContentFlagsBulk
// @Summary Gets ContentFlags in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ContentFlag
// @Success 200 {array} models.ContentFlag
// @Failure 500 {string} string "Bad query request"
// @Router /content_flags/bulk [post]
func (e *ContentFlagController) getContentFlagsBulk(c echo.Context) error {
	var results []models.ContentFlag

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

	err := e.db.QueryContext(models.ContentFlag{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
