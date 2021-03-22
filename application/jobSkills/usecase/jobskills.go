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

//func (j JobSkills) CheckJob(uid,jID int) (*models.Job, error) {
//	//jb, err := j.jobSkills.GetJobByName(jobName)
//	//if err != nil {
//	//	return j.jobSkills.CreateJob(jobName)
//	//}
//	//return jb, nil
//	return j.jobSkills.GetJobByID(jID)
//}

func (j JobSkills) RemoveJob(uid, jid int) error {
	return j.jobSkills.RemoveJob(uid, jid)
}

func (j JobSkills) AddSkill(uid int, params *models.AddSkillIDArr) error {
	jb, err := j.jobSkills.GetJobByID((*params)[0].JobID)
	if err != nil {
		return err
	}
	//old logic
	//sks, err := j.GetSkillsByJob(jobName)
	//if err != nil {
	//	return err
	//}
	//for i := range *sks {
	//	if (*sks)[i].Name == skillName {
	//		return errors.New("already has that skill")
	//	}
	//}

	err = j.jobSkills.RemoveAllSkills(uid, jb.Id)
	if err != nil {
		return err
	}
	var skillsID []int
	for i := range *params {
		skillsID = append(skillsID, (*params)[i].SkillID)
	}

	//newSkill, err := j.jobSkills.CreateSkill(skillName, jb.Id)
	//if err != nil {
	//	return err
	//}
	return j.jobSkills.AddManySkills(uid, jb.Id, skillsID)
}

func (j JobSkills) RemoveSkill(uid, jbID, skID int) error {
	return j.jobSkills.RemoveSkill(uid, jbID, skID)
}

func (j JobSkills) GetAllJobs() (*[]models.Job, error) {
	return j.jobSkills.GetAllJobs()
}

func (j JobSkills) GetSkillsByJob(jobName string) (*[]models.Skills, error) {
	return j.jobSkills.GetSkillsByJob(jobName)
}
