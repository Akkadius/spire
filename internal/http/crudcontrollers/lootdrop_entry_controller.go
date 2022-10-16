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

type LootdropEntryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewLootdropEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *LootdropEntryController {
	return &LootdropEntryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *LootdropEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "lootdrop_entry/:lootdropId", e.getLootdropEntry, nil),
		routes.RegisterRoute(http.MethodGet, "lootdrop_entries", e.listLootdropEntries, nil),
		routes.RegisterRoute(http.MethodPut, "lootdrop_entry", e.createLootdropEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "lootdrop_entry/:lootdropId", e.deleteLootdropEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "lootdrop_entry/:lootdropId", e.updateLootdropEntry, nil),
		routes.RegisterRoute(http.MethodPost, "lootdrop_entries/bulk", e.getLootdropEntriesBulk, nil),
	}
}

// listLootdropEntries godoc
// @Id listLootdropEntries
// @Summary Lists LootdropEntries
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.AlternateCurrencies.Item<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.AlternateCurrency.Item<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.Items<br>Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Item.Fishings.NpcType.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.Merchantlists<br>Item.Merchantlists.Items<br>Item.Merchantlists.NpcTypes<br>Item.Merchantlists.NpcTypes.AlternateCurrency<br>Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Item.Merchantlists.NpcTypes.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Item.Merchantlists.NpcTypes.Merchantlists<br>Item.Merchantlists.NpcTypes.NpcEmotes<br>Item.Merchantlists.NpcTypes.NpcFactions<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.Merchantlists.NpcTypes.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Merchantlists.NpcTypes.NpcTypesTint<br>Item.Merchantlists.NpcTypes.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>Lootdrop<br>Lootdrop.LootdropEntries<br>Lootdrop.LoottableEntries<br>Lootdrop.LoottableEntries.Lootdrop<br>Lootdrop.LoottableEntries.Loottable<br>Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Doors<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Doors.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.ItemTicks<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Keyrings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.ObjectContents<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TributeLevels<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.DiscoveredItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Doors<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Doors.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.GroundSpawns<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.ItemTicks<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Keyrings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.ObjectContents<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TributeLevels<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entries [get]
func (e *LootdropEntryController) listLootdropEntries(c echo.Context) error {
	var results []models.LootdropEntry
	err := e.db.QueryContext(models.LootdropEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLootdropEntry godoc
// @Id getLootdropEntry
// @Summary Gets LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.AlternateCurrencies.Item<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.AlternateCurrency.Item<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.Items<br>Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Item.Fishings.NpcType.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.Merchantlists<br>Item.Merchantlists.Items<br>Item.Merchantlists.NpcTypes<br>Item.Merchantlists.NpcTypes.AlternateCurrency<br>Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Item.Merchantlists.NpcTypes.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Item.Merchantlists.NpcTypes.Merchantlists<br>Item.Merchantlists.NpcTypes.NpcEmotes<br>Item.Merchantlists.NpcTypes.NpcFactions<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.Merchantlists.NpcTypes.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Merchantlists.NpcTypes.NpcTypesTint<br>Item.Merchantlists.NpcTypes.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>Lootdrop<br>Lootdrop.LootdropEntries<br>Lootdrop.LoottableEntries<br>Lootdrop.LoottableEntries.Lootdrop<br>Lootdrop.LoottableEntries.Loottable<br>Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Doors<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Doors.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.ItemTicks<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Keyrings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.ObjectContents<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TributeLevels<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.DiscoveredItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Doors<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Doors.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.GroundSpawns<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.ItemTicks<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Keyrings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.ObjectContents<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TributeLevels<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LootdropEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entry/{id} [get]
func (e *LootdropEntryController) getLootdropEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LootdropId]"})
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.LootdropId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLootdropEntry godoc
// @Id updateLootdropEntry
// @Summary Updates LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "Id"
// @Param lootdrop_entry body models.LootdropEntry true "LootdropEntry"
// @Success 200 {array} models.LootdropEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /lootdrop_entry/{id} [patch]
func (e *LootdropEntryController) updateLootdropEntry(c echo.Context) error {
	request := new(models.LootdropEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LootdropId]"})
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
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
		event := fmt.Sprintf("Updated [LootdropEntry] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createLootdropEntry godoc
// @Id createLootdropEntry
// @Summary Creates LootdropEntry
// @Accept json
// @Produce json
// @Param lootdrop_entry body models.LootdropEntry true "LootdropEntry"
// @Tags LootdropEntry
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /lootdrop_entry [put]
func (e *LootdropEntryController) createLootdropEntry(c echo.Context) error {
	lootdropEntry := new(models.LootdropEntry)
	if err := c.Bind(lootdropEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LootdropEntry{}, c).Model(&models.LootdropEntry{}).Create(&lootdropEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.LootdropEntry{}, lootdropEntry)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [LootdropEntry] [%v] data [%v]", lootdropEntry.LootdropId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, lootdropEntry)
}

// deleteLootdropEntry godoc
// @Id deleteLootdropEntry
// @Summary Deletes LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "lootdropId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /lootdrop_entry/{id} [delete]
func (e *LootdropEntryController) deleteLootdropEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
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
		event := fmt.Sprintf("Deleted [LootdropEntry] [%v] keys [%v]", result.LootdropId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLootdropEntriesBulk godoc
// @Id getLootdropEntriesBulk
// @Summary Gets LootdropEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LootdropEntry
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entries/bulk [post]
func (e *LootdropEntryController) getLootdropEntriesBulk(c echo.Context) error {
	var results []models.LootdropEntry

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

	err := e.db.QueryContext(models.LootdropEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
