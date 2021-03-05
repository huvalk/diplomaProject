package invite

import "diplomaProject/application/models"

type Repository interface {
	InviteUser(invitation *models.Invitation) error
	InviteTeam(invitation *models.Invitation) error
	GetInvitedUser(invitation *models.Invitation) ([]models.FeedUser, error)
	GetInvitedTeam(invitation *models.Invitation) (models.TeamArr, error)
	GetInvitationUser(invitation *models.Invitation) ([]models.FeedUser, error)
	GetInvitationTeam(invitation *models.Invitation) (models.TeamArr, error)
}
