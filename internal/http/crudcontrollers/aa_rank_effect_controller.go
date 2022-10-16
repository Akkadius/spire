package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type AaRankEffectController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewAaRankEffectController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *AaRankEffectController {
	return &AaRankEffectController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *AaRankEffectController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "aa_rank_effect/:rankId", e.getAaRankEffect, nil),
		routes.RegisterRoute(http.MethodGet, "aa_rank_effects", e.listAaRankEffects, nil),
		routes.RegisterRoute(http.MethodPut, "aa_rank_effect", e.createAaRankEffect, nil),
		routes.RegisterRoute(http.MethodDelete, "aa_rank_effect/:rankId", e.deleteAaRankEffect, nil),
		routes.RegisterRoute(http.MethodPatch, "aa_rank_effect/:rankId", e.updateAaRankEffect, nil),
		routes.RegisterRoute(http.MethodPost, "aa_rank_effects/bulk", e.getAaRankEffectsBulk, nil),
	}
}

// listAaRankEffects godoc
// @Id listAaRankEffects
// @Summary Lists AaRankEffects
// @Accept json
// @Produce json
// @Tags AaRankEffect
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaRankEffect
// @Failure 500 {string} string "Bad query request"
// @Router /aa_rank_effects [get]
func (e *AaRankEffectController) listAaRankEffects(c echo.Context) error {
	var results []models.AaRankEffect
	err := e.db.QueryContext(models.AaRankEffect{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAaRankEffect godoc
// @Id getAaRankEffect
// @Summary Gets AaRankEffect
// @Accept json
// @Produce json
// @Tags AaRankEffect
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AaRankEffect
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /aa_rank_effect/{id} [get]
func (e *AaRankEffectController) getAaRankEffect(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rankId, err := strconv.Atoi(c.Param("rankId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [RankId]"})
	}
	params = append(params, rankId)
	keys = append(keys, "rank_id = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.AaRankEffect
	query := e.db.QueryContext(models.AaRankEffect{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.RankId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAaRankEffect godoc
// @Id updateAaRankEffect
// @Summary Updates AaRankEffect
// @Accept json
// @Produce json
// @Tags AaRankEffect
// @Param id path int true "Id"
// @Param aa_rank_effect body models.AaRankEffect true "AaRankEffect"
// @Success 200 {array} models.AaRankEffect
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /aa_rank_effect/{id} [patch]
func (e *AaRankEffectController) updateAaRankEffect(c echo.Context) error {
	request := new(models.AaRankEffect)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	rankId, err := strconv.Atoi(c.Param("rankId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [RankId]"})
	}
	params = append(params, rankId)
	keys = append(keys, "rank_id = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.AaRankEffect
	query := e.db.QueryContext(models.AaRankEffect{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [AaRankEffect] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createAaRankEffect godoc
// @Id createAaRankEffect
// @Summary Creates AaRankEffect
// @Accept json
// @Produce json
// @Param aa_rank_effect body models.AaRankEffect true "AaRankEffect"
// @Tags AaRankEffect
// @Success 200 {array} models.AaRankEffect
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /aa_rank_effect [put]
func (e *AaRankEffectController) createAaRankEffect(c echo.Context) error {
	aaRankEffect := new(models.AaRankEffect)
	if err := c.Bind(aaRankEffect); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AaRankEffect{}, c).Model(&models.AaRankEffect{}).Create(&aaRankEffect).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.AaRankEffect{}, aaRankEffect)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [AaRankEffect] [%v] data [%v]", aaRankEffect.RankId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, aaRankEffect)
}

// deleteAaRankEffect godoc
// @Id deleteAaRankEffect
// @Summary Deletes AaRankEffect
// @Accept json
// @Produce json
// @Tags AaRankEffect
// @Param id path int true "rankId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /aa_rank_effect/{id} [delete]
func (e *AaRankEffectController) deleteAaRankEffect(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	rankId, err := strconv.Atoi(c.Param("rankId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, rankId)
	keys = append(keys, "rank_id = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.AaRankEffect
	query := e.db.QueryContext(models.AaRankEffect{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [AaRankEffect] [%v] keys [%v]", result.RankId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getAaRankEffectsBulk godoc
// @Id getAaRankEffectsBulk
// @Summary Gets AaRankEffects in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AaRankEffect
// @Success 200 {array} models.AaRankEffect
// @Failure 500 {string} string "Bad query request"
// @Router /aa_rank_effects/bulk [post]
func (e *AaRankEffectController) getAaRankEffectsBulk(c echo.Context) error {
	var results []models.AaRankEffect

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

	err := e.db.QueryContext(models.AaRankEffect{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
