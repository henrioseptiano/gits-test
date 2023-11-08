package repository

import (
	"gits-test/models/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateGatheringRepository(t *testing.T) {

	// Expect that the CreateAttendeeRepository method will be called with the provided attendee
	maem.Mock.On("Create", mock.Anything).Return(&gorm.DB{})

	// Create an AppRepository with the mock

	// Call the method
	gatheringID, _ := uuid.NewUUID()
	scheduledAt := "2023-11-15 15:00:00"
	scheduledAtTime, _ := time.Parse("2006-01-02 15:04:05", scheduledAt)
	gathering := model.Gathering{
		GatheringID: gatheringID.String(),
		Creator:     "joni",
		ScheduledAt: scheduledAtTime,
		Name:        "Gathering 1",
		Location:    "Jakarta",
		Status:      true,
	}
	_, err := mainAppEntity.CreateGatheringRepository(gathering)

	// Assert that the expected method was called
	maem.Mock.AssertExpectations(t)

	// Check if an error occurred
	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
}

func TestCreateGatheringTypeRepository(t *testing.T) {

	// Expect that the CreateAttendeeRepository method will be called with the provided attendee
	maem.Mock.On("Create", mock.Anything).Return(&gorm.DB{})

	// Create an AppRepository with the mock

	// Call the method
	gatheringType := model.GatheringType{
		GatheringID: 2,
		Type:        "concert",
	}
	err := mainAppEntity.CreateGatheringTypeRepository(gatheringType)

	// Assert that the expected method was called
	maem.Mock.AssertExpectations(t)

	// Check if an error occurred
	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
}

func TestGetGatheringByGatheringIDRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Where", "gathering_id = 4ec1b06b-7ddc-11ee-9739-3822e2435ffd").Return(&gorm.DB{})
	maem.Mock.On("First", mock.Anything).Return(&gorm.DB{})
	_, err := mainAppEntity.GetGatheringByGatheringIDRepository("4ec1b06b-7ddc-11ee-9739-3822e2435ffd")

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}

func TestListGatheringRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Table", "gatherings as g").Return(&gorm.DB{})
	maem.Mock.On("Joins", "JOIN gathering_types as gt ON g.id = gt.gathering_id").Return(&gorm.DB{})
	maem.Mock.On("Where", "g.status = 1").Return(&gorm.DB{})
	maem.Mock.On("Select", "g.gathering_id, name as gathering_name, creator, DATE_FORMAT(scheduled_at, '%Y-%m-%d %H:%i:%s') AS scheduled_at, location, gt.type as gathering_type").Return(&gorm.DB{})
	maem.Mock.On("Find", mock.Anything).Return(&gorm.DB{})

	_, err := mainAppEntity.ListAttendeeRepository()

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}
