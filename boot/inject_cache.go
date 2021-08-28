package boot

import (
	"github.com/google/wire"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

// wire set for cache
var cacheSet = wire.NewSet(
	provideCache,
)

// provides cache
func provideCache() *gocache.Cache {
	return gocache.New(5*time.Minute, 10*time.Minute)
}
