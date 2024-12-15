package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	spiremiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"unicode"
)

const (
	LauncherNodeTypeRoot = "root"
	LauncherNodeTypeLeaf = "leaf"
)

// LauncherDistributedNode represents a node in the distributed launcher fleet
// this also includes the root node
type LauncherDistributedNode struct {
	Address          string    // Address of the client
	Hostname         string    // Hostname of the client
	CurrentZoneCount int       // Current number of zones running on this node
	LastSeen         time.Time // Last time this node was seen
	NodeType         string    // Type of node, root or leaf
	TargetZoneCount  int       // (state) Target number of zones to run on this node
	MaxZoneCount     int       // (state) Max zone count for leaf nodes - Loaded from eqemu_config
	AtMaxZoneCount   bool      // (state) If the node is at max zone count
}

// StartRpcServer starts the HTTP RPC server for the launcher
func (l *Launcher) StartRpcServer(port int) error {
	e := echo.New()

	// RPC routes
	e.POST("/api/v1/dzs/register", l.rpcRegisterLeaf)
	e.GET("/api/v1/dzs/zone-count", l.rpcZoneCountDynamic)
	e.POST("/api/v1/dzs/set-zone-count", l.rpcSetZoneCount)
	e.GET("/api/v1/dzs/test", l.rpcTest)

	e.Use(spiremiddleware.LoggerWithConfig(spiremiddleware.LoggerConfig{
		Format: fmt.Sprintf(
			"%sSpire › API ›%s [${status}] [${method}] [${uri}] [${latency_human}]\n",
			console.BoldWhite,
			console.Reset,
		),
		Output: e.Logger.Output(),
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("RPC_KEY")
			cfg, _ := l.serverconfig.Get()

			if key != cfg.Server.World.Key {
				return c.JSON(
					http.StatusUnauthorized,
					echo.Map{
						"error": fmt.Sprintf(
							"Invalid key, unauthorized. Request key [%v] does not match server key [%v]",
							key,
							cfg.Server.World.Key,
						),
					},
				)
			}

			return next(c)
		}
	})

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		_ = c.JSON(422, echo.Map{"error": first(err.Error())})
	}

	e.HideBanner = true
	e.HidePort = true

	l.logger.Info().Any("port", port).Msgf("Starting Spire DZS HTTP RPC server")

	return e.Start(fmt.Sprintf(":%v", port))
}

// rpcTest is a test route
func (l *Launcher) rpcTest(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "Hello, World!"})
}

// rpcRegisterLeaf registers a leaf node with the launcher
func (l *Launcher) rpcRegisterLeaf(c echo.Context) error {
	// bind to RpcClientRegisterRequest
	var req RpcClientRegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": first(err.Error())})
	}

	// check if the client is already registered
	for i, node := range l.nodes {
		if node.Address == req.ClientAddress {

			// reset
			l.nodes[i].LastSeen = time.Now()
			l.nodes[i].Hostname = req.Hostname
			l.nodes[i].TargetZoneCount = 0

			return c.JSON(
				http.StatusOK,
				echo.Map{
					"message": "Client already registered",
				},
			)
		}
	}

	l.nodes = append(
		l.nodes,
		LauncherDistributedNode{
			Address:  req.ClientAddress,
			Hostname: req.Hostname,
			LastSeen: time.Now(),
			NodeType: LauncherNodeTypeLeaf,
		},
	)

	l.logger.Info().
		Any("nodes", l.nodes).
		Any("client_address", req.ClientAddress).
		Any("hostname", req.Hostname).
		Msg("Client registered")

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"message": "Registered",
		},
	)
}

func first(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

// rpcZoneCountDynamic returns the number of zones currently running
func (l *Launcher) rpcZoneCountDynamic(c echo.Context) error {
	l.pollProcessCounts()

	cfg, _ := l.serverconfig.Get()
	maxZoneCount := 0
	if cfg.WebAdmin != nil && cfg.WebAdmin.Launcher != nil {
		maxZoneCount = cfg.WebAdmin.Launcher.DistributedMaxZoneCount
	}

	return c.JSON(
		http.StatusOK,
		RpcZoneCountResponse{
			ZoneCount:    l.bootedTotalDynamics,
			MaxZoneCount: maxZoneCount,
		},
	)
}

// rpcSetZoneCount sets the number of zones to boot
func (l *Launcher) rpcSetZoneCount(c echo.Context) error {
	var req RpcLaunchZonesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": first(err.Error())})
	}

	bootCount := req.ZoneCount - l.bootedTotalDynamics

	var errors []error
	for i := 0; i < bootCount; i++ {
		err := l.startServerProcess(zoneProcessName)
		if err != nil {
			l.logger.Error().Err(err).Msg("Failed to start zone process")
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"errors": errors})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Zones started"})
}
