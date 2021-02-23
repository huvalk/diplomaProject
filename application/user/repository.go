package user

import "diplomaProject/application/models"

type Repository interface {
	GetByID(uid int) (*models.User, error)
	GetByName(name string) (*models.User, error)
	JoinEvent(uid, evtID int) error
}
