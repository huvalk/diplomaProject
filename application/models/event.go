package models

import "time"

type EventDB struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Founder           int       `json:"founder"`
	DateStart         time.Time `json:"dateStart"`
	DateEnd           time.Time `json:"dateEnd"`
	State             string    `json:"state"`
	Place             string    `json:"place"`
	ParticipantsCount int       `json:"participantsCount"`
}

type Event struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Founder           int       `json:"founder"`
	DateStart         time.Time `json:"dateStart"`
	DateEnd           time.Time `json:"dateEnd"`
	State             string    `json:"state"`
	Place             string    `json:"place"`
	Feed              Feed      `json:"feed"`
	ParticipantsCount int       `json:"participantsCount"`
	PrizeList         PrizeArr  `json:"prizeList"`
}

//easyjson:json
type EventArr []Event

func (e *Event) Convert(evt EventDB) {
	e.Id = evt.Id
	e.Name = evt.Name
	e.Description = evt.Description
	e.Founder = evt.Founder
	e.DateStart = evt.DateStart
	e.DateEnd = evt.DateEnd
	e.State = evt.State
	e.Place = evt.Place
	e.ParticipantsCount = evt.ParticipantsCount
}

type Prize struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Place         string `json:"place"`
	Amount        int    `json:"Amount"`
	WinnerTeamIDs []int  `json:"winnerTeamIDs"`
}

//easyjson:json
type PrizeArr []Prize
