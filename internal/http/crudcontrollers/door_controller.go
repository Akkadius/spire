package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type DoorController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewDoorController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *DoorController {
	return &DoorController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *DoorController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "door/:id", e.getDoor, nil),
		routes.RegisterRoute(http.MethodGet, "doors", e.listDoors, nil),
		routes.RegisterRoute(http.MethodPut, "door", e.createDoor, nil),
		routes.RegisterRoute(http.MethodDelete, "door/:id", e.deleteDoor, nil),
		routes.RegisterRoute(http.MethodPatch, "door/:id", e.updateDoor, nil),
		routes.RegisterRoute(http.MethodPost, "doors/bulk", e.getDoorsBulk, nil),
	}
}

// listDoors godoc
// @Id listDoors
// @Summary Lists Doors
// @Accept json
// @Produce json
// @Tags Door
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.AlternateCurrencies.Item<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.AlternateCurrency.Item<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.Items<br>Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Item.Fishings.NpcType.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.Items<br>Item.Merchantlists.NpcTypes<br>Item.Merchantlists.NpcTypes.AlternateCurrency<br>Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Item.Merchantlists.NpcTypes.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Item.Merchantlists.NpcTypes.Merchantlists<br>Item.Merchantlists.NpcTypes.NpcEmotes<br>Item.Merchantlists.NpcTypes.NpcFactions<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.Merchantlists.NpcTypes.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Merchantlists.NpcTypes.NpcTypesTint<br>Item.Merchantlists.NpcTypes.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Door
// @Failure 500 {string} string "Bad query request"
// @Router /doors [get]
func (e *DoorController) listDoors(c echo.Context) error {
	var results []models.Door
	err := e.db.QueryContext(models.Door{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDoor godoc
// @Id getDoor
// @Summary Gets Door
// @Accept json
// @Produce json
// @Tags Door
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.AlternateCurrencies.Item<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.AlternateCurrency.Item<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.Items<br>Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Item.Fishings.NpcType.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpell<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.Items<br>Item.Merchantlists.NpcTypes<br>Item.Merchantlists.NpcTypes.AlternateCurrency<br>Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Item.Merchantlists.NpcTypes.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Item.Merchantlists.NpcTypes.Merchantlists<br>Item.Merchantlists.NpcTypes.NpcEmotes<br>Item.Merchantlists.NpcTypes.NpcFactions<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Item.Merchantlists.NpcTypes.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Item.Merchantlists.NpcTypes.NpcTypesTint<br>Item.Merchantlists.NpcTypes.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Door
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /door/{id} [get]
func (e *DoorController) getDoor(c echo.Context) error {
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
	var result models.Door
	query := e.db.QueryContext(models.Door{}, c)
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

// updateDoor godoc
// @Id updateDoor
// @Summary Updates Door
// @Accept json
// @Produce json
// @Tags Door
// @Param id path int true "Id"
// @Param door body models.Door true "Door"
// @Success 200 {array} models.Door
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /door/{id} [patch]
func (e *DoorController) updateDoor(c echo.Context) error {
	request := new(models.Door)
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
	var result models.Door
	query := e.db.QueryContext(models.Door{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [Door] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createDoor godoc
// @Id createDoor
// @Summary Creates Door
// @Accept json
// @Produce json
// @Param door body models.Door true "Door"
// @Tags Door
// @Success 200 {array} models.Door
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /door [put]
func (e *DoorController) createDoor(c echo.Context) error {
	door := new(models.Door)
	if err := c.Bind(door); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Door{}, c).Model(&models.Door{}).Create(&door).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.Door{}, door)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [Door] [%v] data [%v]", door.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, door)
}

// deleteDoor godoc
// @Id deleteDoor
// @Summary Deletes Door
// @Accept json
// @Produce json
// @Tags Door
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /door/{id} [delete]
func (e *DoorController) deleteDoor(c echo.Context) error {
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
	var result models.Door
	query := e.db.QueryContext(models.Door{}, c)
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

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [Door] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDoorsBulk godoc
// @Id getDoorsBulk
// @Summary Gets Doors in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Door
// @Success 200 {array} models.Door
// @Failure 500 {string} string "Bad query request"
// @Router /doors/bulk [post]
func (e *DoorController) getDoorsBulk(c echo.Context) error {
	var results []models.Door

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

	err := e.db.QueryContext(models.Door{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
