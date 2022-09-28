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

type AccountRewardController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAccountRewardController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AccountRewardController {
	return &AccountRewardController{
		db:	 db,
		logger: logger,
	}
}

func (e *AccountRewardController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "account_reward/:accountId", e.getAccountReward, nil),
		routes.RegisterRoute(http.MethodGet, "account_rewards", e.listAccountRewards, nil),
		routes.RegisterRoute(http.MethodPut, "account_reward", e.createAccountReward, nil),
		routes.RegisterRoute(http.MethodDelete, "account_reward/:accountId", e.deleteAccountReward, nil),
		routes.RegisterRoute(http.MethodPatch, "account_reward/:accountId", e.updateAccountReward, nil),
		routes.RegisterRoute(http.MethodPost, "account_rewards/bulk", e.getAccountRewardsBulk, nil),
	}
}

// listAccountRewards godoc
// @Id listAccountRewards
// @Summary Lists AccountRewards
// @Accept json
// @Produce json
// @Tags AccountReward
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AccountReward
// @Failure 500 {string} string "Bad query request"
// @Router /account_rewards [get]
func (e *AccountRewardController) listAccountRewards(c echo.Context) error {
	var results []models.AccountReward
	err := e.db.QueryContext(models.AccountReward{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAccountReward godoc
// @Id getAccountReward
// @Summary Gets AccountReward
// @Accept json
// @Produce json
// @Tags AccountReward
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.AccountReward
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /account_reward/{id} [get]
func (e *AccountRewardController) getAccountReward(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [AccountId]"})
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [reward_id] position [2] type [int]
	if len(c.QueryParam("reward_id")) > 0 {
		rewardIdParam, err := strconv.Atoi(c.QueryParam("reward_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [reward_id] err [%s]", err.Error())})
		}

		params = append(params, rewardIdParam)
		keys = append(keys, "reward_id = ?")
	}

	// query builder
	var result models.AccountReward
	query := e.db.QueryContext(models.AccountReward{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.AccountId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAccountReward godoc
// @Id updateAccountReward
// @Summary Updates AccountReward
// @Accept json
// @Produce json
// @Tags AccountReward
// @Param id path int true "Id"
// @Param account_reward body models.AccountReward true "AccountReward"
// @Success 200 {array} models.AccountReward
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /account_reward/{id} [patch]
func (e *AccountRewardController) updateAccountReward(c echo.Context) error {
	request := new(models.AccountReward)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [AccountId]"})
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [reward_id] position [2] type [int]
	if len(c.QueryParam("reward_id")) > 0 {
		rewardIdParam, err := strconv.Atoi(c.QueryParam("reward_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [reward_id] err [%s]", err.Error())})
		}

		params = append(params, rewardIdParam)
		keys = append(keys, "reward_id = ?")
	}

	// query builder
	var result models.AccountReward
	query := e.db.QueryContext(models.AccountReward{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.AccountReward{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAccountReward godoc
// @Id createAccountReward
// @Summary Creates AccountReward
// @Accept json
// @Produce json
// @Param account_reward body models.AccountReward true "AccountReward"
// @Tags AccountReward
// @Success 200 {array} models.AccountReward
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /account_reward [put]
func (e *AccountRewardController) createAccountReward(c echo.Context) error {
	accountReward := new(models.AccountReward)
	if err := c.Bind(accountReward); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.AccountReward{}, c).Model(&models.AccountReward{}).Create(&accountReward).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, accountReward)
}

// deleteAccountReward godoc
// @Id deleteAccountReward
// @Summary Deletes AccountReward
// @Accept json
// @Produce json
// @Tags AccountReward
// @Param id path int true "accountId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /account_reward/{id} [delete]
func (e *AccountRewardController) deleteAccountReward(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, accountId)
	keys = append(keys, "account_id = ?")

	// key param [reward_id] position [2] type [int]
	if len(c.QueryParam("reward_id")) > 0 {
		rewardIdParam, err := strconv.Atoi(c.QueryParam("reward_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [reward_id] err [%s]", err.Error())})
		}

		params = append(params, rewardIdParam)
		keys = append(keys, "reward_id = ?")
	}

	// query builder
	var result models.AccountReward
	query := e.db.QueryContext(models.AccountReward{}, c)
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

// getAccountRewardsBulk godoc
// @Id getAccountRewardsBulk
// @Summary Gets AccountRewards in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags AccountReward
// @Success 200 {array} models.AccountReward
// @Failure 500 {string} string "Bad query request"
// @Router /account_rewards/bulk [post]
func (e *AccountRewardController) getAccountRewardsBulk(c echo.Context) error {
	var results []models.AccountReward

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

	err := e.db.QueryContext(models.AccountReward{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
