package event

import "diplomaProject/application/models"

type Repository interface {
	Get(id int) (*models.Event, error)
	Create(newEvent *models.Event) (*models.Event, error)
}
