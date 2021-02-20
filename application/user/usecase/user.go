package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"diplomaProject/pkg/crypto"
	"errors"
	"github.com/google/uuid"
)

type User struct {
	users user.Repository
}

func NewUser(u user.Repository) user.UseCase {
	return &User{users: u}
}

func (u *User) Get(uid int) (*models.VkUser, error) {
	return u.users.GetByID(uid)
}

func (u *User) JoinEvent(uid, evtID int) error {
	return u.users.JoinEvent(uid, evtID)
}

func (u *User) Login(username string, password string) (sessionId string, csrfToken string, err error) {
	if username == "" || password == "" {
		return "", "", errors.New("(((")
	}

	_, err = u.users.GetByName(username)
	if err != nil {
		return "", "", err
	}

	//ok, err := crypto.CheckPassword(password, usr.Password)
	//if err != nil {
	//	return "", "", err
	//}
	//if !ok {
	//	return "", "", errors.NewInvalidArgument("Wrong password")
	//}

	sessionId = uuid.New().String()
	csrfToken = crypto.CreateToken(sessionId)
	return
	//err = uc.sessions.Add(sessionId, usr.Id)
}
