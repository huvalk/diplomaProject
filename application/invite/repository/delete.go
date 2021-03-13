package repository

import (
	"context"
	"diplomaProject/application/models"
)

func (r *InviteRepository) UnInvite(inv *models.Invitation) error {
	query := `delete from invite
			where user_id = $1
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
