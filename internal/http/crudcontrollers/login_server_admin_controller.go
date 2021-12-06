package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type LoginServerAdminController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoginServerAdminController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoginServerAdminController {
	return &LoginServerAdminController{
		db:     db,
		logger: logger,
	}
}

func (e *LoginServerAdminController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "login_server_admin/:login_server_admin", e.deleteLoginServerAdmin, nil),
		routes.RegisterRoute(http.MethodGet, "login_server_admin/:login_server_admin", e.getLoginServerAdmin, nil),
		routes.RegisterRoute(http.MethodGet, "login_server_admins", e.listLoginServerAdmins, nil),
		routes.RegisterRoute(http.MethodPost, "login_server_admins/bulk", e.getLoginServerAdminsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "login_server_admin/:login_server_admin", e.updateLoginServerAdmin, nil),
		routes.RegisterRoute(http.MethodPut, "login_server_admin", e.createLoginServerAdmin, nil),
	}
}

// listLoginServerAdmins godoc
// @Id listLoginServerAdmins
// @Summary Lists LoginServerAdmins
// @Accept json
// @Produce json
// @Tags LoginServerAdmin
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginServerAdmin
// @Failure 500 {string} string "Bad query request"
// @Router /login_server_admins [get]
func (e *LoginServerAdminController) listLoginServerAdmins(c echo.Context) error {
	var results []models.LoginServerAdmin
	err := e.db.QueryContext(models.LoginServerAdmin{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoginServerAdmin godoc
// @Id getLoginServerAdmin
// @Summary Gets LoginServerAdmin
// @Accept json
// @Produce json
// @Tags LoginServerAdmin
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginServerAdmin
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /login_server_admin/{id} [get]
func (e *LoginServerAdminController) getLoginServerAdmin(c echo.Context) error {
	loginServerAdminId, err := strconv.Atoi(c.Param("login_server_admin"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.LoginServerAdmin
	err = e.db.QueryContext(models.LoginServerAdmin{}, c).First(&result, loginServerAdminId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLoginServerAdmin godoc
// @Id updateLoginServerAdmin
// @Summary Updates LoginServerAdmin
// @Accept json
// @Produce json
// @Tags LoginServerAdmin
// @Param id path int true "Id"
// @Param login_server_admin body models.LoginServerAdmin true "LoginServerAdmin"
// @Success 200 {array} models.LoginServerAdmin
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /login_server_admin/{id} [patch]
func (e *LoginServerAdminController) updateLoginServerAdmin(c echo.Context) error {
	loginServerAdmin := new(models.LoginServerAdmin)
	if err := c.Bind(loginServerAdmin); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.LoginServerAdmin{}
	err := e.db.Get(models.LoginServerAdmin{}, c).Model(&models.LoginServerAdmin{}).First(&entity, loginServerAdmin.ID).Error
	if err != nil || loginServerAdmin.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LoginServerAdmin{}, c).Model(&entity).Select("*").Updates(&loginServerAdmin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, loginServerAdmin)
}

// createLoginServerAdmin godoc
// @Id createLoginServerAdmin
// @Summary Creates LoginServerAdmin
// @Accept json
// @Produce json
// @Param login_server_admin body models.LoginServerAdmin true "LoginServerAdmin"
// @Tags LoginServerAdmin
// @Success 200 {array} models.LoginServerAdmin
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /login_server_admin [put]
func (e *LoginServerAdminController) createLoginServerAdmin(c echo.Context) error {
	loginServerAdmin := new(models.LoginServerAdmin)
	if err := c.Bind(loginServerAdmin); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.LoginServerAdmin{}, c).Model(&models.LoginServerAdmin{}).Create(&loginServerAdmin).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, loginServerAdmin)
}

// deleteLoginServerAdmin godoc
// @Id deleteLoginServerAdmin
// @Summary Deletes LoginServerAdmin
// @Accept json
// @Produce json
// @Tags LoginServerAdmin
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /login_server_admin/{id} [delete]
func (e *LoginServerAdminController) deleteLoginServerAdmin(c echo.Context) error {
	loginServerAdminId, err := strconv.Atoi(c.Param("login_server_admin"))
	if err != nil {
		e.logger.Error(err)
	}

	loginServerAdmin := new(models.LoginServerAdmin)
	err = e.db.Get(models.LoginServerAdmin{}, c).Model(&models.LoginServerAdmin{}).First(&loginServerAdmin, loginServerAdminId).Error
	if err != nil || loginServerAdmin.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LoginServerAdmin{}, c).Model(&models.LoginServerAdmin{}).Delete(&loginServerAdmin).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLoginServerAdminsBulk godoc
// @Id getLoginServerAdminsBulk
// @Summary Gets LoginServerAdmins in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoginServerAdmin
// @Success 200 {array} models.LoginServerAdmin
// @Failure 500 {string} string "Bad query request"
// @Router /login_server_admins/bulk [post]
func (e *LoginServerAdminController) getLoginServerAdminsBulk(c echo.Context) error {
	var results []models.LoginServerAdmin

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.LoginServerAdmin{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
