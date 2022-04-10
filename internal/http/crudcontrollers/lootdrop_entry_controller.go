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

type LootdropEntryController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewLootdropEntryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *LootdropEntryController {
	return &LootdropEntryController{
		db:	 db,
		logger: logger,
	}
}

func (e *LootdropEntryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "lootdrop_entry/:lootdropId", e.getLootdropEntry, nil),
		routes.RegisterRoute(http.MethodGet, "lootdrop_entries", e.listLootdropEntries, nil),
		routes.RegisterRoute(http.MethodPut, "lootdrop_entry", e.createLootdropEntry, nil),
		routes.RegisterRoute(http.MethodDelete, "lootdrop_entry/:lootdropId", e.deleteLootdropEntry, nil),
		routes.RegisterRoute(http.MethodPatch, "lootdrop_entry/:lootdropId", e.updateLootdropEntry, nil),
		routes.RegisterRoute(http.MethodPost, "lootdrop_entries/bulk", e.getLootdropEntriesBulk, nil),
	}
}

// listLootdropEntries godoc
// @Id listLootdropEntries
// @Summary Lists LootdropEntries
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.NpcType<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcSpells<br>Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.Merchantlists<br>Item.Merchantlists.NpcType<br>Item.Merchantlists.NpcType.AlternateCurrency<br>Item.Merchantlists.NpcType.Loottable<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcType.Loottable.NpcTypes<br>Item.Merchantlists.NpcType.Merchantlists<br>Item.Merchantlists.NpcType.NpcEmotes<br>Item.Merchantlists.NpcType.NpcFactions<br>Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcType.NpcSpells<br>Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Merchantlists.NpcType.NpcTypesTint<br>Item.Merchantlists.NpcType.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.Tasks<br>Item.Tasks.TaskActivities<br>Item.Tasks.TaskActivities.Goallists<br>Item.Tasks.TaskActivities.NpcType<br>Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Item.Tasks.TaskActivities.NpcType.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>Lootdrop<br>Lootdrop.LootdropEntries<br>Lootdrop.LoottableEntries<br>Lootdrop.LoottableEntries.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable<br>Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entries [get]
func (e *LootdropEntryController) listLootdropEntries(c echo.Context) error {
	var results []models.LootdropEntry
	err := e.db.QueryContext(models.LootdropEntry{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLootdropEntry godoc
// @Id getLootdropEntry
// @Summary Gets LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Item<br>Item.AlternateCurrencies<br>Item.CharacterCorpseItems<br>Item.DiscoveredItems<br>Item.Doors<br>Item.Doors.Item<br>Item.Fishings<br>Item.Fishings.Item<br>Item.Fishings.NpcType<br>Item.Fishings.NpcType.AlternateCurrency<br>Item.Fishings.NpcType.Loottable<br>Item.Fishings.NpcType.Loottable.LoottableEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Fishings.NpcType.Loottable.NpcTypes<br>Item.Fishings.NpcType.Merchantlists<br>Item.Fishings.NpcType.Merchantlists.NpcType<br>Item.Fishings.NpcType.NpcEmotes<br>Item.Fishings.NpcType.NpcFactions<br>Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Item.Fishings.NpcType.NpcSpells<br>Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Fishings.NpcType.NpcTypesTint<br>Item.Fishings.NpcType.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.NpcType<br>Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Fishings.Zone<br>Item.Forages<br>Item.Forages.Item<br>Item.Forages.Zone<br>Item.GroundSpawns<br>Item.GroundSpawns.Zone<br>Item.ItemTicks<br>Item.Keyrings<br>Item.LootdropEntries<br>Item.Merchantlists<br>Item.Merchantlists.NpcType<br>Item.Merchantlists.NpcType.AlternateCurrency<br>Item.Merchantlists.NpcType.Loottable<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Merchantlists.NpcType.Loottable.NpcTypes<br>Item.Merchantlists.NpcType.Merchantlists<br>Item.Merchantlists.NpcType.NpcEmotes<br>Item.Merchantlists.NpcType.NpcFactions<br>Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Item.Merchantlists.NpcType.NpcSpells<br>Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Merchantlists.NpcType.NpcTypesTint<br>Item.Merchantlists.NpcType.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.NpcType<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.ObjectContents<br>Item.Objects<br>Item.Objects.Item<br>Item.Objects.Zone<br>Item.StartingItems<br>Item.StartingItems.Item<br>Item.StartingItems.Zone<br>Item.Tasks<br>Item.Tasks.TaskActivities<br>Item.Tasks.TaskActivities.Goallists<br>Item.Tasks.TaskActivities.NpcType<br>Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Item.Tasks.TaskActivities.NpcType.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Item.Tasks.TaskActivities.NpcType.NpcSpells<br>Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Item.Tasks.Tasksets<br>Item.TradeskillRecipeEntries<br>Item.TradeskillRecipeEntries.TradeskillRecipe<br>Item.TributeLevels<br>Lootdrop<br>Lootdrop.LootdropEntries<br>Lootdrop.LoottableEntries<br>Lootdrop.LoottableEntries.LootdropEntries<br>Lootdrop.LoottableEntries.Loottable<br>Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LootdropEntry
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entry/{id} [get]
func (e *LootdropEntryController) getLootdropEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LootdropId]"})
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.LootdropId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLootdropEntry godoc
// @Id updateLootdropEntry
// @Summary Updates LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "Id"
// @Param lootdrop_entry body models.LootdropEntry true "LootdropEntry"
// @Success 200 {array} models.LootdropEntry
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /lootdrop_entry/{id} [patch]
func (e *LootdropEntryController) updateLootdropEntry(c echo.Context) error {
	request := new(models.LootdropEntry)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LootdropId]"})
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.LootdropEntry{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createLootdropEntry godoc
// @Id createLootdropEntry
// @Summary Creates LootdropEntry
// @Accept json
// @Produce json
// @Param lootdrop_entry body models.LootdropEntry true "LootdropEntry"
// @Tags LootdropEntry
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /lootdrop_entry [put]
func (e *LootdropEntryController) createLootdropEntry(c echo.Context) error {
	lootdropEntry := new(models.LootdropEntry)
	if err := c.Bind(lootdropEntry); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LootdropEntry{}, c).Model(&models.LootdropEntry{}).Create(&lootdropEntry).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, lootdropEntry)
}

// deleteLootdropEntry godoc
// @Id deleteLootdropEntry
// @Summary Deletes LootdropEntry
// @Accept json
// @Produce json
// @Tags LootdropEntry
// @Param id path int true "lootdropId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /lootdrop_entry/{id} [delete]
func (e *LootdropEntryController) deleteLootdropEntry(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	lootdropId, err := strconv.Atoi(c.Param("lootdropId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, lootdropId)
	keys = append(keys, "lootdrop_id = ?")

	// key param [item_id] position [2] type [int]
	if len(c.QueryParam("item_id")) > 0 {
		itemIdParam, err := strconv.Atoi(c.QueryParam("item_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [item_id] err [%s]", err.Error())})
		}

		params = append(params, itemIdParam)
		keys = append(keys, "item_id = ?")
	}

	// query builder
	var result models.LootdropEntry
	query := e.db.QueryContext(models.LootdropEntry{}, c)
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

// getLootdropEntriesBulk godoc
// @Id getLootdropEntriesBulk
// @Summary Gets LootdropEntries in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LootdropEntry
// @Success 200 {array} models.LootdropEntry
// @Failure 500 {string} string "Bad query request"
// @Router /lootdrop_entries/bulk [post]
func (e *LootdropEntryController) getLootdropEntriesBulk(c echo.Context) error {
	var results []models.LootdropEntry

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

	err := e.db.QueryContext(models.LootdropEntry{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
