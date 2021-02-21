package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"fmt"
)

type Event struct {
	events event.Repository
	feeds  feed.Repository
}

func NewEvent(e event.Repository, f feed.Repository) event.UseCase {
	return &Event{events: e, feeds: f}
}

func (e *Event) Get(id int) (*models.Event, error) {
	newEvent, err := e.events.Get(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(newEvent)
	fd, err := e.feeds.GetByEvent(int(newEvent.Id))
	if err != nil {
		return nil, err
	}
	fmt.Println(fd)
	newEvent.Feed = *fd

	return newEvent, nil
}

func (e *Event) Create(newEvent *models.Event) (*models.Event, error) {
	newEvent, err := e.events.Create(newEvent)
	if err != nil {
		return nil, err
	}
	fd, err := e.feeds.Create(int(newEvent.Id))
	if err != nil {
		return nil, err
	}
	newEvent.Feed = *fd

	return newEvent, nil
}
