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

type FriendController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewFriendController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *FriendController {
	return &FriendController{
		db:	 db,
		logger: logger,
	}
}

func (e *FriendController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "friend/:charid", e.getFriend, nil),
		routes.RegisterRoute(http.MethodGet, "friends", e.listFriends, nil),
		routes.RegisterRoute(http.MethodPut, "friend", e.createFriend, nil),
		routes.RegisterRoute(http.MethodDelete, "friend/:charid", e.deleteFriend, nil),
		routes.RegisterRoute(http.MethodPatch, "friend/:charid", e.updateFriend, nil),
		routes.RegisterRoute(http.MethodPost, "friends/bulk", e.getFriendsBulk, nil),
	}
}

// listFriends godoc
// @Id listFriends
// @Summary Lists Friends
// @Accept json
// @Produce json
// @Tags Friend
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Friend
// @Failure 500 {string} string "Bad query request"
// @Router /friends [get]
func (e *FriendController) listFriends(c echo.Context) error {
	var results []models.Friend
	err := e.db.QueryContext(models.Friend{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getFriend godoc
// @Id getFriend
// @Summary Gets Friend
// @Accept json
// @Produce json
// @Tags Friend
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Friend
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /friend/{id} [get]
func (e *FriendController) getFriend(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Charid]"})
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [typeId] position [2] type [tinyint]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// key param [name] position [3] type [varchar]
	if len(c.QueryParam("name")) > 0 {
		nameParam, err := strconv.Atoi(c.QueryParam("name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [name] err [%s]", err.Error())})
		}

		params = append(params, nameParam)
		keys = append(keys, "name = ?")
	}

	// query builder
	var result models.Friend
	query := e.db.QueryContext(models.Friend{}, c)
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

// updateFriend godoc
// @Id updateFriend
// @Summary Updates Friend
// @Accept json
// @Produce json
// @Tags Friend
// @Param id path int true "Id"
// @Param friend body models.Friend true "Friend"
// @Success 200 {array} models.Friend
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /friend/{id} [patch]
func (e *FriendController) updateFriend(c echo.Context) error {
	request := new(models.Friend)
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

	// key param [typeId] position [2] type [tinyint]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// key param [name] position [3] type [varchar]
	if len(c.QueryParam("name")) > 0 {
		nameParam, err := strconv.Atoi(c.QueryParam("name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [name] err [%s]", err.Error())})
		}

		params = append(params, nameParam)
		keys = append(keys, "name = ?")
	}

	// query builder
	var result models.Friend
	query := e.db.QueryContext(models.Friend{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Friend{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createFriend godoc
// @Id createFriend
// @Summary Creates Friend
// @Accept json
// @Produce json
// @Param friend body models.Friend true "Friend"
// @Tags Friend
// @Success 200 {array} models.Friend
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /friend [put]
func (e *FriendController) createFriend(c echo.Context) error {
	friend := new(models.Friend)
	if err := c.Bind(friend); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Friend{}, c).Model(&models.Friend{}).Create(&friend).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, friend)
}

// deleteFriend godoc
// @Id deleteFriend
// @Summary Deletes Friend
// @Accept json
// @Produce json
// @Tags Friend
// @Param id path int true "charid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /friend/{id} [delete]
func (e *FriendController) deleteFriend(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charid, err := strconv.Atoi(c.Param("charid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charid)
	keys = append(keys, "charid = ?")

	// key param [typeId] position [2] type [tinyint]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// key param [name] position [3] type [varchar]
	if len(c.QueryParam("name")) > 0 {
		nameParam, err := strconv.Atoi(c.QueryParam("name"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [name] err [%s]", err.Error())})
		}

		params = append(params, nameParam)
		keys = append(keys, "name = ?")
	}

	// query builder
	var result models.Friend
	query := e.db.QueryContext(models.Friend{}, c)
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

// getFriendsBulk godoc
// @Id getFriendsBulk
// @Summary Gets Friends in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Friend
// @Success 200 {array} models.Friend
// @Failure 500 {string} string "Bad query request"
// @Router /friends/bulk [post]
func (e *FriendController) getFriendsBulk(c echo.Context) error {
	var results []models.Friend

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

	err := e.db.QueryContext(models.Friend{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
