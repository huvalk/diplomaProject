package repository

import (
	"context"
	"diplomaProject/application/models"
)

func (r *InviteRepository) IsInviteUserMutual(invitation *models.Invitation) (is bool, err error) {
	sql := `select exists( 
			select 1 from invite as i
			left join team_users as t
			on t.team_id = i.team_id
			)`

	err = r.conn.QueryRow(context.Background(),
		sql, invitation.GuestID, invitation.OwnerID, invitation.EventID).Scan(&is)
	return is, err
}

func (r *InviteRepository) GetInvitedUser(invitation *models.Invitation) (arr []int, err error) {
	arr = []int{}
	sql := `select distinct guest_user_id
			from invite 
			where user_id = $1
			and event_id = $2
			and guest_team_id is null`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) GetInvitedTeam(invitation *models.Invitation) (arr []int, err error) {
	arr = []int{}
	sql := `select distinct guest_team_id
			from invite 
			where user_id = $1
			and event_id = $2
			and guest_team_id is not null`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) GetUserInvitationFromUser(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct user_id
			from invite 
			where guest_user_id = $1
			and event_id = $2
			and team_id is null
			and guest_team_id is null
			and silent = false`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) GetTeamInvitationFromUser(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct user_id
			from invite 
			where guest_team_id = $1
			and event_id = $2
			and team_id is null
			and silent = false`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) GetUserInvitationFromTeam(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct team_id
			from invite 
			where guest_user_id = $1
			and event_id = $2
			and guest_team_id is null
			and silent = false`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) GetTeamInvitationFromTeam(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct team_id
			from invite 
			where guest_team_id = $1
			and event_id = $2
			and silent = false`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) getIdsByEventAndID(sql string, ID int, eventID int) (arr []int, err error) {
	rows, err := r.conn.Query(context.Background(), sql, ID, eventID)
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

