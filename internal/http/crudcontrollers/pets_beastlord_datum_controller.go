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

type PetsBeastlordDatumController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewPetsBeastlordDatumController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *PetsBeastlordDatumController {
	return &PetsBeastlordDatumController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *PetsBeastlordDatumController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "pets_beastlord_datum/:playerRace", e.getPetsBeastlordDatum, nil),
		routes.RegisterRoute(http.MethodGet, "pets_beastlord_data", e.listPetsBeastlordData, nil),
		routes.RegisterRoute(http.MethodGet, "pets_beastlord_data/count", e.getPetsBeastlordDataCount, nil),
		routes.RegisterRoute(http.MethodPut, "pets_beastlord_datum", e.createPetsBeastlordDatum, nil),
		routes.RegisterRoute(http.MethodDelete, "pets_beastlord_datum/:playerRace", e.deletePetsBeastlordDatum, nil),
		routes.RegisterRoute(http.MethodPatch, "pets_beastlord_datum/:playerRace", e.updatePetsBeastlordDatum, nil),
		routes.RegisterRoute(http.MethodPost, "pets_beastlord_data/bulk", e.getPetsBeastlordDataBulk, nil),
	}
}

// listPetsBeastlordData godoc
// @Id listPetsBeastlordData
// @Summary Lists PetsBeastlordData
// @Accept json
// @Produce json
// @Tags PetsBeastlordDatum
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsBeastlordDatum
// @Failure 500 {string} string "Bad query request"
// @Router /pets_beastlord_data [get]
func (e *PetsBeastlordDatumController) listPetsBeastlordData(c echo.Context) error {
	var results []models.PetsBeastlordDatum
	err := e.db.QueryContext(models.PetsBeastlordDatum{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPetsBeastlordDatum godoc
// @Id getPetsBeastlordDatum
// @Summary Gets PetsBeastlordDatum
// @Accept json
// @Produce json
// @Tags PetsBeastlordDatum
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsBeastlordDatum
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /pets_beastlord_datum/{id} [get]
func (e *PetsBeastlordDatumController) getPetsBeastlordDatum(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	playerRace, err := strconv.Atoi(c.Param("playerRace"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PlayerRace]"})
	}
	params = append(params, playerRace)
	keys = append(keys, "player_race = ?")

	// query builder
	var result models.PetsBeastlordDatum
	query := e.db.QueryContext(models.PetsBeastlordDatum{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.PlayerRace == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updatePetsBeastlordDatum godoc
// @Id updatePetsBeastlordDatum
// @Summary Updates PetsBeastlordDatum
// @Accept json
// @Produce json
// @Tags PetsBeastlordDatum
// @Param id path int true "Id"
// @Param pets_beastlord_datum body models.PetsBeastlordDatum true "PetsBeastlordDatum"
// @Success 200 {array} models.PetsBeastlordDatum
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /pets_beastlord_datum/{id} [patch]
func (e *PetsBeastlordDatumController) updatePetsBeastlordDatum(c echo.Context) error {
	request := new(models.PetsBeastlordDatum)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	playerRace, err := strconv.Atoi(c.Param("playerRace"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PlayerRace]"})
	}
	params = append(params, playerRace)
	keys = append(keys, "player_race = ?")

	// query builder
	var result models.PetsBeastlordDatum
	query := e.db.QueryContext(models.PetsBeastlordDatum{}, c)
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
		event := fmt.Sprintf("Updated [PetsBeastlordDatum] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createPetsBeastlordDatum godoc
// @Id createPetsBeastlordDatum
// @Summary Creates PetsBeastlordDatum
// @Accept json
// @Produce json
// @Param pets_beastlord_datum body models.PetsBeastlordDatum true "PetsBeastlordDatum"
// @Tags PetsBeastlordDatum
// @Success 200 {array} models.PetsBeastlordDatum
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /pets_beastlord_datum [put]
func (e *PetsBeastlordDatumController) createPetsBeastlordDatum(c echo.Context) error {
	petsBeastlordDatum := new(models.PetsBeastlordDatum)
	if err := c.Bind(petsBeastlordDatum); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.PetsBeastlordDatum{}, c).Model(&models.PetsBeastlordDatum{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&petsBeastlordDatum).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.PetsBeastlordDatum{}, petsBeastlordDatum)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [PetsBeastlordDatum] [%v] data [%v]", petsBeastlordDatum.PlayerRace, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, petsBeastlordDatum)
}

// deletePetsBeastlordDatum godoc
// @Id deletePetsBeastlordDatum
// @Summary Deletes PetsBeastlordDatum
// @Accept json
// @Produce json
// @Tags PetsBeastlordDatum
// @Param id path int true "playerRace"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /pets_beastlord_datum/{id} [delete]
func (e *PetsBeastlordDatumController) deletePetsBeastlordDatum(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	playerRace, err := strconv.Atoi(c.Param("playerRace"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, playerRace)
	keys = append(keys, "player_race = ?")

	// query builder
	var result models.PetsBeastlordDatum
	query := e.db.QueryContext(models.PetsBeastlordDatum{}, c)
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
		event := fmt.Sprintf("Deleted [PetsBeastlordDatum] [%v] keys [%v]", result.PlayerRace, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getPetsBeastlordDataBulk godoc
// @Id getPetsBeastlordDataBulk
// @Summary Gets PetsBeastlordData in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags PetsBeastlordDatum
// @Success 200 {array} models.PetsBeastlordDatum
// @Failure 500 {string} string "Bad query request"
// @Router /pets_beastlord_data/bulk [post]
func (e *PetsBeastlordDatumController) getPetsBeastlordDataBulk(c echo.Context) error {
	var results []models.PetsBeastlordDatum

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

	err := e.db.QueryContext(models.PetsBeastlordDatum{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getPetsBeastlordDataCount godoc
// @Id getPetsBeastlordDataCount
// @Summary Counts PetsBeastlordData
// @Accept json
// @Produce json
// @Tags PetsBeastlordDatum
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.PetsBeastlordDatum
// @Failure 500 {string} string "Bad query request"
// @Router /pets_beastlord_data/count [get]
func (e *PetsBeastlordDatumController) getPetsBeastlordDataCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.PetsBeastlordDatum{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}