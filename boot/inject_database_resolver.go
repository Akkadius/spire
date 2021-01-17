package boot

import (
	"eoc/database"
	"github.com/google/wire"
)

// wire set for loading the stores.
var databaseResolverSet = wire.NewSet(
	database.NewDatabaseResolver,
)
