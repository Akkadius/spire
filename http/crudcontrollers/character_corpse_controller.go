package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterCorpseController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterCorpseController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterCorpseController {
	return &CharacterCorpseController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterCorpseController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_corpse/:character_corpse", e.deleteCharacterCorpse, nil),
		routes.RegisterRoute(http.MethodGet, "character_corpse/:character_corpse", e.getCharacterCorpse, nil),
		routes.RegisterRoute(http.MethodGet, "character_corpses", e.listCharacterCorpses, nil),
		routes.RegisterRoute(http.MethodPost, "character_corpses/bulk", e.getCharacterCorpsesBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_corpse/:character_corpse", e.updateCharacterCorpse, nil),
		routes.RegisterRoute(http.MethodPut, "character_corpse", e.createCharacterCorpse, nil),
	}
}

// listCharacterCorpses godoc
// @Id listCharacterCorpses
// @Summary Lists CharacterCorpses
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterCorpse
// @Failure 500 {string} string "Bad query request"
// @Router /character_corpses [get]
func (e *CharacterCorpseController) listCharacterCorpses(c echo.Context) error {
	var results []models.CharacterCorpse
	err := e.db.QueryContext(models.CharacterCorpse{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterCorpse godoc
// @Id getCharacterCorpse
// @Summary Gets CharacterCorpse
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterCorpse
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_corpse/{id} [get]
func (e *CharacterCorpseController) getCharacterCorpse(c echo.Context) error {
	characterCorpseId, err := strconv.Atoi(c.Param("character_corpse"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterCorpse
	err = e.db.QueryContext(models.CharacterCorpse{}, c).First(&result, characterCorpseId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterCorpse godoc
// @Id updateCharacterCorpse
// @Summary Updates CharacterCorpse
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param id path int true "Id"
// @Param character_corpse body models.CharacterCorpse true "CharacterCorpse"
// @Success 200 {array} models.CharacterCorpse
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_corpse/{id} [patch]
func (e *CharacterCorpseController) updateCharacterCorpse(c echo.Context) error {
	characterCorpse := new(models.CharacterCorpse)
	if err := c.Bind(characterCorpse); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.CharacterCorpse{}
	err := e.db.Get(models.CharacterCorpse{}, c).Model(&models.CharacterCorpse{}).First(&entity, characterCorpse.ID).Error
	if err != nil || characterCorpse.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterCorpse{}, c).Model(&entity).Updates(&characterCorpse).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterCorpse)
}

// createCharacterCorpse godoc
// @Id createCharacterCorpse
// @Summary Creates CharacterCorpse
// @Accept json
// @Produce json
// @Param character_corpse body models.CharacterCorpse true "CharacterCorpse"
// @Tags CharacterCorpse
// @Success 200 {array} models.CharacterCorpse
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_corpse [put]
func (e *CharacterCorpseController) createCharacterCorpse(c echo.Context) error {
	characterCorpse := new(models.CharacterCorpse)
	if err := c.Bind(characterCorpse); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterCorpse{}, c).Model(&models.CharacterCorpse{}).Create(&characterCorpse).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, characterCorpse)
}

// deleteCharacterCorpse godoc
// @Id deleteCharacterCorpse
// @Summary Deletes CharacterCorpse
// @Accept json
// @Produce json
// @Tags CharacterCorpse
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_corpse/{id} [delete]
func (e *CharacterCorpseController) deleteCharacterCorpse(c echo.Context) error {
	characterCorpseId, err := strconv.Atoi(c.Param("character_corpse"))
	if err != nil {
		e.logger.Error(err)
	}

	characterCorpse := new(models.CharacterCorpse)
	err = e.db.Get(models.CharacterCorpse{}, c).Model(&models.CharacterCorpse{}).First(&characterCorpse, characterCorpseId).Error
	if err != nil || characterCorpse.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterCorpse{}, c).Model(&models.CharacterCorpse{}).Delete(&characterCorpse).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterCorpsesBulk godoc
// @Id getCharacterCorpsesBulk
// @Summary Gets CharacterCorpses in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterCorpse
// @Success 200 {array} models.CharacterCorpse
// @Failure 500 {string} string "Bad query request"
// @Router /character_corpses/bulk [post]
func (e *CharacterCorpseController) getCharacterCorpsesBulk(c echo.Context) error {
	var results []models.CharacterCorpse

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.CharacterCorpse{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
