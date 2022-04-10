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

type AdventureTemplateController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAdventureTemplateController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AdventureTemplateController {
	return &AdventureTemplateController{
		db:	 db,
		logger: logger,
	}
}

func (e *AdventureTemplateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "adventure_template/:id", e.getAdventureTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "adventure_templates", e.listAdventureTemplates, nil),
		routes.RegisterRoute(http.MethodPut, "adventure_template", e.createAdventureTemplate, nil),
		routes.RegisterRoute(http.MethodDelete, "adventure_template/:id", e.deleteAdventureTemplate, nil),
		routes.RegisterRoute(http.MethodPatch, "adventure_template/:id", e.updateAdventureTemplate, nil),
		routes.RegisterRoute(http.MethodPost, "adventure_templates/bulk", e.getAdventureTemplatesBulk, nil),
	}
}

// listAdventureTemplates godoc
// @Id listAdventureTemplates
// @Summary Lists AdventureTemplates
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_templates [get]
func (e *AdventureTemplateController) listAdventureTemplates(c echo.Context) error {
	var results []models.AdventureTemplate
	err := e.db.QueryContext(models.AdventureTemplate{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAdventureTemplate godoc
// @Id getAdventureTemplate
// @Summary Gets AdventureTemplate
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AdventureTemplate
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_template/{id} [get]
func (e *AdventureTemplateController) getAdventureTemplate(c echo.Context) error {
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
	var result models.AdventureTemplate
	query := e.db.QueryContext(models.AdventureTemplate{}, c)
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

// updateAdventureTemplate godoc
// @Id updateAdventureTemplate
// @Summary Updates AdventureTemplate
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param id path int true "Id"
// @Param adventure_template body models.AdventureTemplate true "AdventureTemplate"
// @Success 200 {array} models.AdventureTemplate
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /adventure_template/{id} [patch]
func (e *AdventureTemplateController) updateAdventureTemplate(c echo.Context) error {
	request := new(models.AdventureTemplate)
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
	var result models.AdventureTemplate
	query := e.db.QueryContext(models.AdventureTemplate{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AdventureTemplate{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAdventureTemplate godoc
// @Id createAdventureTemplate
// @Summary Creates AdventureTemplate
// @Accept json
// @Produce json
// @Param adventure_template body models.AdventureTemplate true "AdventureTemplate"
// @Tags AdventureTemplate
// @Success 200 {array} models.AdventureTemplate
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /adventure_template [put]
func (e *AdventureTemplateController) createAdventureTemplate(c echo.Context) error {
	adventureTemplate := new(models.AdventureTemplate)
	if err := c.Bind(adventureTemplate); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AdventureTemplate{}, c).Model(&models.AdventureTemplate{}).Create(&adventureTemplate).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, adventureTemplate)
}

// deleteAdventureTemplate godoc
// @Id deleteAdventureTemplate
// @Summary Deletes AdventureTemplate
// @Accept json
// @Produce json
// @Tags AdventureTemplate
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /adventure_template/{id} [delete]
func (e *AdventureTemplateController) deleteAdventureTemplate(c echo.Context) error {
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
	var result models.AdventureTemplate
	query := e.db.QueryContext(models.AdventureTemplate{}, c)
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

// getAdventureTemplatesBulk godoc
// @Id getAdventureTemplatesBulk
// @Summary Gets AdventureTemplates in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AdventureTemplate
// @Success 200 {array} models.AdventureTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /adventure_templates/bulk [post]
func (e *AdventureTemplateController) getAdventureTemplatesBulk(c echo.Context) error {
	var results []models.AdventureTemplate

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

	err := e.db.QueryContext(models.AdventureTemplate{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
