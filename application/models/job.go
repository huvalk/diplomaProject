package models

type Job struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//easyjson:json
type JobArr []Job
