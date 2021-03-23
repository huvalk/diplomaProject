package user

import "diplomaProject/application/models"

type Repository interface {
	GetByID(uid int) (*models.User, error)
	GetByName(name string) (*models.User, error)
	Update(usr *models.User) (*models.User, error)
	SetImage(uid int, link string) error
	JoinEvent(uid, evtID int) error
	LeaveEvent(uid, evtID int) error
	GetUserEvents(uid int) (*models.EventArr, error)
	GetUserParams(uid int) (models.Job, []models.Skills, error)
}
