package models

type Invitation struct {
	UserID  int  `json:"userID,omitempty"`
	GuestID int  `json:"guestID,omitempty"`
	EventID int  `json:"eventID,omitempty"`
	Silent  bool `json:"silent,omitempty"`
}
