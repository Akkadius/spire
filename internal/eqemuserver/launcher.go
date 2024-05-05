package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	zoneProcessName        = "zone"
	worldProcessName       = "world"
	ucsProcessName         = "ucs"
	queryServProcessName   = "queryserv"
	loginServerProcessName = "loginserver"
)

type Launcher struct {
	logger       *logger.DebugLogger
	serverconfig *eqemuserverconfig.Config
	settings     *spire.Settings
	pathmgmt     *pathmgmt.PathManagement
	serverApi    *Client

	// properties
	configLastModified   time.Time
	isRunning            bool // this is set by the server config
	runSharedMemory      bool
	runLoginserver       bool
	runQueryServ         bool
	minZoneProcesses     int
	currentZoneDynamics  int
	currentZoneStatics   int
	staticZonesToBoot    []string
	currentOnlineStatics []string
	currentProcessCounts map[string]int
}

func NewLauncher(
	logger *logger.DebugLogger,
	serverconfig *eqemuserverconfig.Config,
	settings *spire.Settings,
	pathmgmt *pathmgmt.PathManagement,
	serverApi *Client,
) *Launcher {
	l := &Launcher{
		logger:               logger,
		serverconfig:         serverconfig,
		settings:             settings,
		pathmgmt:             pathmgmt,
		currentProcessCounts: make(map[string]int),
		serverApi:            serverApi,
	}
	l.Process()

	// temp
	l.Start()

	return l
}

// Process runs the process loop for the launcher
// It only does anything if the server has the launcher set to run
func (l *Launcher) Process() {
	go func() {
		for {
			now := time.Now()
			stat, err := os.Stat(l.pathmgmt.GetEQEmuServerConfigFilePath())
			if err != nil {
				fmt.Printf("Error getting server config file: %v\n", err)
			}

			l.logger.DebugVvv().
				Any("configStatTime", now.Sub(time.Now()).String()).
				Any("configLastModified", l.configLastModified).
				Any("isRunning", l.isRunning).
				Msg("Main launcher process loop")

			if stat.ModTime().After(l.configLastModified) {
				l.configLastModified = stat.ModTime()
				cfg := l.serverconfig.Get()
				if cfg.Spire.LauncherStart {
					l.logger.Debug().
						Msg("Starting server launcher")
				}

				l.isRunning = cfg.Spire.LauncherStart
				l.runSharedMemory = cfg.WebAdmin.Launcher.RunSharedMemory
				l.runLoginserver = cfg.WebAdmin.Launcher.RunLoginserver
				l.runQueryServ = cfg.WebAdmin.Launcher.RunQueryServ
				l.minZoneProcesses = cfg.WebAdmin.Launcher.MinZoneProcesses

				var staticZonesToBoot []string

				// filter duplicate definitions
				for _, z := range strings.Split(cfg.WebAdmin.Launcher.StaticZones, ",") {
					isInList := false
					for _, cz := range staticZonesToBoot {
						if cz == z {
							isInList = true
							break
						}
					}

					if !isInList {
						staticZonesToBoot = append(staticZonesToBoot, z)
					}
				}

				l.staticZonesToBoot = staticZonesToBoot

				l.logger.Debug().
					Any("configLastModified", l.configLastModified).
					Any("statTime", now.Sub(time.Now()).String()).
					Any("isRunning", l.isRunning).
					Any("runSharedMemory", l.runSharedMemory).
					Any("runLoginserver", l.runLoginserver).
					Any("runQueryServ", l.runQueryServ).
					Any("minZoneProcesses", l.minZoneProcesses).
					Any("staticZones", l.staticZonesToBoot).
					Msg("Detected server config change")
			}

			if l.isRunning {
				l.Supervisor()
			}

			time.Sleep(1 * time.Second)
		}
	}()
}

func (l *Launcher) Start() {
	// temp
	cfg := l.serverconfig.Get()
	cfg.Spire.LauncherStart = true
	l.serverconfig.Save(cfg)

}

func (l *Launcher) GetServerProcessNames() []string {
	return []string{
		zoneProcessName,
		worldProcessName,
		ucsProcessName,
		queryServProcessName,
		loginServerProcessName,
	}
}

