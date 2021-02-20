package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"diplomaProject/pkg/infrastructure"
	"errors"
)

type Team struct {
	teams team.Repository
	//users user.Repository
}

func NewTeam(t team.Repository) team.UseCase {
	return &Team{teams: t}
}

func (t *Team) Get(id int) (*models.Team, error) {
	tm, err := t.teams.Get(id)
	if err != nil {
		return nil, err
	}
	members, err := t.teams.GetTeamMembers(id)
	if err != nil {
		return nil, err
	}
	tm.Members = *members
	return tm, err
}

func (t *Team) Create(newTeam *models.Team) error {
	return t.teams.Create(newTeam)
}

func (t *Team) AddMember(tid, uid int) error {
	return t.teams.AddMember(tid, uid)
}

func (t *Team) GetTeamByUser(uid int) (*models.Team, error) {
	for ind := range infrastructure.TeamMembers {
		for i := range infrastructure.TeamMembers[ind] {
			if infrastructure.TeamMembers[ind][i] == uid {
				return t.Get(ind)
			}
		}
	}
	return &models.Team{}, errors.New("no team for user")
}
