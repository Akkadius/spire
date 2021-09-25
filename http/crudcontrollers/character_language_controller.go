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

type CharacterLanguageController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterLanguageController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterLanguageController {
	return &CharacterLanguageController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterLanguageController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_language/:character_language", e.deleteCharacterLanguage, nil),
		routes.RegisterRoute(http.MethodGet, "character_language/:character_language", e.getCharacterLanguage, nil),
		routes.RegisterRoute(http.MethodGet, "character_languages", e.listCharacterLanguages, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getCharacterLanguagesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_language/:character_language", e.updateCharacterLanguage, nil),
		routes.RegisterRoute(http.MethodPut, "character_language", e.createCharacterLanguage, nil),
	}
}

// listCharacterLanguages godoc
// @Id listCharacterLanguages
// @Summary Lists CharacterLanguages
// @Accept json
// @Produce json
// @Tags CharacterLanguage
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterLanguage
// @Failure 500 {string} string "Bad query request"
// @Router /character_languages [get]
func (e *CharacterLanguageController) listCharacterLanguages(c echo.Context) error {
	var results []models.CharacterLanguage
	err := e.db.QueryContext(models.CharacterLanguage{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterLanguage godoc
// @Id getCharacterLanguage
// @Summary Gets CharacterLanguage
// @Accept json
// @Produce json
// @Tags CharacterLanguage
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterLanguage
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_language/{id} [get]
func (e *CharacterLanguageController) getCharacterLanguage(c echo.Context) error {
	characterLanguageId, err := strconv.Atoi(c.Param("character_language"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterLanguage
	err = e.db.QueryContext(models.CharacterLanguage{}, c).First(&result, characterLanguageId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterLanguage godoc
// @Id updateCharacterLanguage
// @Summary Updates CharacterLanguage
// @Accept json
// @Produce json
// @Tags CharacterLanguage
// @Param id path int true "Id"
// @Param character_language body models.CharacterLanguage true "CharacterLanguage"
// @Success 200 {array} models.CharacterLanguage
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_language/{id} [patch]
func (e *CharacterLanguageController) updateCharacterLanguage(c echo.Context) error {
	characterLanguage := new(models.CharacterLanguage)
	if err := c.Bind(characterLanguage); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterLanguage{}, c).Model(&models.CharacterLanguage{}).First(&models.CharacterLanguage{}, characterLanguage.ID).Error
	if err != nil || characterLanguage.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterLanguage{}, c).Model(&models.CharacterLanguage{}).Updates(&characterLanguage).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterLanguage)
}

// createCharacterLanguage godoc
// @Id createCharacterLanguage
// @Summary Creates CharacterLanguage
// @Accept json
// @Produce json
// @Param character_language body models.CharacterLanguage true "CharacterLanguage"
// @Tags CharacterLanguage
// @Success 200 {array} models.CharacterLanguage
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_language [put]
func (e *CharacterLanguageController) createCharacterLanguage(c echo.Context) error {
	characterLanguage := new(models.CharacterLanguage)
	if err := c.Bind(characterLanguage); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterLanguage{}, c).Model(&models.CharacterLanguage{}).Create(&characterLanguage).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, characterLanguage)
}

// deleteCharacterLanguage godoc
// @Id deleteCharacterLanguage
// @Summary Deletes CharacterLanguage
// @Accept json
// @Produce json
// @Tags CharacterLanguage
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_language/{id} [delete]
func (e *CharacterLanguageController) deleteCharacterLanguage(c echo.Context) error {
	characterLanguageId, err := strconv.Atoi(c.Param("character_language"))
	if err != nil {
		e.logger.Error(err)
	}

	characterLanguage := new(models.CharacterLanguage)
	err = e.db.Get(models.CharacterLanguage{}, c).Model(&models.CharacterLanguage{}).First(&characterLanguage, characterLanguageId).Error
	if err != nil || characterLanguage.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterLanguage{}, c).Model(&models.CharacterLanguage{}).Delete(&characterLanguage).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterLanguagesBulk godoc
// @Id getCharacterLanguagesBulk
// @Summary Gets CharacterLanguages in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterLanguage
// @Success 200 {array} models.CharacterLanguage
// @Failure 500 {string} string "Bad query request"
// @Router /character_languages/bulk [post]
func (e *CharacterLanguageController) getCharacterLanguagesBulk(c echo.Context) error {
	var results []models.CharacterLanguage

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.CharacterLanguage{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
