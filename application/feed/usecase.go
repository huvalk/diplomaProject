package feed

import "diplomaProject/application/models"

type UseCase interface {
	Get(feedID int) (*models.Feed, error)
	GetByEvent(eventID int) (*models.Feed, error)
	Create(eventID int) (*models.Feed, error)
	AddUser(uid, eventID int) error
	RemoveUser(uid, eventID int) error
	FilterFeed(eventID int, params map[string][]string) (*models.Feed, error)
}
