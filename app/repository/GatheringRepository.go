package repository

import "gits-test/models/model"

func (ar *AppRepository) CreateGatheringRepository(gatheringModel model.Gathering) (int, error) {
	if err := ar.DB.Create(&gatheringModel).Error; err != nil {
		return 0, err
	}
	return gatheringModel.ID, nil
}

func (ar *AppRepository) CreateGatheringTypeRepository(gatheringTypeModel model.GatheringType) error {
	if err := ar.DB.Create(&gatheringTypeModel).Error; err != nil {
		return err
	}
	return nil
}
