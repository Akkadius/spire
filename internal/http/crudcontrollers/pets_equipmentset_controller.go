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

type PetsEquipmentsetController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewPetsEquipmentsetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *PetsEquipmentsetController {
	return &PetsEquipmentsetController{
		db:	 db,
		logger: logger,
	}
}

func (e *PetsEquipmentsetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "pets_equipmentset/:setId", e.getPetsEquipmentset, nil),
		routes.RegisterRoute(http.MethodGet, "pets_equipmentsets", e.listPetsEquipmentsets, nil),
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>PetsEquipmentsetEntries"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>PetsEquipmentsetEntries"
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

	err = e.db.QueryContext(models.PetsEquipmentset{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
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
