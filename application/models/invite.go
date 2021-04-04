package models

//easyjson:json
type IDArr []int

// Инфайт отправляемый от пользователя или от команды
type Invitation struct {
	OwnerID int  `json:"ownerID,omitempty"`
	GuestID int  `json:"guestID,omitempty"`
	EventID int  `json:"eventID,omitempty"`
	Silent  bool `json:"silent,omitempty"`
}

type IsInvited struct {
	IsInvited bool `json:"isInvited"`
}
