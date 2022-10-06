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

type LoottableEntryController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoottableEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoottableEntryController {
	return &LoottableEntryController{
		db:	    db,
		logger: logger,
	}
}

func (e *LoottableEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "loottable_entry/:loottableId", e.getLoottableEntry, nil),
		routes.RegisterRoute(http.MethodGet, "loottable_entries", e.listLoottableEntries, nil),
		routes.RegisterRoute(http.MethodPut, "loottable_entry", e.createLoottableEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "loottable_entry/:loottableId", e.deleteLoottableEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "loottable_entry/:loottableId", e.updateLoottableEntry, nil),
		routes.RegisterRoute(http.MethodPost, "loottable_entries/bulk", e.getLoottableEntriesBulk, nil),
	}
}

// listLoottableEntries godoc
// @Id listLoottableEntries
// @Summary Lists LoottableEntries
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Lootdrop<br>Lootdrop.LootdropEntries<br>Lootdrop.LootdropEntries.Item<br>Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Lootdrop.LootdropEntries.Item.Doors<br>Lootdrop.LootdropEntries.Item.Doors.Item<br>Lootdrop.LootdropEntries.Item.Fishings<br>Lootdrop.LootdropEntries.Item.Fishings.Item<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency.Item<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Items<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Lootdrop.LootdropEntries.Item.Forages<br>Lootdrop.LootdropEntries.Item.Forages.Item<br>Lootdrop.LootdropEntries.Item.Forages.Zone<br>Lootdrop.LootdropEntries.Item.GroundSpawns<br>Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Lootdrop.LootdropEntries.Item.ItemTicks<br>Lootdrop.LootdropEntries.Item.Keyrings<br>Lootdrop.LootdropEntries.Item.LootdropEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists<br>Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Merchantlists<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcEmotes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcTypesTint<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Lootdrop.LootdropEntries.Item.ObjectContents<br>Lootdrop.LootdropEntries.Item.Objects<br>Lootdrop.LootdropEntries.Item.Objects.Item<br>Lootdrop.LootdropEntries.Item.Objects.Zone<br>Lootdrop.LootdropEntries.Item.StartingItems<br>Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LootdropEntries.Item.TributeLevels<br>Lootdrop.LootdropEntries.Lootdrop<br>Lootdrop.LoottableEntries<br>Loottable<br>Loottable.LoottableEntries<br>Loottable.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency<br>Loottable.NpcTypes.AlternateCurrency.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>Loottable.NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>Loottable.NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.ItemTicks<br>Loottable.NpcTypes.AlternateCurrency.Item.Keyrings<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency.Item.ObjectContents<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.AlternateCurrency.Item.TributeLevels<br>Loottable.NpcTypes.Loottable<br>Loottable.NpcTypes.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>Loottable.NpcTypes.Merchantlists.Items.DiscoveredItems<br>Loottable.NpcTypes.Merchantlists.Items.Doors<br>Loottable.NpcTypes.Merchantlists.Items.Doors.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.NpcType<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Zone<br>Loottable.NpcTypes.Merchantlists.Items.Forages<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Item<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Zone<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.Merchantlists.Items.ItemTicks<br>Loottable.NpcTypes.Merchantlists.Items.Keyrings<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Item<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.Merchantlists.Items.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items.ObjectContents<br>Loottable.NpcTypes.Merchantlists.Items.Objects<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Item<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Zone<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Item<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Zone<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.Merchantlists.Items.TributeLevels<br>Loottable.NpcTypes.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcEmotes<br>Loottable.NpcTypes.NpcFactions<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.NpcTypes.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.NpcTypes.NpcTypesTint<br>Loottable.NpcTypes.Spawnentries<br>Loottable.NpcTypes.Spawnentries.NpcType<br>Loottable.NpcTypes.Spawnentries.Spawngroup<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entries [get]
func (e *LoottableEntryController) listLoottableEntries(c echo.Context) error {
	var results []models.LoottableEntry
	err := e.db.QueryContext(models.LoottableEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoottableEntry godoc
// @Id getLoottableEntry
// @Summary Gets LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Lootdrop<br>Lootdrop.LootdropEntries<br>Lootdrop.LootdropEntries.Item<br>Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Lootdrop.LootdropEntries.Item.Doors<br>Lootdrop.LootdropEntries.Item.Doors.Item<br>Lootdrop.LootdropEntries.Item.Fishings<br>Lootdrop.LootdropEntries.Item.Fishings.Item<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency.Item<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Items<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Lootdrop.LootdropEntries.Item.Forages<br>Lootdrop.LootdropEntries.Item.Forages.Item<br>Lootdrop.LootdropEntries.Item.Forages.Zone<br>Lootdrop.LootdropEntries.Item.GroundSpawns<br>Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Lootdrop.LootdropEntries.Item.ItemTicks<br>Lootdrop.LootdropEntries.Item.Keyrings<br>Lootdrop.LootdropEntries.Item.LootdropEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists<br>Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Merchantlists<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcEmotes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcTypesTint<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Lootdrop.LootdropEntries.Item.ObjectContents<br>Lootdrop.LootdropEntries.Item.Objects<br>Lootdrop.LootdropEntries.Item.Objects.Item<br>Lootdrop.LootdropEntries.Item.Objects.Zone<br>Lootdrop.LootdropEntries.Item.StartingItems<br>Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Lootdrop.LootdropEntries.Item.TributeLevels<br>Lootdrop.LootdropEntries.Lootdrop<br>Lootdrop.LoottableEntries<br>Loottable<br>Loottable.LoottableEntries<br>Loottable.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency<br>Loottable.NpcTypes.AlternateCurrency.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>Loottable.NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>Loottable.NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors<br>Loottable.NpcTypes.AlternateCurrency.Item.Doors.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Forages.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns<br>Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.ItemTicks<br>Loottable.NpcTypes.AlternateCurrency.Item.Keyrings<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Loottable.NpcTypes.AlternateCurrency.Item.ObjectContents<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.Objects.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.AlternateCurrency.Item.TributeLevels<br>Loottable.NpcTypes.Loottable<br>Loottable.NpcTypes.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies<br>Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>Loottable.NpcTypes.Merchantlists.Items.DiscoveredItems<br>Loottable.NpcTypes.Merchantlists.Items.Doors<br>Loottable.NpcTypes.Merchantlists.Items.Doors.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Item<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.NpcType<br>Loottable.NpcTypes.Merchantlists.Items.Fishings.Zone<br>Loottable.NpcTypes.Merchantlists.Items.Forages<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Item<br>Loottable.NpcTypes.Merchantlists.Items.Forages.Zone<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns<br>Loottable.NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.Merchantlists.Items.ItemTicks<br>Loottable.NpcTypes.Merchantlists.Items.Keyrings<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Item<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.Merchantlists.Items.Merchantlists<br>Loottable.NpcTypes.Merchantlists.Items.ObjectContents<br>Loottable.NpcTypes.Merchantlists.Items.Objects<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Item<br>Loottable.NpcTypes.Merchantlists.Items.Objects.Zone<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Item<br>Loottable.NpcTypes.Merchantlists.Items.StartingItems.Zone<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.Merchantlists.Items.TributeLevels<br>Loottable.NpcTypes.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcEmotes<br>Loottable.NpcTypes.NpcFactions<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Loottable.NpcTypes.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpell<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Loottable.NpcTypes.NpcTypesTint<br>Loottable.NpcTypes.Spawnentries<br>Loottable.NpcTypes.Spawnentries.NpcType<br>Loottable.NpcTypes.Spawnentries.Spawngroup<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LoottableEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entry/{id} [get]
func (e *LoottableEntryController) getLoottableEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LoottableId]"})
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.LoottableId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLoottableEntry godoc
// @Id updateLoottableEntry
// @Summary Updates LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "Id"
// @Param loottable_entry body models.LoottableEntry true "LoottableEntry"
// @Success 200 {array} models.LoottableEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /loottable_entry/{id} [patch]
func (e *LoottableEntryController) updateLoottableEntry(c echo.Context) error {
	request := new(models.LoottableEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LoottableId]"})
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(database.ResultDifference(result, request)).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createLoottableEntry godoc
// @Id createLoottableEntry
// @Summary Creates LoottableEntry
// @Accept json
// @Produce json
// @Param loottable_entry body models.LoottableEntry true "LoottableEntry"
// @Tags LoottableEntry
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /loottable_entry [put]
func (e *LoottableEntryController) createLoottableEntry(c echo.Context) error {
	loottableEntry := new(models.LoottableEntry)
	if err := c.Bind(loottableEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LoottableEntry{}, c).Model(&models.LoottableEntry{}).Create(&loottableEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, loottableEntry)
}

// deleteLoottableEntry godoc
// @Id deleteLoottableEntry
// @Summary Deletes LoottableEntry
// @Accept json
// @Produce json
// @Tags LoottableEntry
// @Param id path int true "loottableId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /loottable_entry/{id} [delete]
func (e *LoottableEntryController) deleteLoottableEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	loottableId, err := strconv.Atoi(c.Param("loottableId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, loottableId)
	keys = append(keys, "loottable_id = ?")

	// key param [lootdrop_id] position [2] type [int]
	if len(c.QueryParam("lootdrop_id")) > 0 {
		lootdropIdParam, err := strconv.Atoi(c.QueryParam("lootdrop_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [lootdrop_id] err [%s]", err.Error())})
		}

		params = append(params, lootdropIdParam)
		keys = append(keys, "lootdrop_id = ?")
	}

	// query builder
	var result models.LoottableEntry
	query := e.db.QueryContext(models.LoottableEntry{}, c)
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

// getLoottableEntriesBulk godoc
// @Id getLoottableEntriesBulk
// @Summary Gets LoottableEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LoottableEntry
// @Success 200 {array} models.LoottableEntry
// @Failure 500 {string} string "Bad query request"
// @Router /loottable_entries/bulk [post]
func (e *LoottableEntryController) getLoottableEntriesBulk(c echo.Context) error {
	var results []models.LoottableEntry

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

	err := e.db.QueryContext(models.LoottableEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
