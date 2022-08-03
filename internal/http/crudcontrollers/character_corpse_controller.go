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

type CharacterCorpseController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterCorpseController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterCorpseController {
	return &CharacterCorpseController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterCorpseController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_corpse/:id", e.getCharacterCorpse, nil),
		routes.RegisterRoute(http.MethodGet, "character_corpses", e.listCharacterCorpses, nil),
		routes.RegisterRoute(http.MethodPut, "character_corpse", e.createCharacterCorpse, nil),
		routes.RegisterRoute(http.MethodDelete, "character_corpse/:id", e.deleteCharacterCorpse, nil),
		routes.RegisterRoute(http.MethodPatch, "character_corpse/:id", e.updateCharacterCorpse, nil),
		routes.RegisterRoute(http.MethodPost, "character_corpses/bulk", e.getCharacterCorpsesBulk, nil),
	}
}

// listCharacterCorpses godoc
// @Id listCharacterCorpses
// @Summary Lists CharacterCorpses
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterCorpse
// @Failure 500 {string} string "Bad query request"
// @Router /character_corpses [get]
func (e *CharacterCorpseController) listCharacterCorpses(c echo.Context) error {
	var results []models.CharacterCorpse
	err := e.db.QueryContext(models.CharacterCorpse{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterCorpse godoc
// @Id getCharacterCorpse
// @Summary Gets CharacterCorpse
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterCorpse
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_corpse/{id} [get]
func (e *CharacterCorpseController) getCharacterCorpse(c echo.Context) error {
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
	var result models.CharacterCorpse
	query := e.db.QueryContext(models.CharacterCorpse{}, c)
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

// updateCharacterCorpse godoc
// @Id updateCharacterCorpse
// @Summary Updates CharacterCorpse
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param id path int true "Id"
// @Param character_corpse body models.CharacterCorpse true "CharacterCorpse"
// @Success 200 {array} models.CharacterCorpse
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_corpse/{id} [patch]
func (e *CharacterCorpseController) updateCharacterCorpse(c echo.Context) error {
	request := new(models.CharacterCorpse)
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
	var result models.CharacterCorpse
	query := e.db.QueryContext(models.CharacterCorpse{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterCorpse{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterCorpse godoc
// @Id createCharacterCorpse
// @Summary Creates CharacterCorpse
// @Accept json
// @Produce json
// @Param character_corpse body models.CharacterCorpse true "CharacterCorpse"
// @Tags CharacterCorpse
// @Success 200 {array} models.CharacterCorpse
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_corpse [put]
func (e *CharacterCorpseController) createCharacterCorpse(c echo.Context) error {
	characterCorpse := new(models.CharacterCorpse)
	if err := c.Bind(characterCorpse); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterCorpse{}, c).Model(&models.CharacterCorpse{}).Create(&characterCorpse).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterCorpse)
}

// deleteCharacterCorpse godoc
// @Id deleteCharacterCorpse
// @Summary Deletes CharacterCorpse
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_corpse/{id} [delete]
func (e *CharacterCorpseController) deleteCharacterCorpse(c echo.Context) error {
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
	var result models.CharacterCorpse
	query := e.db.QueryContext(models.CharacterCorpse{}, c)
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

// getCharacterCorpsesBulk godoc
// @Id getCharacterCorpsesBulk
// @Summary Gets CharacterCorpses in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterCorpse
// @Success 200 {array} models.CharacterCorpse
// @Failure 500 {string} string "Bad query request"
// @Router /character_corpses/bulk [post]
func (e *CharacterCorpseController) getCharacterCorpsesBulk(c echo.Context) error {
	var results []models.CharacterCorpse

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

	err := e.db.QueryContext(models.CharacterCorpse{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
