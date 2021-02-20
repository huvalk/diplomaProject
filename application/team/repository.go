package team

import "diplomaProject/application/models"

type Repository interface {
	Get(id int) (*models.Team, error)
	Create(newTeam *models.Team) error
	AddMember(tid, uid int) error
	GetTeamMembers(tid int) (*models.UserArr, error)
}
