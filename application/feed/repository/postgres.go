package repository

import (
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"diplomaProject/pkg/infrastructure"
	"errors"
	"github.com/jinzhu/gorm"
)

type FeedDatabase struct {
	conn *gorm.DB
}

func NewFeedDatabase(db *gorm.DB) feed.Repository {
	return &FeedDatabase{conn: db}
}

func (f FeedDatabase) Get(feedID int) (*models.Feed, error) {
	for ind := range infrastructure.EventFeeds {
		if infrastructure.EventFeeds[ind].Id == feedID {
			return &infrastructure.EventFeeds[ind], nil
		}
	}
	return nil, errors.New("feed with that id not found")
}

func (f FeedDatabase) GetByEvent(eventID int) (*models.Feed, error) {
	for ind := range infrastructure.EventFeeds {
		if infrastructure.EventFeeds[ind].Event == eventID {
			return &infrastructure.EventFeeds[ind], nil
		}
	}
	return nil, errors.New("feed for that event not found")
}

func (f FeedDatabase) Create(eventID int) (*models.Feed, error) {
	infrastructure.EventFeeds = append(infrastructure.EventFeeds, models.Feed{
		Id:    len(infrastructure.EventFeeds) + 1,
		Users: nil,
		Event: eventID,
	})
	return &infrastructure.EventFeeds[len(infrastructure.EventFeeds)-1], nil
}

func (f FeedDatabase) AddUser(uid, eventID int) error {
	fd, err := f.GetByEvent(eventID)
	if err != nil {
		return err
	}
	for i := range infrastructure.Users {
		if uid == infrastructure.Users[i].Id {
			fd.Users = append(fd.Users, infrastructure.Users[i])
			return nil
		}
	}
	return errors.New("user not found")
}

func (f FeedDatabase) RemoveUser(uid, eventID int) error {
	fd, err := f.GetByEvent(eventID)
	if err != nil {
		return err
	}
	for ind := range fd.Users {
		if fd.Users[ind].Id == uid {
			fd.Users = append(fd.Users[:ind], fd.Users[ind:]...)
			return nil
		}
	}
	return errors.New("user with that id in that event not found")
}
