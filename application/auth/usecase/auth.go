package usecase

import (
	"diplomaProject/application/auth"
	"diplomaProject/pkg/oauth"
	"errors"
	"github.com/kataras/golog"
	"os"
)

type UseCase struct {
	ClientID     string
	RedirectURL  string
	State        string
	ClientSecret string
}

func NewUsecase() auth.UseCase {
	return &UseCase{
		ClientID:     os.Getenv("CLIENT_ID"),
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		State:        os.Getenv("STATE"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
}

func (u *UseCase) MakeAuthUrl() string {
	return oauth.VkOAuthURL(u.ClientID, u.RedirectURL, u.State)
}

func (u *UseCase) UpdateUserInfo(code string, state string) error {
	if state == "" || state != u.State {
		return errors.New("state doesnt match")
	}
	token, err := oauth.RetrieveUserToken(code, u.ClientID, u.RedirectURL, u.ClientSecret)
	if err != nil {
		return err
	}

	user, err := oauth.RetrieveProfileInfo(token)
	golog.Error(user)

	return err
}


