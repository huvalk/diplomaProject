package user

import "diplomaProject/application/models"

type UseCase interface {
	Get(uid int) (*models.User, error)
	GetForFeed(uid int) (*models.FeedUser, error)
	Login(username string, password string) (string, string, error)
	JoinEvent(uid, evtID int) error
	LeaveEvent(uid, evtID int) error
	GetUserEvents(uid int) (*models.EventArr, error)
}
