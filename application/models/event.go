package models

type Event struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Founder     string `json:"founder"`
	Feed        Feed   `json:"feed"`
}

//easyjson:json
type EventArr []Event
