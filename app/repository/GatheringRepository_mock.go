package repository

import (
	"gits-test/models/model"
	"gits-test/models/response"
)

func (maem *MainAppEntityMock) CreateGatheringRepository(gatheringModel model.Gathering) (int, error) {
	args := maem.Mock.Called(gatheringModel)

	return args.Get(0).(int), args.Error(1)
}

func (maem *MainAppEntityMock) CreateGatheringTypeRepository(gatheringTypeModel model.GatheringType) error {
	args := maem.Mock.Called(gatheringTypeModel)

	return args.Error(0)
}

func (maem *MainAppEntityMock) GetGatheringByGatheringIDRepository(gatheringID string) (model.Gathering, error) {
	args := maem.Mock.Called(gatheringID)

	return args.Get(0).(model.Gathering), args.Error(1)
}

func (maem *MainAppEntityMock) ListGatheringRepository() ([]response.ListGatheringResponse, error) {
	args := maem.Mock.Called()

	return args.Get(0).([]response.ListGatheringResponse), args.Error(1)
}
