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
	is, banned, err = i.IsInvited(invitation)
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
	return i.getSilentTeamLoudLead(invitation.OwnerID, invitation.GuestID, invitation.EventID)
}

func (i *InviteUseCase) getSilentTeamLoudLead(initiator, victim, event int) (silent []int, loud []int, err error) {
	ownerTeam, err := i.teams.GetTeamByUser(initiator, event)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, nil, err
	} else {
		err = nil
	}
	guestTeam, err := i.teams.GetTeamByUser(victim, event)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, nil, err
	} else {
		err = nil
	}

	// TODO для анонимных инвайтов
	//if invitation.Silent && notify {
	//}
	if ownerTeam != nil {
		members, err := i.teams.GetTeamMembers(ownerTeam.Id)
		if err != nil {
			return nil, nil, err
		}

		for _, member := range members {
			silent = append(silent, member.Id)
		}
	} else {
		silent = append(silent, initiator)
	}

	// TODO для анонимных инвайтов
	//if !invitation.Silent || notify {
	//}
	if guestTeam != nil {
		members, err := i.teams.GetTeamMembers(guestTeam.Id)
		if err != nil {
			return nil, nil, err
		}

		for _, member := range members {
			if member.Id == guestTeam.LeadID {
				loud = append(loud, member.Id)
			}
			silent = append(silent, member.Id)
		}
	} else {
		loud = append(loud, victim)
	}

	return silent, loud, nil
}

func (i *InviteUseCase) getAllUsers(initiator, victim, event int) (silent []int, err error) {
	ownerTeam, err := i.teams.GetTeamByUser(initiator, event)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, err
	} else {
		err = nil
	}
	guestTeam, err := i.teams.GetTeamByUser(victim, event)
	if err != nil && err.Error() != "no rows in result set" {
		return nil, err
	} else {
		err = nil
	}

	// TODO для анонимных инвайтов
	//if invitation.Silent && notify {
	//}
	if ownerTeam != nil {
		members, err := i.teams.GetTeamMembers(ownerTeam.Id)
		if err != nil {
			return nil, err
		}

		for _, member := range members {
			silent = append(silent, member.Id)
		}
	} else {
		silent = append(silent, initiator)
	}

	// TODO для анонимных инвайтов
	//if !invitation.Silent || notify {
	//}
	if guestTeam != nil {
		members, err := i.teams.GetTeamMembers(guestTeam.Id)
		if err != nil {
			return nil, err
		}

		for _, member := range members {
			silent = append(silent, member.Id)
		}
	} else {
		silent = append(silent, victim)
	}

	return silent, nil
}

func (i *InviteUseCase) UnInvite(invitation *models.Invitation) (inviters []int, err error) {
	err = i.invites.UnInvite(invitation)
	if err != nil {
		return nil, err
	}

	u1, err := i.getAllUsers(invitation.OwnerID, invitation.GuestID, invitation.EventID)
	if err != nil {
		return nil, err
	}
	return u1, err
}

func (i *InviteUseCase) Deny(invitation *models.Invitation) (silent []int, loud []int, err error) {
	err = i.invites.Deny(invitation)
	if err != nil {
		return nil, nil, err
	}

	return i.getSilentTeamLoudLead(invitation.GuestID, invitation.OwnerID, invitation.EventID)
}

func (i *InviteUseCase) DenyAndBan(invitation *models.Invitation) (silent []int, loud []int, err error) {
	err = i.invites.DenyAndBan(invitation)
	if err != nil {
		return nil, nil, err
	}

	return i.getSilentTeamLoudLead(invitation.GuestID, invitation.OwnerID, invitation.EventID)
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
