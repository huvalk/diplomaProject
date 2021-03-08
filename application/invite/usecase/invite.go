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

func (i *InviteUseCase) Invite(invitation *models.Invitation) (res bool, err error) {
	invitationCopy := *invitation

	ownerTeam, ownerHasTeamErr := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	guestTeam, guestHasTeamErr := i.teams.GetTeamByUser(invitation.GuestID, invitation.EventID)
	if ownerTeam != nil && ownerHasTeamErr == nil {
		if guestTeam != nil && guestHasTeamErr == nil {
			invitationCopy.OwnerID, invitationCopy.GuestID = ownerTeam.Id, guestTeam.Id
			err = i.invites.TeamInviteTeam(&invitationCopy)
		} else {
			invitationCopy.OwnerID = ownerTeam.Id
			err = i.invites.TeamInviteUser(&invitationCopy)
		}
	} else {
		if guestTeam != nil && guestHasTeamErr == nil {
			invitationCopy.GuestID = guestTeam.Id
			err = i.invites.UserInviteTeam(&invitationCopy)
		} else {
			err = i.invites.UserInviteUser(&invitationCopy)
		}
	}

	if err != nil {
		return false, err
	}

	if ownerTeam != nil && ownerHasTeamErr == nil {
		if guestTeam != nil && guestHasTeamErr == nil {
			return i.invites.TeamMutualTeam(&invitationCopy)
		} else {
			return  i.invites.TeamMutualUser(&invitationCopy)
		}
	} else {
		if guestTeam != nil && guestHasTeamErr == nil {
			return  i.invites.UserMutualTeam(&invitationCopy)
		} else {
			return  i.invites.UserMutualUser(&invitationCopy)
		}
	}
}

func (i *InviteUseCase) isMutual(invitation *models.Invitation, ownerTeam *models.Team,
	guestTeam *models.Team) (res bool, err error) {
	return true, err
}

func (i *InviteUseCase) GetInvitedUser(invitation *models.Invitation) (models.IDArr, error) {
	return i.invites.GetInvitedUser(invitation)
}

func (i *InviteUseCase) GetInvitedTeam(invitation *models.Invitation) (models.IDArr, error) {
	return i.invites.GetInvitedTeam(invitation)
}

func (i *InviteUseCase) GetInvitationUser(invitation *models.Invitation) (arr models.UserArr, err error) {
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

func (i *InviteUseCase) GetInvitationTeam(invitation *models.Invitation) (arr models.TeamArr, err error) {
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