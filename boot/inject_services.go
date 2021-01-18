package boot

import (
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/questapi"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	connection.NewDbConnectionCreateService,
	connection.NewDbConnectionCheckService,
	github.NewGithubSourceDownloader,
	questapi.NewParseService,
	questapi.NewQuestExamplesGithubSourcer,
)
