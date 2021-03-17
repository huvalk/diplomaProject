package repository

import (
	"context"
	"diplomaProject/application/models"
)

// TODO так надо делать, чтобы не ловить баги при смене команды пользователем между запросами
func (r *InviteRepository) IsInvited(invitation *models.Invitation) (is bool, err error) {
	sql := `select exists (
				select 1 from invite
				where ( 
					user_id = $1
					or team_id = (
						select distinct(team_id) 
						from team_users
						where user_id = $1
					)
				)
				and event_id = $2
				and ( 
					guest_user_id = $3
					or guest_team_id = (
						select distinct(team_id) 
						from team_users
						where user_id = $3
					)
				)
				and rejected = false
				and approved = false
			)`

	err = r.conn.QueryRow(context.Background(), sql, invitation.OwnerID, invitation.EventID, invitation.GuestID).
		Scan(&is)
	return is, err
}


func (r *InviteRepository) IsMutual(invitation *models.Invitation) (is bool, err error) {
	reverseInv := &models.Invitation{
		OwnerID: invitation.GuestID,
		GuestID: invitation.OwnerID,
		EventID: invitation.EventID,
	}
	return r.IsInvited(reverseInv)
}

// TODO поправить
func (r *InviteRepository) GetInvitedUser(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct guest_user_id
			from invite 
			where user_id = $1
			and event_id = $2
			and guest_team_id is null`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

// TODO поправить
func (r *InviteRepository) GetInvitedTeam(invitation *models.Invitation) (arr []int, err error) {
	sql := `select distinct guest_team_id
			from invite 
			where user_id = $1
			and event_id = $2
			and guest_team_id is not null`

	return r.getIdsByEventAndID(sql, invitation.OwnerID, invitation.EventID)
}

func (r *InviteRepository) GetInvitationFromUser(invitation *models.Invitation) (arr []int, err error) {
	sql := `WITH guest_user_team(team_id) AS (
				select team_id
				from team_users
				where team_users.user_id = $1
				UNION
				SELECT null
				order by team_id
				limit 1
			)
			select distinct user_id
			from invite, guest_user_team
			where (
				invite.guest_team_id = guest_user_team.team_id
				or (guest_user_team.team_id is null
				and guest_user_id = $1)
			)
			and event_id = $2
			and team_id is null
			and rejected = false
			and approved = false
			and silent = false`

	return r.getIdsByEventAndID(sql, invitation.GuestID, invitation.EventID)
}

func (r *InviteRepository) GetInvitationFromTeam(invitation *models.Invitation) (arr []int, err error) {
	sql := `WITH guest_user_team(team_id) AS (
				select team_id
				from team_users
				where team_users.user_id = $1
				UNION
				SELECT null
				order by team_id
				limit 1
			)
			select distinct team_id
			from invite, guest_user_team
			where (
				invite.guest_team_id = guest_user_team.team_id
				or (guest_user_team.team_id is null
				and guest_user_id = $1)
			)
			and event_id = $2
			and rejected = false
			and approved = false
			and silent = false`

	return r.getIdsByEventAndID(sql, invitation.GuestID, invitation.EventID)
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

