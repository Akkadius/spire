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

type LoginApiTokenController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoginApiTokenController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoginApiTokenController {
	return &LoginApiTokenController{
		db:     db,
		logger: logger,
	}
}

func (e *LoginApiTokenController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "login_api_token/:login_api_token", e.deleteLoginApiToken, nil),
		routes.RegisterRoute(http.MethodGet, "login_api_token/:login_api_token", e.getLoginApiToken, nil),
		routes.RegisterRoute(http.MethodGet, "login_api_tokens", e.listLoginApiTokens, nil),
		routes.RegisterRoute(http.MethodPost, "login_api_tokens/bulk", e.getLoginApiTokensBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "login_api_token/:login_api_token", e.updateLoginApiToken, nil),
		routes.RegisterRoute(http.MethodPut, "login_api_token", e.createLoginApiToken, nil),
	}
}

// listLoginApiTokens godoc
// @Id listLoginApiTokens
// @Summary Lists LoginApiTokens
// @Accept json
// @Produce json
// @Tags LoginApiToken
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginApiToken
// @Failure 500 {string} string "Bad query request"
// @Router /login_api_tokens [get]
func (e *LoginApiTokenController) listLoginApiTokens(c echo.Context) error {
	var results []models.LoginApiToken
	err := e.db.QueryContext(models.LoginApiToken{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoginApiToken godoc
// @Id getLoginApiToken
// @Summary Gets LoginApiToken
// @Accept json
// @Produce json
// @Tags LoginApiToken
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoginApiToken
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /login_api_token/{id} [get]
func (e *LoginApiTokenController) getLoginApiToken(c echo.Context) error {
	loginApiTokenId, err := strconv.Atoi(c.Param("login_api_token"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.LoginApiToken
	err = e.db.QueryContext(models.LoginApiToken{}, c).First(&result, loginApiTokenId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLoginApiToken godoc
// @Id updateLoginApiToken
// @Summary Updates LoginApiToken
// @Accept json
// @Produce json
// @Tags LoginApiToken
// @Param id path int true "Id"
// @Param login_api_token body models.LoginApiToken true "LoginApiToken"
// @Success 200 {array} models.LoginApiToken
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /login_api_token/{id} [patch]
func (e *LoginApiTokenController) updateLoginApiToken(c echo.Context) error {
	loginApiToken := new(models.LoginApiToken)
	if err := c.Bind(loginApiToken); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.LoginApiToken{}
	err := e.db.Get(models.LoginApiToken{}, c).Model(&models.LoginApiToken{}).First(&entity, loginApiToken.ID).Error
	if err != nil || loginApiToken.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LoginApiToken{}, c).Model(&entity).Select("*").Updates(&loginApiToken).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, loginApiToken)
}

// createLoginApiToken godoc
// @Id createLoginApiToken
// @Summary Creates LoginApiToken
// @Accept json
// @Produce json
// @Param login_api_token body models.LoginApiToken true "LoginApiToken"
// @Tags LoginApiToken
// @Success 200 {array} models.LoginApiToken
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /login_api_token [put]
func (e *LoginApiTokenController) createLoginApiToken(c echo.Context) error {
	loginApiToken := new(models.LoginApiToken)
	if err := c.Bind(loginApiToken); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.LoginApiToken{}, c).Model(&models.LoginApiToken{}).Create(&loginApiToken).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, loginApiToken)
}

// deleteLoginApiToken godoc
// @Id deleteLoginApiToken
// @Summary Deletes LoginApiToken
// @Accept json
// @Produce json
// @Tags LoginApiToken
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /login_api_token/{id} [delete]
func (e *LoginApiTokenController) deleteLoginApiToken(c echo.Context) error {
	loginApiTokenId, err := strconv.Atoi(c.Param("login_api_token"))
	if err != nil {
		e.logger.Error(err)
	}

	loginApiToken := new(models.LoginApiToken)
	err = e.db.Get(models.LoginApiToken{}, c).Model(&models.LoginApiToken{}).First(&loginApiToken, loginApiTokenId).Error
	if err != nil || loginApiToken.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.LoginApiToken{}, c).Model(&models.LoginApiToken{}).Delete(&loginApiToken).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLoginApiTokensBulk godoc
// @Id getLoginApiTokensBulk
// @Summary Gets LoginApiTokens in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoginApiToken
// @Success 200 {array} models.LoginApiToken
// @Failure 500 {string} string "Bad query request"
// @Router /login_api_tokens/bulk [post]
func (e *LoginApiTokenController) getLoginApiTokensBulk(c echo.Context) error {
	var results []models.LoginApiToken

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

	err := e.db.QueryContext(models.LoginApiToken{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
