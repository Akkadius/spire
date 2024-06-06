package eqemuserver

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/rfsnotify"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type QuestHotReloadWatcher struct {
	logger       *logger.AppLogger
	serverconfig *eqemuserverconfig.Config
	pathmgmt     *pathmgmt.PathManagement
	serverApi    *Client
	db           *database.Resolver

	// properties
	configLastModified time.Time
	isRunning          bool
	watcher            *rfsnotify.RWatcher
	zones              []models.Zone
}

func NewQuestHotReloadWatcher(
	logger *logger.AppLogger,
	serverconfig *eqemuserverconfig.Config,
	pathmgmt *pathmgmt.PathManagement,
	serverApi *Client,
	db *database.Resolver,
) *QuestHotReloadWatcher {
	l := &QuestHotReloadWatcher{
		logger:       logger,
		serverconfig: serverconfig,
		pathmgmt:     pathmgmt,
		serverApi:    serverApi,
		db:           db,
	}

	return l
}

func (l *QuestHotReloadWatcher) Run() {
	if !env.IsAppEnvLocal() {
		return
	}

	l.loadZones()

	go func(l *QuestHotReloadWatcher) {
		l.Process()
	}(l)
}

// Process runs the process loop for the QuestHotReloadWatcher
// It only does anything if the server has the QuestHotReloadWatcher set to run
func (l *QuestHotReloadWatcher) Process() {
	for {
		now := time.Now()
		stat, err := os.Stat(l.pathmgmt.GetEQEmuServerConfigFilePath())
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Any("isRunning", l.isRunning).
				Msg("Error getting server config file stat")
		}

		l.logger.DebugVvv().
			Any("configStatTime", time.Now().Sub(now).String()).
			Any("configLastModified", l.configLastModified).
			Any("isRunning", l.isRunning).
			Msg("Main QuestHotReloadWatcher process loop")

		if stat.ModTime().After(l.configLastModified) {
			l.configLastModified = stat.ModTime()
			l.loadServerConfig()
			l.logger.Info().Any("watching", l.pathmgmt.GetQuestsDir()).Msg("Detected server config change, reloading watcher")
			l.logger.Debug().
				Any("configLastModified", l.configLastModified).
				Any("statTime", time.Now().Sub(now).String()).
				Any("isRunning", l.isRunning).
				Msg("Detected server config change")
		}

		if l.isRunning && l.watcher == nil {
			go func(l *QuestHotReloadWatcher) {
				l.Start()
			}(l)
		}

		if !l.isRunning && l.watcher != nil {
			l.Stop()
		}

		time.Sleep(processLoopTimer)
	}
}

// loadServerConfig loads the server config
// this is called on startup and when the server config changes
func (l *QuestHotReloadWatcher) loadServerConfig() {
	cfg := l.serverconfig.Get()
	l.isRunning = cfg.WebAdmin.Quests.HotReload
}

func (l *QuestHotReloadWatcher) Stop() {
	l.logger.Info().Any("watching", l.pathmgmt.GetQuestsDir()).Msg("Stopping Quest Hot Reload Watcher")

	if l.watcher != nil {
		err := l.watcher.Close()
		if err != nil {
			l.logger.Debug().
				Any("error", err.Error()).
				Msg("Error closing watcher")
		} else {
			l.watcher = nil
		}
	}
}

