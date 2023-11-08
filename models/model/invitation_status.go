package model

// status 1 = created, 2 = accepted, 0 = rejected
type InvitationStatus struct {
	ID           int `json:"id" gorm:"id,primaryKey;autoIncrement,not null"`
	InvitationID int `json:"invitation_id" gorm:"invitation_id,unique,not null"`
	Status       int `json:"status" gorm:"status,not null"`
}
