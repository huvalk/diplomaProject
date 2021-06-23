package usecase

import (
	"diplomaProject/application/auth"
	"diplomaProject/application/event"
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"diplomaProject/pkg/globalVars"
	"diplomaProject/pkg/oauth"
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	url2 "net/url"
	"regexp"
	"strconv"
	"strings"
)

type UseCase struct {
	auths         auth.Repository
	users         user.Repository
	events         event.Repository
	rx *regexp.Regexp
}

func NewUsecase(a auth.Repository, u user.Repository, e event.Repository) auth.UseCase {
	r, _ := regexp.Compile("\\/\\d*$")
	return &UseCase{
		auths: a,
		users: u,
		events: e,
		rx: r,
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

const defaultTitle = "Team Up"
const defaultDescription = "Найди лучшую команду"
const defaultImage = "https://team-up.online/hhton/favicon.ico"
const defaultMeta = "<title>%s</title>" +
	"<meta name=\"description\" content=\"%s\">" +
	"<meta name=\"keywords\" content=\"команда, участники, соревнования, навыки, поиск, хакатон, хак, team, teammates\">" +
	"<meta property=\"og:title\" content=\"%s\">" +
	"<meta property=\"og:description\" content=\"%s\">" +
	"<meta property=\"og:image\" content=\"%s\">" +
	"<meta property=\"og:site_name\" content=\"Team-up.online\">"

func (u *UseCase) GenerateMeta(url string) (string, error) {
	print()
	if strings.Contains(url, "event") {
		id, err := strconv.Atoi(u.rx.FindString(url)[1:])
		if err != nil {
			return "", err
		}

		e, err := u.events.Get(id)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(defaultMeta, e.Name, e.Description, e.Name, e.Description, e.Logo), nil
	} else if strings.Contains(url, "user") {
		id, err := strconv.Atoi(u.rx.FindString(url)[1:])
		if err != nil {
			return "", err
		}

		u, err := u.users.GetByID(id)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(defaultMeta, u.FirstName + " " + u.LastName, "Работает в " + u.WorkPlace + ". " + u.Bio,
			u.FirstName + " " + u.LastName, "Работает в " + u.WorkPlace + ". " + u.Bio, u.Avatar), nil
	} else {
		return fmt.Sprintf(defaultMeta, defaultTitle, defaultDescription, defaultTitle, defaultDescription, defaultImage), nil
	}
}