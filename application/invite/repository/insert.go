package repository

import (
	"context"
	"diplomaProject/application/models"
)

func (r *InviteRepository) Invite(invitation *models.Invitation) error {
	sql := `WITH owner_user_team(team_id) AS (
				select find_users_team($1)
			), guest_user_team(team_id) AS (
				select find_users_team($2)
			)
			insert into invite
			(user_id, team_id, event_id, guest_user_id, guest_team_id, silent)
			select CASE WHEN owner_user_team.team_id is null THEN $1
						ELSE null
					END, 
				   owner_user_team.team_id, $3, 
				   CASE WHEN guest_user_team.team_id is null THEN $2
						ELSE null
					END, guest_user_team.team_id, 
				   $4
			from guest_user_team, owner_user_team;`

	_, err := r.conn.Exec(context.Background(),
		sql, invitation.OwnerID, invitation.GuestID, invitation.EventID, invitation.Silent)
	return err
}
