package models

// Инфайт отправляемый от пользователя или от команды
type Invitation struct {
	OwnerID int  `json:"ownerID,omitempty"`
	GuestID int  `json:"guestID,omitempty"`
	EventID int  `json:"eventID,omitempty"`
	Silent  bool `json:"silent,omitempty"`
}
