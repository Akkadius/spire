package spa

import (
	"github.com/Akkadius/spire/internal/env"
	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
	"strings"
)

type SpirePackagedSpaService struct {
	logger *logrus.Logger
	spa    *PackedSpaService
}

func (s SpirePackagedSpaService) Spa() *PackedSpaService {
	return s.spa
}

// Spire SPA vars
const (
	SpireBasePath      = ""
	SpireLocalBasePath = "../../../frontend/dist/"
	SpireSpaIndex      = "index.html"
)

// New up an instance of Spire
func NewSpirePackagedSpaService(logger *logrus.Logger) *SpirePackagedSpaService {
	// This is merely a no-op and simply informs the packr CLI utility what it needs to bundle
	// since it parses the code separately on its own to know where to bundle
	_ = packr.NewBox(SpireLocalBasePath)

	return &SpirePackagedSpaService{
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
