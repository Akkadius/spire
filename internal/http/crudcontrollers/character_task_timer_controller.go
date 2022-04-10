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

type CharacterTaskTimerController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterTaskTimerController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterTaskTimerController {
	return &CharacterTaskTimerController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterTaskTimerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_task_timer/:id", e.getCharacterTaskTimer, nil),
		routes.RegisterRoute(http.MethodGet, "character_task_timers", e.listCharacterTaskTimers, nil),
		routes.RegisterRoute(http.MethodPut, "character_task_timer", e.createCharacterTaskTimer, nil),
		routes.RegisterRoute(http.MethodDelete, "character_task_timer/:id", e.deleteCharacterTaskTimer, nil),
		routes.RegisterRoute(http.MethodPatch, "character_task_timer/:id", e.updateCharacterTaskTimer, nil),
		routes.RegisterRoute(http.MethodPost, "character_task_timers/bulk", e.getCharacterTaskTimersBulk, nil),
	}
}

// listCharacterTaskTimers godoc
// @Id listCharacterTaskTimers
// @Summary Lists CharacterTaskTimers
// @Accept json
// @Produce json
// @Tags CharacterTaskTimer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterTaskTimer
// @Failure 500 {string} string "Bad query request"
// @Router /character_task_timers [get]
func (e *CharacterTaskTimerController) listCharacterTaskTimers(c echo.Context) error {
	var results []models.CharacterTaskTimer
	err := e.db.QueryContext(models.CharacterTaskTimer{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterTaskTimer godoc
// @Id getCharacterTaskTimer
// @Summary Gets CharacterTaskTimer
// @Accept json
// @Produce json
// @Tags CharacterTaskTimer
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterTaskTimer
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_task_timer/{id} [get]
func (e *CharacterTaskTimerController) getCharacterTaskTimer(c echo.Context) error {
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
	var result models.CharacterTaskTimer
	query := e.db.QueryContext(models.CharacterTaskTimer{}, c)
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

// updateCharacterTaskTimer godoc
// @Id updateCharacterTaskTimer
// @Summary Updates CharacterTaskTimer
// @Accept json
// @Produce json
// @Tags CharacterTaskTimer
// @Param id path int true "Id"
// @Param character_task_timer body models.CharacterTaskTimer true "CharacterTaskTimer"
// @Success 200 {array} models.CharacterTaskTimer
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_task_timer/{id} [patch]
func (e *CharacterTaskTimerController) updateCharacterTaskTimer(c echo.Context) error {
	request := new(models.CharacterTaskTimer)
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
	var result models.CharacterTaskTimer
	query := e.db.QueryContext(models.CharacterTaskTimer{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterTaskTimer{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterTaskTimer godoc
// @Id createCharacterTaskTimer
// @Summary Creates CharacterTaskTimer
// @Accept json
// @Produce json
// @Param character_task_timer body models.CharacterTaskTimer true "CharacterTaskTimer"
// @Tags CharacterTaskTimer
// @Success 200 {array} models.CharacterTaskTimer
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_task_timer [put]
func (e *CharacterTaskTimerController) createCharacterTaskTimer(c echo.Context) error {
	characterTaskTimer := new(models.CharacterTaskTimer)
	if err := c.Bind(characterTaskTimer); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterTaskTimer{}, c).Model(&models.CharacterTaskTimer{}).Create(&characterTaskTimer).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterTaskTimer)
}

// deleteCharacterTaskTimer godoc
// @Id deleteCharacterTaskTimer
// @Summary Deletes CharacterTaskTimer
// @Accept json
// @Produce json
// @Tags CharacterTaskTimer
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_task_timer/{id} [delete]
func (e *CharacterTaskTimerController) deleteCharacterTaskTimer(c echo.Context) error {
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
	var result models.CharacterTaskTimer
	query := e.db.QueryContext(models.CharacterTaskTimer{}, c)
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

// getCharacterTaskTimersBulk godoc
// @Id getCharacterTaskTimersBulk
// @Summary Gets CharacterTaskTimers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterTaskTimer
// @Success 200 {array} models.CharacterTaskTimer
// @Failure 500 {string} string "Bad query request"
// @Router /character_task_timers/bulk [post]
func (e *CharacterTaskTimerController) getCharacterTaskTimersBulk(c echo.Context) error {
	var results []models.CharacterTaskTimer

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

	err := e.db.QueryContext(models.CharacterTaskTimer{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
