package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"diplomaProject/pkg/constants"
	"errors"
)

type Event struct {
	events event.Repository
	feeds  feed.UseCase
}

func NewEvent(e event.Repository, f feed.UseCase) event.UseCase {
	return &Event{events: e, feeds: f}
}

func (e *Event) GetEventWinnerTeams(evtID int) (*models.TeamWinnerArr, error) {
	return e.events.GetEventWinnerTeams(evtID)
}

func (e *Event) GetEventTeams(evtID int) (*models.TeamArr, error) {
	return e.events.GetEventTeams(evtID)
}

func (e *Event) Update(uID int, evt *models.Event) (*models.Event, error) {
	if evt.Founder != uID {
		return nil, errors.New("not founder")
	}
	err := e.events.UpdateEvent(evt)
	if err != nil {
		return nil, err
	}
	for i := range evt.PrizeList {
		err = e.events.UpdatePrize(&evt.PrizeList[i])
		if err != nil {
			return nil, err
		}
	}
	return e.Get(evt.Id)
}

func (e *Event) SelectWinner(uID, evtID, PrizeID, tId int) error {
	ev, err := e.Get(evtID)
	if err != nil {
		return nil
	}
	if ev.Founder != uID {
		return errors.New("not owner")
	}
	if ev.State != constants.EventStatusClosed {
		return errors.New("not finished")
	}
	//TODO:check amount>=1

	//check team in event
	//add to team members prize
	return e.events.SelectWinner(PrizeID, tId)
}

func (e *Event) Finish(uID, evtID int) (*models.Event, error) {
	ev, err := e.Get(evtID)
	if err != nil {
		return nil, err
	}
	if ev.Founder != uID {
		return nil, errors.New("not owner")
	}
	err = e.events.Finish(evtID)
	if err != nil {
		return nil, err
	}
	ev.State = constants.EventStatusClosed
	return ev, nil
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
	prArr, err := e.events.GetEventPrize(id)
	if err != nil || len(*prArr) < 1 {
		evt.PrizeList = models.PrizeArr{}
	} else {
		evt.PrizeList = *prArr
	}
	evt.ParticipantsCount = len(fd.Users)
	return evt, nil
}

func (e *Event) Create(newEvent *models.Event) (*models.Event, error) {
	newEvent.State = constants.EventStatusOpen
	evt, err := e.events.Create(newEvent)
	if err != nil {
		return nil, err
	}
	newEvent.Id = evt.Id
	if newEvent.PrizeList != nil {
		err = e.AddPrize(evt.Id, newEvent.PrizeList)
		if err != nil {
			return nil, err
		}
	}
	fd, err := e.feeds.Create(newEvent.Id)
	if err != nil {
		return nil, err
	}
	newEvent.Feed = *fd
	return newEvent, nil
}

func (e *Event) AddPrize(evtID int, prizeArr models.PrizeArr) error {
	return e.events.AddPrize(evtID, prizeArr)
}
