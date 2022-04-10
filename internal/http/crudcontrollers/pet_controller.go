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

type PetController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewPetController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *PetController {
	return &PetController{
		db:	 db,
		logger: logger,
	}
}

func (e *PetController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "pet/:id", e.getPet, nil),
		routes.RegisterRoute(http.MethodGet, "pets", e.listPets, nil),
		routes.RegisterRoute(http.MethodPut, "pet", e.createPet, nil),
		routes.RegisterRoute(http.MethodDelete, "pet/:id", e.deletePet, nil),
		routes.RegisterRoute(http.MethodPatch, "pet/:id", e.updatePet, nil),
		routes.RegisterRoute(http.MethodPost, "pets/bulk", e.getPetsBulk, nil),
	}
}

// listPets godoc
// @Id listPets
// @Summary Lists Pets
// @Accept json
// @Produce json
// @Tags Pet
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Merchantlists<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcSpells<br>NpcType.NpcSpells.NpcSpellsEntries<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Pet
// @Failure 500 {string} string "Bad query request"
// @Router /pets [get]
func (e *PetController) listPets(c echo.Context) error {
	var results []models.Pet
	err := e.db.QueryContext(models.Pet{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getPet godoc
// @Id getPet
// @Summary Gets Pet
// @Accept json
// @Produce json
// @Tags Pet
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Merchantlists<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcSpells<br>NpcType.NpcSpells.NpcSpellsEntries<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Pet
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /pet/{id} [get]
func (e *PetController) getPet(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.Pet
	query := e.db.QueryContext(models.Pet{}, c)
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

// updatePet godoc
// @Id updatePet
// @Summary Updates Pet
// @Accept json
// @Produce json
// @Tags Pet
// @Param id path int true "Id"
// @Param pet body models.Pet true "Pet"
// @Success 200 {array} models.Pet
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /pet/{id} [patch]
func (e *PetController) updatePet(c echo.Context) error {
	request := new(models.Pet)
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

	// query builder
	var result models.Pet
	query := e.db.QueryContext(models.Pet{}, c)
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

// createPet godoc
// @Id createPet
// @Summary Creates Pet
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Pet"
// @Tags Pet
// @Success 200 {array} models.Pet
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /pet [put]
func (e *PetController) createPet(c echo.Context) error {
	pet := new(models.Pet)
	if err := c.Bind(pet); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Pet{}, c).Model(&models.Pet{}).Create(&pet).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, pet)
}

// deletePet godoc
// @Id deletePet
// @Summary Deletes Pet
// @Accept json
// @Produce json
// @Tags Pet
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /pet/{id} [delete]
func (e *PetController) deletePet(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.Pet
	query := e.db.QueryContext(models.Pet{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.Pet{}, c).Model(&models.Pet{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getPetsBulk godoc
// @Id getPetsBulk
// @Summary Gets Pets in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Pet
// @Success 200 {array} models.Pet
// @Failure 500 {string} string "Bad query request"
// @Router /pets/bulk [post]
func (e *PetController) getPetsBulk(c echo.Context) error {
	var results []models.Pet

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

	err := e.db.QueryContext(models.Pet{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
