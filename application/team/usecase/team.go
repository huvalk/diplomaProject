package usecase

import (
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"diplomaProject/pkg/infrastructure"
	"errors"
	"fmt"
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

func (t *Team) Create(newTeam *models.Team) (*models.Team, error) {
	return t.teams.Create(newTeam)
}

func (t *Team) AddMember(tid int, uid ...int) (*models.Team, error) {
	tm, err := t.teams.AddMember(tid, uid...)
	if err != nil {
		return nil, err
	}
	usrs, err := t.teams.GetTeamMembers(int(tm.Id))
	if err != nil {
		return nil, err
	}
	tm.Members = *usrs
	return tm, nil
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

func (t *Team) Union(uid1, uid2 int) (*models.Team, error) {
	t1, err1 := t.GetTeamByUser(uid1)
	t2, err2 := t.GetTeamByUser(uid2)
	if err1 != nil {
		if err2 != nil {
			newTeam := &models.Team{
				Name: fmt.Sprintf("team-%v-%v", uid1, uid2),
			}
			newTeam, _ = t.Create(newTeam)
			return t.AddMember(int(newTeam.Id), uid1, uid2)
		} else {
			tm, err := t.AddMember(int(t2.Id), uid1)
			if err != nil {
				return nil, err
			}
			return tm, nil
		}
	}

	tm, err := t.AddMember(int(t1.Id), uid2)
	if err != nil {
		return nil, err
	}
	return tm, nil

}
