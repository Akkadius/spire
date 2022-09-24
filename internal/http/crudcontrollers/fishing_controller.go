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

type FishingController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewFishingController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *FishingController {
	return &FishingController{
		db:	 db,
		logger: logger,
	}
}

func (e *FishingController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "fishing/:id", e.getFishing, nil),
		routes.RegisterRoute(http.MethodGet, "fishings", e.listFishings, nil),
		routes.RegisterRoute(http.MethodPut, "fishing", e.createFishing, nil),
		routes.RegisterRoute(http.MethodDelete, "fishing/:id", e.deleteFishing, nil),
		routes.RegisterRoute(http.MethodPatch, "fishing/:id", e.updateFishing, nil),
		routes.RegisterRoute(http.MethodPost, "fishings/bulk", e.getFishingsBulk, nil),
	}
}

// listFishings godoc
// @Id listFishings
// @Summary Lists Fishings
// @Accept json
// @Produce json
// @Tags Fishing
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.AlternateCurrencies.Item<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.Items<br>Item.Merchantlists.NpcTypes<br>Item.Merchantlists.NpcTypes.AlternateCurrency<br>Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Item.Merchantlists.NpcTypes.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Item.Merchantlists.NpcTypes.Merchantlists<br>Item.Merchantlists.NpcTypes.NpcEmotes<br>Item.Merchantlists.NpcTypes.NpcFactions<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.Merchantlists.NpcTypes.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Merchantlists.NpcTypes.NpcTypesTint<br>Item.Merchantlists.NpcTypes.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.Tasks<br>Item.Tasks.AlternateCurrency<br>Item.Tasks.AlternateCurrency.Item<br>Item.Tasks.TaskActivities<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>NpcType<br>NpcType.AlternateCurrency<br>NpcType.AlternateCurrency.Item<br>NpcType.AlternateCurrency.Item.AlternateCurrencies<br>NpcType.AlternateCurrency.Item.CharacterCorpseItems<br>NpcType.AlternateCurrency.Item.DiscoveredItems<br>NpcType.AlternateCurrency.Item.Doors<br>NpcType.AlternateCurrency.Item.Doors.Item<br>NpcType.AlternateCurrency.Item.Fishings<br>NpcType.AlternateCurrency.Item.Forages<br>NpcType.AlternateCurrency.Item.Forages.Item<br>NpcType.AlternateCurrency.Item.Forages.Zone<br>NpcType.AlternateCurrency.Item.GroundSpawns<br>NpcType.AlternateCurrency.Item.GroundSpawns.Zone<br>NpcType.AlternateCurrency.Item.ItemTicks<br>NpcType.AlternateCurrency.Item.Keyrings<br>NpcType.AlternateCurrency.Item.LootdropEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Item<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.AlternateCurrency.Item.Merchantlists<br>NpcType.AlternateCurrency.Item.Merchantlists.Items<br>NpcType.AlternateCurrency.Item.Merchantlists.NpcTypes<br>NpcType.AlternateCurrency.Item.ObjectContents<br>NpcType.AlternateCurrency.Item.Objects<br>NpcType.AlternateCurrency.Item.Objects.Item<br>NpcType.AlternateCurrency.Item.Objects.Zone<br>NpcType.AlternateCurrency.Item.StartingItems<br>NpcType.AlternateCurrency.Item.StartingItems.Item<br>NpcType.AlternateCurrency.Item.StartingItems.Zone<br>NpcType.AlternateCurrency.Item.Tasks<br>NpcType.AlternateCurrency.Item.Tasks.AlternateCurrency<br>NpcType.AlternateCurrency.Item.Tasks.TaskActivities<br>NpcType.AlternateCurrency.Item.Tasks.Tasksets<br>NpcType.AlternateCurrency.Item.TradeskillRecipeEntries<br>NpcType.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.AlternateCurrency.Item.TributeLevels<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.Merchantlists.Items<br>NpcType.Merchantlists.Items.AlternateCurrencies<br>NpcType.Merchantlists.Items.AlternateCurrencies.Item<br>NpcType.Merchantlists.Items.CharacterCorpseItems<br>NpcType.Merchantlists.Items.DiscoveredItems<br>NpcType.Merchantlists.Items.Doors<br>NpcType.Merchantlists.Items.Doors.Item<br>NpcType.Merchantlists.Items.Fishings<br>NpcType.Merchantlists.Items.Forages<br>NpcType.Merchantlists.Items.Forages.Item<br>NpcType.Merchantlists.Items.Forages.Zone<br>NpcType.Merchantlists.Items.GroundSpawns<br>NpcType.Merchantlists.Items.GroundSpawns.Zone<br>NpcType.Merchantlists.Items.ItemTicks<br>NpcType.Merchantlists.Items.Keyrings<br>NpcType.Merchantlists.Items.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Item<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.Merchantlists.Items.Merchantlists<br>NpcType.Merchantlists.Items.ObjectContents<br>NpcType.Merchantlists.Items.Objects<br>NpcType.Merchantlists.Items.Objects.Item<br>NpcType.Merchantlists.Items.Objects.Zone<br>NpcType.Merchantlists.Items.StartingItems<br>NpcType.Merchantlists.Items.StartingItems.Item<br>NpcType.Merchantlists.Items.StartingItems.Zone<br>NpcType.Merchantlists.Items.Tasks<br>NpcType.Merchantlists.Items.Tasks.AlternateCurrency<br>NpcType.Merchantlists.Items.Tasks.AlternateCurrency.Item<br>NpcType.Merchantlists.Items.Tasks.TaskActivities<br>NpcType.Merchantlists.Items.Tasks.Tasksets<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Merchantlists.Items.TributeLevels<br>NpcType.Merchantlists.NpcTypes<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcFactions.NpcFactionEntries.FactionList<br>NpcType.NpcSpell<br>NpcType.NpcSpell.NpcSpell<br>NpcType.NpcSpell.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.NpcType<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Zone"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Bad query request"
// @Router /fishings [get]
func (e *FishingController) listFishings(c echo.Context) error {
	var results []models.Fishing
	err := e.db.QueryContext(models.Fishing{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getFishing godoc
// @Id getFishing
// @Summary Gets Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.AlternateCurrencies.Item<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.Items<br>Item.Merchantlists.NpcTypes<br>Item.Merchantlists.NpcTypes.AlternateCurrency<br>Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Item.Merchantlists.NpcTypes.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Item.Merchantlists.NpcTypes.Merchantlists<br>Item.Merchantlists.NpcTypes.NpcEmotes<br>Item.Merchantlists.NpcTypes.NpcFactions<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.Merchantlists.NpcTypes.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Merchantlists.NpcTypes.NpcTypesTint<br>Item.Merchantlists.NpcTypes.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.Tasks<br>Item.Tasks.AlternateCurrency<br>Item.Tasks.AlternateCurrency.Item<br>Item.Tasks.TaskActivities<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>NpcType<br>NpcType.AlternateCurrency<br>NpcType.AlternateCurrency.Item<br>NpcType.AlternateCurrency.Item.AlternateCurrencies<br>NpcType.AlternateCurrency.Item.CharacterCorpseItems<br>NpcType.AlternateCurrency.Item.DiscoveredItems<br>NpcType.AlternateCurrency.Item.Doors<br>NpcType.AlternateCurrency.Item.Doors.Item<br>NpcType.AlternateCurrency.Item.Fishings<br>NpcType.AlternateCurrency.Item.Forages<br>NpcType.AlternateCurrency.Item.Forages.Item<br>NpcType.AlternateCurrency.Item.Forages.Zone<br>NpcType.AlternateCurrency.Item.GroundSpawns<br>NpcType.AlternateCurrency.Item.GroundSpawns.Zone<br>NpcType.AlternateCurrency.Item.ItemTicks<br>NpcType.AlternateCurrency.Item.Keyrings<br>NpcType.AlternateCurrency.Item.LootdropEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Item<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.AlternateCurrency.Item.Merchantlists<br>NpcType.AlternateCurrency.Item.Merchantlists.Items<br>NpcType.AlternateCurrency.Item.Merchantlists.NpcTypes<br>NpcType.AlternateCurrency.Item.ObjectContents<br>NpcType.AlternateCurrency.Item.Objects<br>NpcType.AlternateCurrency.Item.Objects.Item<br>NpcType.AlternateCurrency.Item.Objects.Zone<br>NpcType.AlternateCurrency.Item.StartingItems<br>NpcType.AlternateCurrency.Item.StartingItems.Item<br>NpcType.AlternateCurrency.Item.StartingItems.Zone<br>NpcType.AlternateCurrency.Item.Tasks<br>NpcType.AlternateCurrency.Item.Tasks.AlternateCurrency<br>NpcType.AlternateCurrency.Item.Tasks.TaskActivities<br>NpcType.AlternateCurrency.Item.Tasks.Tasksets<br>NpcType.AlternateCurrency.Item.TradeskillRecipeEntries<br>NpcType.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.AlternateCurrency.Item.TributeLevels<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency.Item<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.Merchantlists.Items<br>NpcType.Merchantlists.Items.AlternateCurrencies<br>NpcType.Merchantlists.Items.AlternateCurrencies.Item<br>NpcType.Merchantlists.Items.CharacterCorpseItems<br>NpcType.Merchantlists.Items.DiscoveredItems<br>NpcType.Merchantlists.Items.Doors<br>NpcType.Merchantlists.Items.Doors.Item<br>NpcType.Merchantlists.Items.Fishings<br>NpcType.Merchantlists.Items.Forages<br>NpcType.Merchantlists.Items.Forages.Item<br>NpcType.Merchantlists.Items.Forages.Zone<br>NpcType.Merchantlists.Items.GroundSpawns<br>NpcType.Merchantlists.Items.GroundSpawns.Zone<br>NpcType.Merchantlists.Items.ItemTicks<br>NpcType.Merchantlists.Items.Keyrings<br>NpcType.Merchantlists.Items.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Item<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.Merchantlists.Items.Merchantlists<br>NpcType.Merchantlists.Items.ObjectContents<br>NpcType.Merchantlists.Items.Objects<br>NpcType.Merchantlists.Items.Objects.Item<br>NpcType.Merchantlists.Items.Objects.Zone<br>NpcType.Merchantlists.Items.StartingItems<br>NpcType.Merchantlists.Items.StartingItems.Item<br>NpcType.Merchantlists.Items.StartingItems.Zone<br>NpcType.Merchantlists.Items.Tasks<br>NpcType.Merchantlists.Items.Tasks.AlternateCurrency<br>NpcType.Merchantlists.Items.Tasks.AlternateCurrency.Item<br>NpcType.Merchantlists.Items.Tasks.TaskActivities<br>NpcType.Merchantlists.Items.Tasks.Tasksets<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Merchantlists.Items.TributeLevels<br>NpcType.Merchantlists.NpcTypes<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcFactions.NpcFactionEntries.FactionList<br>NpcType.NpcSpell<br>NpcType.NpcSpell.NpcSpell<br>NpcType.NpcSpell.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency.Item<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.NpcType<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Zone"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Fishing
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /fishing/{id} [get]
func (e *FishingController) getFishing(c echo.Context) error {
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
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

// updateFishing godoc
// @Id updateFishing
// @Summary Updates Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "Id"
// @Param fishing body models.Fishing true "Fishing"
// @Success 200 {array} models.Fishing
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /fishing/{id} [patch]
func (e *FishingController) updateFishing(c echo.Context) error {
	request := new(models.Fishing)
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Fishing{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createFishing godoc
// @Id createFishing
// @Summary Creates Fishing
// @Accept json
// @Produce json
// @Param fishing body models.Fishing true "Fishing"
// @Tags Fishing
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /fishing [put]
func (e *FishingController) createFishing(c echo.Context) error {
	fishing := new(models.Fishing)
	if err := c.Bind(fishing); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Fishing{}, c).Model(&models.Fishing{}).Create(&fishing).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, fishing)
}

// deleteFishing godoc
// @Id deleteFishing
// @Summary Deletes Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /fishing/{id} [delete]
func (e *FishingController) deleteFishing(c echo.Context) error {
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
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

// getFishingsBulk godoc
// @Id getFishingsBulk
// @Summary Gets Fishings in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Fishing
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Bad query request"
// @Router /fishings/bulk [post]
func (e *FishingController) getFishingsBulk(c echo.Context) error {
	var results []models.Fishing

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

	err := e.db.QueryContext(models.Fishing{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
