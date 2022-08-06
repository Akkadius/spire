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

type NpcTypeController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewNpcTypeController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *NpcTypeController {
	return &NpcTypeController{
		db:	 db,
		logger: logger,
	}
}

func (e *NpcTypeController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "npc_type/:id", e.getNpcType, nil),
		routes.RegisterRoute(http.MethodGet, "npc_types", e.listNpcTypes, nil),
		routes.RegisterRoute(http.MethodPut, "npc_type", e.createNpcType, nil),
		routes.RegisterRoute(http.MethodDelete, "npc_type/:id", e.deleteNpcType, nil),
		routes.RegisterRoute(http.MethodPatch, "npc_type/:id", e.updateNpcType, nil),
		routes.RegisterRoute(http.MethodPost, "npc_types/bulk", e.getNpcTypesBulk, nil),
	}
}

// listNpcTypes godoc
// @Id listNpcTypes
// @Summary Lists NpcTypes
// @Accept json
// @Produce json
// @Tags NpcType
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrency<br>AlternateCurrency.Item<br>AlternateCurrency.Item.AlternateCurrencies<br>AlternateCurrency.Item.CharacterCorpseItems<br>AlternateCurrency.Item.DiscoveredItems<br>AlternateCurrency.Item.Doors<br>AlternateCurrency.Item.Doors.Item<br>AlternateCurrency.Item.Fishings<br>AlternateCurrency.Item.Fishings.Item<br>AlternateCurrency.Item.Fishings.NpcType<br>AlternateCurrency.Item.Fishings.Zone<br>AlternateCurrency.Item.Forages<br>AlternateCurrency.Item.Forages.Item<br>AlternateCurrency.Item.Forages.Zone<br>AlternateCurrency.Item.GroundSpawns<br>AlternateCurrency.Item.GroundSpawns.Zone<br>AlternateCurrency.Item.ItemTicks<br>AlternateCurrency.Item.Keyrings<br>AlternateCurrency.Item.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Item<br>AlternateCurrency.Item.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>AlternateCurrency.Item.Merchantlists<br>AlternateCurrency.Item.Merchantlists.Items<br>AlternateCurrency.Item.Merchantlists.NpcTypes<br>AlternateCurrency.Item.ObjectContents<br>AlternateCurrency.Item.Objects<br>AlternateCurrency.Item.Objects.Item<br>AlternateCurrency.Item.Objects.Zone<br>AlternateCurrency.Item.StartingItems<br>AlternateCurrency.Item.StartingItems.Item<br>AlternateCurrency.Item.StartingItems.Zone<br>AlternateCurrency.Item.Tasks<br>AlternateCurrency.Item.Tasks.AlternateCurrency<br>AlternateCurrency.Item.Tasks.TaskActivities<br>AlternateCurrency.Item.Tasks.TaskActivities.Goallists<br>AlternateCurrency.Item.Tasks.TaskActivities.NpcType<br>AlternateCurrency.Item.Tasks.Tasksets<br>AlternateCurrency.Item.TradeskillRecipeEntries<br>AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>AlternateCurrency.Item.TributeLevels<br>Loottable<br>Loottable.LoottableEntries<br>Loottable.LoottableEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Loottable.LoottableEntries.Loottable<br>Loottable.NpcTypes<br>Merchantlists<br>Merchantlists.Items<br>Merchantlists.Items.AlternateCurrencies<br>Merchantlists.Items.AlternateCurrencies.Item<br>Merchantlists.Items.CharacterCorpseItems<br>Merchantlists.Items.DiscoveredItems<br>Merchantlists.Items.Doors<br>Merchantlists.Items.Doors.Item<br>Merchantlists.Items.Fishings<br>Merchantlists.Items.Fishings.Item<br>Merchantlists.Items.Fishings.NpcType<br>Merchantlists.Items.Fishings.Zone<br>Merchantlists.Items.Forages<br>Merchantlists.Items.Forages.Item<br>Merchantlists.Items.Forages.Zone<br>Merchantlists.Items.GroundSpawns<br>Merchantlists.Items.GroundSpawns.Zone<br>Merchantlists.Items.ItemTicks<br>Merchantlists.Items.Keyrings<br>Merchantlists.Items.LootdropEntries<br>Merchantlists.Items.LootdropEntries.Item<br>Merchantlists.Items.LootdropEntries.Lootdrop<br>Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Merchantlists.Items.Merchantlists<br>Merchantlists.Items.ObjectContents<br>Merchantlists.Items.Objects<br>Merchantlists.Items.Objects.Item<br>Merchantlists.Items.Objects.Zone<br>Merchantlists.Items.StartingItems<br>Merchantlists.Items.StartingItems.Item<br>Merchantlists.Items.StartingItems.Zone<br>Merchantlists.Items.Tasks<br>Merchantlists.Items.Tasks.AlternateCurrency<br>Merchantlists.Items.Tasks.AlternateCurrency.Item<br>Merchantlists.Items.Tasks.TaskActivities<br>Merchantlists.Items.Tasks.TaskActivities.Goallists<br>Merchantlists.Items.Tasks.TaskActivities.NpcType<br>Merchantlists.Items.Tasks.Tasksets<br>Merchantlists.Items.TradeskillRecipeEntries<br>Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Merchantlists.Items.TributeLevels<br>Merchantlists.NpcTypes<br>NpcEmotes<br>NpcFactions<br>NpcFactions.NpcFactionEntries<br>NpcFactions.NpcFactionEntries.FactionList<br>NpcSpell<br>NpcSpell.NpcSpellsEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcTypesTint<br>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcType
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types [get]
func (e *NpcTypeController) listNpcTypes(c echo.Context) error {
	var results []models.NpcType
	err := e.db.QueryContext(models.NpcType{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getNpcType godoc
// @Id getNpcType
// @Summary Gets NpcType
// @Accept json
// @Produce json
// @Tags NpcType
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrency<br>AlternateCurrency.Item<br>AlternateCurrency.Item.AlternateCurrencies<br>AlternateCurrency.Item.CharacterCorpseItems<br>AlternateCurrency.Item.DiscoveredItems<br>AlternateCurrency.Item.Doors<br>AlternateCurrency.Item.Doors.Item<br>AlternateCurrency.Item.Fishings<br>AlternateCurrency.Item.Fishings.Item<br>AlternateCurrency.Item.Fishings.NpcType<br>AlternateCurrency.Item.Fishings.Zone<br>AlternateCurrency.Item.Forages<br>AlternateCurrency.Item.Forages.Item<br>AlternateCurrency.Item.Forages.Zone<br>AlternateCurrency.Item.GroundSpawns<br>AlternateCurrency.Item.GroundSpawns.Zone<br>AlternateCurrency.Item.ItemTicks<br>AlternateCurrency.Item.Keyrings<br>AlternateCurrency.Item.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Item<br>AlternateCurrency.Item.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>AlternateCurrency.Item.Merchantlists<br>AlternateCurrency.Item.Merchantlists.Items<br>AlternateCurrency.Item.Merchantlists.NpcTypes<br>AlternateCurrency.Item.ObjectContents<br>AlternateCurrency.Item.Objects<br>AlternateCurrency.Item.Objects.Item<br>AlternateCurrency.Item.Objects.Zone<br>AlternateCurrency.Item.StartingItems<br>AlternateCurrency.Item.StartingItems.Item<br>AlternateCurrency.Item.StartingItems.Zone<br>AlternateCurrency.Item.Tasks<br>AlternateCurrency.Item.Tasks.AlternateCurrency<br>AlternateCurrency.Item.Tasks.TaskActivities<br>AlternateCurrency.Item.Tasks.TaskActivities.Goallists<br>AlternateCurrency.Item.Tasks.TaskActivities.NpcType<br>AlternateCurrency.Item.Tasks.Tasksets<br>AlternateCurrency.Item.TradeskillRecipeEntries<br>AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>AlternateCurrency.Item.TributeLevels<br>Loottable<br>Loottable.LoottableEntries<br>Loottable.LoottableEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.AlternateCurrency.Item<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Loottable.LoottableEntries.Loottable<br>Loottable.NpcTypes<br>Merchantlists<br>Merchantlists.Items<br>Merchantlists.Items.AlternateCurrencies<br>Merchantlists.Items.AlternateCurrencies.Item<br>Merchantlists.Items.CharacterCorpseItems<br>Merchantlists.Items.DiscoveredItems<br>Merchantlists.Items.Doors<br>Merchantlists.Items.Doors.Item<br>Merchantlists.Items.Fishings<br>Merchantlists.Items.Fishings.Item<br>Merchantlists.Items.Fishings.NpcType<br>Merchantlists.Items.Fishings.Zone<br>Merchantlists.Items.Forages<br>Merchantlists.Items.Forages.Item<br>Merchantlists.Items.Forages.Zone<br>Merchantlists.Items.GroundSpawns<br>Merchantlists.Items.GroundSpawns.Zone<br>Merchantlists.Items.ItemTicks<br>Merchantlists.Items.Keyrings<br>Merchantlists.Items.LootdropEntries<br>Merchantlists.Items.LootdropEntries.Item<br>Merchantlists.Items.LootdropEntries.Lootdrop<br>Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Merchantlists.Items.Merchantlists<br>Merchantlists.Items.ObjectContents<br>Merchantlists.Items.Objects<br>Merchantlists.Items.Objects.Item<br>Merchantlists.Items.Objects.Zone<br>Merchantlists.Items.StartingItems<br>Merchantlists.Items.StartingItems.Item<br>Merchantlists.Items.StartingItems.Zone<br>Merchantlists.Items.Tasks<br>Merchantlists.Items.Tasks.AlternateCurrency<br>Merchantlists.Items.Tasks.AlternateCurrency.Item<br>Merchantlists.Items.Tasks.TaskActivities<br>Merchantlists.Items.Tasks.TaskActivities.Goallists<br>Merchantlists.Items.Tasks.TaskActivities.NpcType<br>Merchantlists.Items.Tasks.Tasksets<br>Merchantlists.Items.TradeskillRecipeEntries<br>Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>Merchantlists.Items.TributeLevels<br>Merchantlists.NpcTypes<br>NpcEmotes<br>NpcFactions<br>NpcFactions.NpcFactionEntries<br>NpcFactions.NpcFactionEntries.FactionList<br>NpcSpell<br>NpcSpell.NpcSpellsEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency.Item<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcTypesTint<br>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.NpcType
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /npc_type/{id} [get]
func (e *NpcTypeController) getNpcType(c echo.Context) error {
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
	var result models.NpcType
	query := e.db.QueryContext(models.NpcType{}, c)
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

// updateNpcType godoc
// @Id updateNpcType
// @Summary Updates NpcType
// @Accept json
// @Produce json
// @Tags NpcType
// @Param id path int true "Id"
// @Param npc_type body models.NpcType true "NpcType"
// @Success 200 {array} models.NpcType
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /npc_type/{id} [patch]
func (e *NpcTypeController) updateNpcType(c echo.Context) error {
	request := new(models.NpcType)
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
	var result models.NpcType
	query := e.db.QueryContext(models.NpcType{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.NpcType{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createNpcType godoc
// @Id createNpcType
// @Summary Creates NpcType
// @Accept json
// @Produce json
// @Param npc_type body models.NpcType true "NpcType"
// @Tags NpcType
// @Success 200 {array} models.NpcType
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /npc_type [put]
func (e *NpcTypeController) createNpcType(c echo.Context) error {
	npcType := new(models.NpcType)
	if err := c.Bind(npcType); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.NpcType{}, c).Model(&models.NpcType{}).Create(&npcType).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, npcType)
}

// deleteNpcType godoc
// @Id deleteNpcType
// @Summary Deletes NpcType
// @Accept json
// @Produce json
// @Tags NpcType
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /npc_type/{id} [delete]
func (e *NpcTypeController) deleteNpcType(c echo.Context) error {
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
	var result models.NpcType
	query := e.db.QueryContext(models.NpcType{}, c)
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

// getNpcTypesBulk godoc
// @Id getNpcTypesBulk
// @Summary Gets NpcTypes in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags NpcType
// @Success 200 {array} models.NpcType
// @Failure 500 {string} string "Bad query request"
// @Router /npc_types/bulk [post]
func (e *NpcTypeController) getNpcTypesBulk(c echo.Context) error {
	var results []models.NpcType

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

	err := e.db.QueryContext(models.NpcType{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
