package user

import (
	"diplomaProject/application/models"
	"mime/multipart"
)

type UseCase interface {
	Get(uid int) (*models.User, error)
	Update(usr *models.User) (*models.User, error)
	SearchUserByTag(eid int, tag string) (models.UserArr, error)
	SetImage(uid int, avatar *multipart.Form) (string, error)
	GetForFeed(uid int) (*models.FeedUser, error)
	Login(username string, password string) (string, string, error)
	JoinEvent(uid, evtID int) error
	LeaveEvent(uid, evtID int) error
	GetUserEvents(uid int) (*models.EventArr, error)
	GetFounderEvents(userID int) (*models.EventDBArr, error)
	GetBDEvent(evtID int) (*models.EventDB, error)
}
