package repository

import (
	"context"
	"diplomaProject/application/models"
)

func (r *InviteRepository) UserInviteUser(invitation *models.Invitation) error {
	sql := `insert into invite 
			(user_id, event_id, guest_user_id, silent) 
			values ($1, $2, $3, $4)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}

func (r *InviteRepository) UserInviteTeam(invitation *models.Invitation) error {
	sql := `insert into invite
			(user_id, event_id, guest_team_id, silent) 
			values ($1, $2, $3, $4)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}

func (r *InviteRepository) TeamInviteUser(invitation *models.Invitation) error {
	sql := `insert into invite 
			(team_id, event_id, guest_user_id, silent) 
			values ($1, $2, $3, $4)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}

func (r *InviteRepository) TeamInviteTeam(invitation *models.Invitation) error {
	sql := `insert into invite
			(team_id, event_id, guest_team_id, silent) 
			values ($1, $2, $3, false)`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.EventID, invitation.GuestID, invitation.Silent)
	return err
}
