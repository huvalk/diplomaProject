package user

import "diplomaProject/application/models"

type UseCase interface {
	Get(uid int) (*models.VkUser, error)
}
