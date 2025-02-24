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

type CharacterEvolvingItemController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewCharacterEvolvingItemController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *CharacterEvolvingItemController {
	return &CharacterEvolvingItemController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *CharacterEvolvingItemController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_evolving_item/:id", e.getCharacterEvolvingItem, nil),
		routes.RegisterRoute(http.MethodGet, "character_evolving_items", e.listCharacterEvolvingItems, nil),
		routes.RegisterRoute(http.MethodGet, "character_evolving_items/count", e.getCharacterEvolvingItemsCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_evolving_item", e.createCharacterEvolvingItem, nil),
		routes.RegisterRoute(http.MethodDelete, "character_evolving_item/:id", e.deleteCharacterEvolvingItem, nil),
		routes.RegisterRoute(http.MethodPatch, "character_evolving_item/:id", e.updateCharacterEvolvingItem, nil),
		routes.RegisterRoute(http.MethodPost, "character_evolving_items/bulk", e.getCharacterEvolvingItemsBulk, nil),
	}
}

// listCharacterEvolvingItems godoc
// @Id listCharacterEvolvingItems
// @Summary Lists CharacterEvolvingItems
// @Accept json
// @Produce json
// @Tags CharacterEvolvingItem
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterEvolvingItem
// @Failure 500 {string} string "Bad query request"
// @Router /character_evolving_items [get]
func (e *CharacterEvolvingItemController) listCharacterEvolvingItems(c echo.Context) error {
	var results []models.CharacterEvolvingItem
	err := e.db.QueryContext(models.CharacterEvolvingItem{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterEvolvingItem godoc
// @Id getCharacterEvolvingItem
// @Summary Gets CharacterEvolvingItem
// @Accept json
// @Produce json
// @Tags CharacterEvolvingItem
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterEvolvingItem
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_evolving_item/{id} [get]
func (e *CharacterEvolvingItemController) getCharacterEvolvingItem(c echo.Context) error {
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
	var result models.CharacterEvolvingItem
	query := e.db.QueryContext(models.CharacterEvolvingItem{}, c)
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

// updateCharacterEvolvingItem godoc
// @Id updateCharacterEvolvingItem
// @Summary Updates CharacterEvolvingItem
// @Accept json
// @Produce json
// @Tags CharacterEvolvingItem
// @Param id path int true "Id"
// @Param character_evolving_item body models.CharacterEvolvingItem true "CharacterEvolvingItem"
// @Success 200 {array} models.CharacterEvolvingItem
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_evolving_item/{id} [patch]
func (e *CharacterEvolvingItemController) updateCharacterEvolvingItem(c echo.Context) error {
	request := new(models.CharacterEvolvingItem)
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
	var result models.CharacterEvolvingItem
	query := e.db.QueryContext(models.CharacterEvolvingItem{}, c)
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
		event := fmt.Sprintf("Updated [CharacterEvolvingItem] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterEvolvingItem godoc
// @Id createCharacterEvolvingItem
// @Summary Creates CharacterEvolvingItem
// @Accept json
// @Produce json
// @Param character_evolving_item body models.CharacterEvolvingItem true "CharacterEvolvingItem"
// @Tags CharacterEvolvingItem
// @Success 200 {array} models.CharacterEvolvingItem
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_evolving_item [put]
func (e *CharacterEvolvingItemController) createCharacterEvolvingItem(c echo.Context) error {
	characterEvolvingItem := new(models.CharacterEvolvingItem)
	if err := c.Bind(characterEvolvingItem); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterEvolvingItem{}, c).Model(&models.CharacterEvolvingItem{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterEvolvingItem).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterEvolvingItem{}, characterEvolvingItem)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterEvolvingItem] [%v] data [%v]", characterEvolvingItem.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterEvolvingItem)
}

// deleteCharacterEvolvingItem godoc
// @Id deleteCharacterEvolvingItem
// @Summary Deletes CharacterEvolvingItem
// @Accept json
// @Produce json
// @Tags CharacterEvolvingItem
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_evolving_item/{id} [delete]
func (e *CharacterEvolvingItemController) deleteCharacterEvolvingItem(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.CharacterEvolvingItem
	query := e.db.QueryContext(models.CharacterEvolvingItem{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterEvolvingItem] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterEvolvingItemsBulk godoc
// @Id getCharacterEvolvingItemsBulk
// @Summary Gets CharacterEvolvingItems in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterEvolvingItem
// @Success 200 {array} models.CharacterEvolvingItem
// @Failure 500 {string} string "Bad query request"
// @Router /character_evolving_items/bulk [post]
func (e *CharacterEvolvingItemController) getCharacterEvolvingItemsBulk(c echo.Context) error {
	var results []models.CharacterEvolvingItem

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

	err := e.db.QueryContext(models.CharacterEvolvingItem{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterEvolvingItemsCount godoc
// @Id getCharacterEvolvingItemsCount
// @Summary Counts CharacterEvolvingItems
// @Accept json
// @Produce json
// @Tags CharacterEvolvingItem
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterEvolvingItem
// @Failure 500 {string} string "Bad query request"
// @Router /character_evolving_items/count [get]
func (e *CharacterEvolvingItemController) getCharacterEvolvingItemsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterEvolvingItem{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}