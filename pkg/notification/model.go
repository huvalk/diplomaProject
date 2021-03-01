package notification

import (
	"github.com/gorilla/websocket"
	"time"
)

type Notification struct {
	Type    int    `json:"type,omitempty"`
	Message string    `json:"message"`
	UserID  int    `json:"userID,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

type ChannelUser struct {
	ID     int
	Socket *websocket.Conn
	Send   chan []byte
	Chan   Channel
}