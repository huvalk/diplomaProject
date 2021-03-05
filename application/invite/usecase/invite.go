package usecase

import (
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
)

type InviteUseCase struct {
	repo invite.Repository
}

func NewInviteUseCase(n invite.Repository) invite.UseCase {
	return &InviteUseCase{
		repo: n,
	}
}

func (i InviteUseCase) InviteUser(invitation *models.Invitation) error {
	return i.InviteUser(invitation)
}

func (i InviteUseCase) InviteTeam(invitation *models.Invitation) error {
	return i.InviteTeam(invitation)
}

func (i InviteUseCase) GetInvitedUser(invitation *models.Invitation) ([]models.FeedUser, error) {
	return i.GetInvitedUser(invitation)
}

func (i InviteUseCase) GetInvitedTeam(invitation *models.Invitation) (models.TeamArr, error) {
	return i.GetInvitedTeam(invitation)
}

func (i InviteUseCase) GetInvitationUser(invitation *models.Invitation) ([]models.FeedUser, error) {
	return i.GetInvitationUser(invitation)
}

func (i InviteUseCase) GetInvitationTeam(invitation *models.Invitation) (models.TeamArr, error) {
	return i.GetInvitationTeam(invitation)
}