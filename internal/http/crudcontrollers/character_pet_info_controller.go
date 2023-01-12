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

type CharacterPetInfoController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterPetInfoController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterPetInfoController {
	return &CharacterPetInfoController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterPetInfoController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_pet_info/:charId", e.getCharacterPetInfo, nil),
		routes.RegisterRoute(http.MethodGet, "character_pet_infos", e.listCharacterPetInfos, nil),
		routes.RegisterRoute(http.MethodGet, "character_pet_infos/count", e.getCharacterPetInfosCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_pet_info", e.createCharacterPetInfo, nil),
		routes.RegisterRoute(http.MethodDelete, "character_pet_info/:charId", e.deleteCharacterPetInfo, nil),
		routes.RegisterRoute(http.MethodPatch, "character_pet_info/:charId", e.updateCharacterPetInfo, nil),
		routes.RegisterRoute(http.MethodPost, "character_pet_infos/bulk", e.getCharacterPetInfosBulk, nil),
	}
}

// listCharacterPetInfos godoc
// @Id listCharacterPetInfos
// @Summary Lists CharacterPetInfos
// @Accept json
// @Produce json
// @Tags CharacterPetInfo
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetInfo
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_infos [get]
func (e *CharacterPetInfoController) listCharacterPetInfos(c echo.Context) error {
	var results []models.CharacterPetInfo
	err := e.db.QueryContext(models.CharacterPetInfo{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPetInfo godoc
// @Id getCharacterPetInfo
// @Summary Gets CharacterPetInfo
// @Accept json
// @Produce json
// @Tags CharacterPetInfo
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetInfo
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_info/{id} [get]
func (e *CharacterPetInfoController) getCharacterPetInfo(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [pet] position [2] type [int]
	if len(c.QueryParam("pet")) > 0 {
		petParam, err := strconv.Atoi(c.QueryParam("pet"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [pet] err [%s]", err.Error())})
		}

		params = append(params, petParam)
		keys = append(keys, "pet = ?")
	}

	// query builder
	var result models.CharacterPetInfo
	query := e.db.QueryContext(models.CharacterPetInfo{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterPetInfo godoc
// @Id updateCharacterPetInfo
// @Summary Updates CharacterPetInfo
// @Accept json
// @Produce json
// @Tags CharacterPetInfo
// @Param id path int true "Id"
// @Param character_pet_info body models.CharacterPetInfo true "CharacterPetInfo"
// @Success 200 {array} models.CharacterPetInfo
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_pet_info/{id} [patch]
func (e *CharacterPetInfoController) updateCharacterPetInfo(c echo.Context) error {
	request := new(models.CharacterPetInfo)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [pet] position [2] type [int]
	if len(c.QueryParam("pet")) > 0 {
		petParam, err := strconv.Atoi(c.QueryParam("pet"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [pet] err [%s]", err.Error())})
		}

		params = append(params, petParam)
		keys = append(keys, "pet = ?")
	}

	// query builder
	var result models.CharacterPetInfo
	query := e.db.QueryContext(models.CharacterPetInfo{}, c)
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
		event := fmt.Sprintf("Updated [CharacterPetInfo] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterPetInfo godoc
// @Id createCharacterPetInfo
// @Summary Creates CharacterPetInfo
// @Accept json
// @Produce json
// @Param character_pet_info body models.CharacterPetInfo true "CharacterPetInfo"
// @Tags CharacterPetInfo
// @Success 200 {array} models.CharacterPetInfo
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_pet_info [put]
func (e *CharacterPetInfoController) createCharacterPetInfo(c echo.Context) error {
	characterPetInfo := new(models.CharacterPetInfo)
	if err := c.Bind(characterPetInfo); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterPetInfo{}, c).Model(&models.CharacterPetInfo{}).Create(&characterPetInfo).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterPetInfo{}, characterPetInfo)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterPetInfo] [%v] data [%v]", characterPetInfo.CharId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterPetInfo)
}

// deleteCharacterPetInfo godoc
// @Id deleteCharacterPetInfo
// @Summary Deletes CharacterPetInfo
// @Accept json
// @Produce json
// @Tags CharacterPetInfo
// @Param id path int true "charId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_pet_info/{id} [delete]
func (e *CharacterPetInfoController) deleteCharacterPetInfo(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [pet] position [2] type [int]
	if len(c.QueryParam("pet")) > 0 {
		petParam, err := strconv.Atoi(c.QueryParam("pet"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [pet] err [%s]", err.Error())})
		}

		params = append(params, petParam)
		keys = append(keys, "pet = ?")
	}

	// query builder
	var result models.CharacterPetInfo
	query := e.db.QueryContext(models.CharacterPetInfo{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterPetInfo] [%v] keys [%v]", result.CharId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterPetInfosBulk godoc
// @Id getCharacterPetInfosBulk
// @Summary Gets CharacterPetInfos in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterPetInfo
// @Success 200 {array} models.CharacterPetInfo
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_infos/bulk [post]
func (e *CharacterPetInfoController) getCharacterPetInfosBulk(c echo.Context) error {
	var results []models.CharacterPetInfo

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

	err := e.db.QueryContext(models.CharacterPetInfo{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPetInfosCount godoc
// @Id getCharacterPetInfosCount
// @Summary Counts CharacterPetInfos
// @Accept json
// @Produce json
// @Tags CharacterPetInfo
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetInfo
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_infos/count [get]
func (e *CharacterPetInfoController) getCharacterPetInfosCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterPetInfo{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}