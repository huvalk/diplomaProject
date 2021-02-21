package usecase

import (
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"diplomaProject/pkg/crypto"
	"errors"
	"github.com/google/uuid"
)

type User struct {
	users user.Repository
	feeds feed.Repository
}

func NewUser(u user.Repository, f feed.Repository) user.UseCase {
	return &User{users: u, feeds: f}
}

func (u *User) Get(uid int) (*models.VkUser, error) {
	return u.users.GetByID(uid)
}

func (u *User) JoinEvent(uid, evtID int) error {
	err := u.users.JoinEvent(uid, evtID)
	if err != nil {
		return err
	}
	return u.feeds.AddUser(uid, evtID)
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
