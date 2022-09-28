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

type IpExemptionController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewIpExemptionController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *IpExemptionController {
	return &IpExemptionController{
		db:	 db,
		logger: logger,
	}
}

func (e *IpExemptionController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "ip_exemption/:exemptionId", e.getIpExemption, nil),
		routes.RegisterRoute(http.MethodGet, "ip_exemptions", e.listIpExemptions, nil),
		routes.RegisterRoute(http.MethodPut, "ip_exemption", e.createIpExemption, nil),
		routes.RegisterRoute(http.MethodDelete, "ip_exemption/:exemptionId", e.deleteIpExemption, nil),
		routes.RegisterRoute(http.MethodPatch, "ip_exemption/:exemptionId", e.updateIpExemption, nil),
		routes.RegisterRoute(http.MethodPost, "ip_exemptions/bulk", e.getIpExemptionsBulk, nil),
	}
}

// listIpExemptions godoc
// @Id listIpExemptions
// @Summary Lists IpExemptions
// @Accept json
// @Produce json
// @Tags IpExemption
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.IpExemption
// @Failure 500 {string} string "Bad query request"
// @Router /ip_exemptions [get]
func (e *IpExemptionController) listIpExemptions(c echo.Context) error {
	var results []models.IpExemption
	err := e.db.QueryContext(models.IpExemption{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getIpExemption godoc
// @Id getIpExemption
// @Summary Gets IpExemption
// @Accept json
// @Produce json
// @Tags IpExemption
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.IpExemption
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /ip_exemption/{id} [get]
func (e *IpExemptionController) getIpExemption(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	exemptionId, err := strconv.Atoi(c.Param("exemptionId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ExemptionId]"})
	}
	params = append(params, exemptionId)
	keys = append(keys, "exemption_id = ?")

	// query builder
	var result models.IpExemption
	query := e.db.QueryContext(models.IpExemption{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ExemptionId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateIpExemption godoc
// @Id updateIpExemption
// @Summary Updates IpExemption
// @Accept json
// @Produce json
// @Tags IpExemption
// @Param id path int true "Id"
// @Param ip_exemption body models.IpExemption true "IpExemption"
// @Success 200 {array} models.IpExemption
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /ip_exemption/{id} [patch]
func (e *IpExemptionController) updateIpExemption(c echo.Context) error {
	request := new(models.IpExemption)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	exemptionId, err := strconv.Atoi(c.Param("exemptionId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ExemptionId]"})
	}
	params = append(params, exemptionId)
	keys = append(keys, "exemption_id = ?")

	// query builder
	var result models.IpExemption
	query := e.db.QueryContext(models.IpExemption{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.IpExemption{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createIpExemption godoc
// @Id createIpExemption
// @Summary Creates IpExemption
// @Accept json
// @Produce json
// @Param ip_exemption body models.IpExemption true "IpExemption"
// @Tags IpExemption
// @Success 200 {array} models.IpExemption
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /ip_exemption [put]
func (e *IpExemptionController) createIpExemption(c echo.Context) error {
	ipExemption := new(models.IpExemption)
	if err := c.Bind(ipExemption); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.IpExemption{}, c).Model(&models.IpExemption{}).Create(&ipExemption).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, ipExemption)
}

// deleteIpExemption godoc
// @Id deleteIpExemption
// @Summary Deletes IpExemption
// @Accept json
// @Produce json
// @Tags IpExemption
// @Param id path int true "exemptionId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /ip_exemption/{id} [delete]
func (e *IpExemptionController) deleteIpExemption(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	exemptionId, err := strconv.Atoi(c.Param("exemptionId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, exemptionId)
	keys = append(keys, "exemption_id = ?")

	// query builder
	var result models.IpExemption
	query := e.db.QueryContext(models.IpExemption{}, c)
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

// getIpExemptionsBulk godoc
// @Id getIpExemptionsBulk
// @Summary Gets IpExemptions in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags IpExemption
// @Success 200 {array} models.IpExemption
// @Failure 500 {string} string "Bad query request"
// @Router /ip_exemptions/bulk [post]
func (e *IpExemptionController) getIpExemptionsBulk(c echo.Context) error {
	var results []models.IpExemption

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

	err := e.db.QueryContext(models.IpExemption{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
