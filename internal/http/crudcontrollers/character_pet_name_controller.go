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

type CharacterPetNameController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewCharacterPetNameController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *CharacterPetNameController {
	return &CharacterPetNameController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *CharacterPetNameController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_pet_name/:characterId", e.getCharacterPetName, nil),
		routes.RegisterRoute(http.MethodGet, "character_pet_names", e.listCharacterPetNames, nil),
		routes.RegisterRoute(http.MethodGet, "character_pet_names/count", e.getCharacterPetNamesCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_pet_name", e.createCharacterPetName, nil),
		routes.RegisterRoute(http.MethodDelete, "character_pet_name/:characterId", e.deleteCharacterPetName, nil),
		routes.RegisterRoute(http.MethodPatch, "character_pet_name/:characterId", e.updateCharacterPetName, nil),
		routes.RegisterRoute(http.MethodPost, "character_pet_names/bulk", e.getCharacterPetNamesBulk, nil),
	}
}

// listCharacterPetNames godoc
// @Id listCharacterPetNames
// @Summary Lists CharacterPetNames
// @Accept json
// @Produce json
// @Tags CharacterPetName
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetName
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_names [get]
func (e *CharacterPetNameController) listCharacterPetNames(c echo.Context) error {
	var results []models.CharacterPetName
	err := e.db.QueryContext(models.CharacterPetName{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPetName godoc
// @Id getCharacterPetName
// @Summary Gets CharacterPetName
// @Accept json
// @Produce json
// @Tags CharacterPetName
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetName
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_name/{id} [get]
func (e *CharacterPetNameController) getCharacterPetName(c echo.Context) error {
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
	var result models.CharacterPetName
	query := e.db.QueryContext(models.CharacterPetName{}, c)
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

// updateCharacterPetName godoc
// @Id updateCharacterPetName
// @Summary Updates CharacterPetName
// @Accept json
// @Produce json
// @Tags CharacterPetName
// @Param id path int true "Id"
// @Param character_pet_name body models.CharacterPetName true "CharacterPetName"
// @Success 200 {array} models.CharacterPetName
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_pet_name/{id} [patch]
func (e *CharacterPetNameController) updateCharacterPetName(c echo.Context) error {
	request := new(models.CharacterPetName)
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
	var result models.CharacterPetName
	query := e.db.QueryContext(models.CharacterPetName{}, c)
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
		event := fmt.Sprintf("Updated [CharacterPetName] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterPetName godoc
// @Id createCharacterPetName
// @Summary Creates CharacterPetName
// @Accept json
// @Produce json
// @Param character_pet_name body models.CharacterPetName true "CharacterPetName"
// @Tags CharacterPetName
// @Success 200 {array} models.CharacterPetName
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_pet_name [put]
func (e *CharacterPetNameController) createCharacterPetName(c echo.Context) error {
	characterPetName := new(models.CharacterPetName)
	if err := c.Bind(characterPetName); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterPetName{}, c).Model(&models.CharacterPetName{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterPetName).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterPetName{}, characterPetName)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterPetName] [%v] data [%v]", characterPetName.CharacterId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterPetName)
}

// deleteCharacterPetName godoc
// @Id deleteCharacterPetName
// @Summary Deletes CharacterPetName
// @Accept json
// @Produce json
// @Tags CharacterPetName
// @Param id path int true "characterId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_pet_name/{id} [delete]
func (e *CharacterPetNameController) deleteCharacterPetName(c echo.Context) error {
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
	var result models.CharacterPetName
	query := e.db.QueryContext(models.CharacterPetName{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterPetName] [%v] keys [%v]", result.CharacterId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterPetNamesBulk godoc
// @Id getCharacterPetNamesBulk
// @Summary Gets CharacterPetNames in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterPetName
// @Success 200 {array} models.CharacterPetName
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_names/bulk [post]
func (e *CharacterPetNameController) getCharacterPetNamesBulk(c echo.Context) error {
	var results []models.CharacterPetName

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

	err := e.db.QueryContext(models.CharacterPetName{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPetNamesCount godoc
// @Id getCharacterPetNamesCount
// @Summary Counts CharacterPetNames
// @Accept json
// @Produce json
// @Tags CharacterPetName
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPetName
// @Failure 500 {string} string "Bad query request"
// @Router /character_pet_names/count [get]
func (e *CharacterPetNameController) getCharacterPetNamesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterPetName{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}