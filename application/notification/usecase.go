package notification

import (
	"diplomaProject/application/models"
	"github.com/gorilla/websocket"
)

type UseCase interface {
	SendInviteNotification(users []int) error
	SendDenyNotification(users []int) error
	SendPendingNotification(userID int) error
	GetPendingNotification(userID int) (models.NotificationArr, error)
	EnterChannel(userID int, socket *websocket.Conn) error
}
