package repository

import (
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"github.com/jinzhu/gorm"
)

type UserDatabase struct {
	conn *gorm.DB
}

func NewUserDatabase(db *gorm.DB) user.Repository {
	return &UserDatabase{conn: db}
}

func (udb *UserDatabase) GetByID(uid int) (usr *models.VkUser, err error) {
	usr = &models.VkUser{
		Id:        int64(uid),
		FirstName: "2",
		LastName:  "3",
		Email:     "4",
	}

	return usr, nil
}
