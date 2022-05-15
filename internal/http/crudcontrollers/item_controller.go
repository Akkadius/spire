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

type ItemController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewItemController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ItemController {
	return &ItemController{
		db:	 db,
		logger: logger,
	}
}

func (e *ItemController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "item/:id", e.getItem, nil),
		routes.RegisterRoute(http.MethodGet, "items", e.listItems, nil),
		routes.RegisterRoute(http.MethodPut, "item", e.createItem, nil),
		routes.RegisterRoute(http.MethodDelete, "item/:id", e.deleteItem, nil),
		routes.RegisterRoute(http.MethodPatch, "item/:id", e.updateItem, nil),
		routes.RegisterRoute(http.MethodPost, "items/bulk", e.getItemsBulk, nil),
	}
}

// listItems godoc
// @Id listItems
// @Summary Lists Items
// @Accept json
// @Produce json
// @Tags Item
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrencies<br>CharacterCorpseItems<br>DiscoveredItems<br>Doors<br>Doors.Item<br>Fishings<br>Fishings.Item<br>Fishings.NpcType<br>Fishings.NpcType.AlternateCurrency<br>Fishings.NpcType.Loottable<br>Fishings.NpcType.Loottable.LoottableEntries<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Fishings.NpcType.Loottable.NpcTypes<br>Fishings.NpcType.Merchantlists<br>Fishings.NpcType.Merchantlists.NpcType<br>Fishings.NpcType.NpcEmotes<br>Fishings.NpcType.NpcFactions<br>Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Fishings.NpcType.NpcSpells<br>Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Fishings.NpcType.NpcTypesTint<br>Fishings.NpcType.Spawnentries<br>Fishings.NpcType.Spawnentries.NpcType<br>Fishings.NpcType.Spawnentries.Spawngroup<br>Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Fishings.Zone<br>Forages<br>Forages.Item<br>Forages.Zone<br>GroundSpawns<br>GroundSpawns.Zone<br>ItemTicks<br>Keyrings<br>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Lootdrop<br>LootdropEntries.Lootdrop.LootdropEntries<br>LootdropEntries.Lootdrop.LoottableEntries<br>LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Merchantlists<br>Merchantlists.NpcType<br>Merchantlists.NpcType.AlternateCurrency<br>Merchantlists.NpcType.Loottable<br>Merchantlists.NpcType.Loottable.LoottableEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Merchantlists.NpcType.Loottable.NpcTypes<br>Merchantlists.NpcType.Merchantlists<br>Merchantlists.NpcType.NpcEmotes<br>Merchantlists.NpcType.NpcFactions<br>Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Merchantlists.NpcType.NpcSpells<br>Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Merchantlists.NpcType.NpcTypesTint<br>Merchantlists.NpcType.Spawnentries<br>Merchantlists.NpcType.Spawnentries.NpcType<br>Merchantlists.NpcType.Spawnentries.Spawngroup<br>Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>ObjectContents<br>Objects<br>Objects.Item<br>Objects.Zone<br>StartingItems<br>StartingItems.Item<br>StartingItems.Zone<br>Tasks<br>Tasks.TaskActivities<br>Tasks.TaskActivities.Goallists<br>Tasks.TaskActivities.NpcType<br>Tasks.TaskActivities.NpcType.AlternateCurrency<br>Tasks.TaskActivities.NpcType.Loottable<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Tasks.TaskActivities.NpcType.Merchantlists<br>Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Tasks.TaskActivities.NpcType.NpcEmotes<br>Tasks.TaskActivities.NpcType.NpcFactions<br>Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Tasks.TaskActivities.NpcType.NpcSpells<br>Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Tasks.TaskActivities.NpcType.NpcTypesTint<br>Tasks.TaskActivities.NpcType.Spawnentries<br>Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Tasks.Tasksets<br>TradeskillRecipeEntries<br>TradeskillRecipeEntries.TradeskillRecipe<br>TributeLevels"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Item
// @Failure 500 {string} string "Bad query request"
// @Router /items [get]
func (e *ItemController) listItems(c echo.Context) error {
	var results []models.Item
	err := e.db.QueryContext(models.Item{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getItem godoc
// @Id getItem
// @Summary Gets Item
// @Accept json
// @Produce json
// @Tags Item
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AlternateCurrencies<br>CharacterCorpseItems<br>DiscoveredItems<br>Doors<br>Doors.Item<br>Fishings<br>Fishings.Item<br>Fishings.NpcType<br>Fishings.NpcType.AlternateCurrency<br>Fishings.NpcType.Loottable<br>Fishings.NpcType.Loottable.LoottableEntries<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Fishings.NpcType.Loottable.NpcTypes<br>Fishings.NpcType.Merchantlists<br>Fishings.NpcType.Merchantlists.NpcType<br>Fishings.NpcType.NpcEmotes<br>Fishings.NpcType.NpcFactions<br>Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Fishings.NpcType.NpcSpells<br>Fishings.NpcType.NpcSpells.NpcSpellsEntries<br>Fishings.NpcType.NpcTypesTint<br>Fishings.NpcType.Spawnentries<br>Fishings.NpcType.Spawnentries.NpcType<br>Fishings.NpcType.Spawnentries.Spawngroup<br>Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Fishings.Zone<br>Forages<br>Forages.Item<br>Forages.Zone<br>GroundSpawns<br>GroundSpawns.Zone<br>ItemTicks<br>Keyrings<br>LootdropEntries<br>LootdropEntries.Item<br>LootdropEntries.Lootdrop<br>LootdropEntries.Lootdrop.LootdropEntries<br>LootdropEntries.Lootdrop.LoottableEntries<br>LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Merchantlists<br>Merchantlists.NpcType<br>Merchantlists.NpcType.AlternateCurrency<br>Merchantlists.NpcType.Loottable<br>Merchantlists.NpcType.Loottable.LoottableEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Merchantlists.NpcType.Loottable.LoottableEntries.Loottable<br>Merchantlists.NpcType.Loottable.NpcTypes<br>Merchantlists.NpcType.Merchantlists<br>Merchantlists.NpcType.NpcEmotes<br>Merchantlists.NpcType.NpcFactions<br>Merchantlists.NpcType.NpcFactions.NpcFactionEntries<br>Merchantlists.NpcType.NpcSpells<br>Merchantlists.NpcType.NpcSpells.NpcSpellsEntries<br>Merchantlists.NpcType.NpcTypesTint<br>Merchantlists.NpcType.Spawnentries<br>Merchantlists.NpcType.Spawnentries.NpcType<br>Merchantlists.NpcType.Spawnentries.Spawngroup<br>Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2<br>Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>ObjectContents<br>Objects<br>Objects.Item<br>Objects.Zone<br>StartingItems<br>StartingItems.Item<br>StartingItems.Zone<br>Tasks<br>Tasks.TaskActivities<br>Tasks.TaskActivities.Goallists<br>Tasks.TaskActivities.NpcType<br>Tasks.TaskActivities.NpcType.AlternateCurrency<br>Tasks.TaskActivities.NpcType.Loottable<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Tasks.TaskActivities.NpcType.Merchantlists<br>Tasks.TaskActivities.NpcType.Merchantlists.NpcType<br>Tasks.TaskActivities.NpcType.NpcEmotes<br>Tasks.TaskActivities.NpcType.NpcFactions<br>Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Tasks.TaskActivities.NpcType.NpcSpells<br>Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries<br>Tasks.TaskActivities.NpcType.NpcTypesTint<br>Tasks.TaskActivities.NpcType.Spawnentries<br>Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Tasks.Tasksets<br>TradeskillRecipeEntries<br>TradeskillRecipeEntries.TradeskillRecipe<br>TributeLevels"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Item
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /item/{id} [get]
func (e *ItemController) getItem(c echo.Context) error {
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
	var result models.Item
	query := e.db.QueryContext(models.Item{}, c)
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

// updateItem godoc
// @Id updateItem
// @Summary Updates Item
// @Accept json
// @Produce json
// @Tags Item
// @Param id path int true "Id"
// @Param item body models.Item true "Item"
// @Success 200 {array} models.Item
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /item/{id} [patch]
func (e *ItemController) updateItem(c echo.Context) error {
	request := new(models.Item)
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
	var result models.Item
	query := e.db.QueryContext(models.Item{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Item{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createItem godoc
// @Id createItem
// @Summary Creates Item
// @Accept json
// @Produce json
// @Param item body models.Item true "Item"
// @Tags Item
// @Success 200 {array} models.Item
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /item [put]
func (e *ItemController) createItem(c echo.Context) error {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Item{}, c).Model(&models.Item{}).Create(&item).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, item)
}

// deleteItem godoc
// @Id deleteItem
// @Summary Deletes Item
// @Accept json
// @Produce json
// @Tags Item
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /item/{id} [delete]
func (e *ItemController) deleteItem(c echo.Context) error {
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
	var result models.Item
	query := e.db.QueryContext(models.Item{}, c)
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

// getItemsBulk godoc
// @Id getItemsBulk
// @Summary Gets Items in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Item
// @Success 200 {array} models.Item
// @Failure 500 {string} string "Bad query request"
// @Router /items/bulk [post]
func (e *ItemController) getItemsBulk(c echo.Context) error {
	var results []models.Item

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

	err := e.db.QueryContext(models.Item{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
