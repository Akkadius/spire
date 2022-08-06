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

type CharacterDatumController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterDatumController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterDatumController {
	return &CharacterDatumController{
		db:	 db,
		logger: logger,
	}
}

func (e *CharacterDatumController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "character_datum/:id", e.getCharacterDatum, nil),
		routes.RegisterRoute(http.MethodGet, "character_data", e.listCharacterData, nil),
		routes.RegisterRoute(http.MethodPut, "character_datum", e.createCharacterDatum, nil),
		routes.RegisterRoute(http.MethodDelete, "character_datum/:id", e.deleteCharacterDatum, nil),
		routes.RegisterRoute(http.MethodPatch, "character_datum/:id", e.updateCharacterDatum, nil),
		routes.RegisterRoute(http.MethodPost, "character_data/bulk", e.getCharacterDataBulk, nil),
	}
}

// listCharacterData godoc
// @Id listCharacterData
// @Summary Lists CharacterData
// @Accept json
// @Produce json
// @Tags CharacterDatum
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AdventureStats<br>Buyers<br>CharRecipeLists<br>CharacterActivities<br>CharacterAltCurrencies<br>CharacterAlternateAbilities<br>CharacterAuras<br>CharacterBandoliers<br>CharacterBinds<br>CharacterBuffs<br>CharacterCorpses<br>CharacterCurrencies<br>CharacterDisciplines<br>CharacterEnabledtasks<br>CharacterInspectMessages<br>CharacterItemRecasts<br>CharacterLanguages<br>CharacterLeadershipAbilities<br>CharacterMaterials<br>CharacterMemmedSpells<br>CharacterPetBuffs<br>CharacterPetInfos<br>CharacterPetInventories<br>CharacterPotionbelts<br>CharacterSkills<br>CharacterSpells<br>CharacterTasks<br>CharacterTributes<br>CompletedTasks<br>DataBuckets<br>FactionValues<br>Friends<br>Guild<br>Guild.GuildBanks<br>Guild.GuildMembers<br>Guild.GuildRanks<br>GuildMembers<br>InstanceListPlayers<br>Inventories<br>Inventories.Item<br>Inventories.Item.AlternateCurrencies<br>Inventories.Item.AlternateCurrencies.Item<br>Inventories.Item.CharacterCorpseItems<br>Inventories.Item.DiscoveredItems<br>Inventories.Item.Doors<br>Inventories.Item.Doors.Item<br>Inventories.Item.Fishings<br>Inventories.Item.Fishings.Item<br>Inventories.Item.Fishings.NpcType<br>Inventories.Item.Fishings.NpcType.AlternateCurrency<br>Inventories.Item.Fishings.NpcType.AlternateCurrency.Item<br>Inventories.Item.Fishings.NpcType.Loottable<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Inventories.Item.Fishings.NpcType.Loottable.NpcTypes<br>Inventories.Item.Fishings.NpcType.Merchantlists<br>Inventories.Item.Fishings.NpcType.Merchantlists.Items<br>Inventories.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Inventories.Item.Fishings.NpcType.NpcEmotes<br>Inventories.Item.Fishings.NpcType.NpcFactions<br>Inventories.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Inventories.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.Fishings.NpcType.NpcSpell<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.Fishings.NpcType.NpcTypesTint<br>Inventories.Item.Fishings.NpcType.Spawnentries<br>Inventories.Item.Fishings.NpcType.Spawnentries.NpcType<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.Fishings.Zone<br>Inventories.Item.Forages<br>Inventories.Item.Forages.Item<br>Inventories.Item.Forages.Zone<br>Inventories.Item.GroundSpawns<br>Inventories.Item.GroundSpawns.Zone<br>Inventories.Item.ItemTicks<br>Inventories.Item.Keyrings<br>Inventories.Item.LootdropEntries<br>Inventories.Item.LootdropEntries.Item<br>Inventories.Item.LootdropEntries.Lootdrop<br>Inventories.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.Merchantlists<br>Inventories.Item.Merchantlists.Items<br>Inventories.Item.Merchantlists.NpcTypes<br>Inventories.Item.Merchantlists.NpcTypes.AlternateCurrency<br>Inventories.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Inventories.Item.Merchantlists.NpcTypes.Loottable<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Inventories.Item.Merchantlists.NpcTypes.Merchantlists<br>Inventories.Item.Merchantlists.NpcTypes.NpcEmotes<br>Inventories.Item.Merchantlists.NpcTypes.NpcFactions<br>Inventories.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Inventories.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.Merchantlists.NpcTypes.NpcTypesTint<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.ObjectContents<br>Inventories.Item.Objects<br>Inventories.Item.Objects.Item<br>Inventories.Item.Objects.Zone<br>Inventories.Item.StartingItems<br>Inventories.Item.StartingItems.Item<br>Inventories.Item.StartingItems.Zone<br>Inventories.Item.Tasks<br>Inventories.Item.Tasks.AlternateCurrency<br>Inventories.Item.Tasks.AlternateCurrency.Item<br>Inventories.Item.Tasks.TaskActivities<br>Inventories.Item.Tasks.TaskActivities.Goallists<br>Inventories.Item.Tasks.TaskActivities.NpcType<br>Inventories.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Inventories.Item.Tasks.TaskActivities.NpcType.AlternateCurrency.Item<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists.Items<br>Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcTypes<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.Tasks.Tasksets<br>Inventories.Item.TradeskillRecipeEntries<br>Inventories.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Inventories.Item.TributeLevels<br>Keyrings<br>Mail<br>PlayerTitlesets<br>QuestGlobals<br>Timers<br>Titles<br>Traders<br>ZoneFlags"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterDatum
// @Failure 500 {string} string "Bad query request"
// @Router /character_data [get]
func (e *CharacterDatumController) listCharacterData(c echo.Context) error {
	var results []models.CharacterDatum
	err := e.db.QueryContext(models.CharacterDatum{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getCharacterDatum godoc
// @Id getCharacterDatum
// @Summary Gets CharacterDatum
// @Accept json
// @Produce json
// @Tags CharacterDatum
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AdventureStats<br>Buyers<br>CharRecipeLists<br>CharacterActivities<br>CharacterAltCurrencies<br>CharacterAlternateAbilities<br>CharacterAuras<br>CharacterBandoliers<br>CharacterBinds<br>CharacterBuffs<br>CharacterCorpses<br>CharacterCurrencies<br>CharacterDisciplines<br>CharacterEnabledtasks<br>CharacterInspectMessages<br>CharacterItemRecasts<br>CharacterLanguages<br>CharacterLeadershipAbilities<br>CharacterMaterials<br>CharacterMemmedSpells<br>CharacterPetBuffs<br>CharacterPetInfos<br>CharacterPetInventories<br>CharacterPotionbelts<br>CharacterSkills<br>CharacterSpells<br>CharacterTasks<br>CharacterTributes<br>CompletedTasks<br>DataBuckets<br>FactionValues<br>Friends<br>Guild<br>Guild.GuildBanks<br>Guild.GuildMembers<br>Guild.GuildRanks<br>GuildMembers<br>InstanceListPlayers<br>Inventories<br>Inventories.Item<br>Inventories.Item.AlternateCurrencies<br>Inventories.Item.AlternateCurrencies.Item<br>Inventories.Item.CharacterCorpseItems<br>Inventories.Item.DiscoveredItems<br>Inventories.Item.Doors<br>Inventories.Item.Doors.Item<br>Inventories.Item.Fishings<br>Inventories.Item.Fishings.Item<br>Inventories.Item.Fishings.NpcType<br>Inventories.Item.Fishings.NpcType.AlternateCurrency<br>Inventories.Item.Fishings.NpcType.AlternateCurrency.Item<br>Inventories.Item.Fishings.NpcType.Loottable<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable<br>Inventories.Item.Fishings.NpcType.Loottable.NpcTypes<br>Inventories.Item.Fishings.NpcType.Merchantlists<br>Inventories.Item.Fishings.NpcType.Merchantlists.Items<br>Inventories.Item.Fishings.NpcType.Merchantlists.NpcTypes<br>Inventories.Item.Fishings.NpcType.NpcEmotes<br>Inventories.Item.Fishings.NpcType.NpcFactions<br>Inventories.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries<br>Inventories.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.Fishings.NpcType.NpcSpell<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.Fishings.NpcType.NpcTypesTint<br>Inventories.Item.Fishings.NpcType.Spawnentries<br>Inventories.Item.Fishings.NpcType.Spawnentries.NpcType<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.Fishings.Zone<br>Inventories.Item.Forages<br>Inventories.Item.Forages.Item<br>Inventories.Item.Forages.Zone<br>Inventories.Item.GroundSpawns<br>Inventories.Item.GroundSpawns.Zone<br>Inventories.Item.ItemTicks<br>Inventories.Item.Keyrings<br>Inventories.Item.LootdropEntries<br>Inventories.Item.LootdropEntries.Item<br>Inventories.Item.LootdropEntries.Lootdrop<br>Inventories.Item.LootdropEntries.Lootdrop.LootdropEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.Merchantlists<br>Inventories.Item.Merchantlists.Items<br>Inventories.Item.Merchantlists.NpcTypes<br>Inventories.Item.Merchantlists.NpcTypes.AlternateCurrency<br>Inventories.Item.Merchantlists.NpcTypes.AlternateCurrency.Item<br>Inventories.Item.Merchantlists.NpcTypes.Loottable<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable<br>Inventories.Item.Merchantlists.NpcTypes.Loottable.NpcTypes<br>Inventories.Item.Merchantlists.NpcTypes.Merchantlists<br>Inventories.Item.Merchantlists.NpcTypes.NpcEmotes<br>Inventories.Item.Merchantlists.NpcTypes.NpcFactions<br>Inventories.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries<br>Inventories.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.Merchantlists.NpcTypes.NpcTypesTint<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.NpcType<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.ObjectContents<br>Inventories.Item.Objects<br>Inventories.Item.Objects.Item<br>Inventories.Item.Objects.Zone<br>Inventories.Item.StartingItems<br>Inventories.Item.StartingItems.Item<br>Inventories.Item.StartingItems.Zone<br>Inventories.Item.Tasks<br>Inventories.Item.Tasks.AlternateCurrency<br>Inventories.Item.Tasks.AlternateCurrency.Item<br>Inventories.Item.Tasks.TaskActivities<br>Inventories.Item.Tasks.TaskActivities.Goallists<br>Inventories.Item.Tasks.TaskActivities.NpcType<br>Inventories.Item.Tasks.TaskActivities.NpcType.AlternateCurrency<br>Inventories.Item.Tasks.TaskActivities.NpcType.AlternateCurrency.Item<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable<br>Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes<br>Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists<br>Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists.Items<br>Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcTypes<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcEmotes<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals<br>Inventories.Item.Tasks.TaskActivities.NpcType.NpcTypesTint<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup<br>Inventories.Item.Tasks.Tasksets<br>Inventories.Item.TradeskillRecipeEntries<br>Inventories.Item.TradeskillRecipeEntries.TradeskillRecipe<br>Inventories.Item.TributeLevels<br>Keyrings<br>Mail<br>PlayerTitlesets<br>QuestGlobals<br>Timers<br>Titles<br>Traders<br>ZoneFlags"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterDatum
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_datum/{id} [get]
func (e *CharacterDatumController) getCharacterDatum(c echo.Context) error {
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
	var result models.CharacterDatum
	query := e.db.QueryContext(models.CharacterDatum{}, c)
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

// updateCharacterDatum godoc
// @Id updateCharacterDatum
// @Summary Updates CharacterDatum
// @Accept json
// @Produce json
// @Tags CharacterDatum
// @Param id path int true "Id"
// @Param character_datum body models.CharacterDatum true "CharacterDatum"
// @Success 200 {array} models.CharacterDatum
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /character_datum/{id} [patch]
func (e *CharacterDatumController) updateCharacterDatum(c echo.Context) error {
	request := new(models.CharacterDatum)
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
	var result models.CharacterDatum
	query := e.db.QueryContext(models.CharacterDatum{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.CharacterDatum{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createCharacterDatum godoc
// @Id createCharacterDatum
// @Summary Creates CharacterDatum
// @Accept json
// @Produce json
// @Param character_datum body models.CharacterDatum true "CharacterDatum"
// @Tags CharacterDatum
// @Success 200 {array} models.CharacterDatum
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /character_datum [put]
func (e *CharacterDatumController) createCharacterDatum(c echo.Context) error {
	characterDatum := new(models.CharacterDatum)
	if err := c.Bind(characterDatum); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.CharacterDatum{}, c).Model(&models.CharacterDatum{}).Create(&characterDatum).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, characterDatum)
}

// deleteCharacterDatum godoc
// @Id deleteCharacterDatum
// @Summary Deletes CharacterDatum
// @Accept json
// @Produce json
// @Tags CharacterDatum
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_datum/{id} [delete]
func (e *CharacterDatumController) deleteCharacterDatum(c echo.Context) error {
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
	var result models.CharacterDatum
	query := e.db.QueryContext(models.CharacterDatum{}, c)
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

// getCharacterDataBulk godoc
// @Id getCharacterDataBulk
// @Summary Gets CharacterData in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags CharacterDatum
// @Success 200 {array} models.CharacterDatum
// @Failure 500 {string} string "Bad query request"
// @Router /character_data/bulk [post]
func (e *CharacterDatumController) getCharacterDataBulk(c echo.Context) error {
	var results []models.CharacterDatum

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

	err := e.db.QueryContext(models.CharacterDatum{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
