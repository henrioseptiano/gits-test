package model

import "time"

type Member struct {
	ID          int        `json:"id" gorm:"id,primaryKey,autoIncrement"`
	MemberID    string     `json:"member_id" gorm:"member_id,unique,not null"`
	FirstName   string     `json:"first_name" gorm:"first_name,not null"`
	LastName    string     `json:"last_name" gorm:"last_name"`
	Email       string     `json:"email" gorm:"email,unique"`
	Address     string     `json:"address" gorm:"address,not null"`
	PhoneNumber string     `json:"phone_number" gorm:"phone_number,unique,not null"`
	CreatedAt   time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"updated_at"`
}
