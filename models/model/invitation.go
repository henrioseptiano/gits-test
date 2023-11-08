package model

import "time"

type Invitation struct {
	ID           int        `json:"id" gorm:"id,primaryKey;autoIncrement,not null"`
	InvitationID string     `json:"invitation_id" gorm:"invitation_id,unique,not null"`
	MemberID     int        `json:"member_id" gorm:"member_id,not null"`
	GatheringID  int        `json:"gathering_id" gorm:"gathering_id,unique,not null"`
	CreatedAt    time.Time  `json:"created_at" gorm:"created_at,not null"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"updated_at"`
}
