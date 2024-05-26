package expansions

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) GetExpansions() []Expansion {
	return expansions
}
