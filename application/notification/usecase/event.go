package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"github.com/gorilla/websocket"
)

type NotificationUseCase struct {
	notification notification.Repository
}

func NewNotification(n notification.Repository) notification.UseCase {
	return &NotificationUseCase{
		notification: n,
	}
}

func (n *NotificationUseCase) SendInviteNotificationTo(userID int, message string) error {
	return nil
}

func (n *NotificationUseCase) GetPendingNotification(userID int) (models.NotificationArr, error) {
	return nil, nil
}

func (n *NotificationUseCase) EnterChannel(userID int, socket *websocket.Conn) error {
	return nil
}

//func (e *Event) Get(id int) (*models.Event, error) {
//	newEvent, err := e.events.Get(id)
//	if err != nil {
//		return nil, err
//	}
//	evt := &models.Event{}
//	evt.Convert(*newEvent)
//	fd, err := e.feeds.GetByEvent(newEvent.Id)
//	if err != nil {
//		return nil, err
//	}
//	evt.Feed = *fd
//
//	return evt, nil
//}
//
//func (e *Event) Create(newEvent *models.Event) (*models.Event, error) {
//	evt, err := e.events.Create(newEvent)
//	if err != nil {
//		return nil, err
//	}
//	newEvent.Convert(*evt)
//	fd, err := e.feeds.Create(newEvent.Id)
//	if err != nil {
//		return nil, err
//	}
//	newEvent.Feed = *fd
//	return newEvent, nil
//}
