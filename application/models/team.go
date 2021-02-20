package models

type Team struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Members UserArr `json:"members"`
}

//easyjson:json
type TeamArr []Team

type AddToTeam struct {
	TID int `json:"tid"`
	UID int `json:"uid"`
}

type AddToUser struct {
	UID1 int `json:"uid1"`
	UID2 int `json:"uid2"`
}
