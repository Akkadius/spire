package crudcontrollers

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CharacterDatumController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewCharacterDatumController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *CharacterDatumController {
	return &CharacterDatumController{
		db:     db,
		logger: logger,
	}
}

func (e *CharacterDatumController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "character_datum/:character_datum", e.deleteCharacterDatum, nil),
		routes.RegisterRoute(http.MethodGet, "character_datum/:character_datum", e.getCharacterDatum, nil),
		routes.RegisterRoute(http.MethodGet, "character_data", e.listCharacterData, nil),
		routes.RegisterRoute(http.MethodPost, "spells_news/bulk", e.getCharacterDataBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "character_datum/:character_datum", e.updateCharacterDatum, nil),
		routes.RegisterRoute(http.MethodPut, "character_datum", e.createCharacterDatum, nil),
	}
}

// listCharacterData godoc
// @Id listCharacterData
// @Summary Lists CharacterData
// @Accept json
// @Produce json
// @Tags CharacterDatum
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AdventureStats<br>Buyers<br>CharRecipeLists<br>CharacterActivities<br>CharacterAltCurrencies<br>CharacterAlternateAbilities<br>CharacterAuras<br>CharacterBandoliers<br>CharacterBinds<br>CharacterBuffs<br>CharacterCorpses<br>CharacterCurrencies<br>CharacterDisciplines<br>CharacterEnabledtasks<br>CharacterInspectMessages<br>CharacterItemRecasts<br>CharacterLanguages<br>CharacterLeadershipAbilities<br>CharacterMaterials<br>CharacterMemmedSpells<br>CharacterPetBuffs<br>CharacterPetInfos<br>CharacterPetInventories<br>CharacterPotionbelts<br>CharacterSkills<br>CharacterSpells<br>CharacterTasks<br>CharacterTributes<br>CompletedTasks<br>DataBuckets<br>FactionValues<br>Friends<br>Guild<br>Guild.GuildBanks<br>Guild.GuildMembers<br>Guild.GuildRanks<br>GuildMembers<br>InstanceListPlayers<br>Inventories<br>Inventories.Item<br>Inventories.Item.DiscoveredItems<br>Keyrings<br>Mail<br>PlayerTitlesets<br>QuestGlobals<br>Timers<br>Titles<br>Traders<br>ZoneFlags"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>AdventureStats<br>Buyers<br>CharRecipeLists<br>CharacterActivities<br>CharacterAltCurrencies<br>CharacterAlternateAbilities<br>CharacterAuras<br>CharacterBandoliers<br>CharacterBinds<br>CharacterBuffs<br>CharacterCorpses<br>CharacterCurrencies<br>CharacterDisciplines<br>CharacterEnabledtasks<br>CharacterInspectMessages<br>CharacterItemRecasts<br>CharacterLanguages<br>CharacterLeadershipAbilities<br>CharacterMaterials<br>CharacterMemmedSpells<br>CharacterPetBuffs<br>CharacterPetInfos<br>CharacterPetInventories<br>CharacterPotionbelts<br>CharacterSkills<br>CharacterSpells<br>CharacterTasks<br>CharacterTributes<br>CompletedTasks<br>DataBuckets<br>FactionValues<br>Friends<br>Guild<br>Guild.GuildBanks<br>Guild.GuildMembers<br>Guild.GuildRanks<br>GuildMembers<br>InstanceListPlayers<br>Inventories<br>Inventories.Item<br>Inventories.Item.DiscoveredItems<br>Keyrings<br>Mail<br>PlayerTitlesets<br>QuestGlobals<br>Timers<br>Titles<br>Traders<br>ZoneFlags"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.CharacterDatum
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /character_datum/{id} [get]
func (e *CharacterDatumController) getCharacterDatum(c echo.Context) error {
	characterDatumId, err := strconv.Atoi(c.Param("character_datum"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.CharacterDatum
	err = e.db.QueryContext(models.CharacterDatum{}, c).First(&result, characterDatumId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

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
	characterDatum := new(models.CharacterDatum)
	if err := c.Bind(characterDatum); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterDatum{}, c).Model(&models.CharacterDatum{}).First(&models.CharacterDatum{}, characterDatum.ID).Error
	if err != nil || characterDatum.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterDatum{}, c).Model(&models.CharacterDatum{}).Updates(&characterDatum).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, characterDatum)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.CharacterDatum{}, c).Model(&models.CharacterDatum{}).Create(&characterDatum).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
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
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /character_datum/{id} [delete]
func (e *CharacterDatumController) deleteCharacterDatum(c echo.Context) error {
	characterDatumId, err := strconv.Atoi(c.Param("character_datum"))
	if err != nil {
		e.logger.Error(err)
	}

	characterDatum := new(models.CharacterDatum)
	err = e.db.Get(models.CharacterDatum{}, c).Model(&models.CharacterDatum{}).First(&characterDatum, characterDatumId).Error
	if err != nil || characterDatum.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.CharacterDatum{}, c).Model(&models.CharacterDatum{}).Delete(&characterDatum).Error
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.CharacterDatum{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
