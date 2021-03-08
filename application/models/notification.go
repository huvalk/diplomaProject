package models

import (
	"diplomaProject/pkg/channel"
)

type Notification struct {
	UserID  int    `json:"userID"`
	Type    int    `json:"type"`
	Message string `json:"message"`
	Status string `json:"status"`
}

//easyjson:json
type NotificationArr []Notification

func NotificationFromChannel(n *channel.Notification) *Notification {
	return &Notification {
		UserID:  n.UserID,
		Type:    n.Type,
		Message: n.Message,
		Status: n.Status,
	}
}

func NotificationFromChannelArr(ns []channel.Notification) NotificationArr {
	arr := NotificationArr{}
	for _, n := range ns {
		arr = append(arr, Notification {
			UserID:  n.UserID,
			Type:    n.Type,
			Message: n.Message,
			Status: n.Status,
		})
	}

	return arr
}

func NotificationToChannel(n *Notification) *channel.Notification {
	return &channel.Notification{
		ID: 0,
		Type:    n.Type,
		Message: n.Message,
		UserID:  n.UserID,
		Status: n.Status,
	}
}

func NotificationToChannelArr(ns NotificationArr) (arr []channel.Notification) {
	arr = []channel.Notification{}
	for _, n := range ns {
		arr = append(arr, channel.Notification {
			ID: 0,
			UserID:  n.UserID,
			Type:    n.Type,
			Message: n.Message,
			Status: n.Status,
		})
	}

	return arr
}