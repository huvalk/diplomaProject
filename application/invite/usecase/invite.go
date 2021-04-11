package usecase

import (
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"diplomaProject/application/user"
)

type InviteUseCase struct {
	invites invite.Repository
	users   user.Repository
	teams   team.Repository
}

func NewInviteUseCase(inv invite.Repository, u user.Repository, t team.Repository) invite.UseCase {
	return &InviteUseCase{
		invites: inv,
		users:   u,
		teams:   t,
	}
}

func (i *InviteUseCase) Invite(invitation *models.Invitation) (inviters []int, invitees []int, err error) {
	mutualInvite := models.Invitation{
		OwnerID: invitation.GuestID,
		GuestID: invitation.OwnerID,
		EventID: invitation.EventID,
		Silent:  false,
	}
	is, banned, err := i.IsInvited(&mutualInvite)
	if err != nil || is || banned {
		return nil, nil, err
	}
	err = i.invites.Invite(invitation)
	if err != nil {
		return nil, nil, err
	}

	// TODO для анонимных инвайтов
	//notify, err := i.invites.MakeMutual(invitation)
	//if err != nil {
	//	return nil, nil, err
	//}

	ownerTeam, err := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, nil, err
	}
	guestTeam, err := i.teams.GetTeamByUser(invitation.GuestID, invitation.EventID)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, nil, err
	}

	var inviterIDs []int
	// TODO для анонимных инвайтов
	//if invitation.Silent && notify {
	//}
	if ownerTeam != nil {
		members, err := i.teams.GetTeamMembers(ownerTeam.Id)
		if err != nil {
			return nil, nil, err
		}

		for _, member := range members {
			//if member.ID == invitation.OwnerID {
			//	continue
			//}
			inviterIDs = append(inviterIDs, member.Id)
		}
	} else {
		inviterIDs = append(inviterIDs, invitation.OwnerID)
	}

	var inviteeIDs []int
	// TODO для анонимных инвайтов
	//if !invitation.Silent || notify {
	//}
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
	if err != nil && err.Error() != "no rows in result set" {
		return nil, err
	}

	var inviterIDs []int
	if ownerTeam != nil {
		members, err := i.teams.GetTeamMembers(ownerTeam.Id)
		if err != nil {
			return nil, err
		}

		for _, member := range members {
			inviterIDs = append(inviterIDs, member.Id)
		}
	} else {
		inviterIDs = append(inviterIDs, invitation.OwnerID)
	}

	return inviterIDs, nil
}

func (i *InviteUseCase) DenyAndBan(invitation *models.Invitation) (invitersIDs []int, err error) {
	err = i.invites.DenyAndBan(invitation)
	if err != nil {
		return nil, err
	}

	ownerTeam, err := i.teams.GetTeamByUser(invitation.OwnerID, invitation.EventID)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, err
	}

	var inviterIDs []int
	if ownerTeam != nil {
		members, err := i.teams.GetTeamMembers(ownerTeam.Id)
		if err != nil {
			return nil, err
		}

		for _, member := range members {
			inviterIDs = append(inviterIDs, member.Id)
		}
	} else {
		inviterIDs = append(inviterIDs, invitation.OwnerID)
	}

	return inviterIDs, nil
}

func (i *InviteUseCase) IsInvited(invitation *models.Invitation) (bool, bool, error) {
	return i.invites.IsInvited(invitation)
}

func (i *InviteUseCase) GetInvitedUser(invitation *models.Invitation, declined bool) (models.IDArr, error) {
	return i.invites.GetInvitedUser(invitation, declined)
}

func (i *InviteUseCase) GetInvitedTeam(invitation *models.Invitation, declined bool) (models.IDArr, error) {
	return i.invites.GetInvitedTeam(invitation, declined)
}

func (i *InviteUseCase) GetInvitationUser(invitation *models.Invitation) (arr models.IDArr, err error) {
	return i.invites.GetInvitationFromUser(invitation)
}

func (i *InviteUseCase) GetInvitationTeam(invitation *models.Invitation) (arr models.IDArr, err error) {
	return i.invites.GetInvitationFromTeam(invitation)
}
