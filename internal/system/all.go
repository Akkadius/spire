package system

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"os"
	"path/filepath"
)

// AllResponse is the response for the systemAll endpoint
type AllResponse struct {
	Hostname   string                `json:"hostname"`
	Cpu        float64               `json:"cpu"`
	MemPercent float64               `json:"mem_percent"`
	Disk       []disk.IOCountersStat `json:"disk"`
	Net        []net.IOCountersStat  `json:"net"`
}

// All returns all system information
func All() (AllResponse, error) {
	// cpu
	ci, err := cpu.Percent(0, false)
	if err != nil {
		return AllResponse{}, err
	}

	// memory
	memory, err := mem.VirtualMemory()
	if err != nil {
		return AllResponse{}, err
	}

	// disk
	var ioCounters []disk.IOCountersStat
	part, _ := disk.Partitions(false)
	var uniqueDeviceNames []string

	for _, p := range part {
		counters, _ := disk.IOCounters(p.Device)
		if !contains(uniqueDeviceNames, p.Device) && len(counters[filepath.Base(p.Device)].Name) != 0 {
			ioCounters = append(ioCounters, counters[filepath.Base(p.Device)])
			uniqueDeviceNames = append(uniqueDeviceNames, p.Device)
		}
	}

	n, err := net.IOCounters(false)
	if err != nil {
		return AllResponse{}, err
	}

	hostname, _ := os.Hostname()

	return AllResponse{
		Hostname:   hostname,
		Disk:       ioCounters,
		Cpu:        ci[0],
		MemPercent: memory.UsedPercent,
		Net:        n,
	}, nil
}
