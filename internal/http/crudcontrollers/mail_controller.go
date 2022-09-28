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

type MailController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewMailController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *MailController {
	return &MailController{
		db:	 db,
		logger: logger,
	}
}

func (e *MailController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "mail/:msgid", e.getMail, nil),
		routes.RegisterRoute(http.MethodGet, "mail", e.listMail, nil),
		routes.RegisterRoute(http.MethodPut, "mail", e.createMail, nil),
		routes.RegisterRoute(http.MethodDelete, "mail/:msgid", e.deleteMail, nil),
		routes.RegisterRoute(http.MethodPatch, "mail/:msgid", e.updateMail, nil),
		routes.RegisterRoute(http.MethodPost, "mail/bulk", e.getMailBulk, nil),
	}
}

// listMail godoc
// @Id listMail
// @Summary Lists Mail
// @Accept json
// @Produce json
// @Tags Mail
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Mail
// @Failure 500 {string} string "Bad query request"
// @Router /mail [get]
func (e *MailController) listMail(c echo.Context) error {
	var results []models.Mail
	err := e.db.QueryContext(models.Mail{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getMail godoc
// @Id getMail
// @Summary Gets Mail
// @Accept json
// @Produce json
// @Tags Mail
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Mail
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /mail/{id} [get]
func (e *MailController) getMail(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	msgid, err := strconv.Atoi(c.Param("msgid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Msgid]"})
	}
	params = append(params, msgid)
	keys = append(keys, "msgid = ?")

	// query builder
	var result models.Mail
	query := e.db.QueryContext(models.Mail{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Msgid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateMail godoc
// @Id updateMail
// @Summary Updates Mail
// @Accept json
// @Produce json
// @Tags Mail
// @Param id path int true "Id"
// @Param mail body models.Mail true "Mail"
// @Success 200 {array} models.Mail
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /mail/{id} [patch]
func (e *MailController) updateMail(c echo.Context) error {
	request := new(models.Mail)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	msgid, err := strconv.Atoi(c.Param("msgid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Msgid]"})
	}
	params = append(params, msgid)
	keys = append(keys, "msgid = ?")

	// query builder
	var result models.Mail
	query := e.db.QueryContext(models.Mail{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Mail{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createMail godoc
// @Id createMail
// @Summary Creates Mail
// @Accept json
// @Produce json
// @Param mail body models.Mail true "Mail"
// @Tags Mail
// @Success 200 {array} models.Mail
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /mail [put]
func (e *MailController) createMail(c echo.Context) error {
	mail := new(models.Mail)
	if err := c.Bind(mail); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Mail{}, c).Model(&models.Mail{}).Create(&mail).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, mail)
}

// deleteMail godoc
// @Id deleteMail
// @Summary Deletes Mail
// @Accept json
// @Produce json
// @Tags Mail
// @Param id path int true "msgid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /mail/{id} [delete]
func (e *MailController) deleteMail(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	msgid, err := strconv.Atoi(c.Param("msgid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, msgid)
	keys = append(keys, "msgid = ?")

	// query builder
	var result models.Mail
	query := e.db.QueryContext(models.Mail{}, c)
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

// getMailBulk godoc
// @Id getMailBulk
// @Summary Gets Mail in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Mail
// @Success 200 {array} models.Mail
// @Failure 500 {string} string "Bad query request"
// @Router /mail/bulk [post]
func (e *MailController) getMailBulk(c echo.Context) error {
	var results []models.Mail

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

	err := e.db.QueryContext(models.Mail{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
