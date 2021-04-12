package usecase

import (
	"diplomaProject/application/event"
	"diplomaProject/application/invite"
	"diplomaProject/application/models"
	"diplomaProject/application/notification"
	"diplomaProject/application/team"
	"errors"

	//"errors"
	"fmt"
)

type Team struct {
	teams   team.Repository
	events  event.Repository
	notif   notification.UseCase
	invites invite.Repository
}

func NewTeam(t team.Repository, e event.Repository,
	n notification.UseCase, i invite.Repository) team.UseCase {
	return &Team{teams: t, events: e, notif: n, invites: i}
}

func (t *Team) GetVote(uId, tID int) (*models.Vote, error) {
	return t.teams.GetVote(uId, tID)
}

func (t *Team) TeamVotes(teamID int) (*models.TeamVotesArr, error) {
	return t.teams.TeamVotes(teamID)
}

func (t *Team) SendVote(vote *models.Vote) (*models.Team, error) {
	var err error
	tm, err := t.Get(vote.TeamID)
	if err != nil {
		return nil, err
	}
	vt, err := t.teams.GetVote(vote.WhoID, vote.TeamID)
	if err == nil && vote.State == -1 {
		err = t.teams.CancelVote(vote)
		if err != nil {
			return nil, err
		}
		err = t.teams.ChangeUserVotesCount(vote.TeamID, vote.ForWhomID, vote.State)
	} else if err == nil && vote.State == 1 {
		err = t.teams.CancelVote(vt)
		if err != nil {
			return nil, err
		}
		err = t.teams.ChangeUserVotesCount(vt.TeamID, vt.ForWhomID, -1)
	}
	if vote.State == 1 {
		err = t.teams.AddVote(vote)
		if err != nil {
			return nil, err
		}
		err = t.teams.ChangeUserVotesCount(vote.TeamID, vote.ForWhomID, vote.State)
	}
	if err != nil {
		return nil, err
	}
	var teamIDs []int
	for i := range tm.Members {
		teamIDs = append(teamIDs, tm.Members[i].Id)
	}
	err = t.notif.SendVoteNotification(teamIDs, vote.EventID)
	if err != nil {
		return nil, err
	}
	leadID, err := t.teams.SelectLead(tm)
	if err != nil {
		return nil, err
	}
	if leadID != tm.LeadID {
		err = t.notif.SendTeamLeadNotification(teamIDs, vote.EventID)
		if err != nil {
			return nil, err
		}
	}
	tm.LeadID = leadID
	return tm, nil
}

func (t *Team) SetName(newTeam *models.Team) (*models.Team, error) {
	err := t.teams.SetName(newTeam)
	if err != nil {
		return nil, errors.New("can't update team name: " + err.Error())
	}

	return t.Get(newTeam.Id)
}

