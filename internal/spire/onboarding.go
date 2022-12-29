package spire

import "github.com/Akkadius/spire/internal/database"

type Onboarding struct {
	connections *database.Connections
}

func NewOnboarding(connections *database.Connections) *Onboarding {
	return &Onboarding{
		connections: connections,
	}
}

func (o Onboarding) InitSpire() {
	o.connections.SpireMigrate(false)
}
