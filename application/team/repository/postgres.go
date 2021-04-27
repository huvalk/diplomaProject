package repository

import (
	"context"
	sql2 "database/sql"
	"diplomaProject/application/invite"
	"diplomaProject/application/invite/repository"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"log"
)

type TeamDatabase struct {
	conn *pgxpool.Pool
}

var invRepo invite.Repository

func NewTeamDatabase(db *pgxpool.Pool) team.Repository {
	invRepo = repository.NewInviteRepository(db)

	return &TeamDatabase{conn: db}
}

func (t TeamDatabase) RemoveTeam(tid int) error {
	sql := `Delete from team t1 
where t1.id=$1`
	queryResult, err := t.conn.Exec(context.Background(), sql, tid)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected == 0 {
		return errors.New("team not found")
	}

	return nil
}

func (t TeamDatabase) TeamVotes(teamID int) (*models.TeamVotesArr, error) {
	var tmVotes models.TeamVotesArr
	sql := `SELECT user_id,votes from team_users where team_id = $1`

	err := pgxscan.Select(context.Background(), t.conn, &tmVotes,
		sql, teamID)

	if err != nil {
		return &models.TeamVotesArr{}, err
	}

	return &tmVotes, nil
}

func (t TeamDatabase) SelectLead(tm *models.Team) (int, error) {
	userID := 0
	votes := 0
	leadVotes := -1
	maxVotes := -1
	maxVotesID := 0
	sql := `Select user_id,votes from team_users tu1
where votes = (select max(tu2.votes) from team_users tu2 where tu2.team_id = $1)
and team_id=$1 or (team_id=$1 AND user_id=$2)`

	queryResult, err := t.conn.Query(context.Background(), sql, tm.Id, tm.LeadID)
	if err != nil {
		return 0, errors.New("select votes fail" + err.Error())
	}
	for queryResult.Next() {
		err = queryResult.Scan(&userID, &votes)
		if err != nil {
			return 0, errors.New("can't find team votes and lead votes" + err.Error())
		}
		if votes > maxVotes {
			maxVotes = votes
			maxVotesID = userID
		}
		if userID == tm.LeadID {
			leadVotes = votes
		}
	}
	queryResult.Close()
	if maxVotes > leadVotes {
		err = t.UpdateLead(tm.Id, maxVotesID)
		if err != nil {
			return 0, err
		}
		return maxVotesID, nil
	}

	return tm.LeadID, nil
}

func (t TeamDatabase) UpdateLead(tmID, newLeadID int) error {
	sql := `update team set lead_id = $1 
where id = $2 `

	queryResult, err := t.conn.Exec(context.Background(), sql, newLeadID, tmID)
	if err != nil {
		return errors.New("update lead fail" + err.Error())
	}
	affected := queryResult.RowsAffected()
	log.Println(affected)
	return nil
}

func (t TeamDatabase) ChangeUserVotesCount(tID, uID, state int) error {
	sql := `update team_users set votes = team_users.votes ` + fmt.Sprintf("%+d", state) +
		` where team_id=$1 AND user_id=$2`
	fmt.Println(sql)
	queryResult, err := t.conn.Exec(context.Background(), sql,
		tID, uID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("can't vote")
	}
	return nil
}

func (t TeamDatabase) AddVote(vote *models.Vote) error {
	sql := `insert into votes values($1,$2,$3,$4)`
	queryResult, err := t.conn.Exec(context.Background(), sql,
		vote.EventID, vote.TeamID, vote.WhoID, vote.ForWhomID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("already voted")
	}
	return nil
}

func (t TeamDatabase) GetVote(uId, tID int) (*models.Vote, error) {
	vt := models.Vote{}
	sql := `select * from votes where who_id = $1 AND team_id = $2`

	queryResult := t.conn.QueryRow(context.Background(), sql, uId, tID)
	err := queryResult.Scan(&vt.EventID, &vt.TeamID, &vt.WhoID, &vt.ForWhomID)
	if err != nil {
		return nil, err
	}
	return &vt, err
}

