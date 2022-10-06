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

type LoottableController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLoottableController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LoottableController {
	return &LoottableController{
		db:	    db,
		logger: logger,
	}
}

func (e *LoottableController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "loottable/:id", e.getLoottable, nil),
		routes.RegisterRoute(http.MethodGet, "loottables", e.listLoottables, nil),
		routes.RegisterRoute(http.MethodPut, "loottable", e.createLoottable, nil),
		routes.RegisterRoute(http.MethodDelete, "loottable/:id", e.deleteLoottable, nil),
		routes.RegisterRoute(http.MethodPatch, "loottable/:id", e.updateLoottable, nil),
		routes.RegisterRoute(http.MethodPost, "loottables/bulk", e.getLoottablesBulk, nil),
	}
}

// listLoottables godoc
// @Id listLoottables
// @Summary Lists Loottables
// @Accept json
// @Produce json
// @Tags Loottable
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LoottableEntries<br>LoottableEntries.Lootdrop<br>LoottableEntries.Lootdrop.LootdropEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Merchantlists<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcEmotes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcTypesTint<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>LoottableEntries.Lootdrop.LoottableEntries<br>LoottableEntries.Loottable<br>NpcTypes<br>NpcTypes.AlternateCurrency<br>NpcTypes.AlternateCurrency.Item<br>NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>NpcTypes.AlternateCurrency.Item.Doors<br>NpcTypes.AlternateCurrency.Item.Doors.Item<br>NpcTypes.AlternateCurrency.Item.Fishings<br>NpcTypes.AlternateCurrency.Item.Fishings.Item<br>NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>NpcTypes.AlternateCurrency.Item.Forages<br>NpcTypes.AlternateCurrency.Item.Forages.Item<br>NpcTypes.AlternateCurrency.Item.Forages.Zone<br>NpcTypes.AlternateCurrency.Item.GroundSpawns<br>NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>NpcTypes.AlternateCurrency.Item.ItemTicks<br>NpcTypes.AlternateCurrency.Item.Keyrings<br>NpcTypes.AlternateCurrency.Item.LootdropEntries<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Item<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.AlternateCurrency.Item.Merchantlists<br>NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>NpcTypes.AlternateCurrency.Item.ObjectContents<br>NpcTypes.AlternateCurrency.Item.Objects<br>NpcTypes.AlternateCurrency.Item.Objects.Item<br>NpcTypes.AlternateCurrency.Item.Objects.Zone<br>NpcTypes.AlternateCurrency.Item.StartingItems<br>NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.AlternateCurrency.Item.TributeLevels<br>NpcTypes.Loottable<br>NpcTypes.Merchantlists<br>NpcTypes.Merchantlists.Items<br>NpcTypes.Merchantlists.Items.AlternateCurrencies<br>NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>NpcTypes.Merchantlists.Items.DiscoveredItems<br>NpcTypes.Merchantlists.Items.Doors<br>NpcTypes.Merchantlists.Items.Doors.Item<br>NpcTypes.Merchantlists.Items.Fishings<br>NpcTypes.Merchantlists.Items.Fishings.Item<br>NpcTypes.Merchantlists.Items.Fishings.NpcType<br>NpcTypes.Merchantlists.Items.Fishings.Zone<br>NpcTypes.Merchantlists.Items.Forages<br>NpcTypes.Merchantlists.Items.Forages.Item<br>NpcTypes.Merchantlists.Items.Forages.Zone<br>NpcTypes.Merchantlists.Items.GroundSpawns<br>NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>NpcTypes.Merchantlists.Items.ItemTicks<br>NpcTypes.Merchantlists.Items.Keyrings<br>NpcTypes.Merchantlists.Items.LootdropEntries<br>NpcTypes.Merchantlists.Items.LootdropEntries.Item<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.Merchantlists.Items.Merchantlists<br>NpcTypes.Merchantlists.Items.ObjectContents<br>NpcTypes.Merchantlists.Items.Objects<br>NpcTypes.Merchantlists.Items.Objects.Item<br>NpcTypes.Merchantlists.Items.Objects.Zone<br>NpcTypes.Merchantlists.Items.StartingItems<br>NpcTypes.Merchantlists.Items.StartingItems.Item<br>NpcTypes.Merchantlists.Items.StartingItems.Zone<br>NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.Merchantlists.Items.TributeLevels<br>NpcTypes.Merchantlists.NpcTypes<br>NpcTypes.NpcEmotes<br>NpcTypes.NpcFactions<br>NpcTypes.NpcFactions.NpcFactionEntries<br>NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcTypes.NpcSpell<br>NpcTypes.NpcSpell.NpcSpell<br>NpcTypes.NpcSpell.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcTypes.NpcTypesTint<br>NpcTypes.Spawnentries<br>NpcTypes.Spawnentries.NpcType<br>NpcTypes.Spawnentries.Spawngroup<br>NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Loottable
// @Failure 500 {string} string "Bad query request"
// @Router /loottables [get]
func (e *LoottableController) listLoottables(c echo.Context) error {
	var results []models.Loottable
	err := e.db.QueryContext(models.Loottable{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLoottable godoc
// @Id getLoottable
// @Summary Gets Loottable
// @Accept json
// @Produce json
// @Tags Loottable
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LoottableEntries<br>LoottableEntries.Lootdrop<br>LoottableEntries.Lootdrop.LootdropEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies<br>LoottableEntries.Lootdrop.LootdropEntries.Item.AlternateCurrencies.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.CharacterCorpseItems<br>LoottableEntries.Lootdrop.LootdropEntries.Item.DiscoveredItems<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Doors<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Doors.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Loottable<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Fishings.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Forages<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Forages.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns<br>LoottableEntries.Lootdrop.LootdropEntries.Item.GroundSpawns.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.ItemTicks<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Keyrings<br>LoottableEntries.Lootdrop.LootdropEntries.Item.LootdropEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Loottable<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Merchantlists<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcEmotes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.NpcTypesTint<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.Lootdrop.LootdropEntries.Item.ObjectContents<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Objects<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.Objects.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems<br>LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Item<br>LoottableEntries.Lootdrop.LootdropEntries.Item.StartingItems.Zone<br>LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries<br>LoottableEntries.Lootdrop.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>LoottableEntries.Lootdrop.LootdropEntries.Item.TributeLevels<br>LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>LoottableEntries.Lootdrop.LoottableEntries<br>LoottableEntries.Loottable<br>NpcTypes<br>NpcTypes.AlternateCurrency<br>NpcTypes.AlternateCurrency.Item<br>NpcTypes.AlternateCurrency.Item.AlternateCurrencies<br>NpcTypes.AlternateCurrency.Item.CharacterCorpseItems<br>NpcTypes.AlternateCurrency.Item.DiscoveredItems<br>NpcTypes.AlternateCurrency.Item.Doors<br>NpcTypes.AlternateCurrency.Item.Doors.Item<br>NpcTypes.AlternateCurrency.Item.Fishings<br>NpcTypes.AlternateCurrency.Item.Fishings.Item<br>NpcTypes.AlternateCurrency.Item.Fishings.NpcType<br>NpcTypes.AlternateCurrency.Item.Fishings.Zone<br>NpcTypes.AlternateCurrency.Item.Forages<br>NpcTypes.AlternateCurrency.Item.Forages.Item<br>NpcTypes.AlternateCurrency.Item.Forages.Zone<br>NpcTypes.AlternateCurrency.Item.GroundSpawns<br>NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone<br>NpcTypes.AlternateCurrency.Item.ItemTicks<br>NpcTypes.AlternateCurrency.Item.Keyrings<br>NpcTypes.AlternateCurrency.Item.LootdropEntries<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Item<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.AlternateCurrency.Item.Merchantlists<br>NpcTypes.AlternateCurrency.Item.Merchantlists.Items<br>NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes<br>NpcTypes.AlternateCurrency.Item.ObjectContents<br>NpcTypes.AlternateCurrency.Item.Objects<br>NpcTypes.AlternateCurrency.Item.Objects.Item<br>NpcTypes.AlternateCurrency.Item.Objects.Zone<br>NpcTypes.AlternateCurrency.Item.StartingItems<br>NpcTypes.AlternateCurrency.Item.StartingItems.Item<br>NpcTypes.AlternateCurrency.Item.StartingItems.Zone<br>NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries<br>NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.AlternateCurrency.Item.TributeLevels<br>NpcTypes.Loottable<br>NpcTypes.Merchantlists<br>NpcTypes.Merchantlists.Items<br>NpcTypes.Merchantlists.Items.AlternateCurrencies<br>NpcTypes.Merchantlists.Items.AlternateCurrencies.Item<br>NpcTypes.Merchantlists.Items.CharacterCorpseItems<br>NpcTypes.Merchantlists.Items.DiscoveredItems<br>NpcTypes.Merchantlists.Items.Doors<br>NpcTypes.Merchantlists.Items.Doors.Item<br>NpcTypes.Merchantlists.Items.Fishings<br>NpcTypes.Merchantlists.Items.Fishings.Item<br>NpcTypes.Merchantlists.Items.Fishings.NpcType<br>NpcTypes.Merchantlists.Items.Fishings.Zone<br>NpcTypes.Merchantlists.Items.Forages<br>NpcTypes.Merchantlists.Items.Forages.Item<br>NpcTypes.Merchantlists.Items.Forages.Zone<br>NpcTypes.Merchantlists.Items.GroundSpawns<br>NpcTypes.Merchantlists.Items.GroundSpawns.Zone<br>NpcTypes.Merchantlists.Items.ItemTicks<br>NpcTypes.Merchantlists.Items.Keyrings<br>NpcTypes.Merchantlists.Items.LootdropEntries<br>NpcTypes.Merchantlists.Items.LootdropEntries.Item<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.Merchantlists.Items.Merchantlists<br>NpcTypes.Merchantlists.Items.ObjectContents<br>NpcTypes.Merchantlists.Items.Objects<br>NpcTypes.Merchantlists.Items.Objects.Item<br>NpcTypes.Merchantlists.Items.Objects.Zone<br>NpcTypes.Merchantlists.Items.StartingItems<br>NpcTypes.Merchantlists.Items.StartingItems.Item<br>NpcTypes.Merchantlists.Items.StartingItems.Zone<br>NpcTypes.Merchantlists.Items.TradeskillRecipeEntries<br>NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.Merchantlists.Items.TributeLevels<br>NpcTypes.Merchantlists.NpcTypes<br>NpcTypes.NpcEmotes<br>NpcTypes.NpcFactions<br>NpcTypes.NpcFactions.NpcFactionEntries<br>NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>NpcTypes.NpcSpell<br>NpcTypes.NpcSpell.NpcSpell<br>NpcTypes.NpcSpell.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>NpcTypes.NpcTypesTint<br>NpcTypes.Spawnentries<br>NpcTypes.Spawnentries.NpcType<br>NpcTypes.Spawnentries.Spawngroup<br>NpcTypes.Spawnentries.Spawngroup.Spawn2<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Loottable
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /loottable/{id} [get]
func (e *LoottableController) getLoottable(c echo.Context) error {
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
	var result models.Loottable
	query := e.db.QueryContext(models.Loottable{}, c)
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

// updateLoottable godoc
// @Id updateLoottable
// @Summary Updates Loottable
// @Accept json
// @Produce json
// @Tags Loottable
// @Param id path int true "Id"
// @Param loottable body models.Loottable true "Loottable"
// @Success 200 {array} models.Loottable
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /loottable/{id} [patch]
func (e *LoottableController) updateLoottable(c echo.Context) error {
	request := new(models.Loottable)
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
	var result models.Loottable
	query := e.db.QueryContext(models.Loottable{}, c)
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

// createLoottable godoc
// @Id createLoottable
// @Summary Creates Loottable
// @Accept json
// @Produce json
// @Param loottable body models.Loottable true "Loottable"
// @Tags Loottable
// @Success 200 {array} models.Loottable
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /loottable [put]
func (e *LoottableController) createLoottable(c echo.Context) error {
	loottable := new(models.Loottable)
	if err := c.Bind(loottable); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Loottable{}, c).Model(&models.Loottable{}).Create(&loottable).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, loottable)
}

// deleteLoottable godoc
// @Id deleteLoottable
// @Summary Deletes Loottable
// @Accept json
// @Produce json
// @Tags Loottable
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /loottable/{id} [delete]
func (e *LoottableController) deleteLoottable(c echo.Context) error {
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
	var result models.Loottable
	query := e.db.QueryContext(models.Loottable{}, c)
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

// getLoottablesBulk godoc
// @Id getLoottablesBulk
// @Summary Gets Loottables in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Loottable
// @Success 200 {array} models.Loottable
// @Failure 500 {string} string "Bad query request"
// @Router /loottables/bulk [post]
func (e *LoottableController) getLoottablesBulk(c echo.Context) error {
	var results []models.Loottable

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

	err := e.db.QueryContext(models.Loottable{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
