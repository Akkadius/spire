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

type GlobalLootController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewGlobalLootController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GlobalLootController {
	return &GlobalLootController{
		db:	 db,
		logger: logger,
	}
}

func (e *GlobalLootController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "global_loot/:id", e.getGlobalLoot, nil),
		routes.RegisterRoute(http.MethodGet, "global_loots", e.listGlobalLoots, nil),
		routes.RegisterRoute(http.MethodPut, "global_loot", e.createGlobalLoot, nil),
		routes.RegisterRoute(http.MethodDelete, "global_loot/:id", e.deleteGlobalLoot, nil),
		routes.RegisterRoute(http.MethodPatch, "global_loot/:id", e.updateGlobalLoot, nil),
		routes.RegisterRoute(http.MethodPost, "global_loots/bulk", e.getGlobalLootsBulk, nil),
	}
}

// listGlobalLoots godoc
// @Id listGlobalLoots
// @Summary Lists GlobalLoots
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Loottable<br>Loottable.LoottableEntries<br>Loottable.LoottableEntries.LootdropEntries<br>Loottable.LoottableEntries.LootdropEntries.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.LoottableEntries.Loottable<br>Loottable.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency<br>Loottable.NpcTypes.Loottable<br>Loottable.NpcTypes.Merchantlists<br>Loottable.NpcTypes.Merchantlists.NpcType<br>Loottable.NpcTypes.NpcEmotes<br>Loottable.NpcTypes.NpcFactions<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.NpcTypes.NpcSpells<br>Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Loottable.NpcTypes.NpcTypesTint<br>Loottable.NpcTypes.Spawnentries<br>Loottable.NpcTypes.Spawnentries.NpcType<br>Loottable.NpcTypes.Spawnentries.Spawngroup<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GlobalLoot
// @Failure 500 {string} string "Bad query request"
// @Router /global_loots [get]
func (e *GlobalLootController) listGlobalLoots(c echo.Context) error {
	var results []models.GlobalLoot
	err := e.db.QueryContext(models.GlobalLoot{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getGlobalLoot godoc
// @Id getGlobalLoot
// @Summary Gets GlobalLoot
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Loottable<br>Loottable.LoottableEntries<br>Loottable.LoottableEntries.LootdropEntries<br>Loottable.LoottableEntries.LootdropEntries.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.LoottableEntries.Loottable<br>Loottable.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency<br>Loottable.NpcTypes.Loottable<br>Loottable.NpcTypes.Merchantlists<br>Loottable.NpcTypes.Merchantlists.NpcType<br>Loottable.NpcTypes.NpcEmotes<br>Loottable.NpcTypes.NpcFactions<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.NpcTypes.NpcSpells<br>Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Loottable.NpcTypes.NpcTypesTint<br>Loottable.NpcTypes.Spawnentries<br>Loottable.NpcTypes.Spawnentries.NpcType<br>Loottable.NpcTypes.Spawnentries.Spawngroup<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.GlobalLoot
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /global_loot/{id} [get]
func (e *GlobalLootController) getGlobalLoot(c echo.Context) error {
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
	var result models.GlobalLoot
	query := e.db.QueryContext(models.GlobalLoot{}, c)
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

// updateGlobalLoot godoc
// @Id updateGlobalLoot
// @Summary Updates GlobalLoot
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param id path int true "Id"
// @Param global_loot body models.GlobalLoot true "GlobalLoot"
// @Success 200 {array} models.GlobalLoot
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /global_loot/{id} [patch]
func (e *GlobalLootController) updateGlobalLoot(c echo.Context) error {
	request := new(models.GlobalLoot)
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
	var result models.GlobalLoot
	query := e.db.QueryContext(models.GlobalLoot{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.GlobalLoot{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createGlobalLoot godoc
// @Id createGlobalLoot
// @Summary Creates GlobalLoot
// @Accept json
// @Produce json
// @Param global_loot body models.GlobalLoot true "GlobalLoot"
// @Tags GlobalLoot
// @Success 200 {array} models.GlobalLoot
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /global_loot [put]
func (e *GlobalLootController) createGlobalLoot(c echo.Context) error {
	globalLoot := new(models.GlobalLoot)
	if err := c.Bind(globalLoot); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.GlobalLoot{}, c).Model(&models.GlobalLoot{}).Create(&globalLoot).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, globalLoot)
}

// deleteGlobalLoot godoc
// @Id deleteGlobalLoot
// @Summary Deletes GlobalLoot
// @Accept json
// @Produce json
// @Tags GlobalLoot
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /global_loot/{id} [delete]
func (e *GlobalLootController) deleteGlobalLoot(c echo.Context) error {
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
	var result models.GlobalLoot
	query := e.db.QueryContext(models.GlobalLoot{}, c)
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

// getGlobalLootsBulk godoc
// @Id getGlobalLootsBulk
// @Summary Gets GlobalLoots in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags GlobalLoot
// @Success 200 {array} models.GlobalLoot
// @Failure 500 {string} string "Bad query request"
// @Router /global_loots/bulk [post]
func (e *GlobalLootController) getGlobalLootsBulk(c echo.Context) error {
	var results []models.GlobalLoot

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

	err := e.db.QueryContext(models.GlobalLoot{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
