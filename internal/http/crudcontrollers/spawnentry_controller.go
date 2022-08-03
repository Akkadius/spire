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

type SpawnentryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawnentryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpawnentryController {
	return &SpawnentryController{
		db:	 db,
		logger: logger,
	}
}

func (e *SpawnentryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spawnentry/:spawngroupID", e.getSpawnentry, nil),
		routes.RegisterRoute(http.MethodGet, "spawnentries", e.listSpawnentries, nil),
		routes.RegisterRoute(http.MethodPut, "spawnentry", e.createSpawnentry, nil),
		routes.RegisterRoute(http.MethodDelete, "spawnentry/:spawngroupID", e.deleteSpawnentry, nil),
		routes.RegisterRoute(http.MethodPatch, "spawnentry/:spawngroupID", e.updateSpawnentry, nil),
		routes.RegisterRoute(http.MethodPost, "spawnentries/bulk", e.getSpawnentriesBulk, nil),
	}
}

// listSpawnentries godoc
// @Id listSpawnentries
// @Summary Lists Spawnentries
// @Accept json
// @Produce json
// @Tags Spawnentry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.Merchantlists.Items<br>NpcType.Merchantlists.Items.AlternateCurrencies<br>NpcType.Merchantlists.Items.CharacterCorpseItems<br>NpcType.Merchantlists.Items.DiscoveredItems<br>NpcType.Merchantlists.Items.Doors<br>NpcType.Merchantlists.Items.Doors.Item<br>NpcType.Merchantlists.Items.Fishings<br>NpcType.Merchantlists.Items.Fishings.Item<br>NpcType.Merchantlists.Items.Fishings.NpcType<br>NpcType.Merchantlists.Items.Fishings.Zone<br>NpcType.Merchantlists.Items.Forages<br>NpcType.Merchantlists.Items.Forages.Item<br>NpcType.Merchantlists.Items.Forages.Zone<br>NpcType.Merchantlists.Items.GroundSpawns<br>NpcType.Merchantlists.Items.GroundSpawns.Zone<br>NpcType.Merchantlists.Items.ItemTicks<br>NpcType.Merchantlists.Items.Keyrings<br>NpcType.Merchantlists.Items.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Item<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.Merchantlists.Items.Merchantlists<br>NpcType.Merchantlists.Items.ObjectContents<br>NpcType.Merchantlists.Items.Objects<br>NpcType.Merchantlists.Items.Objects.Item<br>NpcType.Merchantlists.Items.Objects.Zone<br>NpcType.Merchantlists.Items.StartingItems<br>NpcType.Merchantlists.Items.StartingItems.Item<br>NpcType.Merchantlists.Items.StartingItems.Zone<br>NpcType.Merchantlists.Items.Tasks<br>NpcType.Merchantlists.Items.Tasks.TaskActivities<br>NpcType.Merchantlists.Items.Tasks.TaskActivities.Goallists<br>NpcType.Merchantlists.Items.Tasks.TaskActivities.NpcType<br>NpcType.Merchantlists.Items.Tasks.Tasksets<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Merchantlists.Items.TributeLevels<br>NpcType.Merchantlists.NpcTypes<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcFactions.NpcFactionEntries.FactionList<br>NpcType.NpcSpell<br>NpcType.NpcSpell.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>Spawngroup<br>Spawngroup.Spawn2<br>Spawngroup.Spawn2.Spawnentries<br>Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawnentry
// @Failure 500 {string} string "Bad query request"
// @Router /spawnentries [get]
func (e *SpawnentryController) listSpawnentries(c echo.Context) error {
	var results []models.Spawnentry
	err := e.db.QueryContext(models.Spawnentry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawnentry godoc
// @Id getSpawnentry
// @Summary Gets Spawnentry
// @Accept json
// @Produce json
// @Tags Spawnentry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.Merchantlists.Items<br>NpcType.Merchantlists.Items.AlternateCurrencies<br>NpcType.Merchantlists.Items.CharacterCorpseItems<br>NpcType.Merchantlists.Items.DiscoveredItems<br>NpcType.Merchantlists.Items.Doors<br>NpcType.Merchantlists.Items.Doors.Item<br>NpcType.Merchantlists.Items.Fishings<br>NpcType.Merchantlists.Items.Fishings.Item<br>NpcType.Merchantlists.Items.Fishings.NpcType<br>NpcType.Merchantlists.Items.Fishings.Zone<br>NpcType.Merchantlists.Items.Forages<br>NpcType.Merchantlists.Items.Forages.Item<br>NpcType.Merchantlists.Items.Forages.Zone<br>NpcType.Merchantlists.Items.GroundSpawns<br>NpcType.Merchantlists.Items.GroundSpawns.Zone<br>NpcType.Merchantlists.Items.ItemTicks<br>NpcType.Merchantlists.Items.Keyrings<br>NpcType.Merchantlists.Items.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Item<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.Merchantlists.Items.Merchantlists<br>NpcType.Merchantlists.Items.ObjectContents<br>NpcType.Merchantlists.Items.Objects<br>NpcType.Merchantlists.Items.Objects.Item<br>NpcType.Merchantlists.Items.Objects.Zone<br>NpcType.Merchantlists.Items.StartingItems<br>NpcType.Merchantlists.Items.StartingItems.Item<br>NpcType.Merchantlists.Items.StartingItems.Zone<br>NpcType.Merchantlists.Items.Tasks<br>NpcType.Merchantlists.Items.Tasks.TaskActivities<br>NpcType.Merchantlists.Items.Tasks.TaskActivities.Goallists<br>NpcType.Merchantlists.Items.Tasks.TaskActivities.NpcType<br>NpcType.Merchantlists.Items.Tasks.Tasksets<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Merchantlists.Items.TributeLevels<br>NpcType.Merchantlists.NpcTypes<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcFactions.NpcFactionEntries.FactionList<br>NpcType.NpcSpell<br>NpcType.NpcSpell.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>Spawngroup<br>Spawngroup.Spawn2<br>Spawngroup.Spawn2.Spawnentries<br>Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawnentry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawnentry/{id} [get]
func (e *SpawnentryController) getSpawnentry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	spawngroupID, err := strconv.Atoi(c.Param("spawngroupID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SpawngroupID]"})
	}
	params = append(params, spawngroupID)
	keys = append(keys, "spawngroupID = ?")

	// key param [npcID] position [2] type [int]
	if len(c.QueryParam("npcID")) > 0 {
		npcIDParam, err := strconv.Atoi(c.QueryParam("npcID"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [npcID] err [%s]", err.Error())})
		}

		params = append(params, npcIDParam)
		keys = append(keys, "npcID = ?")
	}

	// query builder
	var result models.Spawnentry
	query := e.db.QueryContext(models.Spawnentry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.SpawngroupID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSpawnentry godoc
// @Id updateSpawnentry
// @Summary Updates Spawnentry
// @Accept json
// @Produce json
// @Tags Spawnentry
// @Param id path int true "Id"
// @Param spawnentry body models.Spawnentry true "Spawnentry"
// @Success 200 {array} models.Spawnentry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawnentry/{id} [patch]
func (e *SpawnentryController) updateSpawnentry(c echo.Context) error {
	request := new(models.Spawnentry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	spawngroupID, err := strconv.Atoi(c.Param("spawngroupID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [SpawngroupID]"})
	}
	params = append(params, spawngroupID)
	keys = append(keys, "spawngroupID = ?")

	// key param [npcID] position [2] type [int]
	if len(c.QueryParam("npcID")) > 0 {
		npcIDParam, err := strconv.Atoi(c.QueryParam("npcID"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [npcID] err [%s]", err.Error())})
		}

		params = append(params, npcIDParam)
		keys = append(keys, "npcID = ?")
	}

	// query builder
	var result models.Spawnentry
	query := e.db.QueryContext(models.Spawnentry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Spawnentry{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createSpawnentry godoc
// @Id createSpawnentry
// @Summary Creates Spawnentry
// @Accept json
// @Produce json
// @Param spawnentry body models.Spawnentry true "Spawnentry"
// @Tags Spawnentry
// @Success 200 {array} models.Spawnentry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawnentry [put]
func (e *SpawnentryController) createSpawnentry(c echo.Context) error {
	spawnentry := new(models.Spawnentry)
	if err := c.Bind(spawnentry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Spawnentry{}, c).Model(&models.Spawnentry{}).Create(&spawnentry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spawnentry)
}

// deleteSpawnentry godoc
// @Id deleteSpawnentry
// @Summary Deletes Spawnentry
// @Accept json
// @Produce json
// @Tags Spawnentry
// @Param id path int true "spawngroupID"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawnentry/{id} [delete]
func (e *SpawnentryController) deleteSpawnentry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	spawngroupID, err := strconv.Atoi(c.Param("spawngroupID"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, spawngroupID)
	keys = append(keys, "spawngroupID = ?")

	// key param [npcID] position [2] type [int]
	if len(c.QueryParam("npcID")) > 0 {
		npcIDParam, err := strconv.Atoi(c.QueryParam("npcID"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [npcID] err [%s]", err.Error())})
		}

		params = append(params, npcIDParam)
		keys = append(keys, "npcID = ?")
	}

	// query builder
	var result models.Spawnentry
	query := e.db.QueryContext(models.Spawnentry{}, c)
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

// getSpawnentriesBulk godoc
// @Id getSpawnentriesBulk
// @Summary Gets Spawnentries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Spawnentry
// @Success 200 {array} models.Spawnentry
// @Failure 500 {string} string "Bad query request"
// @Router /spawnentries/bulk [post]
func (e *SpawnentryController) getSpawnentriesBulk(c echo.Context) error {
	var results []models.Spawnentry

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

	err := e.db.QueryContext(models.Spawnentry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
