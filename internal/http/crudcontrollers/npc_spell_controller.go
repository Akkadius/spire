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

type NpcSpellController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcSpellController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcSpellController {
	return &NpcSpellController{
		db:	 db,
		logger: logger,
	}
}

func (e *NpcSpellController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_spell/:id", e.getNpcSpell, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells", e.listNpcSpells, nil),
		routes.RegisterRoute(http.MethodPut, "npc_spell", e.createNpcSpell, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_spell/:id", e.deleteNpcSpell, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_spell/:id", e.updateNpcSpell, nil),
		routes.RegisterRoute(http.MethodPost, "npc_spells/bulk", e.getNpcSpellsBulk, nil),
	}
}

// listNpcSpells godoc
// @Id listNpcSpells
// @Summary Lists NpcSpells
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcSpell<br>NpcSpellsEntries<br>NpcSpellsEntries.SpellsNew<br>NpcSpellsEntries.SpellsNew.Aura<br>NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcSpellsEntries.SpellsNew.Items<br>NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.Items<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcEmotes<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcSpell<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcTypesTint<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcSpellsEntries.SpellsNew.SpellGlobals"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpell
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells [get]
func (e *NpcSpellController) listNpcSpells(c echo.Context) error {
	var results []models.NpcSpell
	err := e.db.QueryContext(models.NpcSpell{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcSpell godoc
// @Id getNpcSpell
// @Summary Gets NpcSpell
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcSpell<br>NpcSpellsEntries<br>NpcSpellsEntries.SpellsNew<br>NpcSpellsEntries.SpellsNew.Aura<br>NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcSpellsEntries.SpellsNew.Items<br>NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.Items<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcEmotes<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcSpell<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcTypesTint<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcSpellsEntries.SpellsNew.SpellGlobals"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpell
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spell/{id} [get]
func (e *NpcSpellController) getNpcSpell(c echo.Context) error {
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
	var result models.NpcSpell
	query := e.db.QueryContext(models.NpcSpell{}, c)
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

// updateNpcSpell godoc
// @Id updateNpcSpell
// @Summary Updates NpcSpell
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param id path int true "Id"
// @Param npc_spell body models.NpcSpell true "NpcSpell"
// @Success 200 {array} models.NpcSpell
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_spell/{id} [patch]
func (e *NpcSpellController) updateNpcSpell(c echo.Context) error {
	request := new(models.NpcSpell)
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
	var result models.NpcSpell
	query := e.db.QueryContext(models.NpcSpell{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.NpcSpell{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcSpell godoc
// @Id createNpcSpell
// @Summary Creates NpcSpell
// @Accept json
// @Produce json
// @Param npc_spell body models.NpcSpell true "NpcSpell"
// @Tags NpcSpell
// @Success 200 {array} models.NpcSpell
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_spell [put]
func (e *NpcSpellController) createNpcSpell(c echo.Context) error {
	npcSpell := new(models.NpcSpell)
	if err := c.Bind(npcSpell); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcSpell{}, c).Model(&models.NpcSpell{}).Create(&npcSpell).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcSpell)
}

// deleteNpcSpell godoc
// @Id deleteNpcSpell
// @Summary Deletes NpcSpell
// @Accept json
// @Produce json
// @Tags NpcSpell
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_spell/{id} [delete]
func (e *NpcSpellController) deleteNpcSpell(c echo.Context) error {
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
	var result models.NpcSpell
	query := e.db.QueryContext(models.NpcSpell{}, c)
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

// getNpcSpellsBulk godoc
// @Id getNpcSpellsBulk
// @Summary Gets NpcSpells in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcSpell
// @Success 200 {array} models.NpcSpell
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells/bulk [post]
func (e *NpcSpellController) getNpcSpellsBulk(c echo.Context) error {
	var results []models.NpcSpell

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

	err := e.db.QueryContext(models.NpcSpell{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
