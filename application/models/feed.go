package models

type Feed struct {
	Id    int     `json:"id"`
	Users UserArr `json:"users"`
	Event int     `json:"event"`
}

//easyjson:json
type FeedArr []Feed
