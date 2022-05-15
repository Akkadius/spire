package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Spawn2Controller struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawn2Controller(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *Spawn2Controller {
	return &Spawn2Controller{
		db:	 db,
		logger: logger,
	}
}

func (e *Spawn2Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spawn_2/:id", e.getSpawn2, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_2s", e.listSpawn2s, nil),
		routes.RegisterRoute(http.MethodPut, "spawn_2", e.createSpawn2, nil),
		routes.RegisterRoute(http.MethodDelete, "spawn_2/:id", e.deleteSpawn2, nil),
		routes.RegisterRoute(http.MethodPatch, "spawn_2/:id", e.updateSpawn2, nil),
		routes.RegisterRoute(http.MethodPost, "spawn_2s/bulk", e.getSpawn2sBulk, nil),
	}
}

// listSpawn2s godoc
// @Id listSpawn2s
// @Summary Lists Spawn2s
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.NpcType.AlternateCurrency<br>Spawnentries.NpcType.Loottable<br>Spawnentries.NpcType.Loottable.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Loottable<br>Spawnentries.NpcType.Loottable.NpcTypes<br>Spawnentries.NpcType.Merchantlists<br>Spawnentries.NpcType.Merchantlists.NpcType<br>Spawnentries.NpcType.NpcEmotes<br>Spawnentries.NpcType.NpcFactions<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawnentries.NpcType.NpcSpells<br>Spawnentries.NpcType.NpcSpells.NpcSpellsEntries<br>Spawnentries.NpcType.NpcTypesTint<br>Spawnentries.NpcType.Spawnentries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawngroup<br>Spawngroup.Spawn2"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawn2
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2s [get]
func (e *Spawn2Controller) listSpawn2s(c echo.Context) error {
	var results []models.Spawn2
	err := e.db.QueryContext(models.Spawn2{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawn2 godoc
// @Id getSpawn2
// @Summary Gets Spawn2
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.NpcType.AlternateCurrency<br>Spawnentries.NpcType.Loottable<br>Spawnentries.NpcType.Loottable.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Loottable<br>Spawnentries.NpcType.Loottable.NpcTypes<br>Spawnentries.NpcType.Merchantlists<br>Spawnentries.NpcType.Merchantlists.NpcType<br>Spawnentries.NpcType.NpcEmotes<br>Spawnentries.NpcType.NpcFactions<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawnentries.NpcType.NpcSpells<br>Spawnentries.NpcType.NpcSpells.NpcSpellsEntries<br>Spawnentries.NpcType.NpcTypesTint<br>Spawnentries.NpcType.Spawnentries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawngroup<br>Spawngroup.Spawn2"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawn2
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2/{id} [get]
func (e *Spawn2Controller) getSpawn2(c echo.Context) error {
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
	var result models.Spawn2
	query := e.db.QueryContext(models.Spawn2{}, c)
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

// updateSpawn2 godoc
// @Id updateSpawn2
// @Summary Updates Spawn2
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param id path int true "Id"
// @Param spawn_2 body models.Spawn2 true "Spawn2"
// @Success 200 {array} models.Spawn2
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawn_2/{id} [patch]
func (e *Spawn2Controller) updateSpawn2(c echo.Context) error {
	request := new(models.Spawn2)
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
	var result models.Spawn2
	query := e.db.QueryContext(models.Spawn2{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Spawn2{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createSpawn2 godoc
// @Id createSpawn2
// @Summary Creates Spawn2
// @Accept json
// @Produce json
// @Param spawn_2 body models.Spawn2 true "Spawn2"
// @Tags Spawn2
// @Success 200 {array} models.Spawn2
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawn_2 [put]
func (e *Spawn2Controller) createSpawn2(c echo.Context) error {
	spawn2 := new(models.Spawn2)
	if err := c.Bind(spawn2); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Spawn2{}, c).Model(&models.Spawn2{}).Create(&spawn2).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spawn2)
}

// deleteSpawn2 godoc
// @Id deleteSpawn2
// @Summary Deletes Spawn2
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawn_2/{id} [delete]
func (e *Spawn2Controller) deleteSpawn2(c echo.Context) error {
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
	var result models.Spawn2
	query := e.db.QueryContext(models.Spawn2{}, c)
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

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpawn2sBulk godoc
// @Id getSpawn2sBulk
// @Summary Gets Spawn2s in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Spawn2
// @Success 200 {array} models.Spawn2
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2s/bulk [post]
func (e *Spawn2Controller) getSpawn2sBulk(c echo.Context) error {
	var results []models.Spawn2

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

	err := e.db.QueryContext(models.Spawn2{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
