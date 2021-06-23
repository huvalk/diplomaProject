package channel

import (
	"github.com/gorilla/websocket"
	"time"
)

//easyjson:json
type Notification struct {
	ID int `json:"-"`
	// id евента
	Type string `json:"type,omitempty"`
	// Изменение произошло в
	Status  string    `json:"status,omitempty"`
	Message string    `json:"message"`
	UserID  int       `json:"-"`
	Created time.Time `json:"-"`
	Watched bool      `json:"-"`
}

//easyjson:skip
type ConnectedUser struct {
	ID     int
	Socket *websocket.Conn
	Send   chan []byte
	Chan   Channel
}
