package repository

import (
	"context"
	jobskills "diplomaProject/application/jobSkills"
	"diplomaProject/application/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type JobSkillsDatabase struct {
	conn *pgxpool.Pool
}

func NewJobSkillsDatabase(db *pgxpool.Pool) jobskills.Repository {
	return &JobSkillsDatabase{conn: db}
}

func (j JobSkillsDatabase) AddJob(uid int, newJob *models.Job) error {
	panic("implement me")
}

func (j JobSkillsDatabase) RemoveJob(uid, jid int) error {
	panic("implement me")
}

func (j JobSkillsDatabase) AddSkill(uid int, newSkill *models.Skills) error {
	panic("implement me")
}

func (j JobSkillsDatabase) RemoveSkill(uid, skid int) error {
	panic("implement me")
}

func (j JobSkillsDatabase) GetAllJobs() (*[]models.Job, error) {
	var jArr []models.Job
	jb := models.Job{}
	sql := `select * from job`
	queryResult, err := j.conn.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&jb.Id, &jb.Name)
		if err != nil {
			return nil, err
		}
		jArr = append(jArr, jb)
	}
	queryResult.Close()

	return &jArr, nil
}

func (j JobSkillsDatabase) GetSkillsByJob(jobName string) (*[]models.Skills, error) {
	var skillArr []models.Skills
	sk := models.Skills{}
	sql := `select s1.* from skills s1
join job_skills_users jsu1 on s1.id=jsu1.skill_id
join job j1 on jsu1.job_id=j1.id
where j1.name = $1`
	queryResult, err := j.conn.Query(context.Background(), sql, jobName)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&sk.Id, &sk.Name, &sk.JobID)
		if err != nil {
			return nil, err
		}
		skillArr = append(skillArr, sk)
	}
	queryResult.Close()

	return &skillArr, nil
}
