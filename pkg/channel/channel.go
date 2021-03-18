package channel

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

type Instance struct {
	forwardChan  chan []byte
	joinChan     chan *ConnectedUser
	leaveChan    chan *ConnectedUser
	ChannelUsers map[int]map[*websocket.Conn]*ConnectedUser
}

func NewChannel() *Instance {
	return &Instance{
		forwardChan:  make(chan []byte),
		joinChan:     make(chan *ConnectedUser),
		leaveChan:    make(chan *ConnectedUser),
		ChannelUsers: make(map[int]map[*websocket.Conn]*ConnectedUser),
	}
}

func (r *Instance) Forward(input []byte) {
	r.forwardChan <- input
}

func (r *Instance) Join(character *ConnectedUser) {
	r.joinChan <- character
}

func (r *Instance) Leave(character *ConnectedUser) {
	r.leaveChan <- character
}

func (r *Instance) Run() {
	golog.Info("Running chat room")
	for {
		select {
		case chatter := <-r.joinChan:
			golog.Infof("New chat user in channel: %d", chatter.ID)
			if r.ChannelUsers[chatter.ID] == nil {
				r.ChannelUsers[chatter.ID] = make(map[*websocket.Conn]*ConnectedUser)
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

func (r *Instance) SendNotification(n *Notification) (send bool, err error) {
	receiverConnections, existReceiver := r.ChannelUsers[n.UserID]
	send = false

	if existReceiver {
		rawMessage, err := json.Marshal(n)
		if err == nil {
			for _, connection := range receiverConnections {
				select {
				case connection.Send <- rawMessage:
					send = true
				default:
					delete(r.ChannelUsers[connection.ID], connection.Socket)
					if len(r.ChannelUsers[connection.ID]) == 0 {
						delete(r.ChannelUsers, connection.ID)
					}
					close(connection.Send)
				}
			}
		} else {
			golog.Errorf("Broken message: %+v", n)
			return false, err
		}
	} else {
		golog.Errorf("receiver not connected: %+v", n)
	}

	return send, nil
}

func (r *Instance) HandleMessage(rawMessage []byte) {
	var message Notification
	err := json.Unmarshal(rawMessage, &message)

	if err != nil {
		golog.Errorf("error while parsing message: %s, \n err: %s", string(rawMessage), err.Error())
		return
	}

	_, _ = r.SendNotification(&message)
}
