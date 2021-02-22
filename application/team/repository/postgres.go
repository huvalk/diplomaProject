package repository

import (
	"diplomaProject/application/models"
	"diplomaProject/application/team"
	"diplomaProject/application/user/repository"
	"diplomaProject/pkg/infrastructure"
	"errors"
	"github.com/jinzhu/gorm"
)

type TeamDatabase struct {
	conn *gorm.DB
}

func NewTeamDatabase(db *gorm.DB) team.Repository {
	return &TeamDatabase{conn: db}
}

func (e TeamDatabase) Get(id int) (*models.Team, error) {
	for ind := range infrastructure.MockTeams {
		if infrastructure.MockTeams[ind].Id == id {
			return &infrastructure.MockTeams[ind], nil
		}
	}
	return &models.Team{}, errors.New("team not found")
}

func (e TeamDatabase) Create(newTeam *models.Team, evtID int) (*models.Team, error) {
	newTeam.Id = len(infrastructure.MockTeams) + 1
	newTeam.EventID = evtID
	infrastructure.MockTeams = append(infrastructure.MockTeams, *newTeam)
	infrastructure.TeamMembers[newTeam.Id] = []int{}
	return newTeam, nil
}

func (e TeamDatabase) AddMember(tid int, uid ...int) (*models.Team, error) {
	teamIDs, ok := infrastructure.TeamMembers[tid]
	if !ok {
		return nil, errors.New("team not found")
	}
	for i := range uid {
		teamIDs = append(teamIDs, uid[i])
	}
	infrastructure.TeamMembers[tid] = teamIDs
	return e.Get(tid)
}

func (e TeamDatabase) GetTeamMembers(tid int) (*models.UserArr, error) {
	users := models.UserArr{}
	userDB := repository.NewUserDatabase(e.conn)
	membersID := infrastructure.TeamMembers[tid]
	for ind := range membersID {
		user, err := userDB.GetByID(membersID[ind])
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return &users, nil
}
