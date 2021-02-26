package notification

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/kataras/golog"
	"time"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type ChannelInstance struct {
	forwardChan  chan []byte
	joinChan     chan *ChannelUser
	leaveChan    chan *ChannelUser
	ChannelUsers map[uint64]map[*websocket.Conn]*ChannelUser
}

func NewChannel() *ChannelInstance {
	return &ChannelInstance{
		forwardChan:  make(chan []byte),
		joinChan:     make(chan *ChannelUser),
		leaveChan:    make(chan *ChannelUser),
		ChannelUsers: make(map[uint64]map[*websocket.Conn]*ChannelUser),
	}
}

func (r *ChannelInstance) Forward(input []byte) {
	r.forwardChan <- input
}

func (r *ChannelInstance) Join(character *ChannelUser) {
	r.joinChan <- character
}

func (r *ChannelInstance) Leave(character *ChannelUser) {
	r.leaveChan <- character
}

func (r *ChannelInstance) Run() {
	golog.Info("Running chat room")
	for {
		select {
		case chatter := <-r.joinChan:
			golog.Infof("New chat user in channel: %d", chatter.ID)
			if r.ChannelUsers[chatter.ID] == nil {
				r.ChannelUsers[chatter.ID] = make(map[*websocket.Conn]*ChannelUser)
			}
			r.ChannelUsers[chatter.ID][chatter.Socket] = chatter
		case chatter := <-r.leaveChan:
			golog.Infof("chat user channel: %d", chatter.ID)
			delete(r.ChannelUsers[chatter.ID], chatter.Socket)
			if len(r.ChannelUsers[chatter.ID]) == 0 {
				delete(r.ChannelUsers, chatter.ID)
			}
			close(chatter.Send)
		case rawMessage := <-r.forwardChan:
			r.HandleMessage(rawMessage)
		}
	}
}

func (r *ChannelInstance) SendNotification(n *Notification) error {
	receivers, existReceivers := r.ChannelUsers[n.UserID]
	if existReceivers {
		rawMessage, err := json.Marshal(n)
		if err == nil {
			for _, receiver := range receivers {
				select {
				case receiver.Send <- rawMessage:
				default:
					delete(r.ChannelUsers[receiver.ID], receiver.Socket)
					if len(r.ChannelUsers[receiver.ID]) == 0 {
						delete(r.ChannelUsers, receiver.ID)
					}
					close(receiver.Send)
				}
			}
		} else {
			golog.Errorf("Broken message: %+v", n)
			return err
		}
	} else {
		golog.Errorf("unknown receiver: %+v", n)
	}

	return nil
}

func (r *ChannelInstance) HandleMessage(rawMessage []byte) {
	var message *map[string]json.RawMessage
	err := json.Unmarshal(rawMessage, &message)

	if err != nil {
		golog.Errorf("error while parsing message: %s, \n err: %s", string(rawMessage), err.Error())
		return
	}

	golog.Warn("we could handle this message, but we didnt: %s", string(rawMessage))
}

func (c *ChannelUser) Read() {
	var err error

	err = c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		golog.Error("Cannot set read deadline: ", err)
	}
	c.Socket.SetPongHandler(func(string) error {
		err = c.Socket.SetReadDeadline(time.Now().Add(pongWait))
		golog.Error("Ping")
		if err != nil {
			golog.Error("Cannot set read deadline: ", err)
		}
		return err
	})

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

func (c *ChannelUser) Write() {
	var err error
	ticker := time.NewTicker(pingPeriod)

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
