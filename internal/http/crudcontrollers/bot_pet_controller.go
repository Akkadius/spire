package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type BotPetController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewBotPetController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *BotPetController {
	return &BotPetController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *BotPetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bot_pet/:petsIndex", e.getBotPet, nil),
		routes.RegisterRoute(http.MethodGet, "bot_pets", e.listBotPets, nil),
		routes.RegisterRoute(http.MethodGet, "bot_pets/count", e.getBotPetsCount, nil),
		routes.RegisterRoute(http.MethodPut, "bot_pet", e.createBotPet, nil),
		routes.RegisterRoute(http.MethodDelete, "bot_pet/:petsIndex", e.deleteBotPet, nil),
		routes.RegisterRoute(http.MethodPatch, "bot_pet/:petsIndex", e.updateBotPet, nil),
		routes.RegisterRoute(http.MethodPost, "bot_pets/bulk", e.getBotPetsBulk, nil),
	}
}

// listBotPets godoc
// @Id listBotPets
// @Summary Lists BotPets
// @Accept json
// @Produce json
// @Tags BotPet
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotPet
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pets [get]
func (e *BotPetController) listBotPets(c echo.Context) error {
	var results []models.BotPet
	err := e.db.QueryContext(models.BotPet{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotPet godoc
// @Id getBotPet
// @Summary Gets BotPet
// @Accept json
// @Produce json
// @Tags BotPet
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotPet
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pet/{id} [get]
func (e *BotPetController) getBotPet(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	petsIndex, err := strconv.Atoi(c.Param("petsIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PetsIndex]"})
	}
	params = append(params, petsIndex)
	keys = append(keys, "pets_index = ?")

	// query builder
	var result models.BotPet
	query := e.db.QueryContext(models.BotPet{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.PetsIndex == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBotPet godoc
// @Id updateBotPet
// @Summary Updates BotPet
// @Accept json
// @Produce json
// @Tags BotPet
// @Param id path int true "Id"
// @Param bot_pet body models.BotPet true "BotPet"
// @Success 200 {array} models.BotPet
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bot_pet/{id} [patch]
func (e *BotPetController) updateBotPet(c echo.Context) error {
	request := new(models.BotPet)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	petsIndex, err := strconv.Atoi(c.Param("petsIndex"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [PetsIndex]"})
	}
	params = append(params, petsIndex)
	keys = append(keys, "pets_index = ?")

	// query builder
	var result models.BotPet
	query := e.db.QueryContext(models.BotPet{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [BotPet] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createBotPet godoc
// @Id createBotPet
// @Summary Creates BotPet
// @Accept json
// @Produce json
// @Param bot_pet body models.BotPet true "BotPet"
// @Tags BotPet
// @Success 200 {array} models.BotPet
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bot_pet [put]
func (e *BotPetController) createBotPet(c echo.Context) error {
	botPet := new(models.BotPet)
	if err := c.Bind(botPet); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.BotPet{}, c).Model(&models.BotPet{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&botPet).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.BotPet{}, botPet)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [BotPet] [%v] data [%v]", botPet.PetsIndex, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, botPet)
}

// deleteBotPet godoc
// @Id deleteBotPet
// @Summary Deletes BotPet
// @Accept json
// @Produce json
// @Tags BotPet
// @Param id path int true "petsIndex"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bot_pet/{id} [delete]
func (e *BotPetController) deleteBotPet(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	petsIndex, err := strconv.Atoi(c.Param("petsIndex"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, petsIndex)
	keys = append(keys, "pets_index = ?")

	// query builder
	var result models.BotPet
	query := e.db.QueryContext(models.BotPet{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [BotPet] [%v] keys [%v]", result.PetsIndex, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBotPetsBulk godoc
// @Id getBotPetsBulk
// @Summary Gets BotPets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BotPet
// @Success 200 {array} models.BotPet
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pets/bulk [post]
func (e *BotPetController) getBotPetsBulk(c echo.Context) error {
	var results []models.BotPet

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

	err := e.db.QueryContext(models.BotPet{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getBotPetsCount godoc
// @Id getBotPetsCount
// @Summary Counts BotPets
// @Accept json
// @Produce json
// @Tags BotPet
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BotPet
// @Failure 500 {string} string "Bad query request"
// @Router /bot_pets/count [get]
func (e *BotPetController) getBotPetsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.BotPet{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}