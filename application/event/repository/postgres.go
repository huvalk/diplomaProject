package repository

import (
	"diplomaProject/application/event"
	"diplomaProject/application/models"
	"diplomaProject/pkg/infrastructure"
	"errors"
	"github.com/jinzhu/gorm"
)

type EventDatabase struct {
	conn *gorm.DB
}

func NewEventDatabase(db *gorm.DB) event.Repository {
	return &EventDatabase{conn: db}
}

func (e EventDatabase) Get(id int) (*models.Event, error) {
	for ind := range infrastructure.MockEvents {
		if infrastructure.MockEvents[ind].Id == int64(id) {
			return &infrastructure.MockEvents[ind], nil
		}
	}
	return &models.Event{}, errors.New("event not found")
}

func (e EventDatabase) Create(newEvent *models.Event) error {
	newEvent.Id = int64(len(infrastructure.MockEvents))
	infrastructure.MockEvents = append(infrastructure.MockEvents, *newEvent)
	return nil
}
