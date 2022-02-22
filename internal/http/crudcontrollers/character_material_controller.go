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

type CharacterMaterialController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterMaterialController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterMaterialController {
	return &CharacterMaterialController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterMaterialController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_material/:id", e.getCharacterMaterial, nil),
		routes.RegisterRoute(http.MethodGet, "character_materials", e.listCharacterMaterials, nil),
		routes.RegisterRoute(http.MethodPut, "character_material", e.createCharacterMaterial, nil),
		routes.RegisterRoute(http.MethodDelete, "character_material/:id", e.deleteCharacterMaterial, nil),
		routes.RegisterRoute(http.MethodPatch, "character_material/:id", e.updateCharacterMaterial, nil),
		routes.RegisterRoute(http.MethodPost, "character_materials/bulk", e.getCharacterMaterialsBulk, nil),
	}
}

// listCharacterMaterials godoc
// @Id listCharacterMaterials
// @Summary Lists CharacterMaterials
// @Accept json
// @Produce json
// @Tags CharacterMaterial
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterMaterial
// @Failure 500 {string} string "Bad query request"
// @Router /character_materials [get]
func (e *CharacterMaterialController) listCharacterMaterials(c echo.Context) error {
	var results []models.CharacterMaterial
	err := e.db.QueryContext(models.CharacterMaterial{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterMaterial godoc
// @Id getCharacterMaterial
// @Summary Gets CharacterMaterial
// @Accept json
// @Produce json
// @Tags CharacterMaterial
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterMaterial
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_material/{id} [get]
func (e *CharacterMaterialController) getCharacterMaterial(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterMaterial
	query := e.db.QueryContext(models.CharacterMaterial{}, c)
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

// updateCharacterMaterial godoc
// @Id updateCharacterMaterial
// @Summary Updates CharacterMaterial
// @Accept json
// @Produce json
// @Tags CharacterMaterial
// @Param id path int true "Id"
// @Param character_material body models.CharacterMaterial true "CharacterMaterial"
// @Success 200 {array} models.CharacterMaterial
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_material/{id} [patch]
func (e *CharacterMaterialController) updateCharacterMaterial(c echo.Context) error {
	request := new(models.CharacterMaterial)
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

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterMaterial
	query := e.db.QueryContext(models.CharacterMaterial{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterMaterial godoc
// @Id createCharacterMaterial
// @Summary Creates CharacterMaterial
// @Accept json
// @Produce json
// @Param character_material body models.CharacterMaterial true "CharacterMaterial"
// @Tags CharacterMaterial
// @Success 200 {array} models.CharacterMaterial
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_material [put]
func (e *CharacterMaterialController) createCharacterMaterial(c echo.Context) error {
	characterMaterial := new(models.CharacterMaterial)
	if err := c.Bind(characterMaterial); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterMaterial{}, c).Model(&models.CharacterMaterial{}).Create(&characterMaterial).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterMaterial)
}

// deleteCharacterMaterial godoc
// @Id deleteCharacterMaterial
// @Summary Deletes CharacterMaterial
// @Accept json
// @Produce json
// @Tags CharacterMaterial
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_material/{id} [delete]
func (e *CharacterMaterialController) deleteCharacterMaterial(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [slot] position [2] type [tinyint]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.CharacterMaterial
	query := e.db.QueryContext(models.CharacterMaterial{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.CharacterMaterial{}, c).Model(&models.CharacterMaterial{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getCharacterMaterialsBulk godoc
// @Id getCharacterMaterialsBulk
// @Summary Gets CharacterMaterials in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterMaterial
// @Success 200 {array} models.CharacterMaterial
// @Failure 500 {string} string "Bad query request"
// @Router /character_materials/bulk [post]
func (e *CharacterMaterialController) getCharacterMaterialsBulk(c echo.Context) error {
	var results []models.CharacterMaterial

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

	err := e.db.QueryContext(models.CharacterMaterial{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
