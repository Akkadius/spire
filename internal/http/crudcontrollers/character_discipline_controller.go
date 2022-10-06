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

type CharacterDisciplineController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterDisciplineController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterDisciplineController {
	return &CharacterDisciplineController{
		db:	    db,
		logger: logger,
	}
}

func (e *CharacterDisciplineController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_discipline/:id", e.getCharacterDiscipline, nil),
		routes.RegisterRoute(http.MethodGet, "character_disciplines", e.listCharacterDisciplines, nil),
		routes.RegisterRoute(http.MethodPut, "character_discipline", e.createCharacterDiscipline, nil),
		routes.RegisterRoute(http.MethodDelete, "character_discipline/:id", e.deleteCharacterDiscipline, nil),
		routes.RegisterRoute(http.MethodPatch, "character_discipline/:id", e.updateCharacterDiscipline, nil),
		routes.RegisterRoute(http.MethodPost, "character_disciplines/bulk", e.getCharacterDisciplinesBulk, nil),
	}
}

// listCharacterDisciplines godoc
// @Id listCharacterDisciplines
// @Summary Lists CharacterDisciplines
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterDiscipline
// @Failure 500 {string} string "Bad query request"
// @Router /character_disciplines [get]
func (e *CharacterDisciplineController) listCharacterDisciplines(c echo.Context) error {
	var results []models.CharacterDiscipline
	err := e.db.QueryContext(models.CharacterDiscipline{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterDiscipline godoc
// @Id getCharacterDiscipline
// @Summary Gets CharacterDiscipline
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterDiscipline
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_discipline/{id} [get]
func (e *CharacterDisciplineController) getCharacterDiscipline(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot_id] position [2] type [smallint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterDiscipline
	query := e.db.QueryContext(models.CharacterDiscipline{}, c)
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

// updateCharacterDiscipline godoc
// @Id updateCharacterDiscipline
// @Summary Updates CharacterDiscipline
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param id path int true "Id"
// @Param character_discipline body models.CharacterDiscipline true "CharacterDiscipline"
// @Success 200 {array} models.CharacterDiscipline
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_discipline/{id} [patch]
func (e *CharacterDisciplineController) updateCharacterDiscipline(c echo.Context) error {
	request := new(models.CharacterDiscipline)
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

	// key param [slot_id] position [2] type [smallint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterDiscipline
	query := e.db.QueryContext(models.CharacterDiscipline{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterDiscipline godoc
// @Id createCharacterDiscipline
// @Summary Creates CharacterDiscipline
// @Accept json
// @Produce json
// @Param character_discipline body models.CharacterDiscipline true "CharacterDiscipline"
// @Tags CharacterDiscipline
// @Success 200 {array} models.CharacterDiscipline
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_discipline [put]
func (e *CharacterDisciplineController) createCharacterDiscipline(c echo.Context) error {
	characterDiscipline := new(models.CharacterDiscipline)
	if err := c.Bind(characterDiscipline); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterDiscipline{}, c).Model(&models.CharacterDiscipline{}).Create(&characterDiscipline).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterDiscipline)
}

// deleteCharacterDiscipline godoc
// @Id deleteCharacterDiscipline
// @Summary Deletes CharacterDiscipline
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_discipline/{id} [delete]
func (e *CharacterDisciplineController) deleteCharacterDiscipline(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot_id] position [2] type [smallint]
	if len(c.QueryParam("slot_id")) > 0 {
		slotIdParam, err := strconv.Atoi(c.QueryParam("slot_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot_id] err [%s]", err.Error())})
		}

		params = append(params, slotIdParam)
		keys = append(keys, "slot_id = ?")
	}

	// query builder
	var result models.CharacterDiscipline
	query := e.db.QueryContext(models.CharacterDiscipline{}, c)
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

// getCharacterDisciplinesBulk godoc
// @Id getCharacterDisciplinesBulk
// @Summary Gets CharacterDisciplines in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterDiscipline
// @Success 200 {array} models.CharacterDiscipline
// @Failure 500 {string} string "Bad query request"
// @Router /character_disciplines/bulk [post]
func (e *CharacterDisciplineController) getCharacterDisciplinesBulk(c echo.Context) error {
	var results []models.CharacterDiscipline

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

	err := e.db.QueryContext(models.CharacterDiscipline{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
