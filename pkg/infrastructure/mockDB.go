package infrastructure

import "diplomaProject/application/models"

var Usr1 = models.User{
	Id:        1,
	FirstName: "FName1",
	LastName:  "LNAME1",
	Email:     "EMAIL1",
}

var Usr2 = models.User{
	Id:        2,
	FirstName: "FName2",
	LastName:  "LNAME2",
	Email:     "EMAIL2",
}

var Usr3 = models.User{
	Id:        3,
	FirstName: "FName3",
	LastName:  "LNAME3",
	Email:     "EMAIL3",
}

var Users = []models.User{Usr1, Usr2, Usr3}

var MockTeam1 = &models.Team{
	Id:      1,
	Name:    "NAME1",
	Members: nil,
	EventID: 1,
}

var MockTeam2 = &models.Team{
	Id:      2,
	Name:    "NAME2",
	Members: nil,
	EventID: 1,
}

//map[team]  = []int{membersID...}
var TeamMembers = map[int][]int{1: nil, 2: nil}

var MockTeams = []models.Team{*MockTeam1, *MockTeam2}

var MockEvent1 = &models.Event{
	Id:          1,
	Name:        "EVENT1",
	Description: "Description1",
	Founder:     "Sber1",
}

var MockEvent2 = &models.Event{
	Id:          2,
	Name:        "EVENT2",
	Description: "Description2",
	Founder:     "Sber2",
}

var MockEvents = []models.Event{*MockEvent1, *MockEvent2}

//map[event]  = []int{userID...}
var EventMembers = map[int][]int{1: nil, 2: nil}

var EventFeed1 = &models.Feed{
	Id:    1,
	Users: nil,
	Event: 1,
}

var EventFeed2 = &models.Feed{
	Id:    2,
	Users: nil,
	Event: 2,
}

var EventFeeds = models.FeedArr{*EventFeed1, *EventFeed2}
