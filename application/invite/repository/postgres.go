package repository

import (
	"context"
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type InviteRepository struct {
	conn *pgxpool.Pool
}

func NewInviteRepository(db *pgxpool.Pool) invite.Repository {
	return &InviteRepository{conn: db}
}

func (r InviteRepository) UserInviteUser(invitation *models.Invitation) error {
	sql := `insert into invite 
			(user_id, event_id, guest_user_id, silent) 
			values ($1, $2, $3, $4)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}

func (r InviteRepository) UserInviteTeam(invitation *models.Invitation) error {
	sql := `insert into invite
			(user_id, event_id, guest_team_id, silent) 
			values ($1, $2, $3, $4)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}

func (r InviteRepository) TeamInviteUser(invitation *models.Invitation) error {
	sql := `insert into invite 
			(team_id, event_id, guest_user_id, silent) 
			values ($1, $2, $3, $4)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}

func (r InviteRepository) TeamInviteTeam(invitation *models.Invitation) error {
	sql := `insert into invite
			(team_id, event_id, guest_team_id, silent) 
			values ($1, $2, $3, false)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}

func (r InviteRepository) IsInviteUserMutual(invitation *models.Invitation) (is bool, err error) {
	sql := `select exists( select 1 from invite 
			where user_id = $1
			and guest_user_id = $2)`

	err = r.conn.QueryRow(context.Background(),
		sql, invitation.GuestID, invitation.OwnerID).Scan(&is)
	return is, err
}

func (r InviteRepository) GetInvitedUser(invitation *models.Invitation) (arr []int, err error) {
	arr = []int{}
	sql := `select (guest_user_id)
			from invite 
			where user_id = $1
			and event_id = $2
			and silent = false
			and guest_team_id is null`

	rows, err := r.conn.Query(context.Background(), sql, invitation.OwnerID, invitation.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return nil , err
		}

		arr = append(arr, id)
	}
	return arr, err
}

func (r InviteRepository) GetInvitedTeam(invitation *models.Invitation) (arr []int, err error) {
	arr = []int{}
	sql := `select (guest_team_id)
			from invite 
			where user_id = $1
			and event_id = $2
			and silent = false
			and guest_user_id is null`

	rows, err := r.conn.Query(context.Background(), sql, invitation.OwnerID, invitation.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return nil , err
		}

		arr = append(arr, id)
	}
	return arr, err
}

func (r InviteRepository) GetUserInvitationFromUser(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct user_id
			from invite 
			where guest_user_id = $1
			and event_id = $2
			and team_id is null
			and guest_team_id is null
			and silent = false`

	rows, err := r.conn.Query(context.Background(), sql, invitation.OwnerID, invitation.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return nil , err
		}

		arr = append(arr, id)
	}
	return arr, err
}

func (r InviteRepository) GetTeamInvitationFromUser(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct user_id
			from invite 
			where guest_team_id = $1
			and event_id = $2
			and team_id is null
			and silent = false`

	rows, err := r.conn.Query(context.Background(), sql, invitation.OwnerID, invitation.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return nil , err
		}

		arr = append(arr, id)
	}
	return arr, err
}

func (r InviteRepository) GetUserInvitationFromTeam(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct team_id
			from invite 
			where guest_user_id = $1
			and event_id = $2
			and guest_team_id is null
			and silent = false`

	rows, err := r.conn.Query(context.Background(), sql, invitation.OwnerID, invitation.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return nil , err
		}

		arr = append(arr, id)
	}
	return arr, err
}

func (r InviteRepository) GetTeamInvitationFromTeam(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct team_id
			from invite 
			where guest_team_id = $1
			and event_id = $2
			and silent = false`

	rows, err := r.conn.Query(context.Background(), sql, invitation.OwnerID, invitation.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return nil , err
		}

		arr = append(arr, id)
	}
	return arr, err
}