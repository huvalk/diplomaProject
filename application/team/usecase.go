package team

import "diplomaProject/application/models"

type UseCase interface {
	Get(id int) (*models.Team, error)
	Create(newTeam *models.Team) error
	AddMember(tid, uid int) error
	GetTeamByUser(uid int) (*models.Team, error)
}
