package event

import "diplomaProject/application/models"

type UseCase interface {
	Get(id int) (*models.Event, error)
	Create(newEvent *models.Event) error
}
