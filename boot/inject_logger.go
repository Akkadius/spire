package boot

import (
	"github.com/Akkadius/spire/internal/logger"
	"github.com/google/wire"
)

// wire set for loading the stores.
var loggerSet = wire.NewSet(
	logger.ProvideAppLogger,
)
