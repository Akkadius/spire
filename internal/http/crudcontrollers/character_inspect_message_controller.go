package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterInspectMessageController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterInspectMessageController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterInspectMessageController {
	return &CharacterInspectMessageController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterInspectMessageController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_inspect_message/:id", e.getCharacterInspectMessage, nil),
		routes.RegisterRoute(http.MethodGet, "character_inspect_messages", e.listCharacterInspectMessages, nil),
		routes.RegisterRoute(http.MethodPut, "character_inspect_message", e.createCharacterInspectMessage, nil),
		routes.RegisterRoute(http.MethodDelete, "character_inspect_message/:id", e.deleteCharacterInspectMessage, nil),
		routes.RegisterRoute(http.MethodPatch, "character_inspect_message/:id", e.updateCharacterInspectMessage, nil),
		routes.RegisterRoute(http.MethodPost, "character_inspect_messages/bulk", e.getCharacterInspectMessagesBulk, nil),
	}
}

// listCharacterInspectMessages godoc
// @Id listCharacterInspectMessages
// @Summary Lists CharacterInspectMessages
// @Accept json
// @Produce json
// @Tags CharacterInspectMessage
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterInspectMessage
// @Failure 500 {string} string "Bad query request"
// @Router /character_inspect_messages [get]
func (e *CharacterInspectMessageController) listCharacterInspectMessages(c echo.Context) error {
	var results []models.CharacterInspectMessage
	err := e.db.QueryContext(models.CharacterInspectMessage{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterInspectMessage godoc
// @Id getCharacterInspectMessage
// @Summary Gets CharacterInspectMessage
// @Accept json
// @Produce json
// @Tags CharacterInspectMessage
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterInspectMessage
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_inspect_message/{id} [get]
func (e *CharacterInspectMessageController) getCharacterInspectMessage(c echo.Context) error {
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
	var result models.CharacterInspectMessage
	query := e.db.QueryContext(models.CharacterInspectMessage{}, c)
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

// updateCharacterInspectMessage godoc
// @Id updateCharacterInspectMessage
// @Summary Updates CharacterInspectMessage
// @Accept json
// @Produce json
// @Tags CharacterInspectMessage
// @Param id path int true "Id"
// @Param character_inspect_message body models.CharacterInspectMessage true "CharacterInspectMessage"
// @Success 200 {array} models.CharacterInspectMessage
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_inspect_message/{id} [patch]
func (e *CharacterInspectMessageController) updateCharacterInspectMessage(c echo.Context) error {
	request := new(models.CharacterInspectMessage)
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
	var result models.CharacterInspectMessage
	query := e.db.QueryContext(models.CharacterInspectMessage{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterInspectMessage godoc
// @Id createCharacterInspectMessage
// @Summary Creates CharacterInspectMessage
// @Accept json
// @Produce json
// @Param character_inspect_message body models.CharacterInspectMessage true "CharacterInspectMessage"
// @Tags CharacterInspectMessage
// @Success 200 {array} models.CharacterInspectMessage
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_inspect_message [put]
func (e *CharacterInspectMessageController) createCharacterInspectMessage(c echo.Context) error {
	characterInspectMessage := new(models.CharacterInspectMessage)
	if err := c.Bind(characterInspectMessage); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterInspectMessage{}, c).Model(&models.CharacterInspectMessage{}).Create(&characterInspectMessage).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterInspectMessage)
}

// deleteCharacterInspectMessage godoc
// @Id deleteCharacterInspectMessage
// @Summary Deletes CharacterInspectMessage
// @Accept json
// @Produce json
// @Tags CharacterInspectMessage
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_inspect_message/{id} [delete]
func (e *CharacterInspectMessageController) deleteCharacterInspectMessage(c echo.Context) error {
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
	var result models.CharacterInspectMessage
	query := e.db.QueryContext(models.CharacterInspectMessage{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.CharacterInspectMessage{}, c).Model(&models.CharacterInspectMessage{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterInspectMessagesBulk godoc
// @Id getCharacterInspectMessagesBulk
// @Summary Gets CharacterInspectMessages in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterInspectMessage
// @Success 200 {array} models.CharacterInspectMessage
// @Failure 500 {string} string "Bad query request"
// @Router /character_inspect_messages/bulk [post]
func (e *CharacterInspectMessageController) getCharacterInspectMessagesBulk(c echo.Context) error {
	var results []models.CharacterInspectMessage

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

	err := e.db.QueryContext(models.CharacterInspectMessage{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
