package team

import "diplomaProject/application/models"

type Repository interface {
	Get(id int) (*models.Team, error)
	Create(newTeam *models.Team, evtID int) (*models.Team, error)
	SetName(newTeam *models.Team) error
	RemoveTeam(tid int) error
	RemoveAllUsers(tid int) error
	RemoveMember(tid, uid int) error
	AddMember(tid int, uid ...int) (*models.Team, error)
	GetTeamMembers(tid int) ([]models.User, error)
	GetTeamByUser(uid, evtID int) (*models.Team, error)
	CheckInviteStatus(uid1, uid2, evtID int) (bool, error)
	UpdateUserJoinedTeam(uid1, tid, evtID int) error
	AcceptInvite(uid1, uid2, evtID int) error
	UpdateTeamMerged(tid1, tid2, tid3, evtID int) error
	AddVote(vote *models.Vote) error
	CancelVote(vote *models.Vote) error
	CancelForUserVotes(teamID, userID int) error
	ChangeUserVotesCount(tID, uID, state int) error
	SelectLead(tm *models.Team) (int, error)
	TeamVotes(teamID int) (*models.TeamVotesArr, error)
	GetVote(uId, tID int) (*models.Vote, error)
}
