package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CharacterExpeditionLockoutController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterExpeditionLockoutController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterExpeditionLockoutController {
	return &CharacterExpeditionLockoutController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterExpeditionLockoutController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_expedition_lockout/:id", e.getCharacterExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodGet, "character_expedition_lockouts", e.listCharacterExpeditionLockouts, nil),
		routes.RegisterRoute(http.MethodPut, "character_expedition_lockout", e.createCharacterExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodDelete, "character_expedition_lockout/:id", e.deleteCharacterExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodPatch, "character_expedition_lockout/:id", e.updateCharacterExpeditionLockout, nil),
		routes.RegisterRoute(http.MethodPost, "character_expedition_lockouts/bulk", e.getCharacterExpeditionLockoutsBulk, nil),
	}
}

// listCharacterExpeditionLockouts godoc
// @Id listCharacterExpeditionLockouts
// @Summary Lists CharacterExpeditionLockouts
// @Accept json
// @Produce json
// @Tags CharacterExpeditionLockout
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterExpeditionLockout
// @Failure 500 {string} string "Bad query request"
// @Router /character_expedition_lockouts [get]
func (e *CharacterExpeditionLockoutController) listCharacterExpeditionLockouts(c echo.Context) error {
	var results []models.CharacterExpeditionLockout
	err := e.db.QueryContext(models.CharacterExpeditionLockout{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterExpeditionLockout godoc
// @Id getCharacterExpeditionLockout
// @Summary Gets CharacterExpeditionLockout
// @Accept json
// @Produce json
// @Tags CharacterExpeditionLockout
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterExpeditionLockout
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_expedition_lockout/{id} [get]
func (e *CharacterExpeditionLockoutController) getCharacterExpeditionLockout(c echo.Context) error {
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
	var result models.CharacterExpeditionLockout
	query := e.db.QueryContext(models.CharacterExpeditionLockout{}, c)
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

// updateCharacterExpeditionLockout godoc
// @Id updateCharacterExpeditionLockout
// @Summary Updates CharacterExpeditionLockout
// @Accept json
// @Produce json
// @Tags CharacterExpeditionLockout
// @Param id path int true "Id"
// @Param character_expedition_lockout body models.CharacterExpeditionLockout true "CharacterExpeditionLockout"
// @Success 200 {array} models.CharacterExpeditionLockout
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_expedition_lockout/{id} [patch]
func (e *CharacterExpeditionLockoutController) updateCharacterExpeditionLockout(c echo.Context) error {
	request := new(models.CharacterExpeditionLockout)
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
	var result models.CharacterExpeditionLockout
	query := e.db.QueryContext(models.CharacterExpeditionLockout{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterExpeditionLockout{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterExpeditionLockout godoc
// @Id createCharacterExpeditionLockout
// @Summary Creates CharacterExpeditionLockout
// @Accept json
// @Produce json
// @Param character_expedition_lockout body models.CharacterExpeditionLockout true "CharacterExpeditionLockout"
// @Tags CharacterExpeditionLockout
// @Success 200 {array} models.CharacterExpeditionLockout
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_expedition_lockout [put]
func (e *CharacterExpeditionLockoutController) createCharacterExpeditionLockout(c echo.Context) error {
	characterExpeditionLockout := new(models.CharacterExpeditionLockout)
	if err := c.Bind(characterExpeditionLockout); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterExpeditionLockout{}, c).Model(&models.CharacterExpeditionLockout{}).Create(&characterExpeditionLockout).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterExpeditionLockout)
}

// deleteCharacterExpeditionLockout godoc
// @Id deleteCharacterExpeditionLockout
// @Summary Deletes CharacterExpeditionLockout
// @Accept json
// @Produce json
// @Tags CharacterExpeditionLockout
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_expedition_lockout/{id} [delete]
func (e *CharacterExpeditionLockoutController) deleteCharacterExpeditionLockout(c echo.Context) error {
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
	var result models.CharacterExpeditionLockout
	query := e.db.QueryContext(models.CharacterExpeditionLockout{}, c)
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

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterExpeditionLockoutsBulk godoc
// @Id getCharacterExpeditionLockoutsBulk
// @Summary Gets CharacterExpeditionLockouts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterExpeditionLockout
// @Success 200 {array} models.CharacterExpeditionLockout
// @Failure 500 {string} string "Bad query request"
// @Router /character_expedition_lockouts/bulk [post]
func (e *CharacterExpeditionLockoutController) getCharacterExpeditionLockoutsBulk(c echo.Context) error {
	var results []models.CharacterExpeditionLockout

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

	err := e.db.QueryContext(models.CharacterExpeditionLockout{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
