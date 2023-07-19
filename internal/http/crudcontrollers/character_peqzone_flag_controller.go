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

type CharacterPeqzoneFlagController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterPeqzoneFlagController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterPeqzoneFlagController {
	return &CharacterPeqzoneFlagController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterPeqzoneFlagController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_peqzone_flag/:id", e.getCharacterPeqzoneFlag, nil),
		routes.RegisterRoute(http.MethodGet, "character_peqzone_flags", e.listCharacterPeqzoneFlags, nil),
		routes.RegisterRoute(http.MethodGet, "character_peqzone_flags/count", e.getCharacterPeqzoneFlagsCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_peqzone_flag", e.createCharacterPeqzoneFlag, nil),
		routes.RegisterRoute(http.MethodDelete, "character_peqzone_flag/:id", e.deleteCharacterPeqzoneFlag, nil),
		routes.RegisterRoute(http.MethodPatch, "character_peqzone_flag/:id", e.updateCharacterPeqzoneFlag, nil),
		routes.RegisterRoute(http.MethodPost, "character_peqzone_flags/bulk", e.getCharacterPeqzoneFlagsBulk, nil),
	}
}

// listCharacterPeqzoneFlags godoc
// @Id listCharacterPeqzoneFlags
// @Summary Lists CharacterPeqzoneFlags
// @Accept json
// @Produce json
// @Tags CharacterPeqzoneFlag
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPeqzoneFlag
// @Failure 500 {string} string "Bad query request"
// @Router /character_peqzone_flags [get]
func (e *CharacterPeqzoneFlagController) listCharacterPeqzoneFlags(c echo.Context) error {
	var results []models.CharacterPeqzoneFlag
	err := e.db.QueryContext(models.CharacterPeqzoneFlag{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPeqzoneFlag godoc
// @Id getCharacterPeqzoneFlag
// @Summary Gets CharacterPeqzoneFlag
// @Accept json
// @Produce json
// @Tags CharacterPeqzoneFlag
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPeqzoneFlag
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_peqzone_flag/{id} [get]
func (e *CharacterPeqzoneFlagController) getCharacterPeqzoneFlag(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [zone_id] position [2] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// query builder
	var result models.CharacterPeqzoneFlag
	query := e.db.QueryContext(models.CharacterPeqzoneFlag{}, c)
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

// updateCharacterPeqzoneFlag godoc
// @Id updateCharacterPeqzoneFlag
// @Summary Updates CharacterPeqzoneFlag
// @Accept json
// @Produce json
// @Tags CharacterPeqzoneFlag
// @Param id path int true "Id"
// @Param character_peqzone_flag body models.CharacterPeqzoneFlag true "CharacterPeqzoneFlag"
// @Success 200 {array} models.CharacterPeqzoneFlag
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_peqzone_flag/{id} [patch]
func (e *CharacterPeqzoneFlagController) updateCharacterPeqzoneFlag(c echo.Context) error {
	request := new(models.CharacterPeqzoneFlag)
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

	// key param [zone_id] position [2] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// query builder
	var result models.CharacterPeqzoneFlag
	query := e.db.QueryContext(models.CharacterPeqzoneFlag{}, c)
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
		event := fmt.Sprintf("Updated [CharacterPeqzoneFlag] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterPeqzoneFlag godoc
// @Id createCharacterPeqzoneFlag
// @Summary Creates CharacterPeqzoneFlag
// @Accept json
// @Produce json
// @Param character_peqzone_flag body models.CharacterPeqzoneFlag true "CharacterPeqzoneFlag"
// @Tags CharacterPeqzoneFlag
// @Success 200 {array} models.CharacterPeqzoneFlag
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_peqzone_flag [put]
func (e *CharacterPeqzoneFlagController) createCharacterPeqzoneFlag(c echo.Context) error {
	characterPeqzoneFlag := new(models.CharacterPeqzoneFlag)
	if err := c.Bind(characterPeqzoneFlag); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterPeqzoneFlag{}, c).Model(&models.CharacterPeqzoneFlag{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterPeqzoneFlag).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterPeqzoneFlag{}, characterPeqzoneFlag)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterPeqzoneFlag] [%v] data [%v]", characterPeqzoneFlag.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterPeqzoneFlag)
}

// deleteCharacterPeqzoneFlag godoc
// @Id deleteCharacterPeqzoneFlag
// @Summary Deletes CharacterPeqzoneFlag
// @Accept json
// @Produce json
// @Tags CharacterPeqzoneFlag
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_peqzone_flag/{id} [delete]
func (e *CharacterPeqzoneFlagController) deleteCharacterPeqzoneFlag(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [zone_id] position [2] type [int]
	if len(c.QueryParam("zone_id")) > 0 {
		zoneIdParam, err := strconv.Atoi(c.QueryParam("zone_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zone_id] err [%s]", err.Error())})
		}

		params = append(params, zoneIdParam)
		keys = append(keys, "zone_id = ?")
	}

	// query builder
	var result models.CharacterPeqzoneFlag
	query := e.db.QueryContext(models.CharacterPeqzoneFlag{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterPeqzoneFlag] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterPeqzoneFlagsBulk godoc
// @Id getCharacterPeqzoneFlagsBulk
// @Summary Gets CharacterPeqzoneFlags in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterPeqzoneFlag
// @Success 200 {array} models.CharacterPeqzoneFlag
// @Failure 500 {string} string "Bad query request"
// @Router /character_peqzone_flags/bulk [post]
func (e *CharacterPeqzoneFlagController) getCharacterPeqzoneFlagsBulk(c echo.Context) error {
	var results []models.CharacterPeqzoneFlag

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

	err := e.db.QueryContext(models.CharacterPeqzoneFlag{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPeqzoneFlagsCount godoc
// @Id getCharacterPeqzoneFlagsCount
// @Summary Counts CharacterPeqzoneFlags
// @Accept json
// @Produce json
// @Tags CharacterPeqzoneFlag
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPeqzoneFlag
// @Failure 500 {string} string "Bad query request"
// @Router /character_peqzone_flags/count [get]
func (e *CharacterPeqzoneFlagController) getCharacterPeqzoneFlagsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterPeqzoneFlag{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}