package eqemuserverapi

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Controller struct {
	db             *database.DatabaseResolver
	logger         *logrus.Logger
	eqemuserverapi *Client
	serverconfig   *serverconfig.EQEmuServerConfig
}

func NewController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	api *Client,
	serverconfig *serverconfig.EQEmuServerConfig,
) *Controller {
	return &Controller{
		db:             db,
		logger:         logger,
		eqemuserverapi: api,
		serverconfig:   serverconfig,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "eqemuserver/zone-list", a.getZoneList, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/client-list", a.getClientList, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/server-stats", a.getServerStats, nil),
		routes.RegisterRoute(http.MethodGet, "eqemuserver/reload-types", a.getReloadTypes, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/reload/:type", a.reload, nil),
	}
}

func (a *Controller) getReloadTypes(c echo.Context) error {
	types, err := a.eqemuserverapi.GetReloadTypes()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, types)
}

func (a *Controller) reload(c echo.Context) error {
	reloadType := c.Param("type")
	r, err := a.eqemuserverapi.Reload(reloadType)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, r)
}

type ServerStatsResponse struct {
	ServerName      string          `json:"server_name"`
	LauncherOnline  bool            `json:"launcher_online"`
	UcsOnline       bool            `json:"ucs_online"`
	LoginOnline     bool            `json:"login_online"`
	WorldOnline     bool            `json:"world_online"`
	QueryServOnline bool            `json:"query_serv_online"`
	ZoneList        WorldZoneList   `json:"zone_list"`
	PlayersOnline   WorldClientList `json:"client_list"`
}

func (a *Controller) getServerStats(c echo.Context) error {
	var r ServerStatsResponse

	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			return err
		}

		if strings.Contains(cmdline, "server-launcher") {
			r.LauncherOnline = true
		}
		if strings.Contains(cmdline, "world") {
			r.WorldOnline = true
		}
		if strings.Contains(cmdline, "ucs") {
			r.UcsOnline = true
		}
		if strings.Contains(cmdline, "loginserver") {
			r.LoginOnline = true
		}
		if strings.Contains(cmdline, "queryserv") {
			r.QueryServOnline = true
		}
	}

	zoneList, _ := a.eqemuserverapi.GetZoneList()
	if len(zoneList.Data) > 0 {
		r.ZoneList = zoneList
	}
	clientList, _ := a.eqemuserverapi.GetWorldClientList()
	if len(clientList.Data) > 0 {
		r.PlayersOnline = clientList
	}

	r.ServerName = a.serverconfig.Get().Server.World.Longname

	return c.JSON(http.StatusOK, r)
}

func (a *Controller) getClientList(c echo.Context) error {
	types, err := a.eqemuserverapi.GetWorldClientList()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, types)
}

func (a *Controller) getZoneList(c echo.Context) error {
	types, err := a.eqemuserverapi.GetZoneList()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Failed to connect to gameserver [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, types)
}
