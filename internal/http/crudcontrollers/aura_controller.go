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

type AuraController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewAuraController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *AuraController {
	return &AuraController{
		db:	 db,
		logger: logger,
	}
}

func (e *AuraController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "aura/:typeId", e.getAura, nil),
		routes.RegisterRoute(http.MethodGet, "auras", e.listAuras, nil),
		routes.RegisterRoute(http.MethodPut, "aura", e.createAura, nil),
		routes.RegisterRoute(http.MethodDelete, "aura/:typeId", e.deleteAura, nil),
		routes.RegisterRoute(http.MethodPatch, "aura/:typeId", e.updateAura, nil),
		routes.RegisterRoute(http.MethodPost, "auras/bulk", e.getAurasBulk, nil),
	}
}

// listAuras godoc
// @Id listAuras
// @Summary Lists Auras
// @Accept json
// @Produce json
// @Tags Aura
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>SpellsNew<br>SpellsNew.Aura<br>SpellsNew.BlockedSpells<br>SpellsNew.Damageshieldtypes<br>SpellsNew.Items<br>SpellsNew.Items.AlternateCurrencies<br>SpellsNew.Items.AlternateCurrencies.Item<br>SpellsNew.Items.CharacterCorpseItems<br>SpellsNew.Items.DiscoveredItems<br>SpellsNew.Items.Doors<br>SpellsNew.Items.Doors.Item<br>SpellsNew.Items.Fishings<br>SpellsNew.Items.Fishings.Item<br>SpellsNew.Items.Fishings.NpcType<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.Merchantlists<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.Items<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.NpcEmotes<br>SpellsNew.Items.Fishings.NpcType.NpcFactions<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Fishings.NpcType.NpcSpell<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpell<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>SpellsNew.Items.Fishings.NpcType.NpcTypesTint<br>SpellsNew.Items.Fishings.NpcType.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Fishings.Zone<br>SpellsNew.Items.Forages<br>SpellsNew.Items.Forages.Item<br>SpellsNew.Items.Forages.Zone<br>SpellsNew.Items.GroundSpawns<br>SpellsNew.Items.GroundSpawns.Zone<br>SpellsNew.Items.ItemTicks<br>SpellsNew.Items.Keyrings<br>SpellsNew.Items.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Merchantlists<br>SpellsNew.Items.Merchantlists.Items<br>SpellsNew.Items.Merchantlists.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.ObjectContents<br>SpellsNew.Items.Objects<br>SpellsNew.Items.Objects.Item<br>SpellsNew.Items.Objects.Zone<br>SpellsNew.Items.StartingItems<br>SpellsNew.Items.StartingItems.Item<br>SpellsNew.Items.StartingItems.Zone<br>SpellsNew.Items.TradeskillRecipeEntries<br>SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>SpellsNew.Items.TributeLevels<br>SpellsNew.NpcSpellsEntries<br>SpellsNew.NpcSpellsEntries.SpellsNew<br>SpellsNew.SpellBuckets<br>SpellsNew.SpellGlobals"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Aura
// @Failure 500 {string} string "Bad query request"
// @Router /auras [get]
func (e *AuraController) listAuras(c echo.Context) error {
	var results []models.Aura
	err := e.db.QueryContext(models.Aura{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getAura godoc
// @Id getAura
// @Summary Gets Aura
// @Accept json
// @Produce json
// @Tags Aura
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>SpellsNew<br>SpellsNew.Aura<br>SpellsNew.BlockedSpells<br>SpellsNew.Damageshieldtypes<br>SpellsNew.Items<br>SpellsNew.Items.AlternateCurrencies<br>SpellsNew.Items.AlternateCurrencies.Item<br>SpellsNew.Items.CharacterCorpseItems<br>SpellsNew.Items.DiscoveredItems<br>SpellsNew.Items.Doors<br>SpellsNew.Items.Doors.Item<br>SpellsNew.Items.Fishings<br>SpellsNew.Items.Fishings.Item<br>SpellsNew.Items.Fishings.NpcType<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency<br>SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.Merchantlists<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.Items<br>SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes<br>SpellsNew.Items.Fishings.NpcType.NpcEmotes<br>SpellsNew.Items.Fishings.NpcType.NpcFactions<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Fishings.NpcType.NpcSpell<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpell<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>SpellsNew.Items.Fishings.NpcType.NpcTypesTint<br>SpellsNew.Items.Fishings.NpcType.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Fishings.Zone<br>SpellsNew.Items.Forages<br>SpellsNew.Items.Forages.Item<br>SpellsNew.Items.Forages.Zone<br>SpellsNew.Items.GroundSpawns<br>SpellsNew.Items.GroundSpawns.Zone<br>SpellsNew.Items.ItemTicks<br>SpellsNew.Items.Keyrings<br>SpellsNew.Items.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.Merchantlists<br>SpellsNew.Items.Merchantlists.Items<br>SpellsNew.Items.Merchantlists.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency<br>SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>SpellsNew.Items.ObjectContents<br>SpellsNew.Items.Objects<br>SpellsNew.Items.Objects.Item<br>SpellsNew.Items.Objects.Zone<br>SpellsNew.Items.StartingItems<br>SpellsNew.Items.StartingItems.Item<br>SpellsNew.Items.StartingItems.Zone<br>SpellsNew.Items.TradeskillRecipeEntries<br>SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe<br>SpellsNew.Items.TributeLevels<br>SpellsNew.NpcSpellsEntries<br>SpellsNew.NpcSpellsEntries.SpellsNew<br>SpellsNew.SpellBuckets<br>SpellsNew.SpellGlobals"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Aura
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /aura/{id} [get]
func (e *AuraController) getAura(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	typeId, err := strconv.Atoi(c.Param("typeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TypeId]"})
	}
	params = append(params, typeId)
	keys = append(keys, "type = ?")

	// query builder
	var result models.Aura
	query := e.db.QueryContext(models.Aura{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Type == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateAura godoc
// @Id updateAura
// @Summary Updates Aura
// @Accept json
// @Produce json
// @Tags Aura
// @Param id path int true "Id"
// @Param aura body models.Aura true "Aura"
// @Success 200 {array} models.Aura
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /aura/{id} [patch]
func (e *AuraController) updateAura(c echo.Context) error {
	request := new(models.Aura)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	typeId, err := strconv.Atoi(c.Param("typeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TypeId]"})
	}
	params = append(params, typeId)
	keys = append(keys, "type = ?")

	// query builder
	var result models.Aura
	query := e.db.QueryContext(models.Aura{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Aura{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createAura godoc
// @Id createAura
// @Summary Creates Aura
// @Accept json
// @Produce json
// @Param aura body models.Aura true "Aura"
// @Tags Aura
// @Success 200 {array} models.Aura
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /aura [put]
func (e *AuraController) createAura(c echo.Context) error {
	aura := new(models.Aura)
	if err := c.Bind(aura); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Aura{}, c).Model(&models.Aura{}).Create(&aura).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, aura)
}

// deleteAura godoc
// @Id deleteAura
// @Summary Deletes Aura
// @Accept json
// @Produce json
// @Tags Aura
// @Param id path int true "typeId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /aura/{id} [delete]
func (e *AuraController) deleteAura(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	typeId, err := strconv.Atoi(c.Param("typeId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, typeId)
	keys = append(keys, "type = ?")

	// query builder
	var result models.Aura
	query := e.db.QueryContext(models.Aura{}, c)
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

// getAurasBulk godoc
// @Id getAurasBulk
// @Summary Gets Auras in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Aura
// @Success 200 {array} models.Aura
// @Failure 500 {string} string "Bad query request"
// @Router /auras/bulk [post]
func (e *AuraController) getAurasBulk(c echo.Context) error {
	var results []models.Aura

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

	err := e.db.QueryContext(models.Aura{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
