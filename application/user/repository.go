package user

import "diplomaProject/application/models"

type Repository interface {
	GetByID(uid int) (*models.VkUser, error)
}
