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
		if infrastructure.MockEvents[ind].Id == id {
			return &infrastructure.MockEvents[ind], nil
		}
	}
	return &models.Event{}, errors.New("event not found")
}

func (e EventDatabase) Create(newEvent *models.Event) (*models.Event, error) {
	newEvent.Id = len(infrastructure.MockEvents) + 1
	infrastructure.MockEvents = append(infrastructure.MockEvents, *newEvent)
	return newEvent, nil
}

func (e EventDatabase) CheckUser(evtID, uid int) bool {
	for ind := range infrastructure.EventMembers[evtID] {
		if infrastructure.EventMembers[evtID][ind] == uid {
			return true
		}
	}
	return false
}
