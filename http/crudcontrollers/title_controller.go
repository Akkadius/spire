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

type TitleController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTitleController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TitleController {
	return &TitleController {
		db:     db,
		logger: logger,
	}
}

func (e *TitleController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "title/:title", e.deleteTitle, nil),
		routes.RegisterRoute(http.MethodGet, "title/:title", e.getTitle, nil),
		routes.RegisterRoute(http.MethodGet, "titles", e.listTitles, nil),
		routes.RegisterRoute(http.MethodPatch, "title/:title", e.updateTitle, nil),
		routes.RegisterRoute(http.MethodPut, "title", e.createTitle, nil),
	}
}

// listTitles godoc
// @Id listTitles
// @Summary Lists Titles
// @Accept json
// @Produce json
// @Tags Title
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Title
// @Failure 500 {string} string "Bad query request"
// @Router /titles [get]
func (e *TitleController) listTitles(c echo.Context) error {
	var results []models.Title
	err := e.db.QueryContext(models.Title{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTitle godoc
// @Id getTitle
// @Summary Gets Title
// @Accept json
// @Produce json
// @Tags Title
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Title
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /title/{id} [get]
func (e *TitleController) getTitle(c echo.Context) error {
	titleId, err := strconv.Atoi(c.Param("title"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Title
	err = e.db.QueryContext(models.Title{}, c).First(&result, titleId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTitle godoc
// @Id updateTitle
// @Summary Updates Title
// @Accept json
// @Produce json
// @Tags Title
// @Param id path int true "Id"
// @Param title body models.Title true "Title"
// @Success 200 {array} models.Title
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /title/{id} [patch]
func (e *TitleController) updateTitle(c echo.Context) error {
	title := new(models.Title)
	if err := c.Bind(title); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Title{}, c).Model(&models.Title{}).First(&models.Title{}, title.ID).Error
	if err != nil || title.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Title{}, c).Model(&models.Title{}).Update(&title).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, title)
}

// createTitle godoc
// @Id createTitle
// @Summary Creates Title
// @Accept json
// @Produce json
// @Param title body models.Title true "Title"
// @Tags Title
// @Success 200 {array} models.Title
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /title [put]
func (e *TitleController) createTitle(c echo.Context) error {
	title := new(models.Title)
	if err := c.Bind(title); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.Title{}, c).Model(&models.Title{}).Create(&title).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, title)
}

// deleteTitle godoc
// @Id deleteTitle
// @Summary Deletes Title
// @Accept json
// @Produce json
// @Tags Title
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /title/{id} [delete]
func (e *TitleController) deleteTitle(c echo.Context) error {
	titleId, err := strconv.Atoi(c.Param("title"))
	if err != nil {
		e.logger.Error(err)
	}

	title := new(models.Title)
	err = e.db.Get(models.Title{}, c).Model(&models.Title{}).First(&title, titleId).Error
	if err != nil || title.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Title{}, c).Model(&models.Title{}).Delete(&title).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
