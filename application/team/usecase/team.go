package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"errors"
	"fmt"
)

type Team struct {
	teams  team.Repository
	events event.Repository
}

func NewTeam(t team.Repository, e event.Repository) team.UseCase {
	return &Team{teams: t, events: e}
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
	tm.Members = members
	return tm, err
}

func (t *Team) Create(newTeam *models.Team, evtID int) (*models.Team, error) {
	return t.teams.Create(newTeam, evtID)
}

func (t *Team) AddMember(tid int, uid ...int) (*models.Team, error) {
	tm, err := t.teams.AddMember(tid, uid...)
	if err != nil {
		return nil, err
	}
	usrs, err := t.teams.GetTeamMembers(tm.Id)
	if err != nil {
		return nil, err
	}
	tm.Members = usrs
	return tm, nil
}

func (t *Team) GetTeamByUser(uid, evtID int) (*models.Team, error) {
	tm, err := t.teams.GetTeamByUser(uid, evtID)
	if err != nil {
		return nil, err
	}
	members, err := t.teams.GetTeamMembers(tm.Id)
	if err != nil {
		return nil, err
	}
	tm.Members = members
	return tm, err
}

func (t *Team) Union(uid1, uid2, evtID int) (*models.Team, error) {
	if !t.events.CheckUser(evtID, uid1) || !t.events.CheckUser(evtID, uid2) {
		return nil, errors.New("user does not join event")
	}
	t1, err1 := t.GetTeamByUser(uid1, evtID)
	t2, err2 := t.GetTeamByUser(uid2, evtID)
	if err1 != nil {
		if err2 != nil {
			newTeam := &models.Team{
				Name: fmt.Sprintf("team-%v-%v", uid1, uid2),
			}
			newTeam, err2 = t.Create(newTeam, evtID)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			}
			return t.AddMember(newTeam.Id, uid1, uid2)
		} else {
			tm, err := t.AddMember(t2.Id, uid1)
			if err != nil {
				return nil, err
			}
			return tm, nil
		}
	}
	if err2 == nil {
		if t1.Id == t2.Id {
			return t.Get(t1.Id)
		}
	}
	tm, err := t.AddMember(t1.Id, uid2)
	if err != nil {
		return nil, err
	}
	return tm, nil

}
