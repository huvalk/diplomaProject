package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"diplomaProject/application/team"
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
	teams team.Repository
	channel       channel.Channel
}

func NewNotificationUsecase(n notification.Repository, t team.Repository) notification.UseCase {
	ch := channel.NewChannel()
	go ch.Run()

	return &NotificationUseCase{
		channel:       ch,
		teams: t,
		notifications: n,
	}
}

//func (n *NotificationUseCase) SetChannel(channel channel.Channel) {
//	n.channel = channel
//	go n.channel.Run()
//}

func (n *NotificationUseCase) SendInviteNotification(inv models.Invitation) (err error) {
	message := "Оповещение о приглашении"
	userTeam, err := n.teams.GetTeamByUser(inv.GuestID, inv.EventID)

	var teammates models.UserArr
	if userTeam != nil && err == nil {
		teammates, err = n.teams.GetTeamMembers(userTeam.Id)

		if err != nil {
			return err
		}
	} else {
		teammates = append(teammates, models.User{Id: inv.GuestID})
	}

	for _, user := range teammates {
		newNot := &channel.Notification{
			Type:    0,
			Message: message,
			UserID:  user.Id,
			Created: time.Time{},
			Status: "good",
			Watched: false,
		}

		newNot.Watched, err = n.channel.SendNotification(newNot)
		if err == nil {
			return err
		}

		err = n.notifications.SaveNotification(newNot)
		if err == nil {
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