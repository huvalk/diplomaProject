package usecase

import (
	"diplomaProject/application/auth"
	"diplomaProject/application/models"
	"diplomaProject/pkg/globalVars"
	"diplomaProject/pkg/oauth"
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	url2 "net/url"
)

type UseCase struct {
	auths         auth.Repository
}

func NewUsecase(a auth.Repository) auth.UseCase {
	return &UseCase{
		auths: a,

	}
}

func (u *UseCase) MakeAuthUrl(backTo string) string {
	return oauth.VkOAuthURL(globalVars.CLIENT_ID, globalVars.BACKEND_URI+"auth?backTo="+backTo, globalVars.STATE)
}

func (u *UseCase) UpdateUserInfo(code string, state string, backTo string) (int, error) {
	if state == "" || state != globalVars.STATE {
		return 0, errors.New("state doesnt match")
	}
	token, err := oauth.RetrieveUserToken(code, globalVars.CLIENT_ID,
		globalVars.BACKEND_URI+"auth?backTo="+url2.QueryEscape(backTo),
		globalVars.CLIENT_SECRET)
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
		Avatar:    userResponse.PhotoMaxOrig,
		Email:     token.Email,
		Vk:        userResponse.ScreenName,
	}
	if len(user.Vk) == 0 || len(user.Vk) > 50 {
		user.Vk = fmt.Sprintf("id%d", user.Id)
	}

	return user.Id, u.auths.UpdateUserInfo(user)
}
