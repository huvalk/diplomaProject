package models

type Skills struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

//easyjson:json
type SkillsArr []Skills
