package invite

import "diplomaProject/application/models"

type UseCase interface {
	Invite(invitation *models.Invitation) (bool, error)
	UnInvite(invitation *models.Invitation) error
	Deny(invitation *models.Invitation) error
	IsInvited(invitation *models.Invitation) (bool, error)
	GetInvitedUser(invitation *models.Invitation) (models.IDArr, error)
	GetInvitedTeam(invitation *models.Invitation) (models.IDArr, error)
	GetInvitationUser(invitation *models.Invitation) (models.UserArr, error)
	GetInvitationTeam(invitation *models.Invitation) (models.TeamArr, error)
}