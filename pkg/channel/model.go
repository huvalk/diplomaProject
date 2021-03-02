package channel

import (
	"github.com/gorilla/websocket"
	"time"
)

type Notification struct {
	ID  int       `json:"ID,omitempty"`
	Type    int       `json:"type,omitempty"`
	Message string    `json:"message"`
	UserID  int       `json:"userID,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Watched bool      `json:"watched,omitempty"`
}

type ConnectedUser struct {
	ID     int
	Socket *websocket.Conn
	Send   chan []byte
	Chan   Channel
}
