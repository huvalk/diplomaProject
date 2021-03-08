package repository

import (
	"context"
	"diplomaProject/application/models"
)

func (r *InviteRepository) setMutual(checkMutual string, updateSilent string,
	invitation *models.Invitation) (is bool, err error) {
	err = r.conn.QueryRow(context.Background(),
		checkMutual, invitation.GuestID, invitation.OwnerID, invitation.EventID).Scan(&is)

	if is {
		_, err = r.conn.Exec(context.Background(), updateSilent, invitation.GuestID,
			invitation.OwnerID, invitation.EventID)
	}

	return is, err
}

func (r *InviteRepository) UserMutualUser(invitation *models.Invitation) (is bool, err error) {
	checkMutual := `select exists( 
			select 1 from invite as i
			where user_id = $1
			and guest_user_id = $2
			and event_id = $3
			)`

	updateSilent := `update invite 
			set silent = false
			where ((user_id = $1
			and guest_user_id = $2)
			or (guest_user_id = $1
			and user_id = $2))
			and event_id = $3
			and silent = true`

	return r.setMutual(checkMutual, updateSilent, invitation)
}

func (r *InviteRepository) UserMutualTeam(invitation *models.Invitation) (is bool, err error) {
	checkMutual := `select exists( 
			select 1 from invite as i
			where team_id = $1
			and guest_user_id = $2
			and event_id = $3
			)`

	updateSilent := `update invite 
			set silent = false
			where ((team_id = $1
			and guest_user_id = $2)
			or (guest_team_id = $1
			and user_id = $2))
			and event_id = $3
			and silent = true`

	return r.setMutual(checkMutual, updateSilent, invitation)
}

func (r *InviteRepository) TeamMutualUser(invitation *models.Invitation) (is bool, err error) {
	checkMutual := `select exists( 
			select 1 from invite as i
			where user_id = $1
			and guest_team_id = $2
			and event_id = $3
			)`

	updateSilent := `update invite 
			set silent = false
			where ((user_id = $1
			and guest_team_id = $2)
			or (guest_user_id = $1
			and team_id = $2))
			and event_id = $3
			and silent = true`

	return r.setMutual(checkMutual, updateSilent, invitation)
}

func (r *InviteRepository) TeamMutualTeam(invitation *models.Invitation) (is bool, err error) {
	checkMutual := `select exists( 
			select 1 from invite as i
			where team_id = $1
			and guest_team_id = $2
			and event_id = $3
			)`

	updateSilent := `update invite 
			set silent = false
			where ((team_id = $1
			and guest_team_id = $2)
			or (guest_team_id = $1
			and team_id = $2))
			and event_id = $3
			and silent = true`

	return r.setMutual(checkMutual, updateSilent, invitation)
}
