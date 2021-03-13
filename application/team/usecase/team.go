package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	//"errors"
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

//chekc invite
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

//на успешный добавление , апдейт юзер
func (t *Team) Union(uid1, uid2, evtID int) (*models.Team, error) {

	//if !t.events.CheckUser(evtID, uid1) || !t.events.CheckUser(evtID, uid2) {
	//	return nil, errors.New("user does not join event")
	//}

	// есть ли инвайт на добавление
	//hasInvite, err := t.teams.CheckInviteStatus(uid1, uid2, evtID)
	//if !hasInvite || err != nil {
	//	return nil, errors.New("user has not got invite")
	//}
	t1, err1 := t.GetTeamByUser(uid1, evtID)
	t2, err2 := t.GetTeamByUser(uid2, evtID)
	if err1 != nil {
		if err2 != nil {
			//both users have no team
			newTeam := &models.Team{
				Name: fmt.Sprintf("team-%v-%v", uid1, uid2),
			}
			newTeam, err2 = t.Create(newTeam, evtID)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			}
			err := t.teams.UpdateUserJoinedTeam(uid1, uid2, newTeam.Id, evtID)
			if err != nil {
				return nil, err
			}
			err = t.teams.UpdateUserJoinedTeam(uid2, uid1, newTeam.Id, evtID)
			if err != nil {
				return nil, err
			}
			return t.AddMember(newTeam.Id, uid1, uid2)
		} else {
			// 2 user has team
			tm, err := t.AddMember(t2.Id, uid1)
			if err != nil {
				return nil, err
			}
			err = t.teams.UpdateUserJoinedTeam(uid2, uid1, t2.Id, evtID)
			if err != nil {
				return nil, err
			}
			return tm, nil
		}
	}
	if err2 != nil {
		//1 user has team
		tm, err := t.AddMember(t1.Id, uid2)
		if err != nil {
			return nil, err
		}
		err = t.teams.UpdateUserJoinedTeam(uid1, uid2, t1.Id, evtID)
		if err != nil {
			return nil, err
		}
		return tm, nil
	}

	if t1.Id == t2.Id {
		//same team
		return t.Get(t1.Id)
	}

	//merge teams
	newTeam := &models.Team{
		Name:    t1.Name + t2.Name,
		EventID: evtID,
	}
	newTeam, err := t.Create(newTeam, evtID)
	if err != nil {
		return nil, err
	}
	var newTeamIDS []int
	for i := range t1.Members {
		newTeamIDS = append(newTeamIDS, t1.Members[i].Id)
	}
	for i := range t2.Members {
		newTeamIDS = append(newTeamIDS, t2.Members[i].Id)
	}
	err = t.teams.RemoveUsers(t1.Id)
	if err != nil {
		return nil, err
	}
	err = t.teams.RemoveUsers(t2.Id)
	if err != nil {
		return nil, err
	}
	//teamjointeam
	return t.AddMember(newTeam.Id, newTeamIDS...)

}
