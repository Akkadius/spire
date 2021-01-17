package boot

import (
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/questapi"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	connection.NewDbConnectionCreateService,
	connection.NewDbConnectionCheckService,
	questapi.NewParseService,
)
