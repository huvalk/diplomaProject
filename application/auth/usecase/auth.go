package usecase

import (
	"diplomaProject/application/auth"
	"diplomaProject/application/models"
	"diplomaProject/pkg/oauth"
	"errors"
	"os"
)

type UseCase struct {
	auths        auth.Repository
	ClientID     string
	RedirectURL  string
	State        string
	ClientSecret string
}

func NewUsecase(a auth.Repository) auth.UseCase {
	return &UseCase{
		auths:        a,
		ClientID:     os.Getenv("CLIENT_ID"),
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		State:        os.Getenv("STATE"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
}

func (u *UseCase) MakeAuthUrl() string {
	return oauth.VkOAuthURL(u.ClientID, u.RedirectURL, u.State)
}

func (u *UseCase) UpdateUserInfo(code string, state string) (int, error) {
	if state == "" || state != u.State {
		return 0, errors.New("state doesnt match")
	}
	token, err := oauth.RetrieveUserToken(code, u.ClientID, u.RedirectURL, u.ClientSecret)
	if err != nil {
		return 0, err
	}

	vkUser, err := oauth.RetrieveProfileInfo(token.AccessToken)
	if err != nil {
		return 0, err
	}
	user := &models.User{
		Id:        vkUser.Id,
		FirstName: vkUser.FirstName,
		LastName:  vkUser.LastName,
		Email:     token.Email,
	}

	return user.Id, u.auths.UpdateUserInfo(user)
}
