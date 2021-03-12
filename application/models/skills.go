package models

type Skills struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	JobID int    `json:"jobId"`
}

//easyjson:json
type SkillsArr []Skills
