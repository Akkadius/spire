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

type CharacterBindController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterBindController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterBindController {
	return &CharacterBindController {
		db:     db,
		logger: logger,
	}
}

func (e *CharacterBindController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_bind/:character_bind", e.deleteCharacterBind, nil),
		routes.RegisterRoute(http.MethodGet, "character_bind/:character_bind", e.getCharacterBind, nil),
		routes.RegisterRoute(http.MethodGet, "character_binds", e.listCharacterBinds, nil),
		routes.RegisterRoute(http.MethodPatch, "character_bind/:character_bind", e.updateCharacterBind, nil),
		routes.RegisterRoute(http.MethodPut, "character_bind", e.createCharacterBind, nil),
	}
}

// listCharacterBinds godoc
// @Id listCharacterBinds
// @Summary Lists CharacterBinds
// @Accept json
// @Produce json
// @Tags CharacterBind
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterBind
// @Failure 500 {string} string "Bad query request"
// @Router /character_binds [get]
func (e *CharacterBindController) listCharacterBinds(c echo.Context) error {
	var results []models.CharacterBind
	err := e.db.QueryContext(models.CharacterBind{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterBind godoc
// @Id getCharacterBind
// @Summary Gets CharacterBind
// @Accept json
// @Produce json
// @Tags CharacterBind
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterBind
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_bind/{id} [get]
func (e *CharacterBindController) getCharacterBind(c echo.Context) error {
	characterBindId, err := strconv.Atoi(c.Param("character_bind"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterBind
	err = e.db.QueryContext(models.CharacterBind{}, c).First(&result, characterBindId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterBind godoc
// @Id updateCharacterBind
// @Summary Updates CharacterBind
// @Accept json
// @Produce json
// @Tags CharacterBind
// @Param id path int true "Id"
// @Param character_bind body models.CharacterBind true "CharacterBind"
// @Success 200 {array} models.CharacterBind
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_bind/{id} [patch]
func (e *CharacterBindController) updateCharacterBind(c echo.Context) error {
	characterBind := new(models.CharacterBind)
	if err := c.Bind(characterBind); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.CharacterBind{}, c).Model(&models.CharacterBind{}).First(&models.CharacterBind{}, characterBind.ID).Error
	if err != nil || characterBind.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterBind{}, c).Model(&models.CharacterBind{}).Update(&characterBind).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterBind)
}

// createCharacterBind godoc
// @Id createCharacterBind
// @Summary Creates CharacterBind
// @Accept json
// @Produce json
// @Param character_bind body models.CharacterBind true "CharacterBind"
// @Tags CharacterBind
// @Success 200 {array} models.CharacterBind
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_bind [put]
func (e *CharacterBindController) createCharacterBind(c echo.Context) error {
	characterBind := new(models.CharacterBind)
	if err := c.Bind(characterBind); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)})
	}

	err := e.db.Get(models.CharacterBind{}, c).Model(&models.CharacterBind{}).Create(&characterBind).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterBind)
}

// deleteCharacterBind godoc
// @Id deleteCharacterBind
// @Summary Deletes CharacterBind
// @Accept json
// @Produce json
// @Tags CharacterBind
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_bind/{id} [delete]
func (e *CharacterBindController) deleteCharacterBind(c echo.Context) error {
	characterBindId, err := strconv.Atoi(c.Param("character_bind"))
	if err != nil {
		e.logger.Error(err)
	}

	characterBind := new(models.CharacterBind)
	err = e.db.Get(models.CharacterBind{}, c).Model(&models.CharacterBind{}).First(&characterBind, characterBindId).Error
	if err != nil || characterBind.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterBind{}, c).Model(&models.CharacterBind{}).Delete(&characterBind).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}
