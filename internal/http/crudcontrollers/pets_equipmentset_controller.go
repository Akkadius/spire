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
	"net/http"
	"strconv"
	"strings"
)

type PetsEquipmentsetController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewPetsEquipmentsetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *PetsEquipmentsetController {
	return &PetsEquipmentsetController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *PetsEquipmentsetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "pets_equipmentset/:setId", e.getPetsEquipmentset, nil),
		routes.RegisterRoute(http.MethodGet, "pets_equipmentsets", e.listPetsEquipmentsets, nil),
		routes.RegisterRoute(http.MethodGet, "pets_equipmentsets/count", e.getPetsEquipmentsetsCount, nil),
		routes.RegisterRoute(http.MethodPut, "pets_equipmentset", e.createPetsEquipmentset, nil),
		routes.RegisterRoute(http.MethodDelete, "pets_equipmentset/:setId", e.deletePetsEquipmentset, nil),
		routes.RegisterRoute(http.MethodPatch, "pets_equipmentset/:setId", e.updatePetsEquipmentset, nil),
		routes.RegisterRoute(http.MethodPost, "pets_equipmentsets/bulk", e.getPetsEquipmentsetsBulk, nil),
	}
}

// listPetsEquipmentsets godoc
// @Id listPetsEquipmentsets
// @Summary Lists PetsEquipmentsets
// @Accept json
// @Produce json
// @Tags PetsEquipmentset
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsEquipmentset
// @Failure 500 {string} string "Bad query request"
// @Router /pets_equipmentsets [get]
func (e *PetsEquipmentsetController) listPetsEquipmentsets(c echo.Context) error {
	var results []models.PetsEquipmentset
	err := e.db.QueryContext(models.PetsEquipmentset{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getPetsEquipmentset godoc
// @Id getPetsEquipmentset
// @Summary Gets PetsEquipmentset
// @Accept json
// @Produce json
// @Tags PetsEquipmentset
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsEquipmentset
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /pets_equipmentset/{id} [get]
func (e *PetsEquipmentsetController) getPetsEquipmentset(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	setId, err := strconv.Atoi(c.Param("setId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SetId]"})
	}
	params = append(params, setId)
	keys = append(keys, "set_id = ?")

	// query builder
	var result models.PetsEquipmentset
	query := e.db.QueryContext(models.PetsEquipmentset{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.SetId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updatePetsEquipmentset godoc
// @Id updatePetsEquipmentset
// @Summary Updates PetsEquipmentset
// @Accept json
// @Produce json
// @Tags PetsEquipmentset
// @Param id path int true "Id"
// @Param pets_equipmentset body models.PetsEquipmentset true "PetsEquipmentset"
// @Success 200 {array} models.PetsEquipmentset
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /pets_equipmentset/{id} [patch]
func (e *PetsEquipmentsetController) updatePetsEquipmentset(c echo.Context) error {
	request := new(models.PetsEquipmentset)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	setId, err := strconv.Atoi(c.Param("setId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SetId]"})
	}
	params = append(params, setId)
	keys = append(keys, "set_id = ?")

	// query builder
	var result models.PetsEquipmentset
	query := e.db.QueryContext(models.PetsEquipmentset{}, c)
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
		event := fmt.Sprintf("Updated [PetsEquipmentset] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createPetsEquipmentset godoc
// @Id createPetsEquipmentset
// @Summary Creates PetsEquipmentset
// @Accept json
// @Produce json
// @Param pets_equipmentset body models.PetsEquipmentset true "PetsEquipmentset"
// @Tags PetsEquipmentset
// @Success 200 {array} models.PetsEquipmentset
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /pets_equipmentset [put]
func (e *PetsEquipmentsetController) createPetsEquipmentset(c echo.Context) error {
	petsEquipmentset := new(models.PetsEquipmentset)
	if err := c.Bind(petsEquipmentset); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.PetsEquipmentset{}, c).Model(&models.PetsEquipmentset{}).Create(&petsEquipmentset).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.PetsEquipmentset{}, petsEquipmentset)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [PetsEquipmentset] [%v] data [%v]", petsEquipmentset.SetId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, petsEquipmentset)
}

// deletePetsEquipmentset godoc
// @Id deletePetsEquipmentset
// @Summary Deletes PetsEquipmentset
// @Accept json
// @Produce json
// @Tags PetsEquipmentset
// @Param id path int true "setId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /pets_equipmentset/{id} [delete]
func (e *PetsEquipmentsetController) deletePetsEquipmentset(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	setId, err := strconv.Atoi(c.Param("setId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, setId)
	keys = append(keys, "set_id = ?")

	// query builder
	var result models.PetsEquipmentset
	query := e.db.QueryContext(models.PetsEquipmentset{}, c)
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
		event := fmt.Sprintf("Deleted [PetsEquipmentset] [%v] keys [%v]", result.SetId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getPetsEquipmentsetsBulk godoc
// @Id getPetsEquipmentsetsBulk
// @Summary Gets PetsEquipmentsets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PetsEquipmentset
// @Success 200 {array} models.PetsEquipmentset
// @Failure 500 {string} string "Bad query request"
// @Router /pets_equipmentsets/bulk [post]
func (e *PetsEquipmentsetController) getPetsEquipmentsetsBulk(c echo.Context) error {
	var results []models.PetsEquipmentset

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

	err := e.db.QueryContext(models.PetsEquipmentset{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPetsEquipmentsetsCount godoc
// @Id getPetsEquipmentsetsCount
// @Summary Counts PetsEquipmentsets
// @Accept json
// @Produce json
// @Tags PetsEquipmentset
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsEquipmentset
// @Failure 500 {string} string "Bad query request"
// @Router /pets_equipmentsets/count [get]
func (e *PetsEquipmentsetController) getPetsEquipmentsetsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.PetsEquipmentset{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}