package models

type Team struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Members UserArr `json:"members"`
}

//easyjson:json
type TeamArr []Team

type Adder struct {
	TID int `json:"tid"`
	UID int `json:"uid"`
}
