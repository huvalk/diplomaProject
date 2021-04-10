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
	Logo              string    `json:"logo"`
	Background        string    `json:"background"`
	Site              string    `json:"site"`
	TeamSize          int       `json:"teamSize"`
}

//easyjson:json
type EventDBArr []EventDB

type Event struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Founder           int       `json:"founder"`
	DateStart         time.Time `json:"dateStart"`
	DateEnd           time.Time `json:"dateEnd"`
	State             string    `json:"state"`
	Place             string    `json:"place"`
	ParticipantsCount int       `json:"participantsCount"`
	Logo              string    `json:"logo"`
	Background        string    `json:"background"`
	Site              string    `json:"site"`
	TeamSize          int       `json:"teamSize"`
	Feed              Feed      `json:"feed"`
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
	e.Logo = evt.Logo
	e.Background = evt.Background
	e.Site = evt.Site
	e.TeamSize = evt.TeamSize
}
