package boot

import (
	"eoc/internal/encryption"
	"github.com/google/wire"
)

var encryptionSet = wire.NewSet(
	encryption.NewEncrypter,
)
