package team

import "diplomaProject/application/models"

type UseCase interface {
	Get(id int) (*models.Team, error)
	Create(newTeam *models.Team) (*models.Team, error)
	AddMember(tid int, uid ...int) (*models.Team, error)
	Union(uid1, uid2 int) (*models.Team, error)
	GetTeamByUser(uid int) (*models.Team, error)
}
