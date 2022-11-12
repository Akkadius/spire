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
	"net/http"
	"strconv"
	"strings"
)

type GlobalLootController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewGlobalLootController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *GlobalLootController {
	return &GlobalLootController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Loottable<br>Loottable.LoottableEntries<br>Loottable.LoottableEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcEmotes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcTypesTint<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Loottable.LoottableEntries.Loottable<br>Loottable.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency<br>Loottable.NpcTypes.AlternateCurrency.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>Loottable.NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>Loottable.NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.ItemTicks<br>Loottable.NpcTypes.AlternateCurrency.Item.Keyrings<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency.Item.ObjectContents<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.AlternateCurrency.Item.TributeLevels<br>Loottable.NpcTypes.Loottable<br>Loottable.NpcTypes.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>Loottable.NpcTypes.Merchantlists.Items.DiscoveredItems<br>Loottable.NpcTypes.Merchantlists.Items.Doors<br>Loottable.NpcTypes.Merchantlists.Items.Doors.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.NpcType<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Zone<br>Loottable.NpcTypes.Merchantlists.Items.Forages<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Item<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Zone<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.Merchantlists.Items.ItemTicks<br>Loottable.NpcTypes.Merchantlists.Items.Keyrings<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Item<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.Merchantlists.Items.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items.ObjectContents<br>Loottable.NpcTypes.Merchantlists.Items.Objects<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Item<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Zone<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Item<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Zone<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.Merchantlists.Items.TributeLevels<br>Loottable.NpcTypes.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcEmotes<br>Loottable.NpcTypes.NpcFactions<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.NpcTypes.NpcSpell<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.ItemTicks<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Keyrings<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.ObjectContents<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.TributeLevels<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.NpcTypes.NpcSpell.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.NpcTypes.NpcTypesTint<br>Loottable.NpcTypes.Spawnentries<br>Loottable.NpcTypes.Spawnentries.NpcType<br>Loottable.NpcTypes.Spawnentries.Spawngroup<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Loottable<br>Loottable.LoottableEntries<br>Loottable.LoottableEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcEmotes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcTypesTint<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Loottable.LoottableEntries.Loottable<br>Loottable.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency<br>Loottable.NpcTypes.AlternateCurrency.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>Loottable.NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>Loottable.NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.ItemTicks<br>Loottable.NpcTypes.AlternateCurrency.Item.Keyrings<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency.Item.ObjectContents<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.AlternateCurrency.Item.TributeLevels<br>Loottable.NpcTypes.Loottable<br>Loottable.NpcTypes.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>Loottable.NpcTypes.Merchantlists.Items.DiscoveredItems<br>Loottable.NpcTypes.Merchantlists.Items.Doors<br>Loottable.NpcTypes.Merchantlists.Items.Doors.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.NpcType<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Zone<br>Loottable.NpcTypes.Merchantlists.Items.Forages<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Item<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Zone<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.Merchantlists.Items.ItemTicks<br>Loottable.NpcTypes.Merchantlists.Items.Keyrings<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Item<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.Merchantlists.Items.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items.ObjectContents<br>Loottable.NpcTypes.Merchantlists.Items.Objects<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Item<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Zone<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Item<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Zone<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.Merchantlists.Items.TributeLevels<br>Loottable.NpcTypes.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcEmotes<br>Loottable.NpcTypes.NpcFactions<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.NpcTypes.NpcSpell<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.ItemTicks<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Keyrings<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.ObjectContents<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items.TributeLevels<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.NpcTypes.NpcSpell.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.NpcTypes.NpcTypesTint<br>Loottable.NpcTypes.Spawnentries<br>Loottable.NpcTypes.Spawnentries.NpcType<br>Loottable.NpcTypes.Spawnentries.Spawngroup<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
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
		event := fmt.Sprintf("Updated [GlobalLoot] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
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

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.GlobalLoot{}, globalLoot)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [GlobalLoot] [%v] data [%v]", globalLoot.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [GlobalLoot] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
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
