package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ExpeditionMemberController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewExpeditionMemberController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ExpeditionMemberController {
	return &ExpeditionMemberController {
		db:     db,
		logger: logger,
	}
}

func (e *ExpeditionMemberController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "expedition_member/:expedition_member", e.deleteExpeditionMember, nil),
		routes.RegisterRoute(http.MethodGet, "expedition_member/:expedition_member", e.getExpeditionMember, nil),
		routes.RegisterRoute(http.MethodGet, "expedition_members", e.listExpeditionMembers, nil),
		routes.RegisterRoute(http.MethodPatch, "expedition_member/:expedition_member", e.updateExpeditionMember, nil),
		routes.RegisterRoute(http.MethodPut, "expedition_member", e.createExpeditionMember, nil),
	}
}

// listExpeditionMembers godoc
// @Id listExpeditionMembers
// @Summary Lists ExpeditionMembers
// @Accept json
// @Produce json
// @Tags ExpeditionMember
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ExpeditionMember
// @Failure 500 {string} string "Bad query request"
// @Router /expedition_members [get]
func (e *ExpeditionMemberController) listExpeditionMembers(c echo.Context) error {
	var results []models.ExpeditionMember
	err := e.db.QueryContext(models.ExpeditionMember{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getExpeditionMember godoc
// @Id getExpeditionMember
// @Summary Gets ExpeditionMember
// @Accept json
// @Produce json
// @Tags ExpeditionMember
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.ExpeditionMember
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /expedition_member/{id} [get]
func (e *ExpeditionMemberController) getExpeditionMember(c echo.Context) error {
	expeditionMemberId, err := strconv.Atoi(c.Param("expedition_member"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.ExpeditionMember
	err = e.db.QueryContext(models.ExpeditionMember{}, c).First(&result, expeditionMemberId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateExpeditionMember godoc
// @Id updateExpeditionMember
// @Summary Updates ExpeditionMember
// @Accept json
// @Produce json
// @Tags ExpeditionMember
// @Param id path int true "Id"
// @Param expedition_member body models.ExpeditionMember true "ExpeditionMember"
// @Success 200 {array} models.ExpeditionMember
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /expedition_member/{id} [patch]
func (e *ExpeditionMemberController) updateExpeditionMember(c echo.Context) error {
	expeditionMember := new(models.ExpeditionMember)
	if err := c.Bind(expeditionMember); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.ExpeditionMember{}, c).Model(&models.ExpeditionMember{}).First(&models.ExpeditionMember{}, expeditionMember.ID).Error
	if err != nil || expeditionMember.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ExpeditionMember{}, c).Model(&models.ExpeditionMember{}).Update(&expeditionMember).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, expeditionMember)
}

// createExpeditionMember godoc
// @Id createExpeditionMember
// @Summary Creates ExpeditionMember
// @Accept json
// @Produce json
// @Param expedition_member body models.ExpeditionMember true "ExpeditionMember"
// @Tags ExpeditionMember
// @Success 200 {array} models.ExpeditionMember
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /expedition_member [put]
func (e *ExpeditionMemberController) createExpeditionMember(c echo.Context) error {
	expeditionMember := new(models.ExpeditionMember)
	if err := c.Bind(expeditionMember); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.ExpeditionMember{}, c).Model(&models.ExpeditionMember{}).Create(&expeditionMember).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, expeditionMember)
}

// deleteExpeditionMember godoc
// @Id deleteExpeditionMember
// @Summary Deletes ExpeditionMember
// @Accept json
// @Produce json
// @Tags ExpeditionMember
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /expedition_member/{id} [delete]
func (e *ExpeditionMemberController) deleteExpeditionMember(c echo.Context) error {
	expeditionMemberId, err := strconv.Atoi(c.Param("expedition_member"))
	if err != nil {
		e.logger.Error(err)
	}

	expeditionMember := new(models.ExpeditionMember)
	err = e.db.Get(models.ExpeditionMember{}, c).Model(&models.ExpeditionMember{}).First(&expeditionMember, expeditionMemberId).Error
	if err != nil || expeditionMember.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.ExpeditionMember{}, c).Model(&models.ExpeditionMember{}).Delete(&expeditionMember).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
