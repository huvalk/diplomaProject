package notification

import (
	"diplomaProject/pkg/channel"
)

type Repository interface {
	SaveNotification(n *channel.Notification) error
	MarkAsWatched(notificationID int) error
	GetPendingNotification(userID int) ([]channel.Notification, error)
}
