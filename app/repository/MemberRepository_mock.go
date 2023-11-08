package repository

import (
	"gits-test/models/model"
	"gits-test/models/response"
)

func (maem *MainAppEntityMock) CreateMemberRepository(member model.Member) error {
	args := maem.Mock.Called(member)

	return args.Error(0)
}

func (maem *MainAppEntityMock) GetMemberByMemberIDRepository(memberID string) (model.Member, error) {
	args := maem.Mock.Called(memberID)

	return args.Get(0).(model.Member), args.Error(1)
}

func (maem *MainAppEntityMock) ListMemberRepository() ([]response.ListMemberResponse, error) {
	args := maem.Mock.Called()

	return args.Get(0).([]response.ListMemberResponse), args.Error(1)
}
