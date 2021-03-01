package models

import (
	"diplomaProject/pkg/notification"
)

type Notification struct {
	UserID  int    `json:"userID"`
	Type    int    `json:"type"`
	Message string `json:"message"`
}

//easyjson:json
type NotificationArr []Notification

func FromChannel(n *notification.Notification) Notification {
	return Notification{
		UserID:  n.UserID,
		Type:    n.Type,
		Message: n.Message,
	}
}
