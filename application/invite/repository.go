package invite

import "diplomaProject/application/models"

type Repository interface {
	Invite(invitation *models.Invitation) error
	IsInvited(invitation *models.Invitation) (bool, bool, error)
	UnInvite(invitation *models.Invitation) error
	Deny(invitation *models.Invitation) error
	DenyAndBan(inv *models.Invitation) error
	AcceptInvite(userID1 int, userID2 int, eventID int) error
	UpdateUserJoinedTeam(userID1 int, teamID int, eventID int) error
	UpdateUserLeftTeam(userID int, teamID int, eventID int) error
	UpdateUserChangedTeam(userID int, teamID int, eventID int) error
	UpdateTeamMerged(teamFromID1 int, teamFromID2 int, teamToID int, eventID int) error
	MakeMutual(invitation *models.Invitation) (is bool, err error)
	GetInvitedUser(invitation *models.Invitation, declined bool) ([]int, error)
	GetInvitedTeam(invitation *models.Invitation, declined bool) ([]int, error)
	GetInvitationFromUser(invitation *models.Invitation) ([]int, error)
	GetInvitationFromTeam(invitation *models.Invitation) ([]int, error)
}
