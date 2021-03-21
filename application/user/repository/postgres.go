package repository

import (
	"context"
	"diplomaProject/application/models"
	"diplomaProject/application/user"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserDatabase struct {
	conn *pgxpool.Pool
}

func NewUserDatabase(db *pgxpool.Pool) user.Repository {
	return &UserDatabase{conn: db}
}

func (ud *UserDatabase) GetUserParams(uid int) (models.Job, []models.Skills, error) {
	var skillsArr []models.Skills
	skill := models.Skills{}
	job := models.Job{}
	sql := `select j1.* , s1.* from job_skills_users jsu1
join job j1 on jsu1.job_id=j1.id
join skills s1 on jsu1.skill_id=s1.id
where jsu1.user_id = $1`
	queryResult, err := ud.conn.Query(context.Background(), sql, uid)
	if err != nil {
		return models.Job{}, nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&job.Id, &job.Name, &skill.Id, &skill.Name, &skill.JobID)
		if err != nil {
			return models.Job{}, nil, err
		}
		if skill.JobID != job.Id {
			continue
		}
		skillsArr = append(skillsArr, skill)
	}
	queryResult.Close()

	return job, skillsArr, nil
}

func (ud *UserDatabase) JoinEvent(uid, evtID int) error {
	sql := `insert into event_users values($1,$2)`
	queryResult, err := ud.conn.Exec(context.Background(), sql, evtID, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already join event")
	}
	return nil
}

func (ud *UserDatabase) GetUserEvents(uid int) (*models.EventArr, error) {
	evtArr := models.EventArr{}
	evt := models.Event{}
	sql := `select e1.* from event_users eu1 
join event e1 on eu1.event_id=e1.id
where eu1.user_id = $1`
	queryResult, err := ud.conn.Query(context.Background(), sql, uid)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&evt.Id, &evt.Name, &evt.Description, &evt.Founder, &evt.DateStart, &evt.DateEnd,
			&evt.Place, &evt.Place)
		if err != nil {
			return nil, err
		}
		evtArr = append(evtArr, evt)
	}
	queryResult.Close()
	return &evtArr, nil
}

func (ud *UserDatabase) LeaveEvent(uid, evtID int) error {
	sql := `delete from event_users where event_id=$1 and user_id=$2`
	queryResult, err := ud.conn.Exec(context.Background(), sql, evtID, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already join event")
	}
	return nil
}

func (ud *UserDatabase) GetByID(uid int) (*models.User, error) {
	u := models.User{}
	sql := `select * from users where id = $1`
	queryResult := ud.conn.QueryRow(context.Background(), sql, uid)
	err := queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Bio, &u.Description, &u.WorkPlace)
	if err != nil {
		return nil, err
	}
	return &u, err
}

func (ud *UserDatabase) GetByName(name string) (*models.User, error) {
	u := models.User{}
	sql := `select * from users where name = $1`
	queryResult := ud.conn.QueryRow(context.Background(), sql, name)
	err := queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Bio, &u.Description, &u.WorkPlace)
	if err != nil {
		return nil, err
	}
	return &u, err
}
