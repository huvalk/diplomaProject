package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"diplomaProject/pkg/channel"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
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

func (n *NotificationUseCase) SendYouJoinTeamNotification(users []int, evtID int) (err error) {
	message := "Вы теперь в команде"

	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: message,
		Created: time.Time{},
		Status:  "NewTeamNotification",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendNewMemberNotification(users []int, evtID int) (err error) {
	message := "В команде новый участник"

	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: message,
		Created: time.Time{},
		Status:  "NewMembersNotification",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendYouKickedNotification(users []int, evtID int) (err error) {
	message := "Вас кикнули из команды"

	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: message,
		Created: time.Now(),
		Status:  "NewTeamNotification",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendInviteNotification(users []int, evtID int) (err error) {
	message := "У вас новое приглашение, проверьте свою команду"

	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: message,
		Created: time.Now(),
		Status:  "NewInviteNotification",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendDenyNotification(users []int, evtID int) error {
	message := "Похоже Вам кто-то отказал, проверьте свою команду"

	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: message,
		Created: time.Now(),
		Status:  "NewDenyNotification",
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
			err = n.notifications.MarkAsWatched(not.ID)

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

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		user.Write()
		wg.Done()
	}()
	go func() {
		user.Read()
		wg.Done()
	}()

	// Это плохо, потом переделаю
	time.Sleep(2 * time.Second)
	err := n.SendPendingNotification(userID)

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
