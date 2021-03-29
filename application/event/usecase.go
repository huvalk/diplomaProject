package event

import "diplomaProject/application/models"

type UseCase interface {
	Get(id int) (*models.Event, error)
	Update(uID int, evt *models.Event) (*models.Event, error)
	Finish(uID, evtID int) (*models.Event, error)
	GetEventUsers(evtID int) (*models.UserArr, error)
	Create(newEvent *models.Event) (*models.Event, error)
	SelectWinner(uID, evtID, prizeID, tId int) error
}
