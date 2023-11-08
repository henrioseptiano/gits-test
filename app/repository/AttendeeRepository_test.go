package repository

import (
	"gits-test/models/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateAttendeeRepository(t *testing.T) {

	// Expect that the CreateAttendeeRepository method will be called with the provided attendee
	maem.Mock.On("Create", mock.Anything).Return(&gorm.DB{})

	// Create an AppRepository with the mock

	// Call the method
	attendeeID, _ := uuid.NewUUID()
	attendee := model.Attendee{
		AttendeeID:  attendeeID.String(),
		MemberID:    2,
		GatheringID: 1,
	}
	err := mainAppEntity.CreateAttendeeRepository(attendee)

	// Assert that the expected method was called
	maem.Mock.AssertExpectations(t)

	// Check if an error occurred
	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
}

func TestListAttendeeRepository(t *testing.T) {
	// Initialize the mock
	/*dest := []response.ListAttendeeResponse{
		{
			AttendeeID:  "1",
			MemberID:    "1",
			GatheringID: "1",
		},
	}*/
	maem.Mock.On("Table", "attendees as a").Return(&gorm.DB{})
	//maem.Mock.On("Joins", "JOIN gatherings as g ON a.gathering_id = g.id").Return(&gorm.DB{})
	//maem.Mock.On("Joins", "JOIN members as u ON a.member_id = u.id").Return(&gorm.DB{})
	//maem.Mock.On("Select", "a.attendee_id, u.member_id, g.gathering_id").Return(&gorm.DB{})
	//maem.Mock.On("Find", &dest).Return(&gorm.DB{})

	_, err := mainAppEntity.ListAttendeeRepository()

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}
