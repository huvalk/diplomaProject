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
			and event_id = $3`

	_, err := r.conn.Exec(context.Background(), query, teamID, userID, eventID)

	return err
}

func (r *InviteRepository) setGuestUserTeam(userID int, teamID sql.NullInt64, eventID int) error {
	query := `update invite 
			set guest_team_id = $1
			where guest_user_id = $2
			and event_id = $3`

	_, err := r.conn.Exec(context.Background(), query, teamID, userID, eventID)

	return err
}

func (r *InviteRepository) AcceptInvite(userID1 int, userID2 int, eventID int) error {
	query := 	`update invite
				set approved = true
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
				)`

	_, err := r.conn.Exec(context.Background(), query, userID1, eventID, userID2)

	return err
}

func (r *InviteRepository) UpdateUserJoinedTeam(userID1 int, userID2 int, teamID int, eventID int) error {
	nullTeamID := sql.NullInt64{
		Int64: int64(teamID),
		Valid: true,
	}

	err := r.setUserTeam(userID1, nullTeamID, eventID)
	if err != nil {
		return err
	}

	err = r.setGuestUserTeam(userID1, nullTeamID, eventID)
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
	query := 	`update invite
				set team_id = $1
				where team_id = $2
				and event_id = $3`

	if _, err := r.conn.Exec(context.Background(), query, teamToID, teamFromID, eventID); err != nil {
		return err
	}

	query = 	`update invite
				set guest_team_id = $1
				where guest_team_id = $2
				and event_id = $3`

	_, err := r.conn.Exec(context.Background(), query, teamToID, teamFromID, eventID)

	return err
}

func (r *InviteRepository) UpdateTeamMerged(teamFromID1 int, teamFromID2 int, teamToID int, eventID int) error {
	query := 	`update invite
				set team_id = $1,
				guest_team_id = $1,
				approved = true
				where ((
						team_id = $2
						and guest_team_id = $3
					)
					or (
						team_id = $3
						and guest_team_id = $2
					)
				) 
				and event_id = $4`

	_, err := r.conn.Exec(context.Background(), query, teamToID, teamFromID1, teamFromID2, eventID)

	if err != nil {
		return err
	}

	if err = r.changeTeamToTeam(teamFromID1, teamToID, eventID); err != nil {
		return err
	}

	return r.changeTeamToTeam(teamFromID1, teamToID, eventID)
}

func (r *InviteRepository) Deny(inv *models.Invitation) error {
	query := `update invite
			set rejected = true
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
			)`

	_, err := r.conn.Exec(context.Background(), query, inv.OwnerID, inv.EventID, inv.GuestID)

	return err
}
