package usecase

import (
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"diplomaProject/application/user"
)

type InviteUseCase struct {
	invites invite.Repository
	users user.Repository
	teams team.Repository
}

func NewInviteUseCase(inv invite.Repository, u user.Repository, t team.Repository) invite.UseCase {
	return &InviteUseCase{
		invites: inv,
		users: u,
		teams: t,
	}
}

func (i InviteUseCase) InviteUser(invitation *models.Invitation) (res bool, err error) {
	ownerTeam, ownerHasTeamErr := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	guestTeam, guestHasTeamErr := i.teams.GetTeamByUser(invitation.GuestID, invitation.EventID)
	if ownerTeam != nil && ownerHasTeamErr == nil {
		if guestTeam != nil && guestHasTeamErr == nil {
			invitation.OwnerID, invitation.GuestID = ownerTeam.Id, guestTeam.Id
			err = i.invites.TeamInviteTeam(invitation)
		} else {
			invitation.OwnerID = ownerTeam.Id
			err = i.invites.TeamInviteUser(invitation)
		}
	} else {
		if guestTeam != nil && guestHasTeamErr == nil {
			invitation.GuestID = guestTeam.Id
			err = i.invites.UserInviteTeam(invitation)
		} else {
			err = i.invites.UserInviteUser(invitation)
		}
	}

	if err != nil {
		return false, err
	}

	// TODO проверять взаимность
	//if invitation.Silent == false {
	//	if is, err := i.invites.IsInviteUserMutual(invitation); err == nil && is {
	//		res = true
	//	}
	//}

	return true, nil
}

func (i InviteUseCase) InviteTeam(invitation *models.Invitation) (res bool, err error) {
	panic("Not implemented")
}

func (i InviteUseCase) GetInvitedUser(invitation *models.Invitation) (models.UserArr, error) {
	return nil, nil
}

func (i InviteUseCase) GetInvitedTeam(invitation *models.Invitation) (users models.TeamArr, err error) {

	return nil, nil
}

func (i InviteUseCase) GetInvitationUser(invitation *models.Invitation) (arr models.UserArr, err error) {
	var userIds []int
	ownerTeam, ownerHasTeamErr := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	if ownerTeam != nil && ownerHasTeamErr == nil {
		invitation.OwnerID = ownerTeam.Id
		userIds, err = i.invites.GetTeamInvitationFromUser(invitation)
	} else {
		userIds, err = i.invites.GetUserInvitationFromUser(invitation)
	}

	if err != nil {
		return nil, err
	}

	for _, id := range userIds {
		u, err := i.users.GetByID(id)

		if err != nil {
			return nil, err
		}

		arr = append(arr, *u)
	}

	return arr, nil
}

func (i InviteUseCase) GetInvitationTeam(invitation *models.Invitation) (arr models.TeamArr, err error) {
	var userIds []int
	ownerTeam, ownerHasTeamErr := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	if ownerTeam != nil && ownerHasTeamErr == nil {
		invitation.OwnerID = ownerTeam.Id
		userIds, err = i.invites.GetTeamInvitationFromUser(invitation)
	} else {
		userIds, err = i.invites.GetUserInvitationFromUser(invitation)
	}

	if err != nil {
		return nil, err
	}

	for _, id := range userIds {
		t, err := i.teams.Get(id)

		if err != nil {
			return nil, err
		}

		arr = append(arr, *t)
	}

	return arr, nil
}