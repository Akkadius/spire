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

type CharacterBandolierController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterBandolierController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterBandolierController {
	return &CharacterBandolierController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterBandolierController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_bandolier/:id", e.getCharacterBandolier, nil),
		routes.RegisterRoute(http.MethodGet, "character_bandoliers", e.listCharacterBandoliers, nil),
		routes.RegisterRoute(http.MethodPut, "character_bandolier", e.createCharacterBandolier, nil),
		routes.RegisterRoute(http.MethodDelete, "character_bandolier/:id", e.deleteCharacterBandolier, nil),
		routes.RegisterRoute(http.MethodPatch, "character_bandolier/:id", e.updateCharacterBandolier, nil),
		routes.RegisterRoute(http.MethodPost, "character_bandoliers/bulk", e.getCharacterBandoliersBulk, nil),
	}
}

// listCharacterBandoliers godoc
// @Id listCharacterBandoliers
// @Summary Lists CharacterBandoliers
// @Accept json
// @Produce json
// @Tags CharacterBandolier
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterBandolier
// @Failure 500 {string} string "Bad query request"
// @Router /character_bandoliers [get]
func (e *CharacterBandolierController) listCharacterBandoliers(c echo.Context) error {
	var results []models.CharacterBandolier
	err := e.db.QueryContext(models.CharacterBandolier{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterBandolier godoc
// @Id getCharacterBandolier
// @Summary Gets CharacterBandolier
// @Accept json
// @Produce json
// @Tags CharacterBandolier
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterBandolier
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_bandolier/{id} [get]
func (e *CharacterBandolierController) getCharacterBandolier(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [bandolier_id] position [2] type [tinyint]
	if len(c.QueryParam("bandolier_id")) > 0 {
		bandolierIdParam, err := strconv.Atoi(c.QueryParam("bandolier_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bandolier_id] err [%s]", err.Error())})
		}

		params = append(params, bandolierIdParam)
		keys = append(keys, "bandolier_id = ?")
	}

	// key param [bandolier_slot] position [3] type [tinyint]
	if len(c.QueryParam("bandolier_slot")) > 0 {
		bandolierSlotParam, err := strconv.Atoi(c.QueryParam("bandolier_slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bandolier_slot] err [%s]", err.Error())})
		}

		params = append(params, bandolierSlotParam)
		keys = append(keys, "bandolier_slot = ?")
	}

	// query builder
	var result models.CharacterBandolier
	query := e.db.QueryContext(models.CharacterBandolier{}, c)
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

// updateCharacterBandolier godoc
// @Id updateCharacterBandolier
// @Summary Updates CharacterBandolier
// @Accept json
// @Produce json
// @Tags CharacterBandolier
// @Param id path int true "Id"
// @Param character_bandolier body models.CharacterBandolier true "CharacterBandolier"
// @Success 200 {array} models.CharacterBandolier
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_bandolier/{id} [patch]
func (e *CharacterBandolierController) updateCharacterBandolier(c echo.Context) error {
	request := new(models.CharacterBandolier)
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

	// key param [bandolier_id] position [2] type [tinyint]
	if len(c.QueryParam("bandolier_id")) > 0 {
		bandolierIdParam, err := strconv.Atoi(c.QueryParam("bandolier_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bandolier_id] err [%s]", err.Error())})
		}

		params = append(params, bandolierIdParam)
		keys = append(keys, "bandolier_id = ?")
	}

	// key param [bandolier_slot] position [3] type [tinyint]
	if len(c.QueryParam("bandolier_slot")) > 0 {
		bandolierSlotParam, err := strconv.Atoi(c.QueryParam("bandolier_slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bandolier_slot] err [%s]", err.Error())})
		}

		params = append(params, bandolierSlotParam)
		keys = append(keys, "bandolier_slot = ?")
	}

	// query builder
	var result models.CharacterBandolier
	query := e.db.QueryContext(models.CharacterBandolier{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterBandolier{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterBandolier godoc
// @Id createCharacterBandolier
// @Summary Creates CharacterBandolier
// @Accept json
// @Produce json
// @Param character_bandolier body models.CharacterBandolier true "CharacterBandolier"
// @Tags CharacterBandolier
// @Success 200 {array} models.CharacterBandolier
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_bandolier [put]
func (e *CharacterBandolierController) createCharacterBandolier(c echo.Context) error {
	characterBandolier := new(models.CharacterBandolier)
	if err := c.Bind(characterBandolier); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterBandolier{}, c).Model(&models.CharacterBandolier{}).Create(&characterBandolier).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterBandolier)
}

// deleteCharacterBandolier godoc
// @Id deleteCharacterBandolier
// @Summary Deletes CharacterBandolier
// @Accept json
// @Produce json
// @Tags CharacterBandolier
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_bandolier/{id} [delete]
func (e *CharacterBandolierController) deleteCharacterBandolier(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [bandolier_id] position [2] type [tinyint]
	if len(c.QueryParam("bandolier_id")) > 0 {
		bandolierIdParam, err := strconv.Atoi(c.QueryParam("bandolier_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bandolier_id] err [%s]", err.Error())})
		}

		params = append(params, bandolierIdParam)
		keys = append(keys, "bandolier_id = ?")
	}

	// key param [bandolier_slot] position [3] type [tinyint]
	if len(c.QueryParam("bandolier_slot")) > 0 {
		bandolierSlotParam, err := strconv.Atoi(c.QueryParam("bandolier_slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [bandolier_slot] err [%s]", err.Error())})
		}

		params = append(params, bandolierSlotParam)
		keys = append(keys, "bandolier_slot = ?")
	}

	// query builder
	var result models.CharacterBandolier
	query := e.db.QueryContext(models.CharacterBandolier{}, c)
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

// getCharacterBandoliersBulk godoc
// @Id getCharacterBandoliersBulk
// @Summary Gets CharacterBandoliers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterBandolier
// @Success 200 {array} models.CharacterBandolier
// @Failure 500 {string} string "Bad query request"
// @Router /character_bandoliers/bulk [post]
func (e *CharacterBandolierController) getCharacterBandoliersBulk(c echo.Context) error {
	var results []models.CharacterBandolier

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

	err := e.db.QueryContext(models.CharacterBandolier{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
