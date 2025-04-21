package eqemuserver

import "github.com/shirou/gopsutil/v3/process"

// getMemoryUsage returns the memory usage of a process
func getMemoryUsage(p *process.Process) uint64 {
	memory, _ := p.MemoryInfo()
	return memory.RSS
}
