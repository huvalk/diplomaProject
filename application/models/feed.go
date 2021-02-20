package models

type Feed struct {
	Id    int64   `json:"id"`
	Users UserArr `json:"users"`
	Event int     `json:"event"`
}

//easyjson:json
type FeedArr []Feed
