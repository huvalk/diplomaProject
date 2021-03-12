package models

type Skills struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	JobID int    `json:"job_id"`
}

//easyjson:json
type SkillsArr []Skills
