package model

import "time"

type Attendee struct {
	ID          int        `json:"id" gorm:"id,primaryKey;autoIncrement,not null"`
	AttendeeID  string     `json:"attendee_id" gorm:"attendee_id,unique,not null"`
	MemberID    int        `json:"member_id" gorm:"member_id,not null"`
	GatheringID int        `json:"gathering_id" gorm:"gathering_id,not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"created_at,not null"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"updated_at"`
}
