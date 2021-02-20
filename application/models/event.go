package models

type Event struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Founder     string `json:"founder"`
}

//easyjson:json
type EventArr []Event
