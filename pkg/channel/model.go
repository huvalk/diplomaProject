package channel

import (
	"github.com/gorilla/websocket"
	"time"
)

//easyjson:json
type Notification struct {
	ID      int       `json:"ID,omitempty"`
	Type    string    `json:"type,omitempty"`
	Status  string    `json:"status,omitempty"`
	Message string    `json:"message"`
	UserID  int       `json:"userID,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Watched bool      `json:"watched,omitempty"`
}

//easyjson:skip
type ConnectedUser struct {
	ID     int
	Socket *websocket.Conn
	Send   chan []byte
	Chan   Channel
}
