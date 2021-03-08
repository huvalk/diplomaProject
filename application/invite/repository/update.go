package repository

import (
	"context"
	"database/sql"
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

func (r *InviteRepository) UpdateUserJoinedTeam(userID int, teamID int, eventID int) error {
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
