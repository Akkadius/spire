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

type NpcTypesTintController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcTypesTintController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcTypesTintController {
	return &NpcTypesTintController{
		db:	    db,
		logger: logger,
	}
}

func (e *NpcTypesTintController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_types_tint/:id", e.getNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodGet, "npc_types_tints", e.listNpcTypesTints, nil),
		routes.RegisterRoute(http.MethodPut, "npc_types_tint", e.createNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_types_tint/:id", e.deleteNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_types_tint/:id", e.updateNpcTypesTint, nil),
		routes.RegisterRoute(http.MethodPost, "npc_types_tints/bulk", e.getNpcTypesTintsBulk, nil),
	}
}

// listNpcTypesTints godoc
// @Id listNpcTypesTints
// @Summary Lists NpcTypesTints
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcTypesTint
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types_tints [get]
func (e *NpcTypesTintController) listNpcTypesTints(c echo.Context) error {
	var results []models.NpcTypesTint
	err := e.db.QueryContext(models.NpcTypesTint{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcTypesTint godoc
// @Id getNpcTypesTint
// @Summary Gets NpcTypesTint
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcTypesTint
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types_tint/{id} [get]
func (e *NpcTypesTintController) getNpcTypesTint(c echo.Context) error {
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
	var result models.NpcTypesTint
	query := e.db.QueryContext(models.NpcTypesTint{}, c)
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

// updateNpcTypesTint godoc
// @Id updateNpcTypesTint
// @Summary Updates NpcTypesTint
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param id path int true "Id"
// @Param npc_types_tint body models.NpcTypesTint true "NpcTypesTint"
// @Success 200 {array} models.NpcTypesTint
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_types_tint/{id} [patch]
func (e *NpcTypesTintController) updateNpcTypesTint(c echo.Context) error {
	request := new(models.NpcTypesTint)
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
	var result models.NpcTypesTint
	query := e.db.QueryContext(models.NpcTypesTint{}, c)
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

// createNpcTypesTint godoc
// @Id createNpcTypesTint
// @Summary Creates NpcTypesTint
// @Accept json
// @Produce json
// @Param npc_types_tint body models.NpcTypesTint true "NpcTypesTint"
// @Tags NpcTypesTint
// @Success 200 {array} models.NpcTypesTint
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_types_tint [put]
func (e *NpcTypesTintController) createNpcTypesTint(c echo.Context) error {
	npcTypesTint := new(models.NpcTypesTint)
	if err := c.Bind(npcTypesTint); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcTypesTint{}, c).Model(&models.NpcTypesTint{}).Create(&npcTypesTint).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcTypesTint)
}

// deleteNpcTypesTint godoc
// @Id deleteNpcTypesTint
// @Summary Deletes NpcTypesTint
// @Accept json
// @Produce json
// @Tags NpcTypesTint
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_types_tint/{id} [delete]
func (e *NpcTypesTintController) deleteNpcTypesTint(c echo.Context) error {
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
	var result models.NpcTypesTint
	query := e.db.QueryContext(models.NpcTypesTint{}, c)
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

// getNpcTypesTintsBulk godoc
// @Id getNpcTypesTintsBulk
// @Summary Gets NpcTypesTints in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcTypesTint
// @Success 200 {array} models.NpcTypesTint
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types_tints/bulk [post]
func (e *NpcTypesTintController) getNpcTypesTintsBulk(c echo.Context) error {
	var results []models.NpcTypesTint

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

	err := e.db.QueryContext(models.NpcTypesTint{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
