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

type PetitionController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewPetitionController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *PetitionController {
	return &PetitionController{
		db:	    db,
		logger: logger,
	}
}

func (e *PetitionController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "petition/:dib", e.getPetition, nil),
		routes.RegisterRoute(http.MethodGet, "petitions", e.listPetitions, nil),
		routes.RegisterRoute(http.MethodPut, "petition", e.createPetition, nil),
		routes.RegisterRoute(http.MethodDelete, "petition/:dib", e.deletePetition, nil),
		routes.RegisterRoute(http.MethodPatch, "petition/:dib", e.updatePetition, nil),
		routes.RegisterRoute(http.MethodPost, "petitions/bulk", e.getPetitionsBulk, nil),
	}
}

// listPetitions godoc
// @Id listPetitions
// @Summary Lists Petitions
// @Accept json
// @Produce json
// @Tags Petition
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Petition
// @Failure 500 {string} string "Bad query request"
// @Router /petitions [get]
func (e *PetitionController) listPetitions(c echo.Context) error {
	var results []models.Petition
	err := e.db.QueryContext(models.Petition{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getPetition godoc
// @Id getPetition
// @Summary Gets Petition
// @Accept json
// @Produce json
// @Tags Petition
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Petition
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /petition/{id} [get]
func (e *PetitionController) getPetition(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	dib, err := strconv.Atoi(c.Param("dib"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Dib]"})
	}
	params = append(params, dib)
	keys = append(keys, "dib = ?")

	// query builder
	var result models.Petition
	query := e.db.QueryContext(models.Petition{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Dib == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updatePetition godoc
// @Id updatePetition
// @Summary Updates Petition
// @Accept json
// @Produce json
// @Tags Petition
// @Param id path int true "Id"
// @Param petition body models.Petition true "Petition"
// @Success 200 {array} models.Petition
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /petition/{id} [patch]
func (e *PetitionController) updatePetition(c echo.Context) error {
	request := new(models.Petition)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	dib, err := strconv.Atoi(c.Param("dib"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Dib]"})
	}
	params = append(params, dib)
	keys = append(keys, "dib = ?")

	// query builder
	var result models.Petition
	query := e.db.QueryContext(models.Petition{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createPetition godoc
// @Id createPetition
// @Summary Creates Petition
// @Accept json
// @Produce json
// @Param petition body models.Petition true "Petition"
// @Tags Petition
// @Success 200 {array} models.Petition
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /petition [put]
func (e *PetitionController) createPetition(c echo.Context) error {
	petition := new(models.Petition)
	if err := c.Bind(petition); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Petition{}, c).Model(&models.Petition{}).Create(&petition).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, petition)
}

// deletePetition godoc
// @Id deletePetition
// @Summary Deletes Petition
// @Accept json
// @Produce json
// @Tags Petition
// @Param id path int true "dib"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /petition/{id} [delete]
func (e *PetitionController) deletePetition(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	dib, err := strconv.Atoi(c.Param("dib"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, dib)
	keys = append(keys, "dib = ?")

	// query builder
	var result models.Petition
	query := e.db.QueryContext(models.Petition{}, c)
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

// getPetitionsBulk godoc
// @Id getPetitionsBulk
// @Summary Gets Petitions in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Petition
// @Success 200 {array} models.Petition
// @Failure 500 {string} string "Bad query request"
// @Router /petitions/bulk [post]
func (e *PetitionController) getPetitionsBulk(c echo.Context) error {
	var results []models.Petition

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

	err := e.db.QueryContext(models.Petition{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