func (t TeamDatabase) CancelVote(vote *models.Vote) error {
	sql := `delete from votes 
where event_id = $1 AND
team_id = $2 AND
who_id = $3 AND
for_whom_id = $4`
	queryResult, err := t.conn.Exec(context.Background(), sql,
		vote.EventID, vote.TeamID, vote.WhoID, vote.ForWhomID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	if affected != 1 {
		return errors.New("can't find vote or already voted")
	}
	return nil
}

func (t TeamDatabase) CancelForUserVotes(teamID, userID int) error {
	sql := `delete from votes 
	where team_id = $1 AND
	for_whom_id = $2`
	queryResult, err := t.conn.Exec(context.Background(), sql,
		teamID, userID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	log.Println(affected)
	return nil
}

func (t TeamDatabase) SetName(newTeam *models.Team) error {
	sql := `update team set name = $1 
where id = $2 and event = $3 `
	fmt.Println(newTeam.EventID)
	queryResult, err := t.conn.Exec(context.Background(), sql, newTeam.Name, newTeam.Id, newTeam.EventID)
	if err != nil {
		return err
	}
	affected := queryResult.RowsAffected()
	log.Println(affected)
	return err
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

func (t TeamDatabase) CheckInviteStatus(uid1, uid2, evtID int) (is bool, err error) {
	is, _, err = invRepo.IsInvited(&models.Invitation{
		OwnerID: uid1,
		GuestID: uid2,
		EventID: evtID,
	})
	return is, err
}

func (t TeamDatabase) UpdateUserJoinedTeam(uid1, tid, evtID int) error {
	return invRepo.UpdateUserJoinedTeam(uid1, tid, evtID)
}

func (t TeamDatabase) AcceptInvite(uid1, uid2, evtID int) error {
	return invRepo.AcceptInvite(uid1, uid2, evtID)
}

func (t TeamDatabase) UpdateTeamMerged(tid1, tid2, tid3, evtID int) error {
	return invRepo.UpdateTeamMerged(tid1, tid2, tid3, evtID)
}

func (t TeamDatabase) Get(id int) (*models.Team, error) {
	tm := models.Team{}
	sql := `select * from team where id = $1`
	leadID := sql2.NullInt64{}

	queryResult := t.conn.QueryRow(context.Background(), sql, id)
	err := queryResult.Scan(&tm.Id, &tm.Name, &tm.EventID, &leadID)
	if err != nil {
		return nil, err
	}
	tm.LeadID = int(leadID.Int64)
	return &tm, err
}

func (t TeamDatabase) GetTeamByUser(uid, evtID int) (*models.Team, error) {
	tm := models.Team{}
	sql := `select t1.* from team t1 
join team_users tu1 on t1.id=tu1.team_id 
where t1.event = $1 and tu1.user_id=$2`
	leadID := sql2.NullInt64{}

	queryResult := t.conn.QueryRow(context.Background(), sql, evtID, uid)
	err := queryResult.Scan(&tm.Id, &tm.Name, &tm.EventID, &leadID)
	if err != nil {
		return nil, err
	}
	tm.LeadID = int(leadID.Int64)
	return &tm, nil
}

func (t TeamDatabase) Create(newTeam *models.Team, evtID int) (*models.Team, error) {
	sql := `INSERT INTO team VALUES(default,$1,$2,$3)  RETURNING id`
	id := 0
	err := t.conn.QueryRow(context.Background(), sql, newTeam.Name, evtID, newTeam.LeadID).Scan(&id)
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
	sql := `select u1.* from team_users tu1 
join users u1 on tu1.user_id=u1.id where tu1.team_id = $1`

	queryResult, err := t.conn.Query(context.Background(), sql, tid)
	if err != nil {
		return nil, err
	}
	for queryResult.Next() {
		//TODO:scan all
		err = queryResult.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Bio, &u.Description, &u.WorkPlace, &u.Vk,
			&u.Tg, &u.Git, &u.Avatar)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	queryResult.Close()

	return us, err
}
