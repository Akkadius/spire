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

type PetitionController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewPetitionController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *PetitionController {
	return &PetitionController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *PetitionController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "petition/:dib", e.getPetition, nil),
		routes.RegisterRoute(http.MethodGet, "petitions", e.listPetitions, nil),
		routes.RegisterRoute(http.MethodGet, "petitions/count", e.getPetitionsCount, nil),
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
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
		event := fmt.Sprintf("Updated [Petition] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
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

	db := e.db.Get(models.Petition{}, c).Model(&models.Petition{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&petition).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.Petition{}, petition)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [Petition] [%v] data [%v]", petition.Dib, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [Petition] [%v] keys [%v]", result.Dib, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
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

// getPetitionsCount godoc
// @Id getPetitionsCount
// @Summary Counts Petitions
// @Accept json
// @Produce json
// @Tags Petition
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
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
// @Router /petitions/count [get]
func (e *PetitionController) getPetitionsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.Petition{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}