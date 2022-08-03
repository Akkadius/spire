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

type PetsEquipmentsetEntryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewPetsEquipmentsetEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *PetsEquipmentsetEntryController {
	return &PetsEquipmentsetEntryController{
		db:	 db,
		logger: logger,
	}
}

func (e *PetsEquipmentsetEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "pets_equipmentset_entry/:setId", e.getPetsEquipmentsetEntry, nil),
		routes.RegisterRoute(http.MethodGet, "pets_equipmentset_entries", e.listPetsEquipmentsetEntries, nil),
		routes.RegisterRoute(http.MethodPut, "pets_equipmentset_entry", e.createPetsEquipmentsetEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "pets_equipmentset_entry/:setId", e.deletePetsEquipmentsetEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "pets_equipmentset_entry/:setId", e.updatePetsEquipmentsetEntry, nil),
		routes.RegisterRoute(http.MethodPost, "pets_equipmentset_entries/bulk", e.getPetsEquipmentsetEntriesBulk, nil),
	}
}

// listPetsEquipmentsetEntries godoc
// @Id listPetsEquipmentsetEntries
// @Summary Lists PetsEquipmentsetEntries
// @Accept json
// @Produce json
// @Tags PetsEquipmentsetEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsEquipmentsetEntry
// @Failure 500 {string} string "Bad query request"
// @Router /pets_equipmentset_entries [get]
func (e *PetsEquipmentsetEntryController) listPetsEquipmentsetEntries(c echo.Context) error {
	var results []models.PetsEquipmentsetEntry
	err := e.db.QueryContext(models.PetsEquipmentsetEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getPetsEquipmentsetEntry godoc
// @Id getPetsEquipmentsetEntry
// @Summary Gets PetsEquipmentsetEntry
// @Accept json
// @Produce json
// @Tags PetsEquipmentsetEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsEquipmentsetEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /pets_equipmentset_entry/{id} [get]
func (e *PetsEquipmentsetEntryController) getPetsEquipmentsetEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	setId, err := strconv.Atoi(c.Param("setId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SetId]"})
	}
	params = append(params, setId)
	keys = append(keys, "set_id = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.PetsEquipmentsetEntry
	query := e.db.QueryContext(models.PetsEquipmentsetEntry{}, c)
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

// updatePetsEquipmentsetEntry godoc
// @Id updatePetsEquipmentsetEntry
// @Summary Updates PetsEquipmentsetEntry
// @Accept json
// @Produce json
// @Tags PetsEquipmentsetEntry
// @Param id path int true "Id"
// @Param pets_equipmentset_entry body models.PetsEquipmentsetEntry true "PetsEquipmentsetEntry"
// @Success 200 {array} models.PetsEquipmentsetEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /pets_equipmentset_entry/{id} [patch]
func (e *PetsEquipmentsetEntryController) updatePetsEquipmentsetEntry(c echo.Context) error {
	request := new(models.PetsEquipmentsetEntry)
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

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.PetsEquipmentsetEntry
	query := e.db.QueryContext(models.PetsEquipmentsetEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.PetsEquipmentsetEntry{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createPetsEquipmentsetEntry godoc
// @Id createPetsEquipmentsetEntry
// @Summary Creates PetsEquipmentsetEntry
// @Accept json
// @Produce json
// @Param pets_equipmentset_entry body models.PetsEquipmentsetEntry true "PetsEquipmentsetEntry"
// @Tags PetsEquipmentsetEntry
// @Success 200 {array} models.PetsEquipmentsetEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /pets_equipmentset_entry [put]
func (e *PetsEquipmentsetEntryController) createPetsEquipmentsetEntry(c echo.Context) error {
	petsEquipmentsetEntry := new(models.PetsEquipmentsetEntry)
	if err := c.Bind(petsEquipmentsetEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.PetsEquipmentsetEntry{}, c).Model(&models.PetsEquipmentsetEntry{}).Create(&petsEquipmentsetEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, petsEquipmentsetEntry)
}

// deletePetsEquipmentsetEntry godoc
// @Id deletePetsEquipmentsetEntry
// @Summary Deletes PetsEquipmentsetEntry
// @Accept json
// @Produce json
// @Tags PetsEquipmentsetEntry
// @Param id path int true "setId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /pets_equipmentset_entry/{id} [delete]
func (e *PetsEquipmentsetEntryController) deletePetsEquipmentsetEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	setId, err := strconv.Atoi(c.Param("setId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, setId)
	keys = append(keys, "set_id = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.PetsEquipmentsetEntry
	query := e.db.QueryContext(models.PetsEquipmentsetEntry{}, c)
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

// getPetsEquipmentsetEntriesBulk godoc
// @Id getPetsEquipmentsetEntriesBulk
// @Summary Gets PetsEquipmentsetEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PetsEquipmentsetEntry
// @Success 200 {array} models.PetsEquipmentsetEntry
// @Failure 500 {string} string "Bad query request"
// @Router /pets_equipmentset_entries/bulk [post]
func (e *PetsEquipmentsetEntryController) getPetsEquipmentsetEntriesBulk(c echo.Context) error {
	var results []models.PetsEquipmentsetEntry

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

	err := e.db.QueryContext(models.PetsEquipmentsetEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
