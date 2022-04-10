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

type ExpeditionLockoutController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewExpeditionLockoutController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ExpeditionLockoutController {
	return &ExpeditionLockoutController{
		db:	 db,
		logger: logger,
	}
}

func (e *ExpeditionLockoutController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "expedition_lockout/:id", e.getExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodGet, "expedition_lockouts", e.listExpeditionLockouts, nil),
		routes.RegisterRoute(http.MethodPut, "expedition_lockout", e.createExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodDelete, "expedition_lockout/:id", e.deleteExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodPatch, "expedition_lockout/:id", e.updateExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodPost, "expedition_lockouts/bulk", e.getExpeditionLockoutsBulk, nil),
	}
}

// listExpeditionLockouts godoc
// @Id listExpeditionLockouts
// @Summary Lists ExpeditionLockouts
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ExpeditionLockout
// @Failure 500 {string} string "Bad query request"
// @Router /expedition_lockouts [get]
func (e *ExpeditionLockoutController) listExpeditionLockouts(c echo.Context) error {
	var results []models.ExpeditionLockout
	err := e.db.QueryContext(models.ExpeditionLockout{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getExpeditionLockout godoc
// @Id getExpeditionLockout
// @Summary Gets ExpeditionLockout
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ExpeditionLockout
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /expedition_lockout/{id} [get]
func (e *ExpeditionLockoutController) getExpeditionLockout(c echo.Context) error {
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
	var result models.ExpeditionLockout
	query := e.db.QueryContext(models.ExpeditionLockout{}, c)
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

// updateExpeditionLockout godoc
// @Id updateExpeditionLockout
// @Summary Updates ExpeditionLockout
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param id path int true "Id"
// @Param expedition_lockout body models.ExpeditionLockout true "ExpeditionLockout"
// @Success 200 {array} models.ExpeditionLockout
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /expedition_lockout/{id} [patch]
func (e *ExpeditionLockoutController) updateExpeditionLockout(c echo.Context) error {
	request := new(models.ExpeditionLockout)
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
	var result models.ExpeditionLockout
	query := e.db.QueryContext(models.ExpeditionLockout{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.ExpeditionLockout{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createExpeditionLockout godoc
// @Id createExpeditionLockout
// @Summary Creates ExpeditionLockout
// @Accept json
// @Produce json
// @Param expedition_lockout body models.ExpeditionLockout true "ExpeditionLockout"
// @Tags ExpeditionLockout
// @Success 200 {array} models.ExpeditionLockout
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /expedition_lockout [put]
func (e *ExpeditionLockoutController) createExpeditionLockout(c echo.Context) error {
	expeditionLockout := new(models.ExpeditionLockout)
	if err := c.Bind(expeditionLockout); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.ExpeditionLockout{}, c).Model(&models.ExpeditionLockout{}).Create(&expeditionLockout).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, expeditionLockout)
}

// deleteExpeditionLockout godoc
// @Id deleteExpeditionLockout
// @Summary Deletes ExpeditionLockout
// @Accept json
// @Produce json
// @Tags ExpeditionLockout
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /expedition_lockout/{id} [delete]
func (e *ExpeditionLockoutController) deleteExpeditionLockout(c echo.Context) error {
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
	var result models.ExpeditionLockout
	query := e.db.QueryContext(models.ExpeditionLockout{}, c)
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

// getExpeditionLockoutsBulk godoc
// @Id getExpeditionLockoutsBulk
// @Summary Gets ExpeditionLockouts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ExpeditionLockout
// @Success 200 {array} models.ExpeditionLockout
// @Failure 500 {string} string "Bad query request"
// @Router /expedition_lockouts/bulk [post]
func (e *ExpeditionLockoutController) getExpeditionLockoutsBulk(c echo.Context) error {
	var results []models.ExpeditionLockout

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

	err := e.db.QueryContext(models.ExpeditionLockout{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
