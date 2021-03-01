package notification

import (
	"diplomaProject/application/models"
	"github.com/gorilla/websocket"
)

type UseCase interface {
	SendInviteNotificationTo(userID int, message string) error
	GetPendingNotification(userID int) (models.NotificationArr, error)
	EnterChannel(userID int, socket *websocket.Conn) error
}
