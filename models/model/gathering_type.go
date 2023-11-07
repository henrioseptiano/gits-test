package model

type GatheringType struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	GatheringID int    `json:"gathering_id" gorm:"not null"`
	Type        string `json:"type" gorm:"not null"`
}
