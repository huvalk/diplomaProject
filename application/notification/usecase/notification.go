package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"diplomaProject/pkg/channel"
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

const (
	messageBufferSize = 256
)

type NotificationUseCase struct {
	notifications notification.Repository
	channel       channel.Channel
}

func NewNotificationUsecase(n notification.Repository) notification.UseCase {
	ch := channel.NewChannel()
	go ch.Run()

	return &NotificationUseCase{
		channel:       ch,
		notifications: n,
	}
}

//func (n *NotificationUseCase) SetChannel(channel channel.Channel) {
//	n.channel = channel
//	go n.channel.Run()
//}

func (n *NotificationUseCase) SendInviteNotification(users []int) (err error) {
	message := "Оповещение о приглашении"

	for _, userID := range users {
		newNot := &channel.Notification{
			Type:    0,
			Message: message,
			UserID:  userID,
			Created: time.Time{},
			Status: "good",
			Watched: false,
		}

		newNot.Watched, err = n.channel.SendNotification(newNot)
		if err != nil {
			return err
		}

		err = n.notifications.SaveNotification(newNot)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *NotificationUseCase) GetPendingNotification(userID int) (models.NotificationArr, error) {
	res, err := n.notifications.GetPendingNotification(userID)
	if err != nil {
		return nil, err
	}

	return models.NotificationFromChannelArr(res), nil
}

func (n *NotificationUseCase) SendPendingNotification(userID int) error {
	res, err := n.notifications.GetPendingNotification(userID)
	if err != nil {
		return err
	}

	for _, not := range res {
		wasSent, err := n.channel.SendNotification(&not)

		if err != nil {
			return err
		}
		if wasSent {
			err = n.notifications.MarkAsWatched(not.UserID)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (n *NotificationUseCase) EnterChannel(userID int, socket *websocket.Conn) error {
	if n.channel == nil {
		return errors.New("in EnterChannel: no channel set")
	}

	user := &channel.ConnectedUser{
		ID:     userID,
		Socket: socket,
		Send:   make(chan []byte, messageBufferSize),
		Chan:   n.channel,
	}

	n.channel.Join(user)
	defer func() {
		n.channel.Leave(user)
	}()
	go user.Write()
	user.Read()

	return nil
}