package user

import (
	"diplomaProject/application/models"
	"diplomaProject/application/user"
)

type User struct {
	users user.Repository
	//logger *zap.Logger
}

func NewUser(u user.Repository) user.UseCase {
	return &User{users: u}
}

func (u *User) Get(uid int) (*models.VkUser, error) {
	return u.users.GetByID(uid)
}
