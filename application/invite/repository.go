package invite

import "diplomaProject/application/models"

type Repository interface {
	IsInvited(invitation *models.Invitation) (bool, error)
	UnInvite(invitation *models.Invitation) error
	Deny(invitation *models.Invitation) error
	AcceptInvite(userID1 int, userID2 int, eventID int) error
	UpdateUserJoinedTeam(userID1 int, userID2 int, teamID int, eventID int) error
	UpdateUserLeftTeam(userID int, teamID int, eventID int) error
	UpdateUserChangedTeam(userID int, teamID int, eventID int) error
	UpdateTeamMerged(teamFromID1 int, teamFromID2 int, teamToID int, eventID int)
	UserInviteUser(invitation *models.Invitation) error
	TeamInviteUser(invitation *models.Invitation) error
	UserInviteTeam(invitation *models.Invitation) error
	TeamInviteTeam(invitation *models.Invitation) error
	UserMutualUser(invitation *models.Invitation) (is bool, err error)
	TeamMutualUser(invitation *models.Invitation) (is bool, err error)
	UserMutualTeam(invitation *models.Invitation) (is bool, err error)
	TeamMutualTeam(invitation *models.Invitation) (is bool, err error)
	GetInvitedUser(invitation *models.Invitation) ([]int, error)
	GetInvitedTeam(invitation *models.Invitation) ([]int, error)
	GetUserInvitationFromUser(invitation *models.Invitation) ([]int, error)
	GetUserInvitationFromTeam(invitation *models.Invitation) ([]int, error)
	GetTeamInvitationFromUser(invitation *models.Invitation) ([]int, error)
	GetTeamInvitationFromTeam(invitation *models.Invitation) ([]int, error)
}
