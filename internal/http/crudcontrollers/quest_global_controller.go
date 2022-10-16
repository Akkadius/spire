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

type QuestGlobalController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewQuestGlobalController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *QuestGlobalController {
	return &QuestGlobalController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *QuestGlobalController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "quest_global/:charid", e.getQuestGlobal, nil),
		routes.RegisterRoute(http.MethodGet, "quest_globals", e.listQuestGlobals, nil),
		routes.RegisterRoute(http.MethodPut, "quest_global", e.createQuestGlobal, nil),
		routes.RegisterRoute(http.MethodDelete, "quest_global/:charid", e.deleteQuestGlobal, nil),
		routes.RegisterRoute(http.MethodPatch, "quest_global/:charid", e.updateQuestGlobal, nil),
		routes.RegisterRoute(http.MethodPost, "quest_globals/bulk", e.getQuestGlobalsBulk, nil),
	}
}

// listQuestGlobals godoc
// @Id listQuestGlobals
// @Summary Lists QuestGlobals
// @Accept json
// @Produce json
// @Tags QuestGlobal
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.QuestGlobal
// @Failure 500 {string} string "Bad query request"
// @Router /quest_globals [get]
func (e *QuestGlobalController) listQuestGlobals(c echo.Context) error {
	var results []models.QuestGlobal
	err := e.db.QueryContext(models.QuestGlobal{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getQuestGlobal godoc
// @Id getQuestGlobal
// @Summary Gets QuestGlobal
// @Accept json
// @Produce json
// @Tags QuestGlobal
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.QuestGlobal
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /quest_global/{id} [get]
func (e *QuestGlobalController) getQuestGlobal(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [npcid] position [2] type [int]
	if len(c.QueryParam("npcid")) > 0 {
		npcidParam, err := strconv.Atoi(c.QueryParam("npcid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [npcid] err [%s]", err.Error())})
		}

		params = append(params, npcidParam)
		keys = append(keys, "npcid = ?")
	}

	// key param [zoneid] position [3] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [name] position [4] type [varchar]
	if len(c.QueryParam("name")) > 0 {
		nameParam, err := strconv.Atoi(c.QueryParam("name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [name] err [%s]", err.Error())})
		}

		params = append(params, nameParam)
		keys = append(keys, "name = ?")
	}

	// query builder
	var result models.QuestGlobal
	query := e.db.QueryContext(models.QuestGlobal{}, c)
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

// updateQuestGlobal godoc
// @Id updateQuestGlobal
// @Summary Updates QuestGlobal
// @Accept json
// @Produce json
// @Tags QuestGlobal
// @Param id path int true "Id"
// @Param quest_global body models.QuestGlobal true "QuestGlobal"
// @Success 200 {array} models.QuestGlobal
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /quest_global/{id} [patch]
func (e *QuestGlobalController) updateQuestGlobal(c echo.Context) error {
	request := new(models.QuestGlobal)
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

	// key param [npcid] position [2] type [int]
	if len(c.QueryParam("npcid")) > 0 {
		npcidParam, err := strconv.Atoi(c.QueryParam("npcid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [npcid] err [%s]", err.Error())})
		}

		params = append(params, npcidParam)
		keys = append(keys, "npcid = ?")
	}

	// key param [zoneid] position [3] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [name] position [4] type [varchar]
	if len(c.QueryParam("name")) > 0 {
		nameParam, err := strconv.Atoi(c.QueryParam("name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [name] err [%s]", err.Error())})
		}

		params = append(params, nameParam)
		keys = append(keys, "name = ?")
	}

	// query builder
	var result models.QuestGlobal
	query := e.db.QueryContext(models.QuestGlobal{}, c)
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
		event := fmt.Sprintf("Updated [QuestGlobal] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createQuestGlobal godoc
// @Id createQuestGlobal
// @Summary Creates QuestGlobal
// @Accept json
// @Produce json
// @Param quest_global body models.QuestGlobal true "QuestGlobal"
// @Tags QuestGlobal
// @Success 200 {array} models.QuestGlobal
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /quest_global [put]
func (e *QuestGlobalController) createQuestGlobal(c echo.Context) error {
	questGlobal := new(models.QuestGlobal)
	if err := c.Bind(questGlobal); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.QuestGlobal{}, c).Model(&models.QuestGlobal{}).Create(&questGlobal).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.QuestGlobal{}, questGlobal)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [QuestGlobal] [%v] data [%v]", questGlobal.Charid, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, questGlobal)
}

// deleteQuestGlobal godoc
// @Id deleteQuestGlobal
// @Summary Deletes QuestGlobal
// @Accept json
// @Produce json
// @Tags QuestGlobal
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /quest_global/{id} [delete]
func (e *QuestGlobalController) deleteQuestGlobal(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [npcid] position [2] type [int]
	if len(c.QueryParam("npcid")) > 0 {
		npcidParam, err := strconv.Atoi(c.QueryParam("npcid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [npcid] err [%s]", err.Error())})
		}

		params = append(params, npcidParam)
		keys = append(keys, "npcid = ?")
	}

	// key param [zoneid] position [3] type [int]
	if len(c.QueryParam("zoneid")) > 0 {
		zoneidParam, err := strconv.Atoi(c.QueryParam("zoneid"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [zoneid] err [%s]", err.Error())})
		}

		params = append(params, zoneidParam)
		keys = append(keys, "zoneid = ?")
	}

	// key param [name] position [4] type [varchar]
	if len(c.QueryParam("name")) > 0 {
		nameParam, err := strconv.Atoi(c.QueryParam("name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [name] err [%s]", err.Error())})
		}

		params = append(params, nameParam)
		keys = append(keys, "name = ?")
	}

	// query builder
	var result models.QuestGlobal
	query := e.db.QueryContext(models.QuestGlobal{}, c)
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
		event := fmt.Sprintf("Deleted [QuestGlobal] [%v] keys [%v]", result.Charid, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getQuestGlobalsBulk godoc
// @Id getQuestGlobalsBulk
// @Summary Gets QuestGlobals in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags QuestGlobal
// @Success 200 {array} models.QuestGlobal
// @Failure 500 {string} string "Bad query request"
// @Router /quest_globals/bulk [post]
func (e *QuestGlobalController) getQuestGlobalsBulk(c echo.Context) error {
	var results []models.QuestGlobal

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

	err := e.db.QueryContext(models.QuestGlobal{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
