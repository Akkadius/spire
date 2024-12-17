package system

import (
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"net/http"
	"path/filepath"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "admin/system/host", a.getHostInfo, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/cpu", a.getCpuAll, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/cpu-percent", a.getCpuPercent, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/resource-usage-summary", a.getResourceUsageSummary, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/mem", a.getMemory, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/network", a.getNetwork, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/disk", a.getDisk, nil),
		routes.RegisterRoute(http.MethodGet, "admin/system/all", a.systemAll, nil),
	}
}

func (a *Controller) getHostInfo(c echo.Context) error {
	h, err := host.Info()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(http.StatusOK, h)
}

func (a *Controller) getCpuAll(c echo.Context) error {
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

func (a *Controller) getCpuPercent(c echo.Context) error {
	ci, err := cpu.Percent(0, false)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(http.StatusOK, ci[0])
}

func (a *Controller) getResourceUsageSummary(c echo.Context) error {
	ci, err := cpu.Percent(0, false)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	memory, err := mem.VirtualMemory()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{
			"cpu":    ci[0],
			"memory": memory,
		},
	)
}

// systemAll returns all system information
func (a *Controller) systemAll(c echo.Context) error {
	all, err := All()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(http.StatusOK, all)
}

// contains checks if a string is in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
