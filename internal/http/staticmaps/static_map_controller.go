package staticmaps

import (
	_ "embed"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type StaticMapController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewStaticMapController(db *database.DatabaseResolver, logger *logrus.Logger) *StaticMapController {
	return &StaticMapController{
		db:     db,
		logger: logger,
	}
}

func (q *StaticMapController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "static-map/:file", q.getStaticMapFileCompressed, nil),
	}
}

var (
	//go:embed race-inventory-map.json
	raceInventoryMap []byte
)

func (q *StaticMapController) getStaticMapFileCompressed(c echo.Context) error {
	file := []byte{}
	switch c.Param("file") {
	case "race-inventory-map.json":
		file = raceInventoryMap
	default:
		return fmt.Errorf("Invalid file specified")
	}

	return c.JSONBlob(http.StatusOK, file)
}
