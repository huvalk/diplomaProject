package notification

import (
	"diplomaProject/application/models"
	"github.com/gorilla/websocket"
)

type UseCase interface {
	// Оповещение о принятии в команду
	SendYouJoinTeamNotification(users []int, evtID int) error
	// Оповещение о новом тиммейте
	SendNewMemberNotification(users []int, evtID int) error
	// Оповещение об удалении из команды
	SendYouKickedNotification(users []int, evtID int) error
	// Оповещение о новом инвайте
	SendInviteNotification(users []int, evtID int) error
	// Оповещение о новом инвайте
	SendDenyNotification(userID []int, evtID int) error
	// Оповещение о новом голосе
	SendVoteNotification(userID []int, evtID int) error
	// Оповещение о новом тимлиде
	SendTeamLeadNotification(userID []int, evtID int) error
	// Оповещение об удалении инвайта
	SendUnInviteNotification(userID []int, evtID int) error
	// Безшумное оповещение
	SendSilentUpdateNotification(userID []int, evtID int) error
	GetLastNotification(userID int) (models.NotificationArr, error)
	GetPendingNotification(userID int) (models.NotificationArr, error)
	EnterChannel(userID int, socket *websocket.Conn) error
}
