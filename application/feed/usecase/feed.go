package usecase

import (
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
)

type Feed struct {
	feeds feed.Repository
}

func NewFeed(f feed.Repository) feed.UseCase {
	return &Feed{feeds: f}
}

func (f Feed) Get(feedID int) (*models.Feed, error) {
	return f.feeds.Get(feedID)
}

func (f Feed) GetByEvent(eventID int) (*models.Feed, error) {
	return f.feeds.GetByEvent(eventID)
}

func (f Feed) Create(eventID int) (*models.Feed, error) {
	return f.feeds.Create(eventID)
}

func (f Feed) AddUser(uid, eventID int) error {
	return f.feeds.AddUser(uid, eventID)
}

func (f Feed) RemoveUser(uid, eventID int) error {
	return f.feeds.RemoveUser(uid, eventID)
}
