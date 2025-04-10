package connections

import (
	"controle-grupo-danca/pkg/database"
	"sync"

	"github.com/uptrace/bun"
)

type Manager struct {
	Config      database.Config
	mu          sync.Mutex
	connections map[string]*bun.DB
}

func New(config database.Config) *Manager {
	return &Manager{
		connections: make(map[string]*bun.DB),
		Config:      config,
	}
}

func (m *Manager) Get(dbName string) *bun.DB {
	m.mu.Lock()

	defer m.mu.Unlock()

	if m.connections[dbName] == nil {
		cfg := m.Config
		cfg.DBName = dbName
		m.connections[dbName] = database.NewDB(cfg)
	}

	return m.connections[dbName]
}
