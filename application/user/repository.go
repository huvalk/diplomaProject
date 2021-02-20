package user

import "diplomaProject/application/models"

type Repository interface {
	GetByID(uid int) (*models.VkUser, error)
	GetByName(name string) (*models.VkUser, error)
	JoinEvent(uid, evtID int) error
}
