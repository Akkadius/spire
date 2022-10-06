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

type SpellsNewController struct {
	db	   *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpellsNewController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *SpellsNewController {
	return &SpellsNewController{
		db:	    db,
		logger: logger,
	}
}

func (e *SpellsNewController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "spells_new/:id", e.getSpellsNew, nil),
		routes.RegisterRoute(http.MethodGet, "spells_news", e.listSpellsNews, nil),
		routes.RegisterRoute(http.MethodPut, "spells_new", e.createSpellsNew, nil),
		routes.RegisterRoute(http.MethodDelete, "spells_new/:id", e.deleteSpellsNew, nil),
		routes.RegisterRoute(http.MethodPatch, "spells_new/:id", e.updateSpellsNew, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getSpellsNewsBulk, nil),
	}
}

// listSpellsNews godoc
// @Id listSpellsNews
// @Summary Lists SpellsNews
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Aura<br>Aura.SpellsNew<br>BlockedSpells<br>Damageshieldtypes<br>Items<br>Items.AlternateCurrencies<br>Items.AlternateCurrencies.Item<br>Items.CharacterCorpseItems<br>Items.DiscoveredItems<br>Items.Doors<br>Items.Doors.Item<br>Items.Fishings<br>Items.Fishings.Item<br>Items.Fishings.NpcType<br>Items.Fishings.NpcType.AlternateCurrency<br>Items.Fishings.NpcType.AlternateCurrency.Item<br>Items.Fishings.NpcType.Loottable<br>Items.Fishings.NpcType.Loottable.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Items.Fishings.NpcType.Loottable.NpcTypes<br>Items.Fishings.NpcType.Merchantlists<br>Items.Fishings.NpcType.Merchantlists.Items<br>Items.Fishings.NpcType.Merchantlists.NpcTypes<br>Items.Fishings.NpcType.NpcEmotes<br>Items.Fishings.NpcType.NpcFactions<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Items.Fishings.NpcType.NpcSpell<br>Items.Fishings.NpcType.NpcSpell.NpcSpell<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Fishings.NpcType.NpcTypesTint<br>Items.Fishings.NpcType.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.NpcType<br>Items.Fishings.NpcType.Spawnentries.Spawngroup<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Fishings.Zone<br>Items.Forages<br>Items.Forages.Item<br>Items.Forages.Zone<br>Items.GroundSpawns<br>Items.GroundSpawns.Zone<br>Items.ItemTicks<br>Items.Keyrings<br>Items.LootdropEntries<br>Items.LootdropEntries.Item<br>Items.LootdropEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LootdropEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Merchantlists<br>Items.Merchantlists.Items<br>Items.Merchantlists.NpcTypes<br>Items.Merchantlists.NpcTypes.AlternateCurrency<br>Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Items.Merchantlists.NpcTypes.Loottable<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Items.Merchantlists.NpcTypes.Merchantlists<br>Items.Merchantlists.NpcTypes.NpcEmotes<br>Items.Merchantlists.NpcTypes.NpcFactions<br>Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Items.Merchantlists.NpcTypes.NpcSpell<br>Items.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Merchantlists.NpcTypes.NpcTypesTint<br>Items.Merchantlists.NpcTypes.Spawnentries<br>Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.ObjectContents<br>Items.Objects<br>Items.Objects.Item<br>Items.Objects.Zone<br>Items.StartingItems<br>Items.StartingItems.Item<br>Items.StartingItems.Zone<br>Items.TradeskillRecipeEntries<br>Items.TradeskillRecipeEntries.TradeskillRecipe<br>Items.TributeLevels<br>NpcSpellsEntries<br>NpcSpellsEntries.SpellsNew<br>SpellBuckets<br>SpellGlobals"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Bad query request"
// @Router /spells_news [get]
func (e *SpellsNewController) listSpellsNews(c echo.Context) error {
	var results []models.SpellsNew
	err := e.db.QueryContext(models.SpellsNew{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpellsNew godoc
// @Id getSpellsNew
// @Summary Gets SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Aura<br>Aura.SpellsNew<br>BlockedSpells<br>Damageshieldtypes<br>Items<br>Items.AlternateCurrencies<br>Items.AlternateCurrencies.Item<br>Items.CharacterCorpseItems<br>Items.DiscoveredItems<br>Items.Doors<br>Items.Doors.Item<br>Items.Fishings<br>Items.Fishings.Item<br>Items.Fishings.NpcType<br>Items.Fishings.NpcType.AlternateCurrency<br>Items.Fishings.NpcType.AlternateCurrency.Item<br>Items.Fishings.NpcType.Loottable<br>Items.Fishings.NpcType.Loottable.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Items.Fishings.NpcType.Loottable.NpcTypes<br>Items.Fishings.NpcType.Merchantlists<br>Items.Fishings.NpcType.Merchantlists.Items<br>Items.Fishings.NpcType.Merchantlists.NpcTypes<br>Items.Fishings.NpcType.NpcEmotes<br>Items.Fishings.NpcType.NpcFactions<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Items.Fishings.NpcType.NpcSpell<br>Items.Fishings.NpcType.NpcSpell.NpcSpell<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Fishings.NpcType.NpcTypesTint<br>Items.Fishings.NpcType.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.NpcType<br>Items.Fishings.NpcType.Spawnentries.Spawngroup<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Fishings.Zone<br>Items.Forages<br>Items.Forages.Item<br>Items.Forages.Zone<br>Items.GroundSpawns<br>Items.GroundSpawns.Zone<br>Items.ItemTicks<br>Items.Keyrings<br>Items.LootdropEntries<br>Items.LootdropEntries.Item<br>Items.LootdropEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LootdropEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.Merchantlists<br>Items.Merchantlists.Items<br>Items.Merchantlists.NpcTypes<br>Items.Merchantlists.NpcTypes.AlternateCurrency<br>Items.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Items.Merchantlists.NpcTypes.Loottable<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Items.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Items.Merchantlists.NpcTypes.Merchantlists<br>Items.Merchantlists.NpcTypes.NpcEmotes<br>Items.Merchantlists.NpcTypes.NpcFactions<br>Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Items.Merchantlists.NpcTypes.NpcSpell<br>Items.Merchantlists.NpcTypes.NpcSpell.NpcSpell<br>Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Items.Merchantlists.NpcTypes.NpcTypesTint<br>Items.Merchantlists.NpcTypes.Spawnentries<br>Items.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Items.ObjectContents<br>Items.Objects<br>Items.Objects.Item<br>Items.Objects.Zone<br>Items.StartingItems<br>Items.StartingItems.Item<br>Items.StartingItems.Zone<br>Items.TradeskillRecipeEntries<br>Items.TradeskillRecipeEntries.TradeskillRecipe<br>Items.TributeLevels<br>NpcSpellsEntries<br>NpcSpellsEntries.SpellsNew<br>SpellBuckets<br>SpellGlobals"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.SpellsNew
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spells_new/{id} [get]
func (e *SpellsNewController) getSpellsNew(c echo.Context) error {
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
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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

// updateSpellsNew godoc
// @Id updateSpellsNew
// @Summary Updates SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "Id"
// @Param spells_new body models.SpellsNew true "SpellsNew"
// @Success 200 {array} models.SpellsNew
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spells_new/{id} [patch]
func (e *SpellsNewController) updateSpellsNew(c echo.Context) error {
	request := new(models.SpellsNew)
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
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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

// createSpellsNew godoc
// @Id createSpellsNew
// @Summary Creates SpellsNew
// @Accept json
// @Produce json
// @Param spells_new body models.SpellsNew true "SpellsNew"
// @Tags SpellsNew
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spells_new [put]
func (e *SpellsNewController) createSpellsNew(c echo.Context) error {
	spellsNew := new(models.SpellsNew)
	if err := c.Bind(spellsNew); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.SpellsNew{}, c).Model(&models.SpellsNew{}).Create(&spellsNew).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, spellsNew)
}

// deleteSpellsNew godoc
// @Id deleteSpellsNew
// @Summary Deletes SpellsNew
// @Accept json
// @Produce json
// @Tags SpellsNew
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spells_new/{id} [delete]
func (e *SpellsNewController) deleteSpellsNew(c echo.Context) error {
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
	var result models.SpellsNew
	query := e.db.QueryContext(models.SpellsNew{}, c)
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

// getSpellsNewsBulk godoc
// @Id getSpellsNewsBulk
// @Summary Gets SpellsNews in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags SpellsNew
// @Success 200 {array} models.SpellsNew
// @Failure 500 {string} string "Bad query request"
// @Router /spells_news/bulk [post]
func (e *SpellsNewController) getSpellsNewsBulk(c echo.Context) error {
	var results []models.SpellsNew

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

	err := e.db.QueryContext(models.SpellsNew{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
