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

type MerchantlistController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewMerchantlistController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *MerchantlistController {
	return &MerchantlistController{
		db:	 db,
		logger: logger,
	}
}

func (e *MerchantlistController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "merchantlist/:merchantid", e.getMerchantlist, nil),
		routes.RegisterRoute(http.MethodGet, "merchantlists", e.listMerchantlists, nil),
		routes.RegisterRoute(http.MethodPut, "merchantlist", e.createMerchantlist, nil),
		routes.RegisterRoute(http.MethodDelete, "merchantlist/:merchantid", e.deleteMerchantlist, nil),
		routes.RegisterRoute(http.MethodPatch, "merchantlist/:merchantid", e.updateMerchantlist, nil),
		routes.RegisterRoute(http.MethodPost, "merchantlists/bulk", e.getMerchantlistsBulk, nil),
	}
}

// listMerchantlists godoc
// @Id listMerchantlists
// @Summary Lists Merchantlists
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcSpells<br>NpcType.NpcSpells.NpcSpellsEntries<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.NpcType<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Merchantlist
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlists [get]
func (e *MerchantlistController) listMerchantlists(c echo.Context) error {
	var results []models.Merchantlist
	err := e.db.QueryContext(models.Merchantlist{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getMerchantlist godoc
// @Id getMerchantlist
// @Summary Gets Merchantlist
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>NpcType<br>NpcType.AlternateCurrency<br>NpcType.Loottable<br>NpcType.Loottable.LoottableEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries<br>NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries<br>NpcType.Loottable.LoottableEntries.Loottable<br>NpcType.Loottable.NpcTypes<br>NpcType.Merchantlists<br>NpcType.NpcEmotes<br>NpcType.NpcFactions<br>NpcType.NpcFactions.NpcFactionEntries<br>NpcType.NpcSpells<br>NpcType.NpcSpells.NpcSpellsEntries<br>NpcType.NpcTypesTint<br>NpcType.Spawnentries<br>NpcType.Spawnentries.NpcType<br>NpcType.Spawnentries.Spawngroup<br>NpcType.Spawnentries.Spawngroup.Spawn2<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries<br>NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Merchantlist
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlist/{id} [get]
func (e *MerchantlistController) getMerchantlist(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	merchantid, err := strconv.Atoi(c.Param("merchantid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Merchantid]"})
	}
	params = append(params, merchantid)
	keys = append(keys, "merchantid = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.Merchantlist
	query := e.db.QueryContext(models.Merchantlist{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.Merchantid == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateMerchantlist godoc
// @Id updateMerchantlist
// @Summary Updates Merchantlist
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param id path int true "Id"
// @Param merchantlist body models.Merchantlist true "Merchantlist"
// @Success 200 {array} models.Merchantlist
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /merchantlist/{id} [patch]
func (e *MerchantlistController) updateMerchantlist(c echo.Context) error {
	request := new(models.Merchantlist)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	merchantid, err := strconv.Atoi(c.Param("merchantid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Merchantid]"})
	}
	params = append(params, merchantid)
	keys = append(keys, "merchantid = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.Merchantlist
	query := e.db.QueryContext(models.Merchantlist{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = e.db.QueryContext(models.Merchantlist{}, c).Select("*").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createMerchantlist godoc
// @Id createMerchantlist
// @Summary Creates Merchantlist
// @Accept json
// @Produce json
// @Param merchantlist body models.Merchantlist true "Merchantlist"
// @Tags Merchantlist
// @Success 200 {array} models.Merchantlist
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /merchantlist [put]
func (e *MerchantlistController) createMerchantlist(c echo.Context) error {
	merchantlist := new(models.Merchantlist)
	if err := c.Bind(merchantlist); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.Merchantlist{}, c).Model(&models.Merchantlist{}).Create(&merchantlist).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, merchantlist)
}

// deleteMerchantlist godoc
// @Id deleteMerchantlist
// @Summary Deletes Merchantlist
// @Accept json
// @Produce json
// @Tags Merchantlist
// @Param id path int true "merchantid"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /merchantlist/{id} [delete]
func (e *MerchantlistController) deleteMerchantlist(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	merchantid, err := strconv.Atoi(c.Param("merchantid"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, merchantid)
	keys = append(keys, "merchantid = ?")

	// key param [slot] position [2] type [int]
	if len(c.QueryParam("slot")) > 0 {
		slotParam, err := strconv.Atoi(c.QueryParam("slot"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [slot] err [%s]", err.Error())})
		}

		params = append(params, slotParam)
		keys = append(keys, "slot = ?")
	}

	// query builder
	var result models.Merchantlist
	query := e.db.QueryContext(models.Merchantlist{}, c)
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

// getMerchantlistsBulk godoc
// @Id getMerchantlistsBulk
// @Summary Gets Merchantlists in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Merchantlist
// @Success 200 {array} models.Merchantlist
// @Failure 500 {string} string "Bad query request"
// @Router /merchantlists/bulk [post]
func (e *MerchantlistController) getMerchantlistsBulk(c echo.Context) error {
	var results []models.Merchantlist

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

	err := e.db.QueryContext(models.Merchantlist{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
