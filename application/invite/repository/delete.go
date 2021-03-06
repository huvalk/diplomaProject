package repository

import (
	"context"
	"diplomaProject/application/models"
	"errors"
)

func (r *InviteRepository) UnInvite(inv *models.Invitation) error {
	query := `WITH owner_user_team(team_id) AS (
				select find_users_lead_team($1, $2)
			), guest_user_team(team_id) AS (
				select find_users_team($3, $2)
			) delete from invite
			using owner_user_team, guest_user_team
			where (
				(
					user_id = $1
					and invite.team_id is null
				)
				or invite.team_id = owner_user_team.team_id
			)
			and event_id = $2
			and ( 
				guest_user_id = $3
				or guest_team_id = guest_user_team.team_id
			)
			and rejected = false`

	res, err := r.conn.Exec(context.Background(), query, inv.OwnerID, inv.EventID, inv.GuestID)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("no invites to deny")
	}

	return err
}

// Удаляю инвайты
func (r *InviteRepository) Deny(inv *models.Invitation) error {
	deny := `WITH owner_user_team(team_id) AS (
				select find_users_team($1, $3)
			), guest_user_team(team_id) AS (
				select find_users_lead_team($2, $3)
			)
			delete from invite
			using guest_user_team, owner_user_team
			where
			(
				invite.team_id = owner_user_team.team_id
				or user_id = $1
			)
			and (
				(
					guest_user_id = $2
					and guest_team_id is null
				)
				or guest_team_id = guest_user_team.team_id
			)
			and event_id = $3
			and approved = false`
	//			TODO убрать rejected для разбана 'and rejected = false'
	res, err := r.conn.Exec(context.Background(), deny, inv.OwnerID, inv.GuestID, inv.EventID)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("no invite to delete")
	}

	return nil
}

func (r *InviteRepository) UpdateTeamMerged(teamFromID1 int, teamFromID2 int, teamToID int, eventID int) error {
	query := `delete from invite
				where ((
						team_id = $1
						and guest_team_id = $2
					)
					or (
						team_id = $2
						and guest_team_id = $1
					)
				) 
				and event_id = $3
				and rejected = false
				and approved = false`

	_, err := r.conn.Exec(context.Background(), query, teamFromID1, teamFromID2, eventID)

	if err != nil {
		return err
	}

	if err = r.changeTeamToTeam(teamFromID1, teamToID, eventID); err != nil {
		return err
	}

	return r.changeTeamToTeam(teamFromID2, teamToID, eventID)
}

// TODO Возможно лишнее, ничего не удаляет
func (r *InviteRepository) AcceptInvite(userID1 int, userID2 int, eventID int) error {
	// должна быть проверка на лида, но она работает в юзкейсе
	query := `WITH owner_user_team(team_id) AS (
					select find_users_team($1, $2)
				), guest_user_team(team_id) AS (
					select find_users_team($3, $2)
				)
				delete from invite
				using guest_user_team, owner_user_team
				where event_id = $2
				and (( 
					user_id = $1
					or invite.team_id = owner_user_team.team_id
				)
				and ( 
					guest_user_id = $3
					or guest_team_id = guest_user_team.team_id
				) or ( 
					user_id = $3
					or invite.team_id = guest_user_team.team_id
				)
				and ( 
					guest_user_id = $1
					or guest_team_id = owner_user_team.team_id
				))
				and rejected = false
				and approved = false`

	res, err := r.conn.Exec(context.Background(), query, userID1, eventID, userID2)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("no invite to accept")
	}

	return nil
}
