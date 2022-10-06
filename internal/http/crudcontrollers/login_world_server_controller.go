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

type LoginWorldServerController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoginWorldServerController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoginWorldServerController {
	return &LoginWorldServerController{
		db:	    db,
		logger: logger,
	}
}

func (e *LoginWorldServerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "login_world_server/:id", e.getLoginWorldServer, nil),
		routes.RegisterRoute(http.MethodGet, "login_world_servers", e.listLoginWorldServers, nil),
		routes.RegisterRoute(http.MethodPut, "login_world_server", e.createLoginWorldServer, nil),
		routes.RegisterRoute(http.MethodDelete, "login_world_server/:id", e.deleteLoginWorldServer, nil),
		routes.RegisterRoute(http.MethodPatch, "login_world_server/:id", e.updateLoginWorldServer, nil),
		routes.RegisterRoute(http.MethodPost, "login_world_servers/bulk", e.getLoginWorldServersBulk, nil),
	}
}

// listLoginWorldServers godoc
// @Id listLoginWorldServers
// @Summary Lists LoginWorldServers
// @Accept json
// @Produce json
// @Tags LoginWorldServer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginWorldServer
// @Failure 500 {string} string "Bad query request"
// @Router /login_world_servers [get]
func (e *LoginWorldServerController) listLoginWorldServers(c echo.Context) error {
	var results []models.LoginWorldServer
	err := e.db.QueryContext(models.LoginWorldServer{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoginWorldServer godoc
// @Id getLoginWorldServer
// @Summary Gets LoginWorldServer
// @Accept json
// @Produce json
// @Tags LoginWorldServer
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginWorldServer
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /login_world_server/{id} [get]
func (e *LoginWorldServerController) getLoginWorldServer(c echo.Context) error {
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
	var result models.LoginWorldServer
	query := e.db.QueryContext(models.LoginWorldServer{}, c)
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

// updateLoginWorldServer godoc
// @Id updateLoginWorldServer
// @Summary Updates LoginWorldServer
// @Accept json
// @Produce json
// @Tags LoginWorldServer
// @Param id path int true "Id"
// @Param login_world_server body models.LoginWorldServer true "LoginWorldServer"
// @Success 200 {array} models.LoginWorldServer
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /login_world_server/{id} [patch]
func (e *LoginWorldServerController) updateLoginWorldServer(c echo.Context) error {
	request := new(models.LoginWorldServer)
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
	var result models.LoginWorldServer
	query := e.db.QueryContext(models.LoginWorldServer{}, c)
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

// createLoginWorldServer godoc
// @Id createLoginWorldServer
// @Summary Creates LoginWorldServer
// @Accept json
// @Produce json
// @Param login_world_server body models.LoginWorldServer true "LoginWorldServer"
// @Tags LoginWorldServer
// @Success 200 {array} models.LoginWorldServer
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /login_world_server [put]
func (e *LoginWorldServerController) createLoginWorldServer(c echo.Context) error {
	loginWorldServer := new(models.LoginWorldServer)
	if err := c.Bind(loginWorldServer); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LoginWorldServer{}, c).Model(&models.LoginWorldServer{}).Create(&loginWorldServer).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, loginWorldServer)
}

// deleteLoginWorldServer godoc
// @Id deleteLoginWorldServer
// @Summary Deletes LoginWorldServer
// @Accept json
// @Produce json
// @Tags LoginWorldServer
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /login_world_server/{id} [delete]
func (e *LoginWorldServerController) deleteLoginWorldServer(c echo.Context) error {
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
	var result models.LoginWorldServer
	query := e.db.QueryContext(models.LoginWorldServer{}, c)
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

// getLoginWorldServersBulk godoc
// @Id getLoginWorldServersBulk
// @Summary Gets LoginWorldServers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoginWorldServer
// @Success 200 {array} models.LoginWorldServer
// @Failure 500 {string} string "Bad query request"
// @Router /login_world_servers/bulk [post]
func (e *LoginWorldServerController) getLoginWorldServersBulk(c echo.Context) error {
	var results []models.LoginWorldServer

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

	err := e.db.QueryContext(models.LoginWorldServer{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
