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

type NpcSpellsEntryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcSpellsEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcSpellsEntryController {
	return &NpcSpellsEntryController{
		db:	 db,
		logger: logger,
	}
}

func (e *NpcSpellsEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_spells_entry/:id", e.getNpcSpellsEntry, nil),
		routes.RegisterRoute(http.MethodGet, "npc_spells_entries", e.listNpcSpellsEntries, nil),
		routes.RegisterRoute(http.MethodPut, "npc_spells_entry", e.createNpcSpellsEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_spells_entry/:id", e.deleteNpcSpellsEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_spells_entry/:id", e.updateNpcSpellsEntry, nil),
		routes.RegisterRoute(http.MethodPost, "npc_spells_entries/bulk", e.getNpcSpellsEntriesBulk, nil),
	}
}

// listNpcSpellsEntries godoc
// @Id listNpcSpellsEntries
// @Summary Lists NpcSpellsEntries
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>SpellsNew<br>SpellsNew.Aura<br>SpellsNew.Aura.SpellsNew<br>SpellsNew.BlockedSpells<br>SpellsNew.Damageshieldtypes<br>SpellsNew.Items<br>SpellsNew.Items.AlternateCurrencies<br>SpellsNew.Items.AlternateCurrencies.Item<br>SpellsNew.Items.CharacterCorpseItems<br>SpellsNew.Items.DiscoveredItems<br>SpellsNew.Items.Doors<br>SpellsNew.Items.Doors.Item<br>SpellsNew.Items.Fishings<br>SpellsNew.Items.Fishings.Item<br>SpellsNew.Items.Fishings.NpcType<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.Merchantlists<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.Items<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.NpcEmotes<br>SpellsNew.Items.Fishings.NpcType.NpcFactions<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Fishings.NpcType.NpcSpell<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Fishings.NpcType.NpcTypesTint<br>SpellsNew.Items.Fishings.NpcType.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Fishings.Zone<br>SpellsNew.Items.Forages<br>SpellsNew.Items.Forages.Item<br>SpellsNew.Items.Forages.Zone<br>SpellsNew.Items.GroundSpawns<br>SpellsNew.Items.GroundSpawns.Zone<br>SpellsNew.Items.ItemTicks<br>SpellsNew.Items.Keyrings<br>SpellsNew.Items.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Merchantlists<br>SpellsNew.Items.Merchantlists.Items<br>SpellsNew.Items.Merchantlists.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.ObjectContents<br>SpellsNew.Items.Objects<br>SpellsNew.Items.Objects.Item<br>SpellsNew.Items.Objects.Zone<br>SpellsNew.Items.StartingItems<br>SpellsNew.Items.StartingItems.Item<br>SpellsNew.Items.StartingItems.Zone<br>SpellsNew.Items.Tasks<br>SpellsNew.Items.Tasks.AlternateCurrency<br>SpellsNew.Items.Tasks.AlternateCurrency.Item<br>SpellsNew.Items.Tasks.TaskActivities<br>SpellsNew.Items.Tasks.TaskActivities.Goallists<br>SpellsNew.Items.Tasks.TaskActivities.NpcType<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.AlternateCurrency<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.AlternateCurrency.Item<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.Items<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.NpcTypes<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcEmotes<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcTypesTint<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Tasks.Tasksets<br>SpellsNew.Items.TradeskillRecipeEntries<br>SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>SpellsNew.Items.TributeLevels<br>SpellsNew.NpcSpellsEntries<br>SpellsNew.SpellBuckets<br>SpellsNew.SpellGlobals"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_entries [get]
func (e *NpcSpellsEntryController) listNpcSpellsEntries(c echo.Context) error {
	var results []models.NpcSpellsEntry
	err := e.db.QueryContext(models.NpcSpellsEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcSpellsEntry godoc
// @Id getNpcSpellsEntry
// @Summary Gets NpcSpellsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>SpellsNew<br>SpellsNew.Aura<br>SpellsNew.Aura.SpellsNew<br>SpellsNew.BlockedSpells<br>SpellsNew.Damageshieldtypes<br>SpellsNew.Items<br>SpellsNew.Items.AlternateCurrencies<br>SpellsNew.Items.AlternateCurrencies.Item<br>SpellsNew.Items.CharacterCorpseItems<br>SpellsNew.Items.DiscoveredItems<br>SpellsNew.Items.Doors<br>SpellsNew.Items.Doors.Item<br>SpellsNew.Items.Fishings<br>SpellsNew.Items.Fishings.Item<br>SpellsNew.Items.Fishings.NpcType<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.Merchantlists<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.Items<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.NpcEmotes<br>SpellsNew.Items.Fishings.NpcType.NpcFactions<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Fishings.NpcType.NpcSpell<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Fishings.NpcType.NpcTypesTint<br>SpellsNew.Items.Fishings.NpcType.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Fishings.Zone<br>SpellsNew.Items.Forages<br>SpellsNew.Items.Forages.Item<br>SpellsNew.Items.Forages.Zone<br>SpellsNew.Items.GroundSpawns<br>SpellsNew.Items.GroundSpawns.Zone<br>SpellsNew.Items.ItemTicks<br>SpellsNew.Items.Keyrings<br>SpellsNew.Items.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Merchantlists<br>SpellsNew.Items.Merchantlists.Items<br>SpellsNew.Items.Merchantlists.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.ObjectContents<br>SpellsNew.Items.Objects<br>SpellsNew.Items.Objects.Item<br>SpellsNew.Items.Objects.Zone<br>SpellsNew.Items.StartingItems<br>SpellsNew.Items.StartingItems.Item<br>SpellsNew.Items.StartingItems.Zone<br>SpellsNew.Items.Tasks<br>SpellsNew.Items.Tasks.AlternateCurrency<br>SpellsNew.Items.Tasks.AlternateCurrency.Item<br>SpellsNew.Items.Tasks.TaskActivities<br>SpellsNew.Items.Tasks.TaskActivities.Goallists<br>SpellsNew.Items.Tasks.TaskActivities.NpcType<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.AlternateCurrency<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.AlternateCurrency.Item<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.Items<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.NpcTypes<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcEmotes<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcTypesTint<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Tasks.Tasksets<br>SpellsNew.Items.TradeskillRecipeEntries<br>SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>SpellsNew.Items.TributeLevels<br>SpellsNew.NpcSpellsEntries<br>SpellsNew.SpellBuckets<br>SpellsNew.SpellGlobals"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_entry/{id} [get]
func (e *NpcSpellsEntryController) getNpcSpellsEntry(c echo.Context) error {
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
	var result models.NpcSpellsEntry
	query := e.db.QueryContext(models.NpcSpellsEntry{}, c)
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

// updateNpcSpellsEntry godoc
// @Id updateNpcSpellsEntry
// @Summary Updates NpcSpellsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param id path int true "Id"
// @Param npc_spells_entry body models.NpcSpellsEntry true "NpcSpellsEntry"
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_spells_entry/{id} [patch]
func (e *NpcSpellsEntryController) updateNpcSpellsEntry(c echo.Context) error {
	request := new(models.NpcSpellsEntry)
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
	var result models.NpcSpellsEntry
	query := e.db.QueryContext(models.NpcSpellsEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.NpcSpellsEntry{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcSpellsEntry godoc
// @Id createNpcSpellsEntry
// @Summary Creates NpcSpellsEntry
// @Accept json
// @Produce json
// @Param npc_spells_entry body models.NpcSpellsEntry true "NpcSpellsEntry"
// @Tags NpcSpellsEntry
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_spells_entry [put]
func (e *NpcSpellsEntryController) createNpcSpellsEntry(c echo.Context) error {
	npcSpellsEntry := new(models.NpcSpellsEntry)
	if err := c.Bind(npcSpellsEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcSpellsEntry{}, c).Model(&models.NpcSpellsEntry{}).Create(&npcSpellsEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcSpellsEntry)
}

// deleteNpcSpellsEntry godoc
// @Id deleteNpcSpellsEntry
// @Summary Deletes NpcSpellsEntry
// @Accept json
// @Produce json
// @Tags NpcSpellsEntry
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_spells_entry/{id} [delete]
func (e *NpcSpellsEntryController) deleteNpcSpellsEntry(c echo.Context) error {
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
	var result models.NpcSpellsEntry
	query := e.db.QueryContext(models.NpcSpellsEntry{}, c)
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

// getNpcSpellsEntriesBulk godoc
// @Id getNpcSpellsEntriesBulk
// @Summary Gets NpcSpellsEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcSpellsEntry
// @Success 200 {array} models.NpcSpellsEntry
// @Failure 500 {string} string "Bad query request"
// @Router /npc_spells_entries/bulk [post]
func (e *NpcSpellsEntryController) getNpcSpellsEntriesBulk(c echo.Context) error {
	var results []models.NpcSpellsEntry

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

	err := e.db.QueryContext(models.NpcSpellsEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
