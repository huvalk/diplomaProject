package repository

import (
	"context"
	"diplomaProject/application/invite"
	"diplomaProject/application/invite/repository"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TeamDatabase struct {
	conn *pgxpool.Pool
}

var invRepo invite.Repository

func NewTeamDatabase(db *pgxpool.Pool) team.Repository {
	invRepo = repository.NewInviteRepository(db)

	return &TeamDatabase{conn: db}
}

func (t TeamDatabase) RemoveAllUsers(tid int) error {
	sql := `Delete from team_users tu1 
where tu1.team_id=$1`
	queryResult, err := t.conn.Exec(context.Background(), sql, tid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected == 0 {
		return errors.New("user wasn't in team")
	}

	return nil
}

func (t TeamDatabase) RemoveMember(tid, uid int) error {
	sql := `Delete from team_users tu1 
where tu1.team_id=$1 AND tu1.user_id=$2`
	queryResult, err := t.conn.Exec(context.Background(), sql, tid, uid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected == 0 {
		return errors.New("team not found")
	}

	return nil
}

func (t TeamDatabase) CheckInviteStatus(uid1, uid2, evtID int) (bool, error) {
	return invRepo.IsInvited(&models.Invitation{
		OwnerID: uid1,
		GuestID: uid2,
		EventID: evtID,
	})
}

func (t TeamDatabase) UpdateUserJoinedTeam(uid1, uid2, tid, evtID int) error {
	return invRepo.UpdateUserJoinedTeam(uid1, uid2, tid, evtID)
}

func (t TeamDatabase) UpdateTeamMerged(tid1, tid2, tid3, evtID int) error {
	return invRepo.UpdateTeamMerged(tid1, tid2, tid3, evtID)
}

func (t TeamDatabase) Get(id int) (*models.Team, error) {
	tm := models.Team{}
	sql := `select * from team where id = $1`

	queryResult := t.conn.QueryRow(context.Background(), sql, id)
	err := queryResult.Scan(&tm.Id, &tm.Name, &tm.EventID)
	if err != nil {
		return nil, err
	}
	return &tm, err
}

func (t TeamDatabase) GetTeamByUser(uid, evtID int) (*models.Team, error) {
	tm := models.Team{}
	sql := `select t1.* from team t1 join team_users tu1 on t1.id=tu1.team_id 
where t1.event = $1 and tu1.user_id=$2`

	queryResult := t.conn.QueryRow(context.Background(), sql, evtID, uid)
	err := queryResult.Scan(&tm.Id, &tm.Name, &tm.EventID)
	if err != nil {
		return nil, err
	}
	return &tm, nil
}

func (t TeamDatabase) Create(newTeam *models.Team, evtID int) (*models.Team, error) {
	sql := `INSERT INTO team VALUES(default,$1,$2)  RETURNING id`
	id := 0
	err := t.conn.QueryRow(context.Background(), sql, newTeam.Name, evtID).Scan(&id)
	if err != nil {
		return nil, err
	}
	return t.Get(id)
}

func (t TeamDatabase) AddMember(tid int, uid ...int) (*models.Team, error) {
	sql := `INSERT INTO team_users VALUES($1,$2)`
	for i := range uid {
		queryResult, err := t.conn.Exec(context.Background(), sql, tid, uid[i])
		if err != nil {
			return nil, err
		}
		affected := queryResult.RowsAffected()
		if affected != 1 {
			return nil, errors.New("already in team")
		}
	}
	return t.Get(tid)
}

func (t TeamDatabase) GetTeamMembers(tid int) ([]models.User, error) {
	var us []models.User
	u := models.User{}
	sql := `select u1.id,u1.firstname,u1.lastname,u1.email from team t1 
join team_users tu1 on t1.id=tu1.team_id 
join users u1 on tu1.user_id=u1.id where t1.id = $1`

	queryResult, err := t.conn.Query(context.Background(), sql, tid)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		err = queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	queryResult.Close()

	return us, err
}
