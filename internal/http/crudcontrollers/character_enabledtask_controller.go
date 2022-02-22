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

type CharacterEnabledtaskController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterEnabledtaskController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterEnabledtaskController {
	return &CharacterEnabledtaskController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterEnabledtaskController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_enabledtask/:charid", e.getCharacterEnabledtask, nil),
		routes.RegisterRoute(http.MethodGet, "character_enabledtasks", e.listCharacterEnabledtasks, nil),
		routes.RegisterRoute(http.MethodPut, "character_enabledtask", e.createCharacterEnabledtask, nil),
		routes.RegisterRoute(http.MethodDelete, "character_enabledtask/:charid", e.deleteCharacterEnabledtask, nil),
		routes.RegisterRoute(http.MethodPatch, "character_enabledtask/:charid", e.updateCharacterEnabledtask, nil),
		routes.RegisterRoute(http.MethodPost, "character_enabledtasks/bulk", e.getCharacterEnabledtasksBulk, nil),
	}
}

// listCharacterEnabledtasks godoc
// @Id listCharacterEnabledtasks
// @Summary Lists CharacterEnabledtasks
// @Accept json
// @Produce json
// @Tags CharacterEnabledtask
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterEnabledtask
// @Failure 500 {string} string "Bad query request"
// @Router /character_enabledtasks [get]
func (e *CharacterEnabledtaskController) listCharacterEnabledtasks(c echo.Context) error {
	var results []models.CharacterEnabledtask
	err := e.db.QueryContext(models.CharacterEnabledtask{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterEnabledtask godoc
// @Id getCharacterEnabledtask
// @Summary Gets CharacterEnabledtask
// @Accept json
// @Produce json
// @Tags CharacterEnabledtask
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterEnabledtask
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_enabledtask/{id} [get]
func (e *CharacterEnabledtaskController) getCharacterEnabledtask(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [taskid] position [2] type [int]
	if len(c.QueryParam("taskid")) > 0 {
		taskidParam, err := strconv.Atoi(c.QueryParam("taskid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [taskid] err [%s]", err.Error())})
		}

		params = append(params, taskidParam)
		keys = append(keys, "taskid = ?")
	}

	// query builder
	var result models.CharacterEnabledtask
	query := e.db.QueryContext(models.CharacterEnabledtask{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Charid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterEnabledtask godoc
// @Id updateCharacterEnabledtask
// @Summary Updates CharacterEnabledtask
// @Accept json
// @Produce json
// @Tags CharacterEnabledtask
// @Param id path int true "Id"
// @Param character_enabledtask body models.CharacterEnabledtask true "CharacterEnabledtask"
// @Success 200 {array} models.CharacterEnabledtask
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_enabledtask/{id} [patch]
func (e *CharacterEnabledtaskController) updateCharacterEnabledtask(c echo.Context) error {
	request := new(models.CharacterEnabledtask)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [taskid] position [2] type [int]
	if len(c.QueryParam("taskid")) > 0 {
		taskidParam, err := strconv.Atoi(c.QueryParam("taskid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [taskid] err [%s]", err.Error())})
		}

		params = append(params, taskidParam)
		keys = append(keys, "taskid = ?")
	}

	// query builder
	var result models.CharacterEnabledtask
	query := e.db.QueryContext(models.CharacterEnabledtask{}, c)
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

// createCharacterEnabledtask godoc
// @Id createCharacterEnabledtask
// @Summary Creates CharacterEnabledtask
// @Accept json
// @Produce json
// @Param character_enabledtask body models.CharacterEnabledtask true "CharacterEnabledtask"
// @Tags CharacterEnabledtask
// @Success 200 {array} models.CharacterEnabledtask
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_enabledtask [put]
func (e *CharacterEnabledtaskController) createCharacterEnabledtask(c echo.Context) error {
	characterEnabledtask := new(models.CharacterEnabledtask)
	if err := c.Bind(characterEnabledtask); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterEnabledtask{}, c).Model(&models.CharacterEnabledtask{}).Create(&characterEnabledtask).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterEnabledtask)
}

// deleteCharacterEnabledtask godoc
// @Id deleteCharacterEnabledtask
// @Summary Deletes CharacterEnabledtask
// @Accept json
// @Produce json
// @Tags CharacterEnabledtask
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_enabledtask/{id} [delete]
func (e *CharacterEnabledtaskController) deleteCharacterEnabledtask(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [taskid] position [2] type [int]
	if len(c.QueryParam("taskid")) > 0 {
		taskidParam, err := strconv.Atoi(c.QueryParam("taskid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [taskid] err [%s]", err.Error())})
		}

		params = append(params, taskidParam)
		keys = append(keys, "taskid = ?")
	}

	// query builder
	var result models.CharacterEnabledtask
	query := e.db.QueryContext(models.CharacterEnabledtask{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.CharacterEnabledtask{}, c).Model(&models.CharacterEnabledtask{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterEnabledtasksBulk godoc
// @Id getCharacterEnabledtasksBulk
// @Summary Gets CharacterEnabledtasks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterEnabledtask
// @Success 200 {array} models.CharacterEnabledtask
// @Failure 500 {string} string "Bad query request"
// @Router /character_enabledtasks/bulk [post]
func (e *CharacterEnabledtaskController) getCharacterEnabledtasksBulk(c echo.Context) error {
	var results []models.CharacterEnabledtask

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

	err := e.db.QueryContext(models.CharacterEnabledtask{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
