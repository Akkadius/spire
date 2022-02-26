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

type CharacterItemRecastController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterItemRecastController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterItemRecastController {
	return &CharacterItemRecastController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterItemRecastController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_item_recast/:id", e.getCharacterItemRecast, nil),
		routes.RegisterRoute(http.MethodGet, "character_item_recasts", e.listCharacterItemRecasts, nil),
		routes.RegisterRoute(http.MethodPut, "character_item_recast", e.createCharacterItemRecast, nil),
		routes.RegisterRoute(http.MethodDelete, "character_item_recast/:id", e.deleteCharacterItemRecast, nil),
		routes.RegisterRoute(http.MethodPatch, "character_item_recast/:id", e.updateCharacterItemRecast, nil),
		routes.RegisterRoute(http.MethodPost, "character_item_recasts/bulk", e.getCharacterItemRecastsBulk, nil),
	}
}

// listCharacterItemRecasts godoc
// @Id listCharacterItemRecasts
// @Summary Lists CharacterItemRecasts
// @Accept json
// @Produce json
// @Tags CharacterItemRecast
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterItemRecast
// @Failure 500 {string} string "Bad query request"
// @Router /character_item_recasts [get]
func (e *CharacterItemRecastController) listCharacterItemRecasts(c echo.Context) error {
	var results []models.CharacterItemRecast
	err := e.db.QueryContext(models.CharacterItemRecast{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterItemRecast godoc
// @Id getCharacterItemRecast
// @Summary Gets CharacterItemRecast
// @Accept json
// @Produce json
// @Tags CharacterItemRecast
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterItemRecast
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_item_recast/{id} [get]
func (e *CharacterItemRecastController) getCharacterItemRecast(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [recast_type] position [2] type [smallint]
	if len(c.QueryParam("recast_type")) > 0 {
		recastTypeParam, err := strconv.Atoi(c.QueryParam("recast_type"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [recast_type] err [%s]", err.Error())})
		}

		params = append(params, recastTypeParam)
		keys = append(keys, "recast_type = ?")
	}

	// query builder
	var result models.CharacterItemRecast
	query := e.db.QueryContext(models.CharacterItemRecast{}, c)
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

// updateCharacterItemRecast godoc
// @Id updateCharacterItemRecast
// @Summary Updates CharacterItemRecast
// @Accept json
// @Produce json
// @Tags CharacterItemRecast
// @Param id path int true "Id"
// @Param character_item_recast body models.CharacterItemRecast true "CharacterItemRecast"
// @Success 200 {array} models.CharacterItemRecast
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_item_recast/{id} [patch]
func (e *CharacterItemRecastController) updateCharacterItemRecast(c echo.Context) error {
	request := new(models.CharacterItemRecast)
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

	// key param [recast_type] position [2] type [smallint]
	if len(c.QueryParam("recast_type")) > 0 {
		recastTypeParam, err := strconv.Atoi(c.QueryParam("recast_type"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [recast_type] err [%s]", err.Error())})
		}

		params = append(params, recastTypeParam)
		keys = append(keys, "recast_type = ?")
	}

	// query builder
	var result models.CharacterItemRecast
	query := e.db.QueryContext(models.CharacterItemRecast{}, c)
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

// createCharacterItemRecast godoc
// @Id createCharacterItemRecast
// @Summary Creates CharacterItemRecast
// @Accept json
// @Produce json
// @Param character_item_recast body models.CharacterItemRecast true "CharacterItemRecast"
// @Tags CharacterItemRecast
// @Success 200 {array} models.CharacterItemRecast
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_item_recast [put]
func (e *CharacterItemRecastController) createCharacterItemRecast(c echo.Context) error {
	characterItemRecast := new(models.CharacterItemRecast)
	if err := c.Bind(characterItemRecast); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterItemRecast{}, c).Model(&models.CharacterItemRecast{}).Create(&characterItemRecast).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterItemRecast)
}

// deleteCharacterItemRecast godoc
// @Id deleteCharacterItemRecast
// @Summary Deletes CharacterItemRecast
// @Accept json
// @Produce json
// @Tags CharacterItemRecast
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_item_recast/{id} [delete]
func (e *CharacterItemRecastController) deleteCharacterItemRecast(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [recast_type] position [2] type [smallint]
	if len(c.QueryParam("recast_type")) > 0 {
		recastTypeParam, err := strconv.Atoi(c.QueryParam("recast_type"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [recast_type] err [%s]", err.Error())})
		}

		params = append(params, recastTypeParam)
		keys = append(keys, "recast_type = ?")
	}

	// query builder
	var result models.CharacterItemRecast
	query := e.db.QueryContext(models.CharacterItemRecast{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.CharacterItemRecast{}, c).Model(&models.CharacterItemRecast{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterItemRecastsBulk godoc
// @Id getCharacterItemRecastsBulk
// @Summary Gets CharacterItemRecasts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterItemRecast
// @Success 200 {array} models.CharacterItemRecast
// @Failure 500 {string} string "Bad query request"
// @Router /character_item_recasts/bulk [post]
func (e *CharacterItemRecastController) getCharacterItemRecastsBulk(c echo.Context) error {
	var results []models.CharacterItemRecast

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

	err := e.db.QueryContext(models.CharacterItemRecast{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
