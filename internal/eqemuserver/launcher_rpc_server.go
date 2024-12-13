package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/console"
	spiremiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"unicode"
)

// LauncherDistributedNode represents a node in the distributed launcher fleet
// this also includes the root node
type LauncherDistributedNode struct {
	Address             string // Address of the client
	Hostname            string // Hostname of the client
	CurrentZoneCount    int    // Current number of zones running on this node
	TargetZonesToLaunch int    // Target number of zones to launch on this node
}

func (l *Launcher) StartRpcServer(port int) error {
	e := echo.New()
	e.POST("/api/v1/dzs/register", l.rpcRegister)
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

func (l *Launcher) rpcTest(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "Hello, World!"})
}

func (l *Launcher) rpcRegister(c echo.Context) error {
	// bind to RpcClientRegisterRequest
	var req RpcClientRegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": first(err.Error())})
	}

	l.nodes = append(
		l.nodes,
		LauncherDistributedNode{
			Address:  req.ClientAddress,
			Hostname: req.Hostname,
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
