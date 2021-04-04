package repository

import (
	"context"
	"database/sql"
	"diplomaProject/application/models"
)

func (r *InviteRepository) setUserTeam(userID int, teamID sql.NullInt64, eventID int) error {
	query := `update invite 
			set team_id = $1
			where user_id = $2
			and event_id = $3
			and approved = false`

	_, err := r.conn.Exec(context.Background(), query, teamID, userID, eventID)

	return err
}

func (r *InviteRepository) setGuestUserTeam(userID int, teamID sql.NullInt64, eventID int) error {
	query := `update invite 
			set guest_team_id = $1
			where guest_user_id = $2
			and event_id = $3
			and approved = false`

	_, err := r.conn.Exec(context.Background(), query, teamID, userID, eventID)

	return err
}

func (r *InviteRepository) MakeMutual(invitation *models.Invitation) (is bool, err error) {
	isMutual, err := r.IsMutual(invitation)
	if err != nil {
		return false, err
	}

	if !isMutual {
		return false, nil
	}

	updateSilent := `WITH owner_user_team(team_id) AS (
						select find_users_team($1, $3)
					), guest_user_team(team_id) AS (
						select find_users_team($2, $3)
					)
					update invite
					set silent = false
					from guest_user_team, owner_user_team
					where
						(
							(
								invite.team_id = owner_user_team.team_id
								or user_id = $1
							)
							and (
								guest_user_id = $2
								or guest_team_id = guest_user_team.team_id
							)
						)
						or (
							(
								invite.team_id = guest_user_team.team_id
								or user_id = $2
							)
							and (
								guest_user_id = $1
								or guest_team_id = owner_user_team.team_id
							)
						)
					and event_id = $3
					and silent = true
					and rejected = false
					and approved = false`

	_, err = r.conn.Exec(context.Background(), updateSilent, invitation.OwnerID, invitation.GuestID, invitation.EventID)
	return err == nil, err
}

func (r *InviteRepository) UpdateUserJoinedTeam(userID1 int, userID2 int, teamID int, eventID int) error {
	nullTeamID := sql.NullInt64{
		Int64: int64(teamID),
		Valid: true,
	}

	err := r.setGuestUserTeam(userID1, nullTeamID, eventID)
	if err != nil {
		return err
	}

	err = r.setUserTeam(userID2, nullTeamID, eventID)
	if err != nil {
		return err
	}

	return r.AcceptInvite(userID1, userID2, eventID)
}

func (r *InviteRepository) UpdateUserLeftTeam(userID int, teamID int, eventID int) error {
	nullTeamID := sql.NullInt64{
		Int64: int64(teamID),
		Valid: false,
	}

	err := r.setUserTeam(userID, nullTeamID, eventID)
	if err != nil {
		return err
	}

	return r.setGuestUserTeam(userID, nullTeamID, eventID)
}

func (r *InviteRepository) UpdateUserChangedTeam(userID int, teamID int, eventID int) error {
	nullTeamID := sql.NullInt64{
		Int64: int64(teamID),
		Valid: true,
	}

	err := r.setUserTeam(userID, nullTeamID, eventID)
	if err != nil {
		return err
	}

	return r.setGuestUserTeam(userID, nullTeamID, eventID)
}

func (r *InviteRepository) changeTeamToTeam(teamFromID int, teamToID int, eventID int) error {
	query := `update invite
				set team_id = $1
				where team_id = $2
				and event_id = $3
				and approved = false`

	if _, err := r.conn.Exec(context.Background(), query, teamToID, teamFromID, eventID); err != nil {
		return err
	}

	query = `update invite
				set guest_team_id = $1
				where guest_team_id = $2
				and event_id = $3
				and approved = false`

	_, err := r.conn.Exec(context.Background(), query, teamToID, teamFromID, eventID)

	return err
}

// Устанавливая отказ
func (r *InviteRepository) DenyAndBan(inv *models.Invitation) error {
	deny := `WITH owner_user_team(team_id) AS (
				select find_users_team($1, $3)
			), guest_user_team(team_id) AS (
				select find_users_team($2, $3)
			)
			update invite
			set rejected = true
			from guest_user_team, owner_user_team
			where
			(
				(
					invite.team_id = owner_user_team.team_id
					or user_id = $1
				)
				and (
					guest_user_id = $2
					or guest_team_id = guest_user_team.team_id
				)
			)
			or (
				(
					invite.team_id = guest_user_team.team_id
					or user_id = $2
				)
				and (
					guest_user_id = $1
					or guest_team_id = owner_user_team.team_id
				)
			)
			and event_id = $3
			and rejected = false
			and approved = false`
	_, err := r.conn.Exec(context.Background(), deny, inv.OwnerID, inv.GuestID, inv.EventID)

	return err
}
