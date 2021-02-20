package repository

import (
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"diplomaProject/pkg/infrastructure"
	"errors"
	"github.com/jinzhu/gorm"
)

type UserDatabase struct {
	conn *gorm.DB
}

func NewUserDatabase(db *gorm.DB) user.Repository {
	return &UserDatabase{conn: db}
}

func (udb *UserDatabase) GetByID(uid int) (*models.VkUser, error) {
	for i := range infrastructure.Users {
		if int64(uid) == infrastructure.Users[i].Id {
			return &infrastructure.Users[i], nil
		}
	}
	return &models.VkUser{}, errors.New("user with that id not found")
}

func (udb *UserDatabase) GetByName(name string) (*models.VkUser, error) {
	for i := range infrastructure.Users {
		if name == infrastructure.Users[i].FirstName {
			return &infrastructure.Users[i], nil
		}
	}

	return &models.VkUser{}, errors.New("user with that name not found")
}
