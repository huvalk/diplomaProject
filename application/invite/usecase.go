package invite

import "diplomaProject/application/models"

type UseCase interface {
	Invite(invitation *models.Invitation) (bool, error)
	GetInvitedUser(invitation *models.Invitation) (models.UserArr, error)
	GetInvitedTeam(invitation *models.Invitation) (models.TeamArr, error)
	GetInvitationUser(invitation *models.Invitation) (models.UserArr, error)
	GetInvitationTeam(invitation *models.Invitation) (models.TeamArr, error)
}