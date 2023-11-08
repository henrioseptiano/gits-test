package model

type GatheringType struct {
	ID          int    `json:"id" gorm:"id,primaryKey;autoIncrement;not null"`
	GatheringID int    `json:"gathering_id" gorm:"gathering_id,not null"`
	Type        string `json:"type" gorm:"type,not null"`
}
