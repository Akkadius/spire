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

type LevelExpModController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLevelExpModController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LevelExpModController {
	return &LevelExpModController{
		db:	    db,
		logger: logger,
	}
}

func (e *LevelExpModController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "level_exp_mod/:level", e.getLevelExpMod, nil),
		routes.RegisterRoute(http.MethodGet, "level_exp_mods", e.listLevelExpMods, nil),
		routes.RegisterRoute(http.MethodPut, "level_exp_mod", e.createLevelExpMod, nil),
		routes.RegisterRoute(http.MethodDelete, "level_exp_mod/:level", e.deleteLevelExpMod, nil),
		routes.RegisterRoute(http.MethodPatch, "level_exp_mod/:level", e.updateLevelExpMod, nil),
		routes.RegisterRoute(http.MethodPost, "level_exp_mods/bulk", e.getLevelExpModsBulk, nil),
	}
}

// listLevelExpMods godoc
// @Id listLevelExpMods
// @Summary Lists LevelExpMods
// @Accept json
// @Produce json
// @Tags LevelExpMod
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LevelExpMod
// @Failure 500 {string} string "Bad query request"
// @Router /level_exp_mods [get]
func (e *LevelExpModController) listLevelExpMods(c echo.Context) error {
	var results []models.LevelExpMod
	err := e.db.QueryContext(models.LevelExpMod{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLevelExpMod godoc
// @Id getLevelExpMod
// @Summary Gets LevelExpMod
// @Accept json
// @Produce json
// @Tags LevelExpMod
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LevelExpMod
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /level_exp_mod/{id} [get]
func (e *LevelExpModController) getLevelExpMod(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Level]"})
	}
	params = append(params, level)
	keys = append(keys, "level = ?")

	// query builder
	var result models.LevelExpMod
	query := e.db.QueryContext(models.LevelExpMod{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Level == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLevelExpMod godoc
// @Id updateLevelExpMod
// @Summary Updates LevelExpMod
// @Accept json
// @Produce json
// @Tags LevelExpMod
// @Param id path int true "Id"
// @Param level_exp_mod body models.LevelExpMod true "LevelExpMod"
// @Success 200 {array} models.LevelExpMod
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /level_exp_mod/{id} [patch]
func (e *LevelExpModController) updateLevelExpMod(c echo.Context) error {
	request := new(models.LevelExpMod)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Level]"})
	}
	params = append(params, level)
	keys = append(keys, "level = ?")

	// query builder
	var result models.LevelExpMod
	query := e.db.QueryContext(models.LevelExpMod{}, c)
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

// createLevelExpMod godoc
// @Id createLevelExpMod
// @Summary Creates LevelExpMod
// @Accept json
// @Produce json
// @Param level_exp_mod body models.LevelExpMod true "LevelExpMod"
// @Tags LevelExpMod
// @Success 200 {array} models.LevelExpMod
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /level_exp_mod [put]
func (e *LevelExpModController) createLevelExpMod(c echo.Context) error {
	levelExpMod := new(models.LevelExpMod)
	if err := c.Bind(levelExpMod); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LevelExpMod{}, c).Model(&models.LevelExpMod{}).Create(&levelExpMod).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, levelExpMod)
}

// deleteLevelExpMod godoc
// @Id deleteLevelExpMod
// @Summary Deletes LevelExpMod
// @Accept json
// @Produce json
// @Tags LevelExpMod
// @Param id path int true "level"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /level_exp_mod/{id} [delete]
func (e *LevelExpModController) deleteLevelExpMod(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, level)
	keys = append(keys, "level = ?")

	// query builder
	var result models.LevelExpMod
	query := e.db.QueryContext(models.LevelExpMod{}, c)
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

// getLevelExpModsBulk godoc
// @Id getLevelExpModsBulk
// @Summary Gets LevelExpMods in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LevelExpMod
// @Success 200 {array} models.LevelExpMod
// @Failure 500 {string} string "Bad query request"
// @Router /level_exp_mods/bulk [post]
func (e *LevelExpModController) getLevelExpModsBulk(c echo.Context) error {
	var results []models.LevelExpMod

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

	err := e.db.QueryContext(models.LevelExpMod{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
