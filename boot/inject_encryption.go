package boot

import (
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/google/wire"
)

var encryptionSet = wire.NewSet(
	encryption.NewEncrypter,
)
