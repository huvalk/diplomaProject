package usecase

import (
	"diplomaProject/application/jobSkills"
	"diplomaProject/application/models"
)

type JobSkills struct {
	jobSkills jobskills.Repository
}

func NewJobSkills(js jobskills.Repository) jobskills.UseCase {
	return &JobSkills{jobSkills: js}
}

func (j JobSkills) AddJob(uid int, newJob *models.Job) error {
	panic("implement me")
}

func (j JobSkills) RemoveJob(uid, jid int) error {
	panic("implement me")
}

func (j JobSkills) AddSkill(uid int, newSkill *models.Skills) error {
	panic("implement me")
}

func (j JobSkills) RemoveSkill(uid, skid int) error {
	panic("implement me")
}

func (j JobSkills) GetAllJobs() (*[]models.Job, error) {
	return j.jobSkills.GetAllJobs()
}

func (j JobSkills) GetSkillsByJob(jobName string) (*[]models.Skills, error) {
	return j.jobSkills.GetSkillsByJob(jobName)
}
