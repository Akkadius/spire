package cmd

import (
	"fmt"
	"github.com/Akkadius/spire/internal/backup"
	"github.com/k0kubun/pp/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"time"
)

type HelloWorldCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
	backup  *backup.Mysql
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(
	db *gorm.DB,
	logger *logrus.Logger,
	backup *backup.Mysql,
) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "hello:hello-world",
			Short: "Says hello world",
		},
		backup: backup,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, args []string) {

	start := time.Now()

	v, _ := mem.VirtualMemory()
	pp.Printf("Total: %v, Free:%v, UsedPercent:%v\n", v.Total, v.Free, v.UsedPercent)

	ci, _ := cpu.Percent(0, true)
	pp.Println(ci)

	//cpuInfo, _ := cpu.Info()
	//pp.Println(cpuInfo)

	n, _ := net.IOCounters(false)
	pp.Println(n)

	//conn, _ := net.Connections("all")
	//pp.Println(conn)

	h, _ := host.Info()
	pp.Println(h)

	//part, _ := disk.Partitions(true)
	//pp.Println(part)

	processes, _ := process.Processes()
	for _, p := range processes {
		pp.Println(p.Pid)
		pp.Printf("pid [%v]\n", p.Pid)
		n, _ := p.Name()
		pp.Printf("name [%v]\n", n)
		cli, _ := p.Cmdline()
		pp.Printf("cmd [%v]\n", cli)
		per, _ := p.CPUPercent()
		pp.Printf("CPU Percent %v\n", per)
		mper, _ := p.MemoryPercent()
		pp.Printf("Memory Percent %v\n", mper)
		pp.Println("IO Counters")
		pp.Println(p.IOCounters())
		fmt.Println("")
	}

	fmt.Printf("Read operation took %v\n", time.Since(start))
}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
