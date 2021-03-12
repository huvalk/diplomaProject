package jobskills

import "diplomaProject/application/models"

type Repository interface {
	AddJob(uid int, newJob *models.Job) error
	RemoveJob(uid, jid int) error
	AddSkill(uid int, newSkill *models.Skills) error
	RemoveSkill(uid, skid int) error
	GetAllJobs() (*[]models.Job, error)
	GetSkillsByJob(jobName string) (*[]models.Skills, error)
}
