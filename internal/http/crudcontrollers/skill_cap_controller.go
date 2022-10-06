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

type SkillCapController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSkillCapController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SkillCapController {
	return &SkillCapController{
		db:	    db,
		logger: logger,
	}
}

func (e *SkillCapController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "skill_cap/:skillID", e.getSkillCap, nil),
		routes.RegisterRoute(http.MethodGet, "skill_caps", e.listSkillCaps, nil),
		routes.RegisterRoute(http.MethodPut, "skill_cap", e.createSkillCap, nil),
		routes.RegisterRoute(http.MethodDelete, "skill_cap/:skillID", e.deleteSkillCap, nil),
		routes.RegisterRoute(http.MethodPatch, "skill_cap/:skillID", e.updateSkillCap, nil),
		routes.RegisterRoute(http.MethodPost, "skill_caps/bulk", e.getSkillCapsBulk, nil),
	}
}

// listSkillCaps godoc
// @Id listSkillCaps
// @Summary Lists SkillCaps
// @Accept json
// @Produce json
// @Tags SkillCap
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SkillCap
// @Failure 500 {string} string "Bad query request"
// @Router /skill_caps [get]
func (e *SkillCapController) listSkillCaps(c echo.Context) error {
	var results []models.SkillCap
	err := e.db.QueryContext(models.SkillCap{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSkillCap godoc
// @Id getSkillCap
// @Summary Gets SkillCap
// @Accept json
// @Produce json
// @Tags SkillCap
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SkillCap
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /skill_cap/{id} [get]
func (e *SkillCapController) getSkillCap(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	skillID, err := strconv.Atoi(c.Param("skillID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SkillID]"})
	}
	params = append(params, skillID)
	keys = append(keys, "skillID = ?")

	// key param [class] position [2] type [tinyint]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// key param [level] position [3] type [tinyint]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// key param [class_] position [5] type [tinyint]
	if len(c.QueryParam("class_")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class_"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class_] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class_ = ?")
	}

	// query builder
	var result models.SkillCap
	query := e.db.QueryContext(models.SkillCap{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.SkillID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSkillCap godoc
// @Id updateSkillCap
// @Summary Updates SkillCap
// @Accept json
// @Produce json
// @Tags SkillCap
// @Param id path int true "Id"
// @Param skill_cap body models.SkillCap true "SkillCap"
// @Success 200 {array} models.SkillCap
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /skill_cap/{id} [patch]
func (e *SkillCapController) updateSkillCap(c echo.Context) error {
	request := new(models.SkillCap)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	skillID, err := strconv.Atoi(c.Param("skillID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SkillID]"})
	}
	params = append(params, skillID)
	keys = append(keys, "skillID = ?")

	// key param [class] position [2] type [tinyint]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// key param [level] position [3] type [tinyint]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// key param [class_] position [5] type [tinyint]
	if len(c.QueryParam("class_")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class_"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class_] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class_ = ?")
	}

	// query builder
	var result models.SkillCap
	query := e.db.QueryContext(models.SkillCap{}, c)
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

// createSkillCap godoc
// @Id createSkillCap
// @Summary Creates SkillCap
// @Accept json
// @Produce json
// @Param skill_cap body models.SkillCap true "SkillCap"
// @Tags SkillCap
// @Success 200 {array} models.SkillCap
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /skill_cap [put]
func (e *SkillCapController) createSkillCap(c echo.Context) error {
	skillCap := new(models.SkillCap)
	if err := c.Bind(skillCap); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SkillCap{}, c).Model(&models.SkillCap{}).Create(&skillCap).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, skillCap)
}

// deleteSkillCap godoc
// @Id deleteSkillCap
// @Summary Deletes SkillCap
// @Accept json
// @Produce json
// @Tags SkillCap
// @Param id path int true "skillID"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /skill_cap/{id} [delete]
func (e *SkillCapController) deleteSkillCap(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	skillID, err := strconv.Atoi(c.Param("skillID"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, skillID)
	keys = append(keys, "skillID = ?")

	// key param [class] position [2] type [tinyint]
	if len(c.QueryParam("class")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class = ?")
	}

	// key param [level] position [3] type [tinyint]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// key param [class_] position [5] type [tinyint]
	if len(c.QueryParam("class_")) > 0 {
		classParam, err := strconv.Atoi(c.QueryParam("class_"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [class_] err [%s]", err.Error())})
		}

		params = append(params, classParam)
		keys = append(keys, "class_ = ?")
	}

	// query builder
	var result models.SkillCap
	query := e.db.QueryContext(models.SkillCap{}, c)
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

// getSkillCapsBulk godoc
// @Id getSkillCapsBulk
// @Summary Gets SkillCaps in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SkillCap
// @Success 200 {array} models.SkillCap
// @Failure 500 {string} string "Bad query request"
// @Router /skill_caps/bulk [post]
func (e *SkillCapController) getSkillCapsBulk(c echo.Context) error {
	var results []models.SkillCap

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

	err := e.db.QueryContext(models.SkillCap{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
