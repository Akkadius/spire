package spa

import (
	"github.com/Akkadius/spire/internal/env"
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
	"strings"
)

type Spa struct {
	logger *logrus.Logger
	spa    *Packer
}

func (s Spa) Spa() *Packer {
	return s.spa
}

// Spire SPA vars
const (
	SpireBasePath      = ""
	SpireLocalBasePath = "../../../frontend/dist/"
	SpireSpaIndex      = "index.html"
)

func NewSpa(logger *logrus.Logger) *Spa {
	// This is merely a no-op and simply informs the packr CLI utility what it needs to bundle
	// since it parses the code separately on its own to know where to bundle
	_ = packr.NewBox(SpireLocalBasePath)

	return &Spa{
		logger: logger,
		spa: NewPackedSpaService(
			logger, PackedSpaServeConfig{
				BasePath:      SpireBasePath,
				LocalBasePath: SpireLocalBasePath,
				SpaIndex:      SpireSpaIndex,
				SkipPaths:     strings.Split(env.Get("SPA_SKIP_PATH_PREFIXES", "/auth,/api,/swagger,/websocket"), ","),
			},
		),
	}
}
