package repository

import (
	"gits-test/models/model"
	"gits-test/models/response"
)

func (maem *MainAppEntityMock) CreateInvitationRepository(invitation model.Invitation) (int, error) {
	args := maem.Mock.Called(invitation)

	return args.Get(0).(int), args.Error(1)
}

func (maem *MainAppEntityMock) CreateInvitationStatusRepository(invitationStatus model.InvitationStatus) error {
	args := maem.Mock.Called(invitationStatus)

	return args.Error(0)
}

func (maem *MainAppEntityMock) AcceptNotAcceptRepository(invitationStatus model.InvitationStatus) error {
	args := maem.Mock.Called(invitationStatus)

	return args.Error(0)
}

func (maem *MainAppEntityMock) GetInvitationByInvitationIDRepository(invitationID string) (model.Invitation, error) {
	args := maem.Mock.Called(invitationID)

	return args.Get(0).(model.Invitation), args.Error(1)
}

func (maem *MainAppEntityMock) GetInvitationStatusByInvitationIDRepository(invitationID int) (model.InvitationStatus, error) {
	args := maem.Mock.Called(invitationID)

	return args.Get(0).(model.InvitationStatus), args.Error(1)
}

func (maem *MainAppEntityMock) ListInvitationRepository() ([]response.ListInvitationResponse, error) {
	args := maem.Mock.Called()

	return args.Get(0).([]response.ListInvitationResponse), args.Error(1)
}
