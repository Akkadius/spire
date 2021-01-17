package boot

import (
	"eoc/internal/connection"
	"eoc/questapi"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	connection.NewDbConnectionCreateService,
	connection.NewDbConnectionCheckService,
	questapi.NewParseService,
)
