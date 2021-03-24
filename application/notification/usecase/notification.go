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

func (n *NotificationUseCase) SendNotification(notification channel.Notification, users []int) (err error) {
	for _, userID := range users {
		notification.UserID = userID

		notification.Watched, err = n.channel.SendNotification(&notification)
		if err != nil {
			return err
		}

		err = n.notifications.SaveNotification(&notification)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *NotificationUseCase) SendJoinTeamNotification(users []int) (err error) {
	message := "Вы теперь в команде, проверьте"

	newNot := channel.Notification{
		Type:    "notification",
		Message: message,
		Created: time.Time{},
		Status:  "good",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendKickTeamNotification(users []int) (err error) {
	message := "Вас кикнули из команды, проверьте"

	newNot := channel.Notification{
		Type:    "notification",
		Message: message,
		Created: time.Now(),
		Status:  "bad",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendInviteNotification(users []int) (err error) {
	message := "У вас новое приглашение, проверьте"

	newNot := channel.Notification{
		Type:    "notification",
		Message: message,
		Created: time.Now(),
		Status:  "good",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendDenyNotification(users []int) error {
	message := "Похоже Вам кто-то отказал, проверьте"

	newNot := channel.Notification{
		Type:    "notification",
		Message: message,
		Created: time.Now(),
		Status:  "bad",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
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
	// TODO проверять ошибки
	_ = n.SendPendingNotification(userID)
	user.Read()

	return nil
}
