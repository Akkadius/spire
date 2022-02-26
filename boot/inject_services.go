package boot

import (
	"github.com/Akkadius/spire/internal/clientfiles"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/desktop"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/internal/questapi"
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
)