// Supervisor is the main process loop for the server launcher
// It will monitor the server processes and restart them if they crash
// this only runs when the server is set to run the launcher
func (l *Launcher) Supervisor() {
	now := time.Now()

	// reset
	l.currentProcessCounts = make(map[string]int)
	l.currentOnlineStatics = []string{}
	l.currentZoneDynamics = 0
	l.currentZoneStatics = 0

	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				continue
			}

			l.logger.Debug().
				Any("error", err.Error()).
				Any("pid", p.Pid).
				Msg("Error getting process command line")
		}

		cwd, err := p.Cwd()
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				continue
			}
			if strings.Contains(err.Error(), "permission denied") {
				continue
			}
			l.logger.Debug().
				Any("error", err.Error()).
				Any("pid", p.Pid).
				Msg("Error getting process cwd")
		}

		exe, err := p.Exe()
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Any("pid", p.Pid).
				Msg("Error getting process exe")
		}

		baseProcessName := filepath.Base(exe)

		for _, s := range l.GetServerProcessNames() {
			if s == baseProcessName {
				l.logger.DebugVv().
					Any("pid", p.Pid).
					Any("baseProcessName", baseProcessName).
					Any("cmdline", cmdline).
					Msg("Supervisor - Found server process")

				if _, ok := l.currentProcessCounts[s]; !ok {
					l.currentProcessCounts[s] = 0
				}

				if s == zoneProcessName {
					// get arg
					// check if it's in the static zones
					// if it is, add it to the current online statics
					arg := strings.Split(cmdline, " ")
					if len(arg) > 1 {
						for _, z := range l.staticZonesToBoot {
							if z == arg[1] {
								// make sure it's not already in the list
								isInList := false
								for _, cz := range l.currentOnlineStatics {
									if cz == z {
										isInList = true
										break
									}
								}

								if !isInList {
									l.currentOnlineStatics = append(l.currentOnlineStatics, z)
								}
							}
						}
						l.currentZoneStatics++
					} else {
						l.currentZoneDynamics++
					}
				}

				l.currentProcessCounts[s]++
			}
		}

		l.logger.DebugVvv().
			Any("pid", p.Pid).
			Any("cwd", cwd).
			Any("exe", exe).
			Any("cmdline", cmdline).
			Msg("Supervisor - Checking process")
	}

	list, err := l.serverApi.GetZoneList()
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Msg("Error getting zone list")
	}

	zoneAssignedDynamics := 0
	zoneIdleDynamics := 0
	for _, z := range list.Data {
		if !z.IsStaticZone {
			if z.ZoneID == 0 {
				zoneIdleDynamics++
				continue
			}
			zoneAssignedDynamics++
		}
	}

	zoneDynamicsToBoot := l.minZoneProcesses - zoneIdleDynamics
	if zoneDynamicsToBoot > 0 {
		l.logger.Debug().
			Any("zoneDynamicsToBoot", zoneDynamicsToBoot).
			Msg("Supervisor - Booting dynamic zone(s)")

		for i := 0; i < zoneDynamicsToBoot; i++ {
			l.startServerProcess(zoneProcessName)
		}
	}

	// statics
	if len(l.staticZonesToBoot) > 0 {
		for _, z := range l.staticZonesToBoot {
			isInList := false
			for _, cz := range l.currentOnlineStatics {
				if cz == z {
					isInList = true
					break
				}
			}

			if !isInList {
				l.logger.Debug().
					Any("zone", z).
					Msg("Supervisor - Booting static zone")

				l.startServerProcess(zoneProcessName, z)
			}
		}
	}

	if l.currentProcessCounts[worldProcessName] == 0 {
		l.startServerProcess(worldProcessName)
	}

	if l.runLoginserver && l.currentProcessCounts[loginServerProcessName] == 0 {
		l.startServerProcess(loginServerProcessName)
	}

	if l.runQueryServ && l.currentProcessCounts[queryServProcessName] == 0 {
		l.startServerProcess(queryServProcessName)
	}

	l.logger.DebugVv().
		Any("took", time.Now().Sub(now).String()).
		Any("currentProcessCounts", l.currentProcessCounts).
		Any("currentZoneDynamics", l.currentZoneDynamics).
		Any("currentZoneStatics", l.currentZoneStatics).
		Any("zoneDynamicsToBoot", zoneDynamicsToBoot).
		Any("zoneIdleDynamics", zoneIdleDynamics).
		Any("zoneAssignedDynamics", zoneAssignedDynamics).
		Any("statics - staticZonesToBoot", l.staticZonesToBoot).
		Any("statics - staticZonesToBoot (count)", len(l.staticZonesToBoot)).
		Any("statics - currentOnlineStatics (count)", len(l.currentOnlineStatics)).
		Any("statics - currentOnlineStatics", l.currentOnlineStatics).
		Msg("Supervisor - Loop complete")
}
