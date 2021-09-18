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

type CharacterDisciplineController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterDisciplineController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterDisciplineController {
	return &CharacterDisciplineController {
		db:     db,
		logger: logger,
	}
}

func (e *CharacterDisciplineController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_discipline/:character_discipline", e.deleteCharacterDiscipline, nil),
		routes.RegisterRoute(http.MethodGet, "character_discipline/:character_discipline", e.getCharacterDiscipline, nil),
		routes.RegisterRoute(http.MethodGet, "character_disciplines", e.listCharacterDisciplines, nil),
		routes.RegisterRoute(http.MethodPatch, "character_discipline/:character_discipline", e.updateCharacterDiscipline, nil),
		routes.RegisterRoute(http.MethodPut, "character_discipline", e.createCharacterDiscipline, nil),
	}
}

// listCharacterDisciplines godoc
// @Id listCharacterDisciplines
// @Summary Lists CharacterDisciplines
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterDiscipline
// @Failure 500 {string} string "Bad query request"
// @Router /character_disciplines [get]
func (e *CharacterDisciplineController) listCharacterDisciplines(c echo.Context) error {
	var results []models.CharacterDiscipline
	err := e.db.QueryContext(models.CharacterDiscipline{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterDiscipline godoc
// @Id getCharacterDiscipline
// @Summary Gets CharacterDiscipline
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterDiscipline
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_discipline/{id} [get]
func (e *CharacterDisciplineController) getCharacterDiscipline(c echo.Context) error {
	characterDisciplineId, err := strconv.Atoi(c.Param("character_discipline"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterDiscipline
	err = e.db.QueryContext(models.CharacterDiscipline{}, c).First(&result, characterDisciplineId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterDiscipline godoc
// @Id updateCharacterDiscipline
// @Summary Updates CharacterDiscipline
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param id path int true "Id"
// @Param character_discipline body models.CharacterDiscipline true "CharacterDiscipline"
// @Success 200 {array} models.CharacterDiscipline
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_discipline/{id} [patch]
func (e *CharacterDisciplineController) updateCharacterDiscipline(c echo.Context) error {
	characterDiscipline := new(models.CharacterDiscipline)
	if err := c.Bind(characterDiscipline); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.CharacterDiscipline{}, c).Model(&models.CharacterDiscipline{}).First(&models.CharacterDiscipline{}, characterDiscipline.ID).Error
	if err != nil || characterDiscipline.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterDiscipline{}, c).Model(&models.CharacterDiscipline{}).Updates(&characterDiscipline).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterDiscipline)
}

// createCharacterDiscipline godoc
// @Id createCharacterDiscipline
// @Summary Creates CharacterDiscipline
// @Accept json
// @Produce json
// @Param character_discipline body models.CharacterDiscipline true "CharacterDiscipline"
// @Tags CharacterDiscipline
// @Success 200 {array} models.CharacterDiscipline
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_discipline [put]
func (e *CharacterDisciplineController) createCharacterDiscipline(c echo.Context) error {
	characterDiscipline := new(models.CharacterDiscipline)
	if err := c.Bind(characterDiscipline); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.CharacterDiscipline{}, c).Model(&models.CharacterDiscipline{}).Create(&characterDiscipline).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterDiscipline)
}

// deleteCharacterDiscipline godoc
// @Id deleteCharacterDiscipline
// @Summary Deletes CharacterDiscipline
// @Accept json
// @Produce json
// @Tags CharacterDiscipline
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_discipline/{id} [delete]
func (e *CharacterDisciplineController) deleteCharacterDiscipline(c echo.Context) error {
	characterDisciplineId, err := strconv.Atoi(c.Param("character_discipline"))
	if err != nil {
		e.logger.Error(err)
	}

	characterDiscipline := new(models.CharacterDiscipline)
	err = e.db.Get(models.CharacterDiscipline{}, c).Model(&models.CharacterDiscipline{}).First(&characterDiscipline, characterDisciplineId).Error
	if err != nil || characterDiscipline.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterDiscipline{}, c).Model(&models.CharacterDiscipline{}).Delete(&characterDiscipline).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
