package repository

import (
	"gits-test/models/model"
	"gits-test/models/response"
)

func (maem *MainAppEntityMock) CreateAttendeeRepository(attendee model.Attendee) error {
	args := maem.Mock.Called(attendee)

	return args.Error(0)
}

func (maem *MainAppEntityMock) ListAttendeeRepository() ([]response.ListAttendeeResponse, error) {
	args := maem.Mock.Called()

	return args.Get(0).([]response.ListAttendeeResponse), args.Error(1)
}
