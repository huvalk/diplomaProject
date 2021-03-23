package jobskills

import "diplomaProject/application/models"

type UseCase interface {
	//CheckJob(uid int, jobName string) (*models.Job, error)
	RemoveJob(uid, jid int) error
	AddSkill(uid int, params *models.AddSkillIDArr) error
	RemoveSkill(uid, jbID, skID int) error
	GetAllJobs() (*[]models.Job, error)
	GetSkillsByJob(jobName string) (*[]models.Skills, error)
}
