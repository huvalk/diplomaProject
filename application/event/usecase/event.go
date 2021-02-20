package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/models"
)

type Event struct {
	events event.Repository
}

func NewEvent(e event.Repository) event.UseCase {
	return &Event{events: e}
}

func (e *Event) Get(id int) (*models.Event, error) {
	return e.events.Get(id)
}

func (e *Event) Create(newEvent *models.Event) error {
	//TODO:create feed.
	return e.events.Create(newEvent)
}
