package usecase

import (
	"diplomaProject/application/auth"
	"diplomaProject/application/models"
	"diplomaProject/pkg/oauth"
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"os"
)

type UseCase struct {
	auths        auth.Repository
	ClientID     string
	RedirectURL  string
	State        string
	ClientSecret string
	ServiceSecret string
}

func NewUsecase(a auth.Repository) auth.UseCase {
	return &UseCase{
		auths:        a,
		ClientID:     os.Getenv("CLIENT_ID"),
		RedirectURL:  os.Getenv("BACKEND_URI"),
		State:        os.Getenv("STATE"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
}

func (u *UseCase) MakeAuthUrl() string {
	return oauth.VkOAuthURL(u.ClientID, u.RedirectURL + "auth", u.State)
}

func (u *UseCase) UpdateUserInfo(code string, state string) (int, error) {
	if state == "" || state != u.State {
		return 0, errors.New("state doesnt match")
	}
	token, err := oauth.RetrieveUserToken(code, u.ClientID, u.RedirectURL + "auth", u.ClientSecret)
	if err != nil {
		return 0, err
	}

	vk := api.NewVK(token.AccessToken)
	getUserInfoParams := api.Params{"fields": "screen_name, photo_max_orig"}
	response, err := vk.UsersGet(getUserInfoParams)

	if err != nil || len(response) == 0 {
		return 0, err
	}
	userResponse := response[0]
	user := &models.User{
		Id:        userResponse.ID,
		FirstName: userResponse.FirstName,
		LastName:  userResponse.LastName,
		Avatar: userResponse.PhotoMaxOrig,
		Email:     token.Email,
		Vk:     userResponse.ScreenName,
	}
	if user.Vk == "" {
		user.Vk = fmt.Sprintf("id%d", user.Id)
	}

	return user.Id, u.auths.UpdateUserInfo(user)
}
