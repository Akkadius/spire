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

type Spawn2Controller struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewSpawn2Controller(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *Spawn2Controller {
	return &Spawn2Controller{
		db:     db,
		logger: logger,
	}
}

func (e *Spawn2Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "spawn_2/:spawn_2", e.deleteSpawn2, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_2/:spawn_2", e.getSpawn2, nil),
		routes.RegisterRoute(http.MethodGet, "spawn_2s", e.listSpawn2s, nil),
		routes.RegisterRoute(http.MethodPost, "spawn_2s/bulk", e.getSpawn2sBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "spawn_2/:spawn_2", e.updateSpawn2, nil),
		routes.RegisterRoute(http.MethodPut, "spawn_2", e.createSpawn2, nil),
	}
}

// listSpawn2s godoc
// @Id listSpawn2s
// @Summary Lists Spawn2s
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.NpcType.AlternateCurrency<br>Spawnentries.NpcType.Merchantlists<br>Spawnentries.NpcType.NpcFactions<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawnentries.NpcType.NpcSpells<br>Spawnentries.NpcType.NpcSpells.NpcSpellsEntries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawngroup"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawn2
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2s [get]
func (e *Spawn2Controller) listSpawn2s(c echo.Context) error {
	var results []models.Spawn2
	err := e.db.QueryContext(models.Spawn2{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getSpawn2 godoc
// @Id getSpawn2
// @Summary Gets Spawn2
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names <h4>Relationships</h4>Spawnentries<br>Spawnentries.NpcType<br>Spawnentries.NpcType.AlternateCurrency<br>Spawnentries.NpcType.Merchantlists<br>Spawnentries.NpcType.NpcFactions<br>Spawnentries.NpcType.NpcFactions.NpcFactionEntries<br>Spawnentries.NpcType.NpcSpells<br>Spawnentries.NpcType.NpcSpells.NpcSpellsEntries<br>Spawnentries.Spawngroup<br>Spawnentries.Spawngroup.Spawn2<br>Spawngroup"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Spawn2
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2/{id} [get]
func (e *Spawn2Controller) getSpawn2(c echo.Context) error {
	spawn2Id, err := strconv.Atoi(c.Param("spawn_2"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Spawn2
	err = e.db.QueryContext(models.Spawn2{}, c).First(&result, spawn2Id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateSpawn2 godoc
// @Id updateSpawn2
// @Summary Updates Spawn2
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param id path int true "Id"
// @Param spawn_2 body models.Spawn2 true "Spawn2"
// @Success 200 {array} models.Spawn2
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /spawn_2/{id} [patch]
func (e *Spawn2Controller) updateSpawn2(c echo.Context) error {
	spawn2 := new(models.Spawn2)
	if err := c.Bind(spawn2); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Spawn2{}
	err := e.db.Get(models.Spawn2{}, c).Model(&models.Spawn2{}).First(&entity, spawn2.ID).Error
	if err != nil || spawn2.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Spawn2{}, c).Model(&entity).Updates(&spawn2).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, spawn2)
}

// createSpawn2 godoc
// @Id createSpawn2
// @Summary Creates Spawn2
// @Accept json
// @Produce json
// @Param spawn_2 body models.Spawn2 true "Spawn2"
// @Tags Spawn2
// @Success 200 {array} models.Spawn2
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /spawn_2 [put]
func (e *Spawn2Controller) createSpawn2(c echo.Context) error {
	spawn2 := new(models.Spawn2)
	if err := c.Bind(spawn2); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Spawn2{}, c).Model(&models.Spawn2{}).Create(&spawn2).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, spawn2)
}

// deleteSpawn2 godoc
// @Id deleteSpawn2
// @Summary Deletes Spawn2
// @Accept json
// @Produce json
// @Tags Spawn2
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /spawn_2/{id} [delete]
func (e *Spawn2Controller) deleteSpawn2(c echo.Context) error {
	spawn2Id, err := strconv.Atoi(c.Param("spawn_2"))
	if err != nil {
		e.logger.Error(err)
	}

	spawn2 := new(models.Spawn2)
	err = e.db.Get(models.Spawn2{}, c).Model(&models.Spawn2{}).First(&spawn2, spawn2Id).Error
	if err != nil || spawn2.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Spawn2{}, c).Model(&models.Spawn2{}).Delete(&spawn2).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getSpawn2sBulk godoc
// @Id getSpawn2sBulk
// @Summary Gets Spawn2s in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Spawn2
// @Success 200 {array} models.Spawn2
// @Failure 500 {string} string "Bad query request"
// @Router /spawn_2s/bulk [post]
func (e *Spawn2Controller) getSpawn2sBulk(c echo.Context) error {
	var results []models.Spawn2

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.Spawn2{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
