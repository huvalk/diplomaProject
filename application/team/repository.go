package team

import "diplomaProject/application/models"

type Repository interface {
	Get(id int) (*models.Team, error)
	Create(newTeam *models.Team, evtID int) (*models.Team, error)
	AddMember(tid int, uid ...int) (*models.Team, error)
	GetTeamMembers(tid int) ([]models.User, error)
	GetTeamByUser(uid, evtID int) (*models.Team, error)
	CheckInviteStatus(uid1, uid2, evtID int) (bool, error)
}
