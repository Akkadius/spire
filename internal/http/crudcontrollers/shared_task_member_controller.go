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

type SharedTaskMemberController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSharedTaskMemberController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SharedTaskMemberController {
	return &SharedTaskMemberController{
		db:	 db,
		logger: logger,
	}
}

func (e *SharedTaskMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "shared_task_member/:sharedTaskId", e.getSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodGet, "shared_task_members", e.listSharedTaskMembers, nil),
		routes.RegisterRoute(http.MethodPut, "shared_task_member", e.createSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodDelete, "shared_task_member/:sharedTaskId", e.deleteSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodPatch, "shared_task_member/:sharedTaskId", e.updateSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodPost, "shared_task_members/bulk", e.getSharedTaskMembersBulk, nil),
	}
}

// listSharedTaskMembers godoc
// @Id listSharedTaskMembers
// @Summary Lists SharedTaskMembers
// @Accept json
// @Produce json
// @Tags SharedTaskMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTaskMember
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_members [get]
func (e *SharedTaskMemberController) listSharedTaskMembers(c echo.Context) error {
	var results []models.SharedTaskMember
	err := e.db.QueryContext(models.SharedTaskMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSharedTaskMember godoc
// @Id getSharedTaskMember
// @Summary Gets SharedTaskMember
// @Accept json
// @Produce json
// @Tags SharedTaskMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SharedTaskMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_member/{id} [get]
func (e *SharedTaskMemberController) getSharedTaskMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SharedTaskId]"})
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [character_id] position [2] type [bigint]
	if len(c.QueryParam("character_id")) > 0 {
		characterIdParam, err := strconv.Atoi(c.QueryParam("character_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [character_id] err [%s]", err.Error())})
		}

		params = append(params, characterIdParam)
		keys = append(keys, "character_id = ?")
	}

	// query builder
	var result models.SharedTaskMember
	query := e.db.QueryContext(models.SharedTaskMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.SharedTaskId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSharedTaskMember godoc
// @Id updateSharedTaskMember
// @Summary Updates SharedTaskMember
// @Accept json
// @Produce json
// @Tags SharedTaskMember
// @Param id path int true "Id"
// @Param shared_task_member body models.SharedTaskMember true "SharedTaskMember"
// @Success 200 {array} models.SharedTaskMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /shared_task_member/{id} [patch]
func (e *SharedTaskMemberController) updateSharedTaskMember(c echo.Context) error {
	request := new(models.SharedTaskMember)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SharedTaskId]"})
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [character_id] position [2] type [bigint]
	if len(c.QueryParam("character_id")) > 0 {
		characterIdParam, err := strconv.Atoi(c.QueryParam("character_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [character_id] err [%s]", err.Error())})
		}

		params = append(params, characterIdParam)
		keys = append(keys, "character_id = ?")
	}

	// query builder
	var result models.SharedTaskMember
	query := e.db.QueryContext(models.SharedTaskMember{}, c)
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

// createSharedTaskMember godoc
// @Id createSharedTaskMember
// @Summary Creates SharedTaskMember
// @Accept json
// @Produce json
// @Param shared_task_member body models.SharedTaskMember true "SharedTaskMember"
// @Tags SharedTaskMember
// @Success 200 {array} models.SharedTaskMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /shared_task_member [put]
func (e *SharedTaskMemberController) createSharedTaskMember(c echo.Context) error {
	sharedTaskMember := new(models.SharedTaskMember)
	if err := c.Bind(sharedTaskMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SharedTaskMember{}, c).Model(&models.SharedTaskMember{}).Create(&sharedTaskMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, sharedTaskMember)
}

// deleteSharedTaskMember godoc
// @Id deleteSharedTaskMember
// @Summary Deletes SharedTaskMember
// @Accept json
// @Produce json
// @Tags SharedTaskMember
// @Param id path int true "sharedTaskId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /shared_task_member/{id} [delete]
func (e *SharedTaskMemberController) deleteSharedTaskMember(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	sharedTaskId, err := strconv.Atoi(c.Param("sharedTaskId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, sharedTaskId)
	keys = append(keys, "shared_task_id = ?")

	// key param [character_id] position [2] type [bigint]
	if len(c.QueryParam("character_id")) > 0 {
		characterIdParam, err := strconv.Atoi(c.QueryParam("character_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [character_id] err [%s]", err.Error())})
		}

		params = append(params, characterIdParam)
		keys = append(keys, "character_id = ?")
	}

	// query builder
	var result models.SharedTaskMember
	query := e.db.QueryContext(models.SharedTaskMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.SharedTaskMember{}, c).Model(&models.SharedTaskMember{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSharedTaskMembersBulk godoc
// @Id getSharedTaskMembersBulk
// @Summary Gets SharedTaskMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SharedTaskMember
// @Success 200 {array} models.SharedTaskMember
// @Failure 500 {string} string "Bad query request"
// @Router /shared_task_members/bulk [post]
func (e *SharedTaskMemberController) getSharedTaskMembersBulk(c echo.Context) error {
	var results []models.SharedTaskMember

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

	err := e.db.QueryContext(models.SharedTaskMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
