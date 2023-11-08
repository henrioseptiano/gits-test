package model

import "time"

type Gathering struct {
	ID          int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	GatheringID string     `json:"gathering_id" gorm:"unique,not null"`
	Creator     string     `json:"creator" gorm:"not null"`
	ScheduledAt time.Time  `json:"scheduled_at" gorm:"not null"`
	Name        string     `json:"name" gorm:"not null"`
	Location    string     `json:"location" gorm:"not null"`
	Status      bool       `json:"status" gorm:"not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
