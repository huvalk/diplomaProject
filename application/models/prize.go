package models

type Prize struct {
	Id            int    `json:"id"`
	EventID       int    `json:"eventID"`
	Name          string `json:"name"`
	Place         int    `json:"place"`
	Amount        int    `json:"amount"`
	WinnerTeamIDs []int  `json:"winnerTeamIDs"`
}

//easyjson:json
type PrizeArr []Prize

type SelectWinner struct {
	PrizeID int `json:"prizeID"`
	EventID int `json:"eventID"`
	TeamID  int `json:"teamID"`
}