func (l *QuestHotReloadWatcher) Start() {
	var err error

	// Create new watcher.
	l.watcher, err = rfsnotify.NewWatcher()
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Msg("Error creating watcher")
		return
	}

	defer l.watcher.Close()

	var fileSizes map[string]int

	done := make(chan bool)

	// Start listening for events.
	go func(l *QuestHotReloadWatcher) {
		for {
			if l.watcher == nil {
				return
			}

			select {
			case event, ok := <-l.watcher.Events:
				if !ok {
					l.logger.Info().Any("watching", l.pathmgmt.GetQuestsDir()).Msg("Watcher error, closing watcher")
					if l.watcher != nil {
						l.watcher.Close()
					}
					l.watcher = nil
					return
				}

				// if event is create or write
				if event.Op == fsnotify.Create || event.Op == fsnotify.Write {
					if fileSizes == nil {
						fileSizes = make(map[string]int)
					}

					// initialize key if it doesn't exist
					if _, ok := fileSizes[event.Name]; !ok {
						fileSizes[event.Name] = 0
					}

					// get file size and store it
					fileInfo, err := os.Stat(event.Name)
					if err != nil {
						l.logger.Debug().Any("error", err.Error()).Msg("Error getting file info")
						break
					}

					if fileInfo.Size() == 0 {
						l.logger.Debug().Any("file", event.Name).Msg("File size is 0, skipping")
						break
					}

					if _, ok := fileSizes[event.Name]; ok {
						l.logger.Debug().
							Any("current", fileSizes[event.Name]).
							Any("new", int(fileInfo.Size())).
							Any("file", event.Name).
							Msg("File size comparison")
						if fileSizes[event.Name] == int(fileInfo.Size()) {
							l.logger.Debug().Any("file", event.Name).Msg("File size is the same, skipping reload")
							break
						}
					}

					fileSizes[event.Name] = int(fileInfo.Size())

					err = l.reloadQuestForFile(event.Name)
					if err != nil {
						l.logger.Debug().
							Any("error", err.Error()).
							Any("file", event.Name).
							Msg("Error reloading quest")
					}
				}

			case err, ok := <-l.watcher.Errors:
				if !ok {
					l.logger.Info().Any("watching", l.pathmgmt.GetQuestsDir()).Msg("Watcher error, closing watcher")
					if l.watcher != nil {
						l.watcher.Close()
					}
					l.watcher = nil
					return
				}
				log.Println("error:", err)
			}
		}
	}(l)

	l.logger.Info().Any("watching", l.pathmgmt.GetQuestsDir()).Msg("Starting Quest Hot Reload Watcher")

	// walk the path and add all files to the file sizes map
	err = filepath.Walk(l.pathmgmt.GetQuestsDir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if fileSizes == nil {
				fileSizes = make(map[string]int)
			}

			if _, ok := fileSizes[path]; !ok {
				fileSizes[path] = 0
			}

			fileSizes[path] = int(info.Size())
		}

		return nil
	})

	// Add a path.
	err = l.watcher.AddRecursive(l.pathmgmt.GetQuestsDir())
	if err != nil {
		l.logger.Debug().
			Any("error", err.Error()).
			Msg("Error adding path to watcher")
	}

	<-done
}

func (l *QuestHotReloadWatcher) reloadQuestForFile(name string) error {
	// /home/eqemu/server/quests/halas/zone_controller.pl
	// get the zone name from the path, traverse from the end
	parts := strings.Split(name, filepath.Join(string(os.PathSeparator)))
	folderName := parts[len(parts)-2]

	var zone models.Zone
	for _, z := range l.zones {
		if z.ShortName.String == folderName {
			zone = z
			break
		}
	}

	if zone.ID == 0 {
		// regex string match name against "lua_modules" "plugins" "global"
		// if match, reload all quests
		regex := "lua_modules|plugins|global"
		if matched, _ := regexp.MatchString(regex, name); matched {
			l.logger.Info().Any("file", name).Msg("Reloading all quests")
			err := l.serverApi.ReloadQuestsForZone("all")
			if err != nil {
				return err
			}
		}
	}

	if zone.ID > 0 {
		l.logger.Info().Any("file", name).Any("zone", zone.ShortName.String).Msg("Reloading quests")
		err := l.serverApi.ReloadQuestsForZone(zone.ShortName.String)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *QuestHotReloadWatcher) loadZones() {
	var results []models.Zone
	_ = l.db.GetEqemuDb().Find(&results).Error
	l.zones = results
}
