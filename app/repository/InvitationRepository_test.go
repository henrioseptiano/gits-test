package repository

import (
	"gits-test/models/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateInvitationRepository(t *testing.T) {

	maem.Mock.On("Create", mock.Anything).Return(&gorm.DB{})

	invitationID, _ := uuid.NewUUID()
	invitation := model.Invitation{
		InvitationID: invitationID.String(),
		MemberID:     2,
		GatheringID:  1,
	}
	_, err := mainAppEntity.CreateInvitationRepository(invitation)

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
}

func TestCreateInvitationStatusRepository(t *testing.T) {

	maem.Mock.On("Create", mock.Anything).Return(&gorm.DB{})

	invitationStatus := model.InvitationStatus{
		InvitationID: 1,
		Status:       1,
	}
	err := mainAppEntity.CreateInvitationStatusRepository(invitationStatus)

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
}

func TestAcceptNotAcceptRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Save", mock.Anything).Return(&gorm.DB{})
	invitationStatus := model.InvitationStatus{
		ID:           1,
		InvitationID: 1,
		Status:       2,
	}
	err := mainAppEntity.AcceptNotAcceptRepository(invitationStatus)

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}

func TestGetInvitationByInvitationIDRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Where", "invitation_id = 4ec1b06b-7ddc-11ee-9739-3822e2435ffd").Return(&gorm.DB{})
	maem.Mock.On("First", mock.Anything).Return(&gorm.DB{})
	_, err := mainAppEntity.GetInvitationByInvitationIDRepository("4ec1b06b-7ddc-11ee-9739-3822e2435ffd")

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}

func TestGetInvitationStatusByInvitationIDRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Where", "invitation_id = 1").Return(&gorm.DB{})
	maem.Mock.On("First", mock.Anything).Return(&gorm.DB{})
	_, err := mainAppEntity.GetInvitationStatusByInvitationIDRepository(1)

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}

func TestListInvitationRepository(t *testing.T) {
	// Initialize the mock
	maem.Mock.On("Table", "invitations as i").Return(&gorm.DB{})
	maem.Mock.On("Joins", "JOIN invitation_statuses as is1 ON i.id = is1.invitation_id").Return(&gorm.DB{})
	maem.Mock.On("Joins", "JOIN gatherings as g ON i.gathering_id = g.id").Return(&gorm.DB{})
	maem.Mock.On("Joins", "JOIN members as u ON i.member_id = u.id").Return(&gorm.DB{})
	maem.Mock.On("Where", "g.status = 1").Return(&gorm.DB{})
	maem.Mock.On("Select", "i.invitation_id, u.member_id, g.gathering_id, CASE WHEN is1.status = 1 THEN 'Created' WHEN is1.status = 2 THEN 'Accepted' ELSE 'Rejected' END AS invitation_status").Return(&gorm.DB{})
	maem.Mock.On("Find", mock.Anything).Return(&gorm.DB{})

	_, err := mainAppEntity.ListInvitationRepository()

	maem.Mock.AssertExpectations(t)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}

}
