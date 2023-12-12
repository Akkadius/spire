package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type MerchantlistTempController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewMerchantlistTempController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *MerchantlistTempController {
	return &MerchantlistTempController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *MerchantlistTempController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "merchantlist_temp/:npcid", e.getMerchantlistTemp, nil),
		routes.RegisterRoute(http.MethodGet, "merchantlist_temps", e.listMerchantlistTemps, nil),
		routes.RegisterRoute(http.MethodGet, "merchantlist_temps/count", e.getMerchantlistTempsCount, nil),
		routes.RegisterRoute(http.MethodPut, "merchantlist_temp", e.createMerchantlistTemp, nil),
		routes.RegisterRoute(http.MethodDelete, "merchantlist_temp/:npcid", e.deleteMerchantlistTemp, nil),
		routes.RegisterRoute(http.MethodPatch, "merchantlist_temp/:npcid", e.updateMerchantlistTemp, nil),
		routes.RegisterRoute(http.MethodPost, "merchantlist_temps/bulk", e.getMerchantlistTempsBulk, nil),
	}
}

// listMerchantlistTemps godoc
// @Id listMerchantlistTemps
// @Summary Lists MerchantlistTemps
// @Accept json
// @Produce json
// @Tags MerchantlistTemp
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.MerchantlistTemp
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlist_temps [get]
func (e *MerchantlistTempController) listMerchantlistTemps(c echo.Context) error {
	var results []models.MerchantlistTemp
	err := e.db.QueryContext(models.MerchantlistTemp{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getMerchantlistTemp godoc
// @Id getMerchantlistTemp
// @Summary Gets MerchantlistTemp
// @Accept json
// @Produce json
// @Tags MerchantlistTemp
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.MerchantlistTemp
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlist_temp/{id} [get]
func (e *MerchantlistTempController) getMerchantlistTemp(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	npcid, err := strconv.Atoi(c.Param("npcid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Npcid]"})
	}
	params = append(params, npcid)
	keys = append(keys, "npcid = ?")

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// key param [zone_id] position [3] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// key param [instance_id] position [4] type [int]
	if len(c.QueryParam("instance_id")) > 0 {
		instanceIdParam, err := strconv.Atoi(c.QueryParam("instance_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [instance_id] err [%s]", err.Error())})
		}

		params = append(params, instanceIdParam)
		keys = append(keys, "instance_id = ?")
	}

	// query builder
	var result models.MerchantlistTemp
	query := e.db.QueryContext(models.MerchantlistTemp{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Npcid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateMerchantlistTemp godoc
// @Id updateMerchantlistTemp
// @Summary Updates MerchantlistTemp
// @Accept json
// @Produce json
// @Tags MerchantlistTemp
// @Param id path int true "Id"
// @Param merchantlist_temp body models.MerchantlistTemp true "MerchantlistTemp"
// @Success 200 {array} models.MerchantlistTemp
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /merchantlist_temp/{id} [patch]
func (e *MerchantlistTempController) updateMerchantlistTemp(c echo.Context) error {
	request := new(models.MerchantlistTemp)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	npcid, err := strconv.Atoi(c.Param("npcid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Npcid]"})
	}
	params = append(params, npcid)
	keys = append(keys, "npcid = ?")

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// key param [zone_id] position [3] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// key param [instance_id] position [4] type [int]
	if len(c.QueryParam("instance_id")) > 0 {
		instanceIdParam, err := strconv.Atoi(c.QueryParam("instance_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [instance_id] err [%s]", err.Error())})
		}

		params = append(params, instanceIdParam)
		keys = append(keys, "instance_id = ?")
	}

	// query builder
	var result models.MerchantlistTemp
	query := e.db.QueryContext(models.MerchantlistTemp{}, c)
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
		event := fmt.Sprintf("Updated [MerchantlistTemp] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createMerchantlistTemp godoc
// @Id createMerchantlistTemp
// @Summary Creates MerchantlistTemp
// @Accept json
// @Produce json
// @Param merchantlist_temp body models.MerchantlistTemp true "MerchantlistTemp"
// @Tags MerchantlistTemp
// @Success 200 {array} models.MerchantlistTemp
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /merchantlist_temp [put]
func (e *MerchantlistTempController) createMerchantlistTemp(c echo.Context) error {
	merchantlistTemp := new(models.MerchantlistTemp)
	if err := c.Bind(merchantlistTemp); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.MerchantlistTemp{}, c).Model(&models.MerchantlistTemp{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&merchantlistTemp).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.MerchantlistTemp{}, merchantlistTemp)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [MerchantlistTemp] [%v] data [%v]", merchantlistTemp.Npcid, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, merchantlistTemp)
}

// deleteMerchantlistTemp godoc
// @Id deleteMerchantlistTemp
// @Summary Deletes MerchantlistTemp
// @Accept json
// @Produce json
// @Tags MerchantlistTemp
// @Param id path int true "npcid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /merchantlist_temp/{id} [delete]
func (e *MerchantlistTempController) deleteMerchantlistTemp(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	npcid, err := strconv.Atoi(c.Param("npcid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, npcid)
	keys = append(keys, "npcid = ?")

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// key param [zone_id] position [3] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// key param [instance_id] position [4] type [int]
	if len(c.QueryParam("instance_id")) > 0 {
		instanceIdParam, err := strconv.Atoi(c.QueryParam("instance_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [instance_id] err [%s]", err.Error())})
		}

		params = append(params, instanceIdParam)
		keys = append(keys, "instance_id = ?")
	}

	// query builder
	var result models.MerchantlistTemp
	query := e.db.QueryContext(models.MerchantlistTemp{}, c)
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
		event := fmt.Sprintf("Deleted [MerchantlistTemp] [%v] keys [%v]", result.Npcid, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getMerchantlistTempsBulk godoc
// @Id getMerchantlistTempsBulk
// @Summary Gets MerchantlistTemps in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags MerchantlistTemp
// @Success 200 {array} models.MerchantlistTemp
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlist_temps/bulk [post]
func (e *MerchantlistTempController) getMerchantlistTempsBulk(c echo.Context) error {
	var results []models.MerchantlistTemp

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

	err := e.db.QueryContext(models.MerchantlistTemp{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getMerchantlistTempsCount godoc
// @Id getMerchantlistTempsCount
// @Summary Counts MerchantlistTemps
// @Accept json
// @Produce json
// @Tags MerchantlistTemp
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.MerchantlistTemp
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlist_temps/count [get]
func (e *MerchantlistTempController) getMerchantlistTempsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.MerchantlistTemp{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}