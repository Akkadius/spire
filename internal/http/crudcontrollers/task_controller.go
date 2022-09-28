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

type TaskController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewTaskController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *TaskController {
	return &TaskController{
		db:	 db,
		logger: logger,
	}
}

func (e *TaskController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "task/:id", e.getTask, nil),
		routes.RegisterRoute(http.MethodGet, "tasks", e.listTasks, nil),
		routes.RegisterRoute(http.MethodPut, "task", e.createTask, nil),
		routes.RegisterRoute(http.MethodDelete, "task/:id", e.deleteTask, nil),
		routes.RegisterRoute(http.MethodPatch, "task/:id", e.updateTask, nil),
		routes.RegisterRoute(http.MethodPost, "tasks/bulk", e.getTasksBulk, nil),
	}
}

// listTasks godoc
// @Id listTasks
// @Summary Lists Tasks
// @Accept json
// @Produce json
// @Tags Task
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrency<br>AlternateCurrency.Item<br>AlternateCurrency.Item.AlternateCurrencies<br>AlternateCurrency.Item.CharacterCorpseItems<br>AlternateCurrency.Item.DiscoveredItems<br>AlternateCurrency.Item.Doors<br>AlternateCurrency.Item.Doors.Item<br>AlternateCurrency.Item.Fishings<br>AlternateCurrency.Item.Fishings.Item<br>AlternateCurrency.Item.Fishings.NpcType<br>AlternateCurrency.Item.Fishings.NpcType.AlternateCurrency<br>AlternateCurrency.Item.Fishings.NpcType.Loottable<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.NpcTypes<br>AlternateCurrency.Item.Fishings.NpcType.Merchantlists<br>AlternateCurrency.Item.Fishings.NpcType.Merchantlists.Items<br>AlternateCurrency.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>AlternateCurrency.Item.Fishings.NpcType.NpcEmotes<br>AlternateCurrency.Item.Fishings.NpcType.NpcFactions<br>AlternateCurrency.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>AlternateCurrency.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>AlternateCurrency.Item.Fishings.NpcType.NpcTypesTint<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.NpcType<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>AlternateCurrency.Item.Fishings.Zone<br>AlternateCurrency.Item.Forages<br>AlternateCurrency.Item.Forages.Item<br>AlternateCurrency.Item.Forages.Zone<br>AlternateCurrency.Item.GroundSpawns<br>AlternateCurrency.Item.GroundSpawns.Zone<br>AlternateCurrency.Item.ItemTicks<br>AlternateCurrency.Item.Keyrings<br>AlternateCurrency.Item.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Item<br>AlternateCurrency.Item.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>AlternateCurrency.Item.Merchantlists<br>AlternateCurrency.Item.Merchantlists.Items<br>AlternateCurrency.Item.Merchantlists.NpcTypes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.AlternateCurrency<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Merchantlists<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcEmotes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcTypesTint<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>AlternateCurrency.Item.ObjectContents<br>AlternateCurrency.Item.Objects<br>AlternateCurrency.Item.Objects.Item<br>AlternateCurrency.Item.Objects.Zone<br>AlternateCurrency.Item.StartingItems<br>AlternateCurrency.Item.StartingItems.Item<br>AlternateCurrency.Item.StartingItems.Zone<br>AlternateCurrency.Item.TradeskillRecipeEntries<br>AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>AlternateCurrency.Item.TributeLevels<br>TaskActivities<br>Tasksets"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Bad query request"
// @Router /tasks [get]
func (e *TaskController) listTasks(c echo.Context) error {
	var results []models.Task
	err := e.db.QueryContext(models.Task{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTask godoc
// @Id getTask
// @Summary Gets Task
// @Accept json
// @Produce json
// @Tags Task
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrency<br>AlternateCurrency.Item<br>AlternateCurrency.Item.AlternateCurrencies<br>AlternateCurrency.Item.CharacterCorpseItems<br>AlternateCurrency.Item.DiscoveredItems<br>AlternateCurrency.Item.Doors<br>AlternateCurrency.Item.Doors.Item<br>AlternateCurrency.Item.Fishings<br>AlternateCurrency.Item.Fishings.Item<br>AlternateCurrency.Item.Fishings.NpcType<br>AlternateCurrency.Item.Fishings.NpcType.AlternateCurrency<br>AlternateCurrency.Item.Fishings.NpcType.Loottable<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>AlternateCurrency.Item.Fishings.NpcType.Loottable.NpcTypes<br>AlternateCurrency.Item.Fishings.NpcType.Merchantlists<br>AlternateCurrency.Item.Fishings.NpcType.Merchantlists.Items<br>AlternateCurrency.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>AlternateCurrency.Item.Fishings.NpcType.NpcEmotes<br>AlternateCurrency.Item.Fishings.NpcType.NpcFactions<br>AlternateCurrency.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>AlternateCurrency.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpell<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>AlternateCurrency.Item.Fishings.NpcType.NpcTypesTint<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.NpcType<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>AlternateCurrency.Item.Fishings.Zone<br>AlternateCurrency.Item.Forages<br>AlternateCurrency.Item.Forages.Item<br>AlternateCurrency.Item.Forages.Zone<br>AlternateCurrency.Item.GroundSpawns<br>AlternateCurrency.Item.GroundSpawns.Zone<br>AlternateCurrency.Item.ItemTicks<br>AlternateCurrency.Item.Keyrings<br>AlternateCurrency.Item.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Item<br>AlternateCurrency.Item.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>AlternateCurrency.Item.Merchantlists<br>AlternateCurrency.Item.Merchantlists.Items<br>AlternateCurrency.Item.Merchantlists.NpcTypes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.AlternateCurrency<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Merchantlists<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcEmotes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>AlternateCurrency.Item.Merchantlists.NpcTypes.NpcTypesTint<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>AlternateCurrency.Item.ObjectContents<br>AlternateCurrency.Item.Objects<br>AlternateCurrency.Item.Objects.Item<br>AlternateCurrency.Item.Objects.Zone<br>AlternateCurrency.Item.StartingItems<br>AlternateCurrency.Item.StartingItems.Item<br>AlternateCurrency.Item.StartingItems.Zone<br>AlternateCurrency.Item.TradeskillRecipeEntries<br>AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe<br>AlternateCurrency.Item.TributeLevels<br>TaskActivities<br>Tasksets"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Task
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /task/{id} [get]
func (e *TaskController) getTask(c echo.Context) error {
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
	var result models.Task
	query := e.db.QueryContext(models.Task{}, c)
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

// updateTask godoc
// @Id updateTask
// @Summary Updates Task
// @Accept json
// @Produce json
// @Tags Task
// @Param id path int true "Id"
// @Param task body models.Task true "Task"
// @Success 200 {array} models.Task
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /task/{id} [patch]
func (e *TaskController) updateTask(c echo.Context) error {
	request := new(models.Task)
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
	var result models.Task
	query := e.db.QueryContext(models.Task{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Task{}, c).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createTask godoc
// @Id createTask
// @Summary Creates Task
// @Accept json
// @Produce json
// @Param task body models.Task true "Task"
// @Tags Task
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /task [put]
func (e *TaskController) createTask(c echo.Context) error {
	task := new(models.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Task{}, c).Model(&models.Task{}).Create(&task).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, task)
}

// deleteTask godoc
// @Id deleteTask
// @Summary Deletes Task
// @Accept json
// @Produce json
// @Tags Task
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /task/{id} [delete]
func (e *TaskController) deleteTask(c echo.Context) error {
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
	var result models.Task
	query := e.db.QueryContext(models.Task{}, c)
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

// getTasksBulk godoc
// @Id getTasksBulk
// @Summary Gets Tasks in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Task
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Bad query request"
// @Router /tasks/bulk [post]
func (e *TaskController) getTasksBulk(c echo.Context) error {
	var results []models.Task

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

	err := e.db.QueryContext(models.Task{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
