package auth

import "diplomaProject/application/models"

type Repository interface {
	UpdateUserInfo(user *models.User) error
}
