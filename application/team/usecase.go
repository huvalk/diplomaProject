package team

import "diplomaProject/application/models"

type UseCase interface {
	Get(id int) (*models.Team, error)
	Create(newTeam *models.Team, evtID int) (*models.Team, error)
	SetName(newTeam *models.Team) (*models.Team, error)
	AddMember(tid int, uid ...int) (*models.Team, error)
	RemoveMember(tid, uid int) (*models.Team, error)
	Union(uid1, uid2, evtID int) (*models.Team, error)
	GetTeamByUser(uid, evtID int) (*models.Team, error)
	SendVote(vote *models.Vote) (*models.Team, error)
}
