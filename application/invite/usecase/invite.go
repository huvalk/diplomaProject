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

func (i *InviteUseCase) Invite(invitation *models.Invitation) (inviters []int, invitees []int, err error) {
	err = i.invites.Invite(invitation)
	if err != nil {
		return nil, nil, err
	}

	notify, err := i.invites.MakeMutual(invitation)
	if err != nil {
		return nil, nil, err
	}

	ownerTeam, err := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	// TODO ignore no rows error
	//if err != nil {
	//	return nil, nil, err
	//}
	guestTeam, err := i.teams.GetTeamByUser(invitation.GuestID, invitation.EventID)
	// TODO ignore no rows error
	//if err != nil {
	//	return nil, nil, err
	//}

	// TODO Убрать в один запрос поиск членов команды

	var inviterIDs []int
	if invitation.Silent && notify {
		if ownerTeam != nil {
			members, err := i.teams.GetTeamMembers(ownerTeam.Id)
			if err != nil {
				return nil, nil, err
			}

			for _, member := range members {
				//if member.Id == invitation.OwnerID {
				//	continue
				//}
				inviterIDs = append(inviterIDs, member.Id)
			}
		} else {
			inviterIDs = append(inviterIDs, invitation.OwnerID)
		}
	}

	var inviteeIDs []int
	if !invitation.Silent || notify {
		if guestTeam != nil {
			members, err := i.teams.GetTeamMembers(guestTeam.Id)
			if err != nil {
				return nil, nil, err
			}

			for _, member := range members {
				inviteeIDs = append(inviteeIDs, member.Id)
			}
		} else {
			inviteeIDs = append(inviteeIDs, invitation.GuestID)
		}
	}

	return inviterIDs, inviteeIDs, nil
}

func (i *InviteUseCase) UnInvite(invitation *models.Invitation) error {
	return i.invites.UnInvite(invitation)
}

func (i *InviteUseCase) Deny(invitation *models.Invitation) (invitersIDs []int, err error) {
	err = i.invites.Deny(invitation)
	if err != nil {
		return nil, err
	}

	ownerTeam, err := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	if err != nil {
		return nil, err
	}

	var inviterIDs []int
	if ownerTeam != nil {
		for _, member := range ownerTeam.Members {
			inviterIDs = append(inviterIDs, member.Id)
		}
	} else {
		inviterIDs = append(inviterIDs, invitation.OwnerID)
	}

	return inviterIDs, nil
}

func (i *InviteUseCase) IsInvited(invitation *models.Invitation) (bool, error) {
	return i.invites.IsInvited(invitation)
}

func (i *InviteUseCase) GetInvitedUser(invitation *models.Invitation) (models.IDArr, error) {
	return i.invites.GetInvitedUser(invitation)
}

func (i *InviteUseCase) GetInvitedTeam(invitation *models.Invitation) (models.IDArr, error) {
	return i.invites.GetInvitedTeam(invitation)
}

func (i *InviteUseCase) GetInvitationUser(invitation *models.Invitation) (arr models.UserArr, err error) {
	var userIds []int

	userIds, err = i.invites.GetInvitationFromUser(invitation)
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
	userIds, err = i.invites.GetInvitationFromTeam(invitation)
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