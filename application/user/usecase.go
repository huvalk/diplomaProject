package user

import "diplomaProject/application/models"

type UseCase interface {
	Get(uid int) (*models.VkUser, error)
	Login(username string, password string) (string, string, error)
	JoinEvent(uid, evtID int) error
}
