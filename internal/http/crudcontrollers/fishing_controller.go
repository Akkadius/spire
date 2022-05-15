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

type FishingController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewFishingController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *FishingController {
	return &FishingController{
		db:	 db,
		logger: logger,
	}
}

func (e *FishingController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "fishing/:id", e.getFishing, nil),
		routes.RegisterRoute(http.MethodGet, "fishings", e.listFishings, nil),
		routes.RegisterRoute(http.MethodPut, "fishing", e.createFishing, nil),
		routes.RegisterRoute(http.MethodDelete, "fishing/:id", e.deleteFishing, nil),
		routes.RegisterRoute(http.MethodPatch, "fishing/:id", e.updateFishing, nil),
		routes.RegisterRoute(http.MethodPost, "fishings/bulk", e.getFishingsBulk, nil),
	}
}

// listFishings godoc
// @Id listFishings
// @Summary Lists Fishings
// @Accept json
// @Produce json
// @Tags Fishing
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.NpcType<br>Item.Merchantlists.NpcType.AlternateCurrency<br>Item.Merchantlists.NpcType.Loottable<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcType.Loottable.NpcTypes<br>Item.Merchantlists.NpcType.Merchantlists<br>Item.Merchantlists.NpcType.NpcEmotes<br>Item.Merchantlists.NpcType.NpcFactions<br>Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcType.NpcSpells<br>Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Merchantlists.NpcType.NpcTypesTint<br>Item.Merchantlists.NpcType.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.Tasks<br>Item.Tasks.TaskActivities<br>Item.Tasks.TaskActivities.Goallists<br>Item.Tasks.TaskActivities.NpcType<br>Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Item.Tasks.TaskActivities.NpcType.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.Merchantlists.NpcType<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcSpells<br>NpcType.NpcSpells.NpcSpellsEntries<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.NpcType<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Zone"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Bad query request"
// @Router /fishings [get]
func (e *FishingController) listFishings(c echo.Context) error {
	var results []models.Fishing
	err := e.db.QueryContext(models.Fishing{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getFishing godoc
// @Id getFishing
// @Summary Gets Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.NpcType<br>Item.Merchantlists.NpcType.AlternateCurrency<br>Item.Merchantlists.NpcType.Loottable<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcType.Loottable.NpcTypes<br>Item.Merchantlists.NpcType.Merchantlists<br>Item.Merchantlists.NpcType.NpcEmotes<br>Item.Merchantlists.NpcType.NpcFactions<br>Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcType.NpcSpells<br>Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Merchantlists.NpcType.NpcTypesTint<br>Item.Merchantlists.NpcType.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.Tasks<br>Item.Tasks.TaskActivities<br>Item.Tasks.TaskActivities.Goallists<br>Item.Tasks.TaskActivities.NpcType<br>Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Item.Tasks.TaskActivities.NpcType.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.Merchantlists.NpcType<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcSpells<br>NpcType.NpcSpells.NpcSpellsEntries<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.NpcType<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Zone"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Fishing
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /fishing/{id} [get]
func (e *FishingController) getFishing(c echo.Context) error {
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
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

// updateFishing godoc
// @Id updateFishing
// @Summary Updates Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "Id"
// @Param fishing body models.Fishing true "Fishing"
// @Success 200 {array} models.Fishing
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /fishing/{id} [patch]
func (e *FishingController) updateFishing(c echo.Context) error {
	request := new(models.Fishing)
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Fishing{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createFishing godoc
// @Id createFishing
// @Summary Creates Fishing
// @Accept json
// @Produce json
// @Param fishing body models.Fishing true "Fishing"
// @Tags Fishing
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /fishing [put]
func (e *FishingController) createFishing(c echo.Context) error {
	fishing := new(models.Fishing)
	if err := c.Bind(fishing); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Fishing{}, c).Model(&models.Fishing{}).Create(&fishing).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, fishing)
}

// deleteFishing godoc
// @Id deleteFishing
// @Summary Deletes Fishing
// @Accept json
// @Produce json
// @Tags Fishing
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /fishing/{id} [delete]
func (e *FishingController) deleteFishing(c echo.Context) error {
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
	var result models.Fishing
	query := e.db.QueryContext(models.Fishing{}, c)
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

// getFishingsBulk godoc
// @Id getFishingsBulk
// @Summary Gets Fishings in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Fishing
// @Success 200 {array} models.Fishing
// @Failure 500 {string} string "Bad query request"
// @Router /fishings/bulk [post]
func (e *FishingController) getFishingsBulk(c echo.Context) error {
	var results []models.Fishing

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

	err := e.db.QueryContext(models.Fishing{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
