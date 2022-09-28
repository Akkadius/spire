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

type SpawngroupController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawngroupController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpawngroupController {
	return &SpawngroupController{
		db:	 db,
		logger: logger,
	}
}

func (e *SpawngroupController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spawngroup/:id", e.getSpawngroup, nil),
		routes.RegisterRoute(http.MethodGet, "spawngroups", e.listSpawngroups, nil),
		routes.RegisterRoute(http.MethodPut, "spawngroup", e.createSpawngroup, nil),
		routes.RegisterRoute(http.MethodDelete, "spawngroup/:id", e.deleteSpawngroup, nil),
		routes.RegisterRoute(http.MethodPatch, "spawngroup/:id", e.updateSpawngroup, nil),
		routes.RegisterRoute(http.MethodPost, "spawngroups/bulk", e.getSpawngroupsBulk, nil),
	}
}

// listSpawngroups godoc
// @Id listSpawngroups
// @Summary Lists Spawngroups
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawn2<br>Spawn2.Spawnentries<br>Spawn2.Spawnentries.NpcType<br>Spawn2.Spawnentries.NpcType.AlternateCurrency<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Doors<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Doors.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Forages<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Forages.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Forages.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.ItemTicks<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Keyrings<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.ObjectContents<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Objects<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Objects.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Objects.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.StartingItems<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.TributeLevels<br>Spawn2.Spawnentries.NpcType.Loottable<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.Merchantlists<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Doors<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Doors.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Forages<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Forages.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Forages.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.GroundSpawns<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.ItemTicks<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Keyrings<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Merchantlists<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.ObjectContents<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Objects<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Objects.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Objects.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.StartingItems<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.TributeLevels<br>Spawn2.Spawnentries.NpcType.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.NpcEmotes<br>Spawn2.Spawnentries.NpcType.NpcFactions<br>Spawn2.Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawn2.Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Spawn2.Spawnentries.NpcType.NpcSpell<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpell<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Spawn2.Spawnentries.NpcType.NpcTypesTint<br>Spawn2.Spawnentries.NpcType.Spawnentries<br>Spawn2.Spawnentries.Spawngroup<br>Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawngroup
// @Failure 500 {string} string "Bad query request"
// @Router /spawngroups [get]
func (e *SpawngroupController) listSpawngroups(c echo.Context) error {
	var results []models.Spawngroup
	err := e.db.QueryContext(models.Spawngroup{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawngroup godoc
// @Id getSpawngroup
// @Summary Gets Spawngroup
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawn2<br>Spawn2.Spawnentries<br>Spawn2.Spawnentries.NpcType<br>Spawn2.Spawnentries.NpcType.AlternateCurrency<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Doors<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Doors.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Forages<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Forages.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Forages.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.ItemTicks<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Keyrings<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.ObjectContents<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Objects<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Objects.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.Objects.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.StartingItems<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.AlternateCurrency.Item.TributeLevels<br>Spawn2.Spawnentries.NpcType.Loottable<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.Merchantlists<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.AlternateCurrencies.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Doors<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Doors.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Forages<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Forages.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Forages.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.GroundSpawns<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.ItemTicks<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Keyrings<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Merchantlists<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.ObjectContents<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Objects<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Objects.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.Objects.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.StartingItems<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.Merchantlists.Items.TributeLevels<br>Spawn2.Spawnentries.NpcType.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.NpcEmotes<br>Spawn2.Spawnentries.NpcType.NpcFactions<br>Spawn2.Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawn2.Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Spawn2.Spawnentries.NpcType.NpcSpell<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpell<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Spawn2.Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Spawn2.Spawnentries.NpcType.NpcTypesTint<br>Spawn2.Spawnentries.NpcType.Spawnentries<br>Spawn2.Spawnentries.Spawngroup<br>Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawngroup
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawngroup/{id} [get]
func (e *SpawngroupController) getSpawngroup(c echo.Context) error {
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
	var result models.Spawngroup
	query := e.db.QueryContext(models.Spawngroup{}, c)
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

// updateSpawngroup godoc
// @Id updateSpawngroup
// @Summary Updates Spawngroup
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param id path int true "Id"
// @Param spawngroup body models.Spawngroup true "Spawngroup"
// @Success 200 {array} models.Spawngroup
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawngroup/{id} [patch]
func (e *SpawngroupController) updateSpawngroup(c echo.Context) error {
	request := new(models.Spawngroup)
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
	var result models.Spawngroup
	query := e.db.QueryContext(models.Spawngroup{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Spawngroup{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createSpawngroup godoc
// @Id createSpawngroup
// @Summary Creates Spawngroup
// @Accept json
// @Produce json
// @Param spawngroup body models.Spawngroup true "Spawngroup"
// @Tags Spawngroup
// @Success 200 {array} models.Spawngroup
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawngroup [put]
func (e *SpawngroupController) createSpawngroup(c echo.Context) error {
	spawngroup := new(models.Spawngroup)
	if err := c.Bind(spawngroup); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Spawngroup{}, c).Model(&models.Spawngroup{}).Create(&spawngroup).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spawngroup)
}

// deleteSpawngroup godoc
// @Id deleteSpawngroup
// @Summary Deletes Spawngroup
// @Accept json
// @Produce json
// @Tags Spawngroup
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawngroup/{id} [delete]
func (e *SpawngroupController) deleteSpawngroup(c echo.Context) error {
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
	var result models.Spawngroup
	query := e.db.QueryContext(models.Spawngroup{}, c)
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

// getSpawngroupsBulk godoc
// @Id getSpawngroupsBulk
// @Summary Gets Spawngroups in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Spawngroup
// @Success 200 {array} models.Spawngroup
// @Failure 500 {string} string "Bad query request"
// @Router /spawngroups/bulk [post]
func (e *SpawngroupController) getSpawngroupsBulk(c echo.Context) error {
	var results []models.Spawngroup

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

	err := e.db.QueryContext(models.Spawngroup{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
