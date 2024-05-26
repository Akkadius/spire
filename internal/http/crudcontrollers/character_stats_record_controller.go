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

type CharacterStatsRecordController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewCharacterStatsRecordController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *CharacterStatsRecordController {
	return &CharacterStatsRecordController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *CharacterStatsRecordController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_stats_record/:characterId", e.getCharacterStatsRecord, nil),
		routes.RegisterRoute(http.MethodGet, "character_stats_records", e.listCharacterStatsRecords, nil),
		routes.RegisterRoute(http.MethodGet, "character_stats_records/count", e.getCharacterStatsRecordsCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_stats_record", e.createCharacterStatsRecord, nil),
		routes.RegisterRoute(http.MethodDelete, "character_stats_record/:characterId", e.deleteCharacterStatsRecord, nil),
		routes.RegisterRoute(http.MethodPatch, "character_stats_record/:characterId", e.updateCharacterStatsRecord, nil),
		routes.RegisterRoute(http.MethodPost, "character_stats_records/bulk", e.getCharacterStatsRecordsBulk, nil),
	}
}

// listCharacterStatsRecords godoc
// @Id listCharacterStatsRecords
// @Summary Lists CharacterStatsRecords
// @Accept json
// @Produce json
// @Tags CharacterStatsRecord
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterStatsRecord
// @Failure 500 {string} string "Bad query request"
// @Router /character_stats_records [get]
func (e *CharacterStatsRecordController) listCharacterStatsRecords(c echo.Context) error {
	var results []models.CharacterStatsRecord
	err := e.db.QueryContext(models.CharacterStatsRecord{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterStatsRecord godoc
// @Id getCharacterStatsRecord
// @Summary Gets CharacterStatsRecord
// @Accept json
// @Produce json
// @Tags CharacterStatsRecord
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterStatsRecord
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_stats_record/{id} [get]
func (e *CharacterStatsRecordController) getCharacterStatsRecord(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharacterId]"})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// query builder
	var result models.CharacterStatsRecord
	query := e.db.QueryContext(models.CharacterStatsRecord{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharacterId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterStatsRecord godoc
// @Id updateCharacterStatsRecord
// @Summary Updates CharacterStatsRecord
// @Accept json
// @Produce json
// @Tags CharacterStatsRecord
// @Param id path int true "Id"
// @Param character_stats_record body models.CharacterStatsRecord true "CharacterStatsRecord"
// @Success 200 {array} models.CharacterStatsRecord
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_stats_record/{id} [patch]
func (e *CharacterStatsRecordController) updateCharacterStatsRecord(c echo.Context) error {
	request := new(models.CharacterStatsRecord)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharacterId]"})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// query builder
	var result models.CharacterStatsRecord
	query := e.db.QueryContext(models.CharacterStatsRecord{}, c)
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
		event := fmt.Sprintf("Updated [CharacterStatsRecord] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterStatsRecord godoc
// @Id createCharacterStatsRecord
// @Summary Creates CharacterStatsRecord
// @Accept json
// @Produce json
// @Param character_stats_record body models.CharacterStatsRecord true "CharacterStatsRecord"
// @Tags CharacterStatsRecord
// @Success 200 {array} models.CharacterStatsRecord
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_stats_record [put]
func (e *CharacterStatsRecordController) createCharacterStatsRecord(c echo.Context) error {
	characterStatsRecord := new(models.CharacterStatsRecord)
	if err := c.Bind(characterStatsRecord); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterStatsRecord{}, c).Model(&models.CharacterStatsRecord{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterStatsRecord).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterStatsRecord{}, characterStatsRecord)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterStatsRecord] [%v] data [%v]", characterStatsRecord.CharacterId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterStatsRecord)
}

// deleteCharacterStatsRecord godoc
// @Id deleteCharacterStatsRecord
// @Summary Deletes CharacterStatsRecord
// @Accept json
// @Produce json
// @Tags CharacterStatsRecord
// @Param id path int true "characterId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_stats_record/{id} [delete]
func (e *CharacterStatsRecordController) deleteCharacterStatsRecord(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	characterId, err := strconv.Atoi(c.Param("characterId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, characterId)
	keys = append(keys, "character_id = ?")

	// query builder
	var result models.CharacterStatsRecord
	query := e.db.QueryContext(models.CharacterStatsRecord{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterStatsRecord] [%v] keys [%v]", result.CharacterId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterStatsRecordsBulk godoc
// @Id getCharacterStatsRecordsBulk
// @Summary Gets CharacterStatsRecords in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterStatsRecord
// @Success 200 {array} models.CharacterStatsRecord
// @Failure 500 {string} string "Bad query request"
// @Router /character_stats_records/bulk [post]
func (e *CharacterStatsRecordController) getCharacterStatsRecordsBulk(c echo.Context) error {
	var results []models.CharacterStatsRecord

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

	err := e.db.QueryContext(models.CharacterStatsRecord{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterStatsRecordsCount godoc
// @Id getCharacterStatsRecordsCount
// @Summary Counts CharacterStatsRecords
// @Accept json
// @Produce json
// @Tags CharacterStatsRecord
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterStatsRecord
// @Failure 500 {string} string "Bad query request"
// @Router /character_stats_records/count [get]
func (e *CharacterStatsRecordController) getCharacterStatsRecordsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterStatsRecord{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}