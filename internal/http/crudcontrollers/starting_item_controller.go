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

type StartingItemController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewStartingItemController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *StartingItemController {
	return &StartingItemController{
		db:	 db,
		logger: logger,
	}
}

func (e *StartingItemController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "starting_item/:id", e.getStartingItem, nil),
		routes.RegisterRoute(http.MethodGet, "starting_items", e.listStartingItems, nil),
		routes.RegisterRoute(http.MethodPut, "starting_item", e.createStartingItem, nil),
		routes.RegisterRoute(http.MethodDelete, "starting_item/:id", e.deleteStartingItem, nil),
		routes.RegisterRoute(http.MethodPatch, "starting_item/:id", e.updateStartingItem, nil),
		routes.RegisterRoute(http.MethodPost, "starting_items/bulk", e.getStartingItemsBulk, nil),
	}
}

// listStartingItems godoc
// @Id listStartingItems
// @Summary Lists StartingItems
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.NpcType<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcSpells<br>Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.NpcType<br>Item.Merchantlists.NpcType.AlternateCurrency<br>Item.Merchantlists.NpcType.Loottable<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcType.Loottable.NpcTypes<br>Item.Merchantlists.NpcType.Merchantlists<br>Item.Merchantlists.NpcType.NpcEmotes<br>Item.Merchantlists.NpcType.NpcFactions<br>Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcType.NpcSpells<br>Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Merchantlists.NpcType.NpcTypesTint<br>Item.Merchantlists.NpcType.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.Tasks<br>Item.Tasks.TaskActivities<br>Item.Tasks.TaskActivities.Goallists<br>Item.Tasks.TaskActivities.NpcType<br>Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Item.Tasks.TaskActivities.NpcType.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>Zone"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.StartingItem
// @Failure 500 {string} string "Bad query request"
// @Router /starting_items [get]
func (e *StartingItemController) listStartingItems(c echo.Context) error {
	var results []models.StartingItem
	err := e.db.QueryContext(models.StartingItem{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getStartingItem godoc
// @Id getStartingItem
// @Summary Gets StartingItem
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.NpcType<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcSpells<br>Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.LootdropEntries.Item<br>Item.LootdropEntries.Lootdrop<br>Item.LootdropEntries.Lootdrop.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Merchantlists<br>Item.Merchantlists.NpcType<br>Item.Merchantlists.NpcType.AlternateCurrency<br>Item.Merchantlists.NpcType.Loottable<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcType.Loottable.NpcTypes<br>Item.Merchantlists.NpcType.Merchantlists<br>Item.Merchantlists.NpcType.NpcEmotes<br>Item.Merchantlists.NpcType.NpcFactions<br>Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcType.NpcSpells<br>Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Merchantlists.NpcType.NpcTypesTint<br>Item.Merchantlists.NpcType.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.Tasks<br>Item.Tasks.TaskActivities<br>Item.Tasks.TaskActivities.Goallists<br>Item.Tasks.TaskActivities.NpcType<br>Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Item.Tasks.TaskActivities.NpcType.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>Zone"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.StartingItem
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /starting_item/{id} [get]
func (e *StartingItemController) getStartingItem(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [race] position [2] type [int]
	if len(c.QueryParam("race")) > 0 {
		raceParam, err := strconv.Atoi(c.QueryParam("race"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [race] err [%s]", err.Error())})
		}

		params = append(params, raceParam)
		keys = append(keys, "race = ?")
	}

	// query builder
	var result models.StartingItem
	query := e.db.QueryContext(models.StartingItem{}, c)
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

// updateStartingItem godoc
// @Id updateStartingItem
// @Summary Updates StartingItem
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param id path int true "Id"
// @Param starting_item body models.StartingItem true "StartingItem"
// @Success 200 {array} models.StartingItem
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /starting_item/{id} [patch]
func (e *StartingItemController) updateStartingItem(c echo.Context) error {
	request := new(models.StartingItem)
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

	// key param [race] position [2] type [int]
	if len(c.QueryParam("race")) > 0 {
		raceParam, err := strconv.Atoi(c.QueryParam("race"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [race] err [%s]", err.Error())})
		}

		params = append(params, raceParam)
		keys = append(keys, "race = ?")
	}

	// query builder
	var result models.StartingItem
	query := e.db.QueryContext(models.StartingItem{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.StartingItem{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createStartingItem godoc
// @Id createStartingItem
// @Summary Creates StartingItem
// @Accept json
// @Produce json
// @Param starting_item body models.StartingItem true "StartingItem"
// @Tags StartingItem
// @Success 200 {array} models.StartingItem
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /starting_item [put]
func (e *StartingItemController) createStartingItem(c echo.Context) error {
	startingItem := new(models.StartingItem)
	if err := c.Bind(startingItem); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.StartingItem{}, c).Model(&models.StartingItem{}).Create(&startingItem).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, startingItem)
}

// deleteStartingItem godoc
// @Id deleteStartingItem
// @Summary Deletes StartingItem
// @Accept json
// @Produce json
// @Tags StartingItem
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /starting_item/{id} [delete]
func (e *StartingItemController) deleteStartingItem(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [race] position [2] type [int]
	if len(c.QueryParam("race")) > 0 {
		raceParam, err := strconv.Atoi(c.QueryParam("race"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [race] err [%s]", err.Error())})
		}

		params = append(params, raceParam)
		keys = append(keys, "race = ?")
	}

	// query builder
	var result models.StartingItem
	query := e.db.QueryContext(models.StartingItem{}, c)
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

// getStartingItemsBulk godoc
// @Id getStartingItemsBulk
// @Summary Gets StartingItems in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags StartingItem
// @Success 200 {array} models.StartingItem
// @Failure 500 {string} string "Bad query request"
// @Router /starting_items/bulk [post]
func (e *StartingItemController) getStartingItemsBulk(c echo.Context) error {
	var results []models.StartingItem

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

	err := e.db.QueryContext(models.StartingItem{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
