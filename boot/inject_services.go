package boot

import (
	"github.com/Akkadius/spire/internal/assets"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/clientfiles"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/desktop"
	"github.com/Akkadius/spire/internal/eqemuanalytics"
	"github.com/Akkadius/spire/internal/eqemuchangelog"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/Akkadius/spire/internal/questapi"
	"github.com/Akkadius/spire/internal/serverconfig"
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
	serverconfig.NewEQEmuServerConfig,
	pathmgmt.NewPathManagement,
	permissions.NewService,
	pluralize.NewClient,
	auditlog.NewUserEvent,
	assets.NewSpireAssets,
	eqemuchangelog.NewChangelog,
	eqemuanalytics.NewReleases,
)
