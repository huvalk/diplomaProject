package invite

import "diplomaProject/application/models"

type UseCase interface {
	Invite(invitation *models.Invitation) ([]int, []int, error)
	UnInvite(invitation *models.Invitation) ([]int, error)
	Deny(invitation *models.Invitation) ([]int, []int, error)
	DenyAndBan(inv *models.Invitation) ([]int, []int, error)
	IsInvited(invitation *models.Invitation) (bool, bool, error)
	GetInvitedUser(invitation *models.Invitation, declined bool) (models.IDArr, error)
	GetInvitedTeam(invitation *models.Invitation, declined bool) (models.IDArr, error)
	GetInvitationUser(invitation *models.Invitation) (models.IDArr, error)
	GetInvitationTeam(invitation *models.Invitation) (models.IDArr, error)
}
