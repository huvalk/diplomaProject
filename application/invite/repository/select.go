package repository

import (
	"context"
	"diplomaProject/application/models"
)

func (r *InviteRepository) IsInvited(invitation *models.Invitation) (invited bool, banned bool, err error) {
	sqlQuery := `WITH owner_user_team(team_id) AS (
				select find_users_team($1, $2)
			), guest_user_team(team_id) AS (
				select find_users_team($3, $2)
			) select i.rejected from invite i, owner_user_team, guest_user_team
			where ( 
				i.user_id = $1
				or i.team_id = owner_user_team.team_id
			)
			and i.event_id = $2
			and ( 
				i.guest_user_id = $3
				or i.guest_team_id = guest_user_team.team_id
			)
			and i.approved = false
			order by rejected desc
			limit 1`

	err = r.conn.QueryRow(context.Background(), sqlQuery, invitation.OwnerID, invitation.EventID, invitation.GuestID).
		Scan(&banned)
	if err != nil {
		if err.Error() == "no rows in result set" {
			invited = false
			err = nil
		} else {
			return false, false, err
		}
	} else {
		invited = true
	}
	return invited, banned, err
}

// TODO поправить
func (r *InviteRepository) GetInvitedUser(invitation *models.Invitation, declined bool) (arr []int, err error) {
	var sqlQuery string
	if declined {
		sqlQuery = `WITH guest_user_team(team_id) AS (
				select find_users_team($1, $2)
			)
			select distinct user_id
			from invite, guest_user_team
			where ( guest_user_id = $1
				or invite.guest_team_id = guest_user_team.team_id
			)
			and event_id = $2
			and guest_team_id is null
			and guest_user_id is not null
			and approved = false
			and rejected = true`
	} else {
		sqlQuery = `WITH owner_user_team(team_id) AS (
				select find_users_team($1, $2)
			)
			select distinct guest_user_id
			from invite, owner_user_team
			where ( user_id = $1
				or invite.team_id = owner_user_team.team_id
			)
			and event_id = $2
			and guest_team_id is null
			and guest_user_id is not null
			and approved = false
			and rejected = false`
	}

	return r.getIdsByEventAndID(sqlQuery, invitation.OwnerID, invitation.EventID)
}

// TODO поправить
func (r *InviteRepository) GetInvitedTeam(invitation *models.Invitation, declined bool) (arr []int, err error) {
	var sqlQuery string
	if declined {
		sqlQuery = `WITH guest_user_team(team_id) AS (
				select find_users_team($1, $2)
			)
			select distinct invite.team_id
			from invite, guest_user_team
			where ( guest_user_id = $1
				or invite.guest_team_id = guest_user_team.team_id
			)
			and event_id = $2
			and guest_team_id is not null
			and approved = false
			and rejected = true`
	} else {
		sqlQuery = `WITH owner_user_team(team_id) AS (
				select find_users_team($1, $2)
			)
			select distinct guest_team_id
			from invite, owner_user_team
			where ( user_id = $1
				or invite.team_id = owner_user_team.team_id
			)
			and event_id = $2
			and guest_team_id is not null
			and approved = false
			and rejected = false`
	}

	return r.getIdsByEventAndID(sqlQuery, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) GetInvitationFromUser(invitation *models.Invitation) (arr []int, err error) {
	sqlQuery := `WITH guest_user_team(team_id) AS (
				select find_users_team($1, $2)
			)
			select distinct user_id
			from invite, guest_user_team
			where (
				invite.guest_team_id = guest_user_team.team_id
				or (guest_user_team.team_id is null
				and guest_user_id = $1)
			)
			and event_id = $2
			and invite.team_id is null
			and user_id is not null
			and rejected = false
			and approved = false`

	return r.getIdsByEventAndID(sqlQuery, invitation.GuestID, invitation.EventID)
}

func (r *InviteRepository) GetInvitationFromTeam(invitation *models.Invitation) (arr []int, err error) {
	sqlQuery := `WITH guest_user_team(team_id) AS (
				select find_users_team($1, $2)
			)
			select distinct invite.team_id
			from invite, guest_user_team
			where (
				invite.guest_team_id = guest_user_team.team_id
				or (guest_user_team.team_id is null
				and guest_user_id = $1)
			)
			and event_id = $2
			and rejected = false
			and approved = false
			and invite.team_id is not null`

	return r.getIdsByEventAndID(sqlQuery, invitation.GuestID, invitation.EventID)
}

func (r *InviteRepository) getIdsByEventAndID(sqlQuery string, ID int, eventID int) (arr []int, err error) {
	rows, err := r.conn.Query(context.Background(), sqlQuery, ID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		arr = append(arr, id)
	}
	return arr, err
}
