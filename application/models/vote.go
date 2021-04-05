package models

type Vote struct {
	Id        int `json:"id"`
	EventID   int `json:"eventID"`
	TeamID    int `json:"teamID"`
	WhoID     int `json:"whoID"`
	ForWhomID int `json:"forWhomID"`
	State     int `json:"state"` //+1 / -1
}

//easyjson:json
type Votes []Vote