func (t *Team) Get(id int) (*models.Team, error) {
	tm, err := t.teams.Get(id)
	if err != nil {
		return nil, errors.New("can't get team : " + err.Error())
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

func (t *Team) RemoveMember(tid, uid int) (*models.Team, error) {
	tm, err := t.teams.Get(tid)
	if err != nil {
		return nil, err
	}
	err = t.teams.RemoveMember(tid, uid)
	if err != nil {
		return nil, err
	}
	err = t.invites.UpdateUserLeftTeam(uid, tid, tm.EventID)
	if err != nil {
		return nil, err
	}
	return t.Get(tid)
}

func (t *Team) KickMember(tid, leadID, userID int) (*models.Team, error) {
	tm, err := t.teams.Get(tid)
	if err != nil {
		return nil, err
	}
	if leadID != tm.LeadID {
		return nil, errors.New("only lead can kick")
	}
	err = t.teams.RemoveMember(tid, userID)
	if err != nil {
		return nil, err
	}
	err = t.invites.UpdateUserLeftTeam(userID, tid, tm.EventID)
	if err != nil {
		return nil, err
	}
	return t.Get(tid)
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
	hasInvite, err := t.teams.CheckInviteStatus(uid1, uid2, evtID)
	if !hasInvite || err != nil {
		return nil, errors.New("user has not got invite")
	}
	t1, err1 := t.GetTeamByUser(uid1, evtID)
	t2, err2 := t.GetTeamByUser(uid2, evtID)
	if err1 != nil && err1.Error() == "no rows in result set" {
		if err2 != nil && err2.Error() == "no rows in result set" {
			//both users have no team
			//Что-то напутално с условиями
			newTeam := &models.Team{
				Name:   fmt.Sprintf("team-%v-%v", uid1, uid2),
				LeadID: uid1,
			}
			newTeam, err2 = t.Create(newTeam, evtID)
			if err2 != nil {
				fmt.Println(err2)
				return nil, err2
			}
			// Подтверждение инвайта
			err := t.teams.AcceptInvite(uid1, uid2, evtID)
			if err != nil {
				return nil, err
			}
			// Обновление инвайтов обоих пользователей
			err = t.teams.UpdateUserJoinedTeam(uid1, newTeam.Id, evtID)
			if err != nil {
				return nil, err
			}
			err = t.teams.UpdateUserJoinedTeam(uid2, newTeam.Id, evtID)
			if err != nil {
				return nil, err
			}
			err = t.notif.SendYouJoinTeamNotification([]int{uid1, uid2}, evtID)
			if err != nil {
				return nil, err
			}
			return t.AddMember(newTeam.Id, uid1, uid2)
		} else if err2 == nil {
			// 2 user has team
			tm, err := t.AddMember(t2.Id, uid1)
			if err != nil {
				return nil, err
			}
			// Подтверждение инвайта
			err = t.teams.AcceptInvite(uid1, uid2, evtID)
			if err != nil {
				return nil, err
			}
			// Обновление инвайтов 1 пользовотеля
			err = t.teams.UpdateUserJoinedTeam(uid1, t2.Id, evtID)
			if err != nil {
				return nil, err
			}
			err = t.notif.SendYouJoinTeamNotification([]int{uid1}, evtID)
			if err != nil {
				return nil, err
			}
			var teamIDs []int
			for i := range t2.Members {
				teamIDs = append(teamIDs, t2.Members[i].Id)
			}
			err = t.notif.SendNewMemberNotification(teamIDs, evtID)
			if err != nil {
				return nil, err
			}
			return tm, nil
		}
	} else if err1 != nil {
		return nil, err1
	}
	if err2 != nil && err2.Error() == "no rows in result set" {
		//1 user has team
		tm, err := t.AddMember(t1.Id, uid2)
		if err != nil {
			return nil, err
		}
		// Подтверждение инвайта
		err = t.teams.AcceptInvite(uid1, uid2, evtID)
		if err != nil {
			return nil, err
		}
		// Обновление инвайтов пользователей
		err = t.teams.UpdateUserJoinedTeam(uid2, t1.Id, evtID)
		if err != nil {
			return nil, err
		}
		err = t.notif.SendYouJoinTeamNotification([]int{uid2}, evtID)
		if err != nil {
			return nil, err
		}
		var teamIDs []int
		for i := range t1.Members {
			teamIDs = append(teamIDs, t1.Members[i].Id)
		}
		err = t.notif.SendNewMemberNotification(teamIDs, evtID)
		if err != nil {
			return nil, err
		}
		return tm, nil
	} else if err2 != nil {
		return nil, err2
	}

	if t1 == nil || t2 == nil {
		return nil, errors.New("no sql error but no team neither")
	}
	if t1.Id == t2.Id {
		//same team
		return t.Get(t1.Id)
	}

	//merge teams
	//TODO: move votes from one team to another
	newTeam := &models.Team{
		Name:    t1.Name + "_" + t2.Name,
		EventID: evtID,
		LeadID:  t1.LeadID,
	}
	newTeam, err = t.Create(newTeam, evtID)
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
	err = t.teams.RemoveAllUsers(t1.Id)
	if err != nil {
		return nil, err
	}
	err = t.teams.RemoveAllUsers(t2.Id)
	if err != nil {
		return nil, err
	}
	//teamjointeam
	// Подтверждение инвайта и обновление
	err = t.teams.UpdateTeamMerged(t1.Id, t2.Id, newTeam.Id, evtID)
	if err != nil {
		return nil, err
	}
	err = t.notif.SendNewMemberNotification(newTeamIDS, evtID)
	if err != nil {
		return nil, err
	}
	return t.AddMember(newTeam.Id, newTeamIDS...)
}
