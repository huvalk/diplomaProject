package event

import (
	"diplomaProject/application/models"
)

type Repository interface {
	Get(id int) (*models.EventDB, error)
	UpdateEvent(evt *models.Event) error
	UpdatePrize(pr *models.Prize) error
	Finish(id int) error
	Create(newEvent *models.Event) (*models.EventDB, error)
	CheckUser(evtID, uid int) bool
	GetEventUsers(evtID int) (*models.UserArr, error)
	GetEventTeams(evtID int) (*models.TeamArr, error)
	GetEventWinnerTeams(evtID int) (*models.TeamWinnerArr, error)
	GetEventPrize(evtID int) (*models.PrizeArr, error)
	AddPrize(evtID int, prizeArr models.PrizeArr) error
	SelectWinner(prizeID, tId int) error
	SetLogo(uid, eid int, link string) error
	SetBackground(uid, eid int, link string) error
}
