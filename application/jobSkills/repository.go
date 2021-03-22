package jobskills

import "diplomaProject/application/models"

type Repository interface {
	AddJob(uid int, newJob *models.Job) error
	RemoveJob(uid, jid int) error
	AddSkill(uid int, job *models.Job, newSkill *models.Skills) error
	AddManySkills(uid, jbID int, skillsID []int) error
	RemoveSkill(uid, jbID, skID int) error
	RemoveAllSkills(uid, jbID int) error
	GetAllJobs() (*[]models.Job, error)
	GetSkillsByJob(jobName string) (*[]models.Skills, error)
	GetJobByName(jobName string) (*models.Job, error)
	GetJobByID(jobID int) (*models.Job, error)
	CreateJob(jobName string) (*models.Job, error)
	CreateSkill(skillName string, jbID int) (*models.Skills, error)
}
