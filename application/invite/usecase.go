package invite

import "diplomaProject/application/models"

type UseCase interface {
	InviteUser(invitation *models.Invitation) (bool, error)
	InviteTeam(invitation *models.Invitation) (bool, error)
	GetInvitedUser(invitation *models.Invitation) (models.UserArr, error)
	GetInvitedTeam(invitation *models.Invitation) (models.TeamArr, error)
	GetInvitationUser(invitation *models.Invitation) (models.UserArr, error)
	GetInvitationTeam(invitation *models.Invitation) (models.TeamArr, error)
}