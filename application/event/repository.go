package event

import "diplomaProject/application/models"

type Repository interface {
	Get(id int) (*models.EventDB, error)
	Create(newEvent *models.Event) (*models.EventDB, error)
	CheckUser(evtID, uid int) bool
}
