package repository

import (
	"context"
	jobskills "diplomaProject/application/jobSkills"
	"diplomaProject/application/models"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type JobSkillsDatabase struct {
	conn *pgxpool.Pool
}

func NewJobSkillsDatabase(db *pgxpool.Pool) jobskills.Repository {
	return &JobSkillsDatabase{conn: db}
}

func (j JobSkillsDatabase) AddManySkills(uid int, params *models.AddSkillIDArr) error {
	sql := `INSERT INTO job_skills_users VALUES`
	for i := range *params {
		sql += fmt.Sprintf("(%v,%v,$1),", (*params)[i].JobID, (*params)[i].SkillID)
	}
	println(sql[:len(sql)-1] + `;`)
	queryResult, err := j.conn.Exec(context.Background(), sql[:len(sql)-1]+`;`, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != int64(len(*params)) {
		return errors.New("already has that job/skill")
	}
	return nil
}

func (j JobSkillsDatabase) RemoveAllSkills(uid int) error {
	sql := `DELETE from job_skills_users jsu1 
where jsu1.user_id=$1`
	_, err := j.conn.Exec(context.Background(), sql, uid)
	if err != nil {
		return err
	}
	return nil
}

func (j JobSkillsDatabase) GetJobByID(jobID int) (*models.Job, error) {
	jb := models.Job{}
	sql := `select * from job where id = $1`
	queryResult := j.conn.QueryRow(context.Background(), sql, jobID)
	err := queryResult.Scan(&jb.Id, &jb.Name)
	if err != nil {
		return nil, err
	}
	return &jb, err
}

func (j JobSkillsDatabase) GetJobByName(jobName string) (*models.Job, error) {
	jb := models.Job{}
	sql := `select * from job where lower(name) = lower($1)`
	queryResult := j.conn.QueryRow(context.Background(), sql, jobName)
	err := queryResult.Scan(&jb.Id, &jb.Name)
	if err != nil {
		return nil, err
	}
	return &jb, err
}

func (j JobSkillsDatabase) CreateSkill(skillName string, jbID int) (*models.Skills, error) {
	sql := `INSERT INTO skills VALUES(default,$1,$2) RETURNING id`
	id := 0
	err := j.conn.QueryRow(context.Background(), sql, skillName, jbID).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &models.Skills{
		Id:    id,
		Name:  skillName,
		JobID: jbID,
	}, nil
}

func (j JobSkillsDatabase) CreateJob(jobName string) (*models.Job, error) {
	sql := `INSERT INTO job VALUES(default,$1) RETURNING id`
	id := 0
	err := j.conn.QueryRow(context.Background(), sql, jobName).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &models.Job{
		Id:   id,
		Name: jobName,
	}, nil
}

func (j JobSkillsDatabase) AddJob(uid int, newJob *models.Job) error {
	//sql := `INSERT INTO job_skills_users VALUES($1,$2)`
	//queryResult, err := j.conn.Exec(context.Background(), sql, uid, newJob.Name)
	//if err != nil {
	//	return err
	//}
	//affected := queryResult.RowsAffected()
	//if affected != 1 {
	//	return errors.New("already has that job")
	//}

	return nil
}

func (j JobSkillsDatabase) RemoveJob(uid, jid int) error {
	panic("implement me")
}

func (j JobSkillsDatabase) AddSkill(uid int, job *models.Job, newSkill *models.Skills) error {
	sql := `INSERT INTO job_skills_users VALUES($1,$2,$3)`
	queryResult, err := j.conn.Exec(context.Background(), sql, job.Id, newSkill.Id, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already has that job/skill")
	}
	return nil
}

func (j JobSkillsDatabase) RemoveSkill(uid, jbID, skID int) error {
	sql := `DELETE from job_skills_users jsu1 
where jsu1.user_id=$1
AND jsu1.job_id=$2
AND jsu1.skill_id=$3`
	_, err := j.conn.Exec(context.Background(), sql, uid, jbID, skID)
	if err != nil {
		return err
	}
	//affected := queryResult.RowsAffected()
	//if affected != 1 {
	//	return errors.New("already has that job/skill")
	//}
	return nil
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
join job j1 on s1.job_id=j1.id
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
