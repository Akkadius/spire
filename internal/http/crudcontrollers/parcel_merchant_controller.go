package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type ParcelMerchantController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewParcelMerchantController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *ParcelMerchantController {
	return &ParcelMerchantController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *ParcelMerchantController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "parcel_merchant/:id", e.getParcelMerchant, nil),
		routes.RegisterRoute(http.MethodGet, "parcel_merchants", e.listParcelMerchants, nil),
		routes.RegisterRoute(http.MethodGet, "parcel_merchants/count", e.getParcelMerchantsCount, nil),
		routes.RegisterRoute(http.MethodPut, "parcel_merchant", e.createParcelMerchant, nil),
		routes.RegisterRoute(http.MethodDelete, "parcel_merchant/:id", e.deleteParcelMerchant, nil),
		routes.RegisterRoute(http.MethodPatch, "parcel_merchant/:id", e.updateParcelMerchant, nil),
		routes.RegisterRoute(http.MethodPost, "parcel_merchants/bulk", e.getParcelMerchantsBulk, nil),
	}
}

// listParcelMerchants godoc
// @Id listParcelMerchants
// @Summary Lists ParcelMerchants
// @Accept json
// @Produce json
// @Tags ParcelMerchant
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ParcelMerchant
// @Failure 500 {string} string "Bad query request"
// @Router /parcel_merchants [get]
func (e *ParcelMerchantController) listParcelMerchants(c echo.Context) error {
	var results []models.ParcelMerchant
	err := e.db.QueryContext(models.ParcelMerchant{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getParcelMerchant godoc
// @Id getParcelMerchant
// @Summary Gets ParcelMerchant
// @Accept json
// @Produce json
// @Tags ParcelMerchant
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ParcelMerchant
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /parcel_merchant/{id} [get]
func (e *ParcelMerchantController) getParcelMerchant(c echo.Context) error {
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
	var result models.ParcelMerchant
	query := e.db.QueryContext(models.ParcelMerchant{}, c)
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

// updateParcelMerchant godoc
// @Id updateParcelMerchant
// @Summary Updates ParcelMerchant
// @Accept json
// @Produce json
// @Tags ParcelMerchant
// @Param id path int true "Id"
// @Param parcel_merchant body models.ParcelMerchant true "ParcelMerchant"
// @Success 200 {array} models.ParcelMerchant
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /parcel_merchant/{id} [patch]
func (e *ParcelMerchantController) updateParcelMerchant(c echo.Context) error {
	request := new(models.ParcelMerchant)
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
	var result models.ParcelMerchant
	query := e.db.QueryContext(models.ParcelMerchant{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [ParcelMerchant] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createParcelMerchant godoc
// @Id createParcelMerchant
// @Summary Creates ParcelMerchant
// @Accept json
// @Produce json
// @Param parcel_merchant body models.ParcelMerchant true "ParcelMerchant"
// @Tags ParcelMerchant
// @Success 200 {array} models.ParcelMerchant
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /parcel_merchant [put]
func (e *ParcelMerchantController) createParcelMerchant(c echo.Context) error {
	parcelMerchant := new(models.ParcelMerchant)
	if err := c.Bind(parcelMerchant); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.ParcelMerchant{}, c).Model(&models.ParcelMerchant{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&parcelMerchant).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.ParcelMerchant{}, parcelMerchant)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [ParcelMerchant] [%v] data [%v]", parcelMerchant.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, parcelMerchant)
}

// deleteParcelMerchant godoc
// @Id deleteParcelMerchant
// @Summary Deletes ParcelMerchant
// @Accept json
// @Produce json
// @Tags ParcelMerchant
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /parcel_merchant/{id} [delete]
func (e *ParcelMerchantController) deleteParcelMerchant(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.ParcelMerchant
	query := e.db.QueryContext(models.ParcelMerchant{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [ParcelMerchant] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getParcelMerchantsBulk godoc
// @Id getParcelMerchantsBulk
// @Summary Gets ParcelMerchants in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags ParcelMerchant
// @Success 200 {array} models.ParcelMerchant
// @Failure 500 {string} string "Bad query request"
// @Router /parcel_merchants/bulk [post]
func (e *ParcelMerchantController) getParcelMerchantsBulk(c echo.Context) error {
	var results []models.ParcelMerchant

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

	err := e.db.QueryContext(models.ParcelMerchant{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getParcelMerchantsCount godoc
// @Id getParcelMerchantsCount
// @Summary Counts ParcelMerchants
// @Accept json
// @Produce json
// @Tags ParcelMerchant
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ParcelMerchant
// @Failure 500 {string} string "Bad query request"
// @Router /parcel_merchants/count [get]
func (e *ParcelMerchantController) getParcelMerchantsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.ParcelMerchant{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}