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

type CharacterParcelsContainerController struct {
	db       *database.Resolver
	auditLog *auditlog.UserEvent
}

func NewCharacterParcelsContainerController(
	db *database.Resolver,
	auditLog *auditlog.UserEvent,
) *CharacterParcelsContainerController {
	return &CharacterParcelsContainerController{
		db:       db,
		auditLog: auditLog,
	}
}

func (e *CharacterParcelsContainerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_parcels_container/:id", e.getCharacterParcelsContainer, nil),
		routes.RegisterRoute(http.MethodGet, "character_parcels_containers", e.listCharacterParcelsContainers, nil),
		routes.RegisterRoute(http.MethodGet, "character_parcels_containers/count", e.getCharacterParcelsContainersCount, nil),
		routes.RegisterRoute(http.MethodPut, "character_parcels_container", e.createCharacterParcelsContainer, nil),
		routes.RegisterRoute(http.MethodDelete, "character_parcels_container/:id", e.deleteCharacterParcelsContainer, nil),
		routes.RegisterRoute(http.MethodPatch, "character_parcels_container/:id", e.updateCharacterParcelsContainer, nil),
		routes.RegisterRoute(http.MethodPost, "character_parcels_containers/bulk", e.getCharacterParcelsContainersBulk, nil),
	}
}

// listCharacterParcelsContainers godoc
// @Id listCharacterParcelsContainers
// @Summary Lists CharacterParcelsContainers
// @Accept json
// @Produce json
// @Tags CharacterParcelsContainer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterParcelsContainer
// @Failure 500 {string} string "Bad query request"
// @Router /character_parcels_containers [get]
func (e *CharacterParcelsContainerController) listCharacterParcelsContainers(c echo.Context) error {
	var results []models.CharacterParcelsContainer
	err := e.db.QueryContext(models.CharacterParcelsContainer{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterParcelsContainer godoc
// @Id getCharacterParcelsContainer
// @Summary Gets CharacterParcelsContainer
// @Accept json
// @Produce json
// @Tags CharacterParcelsContainer
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterParcelsContainer
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_parcels_container/{id} [get]
func (e *CharacterParcelsContainerController) getCharacterParcelsContainer(c echo.Context) error {
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
	var result models.CharacterParcelsContainer
	query := e.db.QueryContext(models.CharacterParcelsContainer{}, c)
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

// updateCharacterParcelsContainer godoc
// @Id updateCharacterParcelsContainer
// @Summary Updates CharacterParcelsContainer
// @Accept json
// @Produce json
// @Tags CharacterParcelsContainer
// @Param id path int true "Id"
// @Param character_parcels_container body models.CharacterParcelsContainer true "CharacterParcelsContainer"
// @Success 200 {array} models.CharacterParcelsContainer
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_parcels_container/{id} [patch]
func (e *CharacterParcelsContainerController) updateCharacterParcelsContainer(c echo.Context) error {
	request := new(models.CharacterParcelsContainer)
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
	var result models.CharacterParcelsContainer
	query := e.db.QueryContext(models.CharacterParcelsContainer{}, c)
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
		event := fmt.Sprintf("Updated [CharacterParcelsContainer] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterParcelsContainer godoc
// @Id createCharacterParcelsContainer
// @Summary Creates CharacterParcelsContainer
// @Accept json
// @Produce json
// @Param character_parcels_container body models.CharacterParcelsContainer true "CharacterParcelsContainer"
// @Tags CharacterParcelsContainer
// @Success 200 {array} models.CharacterParcelsContainer
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_parcels_container [put]
func (e *CharacterParcelsContainerController) createCharacterParcelsContainer(c echo.Context) error {
	characterParcelsContainer := new(models.CharacterParcelsContainer)
	if err := c.Bind(characterParcelsContainer); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.CharacterParcelsContainer{}, c).Model(&models.CharacterParcelsContainer{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&characterParcelsContainer).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.CharacterParcelsContainer{}, characterParcelsContainer)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [CharacterParcelsContainer] [%v] data [%v]", characterParcelsContainer.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, characterParcelsContainer)
}

// deleteCharacterParcelsContainer godoc
// @Id deleteCharacterParcelsContainer
// @Summary Deletes CharacterParcelsContainer
// @Accept json
// @Produce json
// @Tags CharacterParcelsContainer
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_parcels_container/{id} [delete]
func (e *CharacterParcelsContainerController) deleteCharacterParcelsContainer(c echo.Context) error {
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
	var result models.CharacterParcelsContainer
	query := e.db.QueryContext(models.CharacterParcelsContainer{}, c)
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
		event := fmt.Sprintf("Deleted [CharacterParcelsContainer] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterParcelsContainersBulk godoc
// @Id getCharacterParcelsContainersBulk
// @Summary Gets CharacterParcelsContainers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterParcelsContainer
// @Success 200 {array} models.CharacterParcelsContainer
// @Failure 500 {string} string "Bad query request"
// @Router /character_parcels_containers/bulk [post]
func (e *CharacterParcelsContainerController) getCharacterParcelsContainersBulk(c echo.Context) error {
	var results []models.CharacterParcelsContainer

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

	err := e.db.QueryContext(models.CharacterParcelsContainer{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterParcelsContainersCount godoc
// @Id getCharacterParcelsContainersCount
// @Summary Counts CharacterParcelsContainers
// @Accept json
// @Produce json
// @Tags CharacterParcelsContainer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterParcelsContainer
// @Failure 500 {string} string "Bad query request"
// @Router /character_parcels_containers/count [get]
func (e *CharacterParcelsContainerController) getCharacterParcelsContainersCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.CharacterParcelsContainer{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}