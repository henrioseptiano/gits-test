package repository

import (
	"fmt"
	"gits-test/models/model"
	"gits-test/models/response"
)

func (ar *AppRepository) CreateAttendeeRepository(attendee model.Attendee) error {
	err := ar.DB.Create(&attendee).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *AppRepository) ListAttendeeRepository() ([]response.ListAttendeeResponse, error) {
	listAttendee := make([]response.ListAttendeeResponse, 0)
	query := ar.DB.Table("attendees as a")
	fmt.Printf("%#v\n", query)
	query.Joins("JOIN gatherings as g ON a.gathering_id = g.id")
	query.Joins("JOIN members as u ON a.member_id = u.id")
	query.Select("a.attendee_id, u.member_id, g.gathering_id")
	query.Find(&listAttendee)
	return listAttendee, nil
}
