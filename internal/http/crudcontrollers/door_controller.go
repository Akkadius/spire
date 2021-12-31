package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DoorController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewDoorController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *DoorController {
	return &DoorController{
		db:     db,
		logger: logger,
	}
}

func (e *DoorController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "door/:door", e.deleteDoor, nil),
		routes.RegisterRoute(http.MethodGet, "door/:door", e.getDoor, nil),
		routes.RegisterRoute(http.MethodGet, "doors", e.listDoors, nil),
		routes.RegisterRoute(http.MethodPost, "doors/bulk", e.getDoorsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "door/:door", e.updateDoor, nil),
		routes.RegisterRoute(http.MethodPut, "door", e.createDoor, nil),
	}
}

// listDoors godoc
// @Id listDoors
// @Summary Lists Doors
// @Accept json
// @Produce json
// @Tags Door
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
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
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Door
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /door/{id} [get]
func (e *DoorController) getDoor(c echo.Context) error {
	doorId, err := strconv.Atoi(c.Param("door"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Door
	err = e.db.QueryContext(models.Door{}, c).First(&result, doorId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

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
	door := new(models.Door)
	if err := c.Bind(door); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Door{}
	err := e.db.Get(models.Door{}, c).Model(&models.Door{}).First(&entity, door.ID).Error
	if err != nil || door.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Door{}, c).Model(&entity).Select("*").Updates(&door).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, door)
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
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Door{}, c).Model(&models.Door{}).Create(&door).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, door)
}

// deleteDoor godoc
// @Id deleteDoor
// @Summary Deletes Door
// @Accept json
// @Produce json
// @Tags Door
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /door/{id} [delete]
func (e *DoorController) deleteDoor(c echo.Context) error {
	doorId, err := strconv.Atoi(c.Param("door"))
	if err != nil {
		e.logger.Error(err)
	}

	door := new(models.Door)
	err = e.db.Get(models.Door{}, c).Model(&models.Door{}).First(&door, doorId).Error
	if err != nil || door.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Door{}, c).Model(&models.Door{}).Delete(&door).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
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
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
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
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
