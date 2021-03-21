package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
)

type Event struct {
	events event.Repository
	feeds  feed.UseCase
}

func NewEvent(e event.Repository, f feed.UseCase) event.UseCase {
	return &Event{events: e, feeds: f}
}

func (e *Event) GetEventUsers(evtID int) (*models.UserArr, error) {
	return e.events.GetEventUsers(evtID)
}

func (e *Event) Get(id int) (*models.Event, error) {
	newEvent, err := e.events.Get(id)
	if err != nil {
		return nil, err
	}
	evt := &models.Event{}
	evt.Convert(*newEvent)
	fd, err := e.feeds.GetByEvent(newEvent.Id)
	if err != nil {
		return nil, err
	}
	evt.Feed = *fd

	return evt, nil
}

func (e *Event) Create(newEvent *models.Event) (*models.Event, error) {
	evt, err := e.events.Create(newEvent)
	if err != nil {
		return nil, err
	}
	newEvent.Convert(*evt)
	fd, err := e.feeds.Create(newEvent.Id)
	if err != nil {
		return nil, err
	}
	newEvent.Feed = *fd
	return newEvent, nil
}
