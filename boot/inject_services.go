package boot

import (
	"github.com/Akkadius/spire/internal/assets"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/backup"
	"github.com/Akkadius/spire/internal/clientfiles"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/desktop"
	"github.com/Akkadius/spire/internal/eqemuanalytics"
	"github.com/Akkadius/spire/internal/eqemuchangelog"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/internal/occulus"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/Akkadius/spire/internal/questapi"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/telnet"
	"github.com/Akkadius/spire/internal/websocket"
	"github.com/gertd/go-pluralize"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	influx.NewClient,
	connection.NewDbConnectionCreateService,
	connection.NewDbConnectionCheckService,
	github.NewGithubSourceDownloader,
	questapi.NewParseService,
	questapi.NewQuestExamplesGithubSourcer,
	desktop.NewWebBoot,
	clientfiles.NewExporter,
	clientfiles.NewImporter,
	eqemuserverconfig.NewConfig,
	pathmgmt.NewPathManagement,
	permissions.NewService,
	pluralize.NewClient,
	auditlog.NewUserEvent,
	assets.NewSpireAssets,
	eqemuchangelog.NewChangelog,
	eqemuanalytics.NewReleases,
	spire.NewUserService,
	spire.NewSettings,
	spire.NewInit,
	occulus.NewProxy,
	occulus.NewProcessManagement,
	telnet.NewClient,
	eqemuserver.NewClient,
	backup.NewMysql,
	websocket.NewSpireHandler,
	eqemuserver.NewUpdater,
)
