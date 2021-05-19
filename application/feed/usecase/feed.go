package usecase

import (
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"math/rand"
	"time"
)

type Feed struct {
	feeds feed.Repository
}

func NewFeed(f feed.Repository) feed.UseCase {
	return &Feed{feeds: f}
}

func (f Feed) Get(feedID int) (*models.Feed, error) {
	fd, err := f.feeds.Get(feedID)
	if err != nil {
		return nil, err
	}
	us, err := f.feeds.GetFeedUsers(feedID)
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(us), func(i, j int) { us[i], us[j] = us[j], us[i] })
	fd.Users = us
	return fd, nil
}

func (f Feed) GetByEvent(eventID int) (*models.Feed, error) {
	fd, err := f.feeds.GetByEvent(eventID)
	if err != nil {
		return nil, err
	}
	us, err := f.feeds.GetFeedUsers(fd.Id)
	if err != nil {
		return nil, err
	}
	fd.Users = us
	return fd, nil
}

func (f Feed) FilterFeed(eventID int, params map[string][]string) (*models.Feed, error) {
	fd, err := f.feeds.GetByEvent(eventID)
	if err != nil {
		return nil, err
	}
	if len(params["job"]) == 0 {
		us, err := f.feeds.GetFeedUsers(fd.Id)
		if err != nil {
			return nil, err
		}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(us), func(i, j int) { us[i], us[j] = us[j], us[i] })
		fd.Users = us
		return fd, nil
	}
	us, err := f.feeds.FilterFeedBySkills(fd.Id, params["job"][0], params["skills"])
	if err != nil {
		return nil, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(us), func(i, j int) { us[i], us[j] = us[j], us[i] })
	fd.Users = us
	return fd, nil
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
