package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterPotionbeltController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterPotionbeltController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterPotionbeltController {
	return &CharacterPotionbeltController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterPotionbeltController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_potionbelt/:character_potionbelt", e.deleteCharacterPotionbelt, nil),
		routes.RegisterRoute(http.MethodGet, "character_potionbelt/:character_potionbelt", e.getCharacterPotionbelt, nil),
		routes.RegisterRoute(http.MethodGet, "character_potionbelts", e.listCharacterPotionbelts, nil),
		routes.RegisterRoute(http.MethodPost, "character_potionbelts/bulk", e.getCharacterPotionbeltsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_potionbelt/:character_potionbelt", e.updateCharacterPotionbelt, nil),
		routes.RegisterRoute(http.MethodPut, "character_potionbelt", e.createCharacterPotionbelt, nil),
	}
}

// listCharacterPotionbelts godoc
// @Id listCharacterPotionbelts
// @Summary Lists CharacterPotionbelts
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 500 {string} string "Bad query request"
// @Router /character_potionbelts [get]
func (e *CharacterPotionbeltController) listCharacterPotionbelts(c echo.Context) error {
	var results []models.CharacterPotionbelt
	err := e.db.QueryContext(models.CharacterPotionbelt{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterPotionbelt godoc
// @Id getCharacterPotionbelt
// @Summary Gets CharacterPotionbelt
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_potionbelt/{id} [get]
func (e *CharacterPotionbeltController) getCharacterPotionbelt(c echo.Context) error {
	characterPotionbeltId, err := strconv.Atoi(c.Param("character_potionbelt"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterPotionbelt
	err = e.db.QueryContext(models.CharacterPotionbelt{}, c).First(&result, characterPotionbeltId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateCharacterPotionbelt godoc
// @Id updateCharacterPotionbelt
// @Summary Updates CharacterPotionbelt
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param id path int true "Id"
// @Param character_potionbelt body models.CharacterPotionbelt true "CharacterPotionbelt"
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_potionbelt/{id} [patch]
func (e *CharacterPotionbeltController) updateCharacterPotionbelt(c echo.Context) error {
	characterPotionbelt := new(models.CharacterPotionbelt)
	if err := c.Bind(characterPotionbelt); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.CharacterPotionbelt{}
	err := e.db.Get(models.CharacterPotionbelt{}, c).Model(&models.CharacterPotionbelt{}).First(&entity, characterPotionbelt.ID).Error
	if err != nil || characterPotionbelt.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterPotionbelt{}, c).Model(&entity).Select("*").Updates(&characterPotionbelt).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterPotionbelt)
}

// createCharacterPotionbelt godoc
// @Id createCharacterPotionbelt
// @Summary Creates CharacterPotionbelt
// @Accept json
// @Produce json
// @Param character_potionbelt body models.CharacterPotionbelt true "CharacterPotionbelt"
// @Tags CharacterPotionbelt
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_potionbelt [put]
func (e *CharacterPotionbeltController) createCharacterPotionbelt(c echo.Context) error {
	characterPotionbelt := new(models.CharacterPotionbelt)
	if err := c.Bind(characterPotionbelt); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterPotionbelt{}, c).Model(&models.CharacterPotionbelt{}).Create(&characterPotionbelt).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, characterPotionbelt)
}

// deleteCharacterPotionbelt godoc
// @Id deleteCharacterPotionbelt
// @Summary Deletes CharacterPotionbelt
// @Accept json
// @Produce json
// @Tags CharacterPotionbelt
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_potionbelt/{id} [delete]
func (e *CharacterPotionbeltController) deleteCharacterPotionbelt(c echo.Context) error {
	characterPotionbeltId, err := strconv.Atoi(c.Param("character_potionbelt"))
	if err != nil {
		e.logger.Error(err)
	}

	characterPotionbelt := new(models.CharacterPotionbelt)
	err = e.db.Get(models.CharacterPotionbelt{}, c).Model(&models.CharacterPotionbelt{}).First(&characterPotionbelt, characterPotionbeltId).Error
	if err != nil || characterPotionbelt.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterPotionbelt{}, c).Model(&models.CharacterPotionbelt{}).Delete(&characterPotionbelt).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterPotionbeltsBulk godoc
// @Id getCharacterPotionbeltsBulk
// @Summary Gets CharacterPotionbelts in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterPotionbelt
// @Success 200 {array} models.CharacterPotionbelt
// @Failure 500 {string} string "Bad query request"
// @Router /character_potionbelts/bulk [post]
func (e *CharacterPotionbeltController) getCharacterPotionbeltsBulk(c echo.Context) error {
	var results []models.CharacterPotionbelt

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

	err := e.db.QueryContext(models.CharacterPotionbelt{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
