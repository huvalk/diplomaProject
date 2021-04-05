package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/feed"
	"diplomaProject/application/models"
	"diplomaProject/pkg/constants"
	"diplomaProject/pkg/sss"
	"errors"
	"mime/multipart"
)

type Event struct {
	events event.Repository
	feeds  feed.UseCase
}

func NewEvent(e event.Repository, f feed.UseCase) event.UseCase {
	return &Event{events: e, feeds: f}
}

func (e *Event) RemovePrize(uID, evtID int, prArr *models.PrizeArr) (*models.Event, error) {
	ev, err := e.Get(evtID)
	if err != nil {
		return nil, err
	}
	if ev.Founder != uID {
		return nil, errors.New("not founder")
	}
	if prArr == nil || len(*prArr) < 1 {
		return e.Get(evtID)
	}
	err = e.events.RemovePrize(evtID, prArr)
	if err != nil {
		return nil, err
	}
	return e.Get(evtID)
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
	var prArr models.PrizeArr
	for i := range evt.PrizeList {
		if evt.PrizeList[i].Id == 0 {
			prArr = append(prArr, evt.PrizeList[i])
		} else {
			err = e.events.UpdatePrize(&evt.PrizeList[i])
			if err != nil {
				return nil, err
			}
		}
	}
	if len(prArr) > 0 {
		err = e.AddPrize(evt.Id, prArr)
		if err != nil {
			return nil, err
		}
	}
	return e.Get(evt.Id)
}

func (e *Event) UnSelectWinner(uID, evtID, PrizeID, tId int) error {
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
	pr, err := e.events.GetPrize(PrizeID)
	if err != nil {
		return errors.New("no such prize")
	}
	if len(pr.WinnerTeamIDs) == 0 {
		return nil
	}
	for i := range pr.WinnerTeamIDs {
		if pr.WinnerTeamIDs[i] == tId {
			pr.WinnerTeamIDs = append(pr.WinnerTeamIDs[:i], pr.WinnerTeamIDs[i+1:]...)
			break
		}
	}
	return e.events.UnSelectWinner(PrizeID, tId, pr.WinnerTeamIDs)
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
	if newEvent.PrizeList != nil && len(newEvent.PrizeList) > 0 {
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
	return e.events.CreatePrize(evtID, prizeArr)
}

func (e *Event) SetLogo(uid, eid int, avatar *multipart.Form) (string, error) {
	link, err := sss.UploadPic(avatar, "")
	if err != nil {
		return "", err
	}
	err = e.events.SetLogo(uid, eid, link)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (e *Event) SetBackground(uid, eid int, avatar *multipart.Form) (string, error) {
	link, err := sss.UploadPic(avatar, "")
	if err != nil {
		return "", err
	}
	err = e.events.SetBackground(uid, eid, link)
	if err != nil {
		return "", err
	}
	return link, nil
}
