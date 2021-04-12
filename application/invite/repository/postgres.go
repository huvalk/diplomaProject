package repository

import (
	"database/sql"
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

func (r *InviteRepository) IsMutual(invitation *models.Invitation) (is bool, err error) {
	reverseInv := &models.Invitation{
		OwnerID: invitation.GuestID,
		GuestID: invitation.OwnerID,
		EventID: invitation.EventID,
	}

	is, _, err = r.IsInvited(reverseInv)
	return is, err
}

func (r *InviteRepository) UpdateUserJoinedTeam(userID1 int, teamID int, eventID int) error {
	nullTeamID := sql.NullInt64{
		Int64: int64(teamID),
		Valid: true,
	}

	err := r.setGuestUserTeam(userID1, nullTeamID, eventID)
	if err != nil {
		return err
	}

	return r.setUserTeam(userID1, nullTeamID, eventID)
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