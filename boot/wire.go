//go:build wireinject
// +build wireinject

package boot

import (
	"github.com/google/wire"
)

func InitializeApplication() (App, error) {
	wire.Build(
		cacheSet,
		databaseSet,
		databaseResolverSet,
		encryptionSet,
		httpCrudControllerSet,
		serviceSet,
		httpSet,
		loggerSet,
		commandSet,
		NewApplication,
	)

	return App{}, nil
}
