package notification

import (
	"diplomaProject/application/models"
	"github.com/gorilla/websocket"
)

type UseCase interface {
	SendInviteNotificationToUser(userID int, message string) error
	SendInviteNotificationToTeamByUser(userID int, message string) error
	SendPendingNotification(userID int) error
	GetPendingNotification(userID int) (models.NotificationArr, error)
	EnterChannel(userID int, socket *websocket.Conn) error
}
