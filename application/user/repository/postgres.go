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

func (ud *UserDatabase) GetUserHistory(uid int) (models.HistoryEventArr, error) {
	historyArr := models.HistoryEventArr{}
	hEvt := models.HistoryEvent{}
	sql := `select e1.id,e1.name,p1.place from prize_users pu1
join prize p1 on pu1.prize_id=p1.id
join event e1 on p1.event_id=e1.id
where pu1.user_id =$1`
	queryResult, err := ud.conn.Query(context.Background(), sql, uid)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&hEvt.Id, &hEvt.Name, &hEvt.UserPlace)
		if err != nil {
			return nil, err
		}
		historyArr = append(historyArr, hEvt)
	}
	queryResult.Close()
	return historyArr, nil
}

func (ud *UserDatabase) SetImage(uid int, link string) error {
	sql := `update users set avatar=$1 where id=$2`

	queryResult, err := ud.conn.Exec(context.Background(), sql, link, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already join event")
	}
	return nil
}

func (ud *UserDatabase) Update(usr *models.User) (*models.User, error) {
	//	update users set workplace = 'wp' , description = 'dr'  where id=4 returning id;
	sql := `update users set  `
	if usr.WorkPlace != "" {
		sql += "workplace = '" + usr.WorkPlace + "', "
	}
	if usr.Description != "" {
		sql += "description = '" + usr.Description + "', "
	}
	if usr.Bio != "" {
		sql += "bio = '" + usr.Bio + "', "
	}
	if usr.Email != "" {
		sql += "email = '" + usr.Email + "', "
	}
	if usr.FirstName != "" {
		sql += "firstname = '" + usr.FirstName + "', "
	}
	if usr.LastName != "" {
		sql += "lastname = '" + usr.LastName + "', "
	}
	if usr.Vk != "" {
		sql += "vk_url = '" + usr.Vk + "', "
	}
	if usr.Git != "" {
		sql += "gh_url = '" + usr.Git + "', "
	}
	if usr.Tg != "" {
		sql += "tg_url = '" + usr.Tg + "', "
	}
	sql = sql[:len(sql)-2] + ` where id=$1 returning id`

	id := 0
	err := ud.conn.QueryRow(context.Background(), sql, usr.Id).Scan(&id)
	if err != nil {
		return nil, err
	}
	return ud.GetByID(id)
}

func (ud *UserDatabase) SearchUserByTag(eid int, tag string) (models.UserArr, error) {
	var users models.UserArr
	sql := `select u.* from users u
			join event_users eu on u.id = eu.user_id
			where (vk_url like concat($1::text, '%')
				or gh_url like concat($1::text, '%')
				or tg_url like concat($1::text, '%'))
   			and event_id = $2
			limit 10`

	queryResult, err := ud.conn.Query(context.Background(), sql, tag, eid)
	if err != nil {
		return nil, err
	}
	defer queryResult.Close()

	for queryResult.Next() {
		var u models.User
		err := queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Bio, &u.Description, &u.WorkPlace, &u.Vk,
			&u.Tg, &u.Git, &u.Avatar)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
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
		err = queryResult.Scan(&evt.Id, &evt.Name, &evt.Description, &evt.Founder,
			&evt.DateStart, &evt.DateEnd, &evt.Place, &evt.Place,
			&evt.ParticipantsCount, &evt.Logo, &evt.Background, &evt.Site, &evt.TeamSize)
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
	err := queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Bio, &u.Description, &u.WorkPlace, &u.Vk,
		&u.Tg, &u.Git, &u.Avatar)
	if err != nil {
		return nil, err
	}
	return &u, err
}

func (ud *UserDatabase) GetByName(name string) (*models.User, error) {
	u := models.User{}
	sql := `select * from users where name = $1`
	queryResult := ud.conn.QueryRow(context.Background(), sql, name)
	err := queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Bio, &u.Description, &u.WorkPlace, &u.Vk,
		&u.Tg, &u.Git, &u.Avatar)
	if err != nil {
		return nil, err
	}
	return &u, err
}
