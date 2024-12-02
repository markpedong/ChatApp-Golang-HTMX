package internal

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

type Manager struct {
	ClientList             []*Client
	ClientListEventChannel chan *ClientListEvent
}

type ClientListEvent struct {
	EventType string
	Client    *Client
}

var (
	upgrader = websocket.Upgrader{}
)

func NewManager() *Manager {
	return &Manager{
		ClientList:             []*Client{},
		ClientListEventChannel: make(chan *ClientListEvent),
	}
}

func (m *Manager) HandleClientListEventChannel(ctx context.Context) {
	for {
		select {
		case clientListEvent, ok := <-m.ClientListEventChannel:
			if !ok {
				return
			}
			switch clientListEvent.EventType {
			case "ADD":
				for _, v := range m.ClientList {
					if v.ID == clientListEvent.Client.ID {
						return
					}
				}

				m.ClientList = append(m.ClientList, clientListEvent.Client)
			case "REMOVE":
				newSlice := []*Client{}
				for _, v := range m.ClientList {
					if v.ID == clientListEvent.Client.ID {
						continue
					}
					newSlice = append(newSlice, v)
				}
				m.ClientList = newSlice
			}
		case <-ctx.Done():
			return
		}
	}
}

func (m *Manager) Handle(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	newClient := NewClient(ws, m)

	m.ClientListEventChannel <- &ClientListEvent{
		EventType: "ADD",
		Client:    newClient,
	}

	go newClient.ReadMessages(r)
	go newClient.WriteMessages(r)
}
