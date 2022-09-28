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

type CompletedSharedTaskMemberController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCompletedSharedTaskMemberController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CompletedSharedTaskMemberController {
	return &CompletedSharedTaskMemberController{
		db:	 db,
		logger: logger,
	}
}

func (e *CompletedSharedTaskMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "completed_shared_task_member/:sharedTaskId", e.getCompletedSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodGet, "completed_shared_task_members", e.listCompletedSharedTaskMembers, nil),
		routes.RegisterRoute(http.MethodPut, "completed_shared_task_member", e.createCompletedSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodDelete, "completed_shared_task_member/:sharedTaskId", e.deleteCompletedSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodPatch, "completed_shared_task_member/:sharedTaskId", e.updateCompletedSharedTaskMember, nil),
		routes.RegisterRoute(http.MethodPost, "completed_shared_task_members/bulk", e.getCompletedSharedTaskMembersBulk, nil),
	}
}

// listCompletedSharedTaskMembers godoc
// @Id listCompletedSharedTaskMembers
// @Summary Lists CompletedSharedTaskMembers
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedSharedTaskMember
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_task_members [get]
func (e *CompletedSharedTaskMemberController) listCompletedSharedTaskMembers(c echo.Context) error {
	var results []models.CompletedSharedTaskMember
	err := e.db.QueryContext(models.CompletedSharedTaskMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCompletedSharedTaskMember godoc
// @Id getCompletedSharedTaskMember
// @Summary Gets CompletedSharedTaskMember
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CompletedSharedTaskMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_task_member/{id} [get]
func (e *CompletedSharedTaskMemberController) getCompletedSharedTaskMember(c echo.Context) error {
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
	var result models.CompletedSharedTaskMember
	query := e.db.QueryContext(models.CompletedSharedTaskMember{}, c)
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

// updateCompletedSharedTaskMember godoc
// @Id updateCompletedSharedTaskMember
// @Summary Updates CompletedSharedTaskMember
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskMember
// @Param id path int true "Id"
// @Param completed_shared_task_member body models.CompletedSharedTaskMember true "CompletedSharedTaskMember"
// @Success 200 {array} models.CompletedSharedTaskMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /completed_shared_task_member/{id} [patch]
func (e *CompletedSharedTaskMemberController) updateCompletedSharedTaskMember(c echo.Context) error {
	request := new(models.CompletedSharedTaskMember)
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
	var result models.CompletedSharedTaskMember
	query := e.db.QueryContext(models.CompletedSharedTaskMember{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CompletedSharedTaskMember{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCompletedSharedTaskMember godoc
// @Id createCompletedSharedTaskMember
// @Summary Creates CompletedSharedTaskMember
// @Accept json
// @Produce json
// @Param completed_shared_task_member body models.CompletedSharedTaskMember true "CompletedSharedTaskMember"
// @Tags CompletedSharedTaskMember
// @Success 200 {array} models.CompletedSharedTaskMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /completed_shared_task_member [put]
func (e *CompletedSharedTaskMemberController) createCompletedSharedTaskMember(c echo.Context) error {
	completedSharedTaskMember := new(models.CompletedSharedTaskMember)
	if err := c.Bind(completedSharedTaskMember); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CompletedSharedTaskMember{}, c).Model(&models.CompletedSharedTaskMember{}).Create(&completedSharedTaskMember).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, completedSharedTaskMember)
}

// deleteCompletedSharedTaskMember godoc
// @Id deleteCompletedSharedTaskMember
// @Summary Deletes CompletedSharedTaskMember
// @Accept json
// @Produce json
// @Tags CompletedSharedTaskMember
// @Param id path int true "sharedTaskId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /completed_shared_task_member/{id} [delete]
func (e *CompletedSharedTaskMemberController) deleteCompletedSharedTaskMember(c echo.Context) error {
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
	var result models.CompletedSharedTaskMember
	query := e.db.QueryContext(models.CompletedSharedTaskMember{}, c)
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

// getCompletedSharedTaskMembersBulk godoc
// @Id getCompletedSharedTaskMembersBulk
// @Summary Gets CompletedSharedTaskMembers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CompletedSharedTaskMember
// @Success 200 {array} models.CompletedSharedTaskMember
// @Failure 500 {string} string "Bad query request"
// @Router /completed_shared_task_members/bulk [post]
func (e *CompletedSharedTaskMemberController) getCompletedSharedTaskMembersBulk(c echo.Context) error {
	var results []models.CompletedSharedTaskMember

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

	err := e.db.QueryContext(models.CompletedSharedTaskMember{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
