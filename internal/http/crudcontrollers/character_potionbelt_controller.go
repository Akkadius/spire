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

type CharacterPotionbeltController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewCharacterPotionbeltController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *CharacterPotionbeltController {
	return &CharacterPotionbeltController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *CharacterPotionbeltController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_potionbelt/:id", e.getCharacterPotionbelt, nil),
		routes.RegisterRoute(http.MethodGet, "character_potionbelts", e.listCharacterPotionbelts, nil),
		routes.RegisterRoute(http.MethodPut, "character_potionbelt", e.createCharacterPotionbelt, nil),
		routes.RegisterRoute(http.MethodDelete, "character_potionbelt/:id", e.deleteCharacterPotionbelt, nil),
		routes.RegisterRoute(http.MethodPatch, "character_potionbelt/:id", e.updateCharacterPotionbelt, nil),
		routes.RegisterRoute(http.MethodPost, "character_potionbelts/bulk", e.getCharacterPotionbeltsBulk, nil),
	}
}

// listCharacterPotionbelts godoc
// @Id listCharacterPotionbelts
// @Summary Lists CharacterPotionbelts
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 500 {string} string "Bad query request"
// @Router /character_potionbelts [get]
func (e *CharacterPotionbeltController) listCharacterPotionbelts(c echo.Context) error {
	var results []models.CharacterPotionbelt
	err := e.db.QueryContext(models.CharacterPotionbelt{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPotionbelt godoc
// @Id getCharacterPotionbelt
// @Summary Gets CharacterPotionbelt
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_potionbelt/{id} [get]
func (e *CharacterPotionbeltController) getCharacterPotionbelt(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [potion_id] position [2] type [tinyint]
	if len(c.QueryParam("potion_id")) > 0 {
		potionIdParam, err := strconv.Atoi(c.QueryParam("potion_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [potion_id] err [%s]", err.Error())})
		}

		params = append(params, potionIdParam)
		keys = append(keys, "potion_id = ?")
	}

	// query builder
	var result models.CharacterPotionbelt
	query := e.db.QueryContext(models.CharacterPotionbelt{}, c)
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

// updateCharacterPotionbelt godoc
// @Id updateCharacterPotionbelt
// @Summary Updates CharacterPotionbelt
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param id path int true "Id"
// @Param character_potionbelt body models.CharacterPotionbelt true "CharacterPotionbelt"
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_potionbelt/{id} [patch]
func (e *CharacterPotionbeltController) updateCharacterPotionbelt(c echo.Context) error {
	request := new(models.CharacterPotionbelt)
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

	// key param [potion_id] position [2] type [tinyint]
	if len(c.QueryParam("potion_id")) > 0 {
		potionIdParam, err := strconv.Atoi(c.QueryParam("potion_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [potion_id] err [%s]", err.Error())})
		}

		params = append(params, potionIdParam)
		keys = append(keys, "potion_id = ?")
	}

	// query builder
	var result models.CharacterPotionbelt
	query := e.db.QueryContext(models.CharacterPotionbelt{}, c)
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
		event := fmt.Sprintf("Updated [CharacterPotionbelt] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterPotionbelt godoc
// @Id createCharacterPotionbelt
// @Summary Creates CharacterPotionbelt
// @Accept json
// @Produce json
// @Param character_potionbelt body models.CharacterPotionbelt true "CharacterPotionbelt"
// @Tags CharacterPotionbelt
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_potionbelt [put]
func (e *CharacterPotionbeltController) createCharacterPotionbelt(c echo.Context) error {
	characterPotionbelt := new(models.CharacterPotionbelt)
	if err := c.Bind(characterPotionbelt); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterPotionbelt{}, c).Model(&models.CharacterPotionbelt{}).Create(&characterPotionbelt).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterPotionbelt{}, characterPotionbelt)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterPotionbelt] [%v] data [%v]", characterPotionbelt.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterPotionbelt)
}

// deleteCharacterPotionbelt godoc
// @Id deleteCharacterPotionbelt
// @Summary Deletes CharacterPotionbelt
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_potionbelt/{id} [delete]
func (e *CharacterPotionbeltController) deleteCharacterPotionbelt(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [potion_id] position [2] type [tinyint]
	if len(c.QueryParam("potion_id")) > 0 {
		potionIdParam, err := strconv.Atoi(c.QueryParam("potion_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [potion_id] err [%s]", err.Error())})
		}

		params = append(params, potionIdParam)
		keys = append(keys, "potion_id = ?")
	}

	// query builder
	var result models.CharacterPotionbelt
	query := e.db.QueryContext(models.CharacterPotionbelt{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterPotionbelt] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterPotionbeltsBulk godoc
// @Id getCharacterPotionbeltsBulk
// @Summary Gets CharacterPotionbelts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterPotionbelt
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 500 {string} string "Bad query request"
// @Router /character_potionbelts/bulk [post]
func (e *CharacterPotionbeltController) getCharacterPotionbeltsBulk(c echo.Context) error {
	var results []models.CharacterPotionbelt

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

	err := e.db.QueryContext(models.CharacterPotionbelt{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
