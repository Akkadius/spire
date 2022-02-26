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

type FishingController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewFishingController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *FishingController {
	return &FishingController{
		db:	 db,
		logger: logger,
	}
}

func (e *FishingController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "fishing/:id", e.getFishing, nil),
		routes.RegisterRoute(http.MethodGet, "fishings", e.listFishings, nil),
		routes.RegisterRoute(http.MethodPut, "fishing", e.createFishing, nil),
		routes.RegisterRoute(http.MethodDelete, "fishing/:id", e.deleteFishing, nil),
		routes.RegisterRoute(http.MethodPatch, "fishing/:id", e.updateFishing, nil),
		routes.RegisterRoute(http.MethodPost, "fishings/bulk", e.getFishingsBulk, nil),
	}
}

// listFishings godoc
// @Id listFishings
// @Summary Lists Fishings
// @Accept json
// @Produce json
// @Tags Fishing
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Bad query request"
// @Router /fishings [get]
func (e *FishingController) listFishings(c echo.Context) error {
	var results []models.Fishing
	err := e.db.QueryContext(models.Fishing{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getFishing godoc
// @Id getFishing
// @Summary Gets Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Fishing
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /fishing/{id} [get]
func (e *FishingController) getFishing(c echo.Context) error {
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
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

// updateFishing godoc
// @Id updateFishing
// @Summary Updates Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "Id"
// @Param fishing body models.Fishing true "Fishing"
// @Success 200 {array} models.Fishing
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /fishing/{id} [patch]
func (e *FishingController) updateFishing(c echo.Context) error {
	request := new(models.Fishing)
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createFishing godoc
// @Id createFishing
// @Summary Creates Fishing
// @Accept json
// @Produce json
// @Param fishing body models.Fishing true "Fishing"
// @Tags Fishing
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /fishing [put]
func (e *FishingController) createFishing(c echo.Context) error {
	fishing := new(models.Fishing)
	if err := c.Bind(fishing); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Fishing{}, c).Model(&models.Fishing{}).Create(&fishing).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, fishing)
}

// deleteFishing godoc
// @Id deleteFishing
// @Summary Deletes Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /fishing/{id} [delete]
func (e *FishingController) deleteFishing(c echo.Context) error {
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.Fishing{}, c).Model(&models.Fishing{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getFishingsBulk godoc
// @Id getFishingsBulk
// @Summary Gets Fishings in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Fishing
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Bad query request"
// @Router /fishings/bulk [post]
func (e *FishingController) getFishingsBulk(c echo.Context) error {
	var results []models.Fishing

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

	err := e.db.QueryContext(models.Fishing{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
