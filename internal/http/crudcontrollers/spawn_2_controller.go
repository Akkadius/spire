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

type Spawn2Controller struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewSpawn2Controller(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *Spawn2Controller {
	return &Spawn2Controller{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.NpcType.AlternateCurrency<br>Spawnentries.NpcType.AlternateCurrency.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.AlternateCurrencies<br>Spawnentries.NpcType.AlternateCurrency.Item.CharacterCorpseItems<br>Spawnentries.NpcType.AlternateCurrency.Item.DiscoveredItems<br>Spawnentries.NpcType.AlternateCurrency.Item.Doors<br>Spawnentries.NpcType.AlternateCurrency.Item.Doors.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings.NpcType<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.Forages<br>Spawnentries.NpcType.AlternateCurrency.Item.Forages.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Forages.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns<br>Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.ItemTicks<br>Spawnentries.NpcType.AlternateCurrency.Item.Keyrings<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists<br>Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.Items<br>Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Spawnentries.NpcType.AlternateCurrency.Item.ObjectContents<br>Spawnentries.NpcType.AlternateCurrency.Item.Objects<br>Spawnentries.NpcType.AlternateCurrency.Item.Objects.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Objects.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.StartingItems<br>Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.AlternateCurrency.Item.TributeLevels<br>Spawnentries.NpcType.Loottable<br>Spawnentries.NpcType.Loottable.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Loottable<br>Spawnentries.NpcType.Loottable.NpcTypes<br>Spawnentries.NpcType.Merchantlists<br>Spawnentries.NpcType.Merchantlists.Items<br>Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies<br>Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies.Item<br>Spawnentries.NpcType.Merchantlists.Items.CharacterCorpseItems<br>Spawnentries.NpcType.Merchantlists.Items.DiscoveredItems<br>Spawnentries.NpcType.Merchantlists.Items.Doors<br>Spawnentries.NpcType.Merchantlists.Items.Doors.Item<br>Spawnentries.NpcType.Merchantlists.Items.Fishings<br>Spawnentries.NpcType.Merchantlists.Items.Fishings.Item<br>Spawnentries.NpcType.Merchantlists.Items.Fishings.NpcType<br>Spawnentries.NpcType.Merchantlists.Items.Fishings.Zone<br>Spawnentries.NpcType.Merchantlists.Items.Forages<br>Spawnentries.NpcType.Merchantlists.Items.Forages.Item<br>Spawnentries.NpcType.Merchantlists.Items.Forages.Zone<br>Spawnentries.NpcType.Merchantlists.Items.GroundSpawns<br>Spawnentries.NpcType.Merchantlists.Items.GroundSpawns.Zone<br>Spawnentries.NpcType.Merchantlists.Items.ItemTicks<br>Spawnentries.NpcType.Merchantlists.Items.Keyrings<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Item<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.Merchantlists.Items.Merchantlists<br>Spawnentries.NpcType.Merchantlists.Items.ObjectContents<br>Spawnentries.NpcType.Merchantlists.Items.Objects<br>Spawnentries.NpcType.Merchantlists.Items.Objects.Item<br>Spawnentries.NpcType.Merchantlists.Items.Objects.Zone<br>Spawnentries.NpcType.Merchantlists.Items.StartingItems<br>Spawnentries.NpcType.Merchantlists.Items.StartingItems.Item<br>Spawnentries.NpcType.Merchantlists.Items.StartingItems.Zone<br>Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.Merchantlists.Items.TributeLevels<br>Spawnentries.NpcType.Merchantlists.NpcTypes<br>Spawnentries.NpcType.NpcEmotes<br>Spawnentries.NpcType.NpcFactions<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Spawnentries.NpcType.NpcSpell<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.NpcSpell<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.ItemTicks<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Keyrings<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.ObjectContents<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TributeLevels<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Spawnentries.NpcType.NpcSpell.NpcSpell<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Spawnentries.NpcType.NpcTypesTint<br>Spawnentries.NpcType.Spawnentries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawngroup<br>Spawngroup.Spawn2"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.NpcType.AlternateCurrency<br>Spawnentries.NpcType.AlternateCurrency.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.AlternateCurrencies<br>Spawnentries.NpcType.AlternateCurrency.Item.CharacterCorpseItems<br>Spawnentries.NpcType.AlternateCurrency.Item.DiscoveredItems<br>Spawnentries.NpcType.AlternateCurrency.Item.Doors<br>Spawnentries.NpcType.AlternateCurrency.Item.Doors.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings.NpcType<br>Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.Forages<br>Spawnentries.NpcType.AlternateCurrency.Item.Forages.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Forages.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns<br>Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.ItemTicks<br>Spawnentries.NpcType.AlternateCurrency.Item.Keyrings<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists<br>Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.Items<br>Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Spawnentries.NpcType.AlternateCurrency.Item.ObjectContents<br>Spawnentries.NpcType.AlternateCurrency.Item.Objects<br>Spawnentries.NpcType.AlternateCurrency.Item.Objects.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.Objects.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.StartingItems<br>Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Item<br>Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Zone<br>Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries<br>Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.AlternateCurrency.Item.TributeLevels<br>Spawnentries.NpcType.Loottable<br>Spawnentries.NpcType.Loottable.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.Loottable.LoottableEntries.Loottable<br>Spawnentries.NpcType.Loottable.NpcTypes<br>Spawnentries.NpcType.Merchantlists<br>Spawnentries.NpcType.Merchantlists.Items<br>Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies<br>Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies.Item<br>Spawnentries.NpcType.Merchantlists.Items.CharacterCorpseItems<br>Spawnentries.NpcType.Merchantlists.Items.DiscoveredItems<br>Spawnentries.NpcType.Merchantlists.Items.Doors<br>Spawnentries.NpcType.Merchantlists.Items.Doors.Item<br>Spawnentries.NpcType.Merchantlists.Items.Fishings<br>Spawnentries.NpcType.Merchantlists.Items.Fishings.Item<br>Spawnentries.NpcType.Merchantlists.Items.Fishings.NpcType<br>Spawnentries.NpcType.Merchantlists.Items.Fishings.Zone<br>Spawnentries.NpcType.Merchantlists.Items.Forages<br>Spawnentries.NpcType.Merchantlists.Items.Forages.Item<br>Spawnentries.NpcType.Merchantlists.Items.Forages.Zone<br>Spawnentries.NpcType.Merchantlists.Items.GroundSpawns<br>Spawnentries.NpcType.Merchantlists.Items.GroundSpawns.Zone<br>Spawnentries.NpcType.Merchantlists.Items.ItemTicks<br>Spawnentries.NpcType.Merchantlists.Items.Keyrings<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Item<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.Merchantlists.Items.Merchantlists<br>Spawnentries.NpcType.Merchantlists.Items.ObjectContents<br>Spawnentries.NpcType.Merchantlists.Items.Objects<br>Spawnentries.NpcType.Merchantlists.Items.Objects.Item<br>Spawnentries.NpcType.Merchantlists.Items.Objects.Zone<br>Spawnentries.NpcType.Merchantlists.Items.StartingItems<br>Spawnentries.NpcType.Merchantlists.Items.StartingItems.Item<br>Spawnentries.NpcType.Merchantlists.Items.StartingItems.Zone<br>Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.Merchantlists.Items.TributeLevels<br>Spawnentries.NpcType.Merchantlists.NpcTypes<br>Spawnentries.NpcType.NpcEmotes<br>Spawnentries.NpcType.NpcFactions<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Spawnentries.NpcType.NpcSpell<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.NpcSpell<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Doors.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Forages.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.ItemTicks<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Keyrings<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.ObjectContents<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.Objects.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items.TributeLevels<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets<br>Spawnentries.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals<br>Spawnentries.NpcType.NpcSpell.NpcSpell<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Spawnentries.NpcType.NpcTypesTint<br>Spawnentries.NpcType.Spawnentries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawngroup<br>Spawngroup.Spawn2"
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
		event := fmt.Sprintf("Updated [Spawn2] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
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

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.Spawn2{}, spawn2)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [Spawn2] [%v] data [%v]", spawn2.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [Spawn2] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
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
