package notification

import (
	"github.com/gorilla/websocket"
	"time"
)

type Notification struct {
	Type    uint64    `json:"type,omitempty"`
	Message string    `json:"message"`
	UserID  uint64    `json:"userID,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

type ChannelUser struct {
	ID     uint64
	Socket *websocket.Conn
	Send   chan []byte
	Chan   Channel
}