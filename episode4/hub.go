package main

import (
	"log"
	"sync"
	"time"
)

type hub struct {
	// the mutex to protect connections
	connectionsMx sync.RWMutex

	// Registered connections.
	connections map[*connection]struct{}

	// Inbound messages from the connections.
	broadcast chan []byte

	logMx sync.RWMutex
	log   [][]byte
}

func newHub() *hub {
	return &hub{
		connectionsMx: sync.RWMutex{},
		broadcast:     make(chan []byte),
		connections:   make(map[*connection]struct{}),
	}
}

func (h *hub) addConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	h.connections[conn] = struct{}{}
}

func (h *hub) removeConnection(conn *connection) {
	h.connectionsMx.Lock()
	defer h.connectionsMx.Unlock()
	if _, ok := h.connections[conn]; ok {
		delete(h.connections, conn)
		close(conn.send)
	}
}

func (h *hub) run() {
	for {
		msg := <-h.broadcast
		for c := range h.connections {
			select {
			case c.send <- msg:

			case <-time.After(1 * time.Second): // bail after trying to send for 1 second
				log.Printf("shutting down connection %s", c)
				h.removeConnection(c)
			}
		}
	}
}
