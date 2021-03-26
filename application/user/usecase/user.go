package usecase

import (
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"diplomaProject/application/user"
	"diplomaProject/pkg/crypto"
	"diplomaProject/pkg/sss"
	"errors"
	"github.com/google/uuid"
	"mime/multipart"
	"regexp"
)

type User struct {
	users     user.Repository
	feeds     feed.Repository
	teams     team.Repository
	tagRegexp *regexp.Regexp
}

func NewUser(u user.Repository, f feed.Repository, t team.Repository) user.UseCase {
	r, _ := regexp.Compile(`([a-zA-Z\\d])+$`)
	return &User{users: u, feeds: f, teams: t, tagRegexp: r}
}

func (u *User) SetImage(uid int, avatar *multipart.Form) (string, error) {
	link, err := sss.UploadPic(avatar, "")
	if err != nil {
		return "", err
	}
	err = u.users.SetImage(uid, link)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (u *User) Update(usr *models.User) (*models.User, error) {
	return u.users.Update(usr)
}

func (u *User) SearchUserByTag(eid int, tag string) (models.UserArr, error) {
	tagExtracted := u.tagRegexp.FindStringSubmatch(tag)
	if len(tagExtracted) < 1 || len(tagExtracted[0]) > 40 || len(tagExtracted[0]) == 0 {
		return nil, errors.New("no valid tag")
	}

	return u.users.SearchUserByTag(eid, tagExtracted[0])
}

func (u *User) GetForFeed(uid int) (*models.FeedUser, error) {
	usr, err := u.Get(uid)
	if err != nil {
		return nil, err
	}
	fdUsr := &models.FeedUser{}
	fdUsr.Convert(*usr)
	_, skills, err := u.users.GetUserParams(uid)
	if err != nil {
		return nil, err
	}
	//fdUsr.JobName = job.Name
	fdUsr.Skills = skills
	return fdUsr, err
}

func (u *User) Get(uid int) (*models.User, error) {
	return u.users.GetByID(uid)
}

func (u *User) JoinEvent(uid, evtID int) error {
	err := u.users.JoinEvent(uid, evtID)
	if err != nil {
		return err
	}
	return u.feeds.AddUser(uid, evtID)
}

func (u *User) LeaveEvent(uid, evtID int) error {
	err := u.users.LeaveEvent(uid, evtID)
	if err != nil {
		return err
	}
	err = u.feeds.RemoveUser(uid, evtID)
	if err != nil {
		return err
	}
	tm, err := u.teams.GetTeamByUser(uid, evtID)
	if err != nil {
		return err
	}
	return u.teams.RemoveMember(tm.Id, uid)
}

func (u *User) GetUserEvents(uid int) (*models.EventArr, error) {
	return u.users.GetUserEvents(uid)
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
	//err = uc.sessions.Add(sessionId, usr.ID)
}
