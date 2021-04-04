package models

type Skills struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	JobID int    `json:"jobId"`
}

//easyjson:json
type SkillsArr []Skills

type AddSkill struct {
	JobName   string `json:"jobName"`
	SkillName string `json:"skillName"`
}

type AddSkillID struct {
	JobID     int    `json:"jobID"`
	SkillName string `json:"skillName"`
	SkillID   int    `json:"skillID"`
}

//easyjson:json
type AddSkillIDArr []AddSkillID
