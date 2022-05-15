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

type BuyerController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewBuyerController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *BuyerController {
	return &BuyerController{
		db:	 db,
		logger: logger,
	}
}

func (e *BuyerController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "buyer/:charid", e.getBuyer, nil),
		routes.RegisterRoute(http.MethodGet, "buyers", e.listBuyers, nil),
		routes.RegisterRoute(http.MethodPut, "buyer", e.createBuyer, nil),
		routes.RegisterRoute(http.MethodDelete, "buyer/:charid", e.deleteBuyer, nil),
		routes.RegisterRoute(http.MethodPatch, "buyer/:charid", e.updateBuyer, nil),
		routes.RegisterRoute(http.MethodPost, "buyers/bulk", e.getBuyersBulk, nil),
	}
}

// listBuyers godoc
// @Id listBuyers
// @Summary Lists Buyers
// @Accept json
// @Produce json
// @Tags Buyer
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Buyer
// @Failure 500 {string} string "Bad query request"
// @Router /buyers [get]
func (e *BuyerController) listBuyers(c echo.Context) error {
	var results []models.Buyer
	err := e.db.QueryContext(models.Buyer{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBuyer godoc
// @Id getBuyer
// @Summary Gets Buyer
// @Accept json
// @Produce json
// @Tags Buyer
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Buyer
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /buyer/{id} [get]
func (e *BuyerController) getBuyer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [buyslot] position [2] type [int]
	if len(c.QueryParam("buyslot")) > 0 {
		buyslotParam, err := strconv.Atoi(c.QueryParam("buyslot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [buyslot] err [%s]", err.Error())})
		}

		params = append(params, buyslotParam)
		keys = append(keys, "buyslot = ?")
	}

	// query builder
	var result models.Buyer
	query := e.db.QueryContext(models.Buyer{}, c)
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

// updateBuyer godoc
// @Id updateBuyer
// @Summary Updates Buyer
// @Accept json
// @Produce json
// @Tags Buyer
// @Param id path int true "Id"
// @Param buyer body models.Buyer true "Buyer"
// @Success 200 {array} models.Buyer
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /buyer/{id} [patch]
func (e *BuyerController) updateBuyer(c echo.Context) error {
	request := new(models.Buyer)
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

	// key param [buyslot] position [2] type [int]
	if len(c.QueryParam("buyslot")) > 0 {
		buyslotParam, err := strconv.Atoi(c.QueryParam("buyslot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [buyslot] err [%s]", err.Error())})
		}

		params = append(params, buyslotParam)
		keys = append(keys, "buyslot = ?")
	}

	// query builder
	var result models.Buyer
	query := e.db.QueryContext(models.Buyer{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Buyer{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createBuyer godoc
// @Id createBuyer
// @Summary Creates Buyer
// @Accept json
// @Produce json
// @Param buyer body models.Buyer true "Buyer"
// @Tags Buyer
// @Success 200 {array} models.Buyer
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /buyer [put]
func (e *BuyerController) createBuyer(c echo.Context) error {
	buyer := new(models.Buyer)
	if err := c.Bind(buyer); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Buyer{}, c).Model(&models.Buyer{}).Create(&buyer).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, buyer)
}

// deleteBuyer godoc
// @Id deleteBuyer
// @Summary Deletes Buyer
// @Accept json
// @Produce json
// @Tags Buyer
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /buyer/{id} [delete]
func (e *BuyerController) deleteBuyer(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [buyslot] position [2] type [int]
	if len(c.QueryParam("buyslot")) > 0 {
		buyslotParam, err := strconv.Atoi(c.QueryParam("buyslot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [buyslot] err [%s]", err.Error())})
		}

		params = append(params, buyslotParam)
		keys = append(keys, "buyslot = ?")
	}

	// query builder
	var result models.Buyer
	query := e.db.QueryContext(models.Buyer{}, c)
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

// getBuyersBulk godoc
// @Id getBuyersBulk
// @Summary Gets Buyers in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Buyer
// @Success 200 {array} models.Buyer
// @Failure 500 {string} string "Bad query request"
// @Router /buyers/bulk [post]
func (e *BuyerController) getBuyersBulk(c echo.Context) error {
	var results []models.Buyer

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

	err := e.db.QueryContext(models.Buyer{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
