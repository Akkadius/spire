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

type LootdropController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLootdropController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LootdropController {
	return &LootdropController{
		db:	 db,
		logger: logger,
	}
}

func (e *LootdropController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "lootdrop/:id", e.getLootdrop, nil),
		routes.RegisterRoute(http.MethodGet, "lootdrops", e.listLootdrops, nil),
		routes.RegisterRoute(http.MethodPut, "lootdrop", e.createLootdrop, nil),
		routes.RegisterRoute(http.MethodDelete, "lootdrop/:id", e.deleteLootdrop, nil),
		routes.RegisterRoute(http.MethodPatch, "lootdrop/:id", e.updateLootdrop, nil),
		routes.RegisterRoute(http.MethodPost, "lootdrops/bulk", e.getLootdropsBulk, nil),
	}
}

// listLootdrops godoc
// @Id listLootdrops
// @Summary Lists Lootdrops
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Item.AlternateCurrencies<br>LootdropEntries.Item.CharacterCorpseItems<br>LootdropEntries.Item.DiscoveredItems<br>LootdropEntries.Item.Doors<br>LootdropEntries.Item.Doors.Item<br>LootdropEntries.Item.Fishings<br>LootdropEntries.Item.Fishings.Item<br>LootdropEntries.Item.Fishings.NpcType<br>LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>LootdropEntries.Item.Fishings.NpcType.Loottable<br>LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries<br>LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes<br>LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType<br>LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>LootdropEntries.Item.Fishings.NpcType.NpcSpells<br>LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LootdropEntries.Item.Fishings.Zone<br>LootdropEntries.Item.Forages<br>LootdropEntries.Item.Forages.Item<br>LootdropEntries.Item.Forages.Zone<br>LootdropEntries.Item.GroundSpawns<br>LootdropEntries.Item.GroundSpawns.Zone<br>LootdropEntries.Item.ItemTicks<br>LootdropEntries.Item.Keyrings<br>LootdropEntries.Item.LootdropEntries<br>LootdropEntries.Item.Merchantlists<br>LootdropEntries.Item.Merchantlists.NpcType<br>LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.NpcTypes<br>LootdropEntries.Item.Merchantlists.NpcType.Merchantlists<br>LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes<br>LootdropEntries.Item.Merchantlists.NpcType.NpcFactions<br>LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>LootdropEntries.Item.Merchantlists.NpcType.NpcSpells<br>LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LootdropEntries.Item.ObjectContents<br>LootdropEntries.Item.Objects<br>LootdropEntries.Item.Objects.Item<br>LootdropEntries.Item.Objects.Zone<br>LootdropEntries.Item.StartingItems<br>LootdropEntries.Item.StartingItems.Item<br>LootdropEntries.Item.StartingItems.Zone<br>LootdropEntries.Item.Tasks<br>LootdropEntries.Item.Tasks.TaskActivities<br>LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LootdropEntries.Item.Tasks.Tasksets<br>LootdropEntries.Item.TradeskillRecipeEntries<br>LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>LootdropEntries.Item.TributeLevels<br>LootdropEntries.Lootdrop<br>LoottableEntries<br>LoottableEntries.LootdropEntries<br>LoottableEntries.LootdropEntries.Item<br>LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>LoottableEntries.LootdropEntries.Item.Doors<br>LoottableEntries.LootdropEntries.Item.Doors.Item<br>LoottableEntries.LootdropEntries.Item.Fishings<br>LoottableEntries.LootdropEntries.Item.Fishings.Item<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>LoottableEntries.LootdropEntries.Item.Forages<br>LoottableEntries.LootdropEntries.Item.Forages.Item<br>LoottableEntries.LootdropEntries.Item.Forages.Zone<br>LoottableEntries.LootdropEntries.Item.GroundSpawns<br>LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>LoottableEntries.LootdropEntries.Item.ItemTicks<br>LoottableEntries.LootdropEntries.Item.Keyrings<br>LoottableEntries.LootdropEntries.Item.LootdropEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable.NpcTypes<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.LootdropEntries.Item.ObjectContents<br>LoottableEntries.LootdropEntries.Item.Objects<br>LoottableEntries.LootdropEntries.Item.Objects.Item<br>LoottableEntries.LootdropEntries.Item.Objects.Zone<br>LoottableEntries.LootdropEntries.Item.StartingItems<br>LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>LoottableEntries.LootdropEntries.Item.Tasks<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>LoottableEntries.LootdropEntries.Item.TributeLevels<br>LoottableEntries.LootdropEntries.Lootdrop<br>LoottableEntries.Loottable<br>LoottableEntries.Loottable.LoottableEntries<br>LoottableEntries.Loottable.NpcTypes<br>LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>LoottableEntries.Loottable.NpcTypes.Loottable<br>LoottableEntries.Loottable.NpcTypes.Merchantlists<br>LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>LoottableEntries.Loottable.NpcTypes.NpcFactions<br>LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>LoottableEntries.Loottable.NpcTypes.NpcSpells<br>LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>LoottableEntries.Loottable.NpcTypes.Spawnentries<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Lootdrop
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrops [get]
func (e *LootdropController) listLootdrops(c echo.Context) error {
	var results []models.Lootdrop
	err := e.db.QueryContext(models.Lootdrop{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLootdrop godoc
// @Id getLootdrop
// @Summary Gets Lootdrop
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Item.AlternateCurrencies<br>LootdropEntries.Item.CharacterCorpseItems<br>LootdropEntries.Item.DiscoveredItems<br>LootdropEntries.Item.Doors<br>LootdropEntries.Item.Doors.Item<br>LootdropEntries.Item.Fishings<br>LootdropEntries.Item.Fishings.Item<br>LootdropEntries.Item.Fishings.NpcType<br>LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>LootdropEntries.Item.Fishings.NpcType.Loottable<br>LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries<br>LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes<br>LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType<br>LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>LootdropEntries.Item.Fishings.NpcType.NpcSpells<br>LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LootdropEntries.Item.Fishings.Zone<br>LootdropEntries.Item.Forages<br>LootdropEntries.Item.Forages.Item<br>LootdropEntries.Item.Forages.Zone<br>LootdropEntries.Item.GroundSpawns<br>LootdropEntries.Item.GroundSpawns.Zone<br>LootdropEntries.Item.ItemTicks<br>LootdropEntries.Item.Keyrings<br>LootdropEntries.Item.LootdropEntries<br>LootdropEntries.Item.Merchantlists<br>LootdropEntries.Item.Merchantlists.NpcType<br>LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>LootdropEntries.Item.Merchantlists.NpcType.Loottable.NpcTypes<br>LootdropEntries.Item.Merchantlists.NpcType.Merchantlists<br>LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes<br>LootdropEntries.Item.Merchantlists.NpcType.NpcFactions<br>LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>LootdropEntries.Item.Merchantlists.NpcType.NpcSpells<br>LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LootdropEntries.Item.ObjectContents<br>LootdropEntries.Item.Objects<br>LootdropEntries.Item.Objects.Item<br>LootdropEntries.Item.Objects.Zone<br>LootdropEntries.Item.StartingItems<br>LootdropEntries.Item.StartingItems.Item<br>LootdropEntries.Item.StartingItems.Zone<br>LootdropEntries.Item.Tasks<br>LootdropEntries.Item.Tasks.TaskActivities<br>LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LootdropEntries.Item.Tasks.Tasksets<br>LootdropEntries.Item.TradeskillRecipeEntries<br>LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>LootdropEntries.Item.TributeLevels<br>LootdropEntries.Lootdrop<br>LoottableEntries<br>LoottableEntries.LootdropEntries<br>LoottableEntries.LootdropEntries.Item<br>LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>LoottableEntries.LootdropEntries.Item.Doors<br>LoottableEntries.LootdropEntries.Item.Doors.Item<br>LoottableEntries.LootdropEntries.Item.Fishings<br>LoottableEntries.LootdropEntries.Item.Fishings.Item<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcEmotes<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>LoottableEntries.LootdropEntries.Item.Forages<br>LoottableEntries.LootdropEntries.Item.Forages.Item<br>LoottableEntries.LootdropEntries.Item.Forages.Zone<br>LoottableEntries.LootdropEntries.Item.GroundSpawns<br>LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>LoottableEntries.LootdropEntries.Item.ItemTicks<br>LoottableEntries.LootdropEntries.Item.Keyrings<br>LoottableEntries.LootdropEntries.Item.LootdropEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable.NpcTypes<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.LootdropEntries.Item.ObjectContents<br>LoottableEntries.LootdropEntries.Item.Objects<br>LoottableEntries.LootdropEntries.Item.Objects.Item<br>LoottableEntries.LootdropEntries.Item.Objects.Zone<br>LoottableEntries.LootdropEntries.Item.StartingItems<br>LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>LoottableEntries.LootdropEntries.Item.Tasks<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>LoottableEntries.LootdropEntries.Item.TributeLevels<br>LoottableEntries.LootdropEntries.Lootdrop<br>LoottableEntries.Loottable<br>LoottableEntries.Loottable.LoottableEntries<br>LoottableEntries.Loottable.NpcTypes<br>LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>LoottableEntries.Loottable.NpcTypes.Loottable<br>LoottableEntries.Loottable.NpcTypes.Merchantlists<br>LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>LoottableEntries.Loottable.NpcTypes.NpcFactions<br>LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>LoottableEntries.Loottable.NpcTypes.NpcSpells<br>LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>LoottableEntries.Loottable.NpcTypes.Spawnentries<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Lootdrop
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop/{id} [get]
func (e *LootdropController) getLootdrop(c echo.Context) error {
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
	var result models.Lootdrop
	query := e.db.QueryContext(models.Lootdrop{}, c)
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

// updateLootdrop godoc
// @Id updateLootdrop
// @Summary Updates Lootdrop
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param id path int true "Id"
// @Param lootdrop body models.Lootdrop true "Lootdrop"
// @Success 200 {array} models.Lootdrop
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /lootdrop/{id} [patch]
func (e *LootdropController) updateLootdrop(c echo.Context) error {
	request := new(models.Lootdrop)
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
	var result models.Lootdrop
	query := e.db.QueryContext(models.Lootdrop{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Lootdrop{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createLootdrop godoc
// @Id createLootdrop
// @Summary Creates Lootdrop
// @Accept json
// @Produce json
// @Param lootdrop body models.Lootdrop true "Lootdrop"
// @Tags Lootdrop
// @Success 200 {array} models.Lootdrop
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /lootdrop [put]
func (e *LootdropController) createLootdrop(c echo.Context) error {
	lootdrop := new(models.Lootdrop)
	if err := c.Bind(lootdrop); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Lootdrop{}, c).Model(&models.Lootdrop{}).Create(&lootdrop).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, lootdrop)
}

// deleteLootdrop godoc
// @Id deleteLootdrop
// @Summary Deletes Lootdrop
// @Accept json
// @Produce json
// @Tags Lootdrop
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /lootdrop/{id} [delete]
func (e *LootdropController) deleteLootdrop(c echo.Context) error {
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
	var result models.Lootdrop
	query := e.db.QueryContext(models.Lootdrop{}, c)
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

// getLootdropsBulk godoc
// @Id getLootdropsBulk
// @Summary Gets Lootdrops in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Lootdrop
// @Success 200 {array} models.Lootdrop
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrops/bulk [post]
func (e *LootdropController) getLootdropsBulk(c echo.Context) error {
	var results []models.Lootdrop

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

	err := e.db.QueryContext(models.Lootdrop{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
