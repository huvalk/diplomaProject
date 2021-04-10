package notification

import (
	"diplomaProject/pkg/channel"
)

type Repository interface {
	SaveNotification(n *channel.Notification) error
	GetEventName(eventID int) (name string, err error)
	MarkAsWatched(notificationID int) error
	GetLastNotification(userID int) ([]channel.Notification, error)
	GetPendingNotification(userID int) ([]channel.Notification, error)
}
