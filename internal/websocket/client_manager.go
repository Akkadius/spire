package websocket

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/logger"
	"golang.org/x/net/websocket"
	"sync"
)

// Client represents a WebSocket client
type Client struct {
	ID string
	WS *websocket.Conn
}

// ClientManager manages all WebSocket clients
type ClientManager struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	lock       sync.Mutex

	// logger
	logger     *logger.AppLogger
	registered map[string]bool
}

// NewClientManager creates a new ClientManager
func NewClientManager(logger *logger.AppLogger) *ClientManager {
	return &ClientManager{
		clients:    make(map[*Client]bool),
		registered: make(map[string]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
		logger:     logger,
	}
}

// Start the manager to handle clients
func (m *ClientManager) Start() {
	for {
		select {
		case client := <-m.register:
			// check if client is already registered
			if _, ok := m.registered[client.ID]; ok {
				m.logger.DebugVvv().Any("client.ID", client.ID).Msg("Client already registered")
				continue
			}

			m.lock.Lock()
			m.clients[client] = true
			m.registered[client.ID] = true
			m.lock.Unlock()
			m.logger.Info().Any("client.ID", client.ID).Msg("Client registered")

		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				m.lock.Lock()
				delete(m.clients, client)
				m.lock.Unlock()
				m.logger.Info().Any("client.ID", client.ID).Msg("Client unregistered")
			}
		case message := <-m.broadcast:
			for client := range m.clients {
				if err := websocket.Message.Send(client.WS, message); err != nil {
					client.WS.Close()
					m.unregister <- client
				}
			}
		}
	}
}

type Message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Broadcast a message to all clients
func (m *ClientManager) Broadcast(msgType string, message string) {
	jsonData, err := json.Marshal(Message{
		Type:    msgType,
		Message: message,
	})
	if err != nil {
		m.logger.Debug().Any("error", err.Error()).Msg("Error marshalling message")
		return
	}

	m.broadcast <- string(jsonData)
}
