package invite

import "diplomaProject/application/models"

type Repository interface {
	UserInviteUser(invitation *models.Invitation) error
	TeamInviteUser(invitation *models.Invitation) error
	UserInviteTeam(invitation *models.Invitation) error
	TeamInviteTeam(invitation *models.Invitation) error
	IsInviteUserMutual(invitation *models.Invitation) (bool, error)
	GetInvitedUser(invitation *models.Invitation) ([]int, error)
	GetInvitedTeam(invitation *models.Invitation) ([]int, error)
	GetUserInvitationFromUser(invitation *models.Invitation) ([]int, error)
	GetUserInvitationFromTeam(invitation *models.Invitation) ([]int, error)
	GetTeamInvitationFromUser(invitation *models.Invitation) ([]int, error)
	GetTeamInvitationFromTeam(invitation *models.Invitation) ([]int, error)
}
