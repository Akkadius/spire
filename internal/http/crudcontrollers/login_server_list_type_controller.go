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

type LoginServerListTypeController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoginServerListTypeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoginServerListTypeController {
	return &LoginServerListTypeController{
		db:	 db,
		logger: logger,
	}
}

func (e *LoginServerListTypeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "login_server_list_type/:id", e.getLoginServerListType, nil),
		routes.RegisterRoute(http.MethodGet, "login_server_list_types", e.listLoginServerListTypes, nil),
		routes.RegisterRoute(http.MethodPut, "login_server_list_type", e.createLoginServerListType, nil),
		routes.RegisterRoute(http.MethodDelete, "login_server_list_type/:id", e.deleteLoginServerListType, nil),
		routes.RegisterRoute(http.MethodPatch, "login_server_list_type/:id", e.updateLoginServerListType, nil),
		routes.RegisterRoute(http.MethodPost, "login_server_list_types/bulk", e.getLoginServerListTypesBulk, nil),
	}
}

// listLoginServerListTypes godoc
// @Id listLoginServerListTypes
// @Summary Lists LoginServerListTypes
// @Accept json
// @Produce json
// @Tags LoginServerListType
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginServerListType
// @Failure 500 {string} string "Bad query request"
// @Router /login_server_list_types [get]
func (e *LoginServerListTypeController) listLoginServerListTypes(c echo.Context) error {
	var results []models.LoginServerListType
	err := e.db.QueryContext(models.LoginServerListType{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoginServerListType godoc
// @Id getLoginServerListType
// @Summary Gets LoginServerListType
// @Accept json
// @Produce json
// @Tags LoginServerListType
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginServerListType
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /login_server_list_type/{id} [get]
func (e *LoginServerListTypeController) getLoginServerListType(c echo.Context) error {
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
	var result models.LoginServerListType
	query := e.db.QueryContext(models.LoginServerListType{}, c)
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

// updateLoginServerListType godoc
// @Id updateLoginServerListType
// @Summary Updates LoginServerListType
// @Accept json
// @Produce json
// @Tags LoginServerListType
// @Param id path int true "Id"
// @Param login_server_list_type body models.LoginServerListType true "LoginServerListType"
// @Success 200 {array} models.LoginServerListType
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /login_server_list_type/{id} [patch]
func (e *LoginServerListTypeController) updateLoginServerListType(c echo.Context) error {
	request := new(models.LoginServerListType)
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
	var result models.LoginServerListType
	query := e.db.QueryContext(models.LoginServerListType{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.LoginServerListType{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createLoginServerListType godoc
// @Id createLoginServerListType
// @Summary Creates LoginServerListType
// @Accept json
// @Produce json
// @Param login_server_list_type body models.LoginServerListType true "LoginServerListType"
// @Tags LoginServerListType
// @Success 200 {array} models.LoginServerListType
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /login_server_list_type [put]
func (e *LoginServerListTypeController) createLoginServerListType(c echo.Context) error {
	loginServerListType := new(models.LoginServerListType)
	if err := c.Bind(loginServerListType); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LoginServerListType{}, c).Model(&models.LoginServerListType{}).Create(&loginServerListType).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, loginServerListType)
}

// deleteLoginServerListType godoc
// @Id deleteLoginServerListType
// @Summary Deletes LoginServerListType
// @Accept json
// @Produce json
// @Tags LoginServerListType
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /login_server_list_type/{id} [delete]
func (e *LoginServerListTypeController) deleteLoginServerListType(c echo.Context) error {
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
	var result models.LoginServerListType
	query := e.db.QueryContext(models.LoginServerListType{}, c)
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

// getLoginServerListTypesBulk godoc
// @Id getLoginServerListTypesBulk
// @Summary Gets LoginServerListTypes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoginServerListType
// @Success 200 {array} models.LoginServerListType
// @Failure 500 {string} string "Bad query request"
// @Router /login_server_list_types/bulk [post]
func (e *LoginServerListTypeController) getLoginServerListTypesBulk(c echo.Context) error {
	var results []models.LoginServerListType

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

	err := e.db.QueryContext(models.LoginServerListType{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
