package channel

import (
	"github.com/gorilla/websocket"
	"github.com/kataras/golog"
	"sync"
	"time"
)

func (c *ConnectedUser) Read(wg *sync.WaitGroup) {
	var err error

	err = c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		golog.Error("Cannot set read deadline: ", err)
	}
	c.Socket.SetPongHandler(func(string) error {
		err = c.Socket.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			golog.Error("Cannot set read deadline: ", err)
		}
		return err
	})

	wg.Done()
	for {
		if _, msg, err := c.Socket.ReadMessage(); err == nil {
			golog.Infof("Read by %d: %s", c.ID, msg)
			if len(msg) != 0 {
				c.Chan.Forward(msg)
			} else {
				golog.Infof("Read empty array by %d", c.ID)
			}
		} else {
			break
		}
	}
	golog.Error("Read from socket terminated: ", err)

	err = c.Socket.Close()
	if err != nil {
		golog.Error("Socket closed with error: ", err)
	}
}

func (c *ConnectedUser) Write(wg *sync.WaitGroup) {
	var err error
	ticker := time.NewTicker(pingPeriod)

	wg.Done()
LOOP:
	for {
		select {
		case message, ok := <-c.Send:
			err = c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				golog.Error("Cannot set write deadline: ", err)
			}

			if !ok {
				err = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				break LOOP
			}

			golog.Errorf("Write by %d: %s", c.ID, message)
			err = c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				golog.Error("Cannot set write deadline: ", err)
			}

			if err = c.Socket.WriteMessage(websocket.TextMessage, message); err != nil {
				break LOOP
			}
		case <-ticker.C:
			err = c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				golog.Error("Cannot set write deadline: ", err)
			}

			if err := c.Socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				break LOOP
			}
		}
	}
	golog.Error("Write to socket terminated: ", err)

	ticker.Stop()
	err = c.Socket.Close()
	if err != nil {
		golog.Error("Socket closed with error: ", err)
	}
}
