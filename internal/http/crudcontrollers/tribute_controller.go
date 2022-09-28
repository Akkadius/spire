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

type TributeController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTributeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TributeController {
	return &TributeController{
		db:	 db,
		logger: logger,
	}
}

func (e *TributeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "tribute/:id", e.getTribute, nil),
		routes.RegisterRoute(http.MethodGet, "tributes", e.listTributes, nil),
		routes.RegisterRoute(http.MethodPut, "tribute", e.createTribute, nil),
		routes.RegisterRoute(http.MethodDelete, "tribute/:id", e.deleteTribute, nil),
		routes.RegisterRoute(http.MethodPatch, "tribute/:id", e.updateTribute, nil),
		routes.RegisterRoute(http.MethodPost, "tributes/bulk", e.getTributesBulk, nil),
	}
}

// listTributes godoc
// @Id listTributes
// @Summary Lists Tributes
// @Accept json
// @Produce json
// @Tags Tribute
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Tribute
// @Failure 500 {string} string "Bad query request"
// @Router /tributes [get]
func (e *TributeController) listTributes(c echo.Context) error {
	var results []models.Tribute
	err := e.db.QueryContext(models.Tribute{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTribute godoc
// @Id getTribute
// @Summary Gets Tribute
// @Accept json
// @Produce json
// @Tags Tribute
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Tribute
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tribute/{id} [get]
func (e *TributeController) getTribute(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [isguild] position [5] type [tinyint]
	if len(c.QueryParam("isguild")) > 0 {
		isguildParam, err := strconv.Atoi(c.QueryParam("isguild"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [isguild] err [%s]", err.Error())})
		}

		params = append(params, isguildParam)
		keys = append(keys, "isguild = ?")
	}

	// query builder
	var result models.Tribute
	query := e.db.QueryContext(models.Tribute{}, c)
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

// updateTribute godoc
// @Id updateTribute
// @Summary Updates Tribute
// @Accept json
// @Produce json
// @Tags Tribute
// @Param id path int true "Id"
// @Param tribute body models.Tribute true "Tribute"
// @Success 200 {array} models.Tribute
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /tribute/{id} [patch]
func (e *TributeController) updateTribute(c echo.Context) error {
	request := new(models.Tribute)
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

	// key param [isguild] position [5] type [tinyint]
	if len(c.QueryParam("isguild")) > 0 {
		isguildParam, err := strconv.Atoi(c.QueryParam("isguild"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [isguild] err [%s]", err.Error())})
		}

		params = append(params, isguildParam)
		keys = append(keys, "isguild = ?")
	}

	// query builder
	var result models.Tribute
	query := e.db.QueryContext(models.Tribute{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Tribute{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createTribute godoc
// @Id createTribute
// @Summary Creates Tribute
// @Accept json
// @Produce json
// @Param tribute body models.Tribute true "Tribute"
// @Tags Tribute
// @Success 200 {array} models.Tribute
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /tribute [put]
func (e *TributeController) createTribute(c echo.Context) error {
	tribute := new(models.Tribute)
	if err := c.Bind(tribute); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Tribute{}, c).Model(&models.Tribute{}).Create(&tribute).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, tribute)
}

// deleteTribute godoc
// @Id deleteTribute
// @Summary Deletes Tribute
// @Accept json
// @Produce json
// @Tags Tribute
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tribute/{id} [delete]
func (e *TributeController) deleteTribute(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [isguild] position [5] type [tinyint]
	if len(c.QueryParam("isguild")) > 0 {
		isguildParam, err := strconv.Atoi(c.QueryParam("isguild"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [isguild] err [%s]", err.Error())})
		}

		params = append(params, isguildParam)
		keys = append(keys, "isguild = ?")
	}

	// query builder
	var result models.Tribute
	query := e.db.QueryContext(models.Tribute{}, c)
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

// getTributesBulk godoc
// @Id getTributesBulk
// @Summary Gets Tributes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Tribute
// @Success 200 {array} models.Tribute
// @Failure 500 {string} string "Bad query request"
// @Router /tributes/bulk [post]
func (e *TributeController) getTributesBulk(c echo.Context) error {
	var results []models.Tribute

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

	err := e.db.QueryContext(models.Tribute{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
