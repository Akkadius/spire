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

type VeteranRewardTemplateController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewVeteranRewardTemplateController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *VeteranRewardTemplateController {
	return &VeteranRewardTemplateController{
		db:	 db,
		logger: logger,
	}
}

func (e *VeteranRewardTemplateController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "veteran_reward_template/:claimId", e.getVeteranRewardTemplate, nil),
		routes.RegisterRoute(http.MethodGet, "veteran_reward_templates", e.listVeteranRewardTemplates, nil),
		routes.RegisterRoute(http.MethodPut, "veteran_reward_template", e.createVeteranRewardTemplate, nil),
		routes.RegisterRoute(http.MethodDelete, "veteran_reward_template/:claimId", e.deleteVeteranRewardTemplate, nil),
		routes.RegisterRoute(http.MethodPatch, "veteran_reward_template/:claimId", e.updateVeteranRewardTemplate, nil),
		routes.RegisterRoute(http.MethodPost, "veteran_reward_templates/bulk", e.getVeteranRewardTemplatesBulk, nil),
	}
}

// listVeteranRewardTemplates godoc
// @Id listVeteranRewardTemplates
// @Summary Lists VeteranRewardTemplates
// @Accept json
// @Produce json
// @Tags VeteranRewardTemplate
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.VeteranRewardTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /veteran_reward_templates [get]
func (e *VeteranRewardTemplateController) listVeteranRewardTemplates(c echo.Context) error {
	var results []models.VeteranRewardTemplate
	err := e.db.QueryContext(models.VeteranRewardTemplate{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getVeteranRewardTemplate godoc
// @Id getVeteranRewardTemplate
// @Summary Gets VeteranRewardTemplate
// @Accept json
// @Produce json
// @Tags VeteranRewardTemplate
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.VeteranRewardTemplate
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /veteran_reward_template/{id} [get]
func (e *VeteranRewardTemplateController) getVeteranRewardTemplate(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	claimId, err := strconv.Atoi(c.Param("claimId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ClaimId]"})
	}
	params = append(params, claimId)
	keys = append(keys, "claim_id = ?")

	// key param [reward_slot] position [5] type [tinyint]
	if len(c.QueryParam("reward_slot")) > 0 {
		rewardSlotParam, err := strconv.Atoi(c.QueryParam("reward_slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [reward_slot] err [%s]", err.Error())})
		}

		params = append(params, rewardSlotParam)
		keys = append(keys, "reward_slot = ?")
	}

	// query builder
	var result models.VeteranRewardTemplate
	query := e.db.QueryContext(models.VeteranRewardTemplate{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ClaimId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateVeteranRewardTemplate godoc
// @Id updateVeteranRewardTemplate
// @Summary Updates VeteranRewardTemplate
// @Accept json
// @Produce json
// @Tags VeteranRewardTemplate
// @Param id path int true "Id"
// @Param veteran_reward_template body models.VeteranRewardTemplate true "VeteranRewardTemplate"
// @Success 200 {array} models.VeteranRewardTemplate
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /veteran_reward_template/{id} [patch]
func (e *VeteranRewardTemplateController) updateVeteranRewardTemplate(c echo.Context) error {
	request := new(models.VeteranRewardTemplate)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	claimId, err := strconv.Atoi(c.Param("claimId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [ClaimId]"})
	}
	params = append(params, claimId)
	keys = append(keys, "claim_id = ?")

	// key param [reward_slot] position [5] type [tinyint]
	if len(c.QueryParam("reward_slot")) > 0 {
		rewardSlotParam, err := strconv.Atoi(c.QueryParam("reward_slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [reward_slot] err [%s]", err.Error())})
		}

		params = append(params, rewardSlotParam)
		keys = append(keys, "reward_slot = ?")
	}

	// query builder
	var result models.VeteranRewardTemplate
	query := e.db.QueryContext(models.VeteranRewardTemplate{}, c)
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

// createVeteranRewardTemplate godoc
// @Id createVeteranRewardTemplate
// @Summary Creates VeteranRewardTemplate
// @Accept json
// @Produce json
// @Param veteran_reward_template body models.VeteranRewardTemplate true "VeteranRewardTemplate"
// @Tags VeteranRewardTemplate
// @Success 200 {array} models.VeteranRewardTemplate
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /veteran_reward_template [put]
func (e *VeteranRewardTemplateController) createVeteranRewardTemplate(c echo.Context) error {
	veteranRewardTemplate := new(models.VeteranRewardTemplate)
	if err := c.Bind(veteranRewardTemplate); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.VeteranRewardTemplate{}, c).Model(&models.VeteranRewardTemplate{}).Create(&veteranRewardTemplate).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, veteranRewardTemplate)
}

// deleteVeteranRewardTemplate godoc
// @Id deleteVeteranRewardTemplate
// @Summary Deletes VeteranRewardTemplate
// @Accept json
// @Produce json
// @Tags VeteranRewardTemplate
// @Param id path int true "claimId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /veteran_reward_template/{id} [delete]
func (e *VeteranRewardTemplateController) deleteVeteranRewardTemplate(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	claimId, err := strconv.Atoi(c.Param("claimId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, claimId)
	keys = append(keys, "claim_id = ?")

	// key param [reward_slot] position [5] type [tinyint]
	if len(c.QueryParam("reward_slot")) > 0 {
		rewardSlotParam, err := strconv.Atoi(c.QueryParam("reward_slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [reward_slot] err [%s]", err.Error())})
		}

		params = append(params, rewardSlotParam)
		keys = append(keys, "reward_slot = ?")
	}

	// query builder
	var result models.VeteranRewardTemplate
	query := e.db.QueryContext(models.VeteranRewardTemplate{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.VeteranRewardTemplate{}, c).Model(&models.VeteranRewardTemplate{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getVeteranRewardTemplatesBulk godoc
// @Id getVeteranRewardTemplatesBulk
// @Summary Gets VeteranRewardTemplates in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags VeteranRewardTemplate
// @Success 200 {array} models.VeteranRewardTemplate
// @Failure 500 {string} string "Bad query request"
// @Router /veteran_reward_templates/bulk [post]
func (e *VeteranRewardTemplateController) getVeteranRewardTemplatesBulk(c echo.Context) error {
	var results []models.VeteranRewardTemplate

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

	err := e.db.QueryContext(models.VeteranRewardTemplate{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
