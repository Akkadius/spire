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

type CharacterInstanceSafereturnController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterInstanceSafereturnController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterInstanceSafereturnController {
	return &CharacterInstanceSafereturnController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterInstanceSafereturnController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_instance_safereturn/:id", e.getCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodGet, "character_instance_safereturns", e.listCharacterInstanceSafereturns, nil),
		routes.RegisterRoute(http.MethodGet, "character_instance_safereturns/count", e.getCharacterInstanceSafereturnsCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_instance_safereturn", e.createCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodDelete, "character_instance_safereturn/:id", e.deleteCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodPatch, "character_instance_safereturn/:id", e.updateCharacterInstanceSafereturn, nil),
		routes.RegisterRoute(http.MethodPost, "character_instance_safereturns/bulk", e.getCharacterInstanceSafereturnsBulk, nil),
	}
}

// listCharacterInstanceSafereturns godoc
// @Id listCharacterInstanceSafereturns
// @Summary Lists CharacterInstanceSafereturns
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 500 {string} string "Bad query request"
// @Router /character_instance_safereturns [get]
func (e *CharacterInstanceSafereturnController) listCharacterInstanceSafereturns(c echo.Context) error {
	var results []models.CharacterInstanceSafereturn
	err := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterInstanceSafereturn godoc
// @Id getCharacterInstanceSafereturn
// @Summary Gets CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_instance_safereturn/{id} [get]
func (e *CharacterInstanceSafereturnController) getCharacterInstanceSafereturn(c echo.Context) error {
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
	var result models.CharacterInstanceSafereturn
	query := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c)
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

// updateCharacterInstanceSafereturn godoc
// @Id updateCharacterInstanceSafereturn
// @Summary Updates CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param id path int true "Id"
// @Param character_instance_safereturn body models.CharacterInstanceSafereturn true "CharacterInstanceSafereturn"
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_instance_safereturn/{id} [patch]
func (e *CharacterInstanceSafereturnController) updateCharacterInstanceSafereturn(c echo.Context) error {
	request := new(models.CharacterInstanceSafereturn)
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
	var result models.CharacterInstanceSafereturn
	query := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c)
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
		event := fmt.Sprintf("Updated [CharacterInstanceSafereturn] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterInstanceSafereturn godoc
// @Id createCharacterInstanceSafereturn
// @Summary Creates CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Param character_instance_safereturn body models.CharacterInstanceSafereturn true "CharacterInstanceSafereturn"
// @Tags CharacterInstanceSafereturn
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_instance_safereturn [put]
func (e *CharacterInstanceSafereturnController) createCharacterInstanceSafereturn(c echo.Context) error {
	characterInstanceSafereturn := new(models.CharacterInstanceSafereturn)
	if err := c.Bind(characterInstanceSafereturn); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterInstanceSafereturn{}, c).Model(&models.CharacterInstanceSafereturn{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterInstanceSafereturn).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterInstanceSafereturn{}, characterInstanceSafereturn)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterInstanceSafereturn] [%v] data [%v]", characterInstanceSafereturn.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterInstanceSafereturn)
}

// deleteCharacterInstanceSafereturn godoc
// @Id deleteCharacterInstanceSafereturn
// @Summary Deletes CharacterInstanceSafereturn
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_instance_safereturn/{id} [delete]
func (e *CharacterInstanceSafereturnController) deleteCharacterInstanceSafereturn(c echo.Context) error {
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
	var result models.CharacterInstanceSafereturn
	query := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterInstanceSafereturn] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterInstanceSafereturnsBulk godoc
// @Id getCharacterInstanceSafereturnsBulk
// @Summary Gets CharacterInstanceSafereturns in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterInstanceSafereturn
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 500 {string} string "Bad query request"
// @Router /character_instance_safereturns/bulk [post]
func (e *CharacterInstanceSafereturnController) getCharacterInstanceSafereturnsBulk(c echo.Context) error {
	var results []models.CharacterInstanceSafereturn

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

	err := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterInstanceSafereturnsCount godoc
// @Id getCharacterInstanceSafereturnsCount
// @Summary Counts CharacterInstanceSafereturns
// @Accept json
// @Produce json
// @Tags CharacterInstanceSafereturn
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterInstanceSafereturn
// @Failure 500 {string} string "Bad query request"
// @Router /character_instance_safereturns/count [get]
func (e *CharacterInstanceSafereturnController) getCharacterInstanceSafereturnsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterInstanceSafereturn{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}