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

type MerchantlistController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewMerchantlistController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *MerchantlistController {
	return &MerchantlistController{
		db:	 db,
		logger: logger,
	}
}

func (e *MerchantlistController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "merchantlist/:merchantid", e.getMerchantlist, nil),
		routes.RegisterRoute(http.MethodGet, "merchantlists", e.listMerchantlists, nil),
		routes.RegisterRoute(http.MethodPut, "merchantlist", e.createMerchantlist, nil),
		routes.RegisterRoute(http.MethodDelete, "merchantlist/:merchantid", e.deleteMerchantlist, nil),
		routes.RegisterRoute(http.MethodPatch, "merchantlist/:merchantid", e.updateMerchantlist, nil),
		routes.RegisterRoute(http.MethodPost, "merchantlists/bulk", e.getMerchantlistsBulk, nil),
	}
}

// listMerchantlists godoc
// @Id listMerchantlists
// @Summary Lists Merchantlists
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Items<br>Items.AlternateCurrencies<br>Items.CharacterCorpseItems<br>Items.DiscoveredItems<br>Items.Doors<br>Items.Doors.Item<br>Items.Fishings<br>Items.Fishings.Item<br>Items.Fishings.NpcType<br>Items.Fishings.NpcType.AlternateCurrency<br>Items.Fishings.NpcType.Loottable<br>Items.Fishings.NpcType.Loottable.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Items.Fishings.NpcType.Loottable.NpcTypes<br>Items.Fishings.NpcType.Merchantlists<br>Items.Fishings.NpcType.NpcEmotes<br>Items.Fishings.NpcType.NpcFactions<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Items.Fishings.NpcType.NpcSpell<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Items.Fishings.NpcType.NpcTypesTint<br>Items.Fishings.NpcType.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.NpcType<br>Items.Fishings.NpcType.Spawnentries.Spawngroup<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Fishings.Zone<br>Items.Forages<br>Items.Forages.Item<br>Items.Forages.Zone<br>Items.GroundSpawns<br>Items.GroundSpawns.Zone<br>Items.ItemTicks<br>Items.Keyrings<br>Items.LootdropEntries<br>Items.LootdropEntries.Item<br>Items.LootdropEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LootdropEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Merchantlists<br>Items.ObjectContents<br>Items.Objects<br>Items.Objects.Item<br>Items.Objects.Zone<br>Items.StartingItems<br>Items.StartingItems.Item<br>Items.StartingItems.Zone<br>Items.Tasks<br>Items.Tasks.TaskActivities<br>Items.Tasks.TaskActivities.Goallists<br>Items.Tasks.TaskActivities.NpcType<br>Items.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Items.Tasks.TaskActivities.NpcType.Loottable<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Items.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Items.Tasks.TaskActivities.NpcType.Merchantlists<br>Items.Tasks.TaskActivities.NpcType.NpcEmotes<br>Items.Tasks.TaskActivities.NpcType.NpcFactions<br>Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Items.Tasks.TaskActivities.NpcType.NpcSpell<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Items.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Items.Tasks.TaskActivities.NpcType.Spawnentries<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Tasks.Tasksets<br>Items.TradeskillRecipeEntries<br>Items.TradeskillRecipeEntries.TradeskillRecipe<br>Items.TributeLevels<br>NpcTypes<br>NpcTypes.AlternateCurrency<br>NpcTypes.Loottable<br>NpcTypes.Loottable.LoottableEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcTypes.Loottable.LoottableEntries.Loottable<br>NpcTypes.Loottable.NpcTypes<br>NpcTypes.Merchantlists<br>NpcTypes.NpcEmotes<br>NpcTypes.NpcFactions<br>NpcTypes.NpcFactions.NpcFactionEntries<br>NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcTypes.NpcSpell<br>NpcTypes.NpcSpell.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcTypes.NpcTypesTint<br>NpcTypes.Spawnentries<br>NpcTypes.Spawnentries.NpcType<br>NpcTypes.Spawnentries.Spawngroup<br>NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Merchantlist
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlists [get]
func (e *MerchantlistController) listMerchantlists(c echo.Context) error {
	var results []models.Merchantlist
	err := e.db.QueryContext(models.Merchantlist{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getMerchantlist godoc
// @Id getMerchantlist
// @Summary Gets Merchantlist
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Items<br>Items.AlternateCurrencies<br>Items.CharacterCorpseItems<br>Items.DiscoveredItems<br>Items.Doors<br>Items.Doors.Item<br>Items.Fishings<br>Items.Fishings.Item<br>Items.Fishings.NpcType<br>Items.Fishings.NpcType.AlternateCurrency<br>Items.Fishings.NpcType.Loottable<br>Items.Fishings.NpcType.Loottable.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Items.Fishings.NpcType.Loottable.NpcTypes<br>Items.Fishings.NpcType.Merchantlists<br>Items.Fishings.NpcType.NpcEmotes<br>Items.Fishings.NpcType.NpcFactions<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Items.Fishings.NpcType.NpcSpell<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Items.Fishings.NpcType.NpcTypesTint<br>Items.Fishings.NpcType.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.NpcType<br>Items.Fishings.NpcType.Spawnentries.Spawngroup<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Fishings.Zone<br>Items.Forages<br>Items.Forages.Item<br>Items.Forages.Zone<br>Items.GroundSpawns<br>Items.GroundSpawns.Zone<br>Items.ItemTicks<br>Items.Keyrings<br>Items.LootdropEntries<br>Items.LootdropEntries.Item<br>Items.LootdropEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LootdropEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Merchantlists<br>Items.ObjectContents<br>Items.Objects<br>Items.Objects.Item<br>Items.Objects.Zone<br>Items.StartingItems<br>Items.StartingItems.Item<br>Items.StartingItems.Zone<br>Items.Tasks<br>Items.Tasks.TaskActivities<br>Items.Tasks.TaskActivities.Goallists<br>Items.Tasks.TaskActivities.NpcType<br>Items.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Items.Tasks.TaskActivities.NpcType.Loottable<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Items.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Items.Tasks.TaskActivities.NpcType.Merchantlists<br>Items.Tasks.TaskActivities.NpcType.NpcEmotes<br>Items.Tasks.TaskActivities.NpcType.NpcFactions<br>Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Items.Tasks.TaskActivities.NpcType.NpcSpell<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Items.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Items.Tasks.TaskActivities.NpcType.Spawnentries<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Tasks.Tasksets<br>Items.TradeskillRecipeEntries<br>Items.TradeskillRecipeEntries.TradeskillRecipe<br>Items.TributeLevels<br>NpcTypes<br>NpcTypes.AlternateCurrency<br>NpcTypes.Loottable<br>NpcTypes.Loottable.LoottableEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.Tasks.Tasksets<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>NpcTypes.Loottable.LoottableEntries.Loottable<br>NpcTypes.Loottable.NpcTypes<br>NpcTypes.Merchantlists<br>NpcTypes.NpcEmotes<br>NpcTypes.NpcFactions<br>NpcTypes.NpcFactions.NpcFactionEntries<br>NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcTypes.NpcSpell<br>NpcTypes.NpcSpell.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcTypes.NpcTypesTint<br>NpcTypes.Spawnentries<br>NpcTypes.Spawnentries.NpcType<br>NpcTypes.Spawnentries.Spawngroup<br>NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Merchantlist
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlist/{id} [get]
func (e *MerchantlistController) getMerchantlist(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	merchantid, err := strconv.Atoi(c.Param("merchantid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Merchantid]"})
	}
	params = append(params, merchantid)
	keys = append(keys, "merchantid = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.Merchantlist
	query := e.db.QueryContext(models.Merchantlist{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Merchantid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateMerchantlist godoc
// @Id updateMerchantlist
// @Summary Updates Merchantlist
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param id path int true "Id"
// @Param merchantlist body models.Merchantlist true "Merchantlist"
// @Success 200 {array} models.Merchantlist
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /merchantlist/{id} [patch]
func (e *MerchantlistController) updateMerchantlist(c echo.Context) error {
	request := new(models.Merchantlist)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	merchantid, err := strconv.Atoi(c.Param("merchantid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Merchantid]"})
	}
	params = append(params, merchantid)
	keys = append(keys, "merchantid = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.Merchantlist
	query := e.db.QueryContext(models.Merchantlist{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Merchantlist{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createMerchantlist godoc
// @Id createMerchantlist
// @Summary Creates Merchantlist
// @Accept json
// @Produce json
// @Param merchantlist body models.Merchantlist true "Merchantlist"
// @Tags Merchantlist
// @Success 200 {array} models.Merchantlist
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /merchantlist [put]
func (e *MerchantlistController) createMerchantlist(c echo.Context) error {
	merchantlist := new(models.Merchantlist)
	if err := c.Bind(merchantlist); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Merchantlist{}, c).Model(&models.Merchantlist{}).Create(&merchantlist).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, merchantlist)
}

// deleteMerchantlist godoc
// @Id deleteMerchantlist
// @Summary Deletes Merchantlist
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param id path int true "merchantid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /merchantlist/{id} [delete]
func (e *MerchantlistController) deleteMerchantlist(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	merchantid, err := strconv.Atoi(c.Param("merchantid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, merchantid)
	keys = append(keys, "merchantid = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.Merchantlist
	query := e.db.QueryContext(models.Merchantlist{}, c)
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

// getMerchantlistsBulk godoc
// @Id getMerchantlistsBulk
// @Summary Gets Merchantlists in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Merchantlist
// @Success 200 {array} models.Merchantlist
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlists/bulk [post]
func (e *MerchantlistController) getMerchantlistsBulk(c echo.Context) error {
	var results []models.Merchantlist

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

	err := e.db.QueryContext(models.Merchantlist{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
