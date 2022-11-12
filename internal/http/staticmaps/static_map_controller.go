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
	//go:embed emitters.json
	emittersMap []byte
	//go:embed item-icons-map.json
	itemIconsMap []byte
	//go:embed monograms-map.json
	monogramsMap []byte
	//go:embed npc-models-map.json
	npcModelsMap []byte
	//go:embed objects-map.json
	objectsMap []byte
	//go:embed player-animations.json
	playerAnimationsMap []byte
	//go:embed race-inventory-map.json
	raceInventoryMap []byte
	//go:embed spell-animations-map.json
	spellAnimationsMap []byte
	//go:embed spell-icons-map.json
	spellIconsMap []byte
	//go:embed spell-icon-anim-name-map.json
	spellIconAnimNameMap []byte
	//go:embed model-relationships.json
	modelRelationships []byte
)

func (q *StaticMapController) getStaticMapFileCompressed(c echo.Context) error {
	file := []byte{}
	switch c.Param("file") {
	case "emitters.json":
		file = emittersMap
	case "item-icons-map.json":
		file = itemIconsMap
	case "monograms-map.json":
		file = monogramsMap
	case "npc-models-map.json":
		file = npcModelsMap
	case "objects-map.json":
		file = objectsMap
	case "player-animations.json":
		file = playerAnimationsMap
	case "spell-animations-map.json":
		file = spellAnimationsMap
	case "spell-icons-map.json":
		file = spellIconsMap
	case "race-inventory-map.json":
		file = raceInventoryMap
	case "spell-icon-anim-name-map.json":
		file = spellIconAnimNameMap
	case "model-relationships.json":
		file = modelRelationships
	default:
		return fmt.Errorf("Invalid file specified")
	}

	return c.JSONBlob(http.StatusOK, file)
}
