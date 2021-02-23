package team

import "diplomaProject/application/models"

type Repository interface {
	Get(id int) (*models.Team, error)
	Create(newTeam *models.Team, evtID int) (*models.Team, error)
	AddMember(tid int, uid ...int) (*models.Team, error)
	GetTeamMembers(tid int) (*models.UserArr, error)
}
