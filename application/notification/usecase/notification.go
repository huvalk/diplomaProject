package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"diplomaProject/pkg/channel"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kataras/golog"
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
	eventName, err := n.notifications.GetEventName(evtID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("%s | Вы теперь в команде", eventName)

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
	eventName, err := n.notifications.GetEventName(evtID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("%s | В команде новый участник", eventName)

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
	eventName, err := n.notifications.GetEventName(evtID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("%s | Вас кикнули из команды", eventName)

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
	eventName, err := n.notifications.GetEventName(evtID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("%s | У вас новое приглашение, проверьте свою команду", eventName)

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
	eventName, err := n.notifications.GetEventName(evtID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("%s | Похоже Вам кто-то отказал, проверьте свою команду", eventName)

	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: message,
		Created: time.Now(),
		Status:  "NewDenyNotification",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendVoteNotification(users []int, evtID int) error {
	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: "",
		Created: time.Now(),
		Status:  "NewVoteNotification",
		Watched: false,
	}
	return n.SendNotification(newNot, users)
}

func (n *NotificationUseCase) SendTeamLeadNotification(users []int, evtID int) error {
	eventName, err := n.notifications.GetEventName(evtID)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("%s | В команде новый тимлид", eventName)

	newNot := channel.Notification{
		Type:    fmt.Sprintf("%d",evtID),
		Message: message,
		Created: time.Now(),
		Status:  "NewTeamLeadNotification",
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

func (n *NotificationUseCase) GetLastNotification(userID int) (models.NotificationArr, error) {
	res, err := n.notifications.GetLastNotification(userID)
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

	var waitToClose sync.WaitGroup
	waitToClose.Add(2)
	var waitToSendPending sync.WaitGroup
	waitToSendPending.Add(2)
	go func() {
		user.Write(&waitToSendPending)
		waitToClose.Done()
	}()
	go func() {
		user.Read(&waitToSendPending)
		waitToClose.Done()
	}()

	// TODO Проверить
	//time.Sleep(2 * time.Second)
	waitToSendPending.Wait()
	err := n.SendPendingNotification(userID)
	golog.Error("Send success")

	if err != nil {
		return err
	}

	waitToClose.Wait()
	return nil
}
