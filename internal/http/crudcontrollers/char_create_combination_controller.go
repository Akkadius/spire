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

type CharCreateCombinationController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharCreateCombinationController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharCreateCombinationController {
	return &CharCreateCombinationController{
		db:	    db,
		logger: logger,
	}
}

func (e *CharCreateCombinationController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "char_create_combination/:race", e.getCharCreateCombination, nil),
		routes.RegisterRoute(http.MethodGet, "char_create_combinations", e.listCharCreateCombinations, nil),
		routes.RegisterRoute(http.MethodPut, "char_create_combination", e.createCharCreateCombination, nil),
		routes.RegisterRoute(http.MethodDelete, "char_create_combination/:race", e.deleteCharCreateCombination, nil),
		routes.RegisterRoute(http.MethodPatch, "char_create_combination/:race", e.updateCharCreateCombination, nil),
		routes.RegisterRoute(http.MethodPost, "char_create_combinations/bulk", e.getCharCreateCombinationsBulk, nil),
	}
}

// listCharCreateCombinations godoc
// @Id listCharCreateCombinations
// @Summary Lists CharCreateCombinations
// @Accept json
// @Produce json
// @Tags CharCreateCombination
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharCreateCombination
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_combinations [get]
func (e *CharCreateCombinationController) listCharCreateCombinations(c echo.Context) error {
	var results []models.CharCreateCombination
	err := e.db.QueryContext(models.CharCreateCombination{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharCreateCombination godoc
// @Id getCharCreateCombination
// @Summary Gets CharCreateCombination
// @Accept json
// @Produce json
// @Tags CharCreateCombination
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharCreateCombination
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_combination/{id} [get]
func (e *CharCreateCombinationController) getCharCreateCombination(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	race, err := strconv.Atoi(c.Param("race"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Race]"})
	}
	params = append(params, race)
	keys = append(keys, "race = ?")

	// key param [class] position [3] type [int]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// key param [deity] position [4] type [int]
	if len(c.QueryParam("deity")) > 0 {
		deityParam, err := strconv.Atoi(c.QueryParam("deity"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [deity] err [%s]", err.Error())})
		}

		params = append(params, deityParam)
		keys = append(keys, "deity = ?")
	}

	// key param [start_zone] position [5] type [int]
	if len(c.QueryParam("start_zone")) > 0 {
		startZoneParam, err := strconv.Atoi(c.QueryParam("start_zone"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [start_zone] err [%s]", err.Error())})
		}

		params = append(params, startZoneParam)
		keys = append(keys, "start_zone = ?")
	}

	// query builder
	var result models.CharCreateCombination
	query := e.db.QueryContext(models.CharCreateCombination{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Race == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharCreateCombination godoc
// @Id updateCharCreateCombination
// @Summary Updates CharCreateCombination
// @Accept json
// @Produce json
// @Tags CharCreateCombination
// @Param id path int true "Id"
// @Param char_create_combination body models.CharCreateCombination true "CharCreateCombination"
// @Success 200 {array} models.CharCreateCombination
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /char_create_combination/{id} [patch]
func (e *CharCreateCombinationController) updateCharCreateCombination(c echo.Context) error {
	request := new(models.CharCreateCombination)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	race, err := strconv.Atoi(c.Param("race"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Race]"})
	}
	params = append(params, race)
	keys = append(keys, "race = ?")

	// key param [class] position [3] type [int]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// key param [deity] position [4] type [int]
	if len(c.QueryParam("deity")) > 0 {
		deityParam, err := strconv.Atoi(c.QueryParam("deity"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [deity] err [%s]", err.Error())})
		}

		params = append(params, deityParam)
		keys = append(keys, "deity = ?")
	}

	// key param [start_zone] position [5] type [int]
	if len(c.QueryParam("start_zone")) > 0 {
		startZoneParam, err := strconv.Atoi(c.QueryParam("start_zone"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [start_zone] err [%s]", err.Error())})
		}

		params = append(params, startZoneParam)
		keys = append(keys, "start_zone = ?")
	}

	// query builder
	var result models.CharCreateCombination
	query := e.db.QueryContext(models.CharCreateCombination{}, c)
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

// createCharCreateCombination godoc
// @Id createCharCreateCombination
// @Summary Creates CharCreateCombination
// @Accept json
// @Produce json
// @Param char_create_combination body models.CharCreateCombination true "CharCreateCombination"
// @Tags CharCreateCombination
// @Success 200 {array} models.CharCreateCombination
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /char_create_combination [put]
func (e *CharCreateCombinationController) createCharCreateCombination(c echo.Context) error {
	charCreateCombination := new(models.CharCreateCombination)
	if err := c.Bind(charCreateCombination); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharCreateCombination{}, c).Model(&models.CharCreateCombination{}).Create(&charCreateCombination).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, charCreateCombination)
}

// deleteCharCreateCombination godoc
// @Id deleteCharCreateCombination
// @Summary Deletes CharCreateCombination
// @Accept json
// @Produce json
// @Tags CharCreateCombination
// @Param id path int true "race"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /char_create_combination/{id} [delete]
func (e *CharCreateCombinationController) deleteCharCreateCombination(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	race, err := strconv.Atoi(c.Param("race"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, race)
	keys = append(keys, "race = ?")

	// key param [class] position [3] type [int]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// key param [deity] position [4] type [int]
	if len(c.QueryParam("deity")) > 0 {
		deityParam, err := strconv.Atoi(c.QueryParam("deity"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [deity] err [%s]", err.Error())})
		}

		params = append(params, deityParam)
		keys = append(keys, "deity = ?")
	}

	// key param [start_zone] position [5] type [int]
	if len(c.QueryParam("start_zone")) > 0 {
		startZoneParam, err := strconv.Atoi(c.QueryParam("start_zone"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [start_zone] err [%s]", err.Error())})
		}

		params = append(params, startZoneParam)
		keys = append(keys, "start_zone = ?")
	}

	// query builder
	var result models.CharCreateCombination
	query := e.db.QueryContext(models.CharCreateCombination{}, c)
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

// getCharCreateCombinationsBulk godoc
// @Id getCharCreateCombinationsBulk
// @Summary Gets CharCreateCombinations in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharCreateCombination
// @Success 200 {array} models.CharCreateCombination
// @Failure 500 {string} string "Bad query request"
// @Router /char_create_combinations/bulk [post]
func (e *CharCreateCombinationController) getCharCreateCombinationsBulk(c echo.Context) error {
	var results []models.CharCreateCombination

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

	err := e.db.QueryContext(models.CharCreateCombination{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
