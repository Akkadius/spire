package system

import (
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
)

type Controller struct {
	logger *logrus.Logger
}

func NewController(
	logger *logrus.Logger,
) *Controller {
	return &Controller{
		logger: logger,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "admin/system/host", a.getHostInfo, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/cpu", a.getCpu, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/mem", a.getMemory, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/network", a.getNetwork, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/disk", a.getDisk, nil),
	}
}

func (a *Controller) getHostInfo(c echo.Context) error {
	host, err := host.Info()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(http.StatusOK, host)
}

func (a *Controller) getCpu(c echo.Context) error {
	ci, err := cpu.Percent(0, true)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"cpu_percents": ci,
			"info":         cpuInfo,
		},
	)
}

func (a *Controller) getMemory(c echo.Context) error {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(http.StatusOK, memory)
}

func (a *Controller) getNetwork(c echo.Context) error {
	n, err := net.IOCounters(false)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(http.StatusOK, n)
}

type DiskInfo struct {
	PartitionStats []disk.PartitionStat  `json:"partition_stats"`
	DiskUsage      []disk.UsageStat      `json:"disk_usage"`
	IOCounters     []disk.IOCountersStat `json:"io_counters"`
}

func (a *Controller) getDisk(c echo.Context) error {
	part, _ := disk.Partitions(true)

	var partitionStats []disk.PartitionStat
	var diskUsage []disk.UsageStat
	var ioCounters []disk.IOCountersStat
	var uniqueDeviceNames []string

	for _, p := range part {
		partitionStats = append(partitionStats, p)

		usage, _ := disk.Usage(p.Mountpoint)
		diskUsage = append(diskUsage, *usage)

		counters, _ := disk.IOCounters(p.Device)
		if !contains(uniqueDeviceNames, p.Device) && len(counters[filepath.Base(p.Device)].Name) != 0 {
			ioCounters = append(ioCounters, counters[filepath.Base(p.Device)])
			uniqueDeviceNames = append(uniqueDeviceNames, p.Device)
		}
	}

	diskInfo := DiskInfo{
		PartitionStats: partitionStats,
		DiskUsage:      diskUsage,
		IOCounters:     ioCounters,
	}

	return c.JSON(http.StatusOK, diskInfo)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
